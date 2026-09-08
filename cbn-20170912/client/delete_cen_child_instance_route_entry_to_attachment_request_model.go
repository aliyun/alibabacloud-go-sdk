// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenChildInstanceRouteEntryToAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetCenId() *string
	SetClientToken(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetClientToken() *string
	SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetRouteTableId() *string
	SetTransitRouterAttachmentId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type DeleteCenChildInstanceRouteEntryToAttachmentRequest struct {
	// The instance ID of the Cloud Enterprise Network (CEN).
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-wgcl0ik5o8jakq****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block of the route entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.1.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform a dry run. The dry run checks parameter validity, user permissions, and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the route entry of the network instance is deleted.
	//
	// - **true**: Sends a check request. Only the check is performed. The route entry of the network instance is not deleted. The check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check succeeds, the corresponding request ID is returned.
	//
	// > This parameter is not yet available.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route table ID of the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1osd9opvegfpowc****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The ID of the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-r1qhupkc19iadz****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenChildInstanceRouteEntryToAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetCenId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetClientToken(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetDestinationCidrBlock(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetDryRun(v bool) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerId(v int64) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetRouteTableId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) SetTransitRouterAttachmentId(v string) *DeleteCenChildInstanceRouteEntryToAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DeleteCenChildInstanceRouteEntryToAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
