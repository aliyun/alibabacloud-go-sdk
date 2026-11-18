// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDownloadSecretKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppDecryptKey(v string) *GenerateDownloadSecretKeyRequest
	GetAppDecryptKey() *string
	SetAppIdentification(v string) *GenerateDownloadSecretKeyRequest
	GetAppIdentification() *string
	SetOwnerId(v int64) *GenerateDownloadSecretKeyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GenerateDownloadSecretKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GenerateDownloadSecretKeyRequest
	GetResourceOwnerId() *int64
}

type GenerateDownloadSecretKeyRequest struct {
	// The custom encryption private key string. The string must be 16 to 32 characters in length, consist of letters and digits, and contain uppercase letters, lowercase letters, and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// AppDecryptKeyAndroid20230101
	AppDecryptKey *string `json:"AppDecryptKey,omitempty" xml:"AppDecryptKey,omitempty"`
	// The unique app identifier.
	//
	// - Android: the SHA-1 value of the keystore file. The obtained SHA-1 value must be a string that contains colons.
	//
	// - iOS: the app bundle ID.
	//
	// - Windows: the serial number of the digital signature certificate.
	//
	// For more information about how to obtain the unique app identifier, see [Obtain the unique app identifier](~~86107#section-wtj-9d7-lg2~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// BB:0D:AC:74:D3:21:09:EF:9C:71:1A:5E:77:2C:8E:BF:03:FD:FA:5A
	AppIdentification    *string `json:"AppIdentification,omitempty" xml:"AppIdentification,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GenerateDownloadSecretKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDownloadSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateDownloadSecretKeyRequest) GetAppDecryptKey() *string {
	return s.AppDecryptKey
}

func (s *GenerateDownloadSecretKeyRequest) GetAppIdentification() *string {
	return s.AppIdentification
}

func (s *GenerateDownloadSecretKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GenerateDownloadSecretKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GenerateDownloadSecretKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GenerateDownloadSecretKeyRequest) SetAppDecryptKey(v string) *GenerateDownloadSecretKeyRequest {
	s.AppDecryptKey = &v
	return s
}

func (s *GenerateDownloadSecretKeyRequest) SetAppIdentification(v string) *GenerateDownloadSecretKeyRequest {
	s.AppIdentification = &v
	return s
}

func (s *GenerateDownloadSecretKeyRequest) SetOwnerId(v int64) *GenerateDownloadSecretKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *GenerateDownloadSecretKeyRequest) SetResourceOwnerAccount(v string) *GenerateDownloadSecretKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GenerateDownloadSecretKeyRequest) SetResourceOwnerId(v int64) *GenerateDownloadSecretKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GenerateDownloadSecretKeyRequest) Validate() error {
	return dara.Validate(s)
}
