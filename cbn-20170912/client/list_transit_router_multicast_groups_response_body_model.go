// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterMulticastGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterMulticastGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterMulticastGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterMulticastGroupsResponseBody
	GetTotalCount() *int32
	SetTransitRouterMulticastGroups(v []*ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) *ListTransitRouterMulticastGroupsResponseBody
	GetTransitRouterMulticastGroups() []*ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups
}

type ListTransitRouterMulticastGroupsResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If **NextToken*	- was not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FB3C4A16-0933-5850-9D43-0C3EA37BCBFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of multicast groups.
	TransitRouterMulticastGroups []*ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups `json:"TransitRouterMulticastGroups,omitempty" xml:"TransitRouterMulticastGroups,omitempty" type:"Repeated"`
}

func (s ListTransitRouterMulticastGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterMulticastGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterMulticastGroupsResponseBody) GetTransitRouterMulticastGroups() []*ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	return s.TransitRouterMulticastGroups
}

func (s *ListTransitRouterMulticastGroupsResponseBody) SetMaxResults(v int32) *ListTransitRouterMulticastGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBody) SetNextToken(v string) *ListTransitRouterMulticastGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBody) SetRequestId(v string) *ListTransitRouterMulticastGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBody) SetTotalCount(v int32) *ListTransitRouterMulticastGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBody) SetTransitRouterMulticastGroups(v []*ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) *ListTransitRouterMulticastGroupsResponseBody {
	s.TransitRouterMulticastGroups = v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBody) Validate() error {
	if s.TransitRouterMulticastGroups != nil {
		for _, item := range s.TransitRouterMulticastGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups struct {
	// The IP address of the multicast group to which the multicast resource belongs.
	//
	// example:
	//
	// 239.XX.XX.2
	GroupIpAddress *string `json:"GroupIpAddress,omitempty" xml:"GroupIpAddress,omitempty"`
	// Indicates whether the multicast resource is a multicast member. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	GroupMember *bool `json:"GroupMember,omitempty" xml:"GroupMember,omitempty"`
	// Indicates whether the multicast resource is a multicast source. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	GroupSource *bool `json:"GroupSource,omitempty" xml:"GroupSource,omitempty"`
	// The type of the multicast source.
	//
	// If the value is **Static**, the multicast source is manually specified.
	//
	// example:
	//
	// Static
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	// The ID of the ENI, which is a multicast resource.
	//
	// example:
	//
	// eni-p0weuda3lszwzjly****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The ID of the multicast domain associated with the multicast resource that is deployed across regions.
	//
	// example:
	//
	// tr-mcast-domain-91wpg6wbhchjeq****
	PeerTransitRouterMulticastDomainId *string `json:"PeerTransitRouterMulticastDomainId,omitempty" xml:"PeerTransitRouterMulticastDomainId,omitempty"`
	// The ID of the resource associated with the multicast resource.
	//
	// example:
	//
	// vpc-p0w9alkte4w2htrqe****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the multicast resource belongs.
	//
	// example:
	//
	// 253460731706911258
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the multicast resource. Valid values:
	//
	// 	- **VPC**: The multicast resource is in a VPC.
	//
	// 	- **TR**: The multicast resource is deployed across regions.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type of the multicast member.
	//
	// If the value is **Static**, the multicast member is manually specified.
	//
	// example:
	//
	// Static
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the multicast resource. Valid values:
	//
	// 	- **Registering**: being created
	//
	// 	- **Registered**: available
	//
	// 	- **Deregistering**: being deleted
	//
	// example:
	//
	// Registered
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-g3kz2k3u76amsk****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the multicast domain.
	//
	// example:
	//
	// tr-mcast-domain-kx0vk0v7fz8kx4****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The ID of the vSwitch to which the multicast resource belongs.
	//
	// example:
	//
	// vsw-p0w9s2ig1jnwgrbzl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetGroupIpAddress() *string {
	return s.GroupIpAddress
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetGroupMember() *bool {
	return s.GroupMember
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetGroupSource() *bool {
	return s.GroupSource
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetMemberType() *string {
	return s.MemberType
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetPeerTransitRouterMulticastDomainId() *string {
	return s.PeerTransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetSourceType() *string {
	return s.SourceType
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetGroupIpAddress(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.GroupIpAddress = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetGroupMember(v bool) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.GroupMember = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetGroupSource(v bool) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.GroupSource = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetMemberType(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.MemberType = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetNetworkInterfaceId(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetPeerTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.PeerTransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetResourceId(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetResourceOwnerId(v int64) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetResourceType(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetSourceType(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.SourceType = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetStatus(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.Status = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetTransitRouterAttachmentId(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) SetVSwitchId(v string) *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups {
	s.VSwitchId = &v
	return s
}

func (s *ListTransitRouterMulticastGroupsResponseBodyTransitRouterMulticastGroups) Validate() error {
	return dara.Validate(s)
}
