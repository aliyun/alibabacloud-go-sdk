// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequiredPhoneCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *RequiredPhoneCodeRequest
	GetOwnerId() *int64
	SetPhoneNo(v string) *RequiredPhoneCodeRequest
	GetPhoneNo() *string
	SetResourceOwnerAccount(v string) *RequiredPhoneCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RequiredPhoneCodeRequest
	GetResourceOwnerId() *int64
}

type RequiredPhoneCodeRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 137****1234
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RequiredPhoneCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RequiredPhoneCodeRequest) GoString() string {
	return s.String()
}

func (s *RequiredPhoneCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RequiredPhoneCodeRequest) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *RequiredPhoneCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RequiredPhoneCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RequiredPhoneCodeRequest) SetOwnerId(v int64) *RequiredPhoneCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *RequiredPhoneCodeRequest) SetPhoneNo(v string) *RequiredPhoneCodeRequest {
	s.PhoneNo = &v
	return s
}

func (s *RequiredPhoneCodeRequest) SetResourceOwnerAccount(v string) *RequiredPhoneCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RequiredPhoneCodeRequest) SetResourceOwnerId(v int64) *RequiredPhoneCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RequiredPhoneCodeRequest) Validate() error {
	return dara.Validate(s)
}
