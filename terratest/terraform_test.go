package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/random"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the simple Terraform module in examples/terraform-basic-example using Terratest.
// Make sure you have the dep binary, https://github.com/golang/dep
// Run 'dep ensure' before run test cases.

func TestTerraformBasicExampleNew(t *testing.T) {
	t.Parallel()

	description := fmt.Sprintf("tf-test-%d", random.Random(100, 1000))

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../example/",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"description": description,
		},

		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: false,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables
	thisDescription := terraform.OutputMap(t, terraformOptions, "this_adb_description")

	// Verify we're getting back the outputs we expect
	assert.Equal(t, thisDescription, description)
}
