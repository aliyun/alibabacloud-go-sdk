// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberOperatorAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeRequest
	GetAuthCode() *string
	SetFlowName(v string) *DescribePhoneNumberOperatorAttributeRequest
	GetFlowName() *string
	SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeRequest
	GetInputNumber() *string
	SetMask(v string) *DescribePhoneNumberOperatorAttributeRequest
	GetMask() *string
	SetOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberOperatorAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest
	GetResourceOwnerId() *int64
	SetResultCount(v string) *DescribePhoneNumberOperatorAttributeRequest
	GetResultCount() *string
}

type DescribePhoneNumberOperatorAttributeRequest struct {
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
	// example:
	//
	// 示例值
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
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
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// 	- **NORMAL**: The phone number is not encrypted.
	//
	// 	- **MD5**: The phone number is MD5-encrypted.
	//
	// 	- **SHA256**: The phone number is SHA256-encrypted.
	//
	// > Letters in the string must be uppercase.
	//
	// This parameter is required.
	//
	// example:
	//
	// MD5
	Mask                 *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ResultCount *string `json:"ResultCount,omitempty" xml:"ResultCount,omitempty"`
}

func (s DescribePhoneNumberOperatorAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberOperatorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberOperatorAttributeRequest) GetResultCount() *string {
	return s.ResultCount
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetAuthCode(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetFlowName(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.FlowName = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetInputNumber(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetMask(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) SetResultCount(v string) *DescribePhoneNumberOperatorAttributeRequest {
	s.ResultCount = &v
	return s
}

func (s *DescribePhoneNumberOperatorAttributeRequest) Validate() error {
	return dara.Validate(s)
}
