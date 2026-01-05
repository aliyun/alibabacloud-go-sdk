// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelProvisionedProductPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *CancelProvisionedProductPlanRequest
	GetPlanId() *string
}

type CancelProvisionedProductPlanRequest struct {
	// The ID of the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s CancelProvisionedProductPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *CancelProvisionedProductPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *CancelProvisionedProductPlanRequest) SetPlanId(v string) *CancelProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

func (s *CancelProvisionedProductPlanRequest) Validate() error {
	return dara.Validate(s)
}
