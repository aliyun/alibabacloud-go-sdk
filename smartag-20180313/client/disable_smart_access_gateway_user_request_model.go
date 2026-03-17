// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmartAccessGatewayUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DisableSmartAccessGatewayUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableSmartAccessGatewayUserRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableSmartAccessGatewayUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisableSmartAccessGatewayUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableSmartAccessGatewayUserRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DisableSmartAccessGatewayUserRequest
	GetSmartAGId() *string
	SetUserName(v string) *DisableSmartAccessGatewayUserRequest
	GetUserName() *string
}

type DisableSmartAccessGatewayUserRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG APP instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG APP instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-va03wf4l4idaj*****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The username of the client account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DisableSmartAccessGatewayUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSmartAccessGatewayUserRequest) GoString() string {
	return s.String()
}

func (s *DisableSmartAccessGatewayUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableSmartAccessGatewayUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableSmartAccessGatewayUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableSmartAccessGatewayUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableSmartAccessGatewayUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableSmartAccessGatewayUserRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DisableSmartAccessGatewayUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *DisableSmartAccessGatewayUserRequest) SetOwnerAccount(v string) *DisableSmartAccessGatewayUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) SetOwnerId(v int64) *DisableSmartAccessGatewayUserRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) SetRegionId(v string) *DisableSmartAccessGatewayUserRequest {
	s.RegionId = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) SetResourceOwnerAccount(v string) *DisableSmartAccessGatewayUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) SetResourceOwnerId(v int64) *DisableSmartAccessGatewayUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) SetSmartAGId(v string) *DisableSmartAccessGatewayUserRequest {
	s.SmartAGId = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) SetUserName(v string) *DisableSmartAccessGatewayUserRequest {
	s.UserName = &v
	return s
}

func (s *DisableSmartAccessGatewayUserRequest) Validate() error {
	return dara.Validate(s)
}
