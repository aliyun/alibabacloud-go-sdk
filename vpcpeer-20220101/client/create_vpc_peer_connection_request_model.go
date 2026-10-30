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
	SetEnableJumboFrame(v bool) *CreateVpcPeerConnectionRequest
	GetEnableJumboFrame() *bool
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
	// The ID of the Alibaba Cloud account to which the acceptor of the VPC peering connection belongs.
	//
	// - If you enter your own Alibaba Cloud account ID, a same-account VPC peering connection is created.
	//
	// - If you enter the ID of another Alibaba Cloud account, a cross-account VPC peering connection is created.
	//
	//
	//
	// > If the acceptor account is a Resource Access Management (RAM) user, you must enter the ID of the Alibaba Cloud account to which the RAM user belongs in **AcceptingAliUid**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	AcceptingAliUid *int64 `json:"AcceptingAliUid,omitempty" xml:"AcceptingAliUid,omitempty"`
	// The ID of the region where the acceptor of the VPC peering connection is deployed.
	//
	// - To create an intra-region VPC peering connection, enter the same region ID as the initiator.
	//
	// - To create an inter-region VPC peering connection, enter a region ID different from that of the initiator.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AcceptingRegionId *string `json:"AcceptingRegionId,omitempty" xml:"AcceptingRegionId,omitempty"`
	// The ID of the VPC that acts as the acceptor of the VPC peering connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1vzjkp2q1xgnind****
	AcceptingVpcId *string `json:"AcceptingVpcId,omitempty" xml:"AcceptingVpcId,omitempty"`
	// The bandwidth of the VPC peering connection. Unit: Mbit/s. The value must be a positive integer.
	//
	// > If you want to specify this parameter, make sure that you are creating an inter-region peering connection.
	//
	// example:
	//
	// 100
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the VPC peering connection. The description must be 2 to 256 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run for this request. Valid values:
	//
	// - **true**: performs a dry run and does not create a VPC peering connection. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the dry run, the corresponding error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable the jumbo frame feature for the VPC peering connection. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// > Note: If you want to specify this parameter, make sure that you are creating an inter-region peering connection.
	//
	// example:
	//
	// true
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// The link type.
	//
	// Valid values: Platinum and Gold. Default value: Gold.
	//
	// > If you want to specify this parameter, make sure that you are creating an inter-region peering connection.
	//
	// example:
	//
	// Gold
	LinkType *string `json:"LinkType,omitempty" xml:"LinkType,omitempty"`
	// The name of the VPC peering connection. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// vpcpeer
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where you want to create a VPC peering connection.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/94475.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag information.
	Tag []*CreateVpcPeerConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID create a VPC that acts as the initiator create a VPC peering connection.
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

func (s *CreateVpcPeerConnectionRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
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

func (s *CreateVpcPeerConnectionRequest) SetEnableJumboFrame(v bool) *CreateVpcPeerConnectionRequest {
	s.EnableJumboFrame = &v
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVpcPeerConnectionRequestTag struct {
	// The tag key of the resource. You must specify at least one tag key and can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You must specify at least one tag value and can specify up to 20 tag values. The tag value can be an empty string.
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
