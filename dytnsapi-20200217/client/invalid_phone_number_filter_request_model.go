// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidPhoneNumberFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *InvalidPhoneNumberFilterRequest
	GetAuthCode() *string
	SetInputNumber(v string) *InvalidPhoneNumberFilterRequest
	GetInputNumber() *string
	SetMask(v string) *InvalidPhoneNumberFilterRequest
	GetMask() *string
	SetOwnerId(v int64) *InvalidPhoneNumberFilterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *InvalidPhoneNumberFilterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *InvalidPhoneNumberFilterRequest
	GetResourceOwnerId() *int64
}

type InvalidPhoneNumberFilterRequest struct {
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
	// The phone number to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number.
	//
	// >  Only the NORMAL encryption method is supported.
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

func (s InvalidPhoneNumberFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s InvalidPhoneNumberFilterRequest) GoString() string {
	return s.String()
}

func (s *InvalidPhoneNumberFilterRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *InvalidPhoneNumberFilterRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *InvalidPhoneNumberFilterRequest) GetMask() *string {
	return s.Mask
}

func (s *InvalidPhoneNumberFilterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InvalidPhoneNumberFilterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *InvalidPhoneNumberFilterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *InvalidPhoneNumberFilterRequest) SetAuthCode(v string) *InvalidPhoneNumberFilterRequest {
	s.AuthCode = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetInputNumber(v string) *InvalidPhoneNumberFilterRequest {
	s.InputNumber = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetMask(v string) *InvalidPhoneNumberFilterRequest {
	s.Mask = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetOwnerId(v int64) *InvalidPhoneNumberFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetResourceOwnerAccount(v string) *InvalidPhoneNumberFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) SetResourceOwnerId(v int64) *InvalidPhoneNumberFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InvalidPhoneNumberFilterRequest) Validate() error {
	return dara.Validate(s)
}
