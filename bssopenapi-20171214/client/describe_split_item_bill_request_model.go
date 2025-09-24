// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSplitItemBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeSplitItemBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *DescribeSplitItemBillRequest
	GetBillingCycle() *string
	SetBillingDate(v string) *DescribeSplitItemBillRequest
	GetBillingDate() *string
	SetGranularity(v string) *DescribeSplitItemBillRequest
	GetGranularity() *string
	SetInstanceID(v string) *DescribeSplitItemBillRequest
	GetInstanceID() *string
	SetIsHideZeroCharge(v bool) *DescribeSplitItemBillRequest
	GetIsHideZeroCharge() *bool
	SetMaxResults(v int32) *DescribeSplitItemBillRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSplitItemBillRequest
	GetNextToken() *string
	SetOwnerId(v int64) *DescribeSplitItemBillRequest
	GetOwnerId() *int64
	SetPipCode(v string) *DescribeSplitItemBillRequest
	GetPipCode() *string
	SetProductCode(v string) *DescribeSplitItemBillRequest
	GetProductCode() *string
	SetProductType(v string) *DescribeSplitItemBillRequest
	GetProductType() *string
	SetSplitItemID(v string) *DescribeSplitItemBillRequest
	GetSplitItemID() *string
	SetSubscriptionType(v string) *DescribeSplitItemBillRequest
	GetSubscriptionType() *string
	SetTagFilter(v []*DescribeSplitItemBillRequestTagFilter) *DescribeSplitItemBillRequest
	GetTagFilter() []*DescribeSplitItemBillRequestTagFilter
}

type DescribeSplitItemBillRequest struct {
	// The ID of the member. If you specify this parameter, the bills of the member are queried. If you do not specify this parameter, the bills of the current account are queried by default.
	//
	// example:
	//
	// 123
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle. Specify the parameter in the YYYY-MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when the Granularity parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2020-03-02
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills on a monthly basis. The data that you query is the same as the data that is queried by billing cycles in the Split Bill module of the User Center console.
	//
	// 	- DAILY: queries bills on a daily basis. The data that you query is the same as the data that is queried by days in the Split Bill module of the User Center console.
	//
	// If you specify DAILY for this parameter, the BillingDate parameter is required.
	//
	// example:
	//
	// Monthly
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-kjhdskjgshfdlkjfdh
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// Specifies whether to filter bills if both the pretax gross amount and pretax amount are 0. Valid values:
	//
	// 	- false: does not filter bills.
	//
	// 	- true: filters bills.
	//
	// example:
	//
	// false
	IsHideZeroCharge *bool `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	// The maximum number of entries to query. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. The parameter must be left empty or set to the value of the NextToken parameter returned in the last call. Otherwise, an error is returned. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// CAESEgoQCg4KCmd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service. The code is the same as that in Cost Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
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
	// The ID of the split item.
	//
	// example:
	//
	// i-kjhdskjgshfdlkjfdh
	SplitItemID *string `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	// The billing method. Valid values: Subscription: the subscription billing method. PayAsYouGo: the pay-as-you-go billing method. This parameter must be used with the ProductCode parameter.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The tags that are used to filter split bills. You can specify multiple tag values. If you specify multiple tag values, split bills that match one of the tag values are queried.
	TagFilter []*DescribeSplitItemBillRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Repeated"`
}

func (s DescribeSplitItemBillRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSplitItemBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeSplitItemBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeSplitItemBillRequest) GetBillingDate() *string {
	return s.BillingDate
}

func (s *DescribeSplitItemBillRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeSplitItemBillRequest) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeSplitItemBillRequest) GetIsHideZeroCharge() *bool {
	return s.IsHideZeroCharge
}

func (s *DescribeSplitItemBillRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSplitItemBillRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSplitItemBillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSplitItemBillRequest) GetPipCode() *string {
	return s.PipCode
}

func (s *DescribeSplitItemBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeSplitItemBillRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeSplitItemBillRequest) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *DescribeSplitItemBillRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeSplitItemBillRequest) GetTagFilter() []*DescribeSplitItemBillRequestTagFilter {
	return s.TagFilter
}

func (s *DescribeSplitItemBillRequest) SetBillOwnerId(v int64) *DescribeSplitItemBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetBillingCycle(v string) *DescribeSplitItemBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetBillingDate(v string) *DescribeSplitItemBillRequest {
	s.BillingDate = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetGranularity(v string) *DescribeSplitItemBillRequest {
	s.Granularity = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetInstanceID(v string) *DescribeSplitItemBillRequest {
	s.InstanceID = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetIsHideZeroCharge(v bool) *DescribeSplitItemBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetMaxResults(v int32) *DescribeSplitItemBillRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetNextToken(v string) *DescribeSplitItemBillRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetOwnerId(v int64) *DescribeSplitItemBillRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetPipCode(v string) *DescribeSplitItemBillRequest {
	s.PipCode = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetProductCode(v string) *DescribeSplitItemBillRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetProductType(v string) *DescribeSplitItemBillRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetSplitItemID(v string) *DescribeSplitItemBillRequest {
	s.SplitItemID = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetSubscriptionType(v string) *DescribeSplitItemBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeSplitItemBillRequest) SetTagFilter(v []*DescribeSplitItemBillRequestTagFilter) *DescribeSplitItemBillRequest {
	s.TagFilter = v
	return s
}

func (s *DescribeSplitItemBillRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeSplitItemBillRequestTagFilter struct {
	// The TagFilter.N parameter is used to query bills that match a specified tag. The value of the TagFilter.N parameter must be a key-value pair. The tag key must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// 	- If only the TagFilter.N.TagKey parameter is specified, all bills associated with the tag key are queried.
	//
	// 	- If you specify multiple tag key-value pairs at the same time, bills that meet any one of the tag key-value pairs are queried.
	//
	// 	- If the tags added to resources change, you can query only the bills that are generated within the period in which the tags and resources are associated.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// You can specify the TagValues.N parameter to query bills that match the specified tag value. The value of the TagValues.N parameter must be a string. The tag value must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// 	- If you specify the TagValues.N parameter, the TagFilter.N.TagKey parameter is required. Otherwise, the error message InvalidParameter.TagValues is returned.
	//
	// 	- If you specify multiple tag values, split bills that match one of the tag values are queried.
	//
	// example:
	//
	// TestValue
	TagValues []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s DescribeSplitItemBillRequestTagFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeSplitItemBillRequestTagFilter) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillRequestTagFilter) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeSplitItemBillRequestTagFilter) GetTagValues() []*string {
	return s.TagValues
}

func (s *DescribeSplitItemBillRequestTagFilter) SetTagKey(v string) *DescribeSplitItemBillRequestTagFilter {
	s.TagKey = &v
	return s
}

func (s *DescribeSplitItemBillRequestTagFilter) SetTagValues(v []*string) *DescribeSplitItemBillRequestTagFilter {
	s.TagValues = v
	return s
}

func (s *DescribeSplitItemBillRequestTagFilter) Validate() error {
	return dara.Validate(s)
}
