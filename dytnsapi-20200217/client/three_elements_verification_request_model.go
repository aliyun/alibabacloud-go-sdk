// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iThreeElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *ThreeElementsVerificationRequest
	GetAuthCode() *string
	SetCertCode(v string) *ThreeElementsVerificationRequest
	GetCertCode() *string
	SetInputNumber(v string) *ThreeElementsVerificationRequest
	GetInputNumber() *string
	SetMask(v string) *ThreeElementsVerificationRequest
	GetMask() *string
	SetName(v string) *ThreeElementsVerificationRequest
	GetName() *string
	SetOwnerId(v int64) *ThreeElementsVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ThreeElementsVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ThreeElementsVerificationRequest
	GetResourceOwnerId() *int64
}

type ThreeElementsVerificationRequest struct {
	// The authorization code.
	//
	// > On the **My Applications*	- page of the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), the authorization ID is the authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The ID card number to be verified.
	//
	// - If Mask is set to NORMAL, this field is in plaintext.
	//
	// - If Mask is set to MD5, encrypt this field by using MD5.
	//
	// - If Mask is set to SHA256, encrypt this field by using SHA256.
	//
	// 	Notice: Letters in the encrypted string are case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 83d8040d3cb2181e04****dc6ff5566d4493876a4a5da782887446356b0a787e
	CertCode *string `json:"CertCode,omitempty" xml:"CertCode,omitempty"`
	// The phone number to be verified.
	//
	// - If Mask is set to NORMAL, this field is in plaintext.
	//
	// - If Mask is set to MD5, encrypt this field by using MD5.
	//
	// - If Mask is set to SHA256, encrypt this field by using SHA256.
	//
	// 	Notice: Letters in the encrypted string are case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method. Valid values:
	//
	// - **NORMAL**: No encryption.
	//
	// - **MD5**
	//
	// - **SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The name to be verified.
	//
	// - If Mask is set to NORMAL, this field is in plaintext.
	//
	// - If Mask is set to MD5, encrypt this field by using MD5.
	//
	// - If Mask is set to SHA256, encrypt this field by using SHA256.
	//
	// 	Notice: Letters in the encrypted string are case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ThreeElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *ThreeElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *ThreeElementsVerificationRequest) GetCertCode() *string {
	return s.CertCode
}

func (s *ThreeElementsVerificationRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *ThreeElementsVerificationRequest) GetMask() *string {
	return s.Mask
}

func (s *ThreeElementsVerificationRequest) GetName() *string {
	return s.Name
}

func (s *ThreeElementsVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ThreeElementsVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ThreeElementsVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ThreeElementsVerificationRequest) SetAuthCode(v string) *ThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetCertCode(v string) *ThreeElementsVerificationRequest {
	s.CertCode = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetInputNumber(v string) *ThreeElementsVerificationRequest {
	s.InputNumber = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetMask(v string) *ThreeElementsVerificationRequest {
	s.Mask = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetName(v string) *ThreeElementsVerificationRequest {
	s.Name = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetOwnerId(v int64) *ThreeElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetResourceOwnerAccount(v string) *ThreeElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ThreeElementsVerificationRequest) SetResourceOwnerId(v int64) *ThreeElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ThreeElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}
