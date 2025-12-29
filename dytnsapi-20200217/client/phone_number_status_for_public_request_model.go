// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForPublicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberStatusForPublicRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberStatusForPublicRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberStatusForPublicRequest
	GetMask() *string
	SetOwnerId(v int64) *PhoneNumberStatusForPublicRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberStatusForPublicRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberStatusForPublicRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberStatusForPublicRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// 	- If the value of Mask is NORMAL, the value of this field is an 11-digit phone number.
	//
	// 	- If the value of Mask is MD5, the value of this field is a 32-bit encrypted string.
	//
	// 	- If the value of Mask is SHA256, the value of this field is a 64-bit encrypted string.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
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

func (s PhoneNumberStatusForPublicRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForPublicRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForPublicRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberStatusForPublicRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberStatusForPublicRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberStatusForPublicRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberStatusForPublicRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberStatusForPublicRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberStatusForPublicRequest) SetAuthCode(v string) *PhoneNumberStatusForPublicRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetInputNumber(v string) *PhoneNumberStatusForPublicRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetMask(v string) *PhoneNumberStatusForPublicRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetOwnerId(v int64) *PhoneNumberStatusForPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForPublicRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberStatusForPublicRequest) Validate() error {
	return dara.Validate(s)
}
