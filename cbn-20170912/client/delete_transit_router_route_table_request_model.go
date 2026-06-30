// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterRouteTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterRouteTableRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *DeleteTransitRouterRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteTransitRouterRouteTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteTransitRouterRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteTransitRouterRouteTableRequest
	GetResourceOwnerId() *int64
	SetTransitRouterRouteTableId(v string) *DeleteTransitRouterRouteTableRequest
	GetTransitRouterRouteTableId() *string
}

type DeleteTransitRouterRouteTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a token from your client to ensure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. A dry run checks for issues such as permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. The custom route table is deleted after the request passes the check.
	//
	// - **true**: Sends a check request. The system checks the required parameters, request format, and permissions. If the check fails, an error message is returned. If the check succeeds, the `DryRunOperation` error code is returned. The custom route table is not deleted.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the custom route table.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1xbcgpgcz9axl9m****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s DeleteTransitRouterRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRouteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterRouteTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteTransitRouterRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteTransitRouterRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteTransitRouterRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteTransitRouterRouteTableRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *DeleteTransitRouterRouteTableRequest) SetClientToken(v string) *DeleteTransitRouterRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetDryRun(v bool) *DeleteTransitRouterRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetOwnerAccount(v string) *DeleteTransitRouterRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetOwnerId(v int64) *DeleteTransitRouterRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetResourceOwnerAccount(v string) *DeleteTransitRouterRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetResourceOwnerId(v int64) *DeleteTransitRouterRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) SetTransitRouterRouteTableId(v string) *DeleteTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *DeleteTransitRouterRouteTableRequest) Validate() error {
	return dara.Validate(s)
}
