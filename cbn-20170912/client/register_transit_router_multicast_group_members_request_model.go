// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterTransitRouterMulticastGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetClientToken() *string
	SetDryRun(v bool) *RegisterTransitRouterMulticastGroupMembersRequest
	GetDryRun() *bool
	SetGroupIpAddress(v string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetGroupIpAddress() *string
	SetNetworkInterfaceIds(v []*string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetNetworkInterfaceIds() []*string
	SetOwnerAccount(v string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RegisterTransitRouterMulticastGroupMembersRequest
	GetOwnerId() *int64
	SetPeerTransitRouterMulticastDomains(v []*string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetPeerTransitRouterMulticastDomains() []*string
	SetResourceOwnerAccount(v string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RegisterTransitRouterMulticastGroupMembersRequest
	GetResourceOwnerId() *int64
	SetTransitRouterMulticastDomainId(v string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetTransitRouterMulticastDomainId() *string
	SetVpcId(v string) *RegisterTransitRouterMulticastGroupMembersRequest
	GetVpcId() *string
}

type RegisterTransitRouterMulticastGroupMembersRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP address of the multicast group to which the multicast members belong. Valid values: **224.0.0.1*	- to **239.255.255.254**.
	//
	// If the multicast group does not exist in the specified multicast domain, the system automatically creates the multicast group in the multicast domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// The IDs of the ENIs.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of inter-region multicast domains.
	PeerTransitRouterMulticastDomains []*string `json:"PeerTransitRouterMulticastDomains,omitempty" xml:"PeerTransitRouterMulticastDomains,omitempty" type:"Repeated"`
	ResourceOwnerAccount              *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                   *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the multicast domain to which the multicast members belong.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-mcast-domain-91wpg6wbhchjeq****
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

func (s RegisterTransitRouterMulticastGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterTransitRouterMulticastGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetGroupIpAddress() *string {
	return s.GroupIpAddress
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetPeerTransitRouterMulticastDomains() []*string {
	return s.PeerTransitRouterMulticastDomains
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetClientToken(v string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetDryRun(v bool) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.DryRun = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetGroupIpAddress(v string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.GroupIpAddress = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetNetworkInterfaceIds(v []*string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetOwnerAccount(v string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetOwnerId(v int64) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetPeerTransitRouterMulticastDomains(v []*string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.PeerTransitRouterMulticastDomains = v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetResourceOwnerAccount(v string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetResourceOwnerId(v int64) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetTransitRouterMulticastDomainId(v string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) SetVpcId(v string) *RegisterTransitRouterMulticastGroupMembersRequest {
	s.VpcId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
