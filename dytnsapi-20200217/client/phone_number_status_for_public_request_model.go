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
	// > On the **My Applications*	- page of the [Phone Number Intelligence console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization ID and use it as the authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried.
	//
	// - If `Mask` is set to `NORMAL`, this parameter must be an 11-digit mobile phone number.
	//
	// - If `Mask` is set to `MD5`, this parameter must be a 32-character encrypted string.
	//
	// - If `Mask` is set to `SHA256`, this parameter must be a 64-character encrypted string.
	//
	// - If `Mask` is set to `SM3`, this parameter must be a 64-character encrypted string.
	//
	// > The encrypted strings are case-insensitive.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// - **NORMAL**: The phone number is not encrypted.
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
