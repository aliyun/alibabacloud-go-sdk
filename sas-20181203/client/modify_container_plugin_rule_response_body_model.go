// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerPluginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyContainerPluginRuleResponseBodyData) *ModifyContainerPluginRuleResponseBody
	GetData() *ModifyContainerPluginRuleResponseBodyData
	SetRequestId(v string) *ModifyContainerPluginRuleResponseBody
	GetRequestId() *string
}

type ModifyContainerPluginRuleResponseBody struct {
	// The defense rule against container escapes.
	Data *ModifyContainerPluginRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContainerPluginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerPluginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerPluginRuleResponseBody) GetData() *ModifyContainerPluginRuleResponseBodyData {
	return s.Data
}

func (s *ModifyContainerPluginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContainerPluginRuleResponseBody) SetData(v *ModifyContainerPluginRuleResponseBodyData) *ModifyContainerPluginRuleResponseBody {
	s.Data = v
	return s
}

func (s *ModifyContainerPluginRuleResponseBody) SetRequestId(v string) *ModifyContainerPluginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContainerPluginRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyContainerPluginRuleResponseBodyData struct {
	// The ID of the rule.
	//
	// example:
	//
	// 600640
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test555
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The switch ID.
	//
	// example:
	//
	// USER-CONTAINER-RULE-SWITCH-TYPE_xxx
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s ModifyContainerPluginRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerPluginRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyContainerPluginRuleResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyContainerPluginRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyContainerPluginRuleResponseBodyData) GetSwitchId() *string {
	return s.SwitchId
}

func (s *ModifyContainerPluginRuleResponseBodyData) SetRuleId(v int64) *ModifyContainerPluginRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *ModifyContainerPluginRuleResponseBodyData) SetRuleName(v string) *ModifyContainerPluginRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *ModifyContainerPluginRuleResponseBodyData) SetSwitchId(v string) *ModifyContainerPluginRuleResponseBodyData {
	s.SwitchId = &v
	return s
}

func (s *ModifyContainerPluginRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
