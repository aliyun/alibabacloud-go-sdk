// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProvisionedProductPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *UpdateProvisionedProductPlanResponseBody
	GetPlanId() *string
	SetProvisionedProductId(v string) *UpdateProvisionedProductPlanResponseBody
	GetProvisionedProductId() *string
	SetRequestId(v string) *UpdateProvisionedProductPlanResponseBody
	GetRequestId() *string
}

type UpdateProvisionedProductPlanResponseBody struct {
	// The ID of the plan.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the product instance.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProvisionedProductPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateProvisionedProductPlanResponseBody) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *UpdateProvisionedProductPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProvisionedProductPlanResponseBody) SetPlanId(v string) *UpdateProvisionedProductPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponseBody) SetProvisionedProductId(v string) *UpdateProvisionedProductPlanResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponseBody) SetRequestId(v string) *UpdateProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
