// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDhcpOptionsSetToVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachDhcpOptionsSetToVpcRequest
	GetClientToken() *string
	SetDhcpOptionsSetId(v string) *AttachDhcpOptionsSetToVpcRequest
	GetDhcpOptionsSetId() *string
	SetDryRun(v bool) *AttachDhcpOptionsSetToVpcRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AttachDhcpOptionsSetToVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachDhcpOptionsSetToVpcRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachDhcpOptionsSetToVpcRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachDhcpOptionsSetToVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachDhcpOptionsSetToVpcRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *AttachDhcpOptionsSetToVpcRequest
	GetVpcId() *string
}

type AttachDhcpOptionsSetToVpcRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DHCP options set.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// **true**: checks the request without performing the operation. The system checks whether your AccessKey pair is valid, whether the Resource Access Management (RAM) user is authorized, and whether the required parameters are set. If the request fails to pass the check, the corresponding error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): sends the request. If the request passes the check, a 2XX HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region to which the DHCP options set belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC to be associated with the DHCP options set.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-sfdkfdjkdf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AttachDhcpOptionsSetToVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDhcpOptionsSetToVpcRequest) GoString() string {
	return s.String()
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachDhcpOptionsSetToVpcRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetClientToken(v string) *AttachDhcpOptionsSetToVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetDhcpOptionsSetId(v string) *AttachDhcpOptionsSetToVpcRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetDryRun(v bool) *AttachDhcpOptionsSetToVpcRequest {
	s.DryRun = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetOwnerAccount(v string) *AttachDhcpOptionsSetToVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetOwnerId(v int64) *AttachDhcpOptionsSetToVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetRegionId(v string) *AttachDhcpOptionsSetToVpcRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetResourceOwnerAccount(v string) *AttachDhcpOptionsSetToVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetResourceOwnerId(v int64) *AttachDhcpOptionsSetToVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) SetVpcId(v string) *AttachDhcpOptionsSetToVpcRequest {
	s.VpcId = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcRequest) Validate() error {
	return dara.Validate(s)
}
