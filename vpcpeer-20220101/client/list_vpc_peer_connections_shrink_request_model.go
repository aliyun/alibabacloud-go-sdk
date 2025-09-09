// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPeerConnectionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVpcPeerConnectionsShrinkRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListVpcPeerConnectionsShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListVpcPeerConnectionsShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListVpcPeerConnectionsShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcPeerConnectionsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpcPeerConnectionsShrinkRequest
	GetResourceGroupId() *string
	SetTags(v []*ListVpcPeerConnectionsShrinkRequestTags) *ListVpcPeerConnectionsShrinkRequest
	GetTags() []*ListVpcPeerConnectionsShrinkRequestTags
	SetVpcIdShrink(v string) *ListVpcPeerConnectionsShrinkRequest
	GetVpcIdShrink() *string
}

type ListVpcPeerConnectionsShrinkRequest struct {
	// The ID of the VPC peering connection that you want to query.
	//
	// example:
	//
	// pcc-lnk0m24khwvtkm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the VPC peering connection that you want to query.
	//
	// example:
	//
	// vpcpeer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where you want to query VPC peering connections.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfm2ggeub5uf3y
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	Tags []*ListVpcPeerConnectionsShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the requester VPC or accepter VPC of the VPC peering connection that you want to query.
	VpcIdShrink *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListVpcPeerConnectionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetTags() []*ListVpcPeerConnectionsShrinkRequestTags {
	return s.Tags
}

func (s *ListVpcPeerConnectionsShrinkRequest) GetVpcIdShrink() *string {
	return s.VpcIdShrink
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetInstanceId(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetMaxResults(v int32) *ListVpcPeerConnectionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetName(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetNextToken(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetRegionId(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetResourceGroupId(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetTags(v []*ListVpcPeerConnectionsShrinkRequestTags) *ListVpcPeerConnectionsShrinkRequest {
	s.Tags = v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) SetVpcIdShrink(v string) *ListVpcPeerConnectionsShrinkRequest {
	s.VpcIdShrink = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type ListVpcPeerConnectionsShrinkRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVpcPeerConnectionsShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) SetKey(v string) *ListVpcPeerConnectionsShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) SetValue(v string) *ListVpcPeerConnectionsShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *ListVpcPeerConnectionsShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
