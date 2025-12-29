// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberEncryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberEncryptRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberEncryptRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberEncryptRequest
	GetMask() *string
	SetOutId(v string) *PhoneNumberEncryptRequest
	GetOutId() *string
	SetOwnerId(v int64) *PhoneNumberEncryptRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberEncryptRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberEncryptRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberEncryptRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// >  You can query only one phone number at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Set the value to **NORMAL**.
	//
	// >  Only the NORMAL encryption method is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberEncryptRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberEncryptRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberEncryptRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberEncryptRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberEncryptRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberEncryptRequest) GetOutId() *string {
	return s.OutId
}

func (s *PhoneNumberEncryptRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberEncryptRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberEncryptRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberEncryptRequest) SetAuthCode(v string) *PhoneNumberEncryptRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetInputNumber(v string) *PhoneNumberEncryptRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetMask(v string) *PhoneNumberEncryptRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetOutId(v string) *PhoneNumberEncryptRequest {
	s.OutId = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetOwnerId(v int64) *PhoneNumberEncryptRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetResourceOwnerAccount(v string) *PhoneNumberEncryptRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberEncryptRequest) SetResourceOwnerId(v int64) *PhoneNumberEncryptRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberEncryptRequest) Validate() error {
	return dara.Validate(s)
}
