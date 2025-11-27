// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuickSaleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodity(v string) *DescribeQuickSaleConfigRequest
	GetCommodity() *string
	SetEngine(v string) *DescribeQuickSaleConfigRequest
	GetEngine() *string
	SetRegionId(v string) *DescribeQuickSaleConfigRequest
	GetRegionId() *string
}

type DescribeQuickSaleConfigRequest struct {
	// The product code. Valid values:
	//
	// 	- rds: The instance is a subscription instance.
	//
	// 	- bards: The instance is a pay-as-you-go instance.
	//
	// example:
	//
	// rds
	Commodity *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeQuickSaleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuickSaleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeQuickSaleConfigRequest) GetCommodity() *string {
	return s.Commodity
}

func (s *DescribeQuickSaleConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeQuickSaleConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeQuickSaleConfigRequest) SetCommodity(v string) *DescribeQuickSaleConfigRequest {
	s.Commodity = &v
	return s
}

func (s *DescribeQuickSaleConfigRequest) SetEngine(v string) *DescribeQuickSaleConfigRequest {
	s.Engine = &v
	return s
}

func (s *DescribeQuickSaleConfigRequest) SetRegionId(v string) *DescribeQuickSaleConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeQuickSaleConfigRequest) Validate() error {
	return dara.Validate(s)
}
