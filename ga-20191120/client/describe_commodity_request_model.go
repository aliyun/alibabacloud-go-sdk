// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommodityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeCommodityRequest
	GetCommodityCode() *string
	SetOrderType(v string) *DescribeCommodityRequest
	GetOrderType() *string
	SetRegionId(v string) *DescribeCommodityRequest
	GetRegionId() *string
}

type DescribeCommodityRequest struct {
	// The commodity code.
	//
	// Valid values on the China site (aliyun.com):
	//
	// 	- **ga_gapluspre_public_cn**: GA instance.
	//
	// 	- **ga_plusbwppre_public_cn**: basic bandwidth plan.
	//
	// Valid values on the international site (alibabacloud.com):
	//
	// 	- **ga_pluspre_public_intl**: GA instance.
	//
	// 	- **ga_bwppreintl_public_intl:*	- basic bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga_gapluspre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The type of the order. Valid values:
	//
	// 	- **BUY**: purchase order.
	//
	// 	- **RENEW**: renewal order.
	//
	// 	- **UPGRADE**: upgrade order.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCommodityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommodityRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCommodityRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeCommodityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommodityRequest) SetCommodityCode(v string) *DescribeCommodityRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCommodityRequest) SetOrderType(v string) *DescribeCommodityRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeCommodityRequest) SetRegionId(v string) *DescribeCommodityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommodityRequest) Validate() error {
	return dara.Validate(s)
}
