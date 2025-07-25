// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpamRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpamRequest
	GetDryRun() *bool
	SetIpamId(v string) *DeleteIpamRequest
	GetIpamId() *string
	SetOwnerAccount(v string) *DeleteIpamRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpamRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpamRequest
	GetResourceOwnerId() *int64
}

type DeleteIpamRequest struct {
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
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId       *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
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

func (s DeleteIpamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpamRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpamRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *DeleteIpamRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpamRequest) SetClientToken(v string) *DeleteIpamRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamRequest) SetDryRun(v bool) *DeleteIpamRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamRequest) SetIpamId(v string) *DeleteIpamRequest {
	s.IpamId = &v
	return s
}

func (s *DeleteIpamRequest) SetOwnerAccount(v string) *DeleteIpamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamRequest) SetOwnerId(v int64) *DeleteIpamRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamRequest) SetRegionId(v string) *DeleteIpamRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamRequest) SetResourceOwnerAccount(v string) *DeleteIpamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamRequest) SetResourceOwnerId(v int64) *DeleteIpamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpamRequest) Validate() error {
	return dara.Validate(s)
}
