// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcPeerConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVpcPeerConnectionsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListVpcPeerConnectionsRequest
	GetMaxResults() *int32
	SetName(v string) *ListVpcPeerConnectionsRequest
	GetName() *string
	SetNextToken(v string) *ListVpcPeerConnectionsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcPeerConnectionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListVpcPeerConnectionsRequest
	GetResourceGroupId() *string
	SetTags(v []*ListVpcPeerConnectionsRequestTags) *ListVpcPeerConnectionsRequest
	GetTags() []*ListVpcPeerConnectionsRequestTags
	SetVpcId(v []*string) *ListVpcPeerConnectionsRequest
	GetVpcId() []*string
}

type ListVpcPeerConnectionsRequest struct {
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
	Tags []*ListVpcPeerConnectionsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the requester VPC or accepter VPC of the VPC peering connection that you want to query.
	VpcId []*string `json:"VpcId,omitempty" xml:"VpcId,omitempty" type:"Repeated"`
}

func (s ListVpcPeerConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVpcPeerConnectionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcPeerConnectionsRequest) GetName() *string {
	return s.Name
}

func (s *ListVpcPeerConnectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcPeerConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcPeerConnectionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVpcPeerConnectionsRequest) GetTags() []*ListVpcPeerConnectionsRequestTags {
	return s.Tags
}

func (s *ListVpcPeerConnectionsRequest) GetVpcId() []*string {
	return s.VpcId
}

func (s *ListVpcPeerConnectionsRequest) SetInstanceId(v string) *ListVpcPeerConnectionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetMaxResults(v int32) *ListVpcPeerConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetName(v string) *ListVpcPeerConnectionsRequest {
	s.Name = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetNextToken(v string) *ListVpcPeerConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetRegionId(v string) *ListVpcPeerConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetResourceGroupId(v string) *ListVpcPeerConnectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetTags(v []*ListVpcPeerConnectionsRequestTags) *ListVpcPeerConnectionsRequest {
	s.Tags = v
	return s
}

func (s *ListVpcPeerConnectionsRequest) SetVpcId(v []*string) *ListVpcPeerConnectionsRequest {
	s.VpcId = v
	return s
}

func (s *ListVpcPeerConnectionsRequest) Validate() error {
	return dara.Validate(s)
}

type ListVpcPeerConnectionsRequestTags struct {
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

func (s ListVpcPeerConnectionsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListVpcPeerConnectionsRequestTags) GoString() string {
	return s.String()
}

func (s *ListVpcPeerConnectionsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListVpcPeerConnectionsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListVpcPeerConnectionsRequestTags) SetKey(v string) *ListVpcPeerConnectionsRequestTags {
	s.Key = &v
	return s
}

func (s *ListVpcPeerConnectionsRequestTags) SetValue(v string) *ListVpcPeerConnectionsRequestTags {
	s.Value = &v
	return s
}

func (s *ListVpcPeerConnectionsRequestTags) Validate() error {
	return dara.Validate(s)
}
