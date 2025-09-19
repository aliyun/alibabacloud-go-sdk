// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBuildRiskDefineRuleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SetBuildRiskDefineRuleConfigRequest
	GetConfig() *string
}

type SetBuildRiskDefineRuleConfigRequest struct {
	// The configuration item for scanning image build command risks. Valid values:
	//
	// 	- **classKey**: Set the value to a valid value of the ClassKey parameter in RuleTree.
	//
	// 	- **ruleList**: Set the value to a valid value of the RuleKey parameter in RuleList.
	//
	// >  You can call the [GetBuildRiskDefineRuleConfig](~~GetBuildRiskDefineRuleConfig~~) operation to query the valid values.
	//
	// example:
	//
	// [
	//
	// 	{
	//
	// 		"classKey": "other",
	//
	// 		"ruleList": [
	//
	// 			"add",
	//
	// 			"apk"
	//
	// 		]
	//
	// 	}
	//
	// ]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s SetBuildRiskDefineRuleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetBuildRiskDefineRuleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetBuildRiskDefineRuleConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *SetBuildRiskDefineRuleConfigRequest) SetConfig(v string) *SetBuildRiskDefineRuleConfigRequest {
	s.Config = &v
	return s
}

func (s *SetBuildRiskDefineRuleConfigRequest) Validate() error {
	return dara.Validate(s)
}
