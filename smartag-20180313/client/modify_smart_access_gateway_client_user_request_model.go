// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ModifySmartAccessGatewayClientUserRequest
	GetBandwidth() *int32
	SetEmail(v string) *ModifySmartAccessGatewayClientUserRequest
	GetEmail() *string
	SetOwnerAccount(v string) *ModifySmartAccessGatewayClientUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySmartAccessGatewayClientUserRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySmartAccessGatewayClientUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySmartAccessGatewayClientUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySmartAccessGatewayClientUserRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySmartAccessGatewayClientUserRequest
	GetSmartAGId() *string
	SetUserName(v string) *ModifySmartAccessGatewayClientUserRequest
	GetUserName() *string
}

type ModifySmartAccessGatewayClientUserRequest struct {
	// The maximum bandwidth allocated to the client account. Unit: Kbit/s.
	//
	// Valid values: **1*	- to **20000**.
	//
	// example:
	//
	// 1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The email address of the client account.
	//
	// example:
	//
	// username@example.com
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
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
	// sag-rz2e23c0e78ema****
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

func (s ModifySmartAccessGatewayClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayClientUserRequest) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetEmail() *string {
	return s.Email
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySmartAccessGatewayClientUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetBandwidth(v int32) *ModifySmartAccessGatewayClientUserRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetEmail(v string) *ModifySmartAccessGatewayClientUserRequest {
	s.Email = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetOwnerAccount(v string) *ModifySmartAccessGatewayClientUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetOwnerId(v int64) *ModifySmartAccessGatewayClientUserRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetRegionId(v string) *ModifySmartAccessGatewayClientUserRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetResourceOwnerAccount(v string) *ModifySmartAccessGatewayClientUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetResourceOwnerId(v int64) *ModifySmartAccessGatewayClientUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetSmartAGId(v string) *ModifySmartAccessGatewayClientUserRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) SetUserName(v string) *ModifySmartAccessGatewayClientUserRequest {
	s.UserName = &v
	return s
}

func (s *ModifySmartAccessGatewayClientUserRequest) Validate() error {
	return dara.Validate(s)
}
