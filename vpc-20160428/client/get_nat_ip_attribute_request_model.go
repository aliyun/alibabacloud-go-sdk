// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatIpAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetNatIpAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *GetNatIpAttributeRequest
	GetDryRun() *bool
	SetNatIpId(v string) *GetNatIpAttributeRequest
	GetNatIpId() *string
	SetOwnerAccount(v string) *GetNatIpAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetNatIpAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetNatIpAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetNatIpAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetNatIpAttributeRequest
	GetResourceOwnerId() *int64
}

type GetNatIpAttributeRequest struct {
	// Client Token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client, ensuring that it is unique across different requests. ClientToken only supports ASCII characters. If not specified, the system automatically uses the **RequestId*	- of the API request as the **ClientToken*	- identifier. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Indicates whether to perform a dry run of the request. Values:
	//
	// - **true**: Sends a check request without querying NAT IP address information. The checks include whether the AccessKey is valid, the RAM user\\"s authorization status, and if all required parameters are filled out. If any check fails, the corresponding error is returned. If all checks pass, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After passing the checks, it returns an HTTP 2xx status code and queries the NAT IP address information.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the NAT IP address instance to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpcnatip-gw8y7q3cpk3fggs87****
	NatIpId      *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway instance to which the NAT IP address to be queried belongs. You can obtain the region ID by calling the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetNatIpAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNatIpAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetNatIpAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetNatIpAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetNatIpAttributeRequest) GetNatIpId() *string {
	return s.NatIpId
}

func (s *GetNatIpAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetNatIpAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetNatIpAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNatIpAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetNatIpAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetNatIpAttributeRequest) SetClientToken(v string) *GetNatIpAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetDryRun(v bool) *GetNatIpAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetNatIpId(v string) *GetNatIpAttributeRequest {
	s.NatIpId = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetOwnerAccount(v string) *GetNatIpAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetOwnerId(v int64) *GetNatIpAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetRegionId(v string) *GetNatIpAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetResourceOwnerAccount(v string) *GetNatIpAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetNatIpAttributeRequest) SetResourceOwnerId(v int64) *GetNatIpAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetNatIpAttributeRequest) Validate() error {
	return dara.Validate(s)
}
