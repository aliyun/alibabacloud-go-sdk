// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnsSaleControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *DeleteEnsSaleControlShrinkRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *DeleteEnsSaleControlShrinkRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DeleteEnsSaleControlShrinkRequest
	GetCustomAccount() *string
	SetSaleControlsShrink(v string) *DeleteEnsSaleControlShrinkRequest
	GetSaleControlsShrink() *string
}

type DeleteEnsSaleControlShrinkRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControlsShrink *string `json:"SaleControls,omitempty" xml:"SaleControls,omitempty"`
}

func (s DeleteEnsSaleControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnsSaleControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnsSaleControlShrinkRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *DeleteEnsSaleControlShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DeleteEnsSaleControlShrinkRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DeleteEnsSaleControlShrinkRequest) GetSaleControlsShrink() *string {
	return s.SaleControlsShrink
}

func (s *DeleteEnsSaleControlShrinkRequest) SetAliUidAccount(v string) *DeleteEnsSaleControlShrinkRequest {
	s.AliUidAccount = &v
	return s
}

func (s *DeleteEnsSaleControlShrinkRequest) SetCommodityCode(v string) *DeleteEnsSaleControlShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *DeleteEnsSaleControlShrinkRequest) SetCustomAccount(v string) *DeleteEnsSaleControlShrinkRequest {
	s.CustomAccount = &v
	return s
}

func (s *DeleteEnsSaleControlShrinkRequest) SetSaleControlsShrink(v string) *DeleteEnsSaleControlShrinkRequest {
	s.SaleControlsShrink = &v
	return s
}

func (s *DeleteEnsSaleControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
