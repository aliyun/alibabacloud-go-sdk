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
	// > The authorization code is the authorization ID that you can find on the **My Applications*	- page of the [Phone Number Encyclopedia console](https://dytns.console.aliyun.com/analysis/apply).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to query.
	//
	// - If you set `Mask` to `NORMAL`, specify an 11-digit mobile number.
	//
	// - If you set `Mask` to `MD5`, specify a 32-bit encrypted string.
	//
	// - If you set `Mask` to `SHA256`, specify a 64-bit encrypted string.
	//
	// - If you set `Mask` to `SM3`, specify a 64-bit encrypted string.
	//
	// 	Notice:
	//
	// The letters in the encrypted string are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// - **NORMAL**: The number is in plaintext.
	//
	// - **MD5**
	//
	// - **SHA256**
	//
	// - **SM3**
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
