// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonLogFieldsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCommonLogFieldsShrinkRequest
	GetInstanceId() *string
	SetIsDefault(v bool) *DescribeCommonLogFieldsShrinkRequest
	GetIsDefault() *bool
	SetIsRequired(v bool) *DescribeCommonLogFieldsShrinkRequest
	GetIsRequired() *bool
	SetLogKeyListShrink(v string) *DescribeCommonLogFieldsShrinkRequest
	GetLogKeyListShrink() *string
	SetRegionId(v string) *DescribeCommonLogFieldsShrinkRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCommonLogFieldsShrinkRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCommonLogFieldsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-l*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// false
	IsRequired       *bool   `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	LogKeyListShrink *string `json:"LogKeyList,omitempty" xml:"LogKeyList,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek2sxgs*****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCommonLogFieldsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonLogFieldsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonLogFieldsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCommonLogFieldsShrinkRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeCommonLogFieldsShrinkRequest) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *DescribeCommonLogFieldsShrinkRequest) GetLogKeyListShrink() *string {
	return s.LogKeyListShrink
}

func (s *DescribeCommonLogFieldsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommonLogFieldsShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCommonLogFieldsShrinkRequest) SetInstanceId(v string) *DescribeCommonLogFieldsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommonLogFieldsShrinkRequest) SetIsDefault(v bool) *DescribeCommonLogFieldsShrinkRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeCommonLogFieldsShrinkRequest) SetIsRequired(v bool) *DescribeCommonLogFieldsShrinkRequest {
	s.IsRequired = &v
	return s
}

func (s *DescribeCommonLogFieldsShrinkRequest) SetLogKeyListShrink(v string) *DescribeCommonLogFieldsShrinkRequest {
	s.LogKeyListShrink = &v
	return s
}

func (s *DescribeCommonLogFieldsShrinkRequest) SetRegionId(v string) *DescribeCommonLogFieldsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommonLogFieldsShrinkRequest) SetResourceManagerResourceGroupId(v string) *DescribeCommonLogFieldsShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCommonLogFieldsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
