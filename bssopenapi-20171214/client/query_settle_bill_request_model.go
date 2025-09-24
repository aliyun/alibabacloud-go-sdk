// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySettleBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *QuerySettleBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *QuerySettleBillRequest
	GetBillingCycle() *string
	SetIsDisplayLocalCurrency(v bool) *QuerySettleBillRequest
	GetIsDisplayLocalCurrency() *bool
	SetIsHideZeroCharge(v bool) *QuerySettleBillRequest
	GetIsHideZeroCharge() *bool
	SetMaxResults(v int32) *QuerySettleBillRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QuerySettleBillRequest
	GetNextToken() *string
	SetOwnerId(v int64) *QuerySettleBillRequest
	GetOwnerId() *int64
	SetProductCode(v string) *QuerySettleBillRequest
	GetProductCode() *string
	SetProductType(v string) *QuerySettleBillRequest
	GetProductType() *string
	SetRecordID(v string) *QuerySettleBillRequest
	GetRecordID() *string
	SetSubscriptionType(v string) *QuerySettleBillRequest
	GetSubscriptionType() *string
	SetType(v string) *QuerySettleBillRequest
	GetType() *string
}

type QuerySettleBillRequest struct {
	// example:
	//
	// 123
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2018-07
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// example:
	//
	// false
	IsDisplayLocalCurrency *bool `json:"IsDisplayLocalCurrency,omitempty" xml:"IsDisplayLocalCurrency,omitempty"`
	// example:
	//
	// true
	IsHideZeroCharge *bool `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 12233
	RecordID *string `json:"RecordID,omitempty" xml:"RecordID,omitempty"`
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// example:
	//
	// SubscriptionOrder
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuerySettleBillRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySettleBillRequest) GoString() string {
	return s.String()
}

func (s *QuerySettleBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *QuerySettleBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QuerySettleBillRequest) GetIsDisplayLocalCurrency() *bool {
	return s.IsDisplayLocalCurrency
}

func (s *QuerySettleBillRequest) GetIsHideZeroCharge() *bool {
	return s.IsHideZeroCharge
}

func (s *QuerySettleBillRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QuerySettleBillRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QuerySettleBillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySettleBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QuerySettleBillRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QuerySettleBillRequest) GetRecordID() *string {
	return s.RecordID
}

func (s *QuerySettleBillRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QuerySettleBillRequest) GetType() *string {
	return s.Type
}

func (s *QuerySettleBillRequest) SetBillOwnerId(v int64) *QuerySettleBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QuerySettleBillRequest) SetBillingCycle(v string) *QuerySettleBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QuerySettleBillRequest) SetIsDisplayLocalCurrency(v bool) *QuerySettleBillRequest {
	s.IsDisplayLocalCurrency = &v
	return s
}

func (s *QuerySettleBillRequest) SetIsHideZeroCharge(v bool) *QuerySettleBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *QuerySettleBillRequest) SetMaxResults(v int32) *QuerySettleBillRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySettleBillRequest) SetNextToken(v string) *QuerySettleBillRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySettleBillRequest) SetOwnerId(v int64) *QuerySettleBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySettleBillRequest) SetProductCode(v string) *QuerySettleBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QuerySettleBillRequest) SetProductType(v string) *QuerySettleBillRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySettleBillRequest) SetRecordID(v string) *QuerySettleBillRequest {
	s.RecordID = &v
	return s
}

func (s *QuerySettleBillRequest) SetSubscriptionType(v string) *QuerySettleBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySettleBillRequest) SetType(v string) *QuerySettleBillRequest {
	s.Type = &v
	return s
}

func (s *QuerySettleBillRequest) Validate() error {
	return dara.Validate(s)
}
