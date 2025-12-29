// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPhoneNumberConvertServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *PhoneNumberConvertServiceRequest
	GetAuthCode() *string
	SetInputNumber(v string) *PhoneNumberConvertServiceRequest
	GetInputNumber() *string
	SetMask(v string) *PhoneNumberConvertServiceRequest
	GetMask() *string
	SetOwnerId(v int64) *PhoneNumberConvertServiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *PhoneNumberConvertServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PhoneNumberConvertServiceRequest
	GetResourceOwnerId() *int64
}

type PhoneNumberConvertServiceRequest struct {
	// This parameter is required.
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PhoneNumberConvertServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s PhoneNumberConvertServiceRequest) GoString() string {
	return s.String()
}

func (s *PhoneNumberConvertServiceRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *PhoneNumberConvertServiceRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *PhoneNumberConvertServiceRequest) GetMask() *string {
	return s.Mask
}

func (s *PhoneNumberConvertServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PhoneNumberConvertServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PhoneNumberConvertServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PhoneNumberConvertServiceRequest) SetAuthCode(v string) *PhoneNumberConvertServiceRequest {
	s.AuthCode = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetInputNumber(v string) *PhoneNumberConvertServiceRequest {
	s.InputNumber = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetMask(v string) *PhoneNumberConvertServiceRequest {
	s.Mask = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetOwnerId(v int64) *PhoneNumberConvertServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetResourceOwnerAccount(v string) *PhoneNumberConvertServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) SetResourceOwnerId(v int64) *PhoneNumberConvertServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PhoneNumberConvertServiceRequest) Validate() error {
	return dara.Validate(s)
}
