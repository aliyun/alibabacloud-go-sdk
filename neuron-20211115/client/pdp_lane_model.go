// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpLane interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PdpLane
	GetAlias() *string
	SetCompanyId(v int64) *PdpLane
	GetCompanyId() *int64
	SetCreatorName(v string) *PdpLane
	GetCreatorName() *string
	SetCreatorUid(v string) *PdpLane
	GetCreatorUid() *string
	SetCustomeMarkingRules(v string) *PdpLane
	GetCustomeMarkingRules() *string
	SetDescription(v string) *PdpLane
	GetDescription() *string
	SetEnv(v string) *PdpLane
	GetEnv() *string
	SetId(v int64) *PdpLane
	GetId() *int64
	SetInitGroupFlag(v bool) *PdpLane
	GetInitGroupFlag() *bool
	SetInletServices(v []*PdpServiceInfo) *PdpLane
	GetInletServices() []*PdpServiceInfo
	SetMarkingRuleType(v string) *PdpLane
	GetMarkingRuleType() *string
	SetName(v string) *PdpLane
	GetName() *string
	SetProductId(v int64) *PdpLane
	GetProductId() *int64
	SetProductName(v string) *PdpLane
	GetProductName() *string
	SetRequestId(v string) *PdpLane
	GetRequestId() *string
	SetRuleConditions(v []*RuleCondition) *PdpLane
	GetRuleConditions() []*RuleCondition
	SetRuleConnectType(v string) *PdpLane
	GetRuleConnectType() *string
	SetServiceGroups(v []*ServiceGroupInfo) *PdpLane
	GetServiceGroups() []*ServiceGroupInfo
	SetType(v string) *PdpLane
	GetType() *string
}

type PdpLane struct {
	// example:
	//
	// 灰度
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 张三
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// 252333049301505383
	CreatorUid *string `json:"creatorUid,omitempty" xml:"creatorUid,omitempty"`
	// This parameter is required.
	//
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
	Description         *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	InitGroupFlag *bool             `json:"initGroupFlag,omitempty" xml:"initGroupFlag,omitempty"`
	InletServices []*PdpServiceInfo `json:"inletServices,omitempty" xml:"inletServices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// TEMPLATE
	MarkingRuleType *string `json:"markingRuleType,omitempty" xml:"markingRuleType,omitempty"`
	// example:
	//
	// gray
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// linkemall
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId      *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RuleConditions []*RuleCondition `json:"ruleConditions,omitempty" xml:"ruleConditions,omitempty" type:"Repeated"`
	// example:
	//
	// and
	RuleConnectType *string             `json:"ruleConnectType,omitempty" xml:"ruleConnectType,omitempty"`
	ServiceGroups   []*ServiceGroupInfo `json:"serviceGroups,omitempty" xml:"serviceGroups,omitempty" type:"Repeated"`
	// example:
	//
	// SYSTEM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PdpLane) String() string {
	return dara.Prettify(s)
}

func (s PdpLane) GoString() string {
	return s.String()
}

func (s *PdpLane) GetAlias() *string {
	return s.Alias
}

func (s *PdpLane) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PdpLane) GetCreatorName() *string {
	return s.CreatorName
}

func (s *PdpLane) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *PdpLane) GetCustomeMarkingRules() *string {
	return s.CustomeMarkingRules
}

func (s *PdpLane) GetDescription() *string {
	return s.Description
}

func (s *PdpLane) GetEnv() *string {
	return s.Env
}

func (s *PdpLane) GetId() *int64 {
	return s.Id
}

func (s *PdpLane) GetInitGroupFlag() *bool {
	return s.InitGroupFlag
}

func (s *PdpLane) GetInletServices() []*PdpServiceInfo {
	return s.InletServices
}

func (s *PdpLane) GetMarkingRuleType() *string {
	return s.MarkingRuleType
}

func (s *PdpLane) GetName() *string {
	return s.Name
}

func (s *PdpLane) GetProductId() *int64 {
	return s.ProductId
}

func (s *PdpLane) GetProductName() *string {
	return s.ProductName
}

func (s *PdpLane) GetRequestId() *string {
	return s.RequestId
}

func (s *PdpLane) GetRuleConditions() []*RuleCondition {
	return s.RuleConditions
}

func (s *PdpLane) GetRuleConnectType() *string {
	return s.RuleConnectType
}

func (s *PdpLane) GetServiceGroups() []*ServiceGroupInfo {
	return s.ServiceGroups
}

func (s *PdpLane) GetType() *string {
	return s.Type
}

func (s *PdpLane) SetAlias(v string) *PdpLane {
	s.Alias = &v
	return s
}

func (s *PdpLane) SetCompanyId(v int64) *PdpLane {
	s.CompanyId = &v
	return s
}

func (s *PdpLane) SetCreatorName(v string) *PdpLane {
	s.CreatorName = &v
	return s
}

func (s *PdpLane) SetCreatorUid(v string) *PdpLane {
	s.CreatorUid = &v
	return s
}

func (s *PdpLane) SetCustomeMarkingRules(v string) *PdpLane {
	s.CustomeMarkingRules = &v
	return s
}

func (s *PdpLane) SetDescription(v string) *PdpLane {
	s.Description = &v
	return s
}

func (s *PdpLane) SetEnv(v string) *PdpLane {
	s.Env = &v
	return s
}

func (s *PdpLane) SetId(v int64) *PdpLane {
	s.Id = &v
	return s
}

func (s *PdpLane) SetInitGroupFlag(v bool) *PdpLane {
	s.InitGroupFlag = &v
	return s
}

func (s *PdpLane) SetInletServices(v []*PdpServiceInfo) *PdpLane {
	s.InletServices = v
	return s
}

func (s *PdpLane) SetMarkingRuleType(v string) *PdpLane {
	s.MarkingRuleType = &v
	return s
}

func (s *PdpLane) SetName(v string) *PdpLane {
	s.Name = &v
	return s
}

func (s *PdpLane) SetProductId(v int64) *PdpLane {
	s.ProductId = &v
	return s
}

func (s *PdpLane) SetProductName(v string) *PdpLane {
	s.ProductName = &v
	return s
}

func (s *PdpLane) SetRequestId(v string) *PdpLane {
	s.RequestId = &v
	return s
}

func (s *PdpLane) SetRuleConditions(v []*RuleCondition) *PdpLane {
	s.RuleConditions = v
	return s
}

func (s *PdpLane) SetRuleConnectType(v string) *PdpLane {
	s.RuleConnectType = &v
	return s
}

func (s *PdpLane) SetServiceGroups(v []*ServiceGroupInfo) *PdpLane {
	s.ServiceGroups = v
	return s
}

func (s *PdpLane) SetType(v string) *PdpLane {
	s.Type = &v
	return s
}

func (s *PdpLane) Validate() error {
	if s.InletServices != nil {
		for _, item := range s.InletServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleConditions != nil {
		for _, item := range s.RuleConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceGroups != nil {
		for _, item := range s.ServiceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
