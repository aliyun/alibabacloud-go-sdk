// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneEncryptionPublicKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetPhoneEncryptionPublicKeyRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetPhoneEncryptionPublicKeyRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetPhoneEncryptionPublicKeyRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetPhoneEncryptionPublicKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPhoneEncryptionPublicKeyRequest
	GetResourceOwnerId() *int64
}

type GetPhoneEncryptionPublicKeyRequest struct {
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
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number. You can view it on the
	//
	// <props="china">[**Channel Management**](https://chatapp.console.aliyun.com/ChannelsManagement)
	//
	// <props="intl">[Channel Management](https://chatapp.console.alibabacloud.com/CustomerList)> **Manage*	- > **WABA Management*	- > **Number Management*	- page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86189********
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetPhoneEncryptionPublicKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPhoneEncryptionPublicKeyRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneEncryptionPublicKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPhoneEncryptionPublicKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetCustSpaceId(v string) *GetPhoneEncryptionPublicKeyRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetOwnerId(v int64) *GetPhoneEncryptionPublicKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetPhoneNumber(v string) *GetPhoneEncryptionPublicKeyRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetResourceOwnerAccount(v string) *GetPhoneEncryptionPublicKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyRequest) SetResourceOwnerId(v int64) *GetPhoneEncryptionPublicKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyRequest) Validate() error {
	return dara.Validate(s)
}
