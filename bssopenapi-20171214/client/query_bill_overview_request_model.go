// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *QueryBillOverviewRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *QueryBillOverviewRequest
	GetBillingCycle() *string
	SetProductCode(v string) *QueryBillOverviewRequest
	GetProductCode() *string
	SetProductType(v string) *QueryBillOverviewRequest
	GetProductType() *string
	SetSubscriptionType(v string) *QueryBillOverviewRequest
	GetSubscriptionType() *string
}

type QueryBillOverviewRequest struct {
	// The ID of the member. If you specify a value for this parameter, you can query the bills of the specified member. If you leave this parameter empty, the bills of the current account are queried by default.
	//
	// example:
	//
	// 1234
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle, in the YYYY-MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-07
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryBillOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBillOverviewRequest) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *QueryBillOverviewRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryBillOverviewRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryBillOverviewRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryBillOverviewRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryBillOverviewRequest) SetBillOwnerId(v int64) *QueryBillOverviewRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QueryBillOverviewRequest) SetBillingCycle(v string) *QueryBillOverviewRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillOverviewRequest) SetProductCode(v string) *QueryBillOverviewRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryBillOverviewRequest) SetProductType(v string) *QueryBillOverviewRequest {
	s.ProductType = &v
	return s
}

func (s *QueryBillOverviewRequest) SetSubscriptionType(v string) *QueryBillOverviewRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillOverviewRequest) Validate() error {
	return dara.Validate(s)
}
