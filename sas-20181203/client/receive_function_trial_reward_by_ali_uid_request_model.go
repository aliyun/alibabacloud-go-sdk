// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveFunctionTrialRewardByAliUidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *ReceiveFunctionTrialRewardByAliUidRequest
	GetFunctionName() *string
	SetLang(v string) *ReceiveFunctionTrialRewardByAliUidRequest
	GetLang() *string
}

type ReceiveFunctionTrialRewardByAliUidRequest struct {
	// The name of the feature for which you want to apply for a free trial. Valid values:
	//
	// 	- **trail_honeypot_reward**: cloud honeypot
	//
	// 	- **trail_file_detect_api_reward**: SDK for malicious file detection
	//
	// example:
	//
	// trail_honeypot_reward
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ReceiveFunctionTrialRewardByAliUidRequest) String() string {
	return dara.Prettify(s)
}

func (s ReceiveFunctionTrialRewardByAliUidRequest) GoString() string {
	return s.String()
}

func (s *ReceiveFunctionTrialRewardByAliUidRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ReceiveFunctionTrialRewardByAliUidRequest) GetLang() *string {
	return s.Lang
}

func (s *ReceiveFunctionTrialRewardByAliUidRequest) SetFunctionName(v string) *ReceiveFunctionTrialRewardByAliUidRequest {
	s.FunctionName = &v
	return s
}

func (s *ReceiveFunctionTrialRewardByAliUidRequest) SetLang(v string) *ReceiveFunctionTrialRewardByAliUidRequest {
	s.Lang = &v
	return s
}

func (s *ReceiveFunctionTrialRewardByAliUidRequest) Validate() error {
	return dara.Validate(s)
}
