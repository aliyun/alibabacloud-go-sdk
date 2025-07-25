// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateIpamResourceDiscoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DissociateIpamResourceDiscoveryRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateIpamResourceDiscoveryRequest
	GetDryRun() *bool
	SetIpamId(v string) *DissociateIpamResourceDiscoveryRequest
	GetIpamId() *string
	SetIpamResourceDiscoveryId(v string) *DissociateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryId() *string
	SetOwnerAccount(v string) *DissociateIpamResourceDiscoveryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DissociateIpamResourceDiscoveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DissociateIpamResourceDiscoveryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DissociateIpamResourceDiscoveryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DissociateIpamResourceDiscoveryRequest
	GetResourceOwnerId() *int64
}

type DissociateIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: Performs a dry run without disassociating the resource discovery and IPAM instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and sends the request. After the request passes the check, an HTTP 2xx status code is returned and the resource discovery and IPAM instances are disassociated.
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
	// The ID of the resource discovery instance.
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

func (s DissociateIpamResourceDiscoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *DissociateIpamResourceDiscoveryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateIpamResourceDiscoveryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateIpamResourceDiscoveryRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *DissociateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *DissociateIpamResourceDiscoveryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DissociateIpamResourceDiscoveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DissociateIpamResourceDiscoveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateIpamResourceDiscoveryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DissociateIpamResourceDiscoveryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DissociateIpamResourceDiscoveryRequest) SetClientToken(v string) *DissociateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetDryRun(v bool) *DissociateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetIpamId(v string) *DissociateIpamResourceDiscoveryRequest {
	s.IpamId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryId(v string) *DissociateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *DissociateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *DissociateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetRegionId(v string) *DissociateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *DissociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *DissociateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryRequest) Validate() error {
	return dara.Validate(s)
}
