// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberAnalysisRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneNumberAnalysisRequest
	GetInputNumber() *string
	SetMask(v string) *DescribePhoneNumberAnalysisRequest
	GetMask() *string
	SetNumberType(v int64) *DescribePhoneNumberAnalysisRequest
	GetNumberType() *int64
	SetOwnerId(v int64) *DescribePhoneNumberAnalysisRequest
	GetOwnerId() *int64
	SetRate(v int64) *DescribePhoneNumberAnalysisRequest
	GetRate() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberAnalysisRequest struct {
	// The authorization code.
	//
	// > Log on to the [Phone Number Service console](https://dytns.console.aliyun.com/analysis/apply), go to the **My Applications*	- page, and obtain the authorization ID, which is the authorization code.
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
	// 187****5620
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The encryption method of the phone number. Valid values:
	//
	// - **NORMAL**: no encryption
	//
	// - **MD5**
	//
	// - **SHA256**
	//
	// example:
	//
	// NORMAL
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The type of the phone number. Valid values:
	//
	// - **0**: mobile phone number
	//
	// - **1**: mobile IMEI number
	//
	// example:
	//
	// 0
	NumberType *int64 `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId    *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The score threshold of the phone number. Valid values: **0 to 100**.
	//
	// 	Notice: Whether the specified score threshold is accepted is determined by the server. When the specified score threshold is not accepted, the data entered in this field is invalid.
	//
	// example:
	//
	// 10
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberAnalysisRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberAnalysisRequest) GetMask() *string {
	return s.Mask
}

func (s *DescribePhoneNumberAnalysisRequest) GetNumberType() *int64 {
	return s.NumberType
}

func (s *DescribePhoneNumberAnalysisRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberAnalysisRequest) GetRate() *int64 {
	return s.Rate
}

func (s *DescribePhoneNumberAnalysisRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberAnalysisRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberAnalysisRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetMask(v string) *DescribePhoneNumberAnalysisRequest {
	s.Mask = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetNumberType(v int64) *DescribePhoneNumberAnalysisRequest {
	s.NumberType = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetRate(v int64) *DescribePhoneNumberAnalysisRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
