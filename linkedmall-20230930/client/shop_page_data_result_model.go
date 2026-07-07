// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShopPageDataResult interface {
	dara.Model
	String() string
	GoString() string
	SetCooperationShops(v []*CooperationShop) *ShopPageDataResult
	GetCooperationShops() []*CooperationShop
	SetEndDate(v string) *ShopPageDataResult
	GetEndDate() *string
	SetPurchaserId(v string) *ShopPageDataResult
	GetPurchaserId() *string
	SetShopId(v string) *ShopPageDataResult
	GetShopId() *string
	SetShopName(v string) *ShopPageDataResult
	GetShopName() *string
	SetShopType(v string) *ShopPageDataResult
	GetShopType() *string
	SetStartDate(v string) *ShopPageDataResult
	GetStartDate() *string
	SetStatus(v string) *ShopPageDataResult
	GetStatus() *string
}

type ShopPageDataResult struct {
	// Partner shops
	//
	// example:
	//
	// 12****01
	CooperationShops []*CooperationShop `json:"cooperationShops,omitempty" xml:"cooperationShops,omitempty" type:"Repeated"`
	// End time
	//
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Purchaser ID
	//
	// example:
	//
	// PID56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
	// Shop ID
	//
	// example:
	//
	// 22****09
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// Shop name
	//
	// example:
	//
	// 儿童座椅分销店铺
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
	// Shop type
	//
	// example:
	//
	// DistributorQYG
	ShopType *string `json:"shopType,omitempty" xml:"shopType,omitempty"`
	// Start time
	//
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Shop status
	//
	// example:
	//
	// Working
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ShopPageDataResult) String() string {
	return dara.Prettify(s)
}

func (s ShopPageDataResult) GoString() string {
	return s.String()
}

func (s *ShopPageDataResult) GetCooperationShops() []*CooperationShop {
	return s.CooperationShops
}

func (s *ShopPageDataResult) GetEndDate() *string {
	return s.EndDate
}

func (s *ShopPageDataResult) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *ShopPageDataResult) GetShopId() *string {
	return s.ShopId
}

func (s *ShopPageDataResult) GetShopName() *string {
	return s.ShopName
}

func (s *ShopPageDataResult) GetShopType() *string {
	return s.ShopType
}

func (s *ShopPageDataResult) GetStartDate() *string {
	return s.StartDate
}

func (s *ShopPageDataResult) GetStatus() *string {
	return s.Status
}

func (s *ShopPageDataResult) SetCooperationShops(v []*CooperationShop) *ShopPageDataResult {
	s.CooperationShops = v
	return s
}

func (s *ShopPageDataResult) SetEndDate(v string) *ShopPageDataResult {
	s.EndDate = &v
	return s
}

func (s *ShopPageDataResult) SetPurchaserId(v string) *ShopPageDataResult {
	s.PurchaserId = &v
	return s
}

func (s *ShopPageDataResult) SetShopId(v string) *ShopPageDataResult {
	s.ShopId = &v
	return s
}

func (s *ShopPageDataResult) SetShopName(v string) *ShopPageDataResult {
	s.ShopName = &v
	return s
}

func (s *ShopPageDataResult) SetShopType(v string) *ShopPageDataResult {
	s.ShopType = &v
	return s
}

func (s *ShopPageDataResult) SetStartDate(v string) *ShopPageDataResult {
	s.StartDate = &v
	return s
}

func (s *ShopPageDataResult) SetStatus(v string) *ShopPageDataResult {
	s.Status = &v
	return s
}

func (s *ShopPageDataResult) Validate() error {
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
