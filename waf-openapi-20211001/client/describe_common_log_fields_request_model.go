// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonLogFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCommonLogFieldsRequest
	GetInstanceId() *string
	SetIsDefault(v bool) *DescribeCommonLogFieldsRequest
	GetIsDefault() *bool
	SetIsRequired(v bool) *DescribeCommonLogFieldsRequest
	GetIsRequired() *bool
	SetLogKeyList(v []*string) *DescribeCommonLogFieldsRequest
	GetLogKeyList() []*string
	SetRegionId(v string) *DescribeCommonLogFieldsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCommonLogFieldsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCommonLogFieldsRequest struct {
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
	LogKeyList []*string `json:"LogKeyList,omitempty" xml:"LogKeyList,omitempty" type:"Repeated"`
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

func (s DescribeCommonLogFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonLogFieldsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommonLogFieldsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCommonLogFieldsRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeCommonLogFieldsRequest) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *DescribeCommonLogFieldsRequest) GetLogKeyList() []*string {
	return s.LogKeyList
}

func (s *DescribeCommonLogFieldsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommonLogFieldsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCommonLogFieldsRequest) SetInstanceId(v string) *DescribeCommonLogFieldsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetIsDefault(v bool) *DescribeCommonLogFieldsRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetIsRequired(v bool) *DescribeCommonLogFieldsRequest {
	s.IsRequired = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetLogKeyList(v []*string) *DescribeCommonLogFieldsRequest {
	s.LogKeyList = v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetRegionId(v string) *DescribeCommonLogFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) SetResourceManagerResourceGroupId(v string) *DescribeCommonLogFieldsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCommonLogFieldsRequest) Validate() error {
	return dara.Validate(s)
}
