// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterMulticastDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AssociateTransitRouterMulticastDomainRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateTransitRouterMulticastDomainRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *AssociateTransitRouterMulticastDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateTransitRouterMulticastDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssociateTransitRouterMulticastDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateTransitRouterMulticastDomainRequest
	GetResourceOwnerId() *int64
	SetTransitRouterAttachmentId(v string) *AssociateTransitRouterMulticastDomainRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterMulticastDomainId(v string) *AssociateTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainId() *string
	SetVSwitchIds(v []*string) *AssociateTransitRouterMulticastDomainRequest
	GetVSwitchIds() []*string
}

type AssociateTransitRouterMulticastDomainRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC connection.
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

func (s AssociateTransitRouterMulticastDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterMulticastDomainRequest) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *AssociateTransitRouterMulticastDomainRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetClientToken(v string) *AssociateTransitRouterMulticastDomainRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetDryRun(v bool) *AssociateTransitRouterMulticastDomainRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetOwnerAccount(v string) *AssociateTransitRouterMulticastDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetOwnerId(v int64) *AssociateTransitRouterMulticastDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetResourceOwnerAccount(v string) *AssociateTransitRouterMulticastDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetResourceOwnerId(v int64) *AssociateTransitRouterMulticastDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetTransitRouterAttachmentId(v string) *AssociateTransitRouterMulticastDomainRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainId(v string) *AssociateTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) SetVSwitchIds(v []*string) *AssociateTransitRouterMulticastDomainRequest {
	s.VSwitchIds = v
	return s
}

func (s *AssociateTransitRouterMulticastDomainRequest) Validate() error {
	return dara.Validate(s)
}
