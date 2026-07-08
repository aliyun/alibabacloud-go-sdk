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
	// > On the **My Applications*	- page in the [Phone Number Intelligence console](https://dytns.console.aliyun.com/analysis/apply), obtain the authorization ID. This ID is the authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to query.
	//
	// - If `Mask` is set to `NORMAL`, this parameter is an 11-digit mobile phone number.
	//
	// - If `Mask` is set to `MD5`, this parameter is a 32-character hashed string.
	//
	// - If `Mask` is set to `SHA256`, this parameter is a 64-character hashed string.
	//
	// - If `Mask` is set to `SM3`, this parameter is a 64-character hashed string.
	//
	// 	Notice:
	//
	// The letters in the hashed string are case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// - **NORMAL**: The number is not encrypted.
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
