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
	// The sales information of cloud service configuration check.
	CheckSale *GetCheckSaleResponseBodyCheckSale `json:"CheckSale,omitempty" xml:"CheckSale,omitempty" type:"Struct"`
	// The ID of the request. The China Cloud generates a unique identifier for the request, which can be used for troubleshooting and diagnostics.
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
	if s.CheckSale != nil {
		if err := s.CheckSale.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCheckSaleResponseBodyCheckSale struct {
	// The number of consumed authorized quotas.
	//
	// example:
	//
	// 500
	ConsumeCount                                *int64 `json:"ConsumeCount,omitempty" xml:"ConsumeCount,omitempty"`
	InstanceConsumeCount                        *int64 `json:"InstanceConsumeCount,omitempty" xml:"InstanceConsumeCount,omitempty"`
	InstanceHybridPostLatestCycledResourceCount *int64 `json:"InstanceHybridPostLatestCycledResourceCount,omitempty" xml:"InstanceHybridPostLatestCycledResourceCount,omitempty"`
	InstancePostConsumeCount                    *int64 `json:"InstancePostConsumeCount,omitempty" xml:"InstancePostConsumeCount,omitempty"`
	InstancePurchaseCount                       *int64 `json:"InstancePurchaseCount,omitempty" xml:"InstancePurchaseCount,omitempty"`
	// Indicates whether the user is an existing user who used the cloud service configuration check feature before the sales feature was released (July 7, 2023). Valid values:
	//
	// - **true**: The user is an existing user.
	//
	// - **false**: The user is not an existing user.
	//
	// example:
	//
	// true
	LoyalUser *bool `json:"LoyalUser,omitempty" xml:"LoyalUser,omitempty"`
	// The number of purchased authorized quotas.
	//
	// example:
	//
	// 1000
	PurchaseCount *int64 `json:"PurchaseCount,omitempty" xml:"PurchaseCount,omitempty"`
	// The sales user type. Valid values:
	//
	// - **1**: Full-feature user. The user can use all check items.
	//
	// - **2**: Upgrade-required user. The user can use only the check items that were available before the sales feature was released (July 7, 2023).
	//
	// - **3**: Purchase-required user. The user cannot use the cloud service configuration check feature.
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

func (s *GetCheckSaleResponseBodyCheckSale) GetInstanceConsumeCount() *int64 {
	return s.InstanceConsumeCount
}

func (s *GetCheckSaleResponseBodyCheckSale) GetInstanceHybridPostLatestCycledResourceCount() *int64 {
	return s.InstanceHybridPostLatestCycledResourceCount
}

func (s *GetCheckSaleResponseBodyCheckSale) GetInstancePostConsumeCount() *int64 {
	return s.InstancePostConsumeCount
}

func (s *GetCheckSaleResponseBodyCheckSale) GetInstancePurchaseCount() *int64 {
	return s.InstancePurchaseCount
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

func (s *GetCheckSaleResponseBodyCheckSale) SetInstanceConsumeCount(v int64) *GetCheckSaleResponseBodyCheckSale {
	s.InstanceConsumeCount = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) SetInstanceHybridPostLatestCycledResourceCount(v int64) *GetCheckSaleResponseBodyCheckSale {
	s.InstanceHybridPostLatestCycledResourceCount = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) SetInstancePostConsumeCount(v int64) *GetCheckSaleResponseBodyCheckSale {
	s.InstancePostConsumeCount = &v
	return s
}

func (s *GetCheckSaleResponseBodyCheckSale) SetInstancePurchaseCount(v int64) *GetCheckSaleResponseBodyCheckSale {
	s.InstancePurchaseCount = &v
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
