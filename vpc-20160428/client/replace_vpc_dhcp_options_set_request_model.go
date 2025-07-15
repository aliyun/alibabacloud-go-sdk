// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceVpcDhcpOptionsSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReplaceVpcDhcpOptionsSetRequest
	GetClientToken() *string
	SetDhcpOptionsSetId(v string) *ReplaceVpcDhcpOptionsSetRequest
	GetDhcpOptionsSetId() *string
	SetDryRun(v bool) *ReplaceVpcDhcpOptionsSetRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ReplaceVpcDhcpOptionsSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReplaceVpcDhcpOptionsSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReplaceVpcDhcpOptionsSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReplaceVpcDhcpOptionsSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReplaceVpcDhcpOptionsSetRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *ReplaceVpcDhcpOptionsSetRequest
	GetVpcId() *string
}

type ReplaceVpcDhcpOptionsSetRequest struct {
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
	// The ID of the new DHCP options set.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks whether your AccessKey pair is valid, whether the Resource Access Management (RAM) user is authorized, and whether the required parameters are set. If the request fails to pass the check, the corresponding error message is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, a 2XX HTTP status code is returned and the operation is performed.
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
	// The ID of the associated VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-dsferghthth****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ReplaceVpcDhcpOptionsSetRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceVpcDhcpOptionsSetRequest) GoString() string {
	return s.String()
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetDhcpOptionsSetId() *string {
	return s.DhcpOptionsSetId
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReplaceVpcDhcpOptionsSetRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetClientToken(v string) *ReplaceVpcDhcpOptionsSetRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetDhcpOptionsSetId(v string) *ReplaceVpcDhcpOptionsSetRequest {
	s.DhcpOptionsSetId = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetDryRun(v bool) *ReplaceVpcDhcpOptionsSetRequest {
	s.DryRun = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetOwnerAccount(v string) *ReplaceVpcDhcpOptionsSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetOwnerId(v int64) *ReplaceVpcDhcpOptionsSetRequest {
	s.OwnerId = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetRegionId(v string) *ReplaceVpcDhcpOptionsSetRequest {
	s.RegionId = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetResourceOwnerAccount(v string) *ReplaceVpcDhcpOptionsSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetResourceOwnerId(v int64) *ReplaceVpcDhcpOptionsSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) SetVpcId(v string) *ReplaceVpcDhcpOptionsSetRequest {
	s.VpcId = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetRequest) Validate() error {
	return dara.Validate(s)
}
