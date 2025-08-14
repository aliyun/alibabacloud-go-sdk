// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterTransitRouterMulticastGroupSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetDryRun() *bool
	SetGroupIpAddress(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetGroupIpAddress() *string
	SetNetworkInterfaceIds(v []*string) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetNetworkInterfaceIds() []*string
	SetOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetResourceOwnerId() *int64
	SetTransitRouterMulticastDomainId(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest
	GetTransitRouterMulticastDomainId() *string
}

type DeregisterTransitRouterMulticastGroupSourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP address of the multicast group to which the multicast source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// A list of multicast source IDs.
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
	// tr-mcast-domain-91wpg6wbhchjeq****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
}

func (s DeregisterTransitRouterMulticastGroupSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeregisterTransitRouterMulticastGroupSourcesRequest) GoString() string {
	return s.String()
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetGroupIpAddress() *string {
	return s.GroupIpAddress
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetClientToken(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetDryRun(v bool) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.DryRun = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetGroupIpAddress(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.GroupIpAddress = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetNetworkInterfaceIds(v []*string) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetOwnerId(v int64) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetResourceOwnerAccount(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetResourceOwnerId(v int64) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) SetTransitRouterMulticastDomainId(v string) *DeregisterTransitRouterMulticastGroupSourcesRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesRequest) Validate() error {
	return dara.Validate(s)
}
