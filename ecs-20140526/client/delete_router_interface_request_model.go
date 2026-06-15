// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouterInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRouterInterfaceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteRouterInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouterInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRouterInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRouterInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouterInterfaceRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *DeleteRouterInterfaceRequest
	GetRouterInterfaceId() *string
	SetUserCidr(v string) *DeleteRouterInterfaceRequest
	GetUserCidr() *string
}

type DeleteRouterInterfaceRequest struct {
	// A client-generated, case-sensitive token used to ensure request idempotency. You must ensure that the token is unique for each request. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the router interface is located.
	//
	// You can call the [DescribeRegions](~~DescribeRegions~~) operation to get the latest list of regions.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the router interface that you want to delete.
	//
	// This parameter is required.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// This parameter is used for auditing and is required in specific scenarios, such as when deleting a router interface for a peer-to-peer connection.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s DeleteRouterInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouterInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouterInterfaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRouterInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouterInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouterInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouterInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouterInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouterInterfaceRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DeleteRouterInterfaceRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *DeleteRouterInterfaceRequest) SetClientToken(v string) *DeleteRouterInterfaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetOwnerAccount(v string) *DeleteRouterInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetOwnerId(v int64) *DeleteRouterInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetRegionId(v string) *DeleteRouterInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetResourceOwnerAccount(v string) *DeleteRouterInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetResourceOwnerId(v int64) *DeleteRouterInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetRouterInterfaceId(v string) *DeleteRouterInterfaceRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) SetUserCidr(v string) *DeleteRouterInterfaceRequest {
	s.UserCidr = &v
	return s
}

func (s *DeleteRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
