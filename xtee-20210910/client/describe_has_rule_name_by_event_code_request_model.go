// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHasRuleNameByEventCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeHasRuleNameByEventCodeRequest
	GetLang() *string
	SetEventCode(v string) *DescribeHasRuleNameByEventCodeRequest
	GetEventCode() *string
	SetExcludeRuleId(v string) *DescribeHasRuleNameByEventCodeRequest
	GetExcludeRuleId() *string
	SetRegId(v string) *DescribeHasRuleNameByEventCodeRequest
	GetRegId() *string
	SetRuleName(v string) *DescribeHasRuleNameByEventCodeRequest
	GetRuleName() *string
}

type DescribeHasRuleNameByEventCodeRequest struct {
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
	// Event code
	//
	// example:
	//
	// de_atvmlf7412
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Excluded policy ID
	//
	// example:
	//
	// 10621
	ExcludeRuleId *string `json:"excludeRuleId,omitempty" xml:"excludeRuleId,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 非常见设备
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
}

func (s DescribeHasRuleNameByEventCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHasRuleNameByEventCodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeHasRuleNameByEventCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHasRuleNameByEventCodeRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeHasRuleNameByEventCodeRequest) GetExcludeRuleId() *string {
	return s.ExcludeRuleId
}

func (s *DescribeHasRuleNameByEventCodeRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeHasRuleNameByEventCodeRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeHasRuleNameByEventCodeRequest) SetLang(v string) *DescribeHasRuleNameByEventCodeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHasRuleNameByEventCodeRequest) SetEventCode(v string) *DescribeHasRuleNameByEventCodeRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeHasRuleNameByEventCodeRequest) SetExcludeRuleId(v string) *DescribeHasRuleNameByEventCodeRequest {
	s.ExcludeRuleId = &v
	return s
}

func (s *DescribeHasRuleNameByEventCodeRequest) SetRegId(v string) *DescribeHasRuleNameByEventCodeRequest {
	s.RegId = &v
	return s
}

func (s *DescribeHasRuleNameByEventCodeRequest) SetRuleName(v string) *DescribeHasRuleNameByEventCodeRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeHasRuleNameByEventCodeRequest) Validate() error {
	return dara.Validate(s)
}
