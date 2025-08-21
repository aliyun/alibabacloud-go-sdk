// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *DescribeEnsSaleControlRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *DescribeEnsSaleControlRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DescribeEnsSaleControlRequest
	GetCustomAccount() *string
	SetModuleCode(v string) *DescribeEnsSaleControlRequest
	GetModuleCode() *string
	SetOrderType(v string) *DescribeEnsSaleControlRequest
	GetOrderType() *string
}

type DescribeEnsSaleControlRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeEnsSaleControlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *DescribeEnsSaleControlRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsSaleControlRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DescribeEnsSaleControlRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeEnsSaleControlRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeEnsSaleControlRequest) SetAliUidAccount(v string) *DescribeEnsSaleControlRequest {
	s.AliUidAccount = &v
	return s
}

func (s *DescribeEnsSaleControlRequest) SetCommodityCode(v string) *DescribeEnsSaleControlRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsSaleControlRequest) SetCustomAccount(v string) *DescribeEnsSaleControlRequest {
	s.CustomAccount = &v
	return s
}

func (s *DescribeEnsSaleControlRequest) SetModuleCode(v string) *DescribeEnsSaleControlRequest {
	s.ModuleCode = &v
	return s
}

func (s *DescribeEnsSaleControlRequest) SetOrderType(v string) *DescribeEnsSaleControlRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeEnsSaleControlRequest) Validate() error {
	return dara.Validate(s)
}
