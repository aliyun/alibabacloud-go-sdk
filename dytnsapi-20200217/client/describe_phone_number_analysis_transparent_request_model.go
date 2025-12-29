// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisTransparentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberAnalysisTransparentRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneNumberAnalysisTransparentRequest
	GetInputNumber() *string
	SetIp(v string) *DescribePhoneNumberAnalysisTransparentRequest
	GetIp() *string
	SetNumberType(v string) *DescribePhoneNumberAnalysisTransparentRequest
	GetNumberType() *string
	SetOwnerId(v int64) *DescribePhoneNumberAnalysisTransparentRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisTransparentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisTransparentRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberAnalysisTransparentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 187****5620
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	NumberType           *string `json:"NumberType,omitempty" xml:"NumberType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisTransparentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisTransparentRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetNumberType() *string {
	return s.NumberType
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetIp(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.Ip = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetNumberType(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.NumberType = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisTransparentRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisTransparentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisTransparentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisTransparentRequest) Validate() error {
	return dara.Validate(s)
}
