// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShop interface {
	dara.Model
	String() string
	GoString() string
	SetCooperationShops(v []*CooperationShop) *Shop
	GetCooperationShops() []*CooperationShop
	SetDistributorId(v string) *Shop
	GetDistributorId() *string
	SetEndDate(v string) *Shop
	GetEndDate() *string
	SetPurchaserId(v string) *Shop
	GetPurchaserId() *string
	SetRequestId(v string) *Shop
	GetRequestId() *string
	SetShopId(v string) *Shop
	GetShopId() *string
	SetShopName(v string) *Shop
	GetShopName() *string
	SetShopType(v string) *Shop
	GetShopType() *string
	SetStartDate(v string) *Shop
	GetStartDate() *string
	SetStatus(v string) *Shop
	GetStatus() *string
}

type Shop struct {
	// The partner shops.
	//
	// example:
	//
	// [{"shopId":"22000019","cooperationCompanyId":"11200003","cooperationShopId":"21000026"}]
	CooperationShops []*CooperationShop `json:"cooperationShops,omitempty" xml:"cooperationShops,omitempty" type:"Repeated"`
	// The distributor ID.
	//
	// example:
	//
	// 12****09
	DistributorId *string `json:"distributorId,omitempty" xml:"distributorId,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The purchaser ID.
	//
	// example:
	//
	// PID56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 48A34399-A72C-1E23-8388-7E63****E927
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The shop ID.
	//
	// example:
	//
	// 22****09
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// The shop name.
	//
	// example:
	//
	// Children\\"s distribution shop
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
	// The shop type.
	//
	// example:
	//
	// DistributorQYG
	ShopType *string `json:"shopType,omitempty" xml:"shopType,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The shop status.
	//
	// example:
	//
	// Working
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Shop) String() string {
	return dara.Prettify(s)
}

func (s Shop) GoString() string {
	return s.String()
}

func (s *Shop) GetCooperationShops() []*CooperationShop {
	return s.CooperationShops
}

func (s *Shop) GetDistributorId() *string {
	return s.DistributorId
}

func (s *Shop) GetEndDate() *string {
	return s.EndDate
}

func (s *Shop) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *Shop) GetRequestId() *string {
	return s.RequestId
}

func (s *Shop) GetShopId() *string {
	return s.ShopId
}

func (s *Shop) GetShopName() *string {
	return s.ShopName
}

func (s *Shop) GetShopType() *string {
	return s.ShopType
}

func (s *Shop) GetStartDate() *string {
	return s.StartDate
}

func (s *Shop) GetStatus() *string {
	return s.Status
}

func (s *Shop) SetCooperationShops(v []*CooperationShop) *Shop {
	s.CooperationShops = v
	return s
}

func (s *Shop) SetDistributorId(v string) *Shop {
	s.DistributorId = &v
	return s
}

func (s *Shop) SetEndDate(v string) *Shop {
	s.EndDate = &v
	return s
}

func (s *Shop) SetPurchaserId(v string) *Shop {
	s.PurchaserId = &v
	return s
}

func (s *Shop) SetRequestId(v string) *Shop {
	s.RequestId = &v
	return s
}

func (s *Shop) SetShopId(v string) *Shop {
	s.ShopId = &v
	return s
}

func (s *Shop) SetShopName(v string) *Shop {
	s.ShopName = &v
	return s
}

func (s *Shop) SetShopType(v string) *Shop {
	s.ShopType = &v
	return s
}

func (s *Shop) SetStartDate(v string) *Shop {
	s.StartDate = &v
	return s
}

func (s *Shop) SetStatus(v string) *Shop {
	s.Status = &v
	return s
}

func (s *Shop) Validate() error {
	if s.CooperationShops != nil {
		for _, item := range s.CooperationShops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
