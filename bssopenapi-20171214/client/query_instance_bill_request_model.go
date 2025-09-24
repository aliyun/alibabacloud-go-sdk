// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *QueryInstanceBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *QueryInstanceBillRequest
	GetBillingCycle() *string
	SetBillingDate(v string) *QueryInstanceBillRequest
	GetBillingDate() *string
	SetGranularity(v string) *QueryInstanceBillRequest
	GetGranularity() *string
	SetIsBillingItem(v bool) *QueryInstanceBillRequest
	GetIsBillingItem() *bool
	SetIsHideZeroCharge(v bool) *QueryInstanceBillRequest
	GetIsHideZeroCharge() *bool
	SetOwnerId(v int64) *QueryInstanceBillRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QueryInstanceBillRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryInstanceBillRequest
	GetPageSize() *int32
	SetProductCode(v string) *QueryInstanceBillRequest
	GetProductCode() *string
	SetProductType(v string) *QueryInstanceBillRequest
	GetProductType() *string
	SetSubscriptionType(v string) *QueryInstanceBillRequest
	GetSubscriptionType() *string
}

type QueryInstanceBillRequest struct {
	// The ID of the member. If you specify a value for this parameter, you can query the bills of the specified member. If you leave this parameter empty, the bills of the current account are queried by default.
	//
	// example:
	//
	// 122
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2020-03-03
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills by month. The data queried is consistent with the data that is displayed for the specified billing cycle on the Billing Details tab of the Bill Details page in User Center.
	//
	// 	- DAILY: queries bills by day. The data queried is consistent with the data that is displayed for the specified day on the Billing Details tab of the Bill Details page in User Center.
	//
	// You must set the **BillingDate*	- parameter before you can set the Granularity parameter to DAILY.
	//
	// example:
	//
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// Specifies whether to query data by billable item. Valid values:
	//
	// 	- false: does not query data by billable item. The data queried is consistent with the data that is displayed for the specified instance on the Billing Details tab of the Bill Details page in User Center.
	//
	// 	- true: queries data by billable item. The data queried is consistent with the data that is displayed for the specified billable item on the Billing Details tab of the Bill Details page in User Center.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsBillingItem *bool `json:"IsBillingItem,omitempty" xml:"IsBillingItem,omitempty"`
	// Specifies whether to filter out a bill whose pretax gross amount and pretax amount are 0. Default value: false.*******	- Valid values:
	//
	// 	- false: does not filter the bill.
	//
	// 	- true: filters the bill.
	//
	// example:
	//
	// false
	IsHideZeroCharge *bool  `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	OwnerId          *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the service. This parameter is required if the ProductCode parameter is set to the service code of Alibaba Cloud Marketplace.
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
	// **
	//
	// ****This parameter must be used together with the **ProductCode*	- parameter.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryInstanceBillRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *QueryInstanceBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryInstanceBillRequest) GetBillingDate() *string {
	return s.BillingDate
}

func (s *QueryInstanceBillRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryInstanceBillRequest) GetIsBillingItem() *bool {
	return s.IsBillingItem
}

func (s *QueryInstanceBillRequest) GetIsHideZeroCharge() *bool {
	return s.IsHideZeroCharge
}

func (s *QueryInstanceBillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryInstanceBillRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryInstanceBillRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInstanceBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryInstanceBillRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryInstanceBillRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryInstanceBillRequest) SetBillOwnerId(v int64) *QueryInstanceBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QueryInstanceBillRequest) SetBillingCycle(v string) *QueryInstanceBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceBillRequest) SetBillingDate(v string) *QueryInstanceBillRequest {
	s.BillingDate = &v
	return s
}

func (s *QueryInstanceBillRequest) SetGranularity(v string) *QueryInstanceBillRequest {
	s.Granularity = &v
	return s
}

func (s *QueryInstanceBillRequest) SetIsBillingItem(v bool) *QueryInstanceBillRequest {
	s.IsBillingItem = &v
	return s
}

func (s *QueryInstanceBillRequest) SetIsHideZeroCharge(v bool) *QueryInstanceBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *QueryInstanceBillRequest) SetOwnerId(v int64) *QueryInstanceBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryInstanceBillRequest) SetPageNum(v int32) *QueryInstanceBillRequest {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceBillRequest) SetPageSize(v int32) *QueryInstanceBillRequest {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceBillRequest) SetProductCode(v string) *QueryInstanceBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceBillRequest) SetProductType(v string) *QueryInstanceBillRequest {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceBillRequest) SetSubscriptionType(v string) *QueryInstanceBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceBillRequest) Validate() error {
	return dara.Validate(s)
}
