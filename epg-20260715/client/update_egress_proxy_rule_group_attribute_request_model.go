// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyRuleGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEgressProxyRuleGroupAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateEgressProxyRuleGroupAttributeRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateEgressProxyRuleGroupAttributeRequest
	GetDryRun() *bool
	SetEgressProxyIpPoolId(v string) *UpdateEgressProxyRuleGroupAttributeRequest
	GetEgressProxyIpPoolId() *string
	SetEgressProxyRuleGroupId(v string) *UpdateEgressProxyRuleGroupAttributeRequest
	GetEgressProxyRuleGroupId() *string
	SetEgressProxyRuleGroupName(v string) *UpdateEgressProxyRuleGroupAttributeRequest
	GetEgressProxyRuleGroupName() *string
}

type UpdateEgressProxyRuleGroupAttributeRequest struct {
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyRuleGroupName *string `json:"EgressProxyRuleGroupName,omitempty" xml:"EgressProxyRuleGroupName,omitempty"`
}

func (s UpdateEgressProxyRuleGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRuleGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) GetEgressProxyRuleGroupName() *string {
	return s.EgressProxyRuleGroupName
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) SetClientToken(v string) *UpdateEgressProxyRuleGroupAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) SetDescription(v string) *UpdateEgressProxyRuleGroupAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) SetDryRun(v bool) *UpdateEgressProxyRuleGroupAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) SetEgressProxyIpPoolId(v string) *UpdateEgressProxyRuleGroupAttributeRequest {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) SetEgressProxyRuleGroupId(v string) *UpdateEgressProxyRuleGroupAttributeRequest {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) SetEgressProxyRuleGroupName(v string) *UpdateEgressProxyRuleGroupAttributeRequest {
	s.EgressProxyRuleGroupName = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
