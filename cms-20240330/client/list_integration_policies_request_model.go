// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *ListIntegrationPoliciesRequest
	GetAddonName() *string
	SetEntityGroupIds(v string) *ListIntegrationPoliciesRequest
	GetEntityGroupIds() *string
	SetFilterRegionIds(v string) *ListIntegrationPoliciesRequest
	GetFilterRegionIds() *string
	SetMaxResults(v int32) *ListIntegrationPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIntegrationPoliciesRequest
	GetNextToken() *string
	SetPolicyId(v string) *ListIntegrationPoliciesRequest
	GetPolicyId() *string
	SetPolicyName(v string) *ListIntegrationPoliciesRequest
	GetPolicyName() *string
	SetPolicyType(v string) *ListIntegrationPoliciesRequest
	GetPolicyType() *string
	SetPrometheusInstanceId(v string) *ListIntegrationPoliciesRequest
	GetPrometheusInstanceId() *string
	SetQuery(v string) *ListIntegrationPoliciesRequest
	GetQuery() *string
	SetResourceGroupId(v string) *ListIntegrationPoliciesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListIntegrationPoliciesRequestTag) *ListIntegrationPoliciesRequest
	GetTag() []*ListIntegrationPoliciesRequestTag
	SetWorkspace(v string) *ListIntegrationPoliciesRequest
	GetWorkspace() *string
}

type ListIntegrationPoliciesRequest struct {
	// example:
	//
	// cs-default
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
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
	ResourceGroupId *string                              `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tag             []*ListIntegrationPoliciesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListIntegrationPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *ListIntegrationPoliciesRequest) GetEntityGroupIds() *string {
	return s.EntityGroupIds
}

func (s *ListIntegrationPoliciesRequest) GetFilterRegionIds() *string {
	return s.FilterRegionIds
}

func (s *ListIntegrationPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIntegrationPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIntegrationPoliciesRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListIntegrationPoliciesRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListIntegrationPoliciesRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListIntegrationPoliciesRequest) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *ListIntegrationPoliciesRequest) GetQuery() *string {
	return s.Query
}

func (s *ListIntegrationPoliciesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIntegrationPoliciesRequest) GetTag() []*ListIntegrationPoliciesRequestTag {
	return s.Tag
}

func (s *ListIntegrationPoliciesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListIntegrationPoliciesRequest) SetAddonName(v string) *ListIntegrationPoliciesRequest {
	s.AddonName = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetEntityGroupIds(v string) *ListIntegrationPoliciesRequest {
	s.EntityGroupIds = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetFilterRegionIds(v string) *ListIntegrationPoliciesRequest {
	s.FilterRegionIds = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetMaxResults(v int32) *ListIntegrationPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetNextToken(v string) *ListIntegrationPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetPolicyId(v string) *ListIntegrationPoliciesRequest {
	s.PolicyId = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetPolicyName(v string) *ListIntegrationPoliciesRequest {
	s.PolicyName = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetPolicyType(v string) *ListIntegrationPoliciesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetPrometheusInstanceId(v string) *ListIntegrationPoliciesRequest {
	s.PrometheusInstanceId = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetQuery(v string) *ListIntegrationPoliciesRequest {
	s.Query = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetResourceGroupId(v string) *ListIntegrationPoliciesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetTag(v []*ListIntegrationPoliciesRequestTag) *ListIntegrationPoliciesRequest {
	s.Tag = v
	return s
}

func (s *ListIntegrationPoliciesRequest) SetWorkspace(v string) *ListIntegrationPoliciesRequest {
	s.Workspace = &v
	return s
}

func (s *ListIntegrationPoliciesRequest) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesRequestTag struct {
	// example:
	//
	// test
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListIntegrationPoliciesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesRequestTag) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListIntegrationPoliciesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListIntegrationPoliciesRequestTag) SetKey(v string) *ListIntegrationPoliciesRequestTag {
	s.Key = &v
	return s
}

func (s *ListIntegrationPoliciesRequestTag) SetValue(v string) *ListIntegrationPoliciesRequestTag {
	s.Value = &v
	return s
}

func (s *ListIntegrationPoliciesRequestTag) Validate() error {
	return dara.Validate(s)
}
