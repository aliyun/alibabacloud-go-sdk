// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *ExecuteTerraformPlanRequest
  GetClientToken() *string 
  SetCode(v string) *ExecuteTerraformPlanRequest
  GetCode() *string 
  SetStateId(v string) *ExecuteTerraformPlanRequest
  GetStateId() *string 
}

type ExecuteTerraformPlanRequest struct {
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

func (s ExecuteTerraformPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformPlanRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformPlanRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteTerraformPlanRequest) GetCode() *string  {
  return s.Code
}

func (s *ExecuteTerraformPlanRequest) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteTerraformPlanRequest) SetClientToken(v string) *ExecuteTerraformPlanRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteTerraformPlanRequest) SetCode(v string) *ExecuteTerraformPlanRequest {
  s.Code = &v
  return s
}

func (s *ExecuteTerraformPlanRequest) SetStateId(v string) *ExecuteTerraformPlanRequest {
  s.StateId = &v
  return s
}

func (s *ExecuteTerraformPlanRequest) Validate() error {
  return dara.Validate(s)
}

