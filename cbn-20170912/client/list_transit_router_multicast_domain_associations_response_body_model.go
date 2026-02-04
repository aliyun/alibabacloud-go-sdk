// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterMulticastDomainAssociationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterMulticastDomainAssociationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterMulticastDomainAssociationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterMulticastDomainAssociationsResponseBody
	GetTotalCount() *int32
	SetTransitRouterMulticastAssociations(v []*ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) *ListTransitRouterMulticastDomainAssociationsResponseBody
	GetTransitRouterMulticastAssociations() []*ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations
}

type ListTransitRouterMulticastDomainAssociationsResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If **NextToken*	- is empty, it indicates that no subsequent query is to be sent.
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
	// 1CD0969B-A605-5D2D-BFF0-699FD182FB7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the resource associated with the multicast domain.
	TransitRouterMulticastAssociations []*ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations `json:"TransitRouterMulticastAssociations,omitempty" xml:"TransitRouterMulticastAssociations,omitempty" type:"Repeated"`
}

func (s ListTransitRouterMulticastDomainAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) GetTransitRouterMulticastAssociations() []*ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	return s.TransitRouterMulticastAssociations
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) SetMaxResults(v int32) *ListTransitRouterMulticastDomainAssociationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) SetNextToken(v string) *ListTransitRouterMulticastDomainAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) SetRequestId(v string) *ListTransitRouterMulticastDomainAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) SetTotalCount(v int32) *ListTransitRouterMulticastDomainAssociationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) SetTransitRouterMulticastAssociations(v []*ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) *ListTransitRouterMulticastDomainAssociationsResponseBody {
	s.TransitRouterMulticastAssociations = v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBody) Validate() error {
	if s.TransitRouterMulticastAssociations != nil {
		for _, item := range s.TransitRouterMulticastAssociations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations struct {
	// The ID of the resource associated with the multicast domain.
	//
	// example:
	//
	// vpc-p0w9b7g9l90yofr0n****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource associated with the multicast domain belongs.
	//
	// example:
	//
	// 1210123456123456
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource associated with the multicast domain.
	//
	// Valid value: **VPC**.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The association status. Valid values:
	//
	// 	- **Associated**: The resource is associated with the multicast domain.
	//
	// 	- **Associating**: The resource is being associated with the multicast domain.
	//
	// 	- **Dissociating**: The resource is being disassociated from the multicast domain.
	//
	// example:
	//
	// Dissociating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-p90y3ymbbwuvy5****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the multicast domain.
	//
	// example:
	//
	// tr-mcast-domain-91wpg6wbhchjeq****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-p0wxk12u6okfkr8xy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetResourceId(v string) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetResourceType(v string) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetStatus(v string) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.Status = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetTransitRouterAttachmentId(v string) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) SetVSwitchId(v string) *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations {
	s.VSwitchId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsResponseBodyTransitRouterMulticastAssociations) Validate() error {
	return dara.Validate(s)
}
