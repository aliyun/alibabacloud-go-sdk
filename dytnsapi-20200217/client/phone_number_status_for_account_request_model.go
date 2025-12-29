// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberStatusForAccountRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberStatusForAccountRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberStatusForAccountRequest
	GetMask() *string
	SetOwnerId(v int64) *PhoneNumberStatusForAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberStatusForAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberStatusForAccountRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberStatusForAccountRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, specify an 11-digit phone number in plaintext.
	//
	// 	- If the value of Mask is MD5, specify a 32-bit string that is encrypted by using MD5.
	//
	// 	- If the value of Mask is SHA256, specify a 64-bit string that is encrypted by using SHA256.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**
	//
	// 	- **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberStatusForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForAccountRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForAccountRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberStatusForAccountRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberStatusForAccountRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberStatusForAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberStatusForAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberStatusForAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberStatusForAccountRequest) SetAuthCode(v string) *PhoneNumberStatusForAccountRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetInputNumber(v string) *PhoneNumberStatusForAccountRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetMask(v string) *PhoneNumberStatusForAccountRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetOwnerId(v int64) *PhoneNumberStatusForAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberStatusForAccountRequest) Validate() error {
	return dara.Validate(s)
}
