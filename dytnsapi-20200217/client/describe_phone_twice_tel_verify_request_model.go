// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneTwiceTelVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneTwiceTelVerifyRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneTwiceTelVerifyRequest
	GetInputNumber() *string
	SetMask(v string) *DescribePhoneTwiceTelVerifyRequest
	GetMask() *string
	SetOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneTwiceTelVerifyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribePhoneTwiceTelVerifyRequest
	GetStartTime() *string
}

type DescribePhoneTwiceTelVerifyRequest struct {
	// The authorization code.
	//
	// >  On the **My Applications*	- page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
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
	// 139*******
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
	// The registration time. Specify the time in the yyyy-MM-dd HH:mm:ss format. This time is the service registration time of the mobile phone user. If the service registration time is later than the time when the phone number is assigned by a carrier, it indicates that the phone number is not a reassigned number. Otherwise, the phone number is a reassigned number.
	//
	// >
	//
	// 	- If a carrier allocates a single number multiple times, the system will determine whether the phone number is a reassigned number based on the time when the carrier last allocated the phone number.
	//
	// 	- The service registration time must be later than 00:00:00 on January 1, 1970.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePhoneTwiceTelVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneTwiceTelVerifyRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneTwiceTelVerifyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetAuthCode(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetInputNumber(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetMask(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetResourceOwnerAccount(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetResourceOwnerId(v int64) *DescribePhoneTwiceTelVerifyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) SetStartTime(v string) *DescribePhoneTwiceTelVerifyRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePhoneTwiceTelVerifyRequest) Validate() error {
	return dara.Validate(s)
}
