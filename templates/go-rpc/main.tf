variable "environment" {}
variable "container_project" {}
variable "project_id" {}
variable "project_number" {}
variable "region" {}
variable "private_network" {}

variable "service_depends_on" {
  type    = any
  default = null
}

resource "google_cloud_run_service" "{{.ServiceName}}" {
  project  = var.project_id
  name     = "{{.ServiceName}}"
  location = var.region

  template {
    spec {
      containers {
        image = "gcr.io/${var.container_project}/services/{{.ServiceName}}:latest"

        env {
          name  = "ENVIRONMENT"
          value = var.environment
        }


        env {
          name = "SERVICE_NAME"
          value = "{{.ServiceName}}"
        }
      }
    }
  }

  autogenerate_revision_name = true

  traffic {
    percent         = 100
    latest_revision = true
  }
}

output "url" {
  value = google_cloud_run_service.{{.ServiceName}}.status[0].url
}

resource "google_cloud_run_service_iam_binding" "binding" {
  location = var.region
  project  = var.project_id
  service  = google_cloud_run_service.{{.ServiceName}}.name
  role     = "roles/run.invoker"
  members = [
    "serviceAccount:service-${var.project_number}@serverless-robot-prod.iam.gserviceaccount.com"
  ]

  depends_on = [google_cloud_run_service.{{.ServiceName}}]
}