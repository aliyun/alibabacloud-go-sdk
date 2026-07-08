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
	// > On the **My Applications*	- page in the [Phone Number Verification Service console](https://dytns.console.aliyun.com/analysis/apply), get the Authorization ID. This ID is your authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to be queried. The number can be a mobile phone number or an encrypted string.
	//
	// - If the value of **Mask*	- is **NORMAL**, **InputNumber*	- is an 11-digit mobile phone number.
	//
	// - If the value of **Mask*	- is **MD5**, **InputNumber*	- is a 32-bit encrypted string.
	//
	// - If the value of **Mask*	- is **SHA256**, **InputNumber*	- is a 64-bit encrypted string.
	//
	// 	Notice:
	//
	// The encrypted string is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139*******
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// - **NORMAL**: The phone number is not encrypted.
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
	// The registration time of a phone number. The time must be in the `yyyy-MM-dd HH:mm:ss` format. The value of this parameter is the registration time of a mobile phone user in your business. If the registration time is later than the time when a carrier assigns a number, the number is not a recycled number. Otherwise, the number is a recycled number.
	//
	// > - If a phone number is assigned for multiple times, the system uses the last assignment time as the criterion.
	//
	// >
	//
	// > - The registration time must be later than `1970-01-01 00:00:00`.
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
