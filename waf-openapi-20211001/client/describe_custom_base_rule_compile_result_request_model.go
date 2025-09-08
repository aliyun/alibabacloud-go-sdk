// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBaseRuleCompileResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCustomBaseRuleCompileResultRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeCustomBaseRuleCompileResultRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCustomBaseRuleCompileResultRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCustomBaseRuleCompileResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCustomBaseRuleCompileResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBaseRuleCompileResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustomBaseRuleCompileResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCustomBaseRuleCompileResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCustomBaseRuleCompileResultRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCustomBaseRuleCompileResultRequest) SetInstanceId(v string) *DescribeCustomBaseRuleCompileResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultRequest) SetRegionId(v string) *DescribeCustomBaseRuleCompileResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultRequest) SetResourceManagerResourceGroupId(v string) *DescribeCustomBaseRuleCompileResultRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultRequest) Validate() error {
	return dara.Validate(s)
}
