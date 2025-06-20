# Terraform Provider: Hello

A minimal custom Terraform provider named `hello`, created using the Terraform Plugin SDK v2. This example demonstrates the structure and basic implementation of a custom provider and resource.

---

## 🗂️ Project Structure

```
hello-provider/
├── go.mod                            # Go module definition
├── go.sum                            # Go dependencies lock file
├── internal/
│   ├── provider/
│   │   └── provider.go               # Provider and resource implementation
│   └── provider.md                   # Optional provider documentation
├── main.go                           # Entry point for the provider binary
└── terraform/
    ├── main.tf                       # Terraform configuration using this provider
    └── terraform.tfstate             # State file (auto-generated)
```

---

## ⚙️ Setup

### 1. Initialize Go Module

```bash
go mod init github.com/j3rryCodes/terraform-provider-hello
go get github.com/hashicorp/terraform-plugin-sdk/v2
go mod tidy
```

### 2. Build the Plugin Binary

```bash
go build -o ~/.terraform.d/plugins/github.com/j3rryCodes/hello/1.0.0/linux_amd64/terraform-provider-hello
```

> 📌 **Note**: Make sure the path follows this format:  
> `~/.terraform.d/plugins/<hostname>/<namespace>/<name>/<version>/<os>_<arch>/terraform-provider-<name>`

---

## 🧪 Usage

### terraform/main.tf

```hcl
terraform {
  required_providers {
    hello = {
      source  = "github.com/j3rryCodes/hello"
      version = "1.0.0"
    }
  }
}

provider "hello" {}

resource "hello_resource" "example" {
  name = "Jerin"
}
```

### Terraform Commands

```bash
cd terraform
terraform init
terraform plan
terraform apply
terraform refresh
terraform destroy
```

---

## ✅ Provider Features

- **Provider**: `hello`
- **Resource**: `hello_resource`
  - **Argument**:
    - `name` (string) — Name to greet.

---

## 📚 Resources

- [Terraform Plugin SDK v2](https://github.com/hashicorp/terraform-plugin-sdk)
- [Custom Provider Dev Guide](https://developer.hashicorp.com/terraform/plugin/framework)

---

## 👨‍💻 Author

Created by [@j3rryCodes](https://github.com/j3rryCodes)

