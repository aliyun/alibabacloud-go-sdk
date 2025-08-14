// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthRulePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAuthRulePageListRequest
	GetLang() *string
	SetEventCode(v string) *DescribeAuthRulePageListRequest
	GetEventCode() *string
	SetRegId(v string) *DescribeAuthRulePageListRequest
	GetRegId() *string
	SetRuleName(v string) *DescribeAuthRulePageListRequest
	GetRuleName() *string
	SetStatus(v string) *DescribeAuthRulePageListRequest
	GetStatus() *string
}

type DescribeAuthRulePageListRequest struct {
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
	// de_arcehq4370
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 分析中心事件测试_策略01
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// Status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeAuthRulePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthRulePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthRulePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAuthRulePageListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeAuthRulePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAuthRulePageListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeAuthRulePageListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAuthRulePageListRequest) SetLang(v string) *DescribeAuthRulePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAuthRulePageListRequest) SetEventCode(v string) *DescribeAuthRulePageListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeAuthRulePageListRequest) SetRegId(v string) *DescribeAuthRulePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAuthRulePageListRequest) SetRuleName(v string) *DescribeAuthRulePageListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeAuthRulePageListRequest) SetStatus(v string) *DescribeAuthRulePageListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAuthRulePageListRequest) Validate() error {
	return dara.Validate(s)
}
