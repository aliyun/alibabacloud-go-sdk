// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpamPoolRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpamPoolRequest
	GetDryRun() *bool
	SetIpamPoolId(v string) *DeleteIpamPoolRequest
	GetIpamPoolId() *string
	SetOwnerAccount(v string) *DeleteIpamPoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpamPoolRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpamPoolRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpamPoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpamPoolRequest
	GetResourceOwnerId() *int64
}

type DeleteIpamPoolRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// **
	//
	// **Usage notes*	- If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, DryRunOperation is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId   *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s DeleteIpamPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpamPoolRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpamPoolRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *DeleteIpamPoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpamPoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpamPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpamPoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpamPoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpamPoolRequest) SetClientToken(v string) *DeleteIpamPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetDryRun(v bool) *DeleteIpamPoolRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetIpamPoolId(v string) *DeleteIpamPoolRequest {
	s.IpamPoolId = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetOwnerAccount(v string) *DeleteIpamPoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetOwnerId(v int64) *DeleteIpamPoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetRegionId(v string) *DeleteIpamPoolRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetResourceOwnerAccount(v string) *DeleteIpamPoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamPoolRequest) SetResourceOwnerId(v int64) *DeleteIpamPoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpamPoolRequest) Validate() error {
	return dara.Validate(s)
}
