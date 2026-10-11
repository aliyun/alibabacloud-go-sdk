// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyRulesAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEgressProxyRulesAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateEgressProxyRulesAttributeRequest
	GetDryRun() *bool
	SetEgressProxyRuleGroupId(v string) *UpdateEgressProxyRulesAttributeRequest
	GetEgressProxyRuleGroupId() *string
	SetEgressProxyRules(v []*UpdateEgressProxyRulesAttributeRequestEgressProxyRules) *UpdateEgressProxyRulesAttributeRequest
	GetEgressProxyRules() []*UpdateEgressProxyRulesAttributeRequestEgressProxyRules
}

type UpdateEgressProxyRulesAttributeRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eprg-gw80bb6a95380edb****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// This parameter is required.
	EgressProxyRules []*UpdateEgressProxyRulesAttributeRequestEgressProxyRules `json:"EgressProxyRules,omitempty" xml:"EgressProxyRules,omitempty" type:"Repeated"`
}

func (s UpdateEgressProxyRulesAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRulesAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRulesAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEgressProxyRulesAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEgressProxyRulesAttributeRequest) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *UpdateEgressProxyRulesAttributeRequest) GetEgressProxyRules() []*UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	return s.EgressProxyRules
}

func (s *UpdateEgressProxyRulesAttributeRequest) SetClientToken(v string) *UpdateEgressProxyRulesAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequest) SetDryRun(v bool) *UpdateEgressProxyRulesAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequest) SetEgressProxyRuleGroupId(v string) *UpdateEgressProxyRulesAttributeRequest {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequest) SetEgressProxyRules(v []*UpdateEgressProxyRulesAttributeRequestEgressProxyRules) *UpdateEgressProxyRulesAttributeRequest {
	s.EgressProxyRules = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequest) Validate() error {
	if s.EgressProxyRules != nil {
		for _, item := range s.EgressProxyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEgressProxyRulesAttributeRequestEgressProxyRules struct {
	// example:
	//
	// Allow
	Action        *string                                                                `json:"Action,omitempty" xml:"Action,omitempty"`
	AddConditions []*UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions `json:"AddConditions,omitempty" xml:"AddConditions,omitempty" type:"Repeated"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// epr-gw80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyRuleName *string `json:"EgressProxyRuleName,omitempty" xml:"EgressProxyRuleName,omitempty"`
	// example:
	//
	// 10
	Priority         *int32                                                                    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RemoveConditions []*UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions `json:"RemoveConditions,omitempty" xml:"RemoveConditions,omitempty" type:"Repeated"`
}

func (s UpdateEgressProxyRulesAttributeRequestEgressProxyRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetAction() *string {
	return s.Action
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetAddConditions() []*UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions {
	return s.AddConditions
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetDescription() *string {
	return s.Description
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) GetRemoveConditions() []*UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions {
	return s.RemoveConditions
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetAction(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.Action = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetAddConditions(v []*UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.AddConditions = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetDescription(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.Description = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetEgressProxyRuleId(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.EgressProxyRuleId = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetEgressProxyRuleName(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.EgressProxyRuleName = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetPriority(v int32) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.Priority = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) SetRemoveConditions(v []*UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) *UpdateEgressProxyRulesAttributeRequestEgressProxyRules {
	s.RemoveConditions = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRules) Validate() error {
	if s.AddConditions != nil {
		for _, item := range s.AddConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RemoveConditions != nil {
		for _, item := range s.RemoveConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions struct {
	// This parameter is required.
	//
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// StringEquals
	ConditionOperator *string `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	// This parameter is required.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) SetConditionKey(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions {
	s.ConditionKey = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) SetConditionOperator(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions {
	s.ConditionOperator = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) SetConditionValues(v []*string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions {
	s.ConditionValues = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesAddConditions) Validate() error {
	return dara.Validate(s)
}

type UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions struct {
	// This parameter is required.
	//
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// StringEquals
	ConditionOperator *string `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	// This parameter is required.
	ConditionValues []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) SetConditionKey(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions {
	s.ConditionKey = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) SetConditionOperator(v string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions {
	s.ConditionOperator = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) SetConditionValues(v []*string) *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions {
	s.ConditionValues = v
	return s
}

func (s *UpdateEgressProxyRulesAttributeRequestEgressProxyRulesRemoveConditions) Validate() error {
	return dara.Validate(s)
}
