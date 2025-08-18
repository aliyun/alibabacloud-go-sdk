// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListWaitingRoomRulesResponseBody
	GetRequestId() *string
	SetWaitingRoomRules(v []*ListWaitingRoomRulesResponseBodyWaitingRoomRules) *ListWaitingRoomRulesResponseBody
	GetWaitingRoomRules() []*ListWaitingRoomRulesResponseBodyWaitingRoomRules
}

type ListWaitingRoomRulesResponseBody struct {
	// Request ID, used for tracking the call status.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of waiting room bypass rules.
	WaitingRoomRules []*ListWaitingRoomRulesResponseBodyWaitingRoomRules `json:"WaitingRoomRules,omitempty" xml:"WaitingRoomRules,omitempty" type:"Repeated"`
}

func (s ListWaitingRoomRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWaitingRoomRulesResponseBody) GetWaitingRoomRules() []*ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	return s.WaitingRoomRules
}

func (s *ListWaitingRoomRulesResponseBody) SetRequestId(v string) *ListWaitingRoomRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBody) SetWaitingRoomRules(v []*ListWaitingRoomRulesResponseBodyWaitingRoomRules) *ListWaitingRoomRulesResponseBody {
	s.WaitingRoomRules = v
	return s
}

func (s *ListWaitingRoomRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWaitingRoomRulesResponseBodyWaitingRoomRules struct {
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: set the value to true
	//
	// - Match specific requests: set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.request.uri.path.file_name eq \\"jpg\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding global configuration. Value range:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter does not need to be set when adding global configuration.
	//
	// example:
	//
	// ip
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 37286782688****
	WaitingRoomRuleId *int64 `json:"WaitingRoomRuleId,omitempty" xml:"WaitingRoomRuleId,omitempty"`
}

func (s ListWaitingRoomRulesResponseBodyWaitingRoomRules) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomRulesResponseBodyWaitingRoomRules) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetRule() *string {
	return s.Rule
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) GetWaitingRoomRuleId() *int64 {
	return s.WaitingRoomRuleId
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRule(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.Rule = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRuleEnable(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.RuleEnable = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetRuleName(v string) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.RuleName = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) SetWaitingRoomRuleId(v int64) *ListWaitingRoomRulesResponseBodyWaitingRoomRules {
	s.WaitingRoomRuleId = &v
	return s
}

func (s *ListWaitingRoomRulesResponseBodyWaitingRoomRules) Validate() error {
	return dara.Validate(s)
}
