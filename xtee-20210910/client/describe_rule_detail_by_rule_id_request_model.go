// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleDetailByRuleIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleDetailByRuleIdRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *DescribeRuleDetailByRuleIdRequest
	GetConsoleRuleId() *int64
	SetRegId(v string) *DescribeRuleDetailByRuleIdRequest
	GetRegId() *string
	SetRuleId(v string) *DescribeRuleDetailByRuleIdRequest
	GetRuleId() *string
	SetRuleVersionId(v int64) *DescribeRuleDetailByRuleIdRequest
	GetRuleVersionId() *int64
}

type DescribeRuleDetailByRuleIdRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID of the policy.
	//
	// example:
	//
	// 7110
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// 102059
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Primary key ID of the policy version.
	//
	// example:
	//
	// 10203
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s DescribeRuleDetailByRuleIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleDetailByRuleIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleDetailByRuleIdRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleDetailByRuleIdRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DescribeRuleDetailByRuleIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleDetailByRuleIdRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeRuleDetailByRuleIdRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DescribeRuleDetailByRuleIdRequest) SetLang(v string) *DescribeRuleDetailByRuleIdRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdRequest) SetConsoleRuleId(v int64) *DescribeRuleDetailByRuleIdRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdRequest) SetRegId(v string) *DescribeRuleDetailByRuleIdRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdRequest) SetRuleId(v string) *DescribeRuleDetailByRuleIdRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdRequest) SetRuleVersionId(v int64) *DescribeRuleDetailByRuleIdRequest {
	s.RuleVersionId = &v
	return s
}

func (s *DescribeRuleDetailByRuleIdRequest) Validate() error {
	return dara.Validate(s)
}
