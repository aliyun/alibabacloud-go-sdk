// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterMulticastDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterMulticastDomainRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterMulticastDomainRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterMulticastDomainRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterMulticastDomainRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterMulticastDomainRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterMulticastDomainRequest
	GetResourceOwnerId() *int64
	SetTransitRouterMulticastDomainId(v string) *DeleteTransitRouterMulticastDomainRequest
	GetTransitRouterMulticastDomainId() *string
}

type DeleteTransitRouterMulticastDomainRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// The ID of the multicast domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-mcast-domain-40cwj0rgzgdtam****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
}

func (s DeleteTransitRouterMulticastDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterMulticastDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterMulticastDomainRequest) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetClientToken(v string) *DeleteTransitRouterMulticastDomainRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetDryRun(v bool) *DeleteTransitRouterMulticastDomainRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetOwnerAccount(v string) *DeleteTransitRouterMulticastDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetOwnerId(v int64) *DeleteTransitRouterMulticastDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterMulticastDomainRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterMulticastDomainRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) SetTransitRouterMulticastDomainId(v string) *DeleteTransitRouterMulticastDomainRequest {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainRequest) Validate() error {
	return dara.Validate(s)
}
