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
}

type DeleteRouterInterfaceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the router interface is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urz****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
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

func (s *DeleteRouterInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
