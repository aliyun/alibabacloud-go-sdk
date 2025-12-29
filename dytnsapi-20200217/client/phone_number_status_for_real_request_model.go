// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberStatusForRealRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberStatusForRealRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberStatusForRealRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberStatusForRealRequest
	GetMask() *string
	SetOwnerId(v int64) *PhoneNumberStatusForRealRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberStatusForRealRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberStatusForRealRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberStatusForRealRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 189****8999
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

func (s PhoneNumberStatusForRealRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberStatusForRealRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberStatusForRealRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberStatusForRealRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberStatusForRealRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberStatusForRealRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberStatusForRealRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberStatusForRealRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberStatusForRealRequest) SetAuthCode(v string) *PhoneNumberStatusForRealRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetInputNumber(v string) *PhoneNumberStatusForRealRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetMask(v string) *PhoneNumberStatusForRealRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetOwnerId(v int64) *PhoneNumberStatusForRealRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetResourceOwnerAccount(v string) *PhoneNumberStatusForRealRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) SetResourceOwnerId(v int64) *PhoneNumberStatusForRealRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberStatusForRealRequest) Validate() error {
	return dara.Validate(s)
}
