// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsCommodityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeRdsCommodityRequest
	GetCommodityCode() *string
	SetDrdsInstanceId(v string) *DescribeRdsCommodityRequest
	GetDrdsInstanceId() *string
	SetOrderType(v string) *DescribeRdsCommodityRequest
	GetOrderType() *string
}

type DescribeRdsCommodityRequest struct {
	// The commodity code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsPost
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The type of the order.
	//
	// example:
	//
	// 1
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeRdsCommodityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsCommodityRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeRdsCommodityRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRdsCommodityRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeRdsCommodityRequest) SetCommodityCode(v string) *DescribeRdsCommodityRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetDrdsInstanceId(v string) *DescribeRdsCommodityRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRdsCommodityRequest) SetOrderType(v string) *DescribeRdsCommodityRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeRdsCommodityRequest) Validate() error {
	return dara.Validate(s)
}
