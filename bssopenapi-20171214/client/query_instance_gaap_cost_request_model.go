// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceGaapCostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingCycle(v string) *QueryInstanceGaapCostRequest
	GetBillingCycle() *string
	SetPageNum(v int32) *QueryInstanceGaapCostRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryInstanceGaapCostRequest
	GetPageSize() *int32
	SetProductCode(v string) *QueryInstanceGaapCostRequest
	GetProductCode() *string
	SetProductType(v string) *QueryInstanceGaapCostRequest
	GetProductType() *string
	SetSubscriptionType(v string) *QueryInstanceGaapCostRequest
	GetSubscriptionType() *string
}

type QueryInstanceGaapCostRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QueryInstanceGaapCostRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceGaapCostRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryInstanceGaapCostRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryInstanceGaapCostRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInstanceGaapCostRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryInstanceGaapCostRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryInstanceGaapCostRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryInstanceGaapCostRequest) SetBillingCycle(v string) *QueryInstanceGaapCostRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetPageNum(v int32) *QueryInstanceGaapCostRequest {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetPageSize(v int32) *QueryInstanceGaapCostRequest {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetProductCode(v string) *QueryInstanceGaapCostRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetProductType(v string) *QueryInstanceGaapCostRequest {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) SetSubscriptionType(v string) *QueryInstanceGaapCostRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceGaapCostRequest) Validate() error {
	return dara.Validate(s)
}
