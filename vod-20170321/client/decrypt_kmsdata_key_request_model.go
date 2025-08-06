// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptKMSDataKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCipherText(v string) *DecryptKMSDataKeyRequest
	GetCipherText() *string
	SetOwnerAccount(v string) *DecryptKMSDataKeyRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DecryptKMSDataKeyRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *DecryptKMSDataKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DecryptKMSDataKeyRequest
	GetResourceOwnerId() *string
}

type DecryptKMSDataKeyRequest struct {
	// The ciphertext to be decrypted.
	//
	// This parameter is required.
	//
	// example:
	//
	// DZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmaaSl+TztSIMe43nbTH/Z1Wr4XfLftKhAciUmDQXuMRl4WTvKhxjMThjK****
	CipherText           *string `json:"CipherText,omitempty" xml:"CipherText,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DecryptKMSDataKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DecryptKMSDataKeyRequest) GoString() string {
	return s.String()
}

func (s *DecryptKMSDataKeyRequest) GetCipherText() *string {
	return s.CipherText
}

func (s *DecryptKMSDataKeyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DecryptKMSDataKeyRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DecryptKMSDataKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DecryptKMSDataKeyRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DecryptKMSDataKeyRequest) SetCipherText(v string) *DecryptKMSDataKeyRequest {
	s.CipherText = &v
	return s
}

func (s *DecryptKMSDataKeyRequest) SetOwnerAccount(v string) *DecryptKMSDataKeyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DecryptKMSDataKeyRequest) SetOwnerId(v string) *DecryptKMSDataKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *DecryptKMSDataKeyRequest) SetResourceOwnerAccount(v string) *DecryptKMSDataKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DecryptKMSDataKeyRequest) SetResourceOwnerId(v string) *DecryptKMSDataKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DecryptKMSDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
