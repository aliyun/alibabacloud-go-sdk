// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteExpressConnectRequest
	GetClientToken() *string
	SetForce(v bool) *DeleteExpressConnectRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteExpressConnectRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteExpressConnectRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteExpressConnectRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteExpressConnectRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteExpressConnectRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *DeleteExpressConnectRequest
	GetRouterInterfaceId() *string
}

type DeleteExpressConnectRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// Use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the route entries associated with the Express Connect instance.
	//
	// 	- **true**: forcefully deletes the snapshot
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Force        *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Express Connect instance is deployed. Call the [DescribeRegion](https://www.alibabacloud.com/help/vpc/developer-reference/api-vpc-2016-04-28-describeregions?spm=a2c63.p38356.0.i2) operation to query the region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the Express Connect instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ri-119mfjz****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
}

func (s DeleteExpressConnectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExpressConnectRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteExpressConnectRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteExpressConnectRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteExpressConnectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExpressConnectRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteExpressConnectRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteExpressConnectRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DeleteExpressConnectRequest) SetClientToken(v string) *DeleteExpressConnectRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetForce(v bool) *DeleteExpressConnectRequest {
	s.Force = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetOwnerAccount(v string) *DeleteExpressConnectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetOwnerId(v int64) *DeleteExpressConnectRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetRegionId(v string) *DeleteExpressConnectRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetResourceOwnerAccount(v string) *DeleteExpressConnectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetResourceOwnerId(v int64) *DeleteExpressConnectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteExpressConnectRequest) SetRouterInterfaceId(v string) *DeleteExpressConnectRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *DeleteExpressConnectRequest) Validate() error {
	return dara.Validate(s)
}
