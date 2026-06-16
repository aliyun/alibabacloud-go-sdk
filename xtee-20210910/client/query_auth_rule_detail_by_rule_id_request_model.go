// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuthRuleDetailByRuleIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryAuthRuleDetailByRuleIdRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *QueryAuthRuleDetailByRuleIdRequest
	GetConsoleRuleId() *int64
	SetRegId(v string) *QueryAuthRuleDetailByRuleIdRequest
	GetRegId() *string
	SetRuleId(v string) *QueryAuthRuleDetailByRuleIdRequest
	GetRuleId() *string
	SetRuleVersionId(v int64) *QueryAuthRuleDetailByRuleIdRequest
	GetRuleVersionId() *int64
}

type QueryAuthRuleDetailByRuleIdRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The primary key ID of the policy.
	//
	// example:
	//
	// 6843
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// The region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// 102224
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The primary key ID of the policy version.
	//
	// example:
	//
	// 11519
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s QueryAuthRuleDetailByRuleIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAuthRuleDetailByRuleIdRequest) GoString() string {
	return s.String()
}

func (s *QueryAuthRuleDetailByRuleIdRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryAuthRuleDetailByRuleIdRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *QueryAuthRuleDetailByRuleIdRequest) GetRegId() *string {
	return s.RegId
}

func (s *QueryAuthRuleDetailByRuleIdRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *QueryAuthRuleDetailByRuleIdRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *QueryAuthRuleDetailByRuleIdRequest) SetLang(v string) *QueryAuthRuleDetailByRuleIdRequest {
	s.Lang = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdRequest) SetConsoleRuleId(v int64) *QueryAuthRuleDetailByRuleIdRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdRequest) SetRegId(v string) *QueryAuthRuleDetailByRuleIdRequest {
	s.RegId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdRequest) SetRuleId(v string) *QueryAuthRuleDetailByRuleIdRequest {
	s.RuleId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdRequest) SetRuleVersionId(v int64) *QueryAuthRuleDetailByRuleIdRequest {
	s.RuleVersionId = &v
	return s
}

func (s *QueryAuthRuleDetailByRuleIdRequest) Validate() error {
	return dara.Validate(s)
}
