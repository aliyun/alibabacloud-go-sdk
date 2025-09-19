// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerPluginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *AddContainerPluginRuleRequest
	GetLang() *string
	SetMode(v int32) *AddContainerPluginRuleRequest
	GetMode() *int32
	SetRuleName(v string) *AddContainerPluginRuleRequest
	GetRuleName() *string
	SetRuleTemplateId(v int32) *AddContainerPluginRuleRequest
	GetRuleTemplateId() *int32
	SetRuleType(v int32) *AddContainerPluginRuleRequest
	GetRuleType() *int32
	SetSelectedPolicy(v []*string) *AddContainerPluginRuleRequest
	GetSelectedPolicy() []*string
	SetWhiteImages(v []*string) *AddContainerPluginRuleRequest
	GetWhiteImages() []*string
}

type AddContainerPluginRuleRequest struct {
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
	// The action that you want to specify for the rule. Valid values:
	//
	// 	- **1**: triggers alerts.
	//
	// 	- **2**: blocks escapes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the rule. The name must be 1 to 80 characters in length, and can contain letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_). The names of rules that are created for the same user must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// tyest111
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The ID of the rule template. You can call the ListSystemClientRules operation to query the ID of the rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86863
	RuleTemplateId *int32 `json:"RuleTemplateId,omitempty" xml:"RuleTemplateId,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **0**: custom rule
	//
	// 	- **1**: system rule
	//
	// example:
	//
	// 0
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The check items that are enabled for the rule.
	//
	// This parameter is required.
	SelectedPolicy []*string `json:"SelectedPolicy,omitempty" xml:"SelectedPolicy,omitempty" type:"Repeated"`
	// The images that are added to the whitelist.
	WhiteImages []*string `json:"WhiteImages,omitempty" xml:"WhiteImages,omitempty" type:"Repeated"`
}

func (s AddContainerPluginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContainerPluginRuleRequest) GoString() string {
	return s.String()
}

func (s *AddContainerPluginRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *AddContainerPluginRuleRequest) GetMode() *int32 {
	return s.Mode
}

func (s *AddContainerPluginRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddContainerPluginRuleRequest) GetRuleTemplateId() *int32 {
	return s.RuleTemplateId
}

func (s *AddContainerPluginRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *AddContainerPluginRuleRequest) GetSelectedPolicy() []*string {
	return s.SelectedPolicy
}

func (s *AddContainerPluginRuleRequest) GetWhiteImages() []*string {
	return s.WhiteImages
}

func (s *AddContainerPluginRuleRequest) SetLang(v string) *AddContainerPluginRuleRequest {
	s.Lang = &v
	return s
}

func (s *AddContainerPluginRuleRequest) SetMode(v int32) *AddContainerPluginRuleRequest {
	s.Mode = &v
	return s
}

func (s *AddContainerPluginRuleRequest) SetRuleName(v string) *AddContainerPluginRuleRequest {
	s.RuleName = &v
	return s
}

func (s *AddContainerPluginRuleRequest) SetRuleTemplateId(v int32) *AddContainerPluginRuleRequest {
	s.RuleTemplateId = &v
	return s
}

func (s *AddContainerPluginRuleRequest) SetRuleType(v int32) *AddContainerPluginRuleRequest {
	s.RuleType = &v
	return s
}

func (s *AddContainerPluginRuleRequest) SetSelectedPolicy(v []*string) *AddContainerPluginRuleRequest {
	s.SelectedPolicy = v
	return s
}

func (s *AddContainerPluginRuleRequest) SetWhiteImages(v []*string) *AddContainerPluginRuleRequest {
	s.WhiteImages = v
	return s
}

func (s *AddContainerPluginRuleRequest) Validate() error {
	return dara.Validate(s)
}
