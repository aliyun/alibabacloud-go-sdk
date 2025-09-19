// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSensitiveDefineRuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SetSensitiveDefineRuleConfigRequest
	GetConfig() *string
	SetEnableNewRule(v int32) *SetSensitiveDefineRuleConfigRequest
	GetEnableNewRule() *int32
	SetSource(v string) *SetSensitiveDefineRuleConfigRequest
	GetSource() *string
}

type SetSensitiveDefineRuleConfigRequest struct {
	// The configurations of the custom check rule. The value is in the JSON format. Valid values of keys:
	//
	// 	- **classKey**: the category keyword of the check rule.
	//
	// 	- **ruleList**: the keyword of the check rule.
	//
	// example:
	//
	// [{\\"classKey\\": \\"password\\", \\"ruleList\\": [\\"huaweicloud_ak\\", \\"ak_leak\\"]}]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether to enable the new ruled for automatic check only on agentless detection. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 1
	EnableNewRule *int32 `json:"EnableNewRule,omitempty" xml:"EnableNewRule,omitempty"`
	// The source of the check rules. Valid values:
	//
	// 	- **image**: image.
	//
	// 	- **agentless**: agentless detection.
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SetSensitiveDefineRuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSensitiveDefineRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetSensitiveDefineRuleConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *SetSensitiveDefineRuleConfigRequest) GetEnableNewRule() *int32 {
	return s.EnableNewRule
}

func (s *SetSensitiveDefineRuleConfigRequest) GetSource() *string {
	return s.Source
}

func (s *SetSensitiveDefineRuleConfigRequest) SetConfig(v string) *SetSensitiveDefineRuleConfigRequest {
	s.Config = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigRequest) SetEnableNewRule(v int32) *SetSensitiveDefineRuleConfigRequest {
	s.EnableNewRule = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigRequest) SetSource(v string) *SetSensitiveDefineRuleConfigRequest {
	s.Source = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigRequest) Validate() error {
	return dara.Validate(s)
}
