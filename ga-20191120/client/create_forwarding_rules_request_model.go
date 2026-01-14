// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateForwardingRulesRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateForwardingRulesRequest
	GetClientToken() *string
	SetForwardingRules(v []*CreateForwardingRulesRequestForwardingRules) *CreateForwardingRulesRequest
	GetForwardingRules() []*CreateForwardingRulesRequestForwardingRules
	SetListenerId(v string) *CreateForwardingRulesRequest
	GetListenerId() *string
	SetRegionId(v string) *CreateForwardingRulesRequest
	GetRegionId() *string
}

type CreateForwardingRulesRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4q****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ForwardingRules []*CreateForwardingRulesRequestForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateForwardingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateForwardingRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateForwardingRulesRequest) GetForwardingRules() []*CreateForwardingRulesRequestForwardingRules {
	return s.ForwardingRules
}

func (s *CreateForwardingRulesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateForwardingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateForwardingRulesRequest) SetAcceleratorId(v string) *CreateForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetClientToken(v string) *CreateForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetForwardingRules(v []*CreateForwardingRulesRequestForwardingRules) *CreateForwardingRulesRequest {
	s.ForwardingRules = v
	return s
}

func (s *CreateForwardingRulesRequest) SetListenerId(v string) *CreateForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateForwardingRulesRequest) SetRegionId(v string) *CreateForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateForwardingRulesRequest) Validate() error {
	if s.ForwardingRules != nil {
		for _, item := range s.ForwardingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateForwardingRulesRequestForwardingRules struct {
	ForwardingRuleName *string `json:"ForwardingRuleName,omitempty" xml:"ForwardingRuleName,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	RuleActions []*CreateForwardingRulesRequestForwardingRulesRuleActions `json:"RuleActions,omitempty" xml:"RuleActions,omitempty" type:"Repeated"`
	// This parameter is required.
	RuleConditions []*CreateForwardingRulesRequestForwardingRulesRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	RuleDirection  *string                                                      `json:"RuleDirection,omitempty" xml:"RuleDirection,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRules) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRules) GetForwardingRuleName() *string {
	return s.ForwardingRuleName
}

func (s *CreateForwardingRulesRequestForwardingRules) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateForwardingRulesRequestForwardingRules) GetRuleActions() []*CreateForwardingRulesRequestForwardingRulesRuleActions {
	return s.RuleActions
}

func (s *CreateForwardingRulesRequestForwardingRules) GetRuleConditions() []*CreateForwardingRulesRequestForwardingRulesRuleConditions {
	return s.RuleConditions
}

func (s *CreateForwardingRulesRequestForwardingRules) GetRuleDirection() *string {
	return s.RuleDirection
}

func (s *CreateForwardingRulesRequestForwardingRules) SetForwardingRuleName(v string) *CreateForwardingRulesRequestForwardingRules {
	s.ForwardingRuleName = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetPriority(v int32) *CreateForwardingRulesRequestForwardingRules {
	s.Priority = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleActions(v []*CreateForwardingRulesRequestForwardingRulesRuleActions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleActions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleConditions(v []*CreateForwardingRulesRequestForwardingRulesRuleConditions) *CreateForwardingRulesRequestForwardingRules {
	s.RuleConditions = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) SetRuleDirection(v string) *CreateForwardingRulesRequestForwardingRules {
	s.RuleDirection = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRules) Validate() error {
	if s.RuleActions != nil {
		for _, item := range s.RuleActions {
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
	return nil
}

type CreateForwardingRulesRequestForwardingRulesRuleActions struct {
	ForwardGroupConfig *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig `json:"ForwardGroupConfig,omitempty" xml:"ForwardGroupConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// This parameter is required.
	RuleActionType  *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	RuleActionValue *string `json:"RuleActionValue,omitempty" xml:"RuleActionValue,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetForwardGroupConfig() *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	return s.ForwardGroupConfig
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetOrder() *int32 {
	return s.Order
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetRuleActionType() *string {
	return s.RuleActionType
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) GetRuleActionValue() *string {
	return s.RuleActionValue
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetForwardGroupConfig(v *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.ForwardGroupConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetOrder(v int32) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.Order = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) SetRuleActionValue(v string) *CreateForwardingRulesRequestForwardingRulesRuleActions {
	s.RuleActionValue = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActions) Validate() error {
	if s.ForwardGroupConfig != nil {
		if err := s.ForwardGroupConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig struct {
	// This parameter is required.
	ServerGroupTuples []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples `json:"ServerGroupTuples,omitempty" xml:"ServerGroupTuples,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) GetServerGroupTuples() []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	return s.ServerGroupTuples
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) SetServerGroupTuples(v []*CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig {
	s.ServerGroupTuples = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfig) Validate() error {
	if s.ServerGroupTuples != nil {
		for _, item := range s.ServerGroupTuples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples struct {
	// This parameter is required.
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) SetEndpointGroupId(v string) *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples {
	s.EndpointGroupId = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleActionsForwardGroupConfigServerGroupTuples) Validate() error {
	return dara.Validate(s)
}

type CreateForwardingRulesRequestForwardingRulesRuleConditions struct {
	HostConfig         *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig `json:"HostConfig,omitempty" xml:"HostConfig,omitempty" type:"Struct"`
	PathConfig         *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig `json:"PathConfig,omitempty" xml:"PathConfig,omitempty" type:"Struct"`
	RuleConditionType  *string                                                              `json:"RuleConditionType,omitempty" xml:"RuleConditionType,omitempty"`
	RuleConditionValue *string                                                              `json:"RuleConditionValue,omitempty" xml:"RuleConditionValue,omitempty"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetHostConfig() *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	return s.HostConfig
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetPathConfig() *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	return s.PathConfig
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetRuleConditionType() *string {
	return s.RuleConditionType
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) GetRuleConditionValue() *string {
	return s.RuleConditionValue
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetHostConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.HostConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetPathConfig(v *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.PathConfig = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionType(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionType = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) SetRuleConditionValue(v string) *CreateForwardingRulesRequestForwardingRulesRuleConditions {
	s.RuleConditionValue = &v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditions) Validate() error {
	if s.HostConfig != nil {
		if err := s.HostConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PathConfig != nil {
		if err := s.PathConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) SetValues(v []*string) *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig {
	s.Values = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsHostConfig) Validate() error {
	return dara.Validate(s)
}

type CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) GetValues() []*string {
	return s.Values
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) SetValues(v []*string) *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig {
	s.Values = v
	return s
}

func (s *CreateForwardingRulesRequestForwardingRulesRuleConditionsPathConfig) Validate() error {
	return dara.Validate(s)
}
