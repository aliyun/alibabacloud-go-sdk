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
	// The space ID of the ISV sub-customer or the instance ID of the direct customer. You can view the Space ID on the
	//
	// <props="china">[Channel Management](https://chatapp.console.aliyun.com/ChannelsManagement)
	//
	// <props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList)
	//
	// page.
	//
	// example:
	//
	// cams-kei****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The public key.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt+PMrYw4zUDEk+VeTrp0\\n8LZaoVpiVFErX7iuoDjUs4F9vkxMQuIABjcXw\\/swzTMEopLORQV28uqN\\/2\\/x9hjU\\****\\/Zwa\\/Vk5Svp4\\niVY4e22FsJCCWUEMvayO8Q+3UGq0eXXQ+8SlUWEMq1VaJ4pwCLsMnmgybA+VmJxi\\nkwIDAQAB\\n-----END PUBLIC KEY-----"
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" xml:"EncryptionPublicKey,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number. You can view the phone number on the
	//
	// <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)
	//
	// <props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList)> **Manage*	- > **WABA Management*	- > **Number Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86158********
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
