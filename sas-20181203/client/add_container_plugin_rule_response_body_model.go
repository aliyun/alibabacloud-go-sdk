// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerPluginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddContainerPluginRuleResponseBodyData) *AddContainerPluginRuleResponseBody
	GetData() *AddContainerPluginRuleResponseBodyData
	SetRequestId(v string) *AddContainerPluginRuleResponseBody
	GetRequestId() *string
}

type AddContainerPluginRuleResponseBody struct {
	// The data returned.
	Data *AddContainerPluginRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddContainerPluginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddContainerPluginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerPluginRuleResponseBody) GetData() *AddContainerPluginRuleResponseBodyData {
	return s.Data
}

func (s *AddContainerPluginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddContainerPluginRuleResponseBody) SetData(v *AddContainerPluginRuleResponseBodyData) *AddContainerPluginRuleResponseBody {
	s.Data = v
	return s
}

func (s *AddContainerPluginRuleResponseBody) SetRequestId(v string) *AddContainerPluginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContainerPluginRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddContainerPluginRuleResponseBodyData struct {
	// The ID of the rule.
	//
	// example:
	//
	// 219
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test2
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the switch.
	//
	// example:
	//
	// USER-CONTAINER-RULE-SWITCH-TYPE_xxxx
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
}

func (s AddContainerPluginRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddContainerPluginRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddContainerPluginRuleResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *AddContainerPluginRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *AddContainerPluginRuleResponseBodyData) GetSwitchId() *string {
	return s.SwitchId
}

func (s *AddContainerPluginRuleResponseBodyData) SetRuleId(v int64) *AddContainerPluginRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *AddContainerPluginRuleResponseBodyData) SetRuleName(v string) *AddContainerPluginRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *AddContainerPluginRuleResponseBodyData) SetSwitchId(v string) *AddContainerPluginRuleResponseBodyData {
	s.SwitchId = &v
	return s
}

func (s *AddContainerPluginRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
