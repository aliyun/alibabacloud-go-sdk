// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleFluctuationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeHitRuleFluctuationRequest
	GetLang() *string
	SetEventCodes(v string) *DescribeHitRuleFluctuationRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeHitRuleFluctuationRequest
	GetRegId() *string
	SetRuleStatus(v string) *DescribeHitRuleFluctuationRequest
	GetRuleStatus() *string
}

type DescribeHitRuleFluctuationRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event codes, separated by commas (,).
	//
	// example:
	//
	// de_ahqhsw7665,de_agbzfi5134
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy status
	//
	// example:
	//
	// DRAFT
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
}

func (s DescribeHitRuleFluctuationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleFluctuationRequest) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleFluctuationRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHitRuleFluctuationRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeHitRuleFluctuationRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeHitRuleFluctuationRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeHitRuleFluctuationRequest) SetLang(v string) *DescribeHitRuleFluctuationRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHitRuleFluctuationRequest) SetEventCodes(v string) *DescribeHitRuleFluctuationRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeHitRuleFluctuationRequest) SetRegId(v string) *DescribeHitRuleFluctuationRequest {
	s.RegId = &v
	return s
}

func (s *DescribeHitRuleFluctuationRequest) SetRuleStatus(v string) *DescribeHitRuleFluctuationRequest {
	s.RuleStatus = &v
	return s
}

func (s *DescribeHitRuleFluctuationRequest) Validate() error {
	return dara.Validate(s)
}
