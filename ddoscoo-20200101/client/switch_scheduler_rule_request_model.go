// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSchedulerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleName(v string) *SwitchSchedulerRuleRequest
	GetRuleName() *string
	SetRuleType(v int32) *SwitchSchedulerRuleRequest
	GetRuleType() *int32
	SetSwitchData(v string) *SwitchSchedulerRuleRequest
	GetSwitchData() *string
}

type SwitchSchedulerRuleRequest struct {
	// The name of the scheduling rule to manage.
	//
	// > You can call the [DescribeSchedulerRules](https://help.aliyun.com/document_detail/157481.html) operation to query the names of all scheduling rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the scheduling rule. Valid values:
	//
	// 	- **2**: tiered protection rule
	//
	// 	- **3**: network acceleration rule
	//
	// 	- **5**: Alibaba Cloud CDN (CDN) interaction rule
	//
	// 	- **6**: cloud service interaction rule
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The configuration that is used to switch service traffic. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that includes the following parameters:
	//
	// 	- **Value**: required. The IP address of the associated resource. Data type: string.
	//
	// 	- **State**: required. The operation type. Data type: integer. Valid values:
	//
	//     	- **0**: switches service traffic from the associated resource to your Anti-DDoS Pro or Anti-DDoS Premium instance for scrubbing.
	//
	//     	- **1**: switches service traffic back to the associated cloud resource.
	//
	// 	- **Interval**: optional. The waiting time that is required before the service traffic is switched back. Unit: minutes. Data type: integer. Usage notes:
	//
	//     	- If the **State*	- parameter is set to **0**, you must set this parameter to \\*\\*-1\\*\\*. Otherwise, the call fails.
	//
	//     	- If the **State*	- parameter is set to **1**, you do not need to set this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Value":"39.104.XX.XX","State":0,"Interval":-1}]
	SwitchData *string `json:"SwitchData,omitempty" xml:"SwitchData,omitempty"`
}

func (s SwitchSchedulerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchSchedulerRuleRequest) GoString() string {
	return s.String()
}

func (s *SwitchSchedulerRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *SwitchSchedulerRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *SwitchSchedulerRuleRequest) GetSwitchData() *string {
	return s.SwitchData
}

func (s *SwitchSchedulerRuleRequest) SetRuleName(v string) *SwitchSchedulerRuleRequest {
	s.RuleName = &v
	return s
}

func (s *SwitchSchedulerRuleRequest) SetRuleType(v int32) *SwitchSchedulerRuleRequest {
	s.RuleType = &v
	return s
}

func (s *SwitchSchedulerRuleRequest) SetSwitchData(v string) *SwitchSchedulerRuleRequest {
	s.SwitchData = &v
	return s
}

func (s *SwitchSchedulerRuleRequest) Validate() error {
	return dara.Validate(s)
}
