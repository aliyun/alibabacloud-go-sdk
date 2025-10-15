// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPoliciesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListIntegrationPoliciesShrinkRequest
	GetAddonName() *string
	SetBindResourceId(v string) *ListIntegrationPoliciesShrinkRequest
	GetBindResourceId() *string
	SetEntityGroupIds(v string) *ListIntegrationPoliciesShrinkRequest
	GetEntityGroupIds() *string
	SetFilterRegionIds(v string) *ListIntegrationPoliciesShrinkRequest
	GetFilterRegionIds() *string
	SetMaxResults(v int32) *ListIntegrationPoliciesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIntegrationPoliciesShrinkRequest
	GetNextToken() *string
	SetPolicyId(v string) *ListIntegrationPoliciesShrinkRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ListIntegrationPoliciesShrinkRequest
	GetPolicyName() *string
	SetPolicyType(v string) *ListIntegrationPoliciesShrinkRequest
	GetPolicyType() *string
	SetPrometheusInstanceId(v string) *ListIntegrationPoliciesShrinkRequest
	GetPrometheusInstanceId() *string
	SetQuery(v string) *ListIntegrationPoliciesShrinkRequest
	GetQuery() *string
	SetResourceGroupId(v string) *ListIntegrationPoliciesShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *ListIntegrationPoliciesShrinkRequest
	GetTagShrink() *string
	SetWorkspace(v string) *ListIntegrationPoliciesShrinkRequest
	GetWorkspace() *string
}

type ListIntegrationPoliciesShrinkRequest struct {
	// example:
	//
	// cs-default
	AddonName      *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	BindResourceId *string `json:"bindResourceId,omitempty" xml:"bindResourceId,omitempty"`
	// example:
	//
	// eg-1,eg-2,eg-3
	EntityGroupIds *string `json:"entityGroupIds,omitempty" xml:"entityGroupIds,omitempty"`
	// example:
	//
	// cn-beijing,cn-hangzhou
	FilterRegionIds *string `json:"filterRegionIds,omitempty" xml:"filterRegionIds,omitempty"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// mvnX6zqg3P
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// policy-93817a401f78435596d745a97d2e85a1
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// example:
	//
	// prod-database
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// example:
	//
	// CS
	PolicyType           *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// example:
	//
	// test
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	TagShrink       *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// demo
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListIntegrationPoliciesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesShrinkRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPoliciesShrinkRequest) GetBindResourceId() *string {
	return s.BindResourceId
}

func (s *ListIntegrationPoliciesShrinkRequest) GetEntityGroupIds() *string {
	return s.EntityGroupIds
}

func (s *ListIntegrationPoliciesShrinkRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListIntegrationPoliciesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIntegrationPoliciesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIntegrationPoliciesShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListIntegrationPoliciesShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListIntegrationPoliciesShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListIntegrationPoliciesShrinkRequest) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *ListIntegrationPoliciesShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *ListIntegrationPoliciesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIntegrationPoliciesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListIntegrationPoliciesShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListIntegrationPoliciesShrinkRequest) SetAddonName(v string) *ListIntegrationPoliciesShrinkRequest {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetBindResourceId(v string) *ListIntegrationPoliciesShrinkRequest {
	s.BindResourceId = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetEntityGroupIds(v string) *ListIntegrationPoliciesShrinkRequest {
	s.EntityGroupIds = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetFilterRegionIds(v string) *ListIntegrationPoliciesShrinkRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetMaxResults(v int32) *ListIntegrationPoliciesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetNextToken(v string) *ListIntegrationPoliciesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetPolicyId(v string) *ListIntegrationPoliciesShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetPolicyName(v string) *ListIntegrationPoliciesShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetPolicyType(v string) *ListIntegrationPoliciesShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetPrometheusInstanceId(v string) *ListIntegrationPoliciesShrinkRequest {
	s.PrometheusInstanceId = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetQuery(v string) *ListIntegrationPoliciesShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetResourceGroupId(v string) *ListIntegrationPoliciesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetTagShrink(v string) *ListIntegrationPoliciesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) SetWorkspace(v string) *ListIntegrationPoliciesShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListIntegrationPoliciesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
