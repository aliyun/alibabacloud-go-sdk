// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsSaleControlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *CreateEnsSaleControlShrinkRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *CreateEnsSaleControlShrinkRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *CreateEnsSaleControlShrinkRequest
	GetCustomAccount() *string
	SetSaleControlsShrink(v string) *CreateEnsSaleControlShrinkRequest
	GetSaleControlsShrink() *string
}

type CreateEnsSaleControlShrinkRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	// This parameter is required.
	SaleControlsShrink *string `json:"SaleControls,omitempty" xml:"SaleControls,omitempty"`
}

func (s CreateEnsSaleControlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlShrinkRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *CreateEnsSaleControlShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateEnsSaleControlShrinkRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *CreateEnsSaleControlShrinkRequest) GetSaleControlsShrink() *string {
	return s.SaleControlsShrink
}

func (s *CreateEnsSaleControlShrinkRequest) SetAliUidAccount(v string) *CreateEnsSaleControlShrinkRequest {
	s.AliUidAccount = &v
	return s
}

func (s *CreateEnsSaleControlShrinkRequest) SetCommodityCode(v string) *CreateEnsSaleControlShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateEnsSaleControlShrinkRequest) SetCustomAccount(v string) *CreateEnsSaleControlShrinkRequest {
	s.CustomAccount = &v
	return s
}

func (s *CreateEnsSaleControlShrinkRequest) SetSaleControlsShrink(v string) *CreateEnsSaleControlShrinkRequest {
	s.SaleControlsShrink = &v
	return s
}

func (s *CreateEnsSaleControlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
