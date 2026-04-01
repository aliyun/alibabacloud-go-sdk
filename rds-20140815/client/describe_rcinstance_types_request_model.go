// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeRCInstanceTypesRequest
	GetCommodityCode() *string
	SetEngine(v string) *DescribeRCInstanceTypesRequest
	GetEngine() *string
	SetInstanceType(v []*string) *DescribeRCInstanceTypesRequest
	GetInstanceType() []*string
	SetInstanceTypeFamily(v string) *DescribeRCInstanceTypesRequest
	GetInstanceTypeFamily() *string
	SetRegionId(v string) *DescribeRCInstanceTypesRequest
	GetRegionId() *string
}

type DescribeRCInstanceTypesRequest struct {
	// The commodity code of the instance.
	//
	// example:
	//
	// rds_customprepaid_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The database engine. Set the value to MySQL.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance types.
	InstanceType []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
	// The instance family. You can call the **DescribeRCInstanceTypeFamilies*	- operation to query the instance families of instances.
	//
	// example:
	//
	// gn8.cm
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeRCInstanceTypesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeRCInstanceTypesRequest) GetInstanceType() []*string {
	return s.InstanceType
}

func (s *DescribeRCInstanceTypesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeRCInstanceTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceTypesRequest) SetCommodityCode(v string) *DescribeRCInstanceTypesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRCInstanceTypesRequest) SetEngine(v string) *DescribeRCInstanceTypesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeRCInstanceTypesRequest) SetInstanceType(v []*string) *DescribeRCInstanceTypesRequest {
	s.InstanceType = v
	return s
}

func (s *DescribeRCInstanceTypesRequest) SetInstanceTypeFamily(v string) *DescribeRCInstanceTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeRCInstanceTypesRequest) SetRegionId(v string) *DescribeRCInstanceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceTypesRequest) Validate() error {
	return dara.Validate(s)
}
