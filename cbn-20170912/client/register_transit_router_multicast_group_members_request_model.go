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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, the multicast member is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP address of the multicast group to which the multicast member belongs. Valid values: **224.0.1.0*	- to **239.255.255.254**.
	//
	// 	Notice: 224.0.0.0 to 224.0.0.127 are system reserved IP addresses and cannot be used as multicast group IP addresses.
	//
	// If the multicast group that you specify does not exist in the current multicast domain, the system automatically creates a new multicast group in the current multicast domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// The list of elastic network interfaces (ENIs) IDs.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of cross-region multicast domain IDs.
	PeerTransitRouterMulticastDomains []*string `json:"PeerTransitRouterMulticastDomains,omitempty" xml:"PeerTransitRouterMulticastDomains,omitempty" type:"Repeated"`
	ResourceOwnerAccount              *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                   *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the multicast domain to which the multicast member belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-mcast-domain-91wpg6wbhchjeq****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The instance ID of the VPC-connected instance to which the elastic network interfaces (ENIs) belongs.
	//
	// - If the network interface controller (NIC) belongs to the same Alibaba Cloud account that you use to logon to call this operation, you do not need to set this parameter.
	//
	// - If the network interface controller (NIC) belongs to a different Alibaba Cloud account from the one you use to logon, you must set this parameter.
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
