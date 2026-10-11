// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyConfigurationAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEgressProxyConfigurationAttributeRequest
	GetClientToken() *string
	SetDefaultRulePhaseActions(v *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) *UpdateEgressProxyConfigurationAttributeRequest
	GetDefaultRulePhaseActions() *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions
	SetDescription(v string) *UpdateEgressProxyConfigurationAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateEgressProxyConfigurationAttributeRequest
	GetDryRun() *bool
	SetEgressProxyConfigurationId(v string) *UpdateEgressProxyConfigurationAttributeRequest
	GetEgressProxyConfigurationId() *string
	SetEgressProxyConfigurationName(v string) *UpdateEgressProxyConfigurationAttributeRequest
	GetEgressProxyConfigurationName() *string
	SetEgressProxyRuleGroups(v []*UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) *UpdateEgressProxyConfigurationAttributeRequest
	GetEgressProxyRuleGroups() []*UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups
}

type UpdateEgressProxyConfigurationAttributeRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken             *string                                                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DefaultRulePhaseActions *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions `json:"DefaultRulePhaseActions,omitempty" xml:"DefaultRulePhaseActions,omitempty" type:"Struct"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// epc-gw80d16a969933e7****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyConfigurationName *string                                                                `json:"EgressProxyConfigurationName,omitempty" xml:"EgressProxyConfigurationName,omitempty"`
	EgressProxyRuleGroups        []*UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups `json:"EgressProxyRuleGroups,omitempty" xml:"EgressProxyRuleGroups,omitempty" type:"Repeated"`
}

func (s UpdateEgressProxyConfigurationAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyConfigurationAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetDefaultRulePhaseActions() *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions {
	return s.DefaultRulePhaseActions
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetEgressProxyConfigurationName() *string {
	return s.EgressProxyConfigurationName
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) GetEgressProxyRuleGroups() []*UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups {
	return s.EgressProxyRuleGroups
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetClientToken(v string) *UpdateEgressProxyConfigurationAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetDefaultRulePhaseActions(v *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) *UpdateEgressProxyConfigurationAttributeRequest {
	s.DefaultRulePhaseActions = v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetDescription(v string) *UpdateEgressProxyConfigurationAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetDryRun(v bool) *UpdateEgressProxyConfigurationAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetEgressProxyConfigurationId(v string) *UpdateEgressProxyConfigurationAttributeRequest {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetEgressProxyConfigurationName(v string) *UpdateEgressProxyConfigurationAttributeRequest {
	s.EgressProxyConfigurationName = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) SetEgressProxyRuleGroups(v []*UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) *UpdateEgressProxyConfigurationAttributeRequest {
	s.EgressProxyRuleGroups = v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequest) Validate() error {
	if s.DefaultRulePhaseActions != nil {
		if err := s.DefaultRulePhaseActions.Validate(); err != nil {
			return err
		}
	}
	if s.EgressProxyRuleGroups != nil {
		for _, item := range s.EgressProxyRuleGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions struct {
	// This parameter is required.
	//
	// example:
	//
	// Allow
	PostResponse *string `json:"PostResponse,omitempty" xml:"PostResponse,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Allow
	PreDns *string `json:"PreDns,omitempty" xml:"PreDns,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Allow
	PreRequest *string `json:"PreRequest,omitempty" xml:"PreRequest,omitempty"`
}

func (s UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) GetPostResponse() *string {
	return s.PostResponse
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) GetPreDns() *string {
	return s.PreDns
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) GetPreRequest() *string {
	return s.PreRequest
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) SetPostResponse(v string) *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions {
	s.PostResponse = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) SetPreDns(v string) *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions {
	s.PreDns = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) SetPreRequest(v string) *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions {
	s.PreRequest = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequestDefaultRulePhaseActions) Validate() error {
	return dara.Validate(s)
}

type UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups struct {
	// This parameter is required.
	//
	// example:
	//
	// eprg-gw80bb6a95380edb****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) SetEgressProxyRuleGroupId(v string) *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) SetPriority(v int32) *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups {
	s.Priority = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeRequestEgressProxyRuleGroups) Validate() error {
	return dara.Validate(s)
}
