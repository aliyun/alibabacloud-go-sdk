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
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-l*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to query for default log fields.
	//
	// - **true**: Queries for default log fields.
	//
	// - **false**: Queries for non-default log fields.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// Specifies whether to query for required log fields.
	//
	// - **true**: Queries for required log fields.
	//
	// - **false**: Queries for non-required log fields.
	//
	// example:
	//
	// false
	IsRequired *bool `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	// The list of log fields to query.
	LogKeyListShrink *string `json:"LogKeyList,omitempty" xml:"LogKeyList,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
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
