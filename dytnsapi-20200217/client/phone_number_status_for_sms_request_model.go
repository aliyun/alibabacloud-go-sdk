// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForSmsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberStatusForSmsRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberStatusForSmsRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberStatusForSmsRequest
	GetMask() *string
	SetOwnerId(v int64) *PhoneNumberStatusForSmsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberStatusForSmsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberStatusForSmsRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberStatusForSmsRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// zf08***pi6
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
	// 181****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: plaintext
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

func (s PhoneNumberStatusForSmsRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForSmsRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForSmsRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberStatusForSmsRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberStatusForSmsRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberStatusForSmsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberStatusForSmsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberStatusForSmsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberStatusForSmsRequest) SetAuthCode(v string) *PhoneNumberStatusForSmsRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetInputNumber(v string) *PhoneNumberStatusForSmsRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetMask(v string) *PhoneNumberStatusForSmsRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetOwnerId(v int64) *PhoneNumberStatusForSmsRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForSmsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForSmsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberStatusForSmsRequest) Validate() error {
	return dara.Validate(s)
}
