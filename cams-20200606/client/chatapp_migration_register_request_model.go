// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappMigrationRegisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ChatappMigrationRegisterRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ChatappMigrationRegisterRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ChatappMigrationRegisterRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ChatappMigrationRegisterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappMigrationRegisterRequest
	GetResourceOwnerId() *int64
}

type ChatappMigrationRegisterRequest struct {
	// None
	//
	// This parameter is required.
	//
	// example:
	//
	// 29348393884****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChatappMigrationRegisterRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationRegisterRequest) GoString() string {
	return s.String()
}

func (s *ChatappMigrationRegisterRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappMigrationRegisterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappMigrationRegisterRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappMigrationRegisterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappMigrationRegisterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappMigrationRegisterRequest) SetCustSpaceId(v string) *ChatappMigrationRegisterRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappMigrationRegisterRequest) SetOwnerId(v int64) *ChatappMigrationRegisterRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappMigrationRegisterRequest) SetPhoneNumber(v string) *ChatappMigrationRegisterRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappMigrationRegisterRequest) SetResourceOwnerAccount(v string) *ChatappMigrationRegisterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappMigrationRegisterRequest) SetResourceOwnerId(v int64) *ChatappMigrationRegisterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappMigrationRegisterRequest) Validate() error {
	return dara.Validate(s)
}
