// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEgressProxyConfigurationRequest
	GetClientToken() *string
	SetDefaultRulePhaseActions(v *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) *CreateEgressProxyConfigurationRequest
	GetDefaultRulePhaseActions() *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions
	SetDescription(v string) *CreateEgressProxyConfigurationRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateEgressProxyConfigurationRequest
	GetDryRun() *bool
	SetEgressProxyConfigurationName(v string) *CreateEgressProxyConfigurationRequest
	GetEgressProxyConfigurationName() *string
	SetResourceGroupId(v string) *CreateEgressProxyConfigurationRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateEgressProxyConfigurationRequestTag) *CreateEgressProxyConfigurationRequest
	GetTag() []*CreateEgressProxyConfigurationRequestTag
}

type CreateEgressProxyConfigurationRequest struct {
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
	// The default actions to be performed in each effective phase (PreDns, PreRequest, and PostResponse) when traffic does not match any rules.
	DefaultRulePhaseActions *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions `json:"DefaultRulePhaseActions,omitempty" xml:"DefaultRulePhaseActions,omitempty" type:"Struct"`
	// The description of the egress proxy configuration. The description must be 2 to 256 characters in length, and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run of the request. Valid values:
	//
	// - **true**: performs a dry run and does not create the egress proxy configuration. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, the corresponding error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the egress proxy configuration. The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyConfigurationName *string `json:"EgressProxyConfigurationName,omitempty" xml:"EgressProxyConfigurationName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4jcyx5e****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tag []*CreateEgressProxyConfigurationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEgressProxyConfigurationRequest) GetDefaultRulePhaseActions() *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions {
	return s.DefaultRulePhaseActions
}

func (s *CreateEgressProxyConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyConfigurationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEgressProxyConfigurationRequest) GetEgressProxyConfigurationName() *string {
	return s.EgressProxyConfigurationName
}

func (s *CreateEgressProxyConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEgressProxyConfigurationRequest) GetTag() []*CreateEgressProxyConfigurationRequestTag {
	return s.Tag
}

func (s *CreateEgressProxyConfigurationRequest) SetClientToken(v string) *CreateEgressProxyConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) SetDefaultRulePhaseActions(v *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) *CreateEgressProxyConfigurationRequest {
	s.DefaultRulePhaseActions = v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) SetDescription(v string) *CreateEgressProxyConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) SetDryRun(v bool) *CreateEgressProxyConfigurationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) SetEgressProxyConfigurationName(v string) *CreateEgressProxyConfigurationRequest {
	s.EgressProxyConfigurationName = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) SetResourceGroupId(v string) *CreateEgressProxyConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) SetTag(v []*CreateEgressProxyConfigurationRequestTag) *CreateEgressProxyConfigurationRequest {
	s.Tag = v
	return s
}

func (s *CreateEgressProxyConfigurationRequest) Validate() error {
	if s.DefaultRulePhaseActions != nil {
		if err := s.DefaultRulePhaseActions.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyConfigurationRequestDefaultRulePhaseActions struct {
	// The default action to be performed in the PostResponse phase (after the response is returned) when traffic does not match any rules. This phase supports only the **Allow*	- value (default), which allows the traffic to pass through.
	//
	// example:
	//
	// Allow
	PostResponse *string `json:"PostResponse,omitempty" xml:"PostResponse,omitempty"`
	// The default action to be performed in the PreDns phase (before DNS resolution) when traffic does not match any rules. Valid values:
	//
	// - **Allow*	- (default): allows the traffic to pass through.
	//
	// - **Deny**: denies the traffic.
	//
	// example:
	//
	// Allow
	PreDns *string `json:"PreDns,omitempty" xml:"PreDns,omitempty"`
	// The default action to be performed in the PreRequest phase (before the request is forwarded) when traffic does not match any rules. Valid values:
	//
	// - **Allow*	- (default): allows the traffic to pass through.
	//
	// - **Deny**: denies the traffic.
	//
	// example:
	//
	// Allow
	PreRequest *string `json:"PreRequest,omitempty" xml:"PreRequest,omitempty"`
}

func (s CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) GetPostResponse() *string {
	return s.PostResponse
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) GetPreDns() *string {
	return s.PreDns
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) GetPreRequest() *string {
	return s.PreRequest
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) SetPostResponse(v string) *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions {
	s.PostResponse = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) SetPreDns(v string) *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions {
	s.PreDns = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) SetPreRequest(v string) *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions {
	s.PreRequest = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequestDefaultRulePhaseActions) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyConfigurationRequestTag struct {
	// The tag key. The key can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 256 characters in length, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEgressProxyConfigurationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyConfigurationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyConfigurationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEgressProxyConfigurationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEgressProxyConfigurationRequestTag) SetKey(v string) *CreateEgressProxyConfigurationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequestTag) SetValue(v string) *CreateEgressProxyConfigurationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEgressProxyConfigurationRequestTag) Validate() error {
	return dara.Validate(s)
}
