// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContainerPluginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteContainerPluginRuleRequest
	GetLang() *string
	SetRuleId(v int32) *DeleteContainerPluginRuleRequest
	GetRuleId() *int32
}

type DeleteContainerPluginRuleRequest struct {
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The rule ID. You can call the addContainerWebDefenseRule operation to obtain the rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200022
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteContainerPluginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContainerPluginRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerPluginRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteContainerPluginRuleRequest) GetRuleId() *int32 {
	return s.RuleId
}

func (s *DeleteContainerPluginRuleRequest) SetLang(v string) *DeleteContainerPluginRuleRequest {
	s.Lang = &v
	return s
}

func (s *DeleteContainerPluginRuleRequest) SetRuleId(v int32) *DeleteContainerPluginRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteContainerPluginRuleRequest) Validate() error {
	return dara.Validate(s)
}
