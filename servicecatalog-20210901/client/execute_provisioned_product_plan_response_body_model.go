// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteProvisionedProductPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetPlanId(v string) *ExecuteProvisionedProductPlanResponseBody
  GetPlanId() *string 
  SetRequestId(v string) *ExecuteProvisionedProductPlanResponseBody
  GetRequestId() *string 
}

type ExecuteProvisionedProductPlanResponseBody struct {
  // The ID of the plan.
  // 
  // example:
  // 
  // plan-bp1jvmdk2k****
  PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteProvisionedProductPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteProvisionedProductPlanResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteProvisionedProductPlanResponseBody) GetPlanId() *string  {
  return s.PlanId
}

func (s *ExecuteProvisionedProductPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteProvisionedProductPlanResponseBody) SetPlanId(v string) *ExecuteProvisionedProductPlanResponseBody {
  s.PlanId = &v
  return s
}

func (s *ExecuteProvisionedProductPlanResponseBody) SetRequestId(v string) *ExecuteProvisionedProductPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteProvisionedProductPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

