// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *DeleteEnsSaleControlRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *DeleteEnsSaleControlRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DeleteEnsSaleControlRequest
	GetCustomAccount() *string
	SetSaleControls(v []*DeleteEnsSaleControlRequestSaleControls) *DeleteEnsSaleControlRequest
	GetSaleControls() []*DeleteEnsSaleControlRequestSaleControls
}

type DeleteEnsSaleControlRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControls []*DeleteEnsSaleControlRequestSaleControls `json:"SaleControls,omitempty" xml:"SaleControls,omitempty" type:"Repeated"`
}

func (s DeleteEnsSaleControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleControlRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *DeleteEnsSaleControlRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DeleteEnsSaleControlRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DeleteEnsSaleControlRequest) GetSaleControls() []*DeleteEnsSaleControlRequestSaleControls {
	return s.SaleControls
}

func (s *DeleteEnsSaleControlRequest) SetAliUidAccount(v string) *DeleteEnsSaleControlRequest {
	s.AliUidAccount = &v
	return s
}

func (s *DeleteEnsSaleControlRequest) SetCommodityCode(v string) *DeleteEnsSaleControlRequest {
	s.CommodityCode = &v
	return s
}

func (s *DeleteEnsSaleControlRequest) SetCustomAccount(v string) *DeleteEnsSaleControlRequest {
	s.CustomAccount = &v
	return s
}

func (s *DeleteEnsSaleControlRequest) SetSaleControls(v []*DeleteEnsSaleControlRequestSaleControls) *DeleteEnsSaleControlRequest {
	s.SaleControls = v
	return s
}

func (s *DeleteEnsSaleControlRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteEnsSaleControlRequestSaleControls struct {
	// This parameter is required.
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// This parameter is required.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DeleteEnsSaleControlRequestSaleControls) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleControlRequestSaleControls) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleControlRequestSaleControls) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DeleteEnsSaleControlRequestSaleControls) GetOrderType() *string {
	return s.OrderType
}

func (s *DeleteEnsSaleControlRequestSaleControls) SetModuleCode(v string) *DeleteEnsSaleControlRequestSaleControls {
	s.ModuleCode = &v
	return s
}

func (s *DeleteEnsSaleControlRequestSaleControls) SetOrderType(v string) *DeleteEnsSaleControlRequestSaleControls {
	s.OrderType = &v
	return s
}

func (s *DeleteEnsSaleControlRequestSaleControls) Validate() error {
	return dara.Validate(s)
}
