// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOnlineTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberOnlineTimeRequest
	GetAuthCode() *string
	SetCarrier(v string) *DescribePhoneNumberOnlineTimeRequest
	GetCarrier() *string
	SetInputNumber(v string) *DescribePhoneNumberOnlineTimeRequest
	GetInputNumber() *string
	SetMask(v string) *DescribePhoneNumberOnlineTimeRequest
	GetMask() *string
	SetOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberOnlineTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberOnlineTimeRequest struct {
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
	// The carrier. Valid values:
	//
	// 	- **MOBILE**: China Mobile
	//
	// 	- **UNICOM**: China Unicom
	//
	// 	- **TELECOM**: China Telecom
	//
	// >  Alibaba Cloud automatically determines the carrier based on the carrier who assigns the phone number. Therefore, the value of this field does not affect the query result.
	//
	// example:
	//
	// UNICOM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
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
	// 1390000****
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

func (s DescribePhoneNumberOnlineTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOnlineTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberOnlineTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetAuthCode(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetCarrier(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.Carrier = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetInputNumber(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetMask(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberOnlineTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberOnlineTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberOnlineTimeRequest) Validate() error {
	return dara.Validate(s)
}
