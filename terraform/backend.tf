terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "Terraform_workspace"

    workspaces {
      name = "Terratest"
    }
  }
}
