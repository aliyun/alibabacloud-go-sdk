// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisPaiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberAnalysisPaiRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneNumberAnalysisPaiRequest
	GetInputNumber() *string
	SetModelConfig(v string) *DescribePhoneNumberAnalysisPaiRequest
	GetModelConfig() *string
	SetOwnerId(v int64) *DescribePhoneNumberAnalysisPaiRequest
	GetOwnerId() *int64
	SetRate(v int64) *DescribePhoneNumberAnalysisPaiRequest
	GetRate() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisPaiRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisPaiRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberAnalysisPaiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 16
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisPaiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisPaiRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetRate() *int64 {
	return s.Rate
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberAnalysisPaiRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetModelConfig(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.ModelConfig = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisPaiRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetRate(v int64) *DescribePhoneNumberAnalysisPaiRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisPaiRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisPaiRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisPaiRequest) Validate() error {
	return dara.Validate(s)
}
