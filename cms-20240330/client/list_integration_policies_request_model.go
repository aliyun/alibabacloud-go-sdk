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
	SetBindResourceId(v string) *ListIntegrationPoliciesRequest
	GetBindResourceId() *string
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
	// Addon Name.
	//
	// example:
	//
	// cs-default
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// Bound Resource ID.
	//
	// example:
	//
	// 622d27c2e87d49debceeebc7c642610e
	BindResourceId *string `json:"bindResourceId,omitempty" xml:"bindResourceId,omitempty"`
	// Filter for entity IDs, separated by commas.
	//
	// example:
	//
	// eg-1,eg-2,eg-3
	EntityGroupIds *string `json:"entityGroupIds,omitempty" xml:"entityGroupIds,omitempty"`
	// Used for Region query, separated by commas.
	//
	// example:
	//
	// cn-beijing,cn-hangzhou
	FilterRegionIds *string `json:"filterRegionIds,omitempty" xml:"filterRegionIds,omitempty"`
	// Maximum number of results to return, default is 30, with a maximum of 100.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Used to return more results. This parameter is not required for the first query; for subsequent queries, use the Token obtained from the previous response.
	//
	// example:
	//
	// mvnX6zqg3P
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// policy-93817a401f78435596d745a97d2e85a1
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// Rule Name.
	//
	// example:
	//
	// prod-database
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// Policy Type
	//
	// example:
	//
	// CS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// cmee-622d27c2e87d49debceeebc7c642610e
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// Used for general queries.
	//
	// example:
	//
	// test
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Resource Group ID.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Tag list.
	Tag []*ListIntegrationPoliciesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// Workspace.
	//
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

func (s *ListIntegrationPoliciesRequest) GetBindResourceId() *string {
	return s.BindResourceId
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

func (s *ListIntegrationPoliciesRequest) SetBindResourceId(v string) *ListIntegrationPoliciesRequest {
	s.BindResourceId = &v
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

type ListIntegrationPoliciesRequestTag struct {
	// Tag key
	//
	// example:
	//
	// test
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Tag value
	//
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
