// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *GenerateModuleResponseBody
	GetModule() *string
	SetProperties(v map[string]interface{}) *GenerateModuleResponseBody
	GetProperties() map[string]interface{}
	SetRequestId(v string) *GenerateModuleResponseBody
	GetRequestId() *string
}

type GenerateModuleResponseBody struct {
	// The generated Terraform HCL template code content.
	//
	// example:
	//
	// terraform {
	//
	//   required_providers {
	//
	//     alicloud = {
	//
	//       source   = "aliyun/alicloud"
	//
	//       version  = "1.260.0"
	//
	//     }
	//
	//   }
	//
	// }
	//
	// resource "alicloud_vpc" "default" {
	//
	//  vpc_name = "vpc-test"
	//
	// }
	Module *string `json:"module,omitempty" xml:"module,omitempty"`
	// The variables and resource properties in the generated template code.
	//
	// example:
	//
	// {"vpc_name":"vpc-test"}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6B40D088-E929-504B-8802-C1759A993FA2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateModuleResponseBody) GetModule() *string {
	return s.Module
}

func (s *GenerateModuleResponseBody) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GenerateModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateModuleResponseBody) SetModule(v string) *GenerateModuleResponseBody {
	s.Module = &v
	return s
}

func (s *GenerateModuleResponseBody) SetProperties(v map[string]interface{}) *GenerateModuleResponseBody {
	s.Properties = v
	return s
}

func (s *GenerateModuleResponseBody) SetRequestId(v string) *GenerateModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
