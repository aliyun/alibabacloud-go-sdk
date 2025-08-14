// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteRuleRequest
	GetLang() *string
	SetConsoleRuleId(v int64) *DeleteRuleRequest
	GetConsoleRuleId() *int64
	SetRegId(v string) *DeleteRuleRequest
	GetRegId() *string
	SetRuleId(v string) *DeleteRuleRequest
	GetRuleId() *string
	SetRuleVersionId(v int64) *DeleteRuleRequest
	GetRuleVersionId() *int64
}

type DeleteRuleRequest struct {
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
	// Policy primary key ID
	//
	// example:
	//
	// 7035
	ConsoleRuleId *int64 `json:"consoleRuleId,omitempty" xml:"consoleRuleId,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy ID
	//
	// example:
	//
	// 102059
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// Policy version primary key ID
	//
	// example:
	//
	// 10203
	RuleVersionId *int64 `json:"ruleVersionId,omitempty" xml:"ruleVersionId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteRuleRequest) GetConsoleRuleId() *int64 {
	return s.ConsoleRuleId
}

func (s *DeleteRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteRuleRequest) GetRuleVersionId() *int64 {
	return s.RuleVersionId
}

func (s *DeleteRuleRequest) SetLang(v string) *DeleteRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRuleRequest) SetConsoleRuleId(v int64) *DeleteRuleRequest {
	s.ConsoleRuleId = &v
	return s
}

func (s *DeleteRuleRequest) SetRegId(v string) *DeleteRuleRequest {
	s.RegId = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v string) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleVersionId(v int64) *DeleteRuleRequest {
	s.RuleVersionId = &v
	return s
}

func (s *DeleteRuleRequest) Validate() error {
	return dara.Validate(s)
}
