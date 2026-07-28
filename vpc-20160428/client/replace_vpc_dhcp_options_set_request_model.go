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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DHCP options set to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// dopt-o6w0df4epg9zo8isy****
	DhcpOptionsSetId *string `json:"DhcpOptionsSetId,omitempty" xml:"DhcpOptionsSetId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized Resource Access Management (RAM) users, and missing parameter values. If the request fails the dry run, the corresponding error is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the request passes the authorization and parameter check, a 2XX HTTP status code is returned and the DHCP options set associated with the VPC is directly changed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the DHCP options set resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC whose association you want to change.
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
