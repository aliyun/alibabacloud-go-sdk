// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePricingModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribePricingModuleRequest
	GetOwnerId() *int64
	SetProductCode(v string) *DescribePricingModuleRequest
	GetProductCode() *string
	SetProductType(v string) *DescribePricingModuleRequest
	GetProductType() *string
	SetSubscriptionType(v string) *DescribePricingModuleRequest
	GetSubscriptionType() *string
}

type DescribePricingModuleRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service code. You can query the service code by calling the **QueryProductList*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service. You can query the service type by calling the **QueryProductList*	- operation.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: subscription
	//
	// 	- PayAsYouGo: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribePricingModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePricingModuleRequest) GoString() string {
	return s.String()
}

func (s *DescribePricingModuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePricingModuleRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePricingModuleRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribePricingModuleRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribePricingModuleRequest) SetOwnerId(v int64) *DescribePricingModuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePricingModuleRequest) SetProductCode(v string) *DescribePricingModuleRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribePricingModuleRequest) SetProductType(v string) *DescribePricingModuleRequest {
	s.ProductType = &v
	return s
}

func (s *DescribePricingModuleRequest) SetSubscriptionType(v string) *DescribePricingModuleRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribePricingModuleRequest) Validate() error {
	return dara.Validate(s)
}
