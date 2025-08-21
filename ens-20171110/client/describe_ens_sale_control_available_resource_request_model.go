// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlAvailableResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeEnsSaleControlAvailableResourceRequest
	GetCommodityCode() *string
	SetCustomAccount(v string) *DescribeEnsSaleControlAvailableResourceRequest
	GetCustomAccount() *string
	SetOrderType(v string) *DescribeEnsSaleControlAvailableResourceRequest
	GetOrderType() *string
}

type DescribeEnsSaleControlAvailableResourceRequest struct {
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CustomAccount *string `json:"CustomAccount,omitempty" xml:"CustomAccount,omitempty"`
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) GetCustomAccount() *string {
	return s.CustomAccount
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) SetCommodityCode(v string) *DescribeEnsSaleControlAvailableResourceRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) SetCustomAccount(v string) *DescribeEnsSaleControlAvailableResourceRequest {
	s.CustomAccount = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) SetOrderType(v string) *DescribeEnsSaleControlAvailableResourceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceRequest) Validate() error {
	return dara.Validate(s)
}
