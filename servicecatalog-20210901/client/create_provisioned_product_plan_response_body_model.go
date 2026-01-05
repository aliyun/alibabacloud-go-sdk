// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProvisionedProductPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *CreateProvisionedProductPlanResponseBody
	GetPlanId() *string
	SetProvisionedProductId(v string) *CreateProvisionedProductPlanResponseBody
	GetProvisionedProductId() *string
	SetRequestId(v string) *CreateProvisionedProductPlanResponseBody
	GetRequestId() *string
}

type CreateProvisionedProductPlanResponseBody struct {
	// The plan ID.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The product instance ID.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProvisionedProductPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateProvisionedProductPlanResponseBody) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *CreateProvisionedProductPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProvisionedProductPlanResponseBody) SetPlanId(v string) *CreateProvisionedProductPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateProvisionedProductPlanResponseBody) SetProvisionedProductId(v string) *CreateProvisionedProductPlanResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *CreateProvisionedProductPlanResponseBody) SetRequestId(v string) *CreateProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProvisionedProductPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
