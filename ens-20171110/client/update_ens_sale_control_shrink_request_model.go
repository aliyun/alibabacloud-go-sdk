// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnsSaleControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *UpdateEnsSaleControlShrinkRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *UpdateEnsSaleControlShrinkRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *UpdateEnsSaleControlShrinkRequest
	GetCustomAccount() *string
	SetSaleControlsShrink(v string) *UpdateEnsSaleControlShrinkRequest
	GetSaleControlsShrink() *string
}

type UpdateEnsSaleControlShrinkRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControlsShrink *string `json:"SaleControls,omitempty" xml:"SaleControls,omitempty"`
}

func (s UpdateEnsSaleControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnsSaleControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnsSaleControlShrinkRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *UpdateEnsSaleControlShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *UpdateEnsSaleControlShrinkRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *UpdateEnsSaleControlShrinkRequest) GetSaleControlsShrink() *string {
	return s.SaleControlsShrink
}

func (s *UpdateEnsSaleControlShrinkRequest) SetAliUidAccount(v string) *UpdateEnsSaleControlShrinkRequest {
	s.AliUidAccount = &v
	return s
}

func (s *UpdateEnsSaleControlShrinkRequest) SetCommodityCode(v string) *UpdateEnsSaleControlShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *UpdateEnsSaleControlShrinkRequest) SetCustomAccount(v string) *UpdateEnsSaleControlShrinkRequest {
	s.CustomAccount = &v
	return s
}

func (s *UpdateEnsSaleControlShrinkRequest) SetSaleControlsShrink(v string) *UpdateEnsSaleControlShrinkRequest {
	s.SaleControlsShrink = &v
	return s
}

func (s *UpdateEnsSaleControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
