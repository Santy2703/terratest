package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestAWSTerraform(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../terraform",
    }

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
}
