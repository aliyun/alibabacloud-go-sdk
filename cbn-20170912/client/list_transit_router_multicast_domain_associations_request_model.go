// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetClientToken() *string
	SetMaxResults(v int64) *ListTransitRouterMulticastDomainAssociationsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterMulticastDomainAssociationsRequest
	GetOwnerId() *int64
	SetResourceId(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainAssociationsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetResourceType() *string
	SetTransitRouterAttachmentId(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetTransitRouterMulticastDomainId() *string
	SetVSwitchIds(v []*string) *ListTransitRouterMulticastDomainAssociationsRequest
	GetVSwitchIds() []*string
}

type ListTransitRouterMulticastDomainAssociationsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The token can only contain ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query and no next queries are to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the parameter to the value of NextToken that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource associated with the multicast domain.
	//
	// example:
	//
	// vpc-p0w9alkte4w2htrqe****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of resource associated with the multicast domain.
	//
	// Valid value: **VPC**.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// The IDs of vSwitches.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ListTransitRouterMulticastDomainAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetClientToken(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetMaxResults(v int64) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetNextToken(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetOwnerAccount(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetOwnerId(v int64) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetResourceId(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetResourceType(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetTransitRouterMulticastDomainId(v string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) SetVSwitchIds(v []*string) *ListTransitRouterMulticastDomainAssociationsRequest {
	s.VSwitchIds = v
	return s
}

func (s *ListTransitRouterMulticastDomainAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
