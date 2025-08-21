// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleConditionControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *DeleteEnsSaleConditionControlShrinkRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *DeleteEnsSaleConditionControlShrinkRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DeleteEnsSaleConditionControlShrinkRequest
	GetCustomAccount() *string
	SetSaleControlsShrink(v string) *DeleteEnsSaleConditionControlShrinkRequest
	GetSaleControlsShrink() *string
}

type DeleteEnsSaleConditionControlShrinkRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControlsShrink *string `json:"SaleControls,omitempty" xml:"SaleControls,omitempty"`
}

func (s DeleteEnsSaleConditionControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleConditionControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) GetSaleControlsShrink() *string {
	return s.SaleControlsShrink
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) SetAliUidAccount(v string) *DeleteEnsSaleConditionControlShrinkRequest {
	s.AliUidAccount = &v
	return s
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) SetCommodityCode(v string) *DeleteEnsSaleConditionControlShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) SetCustomAccount(v string) *DeleteEnsSaleConditionControlShrinkRequest {
	s.CustomAccount = &v
	return s
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) SetSaleControlsShrink(v string) *DeleteEnsSaleConditionControlShrinkRequest {
	s.SaleControlsShrink = &v
	return s
}

func (s *DeleteEnsSaleConditionControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
