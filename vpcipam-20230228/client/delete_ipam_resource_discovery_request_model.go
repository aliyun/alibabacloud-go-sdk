// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamResourceDiscoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteIpamResourceDiscoveryRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteIpamResourceDiscoveryRequest
	GetDryRun() *bool
	SetIpamResourceDiscoveryId(v string) *DeleteIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryId() *string
	SetOwnerAccount(v string) *DeleteIpamResourceDiscoveryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteIpamResourceDiscoveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIpamResourceDiscoveryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteIpamResourceDiscoveryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteIpamResourceDiscoveryRequest
	GetResourceOwnerId() *int64
}

type DeleteIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform the dry run. Valid values:
	//
	// 	- **true**: Performs a dry run without deleting the resource discovery instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and the actual request. If the request passes the check, an HTTP 2xx status code is returned and the resource discovery instance is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
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

func (s DeleteIpamResourceDiscoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpamResourceDiscoveryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteIpamResourceDiscoveryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *DeleteIpamResourceDiscoveryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteIpamResourceDiscoveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIpamResourceDiscoveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIpamResourceDiscoveryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteIpamResourceDiscoveryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteIpamResourceDiscoveryRequest) SetClientToken(v string) *DeleteIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetDryRun(v bool) *DeleteIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *DeleteIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *DeleteIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetOwnerId(v int64) *DeleteIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetRegionId(v string) *DeleteIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *DeleteIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *DeleteIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryRequest) Validate() error {
	return dara.Validate(s)
}
