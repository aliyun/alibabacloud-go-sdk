// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyConfigurationAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetCreateTime() *string
	SetDefaultRulePhaseActions(v *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) *GetEgressProxyConfigurationAttributeResponseBody
	GetDefaultRulePhaseActions() *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions
	SetDescription(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetDescription() *string
	SetEgressProxyConfigurationId(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetEgressProxyConfigurationId() *string
	SetEgressProxyConfigurationName(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetEgressProxyConfigurationName() *string
	SetRequestId(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetResourceGroupId() *string
	SetRuleGroups(v []*GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) *GetEgressProxyConfigurationAttributeResponseBody
	GetRuleGroups() []*GetEgressProxyConfigurationAttributeResponseBodyRuleGroups
	SetStatus(v string) *GetEgressProxyConfigurationAttributeResponseBody
	GetStatus() *string
	SetTags(v []*GetEgressProxyConfigurationAttributeResponseBodyTags) *GetEgressProxyConfigurationAttributeResponseBody
	GetTags() []*GetEgressProxyConfigurationAttributeResponseBodyTags
}

type GetEgressProxyConfigurationAttributeResponseBody struct {
	// The creation time of the egress proxy configuration. The time is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2026-02-02T02:20:24Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default actions executed in each effective phase (PreDns, PreRequest, and PostResponse) when traffic does not match any rules.
	DefaultRulePhaseActions *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions `json:"DefaultRulePhaseActions,omitempty" xml:"DefaultRulePhaseActions,omitempty" type:"Struct"`
	// The description of the egress proxy configuration. The description must be 2 to 256 characters in length and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the egress proxy configuration.
	//
	// example:
	//
	// epc-uf609a6a8fdf9f4d****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The name of the egress proxy configuration. The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testEpg
	EgressProxyConfigurationName *string `json:"EgressProxyConfigurationName,omitempty" xml:"EgressProxyConfigurationName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C79118A9-CA6D-5359-B53E-111BE8A03DCF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzorljcazp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of egress rule group references associated with the egress proxy configuration.
	RuleGroups []*GetEgressProxyConfigurationAttributeResponseBodyRuleGroups `json:"RuleGroups,omitempty" xml:"RuleGroups,omitempty" type:"Repeated"`
	// The status of the egress proxy configuration. Valid values:
	//
	// - **Provisioning**: The configuration is being created.
	//
	// - **Available**: The configuration is available.
	//
	// - **Configuring**: The configuration is being modified.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*GetEgressProxyConfigurationAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetEgressProxyConfigurationAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyConfigurationAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetDefaultRulePhaseActions() *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions {
	return s.DefaultRulePhaseActions
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetEgressProxyConfigurationName() *string {
	return s.EgressProxyConfigurationName
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetRuleGroups() []*GetEgressProxyConfigurationAttributeResponseBodyRuleGroups {
	return s.RuleGroups
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) GetTags() []*GetEgressProxyConfigurationAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetCreateTime(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetDefaultRulePhaseActions(v *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) *GetEgressProxyConfigurationAttributeResponseBody {
	s.DefaultRulePhaseActions = v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetDescription(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetEgressProxyConfigurationId(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetEgressProxyConfigurationName(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.EgressProxyConfigurationName = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetRequestId(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetResourceGroupId(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetRuleGroups(v []*GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) *GetEgressProxyConfigurationAttributeResponseBody {
	s.RuleGroups = v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetStatus(v string) *GetEgressProxyConfigurationAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) SetTags(v []*GetEgressProxyConfigurationAttributeResponseBodyTags) *GetEgressProxyConfigurationAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBody) Validate() error {
	if s.DefaultRulePhaseActions != nil {
		if err := s.DefaultRulePhaseActions.Validate(); err != nil {
			return err
		}
	}
	if s.RuleGroups != nil {
		for _, item := range s.RuleGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions struct {
	// The default action executed in the PostResponse phase (after a response is returned) when traffic does not match any rules. This phase only supports the value **Allow*	- (default), which indicates that the traffic is allowed.
	//
	// example:
	//
	// Allow
	PostResponse *string `json:"PostResponse,omitempty" xml:"PostResponse,omitempty"`
	// The default action executed in the PreDns phase (before DNS resolution) when traffic does not match any rules. Valid values:
	//
	// - **Allow*	- (default): allows the traffic.
	//
	// - **Deny**: denies the traffic.
	//
	// example:
	//
	// Allow
	PreDns *string `json:"PreDns,omitempty" xml:"PreDns,omitempty"`
	// The default action executed in the PreRequest phase (before a request is forwarded) when traffic does not match any rules. Valid values:
	//
	// - **Allow*	- (default): allows the traffic.
	//
	// - **Deny**: denies the traffic.
	//
	// example:
	//
	// Allow
	PreRequest *string `json:"PreRequest,omitempty" xml:"PreRequest,omitempty"`
}

func (s GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) GoString() string {
	return s.String()
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) GetPostResponse() *string {
	return s.PostResponse
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) GetPreDns() *string {
	return s.PreDns
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) GetPreRequest() *string {
	return s.PreRequest
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) SetPostResponse(v string) *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions {
	s.PostResponse = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) SetPreDns(v string) *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions {
	s.PreDns = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) SetPreRequest(v string) *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions {
	s.PreRequest = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyDefaultRulePhaseActions) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyConfigurationAttributeResponseBodyRuleGroups struct {
	// The ID of the egress rule group.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// The priority of the egress rule group in the egress proxy configuration. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) GoString() string {
	return s.String()
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) GetPriority() *int32 {
	return s.Priority
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) SetEgressProxyRuleGroupId(v string) *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) SetPriority(v int32) *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups {
	s.Priority = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyRuleGroups) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyConfigurationAttributeResponseBodyTags struct {
	// The tag key. The key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEgressProxyConfigurationAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyConfigurationAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyTags) SetKey(v string) *GetEgressProxyConfigurationAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyTags) SetValue(v string) *GetEgressProxyConfigurationAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetEgressProxyConfigurationAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
