// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTablePropagationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteTablePropagationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterRouteTablePropagationsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterRouteTablePropagationsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetStatus() *string
	SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterAttachmentResourceId(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetTransitRouterAttachmentResourceId() *string
	SetTransitRouterAttachmentResourceType(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetTransitRouterAttachmentResourceType() *string
	SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablePropagationsRequest
	GetTransitRouterRouteTableId() *string
}

type ListTransitRouterRouteTablePropagationsRequest struct {
	// The number of entries to return on each page. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query.
	//
	// example:
	//
	// dd20****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the route learning correlation. Valid values:
	//
	// 	- **Active**: available
	//
	// 	- **Enabling**: being enabled
	//
	// 	- **Disabling**: being disabled
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the network instance connection.
	//
	// example:
	//
	// tr-attach-vx6iwhjr1x1j78****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-bp1h8vbrbcgohcju5****
	TransitRouterAttachmentResourceId *string `json:"TransitRouterAttachmentResourceId,omitempty" xml:"TransitRouterAttachmentResourceId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **TR**: transit router
	//
	// 	- **VPN**: VPN connection
	//
	// example:
	//
	// VPC
	TransitRouterAttachmentResourceType *string `json:"TransitRouterAttachmentResourceType,omitempty" xml:"TransitRouterAttachmentResourceType,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteTablePropagationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetTransitRouterAttachmentResourceId() *string {
	return s.TransitRouterAttachmentResourceId
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetTransitRouterAttachmentResourceType() *string {
	return s.TransitRouterAttachmentResourceType
}

func (s *ListTransitRouterRouteTablePropagationsRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetMaxResults(v int32) *ListTransitRouterRouteTablePropagationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetNextToken(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetOwnerAccount(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetOwnerId(v int64) *ListTransitRouterRouteTablePropagationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteTablePropagationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetStatus(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.Status = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetTransitRouterAttachmentResourceId(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.TransitRouterAttachmentResourceId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetTransitRouterAttachmentResourceType(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.TransitRouterAttachmentResourceType = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablePropagationsRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsRequest) Validate() error {
	return dara.Validate(s)
}
