// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickOutClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *KickOutClientsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *KickOutClientsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *KickOutClientsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *KickOutClientsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *KickOutClientsRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *KickOutClientsRequest
	GetSmartAGId() *string
	SetUsername(v string) *KickOutClientsRequest
	GetUsername() *string
}

type KickOutClientsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the SAG APP instance is deployed.
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
	// sag-ehjfb*******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The username of the client account. Usernames of client accounts added to the same SAG APP instance are unique.
	//
	// For a client account, if you specify the username, you must also specify the password.
	//
	// This parameter is required.
	//
	// example:
	//
	// doctest
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s KickOutClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s KickOutClientsRequest) GoString() string {
	return s.String()
}

func (s *KickOutClientsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *KickOutClientsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *KickOutClientsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KickOutClientsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *KickOutClientsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *KickOutClientsRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *KickOutClientsRequest) GetUsername() *string {
	return s.Username
}

func (s *KickOutClientsRequest) SetOwnerAccount(v string) *KickOutClientsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *KickOutClientsRequest) SetOwnerId(v int64) *KickOutClientsRequest {
	s.OwnerId = &v
	return s
}

func (s *KickOutClientsRequest) SetRegionId(v string) *KickOutClientsRequest {
	s.RegionId = &v
	return s
}

func (s *KickOutClientsRequest) SetResourceOwnerAccount(v string) *KickOutClientsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *KickOutClientsRequest) SetResourceOwnerId(v int64) *KickOutClientsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *KickOutClientsRequest) SetSmartAGId(v string) *KickOutClientsRequest {
	s.SmartAGId = &v
	return s
}

func (s *KickOutClientsRequest) SetUsername(v string) *KickOutClientsRequest {
	s.Username = &v
	return s
}

func (s *KickOutClientsRequest) Validate() error {
	return dara.Validate(s)
}
