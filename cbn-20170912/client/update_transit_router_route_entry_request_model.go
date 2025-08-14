// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTransitRouterRouteEntryRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTransitRouterRouteEntryRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTransitRouterRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTransitRouterRouteEntryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTransitRouterRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTransitRouterRouteEntryRequest
	GetResourceOwnerId() *int64
	SetTransitRouterRouteEntryDescription(v string) *UpdateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryDescription() *string
	SetTransitRouterRouteEntryId(v string) *UpdateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryId() *string
	SetTransitRouterRouteEntryName(v string) *UpdateTransitRouterRouteEntryRequest
	GetTransitRouterRouteEntryName() *string
}

type UpdateTransitRouterRouteEntryRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- is different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Default values:
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// 	- **true**: performs a dry run. The system checks the required parameters and request syntax. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new description of the route.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteEntryDescription *string `json:"TransitRouterRouteEntryDescription,omitempty" xml:"TransitRouterRouteEntryDescription,omitempty"`
	// The ID of the route.
	//
	// This parameter is required.
	//
	// example:
	//
	// rte-ksssq7kto4wfdx****
	TransitRouterRouteEntryId *string `json:"TransitRouterRouteEntryId,omitempty" xml:"TransitRouterRouteEntryId,omitempty"`
	// The new name of the route.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterRouteEntryName *string `json:"TransitRouterRouteEntryName,omitempty" xml:"TransitRouterRouteEntryName,omitempty"`
}

func (s UpdateTransitRouterRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteEntryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterRouteEntryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTransitRouterRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTransitRouterRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTransitRouterRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryDescription() *string {
	return s.TransitRouterRouteEntryDescription
}

func (s *UpdateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryId() *string {
	return s.TransitRouterRouteEntryId
}

func (s *UpdateTransitRouterRouteEntryRequest) GetTransitRouterRouteEntryName() *string {
	return s.TransitRouterRouteEntryName
}

func (s *UpdateTransitRouterRouteEntryRequest) SetClientToken(v string) *UpdateTransitRouterRouteEntryRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetDryRun(v bool) *UpdateTransitRouterRouteEntryRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetOwnerAccount(v string) *UpdateTransitRouterRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetOwnerId(v int64) *UpdateTransitRouterRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetResourceOwnerAccount(v string) *UpdateTransitRouterRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetResourceOwnerId(v int64) *UpdateTransitRouterRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryDescription(v string) *UpdateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryDescription = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryId(v string) *UpdateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryId = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) SetTransitRouterRouteEntryName(v string) *UpdateTransitRouterRouteEntryRequest {
	s.TransitRouterRouteEntryName = &v
	return s
}

func (s *UpdateTransitRouterRouteEntryRequest) Validate() error {
	return dara.Validate(s)
}
