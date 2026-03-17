// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteSmartAccessGatewayClientUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSmartAccessGatewayClientUserRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSmartAccessGatewayClientUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSmartAccessGatewayClientUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSmartAccessGatewayClientUserRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DeleteSmartAccessGatewayClientUserRequest
	GetSmartAGId() *string
	SetUserName(v string) *DeleteSmartAccessGatewayClientUserRequest
	GetUserName() *string
}

type DeleteSmartAccessGatewayClientUserRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-kzo5dvms3dqii3****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The username of the client account.
	//
	// This parameter is required.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteSmartAccessGatewayClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayClientUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DeleteSmartAccessGatewayClientUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetOwnerAccount(v string) *DeleteSmartAccessGatewayClientUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetOwnerId(v int64) *DeleteSmartAccessGatewayClientUserRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetRegionId(v string) *DeleteSmartAccessGatewayClientUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetResourceOwnerAccount(v string) *DeleteSmartAccessGatewayClientUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetResourceOwnerId(v int64) *DeleteSmartAccessGatewayClientUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetSmartAGId(v string) *DeleteSmartAccessGatewayClientUserRequest {
	s.SmartAGId = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) SetUserName(v string) *DeleteSmartAccessGatewayClientUserRequest {
	s.UserName = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserRequest) Validate() error {
	return dara.Validate(s)
}
