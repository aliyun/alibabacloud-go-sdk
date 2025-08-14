// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateTransitRouterMulticastDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisassociateTransitRouterMulticastDomainRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisassociateTransitRouterMulticastDomainRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DisassociateTransitRouterMulticastDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisassociateTransitRouterMulticastDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisassociateTransitRouterMulticastDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisassociateTransitRouterMulticastDomainRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *DisassociateTransitRouterMulticastDomainRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterMulticastDomainId(v string) *DisassociateTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainId() *string
	SetVSwitchIds(v []*string) *DisassociateTransitRouterMulticastDomainRequest
	GetVSwitchIds() []*string
}

type DisassociateTransitRouterMulticastDomainRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the API request. Valid values:
	//
	// - **true**: prechecks the request but does not disassociate the vSwitch from the multicast domain. The system checks the required parameters, the request format, and the service limits. If the request fails the check, an error message is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends the request. The vSwitch is disassociated from the multicast domain after the request passes the precheck.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the virtual private cloud (VPC) connection.
	//
	// A connection ID is generated after a VPC to which the vSwitch belongs is connected to a transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-g3kz2k3u76amsk****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the multicast domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-mcast-domain-91wpg6wbhchjeq****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
	// The IDs of vSwitches.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s DisassociateTransitRouterMulticastDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateTransitRouterMulticastDomainRequest) GoString() string {
	return s.String()
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *DisassociateTransitRouterMulticastDomainRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetClientToken(v string) *DisassociateTransitRouterMulticastDomainRequest {
	s.ClientToken = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetDryRun(v bool) *DisassociateTransitRouterMulticastDomainRequest {
	s.DryRun = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetOwnerAccount(v string) *DisassociateTransitRouterMulticastDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetOwnerId(v int64) *DisassociateTransitRouterMulticastDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetResourceOwnerAccount(v string) *DisassociateTransitRouterMulticastDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetResourceOwnerId(v int64) *DisassociateTransitRouterMulticastDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetTransitRouterAttachmentId(v string) *DisassociateTransitRouterMulticastDomainRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainId(v string) *DisassociateTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) SetVSwitchIds(v []*string) *DisassociateTransitRouterMulticastDomainRequest {
	s.VSwitchIds = v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainRequest) Validate() error {
	return dara.Validate(s)
}
