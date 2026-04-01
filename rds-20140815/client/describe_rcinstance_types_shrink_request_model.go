// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeRCInstanceTypesShrinkRequest
	GetCommodityCode() *string
	SetEngine(v string) *DescribeRCInstanceTypesShrinkRequest
	GetEngine() *string
	SetInstanceTypeShrink(v string) *DescribeRCInstanceTypesShrinkRequest
	GetInstanceTypeShrink() *string
	SetInstanceTypeFamily(v string) *DescribeRCInstanceTypesShrinkRequest
	GetInstanceTypeFamily() *string
	SetRegionId(v string) *DescribeRCInstanceTypesShrinkRequest
	GetRegionId() *string
}

type DescribeRCInstanceTypesShrinkRequest struct {
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
	InstanceTypeShrink *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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

func (s DescribeRCInstanceTypesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypesShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeRCInstanceTypesShrinkRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeRCInstanceTypesShrinkRequest) GetInstanceTypeShrink() *string {
	return s.InstanceTypeShrink
}

func (s *DescribeRCInstanceTypesShrinkRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeRCInstanceTypesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceTypesShrinkRequest) SetCommodityCode(v string) *DescribeRCInstanceTypesShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRCInstanceTypesShrinkRequest) SetEngine(v string) *DescribeRCInstanceTypesShrinkRequest {
	s.Engine = &v
	return s
}

func (s *DescribeRCInstanceTypesShrinkRequest) SetInstanceTypeShrink(v string) *DescribeRCInstanceTypesShrinkRequest {
	s.InstanceTypeShrink = &v
	return s
}

func (s *DescribeRCInstanceTypesShrinkRequest) SetInstanceTypeFamily(v string) *DescribeRCInstanceTypesShrinkRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeRCInstanceTypesShrinkRequest) SetRegionId(v string) *DescribeRCInstanceTypesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceTypesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
