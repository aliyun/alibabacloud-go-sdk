// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterTransitRouterMulticastGroupSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetClientToken() *string
	SetDryRun(v bool) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetDryRun() *bool
	SetGroupIpAddress(v string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetGroupIpAddress() *string
	SetNetworkInterfaceIds(v []*string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetNetworkInterfaceIds() []*string
	SetOwnerAccount(v string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetResourceOwnerId() *int64
	SetTransitRouterMulticastDomainId(v string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetTransitRouterMulticastDomainId() *string
	SetVpcId(v string) *RegisterTransitRouterMulticastGroupSourcesRequest
	GetVpcId() *string
}

type RegisterTransitRouterMulticastGroupSourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID is different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Default values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP address of the multicast group to which the multicast source belongs. Valid values: **224.0.1.0*	- to **239.255.255.254**.
	//
	// > 224.0.0.0 to 224.0.0.255 are reserved by the system and cannot be used by multicast groups.
	//
	// If the multicast group does not exist in the multicast domain, the system automatically creates the multicast group in the multicast domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// The IDs of the ENIs.
	NetworkInterfaceIds  []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the multicast domain to which the multicast source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-mcast-domain-5mjb5gjb6dgu98****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The ID of the VPC to which the ENI belongs.
	//
	// 	- If the ENI belongs to the current Alibaba Cloud account, ignore this parameter.
	//
	// 	- If the ENI belongs to a different Alibaba Cloud account, you must set this parameter.
	//
	// example:
	//
	// vpc-wz9fusm6zq8uy7cfa****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RegisterTransitRouterMulticastGroupSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterTransitRouterMulticastGroupSourcesRequest) GoString() string {
	return s.String()
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetGroupIpAddress() *string {
	return s.GroupIpAddress
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetClientToken(v string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetDryRun(v bool) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.DryRun = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetGroupIpAddress(v string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.GroupIpAddress = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetNetworkInterfaceIds(v []*string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetOwnerAccount(v string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetOwnerId(v int64) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetResourceOwnerAccount(v string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetResourceOwnerId(v int64) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetTransitRouterMulticastDomainId(v string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) SetVpcId(v string) *RegisterTransitRouterMulticastGroupSourcesRequest {
	s.VpcId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesRequest) Validate() error {
	return dara.Validate(s)
}
