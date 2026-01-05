// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteProvisionedProductPlanRequest interface {
  dara.Model
  String() string
  GoString() string
  SetPlanId(v string) *ExecuteProvisionedProductPlanRequest
  GetPlanId() *string 
}

type ExecuteProvisionedProductPlanRequest struct {
  // The ID of the plan.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // plan-bp1jvmdk2k****
  PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s ExecuteProvisionedProductPlanRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteProvisionedProductPlanRequest) GoString() string {
  return s.String()
}

func (s *ExecuteProvisionedProductPlanRequest) GetPlanId() *string  {
  return s.PlanId
}

func (s *ExecuteProvisionedProductPlanRequest) SetPlanId(v string) *ExecuteProvisionedProductPlanRequest {
  s.PlanId = &v
  return s
}

func (s *ExecuteProvisionedProductPlanRequest) Validate() error {
  return dara.Validate(s)
}

