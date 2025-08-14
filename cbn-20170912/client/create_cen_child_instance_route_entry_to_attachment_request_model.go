// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenChildInstanceRouteEntryToAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetCenId() *string
	SetClientToken(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetClientToken() *string
	SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetDestinationCidrBlock() *string
	SetDryRun(v bool) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetRouteTableId() *string
	SetTransitRouterAttachmentId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest
	GetTransitRouterAttachmentId() *string
}

type CreateCenChildInstanceRouteEntryToAttachmentRequest struct {
	// The CEN instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1n6cbxcszp55vxo****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block of the route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Specifies whether to perform a dry run to check information such as the permissions and the instance status. Valid values:
	//
	// 	- **false**: performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and request syntax. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// >  This parameter is not in use.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route table configured on the network instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vrt-bp1msipdczo9lejup****
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

func (s CreateCenChildInstanceRouteEntryToAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenChildInstanceRouteEntryToAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetCenId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetClientToken(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetDestinationCidrBlock(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetDryRun(v bool) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetOwnerId(v int64) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerAccount(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetResourceOwnerId(v int64) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetRouteTableId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.RouteTableId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) SetTransitRouterAttachmentId(v string) *CreateCenChildInstanceRouteEntryToAttachmentRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *CreateCenChildInstanceRouteEntryToAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
