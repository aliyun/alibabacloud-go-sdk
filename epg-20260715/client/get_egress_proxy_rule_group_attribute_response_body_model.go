// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyRuleGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetDescription() *string
	SetEgressProxyIpPoolId(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetEgressProxyIpPoolId() *string
	SetEgressProxyRuleGroupId(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetEgressProxyRuleGroupId() *string
	SetEgressProxyRuleGroupName(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetEgressProxyRuleGroupName() *string
	SetEgressProxyRules(v *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) *GetEgressProxyRuleGroupAttributeResponseBody
	GetEgressProxyRules() *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules
	SetRequestId(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetEgressProxyRuleGroupAttributeResponseBody
	GetStatus() *string
	SetTags(v []*GetEgressProxyRuleGroupAttributeResponseBodyTags) *GetEgressProxyRuleGroupAttributeResponseBody
	GetTags() []*GetEgressProxyRuleGroupAttributeResponseBodyTags
}

type GetEgressProxyRuleGroupAttributeResponseBody struct {
	// example:
	//
	// 2025-10-16T15:34:56Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyRuleGroupName *string                                                       `json:"EgressProxyRuleGroupName,omitempty" xml:"EgressProxyRuleGroupName,omitempty"`
	EgressProxyRules         *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules `json:"EgressProxyRules,omitempty" xml:"EgressProxyRules,omitempty" type:"Struct"`
	// example:
	//
	// 01A05AB5-5BED-56ED-8EE9-304FBC2F8537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmzvr2nr****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Available
	Status *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*GetEgressProxyRuleGroupAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetEgressProxyRuleGroupName() *string {
	return s.EgressProxyRuleGroupName
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetEgressProxyRules() *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules {
	return s.EgressProxyRules
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) GetTags() []*GetEgressProxyRuleGroupAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetCreateTime(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetDescription(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetEgressProxyIpPoolId(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetEgressProxyRuleGroupId(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetEgressProxyRuleGroupName(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.EgressProxyRuleGroupName = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetEgressProxyRules(v *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.EgressProxyRules = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetRequestId(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetResourceGroupId(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetStatus(v string) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) SetTags(v []*GetEgressProxyRuleGroupAttributeResponseBodyTags) *GetEgressProxyRuleGroupAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBody) Validate() error {
	if s.EgressProxyRules != nil {
		if err := s.EgressProxyRules.Validate(); err != nil {
			return err
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

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules struct {
	PostResponse []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse `json:"PostResponse,omitempty" xml:"PostResponse,omitempty" type:"Repeated"`
	PreDns       []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns       `json:"PreDns,omitempty" xml:"PreDns,omitempty" type:"Repeated"`
	PreRequest   []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest   `json:"PreRequest,omitempty" xml:"PreRequest,omitempty" type:"Repeated"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) GetPostResponse() []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	return s.PostResponse
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) GetPreDns() []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	return s.PreDns
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) GetPreRequest() []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	return s.PreRequest
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) SetPostResponse(v []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules {
	s.PostResponse = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) SetPreDns(v []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules {
	s.PreDns = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) SetPreRequest(v []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules {
	s.PreRequest = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRules) Validate() error {
	if s.PostResponse != nil {
		for _, item := range s.PostResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreDns != nil {
		for _, item := range s.PreDns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreRequest != nil {
		for _, item := range s.PreRequest {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse struct {
	// example:
	//
	// Allow
	Action     *string                                                                               `json:"Action,omitempty" xml:"Action,omitempty"`
	Conditions []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-01-15T11:58:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 66
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetAction() *string {
	return s.Action
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetConditions() []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions {
	return s.Conditions
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) GetPriority() *int32 {
	return s.Priority
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetAction(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.Action = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetConditions(v []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.Conditions = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetCreateTime(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetDescription(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.Description = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetEgressProxyRuleId(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.EgressProxyRuleId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetEgressProxyRuleName(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.EgressProxyRuleName = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) SetPriority(v int32) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse {
	s.Priority = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponse) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions struct {
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// example:
	//
	// StringEquals
	ConditionOperator *string   `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	ConditionValues   []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) SetConditionKey(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions {
	s.ConditionKey = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) SetConditionOperator(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions {
	s.ConditionOperator = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) SetConditionValues(v []*string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions {
	s.ConditionValues = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPostResponseConditions) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns struct {
	// example:
	//
	// Allow
	Action     *string                                                                         `json:"Action,omitempty" xml:"Action,omitempty"`
	Conditions []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-08-19T03:34:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 66
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetAction() *string {
	return s.Action
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetConditions() []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions {
	return s.Conditions
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) GetPriority() *int32 {
	return s.Priority
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetAction(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.Action = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetConditions(v []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.Conditions = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetCreateTime(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetDescription(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.Description = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetEgressProxyRuleId(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.EgressProxyRuleId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetEgressProxyRuleName(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.EgressProxyRuleName = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) SetPriority(v int32) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns {
	s.Priority = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDns) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions struct {
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// example:
	//
	// StringEquals
	ConditionOperator *string   `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	ConditionValues   []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) SetConditionKey(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions {
	s.ConditionKey = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) SetConditionOperator(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions {
	s.ConditionOperator = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) SetConditionValues(v []*string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions {
	s.ConditionValues = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreDnsConditions) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest struct {
	// example:
	//
	// Allow
	Action     *string                                                                             `json:"Action,omitempty" xml:"Action,omitempty"`
	Conditions []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-12-20T15:58:33.343
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 66
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetAction() *string {
	return s.Action
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetConditions() []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions {
	return s.Conditions
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetEgressProxyRuleName() *string {
	return s.EgressProxyRuleName
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetAction(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.Action = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetConditions(v []*GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.Conditions = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetCreateTime(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetDescription(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.Description = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetEgressProxyRuleId(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.EgressProxyRuleId = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetEgressProxyRuleName(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.EgressProxyRuleName = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) SetPriority(v int32) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest {
	s.Priority = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequest) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions struct {
	// example:
	//
	// request:DestinationDomain
	ConditionKey *string `json:"ConditionKey,omitempty" xml:"ConditionKey,omitempty"`
	// example:
	//
	// StringEquals
	ConditionOperator *string   `json:"ConditionOperator,omitempty" xml:"ConditionOperator,omitempty"`
	ConditionValues   []*string `json:"ConditionValues,omitempty" xml:"ConditionValues,omitempty" type:"Repeated"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) GetConditionKey() *string {
	return s.ConditionKey
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) GetConditionOperator() *string {
	return s.ConditionOperator
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) GetConditionValues() []*string {
	return s.ConditionValues
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) SetConditionKey(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions {
	s.ConditionKey = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) SetConditionOperator(v string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions {
	s.ConditionOperator = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) SetConditionValues(v []*string) *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions {
	s.ConditionValues = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyEgressProxyRulesPreRequestConditions) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyRuleGroupAttributeResponseBodyTags struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyTags) SetKey(v string) *GetEgressProxyRuleGroupAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyTags) SetValue(v string) *GetEgressProxyRuleGroupAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
