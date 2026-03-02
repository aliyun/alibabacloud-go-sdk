// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLaneCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpLaneCreateCmd
	GetAlias() *string
	SetCompanyId(v int64) *PdpLaneCreateCmd
	GetCompanyId() *int64
	SetCreatorUid(v string) *PdpLaneCreateCmd
	GetCreatorUid() *string
	SetCustomeMarkingRules(v string) *PdpLaneCreateCmd
	GetCustomeMarkingRules() *string
	SetDescription(v string) *PdpLaneCreateCmd
	GetDescription() *string
	SetEnv(v string) *PdpLaneCreateCmd
	GetEnv() *string
	SetInitGroupFlag(v bool) *PdpLaneCreateCmd
	GetInitGroupFlag() *bool
	SetInletServiceIds(v string) *PdpLaneCreateCmd
	GetInletServiceIds() *string
	SetMarkingRuleType(v string) *PdpLaneCreateCmd
	GetMarkingRuleType() *string
	SetName(v string) *PdpLaneCreateCmd
	GetName() *string
	SetProductId(v int64) *PdpLaneCreateCmd
	GetProductId() *int64
	SetRuleConditions(v []*RuleCondition) *PdpLaneCreateCmd
	GetRuleConditions() []*RuleCondition
	SetRuleConnectType(v string) *PdpLaneCreateCmd
	GetRuleConnectType() *string
	SetServiceGroupIds(v string) *PdpLaneCreateCmd
	GetServiceGroupIds() *string
	SetType(v string) *PdpLaneCreateCmd
	GetType() *string
}

type PdpLaneCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 灰度
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 252333049301505383
	CreatorUid *string `json:"creatorUid,omitempty" xml:"creatorUid,omitempty"`
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
	// 泳道描述信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// true
	InitGroupFlag *bool `json:"initGroupFlag,omitempty" xml:"initGroupFlag,omitempty"`
	// example:
	//
	// 1,2
	InletServiceIds *string `json:"inletServiceIds,omitempty" xml:"inletServiceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEMPLATE
	MarkingRuleType *string `json:"markingRuleType,omitempty" xml:"markingRuleType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gray
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProductId      *int64           `json:"productId,omitempty" xml:"productId,omitempty"`
	RuleConditions []*RuleCondition `json:"ruleConditions,omitempty" xml:"ruleConditions,omitempty" type:"Repeated"`
	// example:
	//
	// and
	RuleConnectType *string `json:"ruleConnectType,omitempty" xml:"ruleConnectType,omitempty"`
	// example:
	//
	// 1,2
	ServiceGroupIds *string `json:"serviceGroupIds,omitempty" xml:"serviceGroupIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 灰度
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpLaneCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PdpLaneCreateCmd) GoString() string {
	return s.String()
}

func (s *PdpLaneCreateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PdpLaneCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpLaneCreateCmd) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *PdpLaneCreateCmd) GetCustomeMarkingRules() *string {
	return s.CustomeMarkingRules
}

func (s *PdpLaneCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *PdpLaneCreateCmd) GetEnv() *string {
	return s.Env
}

func (s *PdpLaneCreateCmd) GetInitGroupFlag() *bool {
	return s.InitGroupFlag
}

func (s *PdpLaneCreateCmd) GetInletServiceIds() *string {
	return s.InletServiceIds
}

func (s *PdpLaneCreateCmd) GetMarkingRuleType() *string {
	return s.MarkingRuleType
}

func (s *PdpLaneCreateCmd) GetName() *string {
	return s.Name
}

func (s *PdpLaneCreateCmd) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpLaneCreateCmd) GetRuleConditions() []*RuleCondition {
	return s.RuleConditions
}

func (s *PdpLaneCreateCmd) GetRuleConnectType() *string {
	return s.RuleConnectType
}

func (s *PdpLaneCreateCmd) GetServiceGroupIds() *string {
	return s.ServiceGroupIds
}

func (s *PdpLaneCreateCmd) GetType() *string {
	return s.Type
}

func (s *PdpLaneCreateCmd) SetAlias(v string) *PdpLaneCreateCmd {
	s.Alias = &v
	return s
}

func (s *PdpLaneCreateCmd) SetCompanyId(v int64) *PdpLaneCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PdpLaneCreateCmd) SetCreatorUid(v string) *PdpLaneCreateCmd {
	s.CreatorUid = &v
	return s
}

func (s *PdpLaneCreateCmd) SetCustomeMarkingRules(v string) *PdpLaneCreateCmd {
	s.CustomeMarkingRules = &v
	return s
}

func (s *PdpLaneCreateCmd) SetDescription(v string) *PdpLaneCreateCmd {
	s.Description = &v
	return s
}

func (s *PdpLaneCreateCmd) SetEnv(v string) *PdpLaneCreateCmd {
	s.Env = &v
	return s
}

func (s *PdpLaneCreateCmd) SetInitGroupFlag(v bool) *PdpLaneCreateCmd {
	s.InitGroupFlag = &v
	return s
}

func (s *PdpLaneCreateCmd) SetInletServiceIds(v string) *PdpLaneCreateCmd {
	s.InletServiceIds = &v
	return s
}

func (s *PdpLaneCreateCmd) SetMarkingRuleType(v string) *PdpLaneCreateCmd {
	s.MarkingRuleType = &v
	return s
}

func (s *PdpLaneCreateCmd) SetName(v string) *PdpLaneCreateCmd {
	s.Name = &v
	return s
}

func (s *PdpLaneCreateCmd) SetProductId(v int64) *PdpLaneCreateCmd {
	s.ProductId = &v
	return s
}

func (s *PdpLaneCreateCmd) SetRuleConditions(v []*RuleCondition) *PdpLaneCreateCmd {
	s.RuleConditions = v
	return s
}

func (s *PdpLaneCreateCmd) SetRuleConnectType(v string) *PdpLaneCreateCmd {
	s.RuleConnectType = &v
	return s
}

func (s *PdpLaneCreateCmd) SetServiceGroupIds(v string) *PdpLaneCreateCmd {
	s.ServiceGroupIds = &v
	return s
}

func (s *PdpLaneCreateCmd) SetType(v string) *PdpLaneCreateCmd {
	s.Type = &v
	return s
}

func (s *PdpLaneCreateCmd) Validate() error {
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
