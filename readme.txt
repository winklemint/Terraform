# Custom Terraform Provider

This is a custom Terraform provider that demonstrates the basic structure and functionality of a custom provider.
It allows you to create a simple resource and execute a local command during the provisioning process.

## Prerequisites

- Go programming language (https://golang.org/doc/install)

## Installation

1. Clone this repository to your local machine:

   ```shell
   git clone <repository-url>

2. Navigate to the project directory:

     cd custom-terraform-provider

3. Compile the go code to generate the provider binary

     go build

4. After the compilation is successful, move the generated binary to a directory that is included in your system's PATH.

# Usage

1. Create a new Terraform configuration file, e.g., main.tf, and define your configuration using the custom provider:

     provider "myprovider" {
     # Provider-specific configuration options
     }

     resource "myprovider_resource" "example" {
     # Resource-specific configuration options
     }

**Note: replace the "myprovider" witht the actual name of your provider

2. Save the main.tf file

3.Open a command prompt or PowerShell window and navigate to the directory where your main.tf file is located.

4.Initialize the Terraform configuration by running:

     terraform init

5. Apply the Terraform configuration to create the resources:

     terraform apply

Terraform will automatically load and use your custom provider based on the configuration in main.tf.

6.Interact with the resources created by your custom provider, modify the configuration, and test different scenarios using various Terraform commands such as terraform plan, terraform apply, terraform destroy, etc.

7.Clean up resources when you're done testing by running:

     terraform destroy

# Customization

To customize the provider for your specific use case, you can modify the main.go file and add more resources, attributes, and functionality based on your requirements.
Refer to the Terraform Plugin SDK documentation for more details on how to implement resource CRUD operations and other provider features.