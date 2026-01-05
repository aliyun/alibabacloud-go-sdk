// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionedProductPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *GetProvisionedProductPlanRequest
	GetPlanId() *string
}

type GetProvisionedProductPlanRequest struct {
	// The ID of the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s GetProvisionedProductPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *GetProvisionedProductPlanRequest) SetPlanId(v string) *GetProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

func (s *GetProvisionedProductPlanRequest) Validate() error {
	return dara.Validate(s)
}
