variable "description" {
  description = "The name of a new description."
  type        = string
  default     = "Test new adb."
}

data "alicloud_zones" "default" {
  available_resource_creation = "ADB"
  enable_details              = true
}

resource "alicloud_vpc" "default" {
  vpc_name   = "terraform_test"
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = "terraform_test"
  cidr_block   = "172.16.0.0/16"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_zones.default.zones.0.id
}

module "example" {
  source            = "../"
  description       = var.description
  vswitch_id        = alicloud_vswitch.default.id
  availability_zone = data.alicloud_zones.default.zones.0.id
}
