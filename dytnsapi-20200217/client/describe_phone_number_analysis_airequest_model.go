// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhoneNumberAnalysisAIRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *DescribePhoneNumberAnalysisAIRequest
	GetAuthCode() *string
	SetInputNumber(v string) *DescribePhoneNumberAnalysisAIRequest
	GetInputNumber() *string
	SetModelConfig(v string) *DescribePhoneNumberAnalysisAIRequest
	GetModelConfig() *string
	SetOwnerId(v int64) *DescribePhoneNumberAnalysisAIRequest
	GetOwnerId() *int64
	SetRate(v int64) *DescribePhoneNumberAnalysisAIRequest
	GetRate() *int64
	SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisAIRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisAIRequest
	GetResourceOwnerId() *int64
}

type DescribePhoneNumberAnalysisAIRequest struct {
	// The authorization code.
	//
	// > In **Cell Phone Number Service*	- -> [**Tag Square**](https://dytns.console.aliyun.com/analysis/square), select a tag and submit a usage application. After the application is approved, you will receive the authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// HwD***nG
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The phone number to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 187****5620
	InputNumber *string `json:"InputNumber,omitempty" xml:"InputNumber,omitempty"`
	// The model parameter configuration (required by some tag capabilities).
	//
	// example:
	//
	// {"trainingJobId": "17**********48"}
	ModelConfig *string `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number score threshold. Valid values: **0 to 100**.
	//
	// > Whether to accept the specified score threshold is determined by the server. When the specified score threshold is not accepted, the data entered in this field is invalid.
	//
	// example:
	//
	// 96
	Rate                 *int64  `json:"Rate,omitempty" xml:"Rate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePhoneNumberAnalysisAIRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhoneNumberAnalysisAIRequest) GoString() string {
	return s.String()
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetInputNumber() *string {
	return s.InputNumber
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetModelConfig() *string {
	return s.ModelConfig
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetRate() *int64 {
	return s.Rate
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhoneNumberAnalysisAIRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetAuthCode(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.AuthCode = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetInputNumber(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.InputNumber = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetModelConfig(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.ModelConfig = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetOwnerId(v int64) *DescribePhoneNumberAnalysisAIRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetRate(v int64) *DescribePhoneNumberAnalysisAIRequest {
	s.Rate = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetResourceOwnerAccount(v string) *DescribePhoneNumberAnalysisAIRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) SetResourceOwnerId(v int64) *DescribePhoneNumberAnalysisAIRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhoneNumberAnalysisAIRequest) Validate() error {
	return dara.Validate(s)
}
