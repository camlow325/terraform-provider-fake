provider "fake" {}

variable "param" {
  default = ""
}

resource "fake_foo" "test" {
  name = "tester"
  parameters = {
    TheParam = var.param
  }
}
