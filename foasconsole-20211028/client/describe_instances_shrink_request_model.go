// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *DescribeInstancesShrinkRequest
	GetArchitectureType() *string
	SetChargeType(v string) *DescribeInstancesShrinkRequest
	GetChargeType() *string
	SetInstanceId(v string) *DescribeInstancesShrinkRequest
	GetInstanceId() *string
	SetPageIndex(v int32) *DescribeInstancesShrinkRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeInstancesShrinkRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeInstancesShrinkRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribeInstancesShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *DescribeInstancesShrinkRequest
	GetTagsShrink() *string
}

type DescribeInstancesShrinkRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagsShrink      *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesShrinkRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesShrinkRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *DescribeInstancesShrinkRequest) SetArchitectureType(v string) *DescribeInstancesShrinkRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetChargeType(v string) *DescribeInstancesShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetInstanceId(v string) *DescribeInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageIndex(v int32) *DescribeInstancesShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetPageSize(v int32) *DescribeInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetRegion(v string) *DescribeInstancesShrinkRequest {
	s.Region = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetResourceGroupId(v string) *DescribeInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) SetTagsShrink(v string) *DescribeInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *DescribeInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
