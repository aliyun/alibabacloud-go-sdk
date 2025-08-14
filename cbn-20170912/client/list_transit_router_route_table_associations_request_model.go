// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTableAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteTableAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterRouteTableAssociationsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterRouteTableAssociationsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetStatus() *string
	SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentResourceId(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetTransitRouterAttachmentResourceId() *string
	SetTransitRouterAttachmentResourceType(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetTransitRouterAttachmentResourceType() *string
	SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTableAssociationsRequest
	GetTransitRouterRouteTableId() *string
}

type ListTransitRouterRouteTableAssociationsRequest struct {
	// The number of entries to return on each page. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query or no subsequent query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// a415****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the associated forwarding correlation. Valid values:
	//
	// 	- **Active**: The associated forwarding correlation is available.
	//
	// 	- **Associating**: The associated forwarding correlation is being created.
	//
	// 	- **Dissociating**: The associated forwarding correlation is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	TransitRouterAttachmentResourceId *string `json:"TransitRouterAttachmentResourceId,omitempty" xml:"TransitRouterAttachmentResourceId,omitempty"`
	// The type of next hop. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **TR**: transit router
	//
	// 	- **VPN**: VPN attachment
	//
	// example:
	//
	// VPC
	TransitRouterAttachmentResourceType *string `json:"TransitRouterAttachmentResourceType,omitempty" xml:"TransitRouterAttachmentResourceType,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTableAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetTransitRouterAttachmentResourceId() *string {
	return s.TransitRouterAttachmentResourceId
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetTransitRouterAttachmentResourceType() *string {
	return s.TransitRouterAttachmentResourceType
}

func (s *ListTransitRouterRouteTableAssociationsRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetMaxResults(v int32) *ListTransitRouterRouteTableAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetNextToken(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetOwnerAccount(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetOwnerId(v int64) *ListTransitRouterRouteTableAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteTableAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetStatus(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.Status = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetTransitRouterAttachmentResourceId(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.TransitRouterAttachmentResourceId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetTransitRouterAttachmentResourceType(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.TransitRouterAttachmentResourceType = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTableAssociationsRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
