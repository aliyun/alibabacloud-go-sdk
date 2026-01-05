// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProvisionedProductPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *DeleteProvisionedProductPlanRequest
	GetPlanId() *string
}

type DeleteProvisionedProductPlanRequest struct {
	// The ID of the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteProvisionedProductPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteProvisionedProductPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteProvisionedProductPlanRequest) SetPlanId(v string) *DeleteProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteProvisionedProductPlanRequest) Validate() error {
	return dara.Validate(s)
}
