// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleConditionControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *DeleteEnsSaleConditionControlRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *DeleteEnsSaleConditionControlRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DeleteEnsSaleConditionControlRequest
	GetCustomAccount() *string
	SetSaleControls(v []*DeleteEnsSaleConditionControlRequestSaleControls) *DeleteEnsSaleConditionControlRequest
	GetSaleControls() []*DeleteEnsSaleConditionControlRequestSaleControls
}

type DeleteEnsSaleConditionControlRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControls []*DeleteEnsSaleConditionControlRequestSaleControls `json:"SaleControls,omitempty" xml:"SaleControls,omitempty" type:"Repeated"`
}

func (s DeleteEnsSaleConditionControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleConditionControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleConditionControlRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *DeleteEnsSaleConditionControlRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DeleteEnsSaleConditionControlRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DeleteEnsSaleConditionControlRequest) GetSaleControls() []*DeleteEnsSaleConditionControlRequestSaleControls {
	return s.SaleControls
}

func (s *DeleteEnsSaleConditionControlRequest) SetAliUidAccount(v string) *DeleteEnsSaleConditionControlRequest {
	s.AliUidAccount = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequest) SetCommodityCode(v string) *DeleteEnsSaleConditionControlRequest {
	s.CommodityCode = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequest) SetCustomAccount(v string) *DeleteEnsSaleConditionControlRequest {
	s.CustomAccount = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequest) SetSaleControls(v []*DeleteEnsSaleConditionControlRequestSaleControls) *DeleteEnsSaleConditionControlRequest {
	s.SaleControls = v
	return s
}

func (s *DeleteEnsSaleConditionControlRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteEnsSaleConditionControlRequestSaleControls struct {
	// This parameter is required.
	ConditionControls []*DeleteEnsSaleConditionControlRequestSaleControlsConditionControls `json:"ConditionControls,omitempty" xml:"ConditionControls,omitempty" type:"Repeated"`
	// This parameter is required.
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// This parameter is required.
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DeleteEnsSaleConditionControlRequestSaleControls) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleConditionControlRequestSaleControls) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) GetConditionControls() []*DeleteEnsSaleConditionControlRequestSaleControlsConditionControls {
	return s.ConditionControls
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) GetOrderType() *string {
	return s.OrderType
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) SetConditionControls(v []*DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) *DeleteEnsSaleConditionControlRequestSaleControls {
	s.ConditionControls = v
	return s
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) SetModuleCode(v string) *DeleteEnsSaleConditionControlRequestSaleControls {
	s.ModuleCode = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) SetOrderType(v string) *DeleteEnsSaleConditionControlRequestSaleControls {
	s.OrderType = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequestSaleControls) Validate() error {
	return dara.Validate(s)
}

type DeleteEnsSaleConditionControlRequestSaleControlsConditionControls struct {
	// This parameter is required.
	ConditionControlModuleCode *string `json:"ConditionControlModuleCode,omitempty" xml:"ConditionControlModuleCode,omitempty"`
	// This parameter is required.
	ConditionControlModuleValue *string `json:"ConditionControlModuleValue,omitempty" xml:"ConditionControlModuleValue,omitempty"`
}

func (s DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) GetConditionControlModuleCode() *string {
	return s.ConditionControlModuleCode
}

func (s *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) GetConditionControlModuleValue() *string {
	return s.ConditionControlModuleValue
}

func (s *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) SetConditionControlModuleCode(v string) *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls {
	s.ConditionControlModuleCode = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) SetConditionControlModuleValue(v string) *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls {
	s.ConditionControlModuleValue = &v
	return s
}

func (s *DeleteEnsSaleConditionControlRequestSaleControlsConditionControls) Validate() error {
	return dara.Validate(s)
}
