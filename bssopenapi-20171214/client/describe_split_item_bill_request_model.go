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
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 123
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when Granularity is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2020-03-02
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// - MONTHLY: monthly. This value corresponds to the monthly bills on the Split Bill page in the User Center.
	//
	// - DAILY: daily. This value corresponds to the daily bills on the Split Bill page in the User Center.
	//
	// > If you set this parameter to DAILY, you must specify the BillingDate parameter.
	//
	// example:
	//
	// Monthly
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-kjhdskjgshfdlkjfdh
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// Specifies whether to filter out bills where both the original price (PretaxGrossAmount) and the payable amount (PretaxAmount) are 0. Valid values:
	//
	// - false: does not filter out the bills.
	//
	// - true: filters out the bills.
	//
	// example:
	//
	// false
	IsHideZeroCharge *bool `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	// The maximum number of entries to return in the current call. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position from which the current call starts to read data. The value must be empty or the value of NextToken returned in the previous response. Otherwise, an error occurs. An empty value indicates that the read operation starts from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCmd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The product code. This value is the same as the product code in the bills of the User Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The product code.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the attached resource.
	//
	// example:
	//
	// i-kjhdskjgshfdlkjfdh
	SplitItemID *string `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	// The subscription type. Valid values:
	//
	// - Subscription: subscription.
	//
	// - PayAsYouGo: pay-as-you-go.
	//
	// > This parameter must be used together with the ProductCode parameter.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The tag filter conditions. Multiple sets of TagValues query conditions are supported. The logical relationship between multiple sets of TagValues conditions is OR.
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
	if s.TagFilter != nil {
		for _, item := range s.TagFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSplitItemBillRequestTagFilter struct {
	// TagFilter.N is used to query bills that match the specified tags. It consists of a key-value pair. The tag key must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// - If you specify only TagFilter.N.TagKey, all bills associated with the tag key are returned.
	//
	// - If you specify multiple tag key-value pairs, bills that match any of the tag key-value pairs are returned.
	//
	// - If the tags on a resource change, you can only query the bills generated during the period when the tags were attached to the resource.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// TagValues.N is used to query bills that match the specified tag values. It consists of a list of strings. The tag value must be 1 to 128 characters in length. Valid values of N: 1 to 20.
	//
	// - If you specify TagValues.N, you must also specify TagFilter.N.TagKey. Otherwise, the InvalidParameter.TagValues error is reported.
	//
	// - If you specify multiple tag values, bills that match any of the tag values are returned.
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
