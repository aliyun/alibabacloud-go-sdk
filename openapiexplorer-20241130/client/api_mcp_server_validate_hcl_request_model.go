// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiMcpServerValidateHclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApiMcpServerValidateHclRequest
	GetCode() *string
}

type ApiMcpServerValidateHclRequest struct {
	// example:
	//
	// variable "name" {
	//
	//   default = "terraform-example"
	//
	// }
	//
	// provider "alicloud" {
	//
	//   region = "cn-beijing"
	//
	// }
	//
	// resource "alicloud_vpc" "default" {
	//
	//   ipv6_isp    = "BGP"
	//
	//   description = "test"
	//
	//   cidr_block  = "10.0.0.0/8"
	//
	//   vpc_name    = var.name
	//
	//   enable_ipv6 = true
	//
	// }
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s ApiMcpServerValidateHclRequest) String() string {
	return dara.Prettify(s)
}

func (s ApiMcpServerValidateHclRequest) GoString() string {
	return s.String()
}

func (s *ApiMcpServerValidateHclRequest) GetCode() *string {
	return s.Code
}

func (s *ApiMcpServerValidateHclRequest) SetCode(v string) *ApiMcpServerValidateHclRequest {
	s.Code = &v
	return s
}

func (s *ApiMcpServerValidateHclRequest) Validate() error {
	return dara.Validate(s)
}
