// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcPeerConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptingAliUid(v int64) *CreateVpcPeerConnectionRequest
	GetAcceptingAliUid() *int64
	SetAcceptingRegionId(v string) *CreateVpcPeerConnectionRequest
	GetAcceptingRegionId() *string
	SetAcceptingVpcId(v string) *CreateVpcPeerConnectionRequest
	GetAcceptingVpcId() *string
	SetBandwidth(v int32) *CreateVpcPeerConnectionRequest
	GetBandwidth() *int32
	SetClientToken(v string) *CreateVpcPeerConnectionRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVpcPeerConnectionRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVpcPeerConnectionRequest
	GetDryRun() *bool
	SetLinkType(v string) *CreateVpcPeerConnectionRequest
	GetLinkType() *string
	SetName(v string) *CreateVpcPeerConnectionRequest
	GetName() *string
	SetRegionId(v string) *CreateVpcPeerConnectionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpcPeerConnectionRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateVpcPeerConnectionRequestTag) *CreateVpcPeerConnectionRequest
	GetTag() []*CreateVpcPeerConnectionRequestTag
	SetVpcId(v string) *CreateVpcPeerConnectionRequest
	GetVpcId() *string
}

type CreateVpcPeerConnectionRequest struct {
	// The ID of the Alibaba Cloud account to which the accepter VPC belongs.
	//
	// 	- To create a VPC peering connection within your Alibaba Cloud account, enter the ID of your Alibaba Cloud account.
	//
	// 	- To create a VPC peering connection between your Alibaba Cloud account and another Alibaba Cloud account, enter the ID of the peer Alibaba Cloud account.
	//
	// >  If the accepter is a RAM user, set **AcceptingAliUid*	- to the ID of the Alibaba Cloud account that created the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	AcceptingAliUid *int64 `json:"AcceptingAliUid,omitempty" xml:"AcceptingAliUid,omitempty"`
	// The region ID of the accepter VPC of the VPC peering connection that you want to create.
	//
	// 	- To create an intra-region VPC peering connection, enter a region ID that is the same as that of the requester VPC.
	//
	// 	- To create an inter-region VPC peering connection, enter a region ID that is different from that of the requester VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The ID of the accepter VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1vzjkp2q1xgnind****
	AcceptingVpcId *string `json:"AcceptingVpcId,omitempty" xml:"AcceptingVpcId,omitempty"`
	// The bandwidth of the VPC peering connection. Unit: Mbit/s. The value must be an integer greater than 0. Before you specify this parameter, make sure that you create an inter-region VPC peering connection.
	//
	// example:
	//
	// 100
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VPC peering connection.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The link type of the VPC peering connection that you want to create. Valid values:
	//
	// - Platinum.
	//
	// - Gold: default value.
	//
	// >
	//
	// > - If you need to specify this parameter, ensure that the VPC peering connection is an inter-region connection.
	//
	// example:
	//
	// Gold
	LinkType *string `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// The name of the VPC peering connection.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// vpcpeer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where you want to create a VPC peering connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateVpcPeerConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the requester VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1gsk7h12ew7oegk****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcPeerConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPeerConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionRequest) GetAcceptingAliUid() *int64 {
	return s.AcceptingAliUid
}

func (s *CreateVpcPeerConnectionRequest) GetAcceptingRegionId() *string {
	return s.AcceptingRegionId
}

func (s *CreateVpcPeerConnectionRequest) GetAcceptingVpcId() *string {
	return s.AcceptingVpcId
}

func (s *CreateVpcPeerConnectionRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateVpcPeerConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcPeerConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVpcPeerConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcPeerConnectionRequest) GetLinkType() *string {
	return s.LinkType
}

func (s *CreateVpcPeerConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateVpcPeerConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcPeerConnectionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcPeerConnectionRequest) GetTag() []*CreateVpcPeerConnectionRequestTag {
	return s.Tag
}

func (s *CreateVpcPeerConnectionRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcPeerConnectionRequest) SetAcceptingAliUid(v int64) *CreateVpcPeerConnectionRequest {
	s.AcceptingAliUid = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetAcceptingRegionId(v string) *CreateVpcPeerConnectionRequest {
	s.AcceptingRegionId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetAcceptingVpcId(v string) *CreateVpcPeerConnectionRequest {
	s.AcceptingVpcId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetBandwidth(v int32) *CreateVpcPeerConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetClientToken(v string) *CreateVpcPeerConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetDescription(v string) *CreateVpcPeerConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetDryRun(v bool) *CreateVpcPeerConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetLinkType(v string) *CreateVpcPeerConnectionRequest {
	s.LinkType = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetName(v string) *CreateVpcPeerConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetRegionId(v string) *CreateVpcPeerConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetResourceGroupId(v string) *CreateVpcPeerConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetTag(v []*CreateVpcPeerConnectionRequestTag) *CreateVpcPeerConnectionRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcPeerConnectionRequest) SetVpcId(v string) *CreateVpcPeerConnectionRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcPeerConnectionRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVpcPeerConnectionRequestTag struct {
	// The tag key. You must specify at least one tag key and at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `acs:` or `aliyun` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You must specify at least one tag value and can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcPeerConnectionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPeerConnectionRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpcPeerConnectionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpcPeerConnectionRequestTag) SetKey(v string) *CreateVpcPeerConnectionRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcPeerConnectionRequestTag) SetValue(v string) *CreateVpcPeerConnectionRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpcPeerConnectionRequestTag) Validate() error {
	return dara.Validate(s)
}
