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
	// > On the **My Applications*	- page of the [Cell Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), obtain the authorization ID, which is the authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// -
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The phone number that you want to query.
	//
	// - If Mask is set to NORMAL, this field is an 11-digit phone number.
	//
	// - If Mask is set to MD5, this field is a 32-character encrypted string.
	//
	// - If Mask is set to SHA256, this field is a 64-character encrypted string.
	//
	// - If Mask is set to SM3, this field is a 64-character encrypted string.
	//
	// 	Notice: The letters in the encrypted string are not case-sensitive.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****1234
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// - **NORMAL**: no encryption
	//
	// - **MD5**: MD5 encryption
	//
	// - **SHA256**: SHA256 encryption
	//
	// - **SM3**: SM3 encryption
	//
	// 	Notice: All letters in the string must be uppercase.</notice>
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
	// A system parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// -
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
