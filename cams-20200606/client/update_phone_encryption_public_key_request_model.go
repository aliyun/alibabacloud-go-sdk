// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneEncryptionPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *UpdatePhoneEncryptionPublicKeyRequest
	GetCustSpaceId() *string
	SetEncryptionPublicKey(v string) *UpdatePhoneEncryptionPublicKeyRequest
	GetEncryptionPublicKey() *string
	SetOwnerId(v int64) *UpdatePhoneEncryptionPublicKeyRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *UpdatePhoneEncryptionPublicKeyRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *UpdatePhoneEncryptionPublicKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdatePhoneEncryptionPublicKeyRequest
	GetResourceOwnerId() *int64
}

type UpdatePhoneEncryptionPublicKeyRequest struct {
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" xml:"EncryptionPublicKey,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdatePhoneEncryptionPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneEncryptionPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) GetEncryptionPublicKey() *string {
	return s.EncryptionPublicKey
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetCustSpaceId(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.CustSpaceId = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetEncryptionPublicKey(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.EncryptionPublicKey = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetOwnerId(v int64) *UpdatePhoneEncryptionPublicKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetPhoneNumber(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetResourceOwnerAccount(v string) *UpdatePhoneEncryptionPublicKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) SetResourceOwnerId(v int64) *UpdatePhoneEncryptionPublicKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
