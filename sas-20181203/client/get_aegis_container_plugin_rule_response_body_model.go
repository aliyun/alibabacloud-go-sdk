// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAegisContainerPluginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAegisContainerPluginRuleResponseBodyData) *GetAegisContainerPluginRuleResponseBody
	GetData() *GetAegisContainerPluginRuleResponseBodyData
	SetRequestId(v string) *GetAegisContainerPluginRuleResponseBody
	GetRequestId() *string
}

type GetAegisContainerPluginRuleResponseBody struct {
	// The returned data.
	Data *GetAegisContainerPluginRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 09969D2C-***0DEF8BF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAegisContainerPluginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleResponseBody) GetData() *GetAegisContainerPluginRuleResponseBodyData {
	return s.Data
}

func (s *GetAegisContainerPluginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAegisContainerPluginRuleResponseBody) SetData(v *GetAegisContainerPluginRuleResponseBodyData) *GetAegisContainerPluginRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBody) SetRequestId(v string) *GetAegisContainerPluginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAegisContainerPluginRuleResponseBodyData struct {
	// The timestamp when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1671607025000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the rule was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1671607025000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The action mode of the rule. Valid values:
	//
	// 	- **0**: allows escape behavior.
	//
	// 	- **1**: triggers alerts.
	//
	// 	- **2**: blocks escape behavior.
	//
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// autoTest
	RuleDescription *string `json:"RuleDescription,omitempty" xml:"RuleDescription,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 21**
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// auto_test_rule-EmzIXZ
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The template ID of the rule.
	//
	// example:
	//
	// 100**
	RuleTemplateId *string `json:"RuleTemplateId,omitempty" xml:"RuleTemplateId,omitempty"`
	// The template name of the rule.
	//
	// example:
	//
	// template01
	RuleTemplateName *string `json:"RuleTemplateName,omitempty" xml:"RuleTemplateName,omitempty"`
	// The rule items.
	SelectedPolicy []*string `json:"SelectedPolicy,omitempty" xml:"SelectedPolicy,omitempty" type:"Repeated"`
	// The ID of the switch.
	//
	// example:
	//
	// USER-CONTAINER-RULE-SWITCH-TYPE_***
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The images that are added to the whitelist.
	WhiteImages []*string `json:"WhiteImages,omitempty" xml:"WhiteImages,omitempty" type:"Repeated"`
}

func (s GetAegisContainerPluginRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAegisContainerPluginRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetMode() *int32 {
	return s.Mode
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetRuleTemplateId() *string {
	return s.RuleTemplateId
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetRuleTemplateName() *string {
	return s.RuleTemplateName
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetSelectedPolicy() []*string {
	return s.SelectedPolicy
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetSwitchId() *string {
	return s.SwitchId
}

func (s *GetAegisContainerPluginRuleResponseBodyData) GetWhiteImages() []*string {
	return s.WhiteImages
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetGmtCreate(v int64) *GetAegisContainerPluginRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetGmtModified(v int64) *GetAegisContainerPluginRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetMode(v int32) *GetAegisContainerPluginRuleResponseBodyData {
	s.Mode = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetRuleDescription(v string) *GetAegisContainerPluginRuleResponseBodyData {
	s.RuleDescription = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetRuleId(v int64) *GetAegisContainerPluginRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetRuleName(v string) *GetAegisContainerPluginRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetRuleTemplateId(v string) *GetAegisContainerPluginRuleResponseBodyData {
	s.RuleTemplateId = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetRuleTemplateName(v string) *GetAegisContainerPluginRuleResponseBodyData {
	s.RuleTemplateName = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetSelectedPolicy(v []*string) *GetAegisContainerPluginRuleResponseBodyData {
	s.SelectedPolicy = v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetSwitchId(v string) *GetAegisContainerPluginRuleResponseBodyData {
	s.SwitchId = &v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) SetWhiteImages(v []*string) *GetAegisContainerPluginRuleResponseBodyData {
	s.WhiteImages = v
	return s
}

func (s *GetAegisContainerPluginRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
