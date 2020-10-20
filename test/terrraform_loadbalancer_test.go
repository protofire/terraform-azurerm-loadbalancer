package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	testStructure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformLoadBalancer(t *testing.T) {
	t.Parallel()

	fixtureFolder := "./fixture"

	// Deploy the example
	testStructure.RunTestStage(t, "setup", func() {
		terraformOptions := configureTerraformOptions(fixtureFolder)

		// Save the options so later test stages can use them
		testStructure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		// This will init and apply the resources and fail the test if there are any errors
		terraform.InitAndApply(t, terraformOptions)
	})

	// Check whether the length of output meets the requirement. In public case, we check whether there occurs a public IP.
	testStructure.RunTestStage(t, "validate", func() {
		terraformOptions := testStructure.LoadTerraformOptions(t, fixtureFolder)

		publicIP := terraform.Output(t, terraformOptions, "public_ip")
		if len(publicIP) <= 0 {
			t.Fatal("Wrong output")
		}
	})

	// At the end of the test, clean up any resources that were created
	testStructure.RunTestStage(t, "teardown", func() {
		terraformOptions := testStructure.LoadTerraformOptions(t, fixtureFolder)
		terraform.Destroy(t, terraformOptions)
	})

}

func configureTerraformOptions(fixtureFolder string) *terraform.Options {

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: fixtureFolder,

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{},
	}

	return terraformOptions
}
