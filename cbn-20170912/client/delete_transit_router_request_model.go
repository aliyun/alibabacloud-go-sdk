// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterRequest
	GetResourceOwnerId() *int64
	SetTransitRouterId(v string) *DeleteTransitRouterRequest
	GetTransitRouterId() *string
}

type DeleteTransitRouterRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether only to precheck the request. Check items include permissions and the status of the transit router. Valid values:
	//
	// 	- **false*	- (default): sends the request. If the request passes the precheck, the transit router is deleted.
	//
	// 	- **true**: prechecks the request but does not delete the transit router. If you use this value, the system checks the required parameters and the request syntax. If the request fails to pass the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DeleteTransitRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DeleteTransitRouterRequest) SetClientToken(v string) *DeleteTransitRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterRequest) SetDryRun(v bool) *DeleteTransitRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterRequest) SetOwnerAccount(v string) *DeleteTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRequest) SetOwnerId(v int64) *DeleteTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterRequest) SetTransitRouterId(v string) *DeleteTransitRouterRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DeleteTransitRouterRequest) Validate() error {
	return dara.Validate(s)
}
