// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEgressProxyRuleGroupsToConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachEgressProxyRuleGroupsToConfigurationRequest
	GetClientToken() *string
	SetDryRun(v bool) *AttachEgressProxyRuleGroupsToConfigurationRequest
	GetDryRun() *bool
	SetEgressProxyConfigurationId(v string) *AttachEgressProxyRuleGroupsToConfigurationRequest
	GetEgressProxyConfigurationId() *string
	SetEgressProxyRuleGroups(v []*AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) *AttachEgressProxyRuleGroupsToConfigurationRequest
	GetEgressProxyRuleGroups() []*AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups
}

type AttachEgressProxyRuleGroupsToConfigurationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > Note If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - **true**: performs a dry run and does not attach the egress proxy rule groups. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the system returns an HTTP 2xx status code and performs the operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// epc-uf609a6a8fdf9f4d****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The list of egress proxy rule group references to be attached to the egress proxy configuration.
	//
	// This parameter is required.
	EgressProxyRuleGroups []*AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups `json:"EgressProxyRuleGroups,omitempty" xml:"EgressProxyRuleGroups,omitempty" type:"Repeated"`
}

func (s AttachEgressProxyRuleGroupsToConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachEgressProxyRuleGroupsToConfigurationRequest) GoString() string {
	return s.String()
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) GetEgressProxyRuleGroups() []*AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups {
	return s.EgressProxyRuleGroups
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) SetClientToken(v string) *AttachEgressProxyRuleGroupsToConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) SetDryRun(v bool) *AttachEgressProxyRuleGroupsToConfigurationRequest {
	s.DryRun = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) SetEgressProxyConfigurationId(v string) *AttachEgressProxyRuleGroupsToConfigurationRequest {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) SetEgressProxyRuleGroups(v []*AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) *AttachEgressProxyRuleGroupsToConfigurationRequest {
	s.EgressProxyRuleGroups = v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequest) Validate() error {
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

type AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups struct {
	// The ID of the egress proxy rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// The priority of the egress proxy rule group in the egress proxy configuration. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) String() string {
	return dara.Prettify(s)
}

func (s AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) GoString() string {
	return s.String()
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) GetPriority() *int32 {
	return s.Priority
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) SetEgressProxyRuleGroupId(v string) *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) SetPriority(v int32) *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups {
	s.Priority = &v
	return s
}

func (s *AttachEgressProxyRuleGroupsToConfigurationRequestEgressProxyRuleGroups) Validate() error {
	return dara.Validate(s)
}
