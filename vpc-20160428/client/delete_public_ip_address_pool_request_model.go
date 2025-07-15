// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicIpAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeletePublicIpAddressPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeletePublicIpAddressPoolRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeletePublicIpAddressPoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePublicIpAddressPoolRequest
	GetOwnerId() *int64
	SetPublicIpAddressPoolId(v string) *DeletePublicIpAddressPoolRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *DeletePublicIpAddressPoolRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePublicIpAddressPoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePublicIpAddressPoolRequest
	GetResourceOwnerId() *int64
}

type DeletePublicIpAddressPoolRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe60000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the IP address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// pippool-6wetvn6fumkgycssx****
	PublicIpAddressPoolId *string `json:"PublicIpAddressPoolId,omitempty" xml:"PublicIpAddressPoolId,omitempty"`
	// The ID of the region where you want to create the IP address pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeletePublicIpAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicIpAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *DeletePublicIpAddressPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletePublicIpAddressPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeletePublicIpAddressPoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePublicIpAddressPoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePublicIpAddressPoolRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *DeletePublicIpAddressPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePublicIpAddressPoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePublicIpAddressPoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePublicIpAddressPoolRequest) SetClientToken(v string) *DeletePublicIpAddressPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetDryRun(v bool) *DeletePublicIpAddressPoolRequest {
	s.DryRun = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetOwnerAccount(v string) *DeletePublicIpAddressPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetOwnerId(v int64) *DeletePublicIpAddressPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetPublicIpAddressPoolId(v string) *DeletePublicIpAddressPoolRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetRegionId(v string) *DeletePublicIpAddressPoolRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetResourceOwnerAccount(v string) *DeletePublicIpAddressPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) SetResourceOwnerId(v int64) *DeletePublicIpAddressPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePublicIpAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}
