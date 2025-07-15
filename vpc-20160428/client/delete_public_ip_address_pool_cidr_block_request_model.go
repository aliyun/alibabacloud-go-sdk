// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicIpAddressPoolCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *DeletePublicIpAddressPoolCidrBlockRequest
	GetCidrBlock() *string
	SetClientToken(v string) *DeletePublicIpAddressPoolCidrBlockRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeletePublicIpAddressPoolCidrBlockRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeletePublicIpAddressPoolCidrBlockRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePublicIpAddressPoolCidrBlockRequest
	GetOwnerId() *int64
	SetPublicIpAddressPoolId(v string) *DeletePublicIpAddressPoolCidrBlockRequest
	GetPublicIpAddressPoolId() *string
	SetRegionId(v string) *DeletePublicIpAddressPoolCidrBlockRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePublicIpAddressPoolCidrBlockRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePublicIpAddressPoolCidrBlockRequest
	GetResourceOwnerId() *int64
}

type DeletePublicIpAddressPoolCidrBlockRequest struct {
	// The CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.0.XX.XX/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
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
	// The region ID of the IP address pool from which you want to delete a CIDR block.
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

func (s DeletePublicIpAddressPoolCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicIpAddressPoolCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetPublicIpAddressPoolId() *string {
	return s.PublicIpAddressPoolId
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetCidrBlock(v string) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.CidrBlock = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetClientToken(v string) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetDryRun(v bool) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.DryRun = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetOwnerAccount(v string) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetOwnerId(v int64) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetPublicIpAddressPoolId(v string) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.PublicIpAddressPoolId = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetRegionId(v string) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetResourceOwnerAccount(v string) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) SetResourceOwnerId(v int64) *DeletePublicIpAddressPoolCidrBlockRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePublicIpAddressPoolCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
