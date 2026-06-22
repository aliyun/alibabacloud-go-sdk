// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappMigrationVerifiedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ChatappMigrationVerifiedRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ChatappMigrationVerifiedRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *ChatappMigrationVerifiedRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *ChatappMigrationVerifiedRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ChatappMigrationVerifiedRequest
	GetResourceOwnerId() *int64
	SetVerifyCode(v string) *ChatappMigrationVerifiedRequest
	GetVerifyCode() *string
}

type ChatappMigrationVerifiedRequest struct {
	// The Space ID of the Independent Software Vendor (ISV) sub-customer.
	//
	// This parameter is required.
	//
	// example:
	//
	// 29348393884****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86138000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8287**
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s ChatappMigrationVerifiedRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatappMigrationVerifiedRequest) GoString() string {
	return s.String()
}

func (s *ChatappMigrationVerifiedRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ChatappMigrationVerifiedRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatappMigrationVerifiedRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappMigrationVerifiedRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ChatappMigrationVerifiedRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ChatappMigrationVerifiedRequest) GetVerifyCode() *string {
	return s.VerifyCode
}

func (s *ChatappMigrationVerifiedRequest) SetCustSpaceId(v string) *ChatappMigrationVerifiedRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetOwnerId(v int64) *ChatappMigrationVerifiedRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetPhoneNumber(v string) *ChatappMigrationVerifiedRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetResourceOwnerAccount(v string) *ChatappMigrationVerifiedRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetResourceOwnerId(v int64) *ChatappMigrationVerifiedRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) SetVerifyCode(v string) *ChatappMigrationVerifiedRequest {
	s.VerifyCode = &v
	return s
}

func (s *ChatappMigrationVerifiedRequest) Validate() error {
	return dara.Validate(s)
}
