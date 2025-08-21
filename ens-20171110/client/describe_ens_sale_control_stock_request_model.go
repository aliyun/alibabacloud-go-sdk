// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlStockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUidAccount(v string) *DescribeEnsSaleControlStockRequest
	GetAliUidAccount() *string
	SetCommodityCode(v string) *DescribeEnsSaleControlStockRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DescribeEnsSaleControlStockRequest
	GetCustomAccount() *string
	SetModuleCode(v string) *DescribeEnsSaleControlStockRequest
	GetModuleCode() *string
	SetOrderType(v string) *DescribeEnsSaleControlStockRequest
	GetOrderType() *string
}

type DescribeEnsSaleControlStockRequest struct {
	AliUidAccount *string `json:"AliUidAccount,omitempty" xml:"AliUidAccount,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeEnsSaleControlStockRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockRequest) GetAliUidAccount() *string {
	return s.AliUidAccount
}

func (s *DescribeEnsSaleControlStockRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsSaleControlStockRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DescribeEnsSaleControlStockRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeEnsSaleControlStockRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeEnsSaleControlStockRequest) SetAliUidAccount(v string) *DescribeEnsSaleControlStockRequest {
	s.AliUidAccount = &v
	return s
}

func (s *DescribeEnsSaleControlStockRequest) SetCommodityCode(v string) *DescribeEnsSaleControlStockRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsSaleControlStockRequest) SetCustomAccount(v string) *DescribeEnsSaleControlStockRequest {
	s.CustomAccount = &v
	return s
}

func (s *DescribeEnsSaleControlStockRequest) SetModuleCode(v string) *DescribeEnsSaleControlStockRequest {
	s.ModuleCode = &v
	return s
}

func (s *DescribeEnsSaleControlStockRequest) SetOrderType(v string) *DescribeEnsSaleControlStockRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeEnsSaleControlStockRequest) Validate() error {
	return dara.Validate(s)
}
