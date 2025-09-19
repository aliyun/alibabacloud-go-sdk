// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerPluginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyContainerPluginRuleRequest
	GetLang() *string
	SetMode(v int32) *ModifyContainerPluginRuleRequest
	GetMode() *int32
	SetRuleId(v int32) *ModifyContainerPluginRuleRequest
	GetRuleId() *int32
	SetRuleName(v string) *ModifyContainerPluginRuleRequest
	GetRuleName() *string
	SetRuleType(v int32) *ModifyContainerPluginRuleRequest
	GetRuleType() *int32
	SetSelectedPolicy(v []*string) *ModifyContainerPluginRuleRequest
	GetSelectedPolicy() []*string
	SetWhiteImages(v []*string) *ModifyContainerPluginRuleRequest
	GetWhiteImages() []*string
}

type ModifyContainerPluginRuleRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The action mode of the rule. Valid values:
	//
	// 	- **1**: alerts
	//
	// 	- **2**: block
	//
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100012
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **0**: user-defined rule
	//
	// 	- **1**: built-in rule
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The rule items.
	SelectedPolicy []*string `json:"SelectedPolicy,omitempty" xml:"SelectedPolicy,omitempty" type:"Repeated"`
	// The images that are added to the whitelist.
	WhiteImages []*string `json:"WhiteImages,omitempty" xml:"WhiteImages,omitempty" type:"Repeated"`
}

func (s ModifyContainerPluginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerPluginRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerPluginRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyContainerPluginRuleRequest) GetMode() *int32 {
	return s.Mode
}

func (s *ModifyContainerPluginRuleRequest) GetRuleId() *int32 {
	return s.RuleId
}

func (s *ModifyContainerPluginRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyContainerPluginRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ModifyContainerPluginRuleRequest) GetSelectedPolicy() []*string {
	return s.SelectedPolicy
}

func (s *ModifyContainerPluginRuleRequest) GetWhiteImages() []*string {
	return s.WhiteImages
}

func (s *ModifyContainerPluginRuleRequest) SetLang(v string) *ModifyContainerPluginRuleRequest {
	s.Lang = &v
	return s
}

func (s *ModifyContainerPluginRuleRequest) SetMode(v int32) *ModifyContainerPluginRuleRequest {
	s.Mode = &v
	return s
}

func (s *ModifyContainerPluginRuleRequest) SetRuleId(v int32) *ModifyContainerPluginRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyContainerPluginRuleRequest) SetRuleName(v string) *ModifyContainerPluginRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyContainerPluginRuleRequest) SetRuleType(v int32) *ModifyContainerPluginRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyContainerPluginRuleRequest) SetSelectedPolicy(v []*string) *ModifyContainerPluginRuleRequest {
	s.SelectedPolicy = v
	return s
}

func (s *ModifyContainerPluginRuleRequest) SetWhiteImages(v []*string) *ModifyContainerPluginRuleRequest {
	s.WhiteImages = v
	return s
}

func (s *ModifyContainerPluginRuleRequest) Validate() error {
	return dara.Validate(s)
}
