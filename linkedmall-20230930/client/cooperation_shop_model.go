// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperationShop interface {
	dara.Model
	String() string
	GoString() string
	SetCooperationCompanyId(v string) *CooperationShop
	GetCooperationCompanyId() *string
	SetCooperationShopId(v string) *CooperationShop
	GetCooperationShopId() *string
	SetShopId(v string) *CooperationShop
	GetShopId() *string
}

type CooperationShop struct {
	CooperationCompanyId *string `json:"cooperationCompanyId,omitempty" xml:"cooperationCompanyId,omitempty"`
	CooperationShopId    *string `json:"cooperationShopId,omitempty" xml:"cooperationShopId,omitempty"`
	ShopId               *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
}

func (s CooperationShop) String() string {
	return dara.Prettify(s)
}

func (s CooperationShop) GoString() string {
	return s.String()
}

func (s *CooperationShop) GetCooperationCompanyId() *string {
	return s.CooperationCompanyId
}

func (s *CooperationShop) GetCooperationShopId() *string {
	return s.CooperationShopId
}

func (s *CooperationShop) GetShopId() *string {
	return s.ShopId
}

func (s *CooperationShop) SetCooperationCompanyId(v string) *CooperationShop {
	s.CooperationCompanyId = &v
	return s
}

func (s *CooperationShop) SetCooperationShopId(v string) *CooperationShop {
	s.CooperationShopId = &v
	return s
}

func (s *CooperationShop) SetShopId(v string) *CooperationShop {
	s.ShopId = &v
	return s
}

func (s *CooperationShop) Validate() error {
	return dara.Validate(s)
}
