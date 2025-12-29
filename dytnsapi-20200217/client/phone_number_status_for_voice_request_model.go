// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForVoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberStatusForVoiceRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberStatusForVoiceRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberStatusForVoiceRequest
	GetMask() *string
	SetOwnerId(v int64) *PhoneNumberStatusForVoiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberStatusForVoiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberStatusForVoiceRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberStatusForVoiceRequest struct {
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
	// The encryption method. Valid values:
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

func (s PhoneNumberStatusForVoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForVoiceRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForVoiceRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberStatusForVoiceRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberStatusForVoiceRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberStatusForVoiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberStatusForVoiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberStatusForVoiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberStatusForVoiceRequest) SetAuthCode(v string) *PhoneNumberStatusForVoiceRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetInputNumber(v string) *PhoneNumberStatusForVoiceRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetMask(v string) *PhoneNumberStatusForVoiceRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetOwnerId(v int64) *PhoneNumberStatusForVoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForVoiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForVoiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberStatusForVoiceRequest) Validate() error {
	return dara.Validate(s)
}
