// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateIpamResourceDiscoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateIpamResourceDiscoveryRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateIpamResourceDiscoveryRequest
	GetDryRun() *bool
	SetIpamId(v string) *AssociateIpamResourceDiscoveryRequest
	GetIpamId() *string
	SetIpamResourceDiscoveryId(v string) *AssociateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryId() *string
	SetOwnerAccount(v string) *AssociateIpamResourceDiscoveryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateIpamResourceDiscoveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateIpamResourceDiscoveryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateIpamResourceDiscoveryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateIpamResourceDiscoveryRequest
	GetResourceOwnerId() *int64
}

type AssociateIpamResourceDiscoveryRequest struct {
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
	// 	- **true**: Performs a dry run without associating the resource discovery and IPAM instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request. After the request passes the check, an HTTP 2xx status code is returned and the resource discovery and IPAM instances are associated.
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
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
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

func (s AssociateIpamResourceDiscoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *AssociateIpamResourceDiscoveryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateIpamResourceDiscoveryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateIpamResourceDiscoveryRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *AssociateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *AssociateIpamResourceDiscoveryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateIpamResourceDiscoveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateIpamResourceDiscoveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateIpamResourceDiscoveryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateIpamResourceDiscoveryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateIpamResourceDiscoveryRequest) SetClientToken(v string) *AssociateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetDryRun(v bool) *AssociateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetIpamId(v string) *AssociateIpamResourceDiscoveryRequest {
	s.IpamId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *AssociateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *AssociateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *AssociateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetRegionId(v string) *AssociateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *AssociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *AssociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryRequest) Validate() error {
	return dara.Validate(s)
}
