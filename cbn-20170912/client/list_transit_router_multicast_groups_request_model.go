// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListTransitRouterMulticastGroupsRequest
	GetClientToken() *string
	SetGroupIpAddress(v string) *ListTransitRouterMulticastGroupsRequest
	GetGroupIpAddress() *string
	SetIsGroupMember(v bool) *ListTransitRouterMulticastGroupsRequest
	GetIsGroupMember() *bool
	SetIsGroupSource(v bool) *ListTransitRouterMulticastGroupsRequest
	GetIsGroupSource() *bool
	SetMaxResults(v int64) *ListTransitRouterMulticastGroupsRequest
	GetMaxResults() *int64
	SetNetworkInterfaceIds(v []*string) *ListTransitRouterMulticastGroupsRequest
	GetNetworkInterfaceIds() []*string
	SetNextToken(v string) *ListTransitRouterMulticastGroupsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterMulticastGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterMulticastGroupsRequest
	GetOwnerId() *int64
	SetPeerTransitRouterMulticastDomains(v []*string) *ListTransitRouterMulticastGroupsRequest
	GetPeerTransitRouterMulticastDomains() []*string
	SetResourceId(v string) *ListTransitRouterMulticastGroupsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterMulticastGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterMulticastGroupsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ListTransitRouterMulticastGroupsRequest
	GetResourceType() *string
	SetTransitRouterAttachmentId(v string) *ListTransitRouterMulticastGroupsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastGroupsRequest
	GetTransitRouterMulticastDomainId() *string
	SetVSwitchIds(v []*string) *ListTransitRouterMulticastGroupsRequest
	GetVSwitchIds() []*string
}

type ListTransitRouterMulticastGroupsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IP address of the multicast group.
	//
	// Each multicast group is identified by a multicast IP address.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// Specifies whether to query multicast members. Valid values:
	//
	// - **false**: no.
	//
	// - **true**: yes.
	//
	// > This parameter is used together with IsGroupSource.
	//
	// > - If neither parameter is configured, both multicast sources and members are queried by default.
	//
	// > - If only one parameter is configured or both are configured, the query is based on the configured parameters.
	//
	// example:
	//
	// false
	IsGroupMember *bool `json:"IsGroupMember,omitempty" xml:"IsGroupMember,omitempty"`
	// Specifies whether to query multicast sources. Valid values:
	//
	// - **false**: no.
	//
	// - **true**: yes.
	//
	// > This parameter is used together with IsGroupMember.
	//
	// > - If neither parameter is configured, both multicast sources and members are queried by default.
	//
	// > - If only one parameter is configured or both are configured, the query is based on the configured parameters.
	//
	// example:
	//
	// true
	IsGroupSource *bool `json:"IsGroupSource,omitempty" xml:"IsGroupSource,omitempty"`
	// The number of entries per page for a paged query. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The list of elastic network interface (ENI) IDs.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - You do not need to specify this parameter for the first request or if no next query exists.
	//
	// - If a next query exists, set the value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The list of cross-region multicast domain IDs.
	PeerTransitRouterMulticastDomains []*string `json:"PeerTransitRouterMulticastDomains,omitempty" xml:"PeerTransitRouterMulticastDomains,omitempty" type:"Repeated"`
	// The resource ID associated with the multicast resource.
	//
	// example:
	//
	// vpc-p0w9alkte4w2htrqe****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the multicast resource. Valid values:
	//
	// - **VPC**: queries multicast resources in a virtual private cloud (VPC).
	//
	// - **TR**: queries cross-region multicast resources.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The network instance connection ID.
	//
	// You must specify at least one of TransitRouterMulticastDomainId and TransitRouterAttachmentId.
	//
	// example:
	//
	// tr-attach-g3kz2k3u76amsk****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The multicast domain ID.
	//
	// You must specify at least one of TransitRouterMulticastDomainId and TransitRouterAttachmentId.
	//
	// example:
	//
	// tr-mcast-domain-5mjb5gjb6dgu98****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ListTransitRouterMulticastGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTransitRouterMulticastGroupsRequest) GetGroupIpAddress() *string {
	return s.GroupIpAddress
}

func (s *ListTransitRouterMulticastGroupsRequest) GetIsGroupMember() *bool {
	return s.IsGroupMember
}

func (s *ListTransitRouterMulticastGroupsRequest) GetIsGroupSource() *bool {
	return s.IsGroupSource
}

func (s *ListTransitRouterMulticastGroupsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastGroupsRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *ListTransitRouterMulticastGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterMulticastGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterMulticastGroupsRequest) GetPeerTransitRouterMulticastDomains() []*string {
	return s.PeerTransitRouterMulticastDomains
}

func (s *ListTransitRouterMulticastGroupsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTransitRouterMulticastGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterMulticastGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterMulticastGroupsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterMulticastGroupsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterMulticastGroupsRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastGroupsRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListTransitRouterMulticastGroupsRequest) SetClientToken(v string) *ListTransitRouterMulticastGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetGroupIpAddress(v string) *ListTransitRouterMulticastGroupsRequest {
	s.GroupIpAddress = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetIsGroupMember(v bool) *ListTransitRouterMulticastGroupsRequest {
	s.IsGroupMember = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetIsGroupSource(v bool) *ListTransitRouterMulticastGroupsRequest {
	s.IsGroupSource = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetMaxResults(v int64) *ListTransitRouterMulticastGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetNetworkInterfaceIds(v []*string) *ListTransitRouterMulticastGroupsRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetNextToken(v string) *ListTransitRouterMulticastGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetOwnerAccount(v string) *ListTransitRouterMulticastGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetOwnerId(v int64) *ListTransitRouterMulticastGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetPeerTransitRouterMulticastDomains(v []*string) *ListTransitRouterMulticastGroupsRequest {
	s.PeerTransitRouterMulticastDomains = v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetResourceId(v string) *ListTransitRouterMulticastGroupsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterMulticastGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetResourceOwnerId(v int64) *ListTransitRouterMulticastGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetResourceType(v string) *ListTransitRouterMulticastGroupsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterMulticastGroupsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastGroupsRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) SetVSwitchIds(v []*string) *ListTransitRouterMulticastGroupsRequest {
	s.VSwitchIds = v
	return s
}

func (s *ListTransitRouterMulticastGroupsRequest) Validate() error {
	return dara.Validate(s)
}
