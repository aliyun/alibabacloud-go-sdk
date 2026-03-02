// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLaneUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpLaneUpdateCmd
	GetAlias() *string
	SetCustomeMarkingRules(v string) *PdpLaneUpdateCmd
	GetCustomeMarkingRules() *string
	SetDescription(v string) *PdpLaneUpdateCmd
	GetDescription() *string
	SetId(v int64) *PdpLaneUpdateCmd
	GetId() *int64
	SetInitGroupFlag(v bool) *PdpLaneUpdateCmd
	GetInitGroupFlag() *bool
	SetInletServiceIds(v string) *PdpLaneUpdateCmd
	GetInletServiceIds() *string
	SetMarkingRuleType(v string) *PdpLaneUpdateCmd
	GetMarkingRuleType() *string
	SetRuleConditions(v []*RuleCondition) *PdpLaneUpdateCmd
	GetRuleConditions() []*RuleCondition
	SetRuleConnectType(v string) *PdpLaneUpdateCmd
	GetRuleConnectType() *string
	SetServiceGroupIds(v string) *PdpLaneUpdateCmd
	GetServiceGroupIds() *string
}

type PdpLaneUpdateCmd struct {
	// example:
	//
	// 灰度
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// rules:
	//
	//     - vars:
	//
	//         - name: userId
	//
	//           position: header
	//
	//           key: x-linkedmall-user-id
	//
	//         - name: userType
	//
	//           position: query
	//
	//           key: userType
	//
	//       expression: "userId % 10 == 0 && userType == A"
	//
	//       tag: "gray"
	//
	//       scope: "pbc1,pbc2"
	CustomeMarkingRules *string `json:"customeMarkingRules,omitempty" xml:"customeMarkingRules,omitempty"`
	// example:
	//
	// 灰度描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id            *int64 `json:"id,omitempty" xml:"id,omitempty"`
	InitGroupFlag *bool  `json:"initGroupFlag,omitempty" xml:"initGroupFlag,omitempty"`
	// example:
	//
	// 1,2
	InletServiceIds *string `json:"inletServiceIds,omitempty" xml:"inletServiceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEMPLATE
	MarkingRuleType *string          `json:"markingRuleType,omitempty" xml:"markingRuleType,omitempty"`
	RuleConditions  []*RuleCondition `json:"ruleConditions,omitempty" xml:"ruleConditions,omitempty" type:"Repeated"`
	// example:
	//
	// and
	RuleConnectType *string `json:"ruleConnectType,omitempty" xml:"ruleConnectType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ServiceGroupIds *string `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty"`
}

func (s PdpLaneUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpLaneUpdateCmd) GoString() string {
	return s.String()
}

func (s *PdpLaneUpdateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PdpLaneUpdateCmd) GetCustomeMarkingRules() *string {
	return s.CustomeMarkingRules
}

func (s *PdpLaneUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *PdpLaneUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *PdpLaneUpdateCmd) GetInitGroupFlag() *bool {
	return s.InitGroupFlag
}

func (s *PdpLaneUpdateCmd) GetInletServiceIds() *string {
	return s.InletServiceIds
}

func (s *PdpLaneUpdateCmd) GetMarkingRuleType() *string {
	return s.MarkingRuleType
}

func (s *PdpLaneUpdateCmd) GetRuleConditions() []*RuleCondition {
	return s.RuleConditions
}

func (s *PdpLaneUpdateCmd) GetRuleConnectType() *string {
	return s.RuleConnectType
}

func (s *PdpLaneUpdateCmd) GetServiceGroupIds() *string {
	return s.ServiceGroupIds
}

func (s *PdpLaneUpdateCmd) SetAlias(v string) *PdpLaneUpdateCmd {
	s.Alias = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetCustomeMarkingRules(v string) *PdpLaneUpdateCmd {
	s.CustomeMarkingRules = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetDescription(v string) *PdpLaneUpdateCmd {
	s.Description = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetId(v int64) *PdpLaneUpdateCmd {
	s.Id = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetInitGroupFlag(v bool) *PdpLaneUpdateCmd {
	s.InitGroupFlag = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetInletServiceIds(v string) *PdpLaneUpdateCmd {
	s.InletServiceIds = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetMarkingRuleType(v string) *PdpLaneUpdateCmd {
	s.MarkingRuleType = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetRuleConditions(v []*RuleCondition) *PdpLaneUpdateCmd {
	s.RuleConditions = v
	return s
}

func (s *PdpLaneUpdateCmd) SetRuleConnectType(v string) *PdpLaneUpdateCmd {
	s.RuleConnectType = &v
	return s
}

func (s *PdpLaneUpdateCmd) SetServiceGroupIds(v string) *PdpLaneUpdateCmd {
	s.ServiceGroupIds = &v
	return s
}

func (s *PdpLaneUpdateCmd) Validate() error {
	if s.RuleConditions != nil {
		for _, item := range s.RuleConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
