// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckSaleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckSale(v *GetCheckSaleResponseBodyCheckSale) *GetCheckSaleResponseBody
	GetCheckSale() *GetCheckSaleResponseBodyCheckSale
	SetRequestId(v string) *GetCheckSaleResponseBody
	GetRequestId() *string
}

type GetCheckSaleResponseBody struct {
	// The sales information about the configuration assessment quota.
	CheckSale *GetCheckSaleResponseBodyCheckSale `json:"CheckSale,omitempty" xml:"CheckSale,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCheckSaleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSaleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckSaleResponseBody) GetCheckSale() *GetCheckSaleResponseBodyCheckSale {
	return s.CheckSale
}

func (s *GetCheckSaleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckSaleResponseBody) SetCheckSale(v *GetCheckSaleResponseBodyCheckSale) *GetCheckSaleResponseBody {
	s.CheckSale = v
	return s
}

func (s *GetCheckSaleResponseBody) SetRequestId(v string) *GetCheckSaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckSaleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCheckSaleResponseBodyCheckSale struct {
	// The consumed quota.
	//
	// example:
	//
	// 500
	ConsumeCount *int64 `json:"ConsumeCount,omitempty" xml:"ConsumeCount,omitempty"`
	// Indicates whether the user is an existing user and whether the user uses the configuration assessment feature before the feature is released for sale on July 07, 2023. Valid values:
	//
	// 	- **true**: existing user
	//
	// 	- **false**: new user
	//
	// example:
	//
	// true
	LoyalUser *bool `json:"LoyalUser,omitempty" xml:"LoyalUser,omitempty"`
	// The purchased quota.
	//
	// example:
	//
	// 1000
	PurchaseCount *int64 `json:"PurchaseCount,omitempty" xml:"PurchaseCount,omitempty"`
	// The type of the user. Valid values:
	//
	// 	- **1**: a user who can use all check items.
	//
	// 	- **2**: an user who can only use the check items before the release of the feature on July 07, 2023. This type of users must upgrade Security Center before the users can use all check items.
	//
	// 	- **3**: a new user who cannot use the configuration assessment feature. This type of users must make a purchase before the users can use the feature.
	//
	// example:
	//
	// 1
	SaleUserType *int32 `json:"SaleUserType,omitempty" xml:"SaleUserType,omitempty"`
}

func (s GetCheckSaleResponseBodyCheckSale) String() string {
	return dara.Prettify(s)
}

func (s GetCheckSaleResponseBodyCheckSale) GoString() string {
	return s.String()
}

func (s *GetCheckSaleResponseBodyCheckSale) GetConsumeCount() *int64 {
	return s.ConsumeCount
}

func (s *GetCheckSaleResponseBodyCheckSale) GetLoyalUser() *bool {
	return s.LoyalUser
}

func (s *GetCheckSaleResponseBodyCheckSale) GetPurchaseCount() *int64 {
	return s.PurchaseCount
}

func (s *GetCheckSaleResponseBodyCheckSale) GetSaleUserType() *int32 {
	return s.SaleUserType
}

func (s *GetCheckSaleResponseBodyCheckSale) SetConsumeCount(v int64) *GetCheckSaleResponseBodyCheckSale {
	s.ConsumeCount = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) SetLoyalUser(v bool) *GetCheckSaleResponseBodyCheckSale {
	s.LoyalUser = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) SetPurchaseCount(v int64) *GetCheckSaleResponseBodyCheckSale {
	s.PurchaseCount = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) SetSaleUserType(v int32) *GetCheckSaleResponseBodyCheckSale {
	s.SaleUserType = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) Validate() error {
	return dara.Validate(s)
}
