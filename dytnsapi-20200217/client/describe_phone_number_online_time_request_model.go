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
	// > Log on to the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), go to the **My Applications*	- page, and obtain the authorization ID, which is the authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The external carrier. Valid values:
	//
	// - **MOBILE**: China Mobile.
	//
	// - **UNICOM**: China Unicom.
	//
	// - **TELECOM**: China Telecom.
	//
	// 	Notice: This parameter is optional. Alibaba Cloud automatically determines the carrier type based on the phone number. The value of this field has no impact on the query result.
	//
	// example:
	//
	// UNICOM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The phone number to be queried.
	//
	// - If Mask is set to NORMAL, this field is an 11-digit phone number.
	//
	// - If Mask is set to MD5, this field is a 32-character encrypted string.
	//
	// - If Mask is set to SHA256, this field is a 64-character encrypted string.
	//
	// 	Notice: Letters in the encrypted string are case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// - **NORMAL**: no encryption
	//
	// - **MD5**
	//
	// - **SHA256**
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
