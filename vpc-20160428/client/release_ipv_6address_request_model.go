// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseIpv6AddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ReleaseIpv6AddressRequest
	GetClientToken() *string
	SetDryRun(v bool) *ReleaseIpv6AddressRequest
	GetDryRun() *bool
	SetIpv6AddressId(v string) *ReleaseIpv6AddressRequest
	GetIpv6AddressId() *string
	SetOwnerAccount(v string) *ReleaseIpv6AddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseIpv6AddressRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReleaseIpv6AddressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReleaseIpv6AddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseIpv6AddressRequest
	GetResourceOwnerId() *int64
}

type ReleaseIpv6AddressRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// The client generates the value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// > If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without actually creating the IP address pool. The system checks whether the required parameters are specified, whether the request format is valid, and whether business limits are met. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the IPv6 address.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6-2zen5j4axcp5l5qyy****
	Ipv6AddressId *string `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseIpv6AddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseIpv6AddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseIpv6AddressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReleaseIpv6AddressRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ReleaseIpv6AddressRequest) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *ReleaseIpv6AddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseIpv6AddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseIpv6AddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseIpv6AddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseIpv6AddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseIpv6AddressRequest) SetClientToken(v string) *ReleaseIpv6AddressRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetDryRun(v bool) *ReleaseIpv6AddressRequest {
	s.DryRun = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetIpv6AddressId(v string) *ReleaseIpv6AddressRequest {
	s.Ipv6AddressId = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetOwnerAccount(v string) *ReleaseIpv6AddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetOwnerId(v int64) *ReleaseIpv6AddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetRegionId(v string) *ReleaseIpv6AddressRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetResourceOwnerAccount(v string) *ReleaseIpv6AddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) SetResourceOwnerId(v int64) *ReleaseIpv6AddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseIpv6AddressRequest) Validate() error {
	return dara.Validate(s)
}
