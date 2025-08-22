// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformApplyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *ExecuteTerraformApplyRequest
  GetClientToken() *string 
  SetCode(v string) *ExecuteTerraformApplyRequest
  GetCode() *string 
  SetStateId(v string) *ExecuteTerraformApplyRequest
  GetStateId() *string 
}

type ExecuteTerraformApplyRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // a65451293e64979ba7a4b573950217fe
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
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
  //       version  = "1.254.0"
  // 
  //     }
  // 
  //   }
  // 
  // }
  // 
  // resource "alicloud_vpc" "default" {
  // 
  //   is_default                                  = false
  // 
  //   enable_ipv6                                 = false
  // 
  //   classic_link_enabled                        = false
  // 
  //   force_delete                                = false
  // 
  //   system_route_table_route_propagation_enable = false
  // 
  //   dry_run                                     = false
  // 
  // }
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // example:
  // 
  // task-xxx
  StateId *string `json:"stateId,omitempty" xml:"stateId,omitempty"`
}

func (s ExecuteTerraformApplyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformApplyRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformApplyRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteTerraformApplyRequest) GetCode() *string  {
  return s.Code
}

func (s *ExecuteTerraformApplyRequest) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteTerraformApplyRequest) SetClientToken(v string) *ExecuteTerraformApplyRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteTerraformApplyRequest) SetCode(v string) *ExecuteTerraformApplyRequest {
  s.Code = &v
  return s
}

func (s *ExecuteTerraformApplyRequest) SetStateId(v string) *ExecuteTerraformApplyRequest {
  s.StateId = &v
  return s
}

func (s *ExecuteTerraformApplyRequest) Validate() error {
  return dara.Validate(s)
}

