// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterTransitRouterMulticastGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetDryRun() *bool
	SetGroupIpAddress(v string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetGroupIpAddress() *string
	SetNetworkInterfaceIds(v []*string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetNetworkInterfaceIds() []*string
	SetOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetOwnerId() *int64
	SetPeerTransitRouterMulticastDomains(v []*string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetPeerTransitRouterMulticastDomains() []*string
	SetResourceOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetResourceOwnerId() *int64
	SetTransitRouterMulticastDomainId(v string) *DeregisterTransitRouterMulticastGroupMembersRequest
	GetTransitRouterMulticastDomainId() *string
}

type DeregisterTransitRouterMulticastGroupMembersRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a unique value from your client for each request. The ClientToken parameter supports only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Sends a check request without deleting the multicast member. The system checks required parameters, the request format, and service limits. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After the request passes the check, the multicast member is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP address of the multicast group to which the multicast member belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// A list of Elastic Network Interface (ENI) IDs.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A list of inter-region multicast domain IDs.
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
}

func (s DeregisterTransitRouterMulticastGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s DeregisterTransitRouterMulticastGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetGroupIpAddress() *string {
	return s.GroupIpAddress
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetPeerTransitRouterMulticastDomains() []*string {
	return s.PeerTransitRouterMulticastDomains
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetClientToken(v string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.ClientToken = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetDryRun(v bool) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.DryRun = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetGroupIpAddress(v string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.GroupIpAddress = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetNetworkInterfaceIds(v []*string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetOwnerId(v int64) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetPeerTransitRouterMulticastDomains(v []*string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.PeerTransitRouterMulticastDomains = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetResourceOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetResourceOwnerId(v int64) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) SetTransitRouterMulticastDomainId(v string) *DeregisterTransitRouterMulticastGroupMembersRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersRequest) Validate() error {
	return dara.Validate(s)
}
