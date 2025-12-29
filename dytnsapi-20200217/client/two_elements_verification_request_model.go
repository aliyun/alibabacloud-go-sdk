// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTwoElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *TwoElementsVerificationRequest
	GetAuthCode() *string
	SetInputNumber(v string) *TwoElementsVerificationRequest
	GetInputNumber() *string
	SetMask(v string) *TwoElementsVerificationRequest
	GetMask() *string
	SetName(v string) *TwoElementsVerificationRequest
	GetName() *string
	SetOwnerId(v int64) *TwoElementsVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TwoElementsVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TwoElementsVerificationRequest
	GetResourceOwnerId() *int64
}

type TwoElementsVerificationRequest struct {
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
	// The phone number to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
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
	// MD5
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The name to be verified.
	//
	// 	- If the value of Mask is NORMAL, specify a value in plaintext for this field.
	//
	// 	- If the value of Mask is MD5, specify a MD5-encrypted value for this field.
	//
	// 	- If the value of Mask is SHA256, specify a SHA256-encrypted value for this field.
	//
	// >  Letters in the encrypted strings are not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TwoElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s TwoElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *TwoElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *TwoElementsVerificationRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *TwoElementsVerificationRequest) GetMask() *string {
	return s.Mask
}

func (s *TwoElementsVerificationRequest) GetName() *string {
	return s.Name
}

func (s *TwoElementsVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TwoElementsVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TwoElementsVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TwoElementsVerificationRequest) SetAuthCode(v string) *TwoElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetInputNumber(v string) *TwoElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetMask(v string) *TwoElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetName(v string) *TwoElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetOwnerId(v int64) *TwoElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetResourceOwnerAccount(v string) *TwoElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TwoElementsVerificationRequest) SetResourceOwnerId(v int64) *TwoElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TwoElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}
