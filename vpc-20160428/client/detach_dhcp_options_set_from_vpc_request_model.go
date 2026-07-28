// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDhcpOptionsSetFromVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachDhcpOptionsSetFromVpcRequest
	GetClientToken() *string
	SetDhcpOptionsSetId(v string) *DetachDhcpOptionsSetFromVpcRequest
	GetDhcpOptionsSetId() *string
	SetDryRun(v bool) *DetachDhcpOptionsSetFromVpcRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DetachDhcpOptionsSetFromVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachDhcpOptionsSetFromVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachDhcpOptionsSetFromVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachDhcpOptionsSetFromVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachDhcpOptionsSetFromVpcRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DetachDhcpOptionsSetFromVpcRequest
	GetVpcId() *string
}

type DetachDhcpOptionsSetFromVpcRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DHCP options set to disassociate from the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized Resource Access Management (RAM) users, and missing parameter values. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): sends a Normal request. If the request passes the authorization and parameter check, a 2XX HTTP status code is returned and the DHCP options set is disassociated from the VPC.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the DHCP options set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC to disassociate from the DHCP options set.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-dfdgrgthhy****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DetachDhcpOptionsSetFromVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachDhcpOptionsSetFromVpcRequest) GoString() string {
	return s.String()
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachDhcpOptionsSetFromVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetClientToken(v string) *DetachDhcpOptionsSetFromVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetDhcpOptionsSetId(v string) *DetachDhcpOptionsSetFromVpcRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetDryRun(v bool) *DetachDhcpOptionsSetFromVpcRequest {
	s.DryRun = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetOwnerAccount(v string) *DetachDhcpOptionsSetFromVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetOwnerId(v int64) *DetachDhcpOptionsSetFromVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetRegionId(v string) *DetachDhcpOptionsSetFromVpcRequest {
	s.RegionId = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetResourceOwnerAccount(v string) *DetachDhcpOptionsSetFromVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetResourceOwnerId(v int64) *DetachDhcpOptionsSetFromVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) SetVpcId(v string) *DetachDhcpOptionsSetFromVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcRequest) Validate() error {
	return dara.Validate(s)
}
