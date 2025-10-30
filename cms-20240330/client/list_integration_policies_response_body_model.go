// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIntegrationPoliciesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListIntegrationPoliciesResponseBody
	GetNextToken() *string
	SetPolicies(v []*ListIntegrationPoliciesResponseBodyPolicies) *ListIntegrationPoliciesResponseBody
	GetPolicies() []*ListIntegrationPoliciesResponseBodyPolicies
	SetRequestId(v string) *ListIntegrationPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIntegrationPoliciesResponseBody
	GetTotalCount() *int32
}

type ListIntegrationPoliciesResponseBody struct {
	// Page size
	//
	// Default value:
	//
	// 	50
	//
	// Maximum value:
	//
	// 	50
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Pagination Token
	//
	// example:
	//
	// 44ANBjKZmQeKnaB1fXRq06w7sFYK3MUcCALMD9qQbmEiE
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Access policy list
	Policies []*ListIntegrationPoliciesResponseBodyPolicies `json:"policies,omitempty" xml:"policies,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Total number of entries
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListIntegrationPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIntegrationPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIntegrationPoliciesResponseBody) GetPolicies() []*ListIntegrationPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *ListIntegrationPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIntegrationPoliciesResponseBody) SetMaxResults(v int32) *ListIntegrationPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBody) SetNextToken(v string) *ListIntegrationPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBody) SetPolicies(v []*ListIntegrationPoliciesResponseBodyPolicies) *ListIntegrationPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListIntegrationPoliciesResponseBody) SetRequestId(v string) *ListIntegrationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBody) SetTotalCount(v int32) *ListIntegrationPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPoliciesResponseBodyPolicies struct {
	// Bound resource information
	BindResource *ListIntegrationPoliciesResponseBodyPoliciesBindResource `json:"bindResource,omitempty" xml:"bindResource,omitempty" type:"Struct"`
	// Cs Umodel Status
	//
	// example:
	//
	// true
	CsUmodelStatus *bool `json:"csUmodelStatus,omitempty" xml:"csUmodelStatus,omitempty"`
	// Entity group
	EntityGroup *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup `json:"entityGroup,omitempty" xml:"entityGroup,omitempty" type:"Struct"`
	FeePackage  *string                                                 `json:"feePackage,omitempty" xml:"feePackage,omitempty"`
	// Policy network management information.
	ManagedInfo *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo `json:"managedInfo,omitempty" xml:"managedInfo,omitempty" type:"Struct"`
	// Policy ID.
	//
	// example:
	//
	// policy-ac38a7cb02d14ff48bc9f97d0a75063e
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// Rule name.
	//
	// example:
	//
	// 6f5HSsg3AP
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// Access center policy type
	//
	// example:
	//
	// ECS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Number of sub-releases
	SubAddonRelease *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease `json:"subAddonRelease,omitempty" xml:"subAddonRelease,omitempty" type:"Struct"`
	// Resource tag key values.
	Tags []*ListIntegrationPoliciesResponseBodyPoliciesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// User ID
	//
	// example:
	//
	// 128470923
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Workspace.
	//
	// example:
	//
	// test-api
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetBindResource() *ListIntegrationPoliciesResponseBodyPoliciesBindResource {
	return s.BindResource
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetCsUmodelStatus() *bool {
	return s.CsUmodelStatus
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetEntityGroup() *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	return s.EntityGroup
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetFeePackage() *string {
	return s.FeePackage
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetManagedInfo() *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo {
	return s.ManagedInfo
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetSubAddonRelease() *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease {
	return s.SubAddonRelease
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetTags() []*ListIntegrationPoliciesResponseBodyPoliciesTags {
	return s.Tags
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetUserId() *string {
	return s.UserId
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetBindResource(v *ListIntegrationPoliciesResponseBodyPoliciesBindResource) *ListIntegrationPoliciesResponseBodyPolicies {
	s.BindResource = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetCsUmodelStatus(v bool) *ListIntegrationPoliciesResponseBodyPolicies {
	s.CsUmodelStatus = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetEntityGroup(v *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) *ListIntegrationPoliciesResponseBodyPolicies {
	s.EntityGroup = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetFeePackage(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.FeePackage = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetManagedInfo(v *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) *ListIntegrationPoliciesResponseBodyPolicies {
	s.ManagedInfo = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetPolicyName(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetPolicyType(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetRegionId(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.RegionId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetResourceGroupId(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetSubAddonRelease(v *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) *ListIntegrationPoliciesResponseBodyPolicies {
	s.SubAddonRelease = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetTags(v []*ListIntegrationPoliciesResponseBodyPoliciesTags) *ListIntegrationPoliciesResponseBodyPolicies {
	s.Tags = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetUserId(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.UserId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) SetWorkspace(v string) *ListIntegrationPoliciesResponseBodyPolicies {
	s.Workspace = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPolicies) Validate() error {
	if s.BindResource != nil {
		if err := s.BindResource.Validate(); err != nil {
			return err
		}
	}
	if s.EntityGroup != nil {
		if err := s.EntityGroup.Validate(); err != nil {
			return err
		}
	}
	if s.ManagedInfo != nil {
		if err := s.ManagedInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SubAddonRelease != nil {
		if err := s.SubAddonRelease.Validate(); err != nil {
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

type ListIntegrationPoliciesResponseBodyPoliciesBindResource struct {
	// Cluster ID.
	//
	// example:
	//
	// cv68tt87d78vyc89zy9v
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// Cluster type.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// VPC CIDR
	//
	// example:
	//
	// 100.100.0.1/16
	VpcCidr *string `json:"vpcCidr,omitempty" xml:"vpcCidr,omitempty"`
	// Virtual Private Cloud (VPC).
	//
	// example:
	//
	// vpc-uf664nyle5khp5d4d7hdo
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesBindResource) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesBindResource) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) GetVpcCidr() *string {
	return s.VpcCidr
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) SetClusterId(v string) *ListIntegrationPoliciesResponseBodyPoliciesBindResource {
	s.ClusterId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) SetClusterType(v string) *ListIntegrationPoliciesResponseBodyPoliciesBindResource {
	s.ClusterType = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) SetVpcCidr(v string) *ListIntegrationPoliciesResponseBodyPoliciesBindResource {
	s.VpcCidr = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) SetVpcId(v string) *ListIntegrationPoliciesResponseBodyPoliciesBindResource {
	s.VpcId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesBindResource) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroup struct {
	// Description.
	//
	// example:
	//
	// xxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Entity group ID
	//
	// example:
	//
	// eg-b79f65d11fb94e779867cf937c3a3002
	EntityGroupId *string `json:"entityGroupId,omitempty" xml:"entityGroupId,omitempty"`
	// Entity group name
	//
	// example:
	//
	// test-eg
	EntityGroupName *string `json:"entityGroupName,omitempty" xml:"entityGroupName,omitempty"`
	// Entity group
	EntityRules *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules `json:"entityRules,omitempty" xml:"entityRules,omitempty" type:"Struct"`
	// Search keyword, supports document library name and description
	//
	// example:
	//
	// 哈弗
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// User ID
	//
	// example:
	//
	// 1236812738
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// Workspace.
	//
	// example:
	//
	// test-api
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetDescription() *string {
	return s.Description
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetEntityGroupId() *string {
	return s.EntityGroupId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetEntityGroupName() *string {
	return s.EntityGroupName
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetEntityRules() *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	return s.EntityRules
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetQuery() *string {
	return s.Query
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetUserId() *string {
	return s.UserId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetDescription(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.Description = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetEntityGroupId(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.EntityGroupId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetEntityGroupName(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.EntityGroupName = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetEntityRules(v *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.EntityRules = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetQuery(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.Query = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetRegionId(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.RegionId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetUserId(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.UserId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) SetWorkspace(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup {
	s.Workspace = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroup) Validate() error {
	if s.EntityRules != nil {
		if err := s.EntityRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules struct {
	// Annotations
	Annotations []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	// List of entity types
	EntityTypes []*string `json:"entityTypes,omitempty" xml:"entityTypes,omitempty" type:"Repeated"`
	// Field rules
	FieldRules []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules `json:"fieldRules,omitempty" xml:"fieldRules,omitempty" type:"Repeated"`
	// Instance IDs.
	InstanceIds []*string `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	// IP match rule
	IpMatchRule *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule `json:"ipMatchRule,omitempty" xml:"ipMatchRule,omitempty" type:"Struct"`
	// Labels
	Labels []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// List of region IDs
	RegionIds []*string `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
	// Resource group ID.
	//
	// example:
	//
	// rg-5i6dbwxfxuqihk7k
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Instance tag information.
	Tags []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetAnnotations() []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations {
	return s.Annotations
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetFieldRules() []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules {
	return s.FieldRules
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetIpMatchRule() *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule {
	return s.IpMatchRule
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetLabels() []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels {
	return s.Labels
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) GetTags() []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags {
	return s.Tags
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetAnnotations(v []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.Annotations = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetEntityTypes(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.EntityTypes = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetFieldRules(v []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.FieldRules = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetInstanceIds(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.InstanceIds = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetIpMatchRule(v *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.IpMatchRule = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetLabels(v []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.Labels = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetRegionIds(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.RegionIds = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetResourceGroupId(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) SetTags(v []*ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules {
	s.Tags = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRules) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FieldRules != nil {
		for _, item := range s.FieldRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IpMatchRule != nil {
		if err := s.IpMatchRule.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
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

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations struct {
	// Operation to be performed.
	//
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// Tag key
	//
	// example:
	//
	// key
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// Tag values
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) GetOp() *string {
	return s.Op
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) GetTagKey() *string {
	return s.TagKey
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) SetOp(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations {
	s.Op = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) SetTagKey(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations {
	s.TagKey = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) SetTagValues(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations {
	s.TagValues = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesAnnotations) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules struct {
	// Unique identifier for the field.
	//
	// example:
	//
	// test
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// Field content, multiple values separated by commas.
	FieldValues []*string `json:"fieldValues,omitempty" xml:"fieldValues,omitempty" type:"Repeated"`
	// Operation to be performed.
	//
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) GetFieldKey() *string {
	return s.FieldKey
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) GetFieldValues() []*string {
	return s.FieldValues
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) GetOp() *string {
	return s.Op
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) SetFieldKey(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules {
	s.FieldKey = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) SetFieldValues(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules {
	s.FieldValues = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) SetOp(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules {
	s.Op = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesFieldRules) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule struct {
	// IP segment
	//
	// example:
	//
	// 100.100.1.0/16
	IpCidr *string `json:"ipCidr,omitempty" xml:"ipCidr,omitempty"`
	// Key of the IP field
	//
	// example:
	//
	// xxxx
	IpFieldKey *string `json:"ipFieldKey,omitempty" xml:"ipFieldKey,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) GetIpCidr() *string {
	return s.IpCidr
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) GetIpFieldKey() *string {
	return s.IpFieldKey
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) SetIpCidr(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule {
	s.IpCidr = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) SetIpFieldKey(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule {
	s.IpFieldKey = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesIpMatchRule) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels struct {
	// Operation to be performed.
	//
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// Tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// Tag values
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) GetOp() *string {
	return s.Op
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) GetTagKey() *string {
	return s.TagKey
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) SetOp(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels {
	s.Op = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) SetTagKey(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels {
	s.TagKey = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) SetTagValues(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels {
	s.TagValues = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesLabels) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags struct {
	// Operation to be performed.
	//
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// Tag key.
	//
	// example:
	//
	// key
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// Tag value.
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) GetOp() *string {
	return s.Op
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) GetTagValues() []*string {
	return s.TagValues
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) SetOp(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags {
	s.Op = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) SetTagKey(v string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags {
	s.TagKey = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) SetTagValues(v []*string) *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags {
	s.TagValues = v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesEntityGroupEntityRulesTags) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesManagedInfo struct {
	// ENI card ID of the managed probe. For example: eni-xxxx.
	//
	// example:
	//
	// eni-12345678
	EniId *string `json:"eniId,omitempty" xml:"eniId,omitempty"`
	// Security group ID
	//
	// example:
	//
	// sg-xxxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// VSwitch ID.
	//
	// example:
	//
	// vsw-xxxxxx
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) GetEniId() *string {
	return s.EniId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) SetEniId(v string) *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo {
	s.EniId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) SetSecurityGroupId(v string) *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) SetVswitchId(v string) *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo {
	s.VswitchId = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesManagedInfo) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease struct {
	// Number of ready sub-releases
	//
	// example:
	//
	// 30
	Ready *int32 `json:"ready,omitempty" xml:"ready,omitempty"`
	// Number of rules.
	//
	// example:
	//
	// 278
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) GetReady() *int32 {
	return s.Ready
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) GetTotal() *int32 {
	return s.Total
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) SetReady(v int32) *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease {
	s.Ready = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) SetTotal(v int32) *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease {
	s.Total = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesSubAddonRelease) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPoliciesResponseBodyPoliciesTags struct {
	// Tag key
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Match value.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListIntegrationPoliciesResponseBodyPoliciesTags) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPoliciesResponseBodyPoliciesTags) GoString() string {
	return s.String()
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesTags) GetKey() *string {
	return s.Key
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesTags) GetValue() *string {
	return s.Value
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesTags) SetKey(v string) *ListIntegrationPoliciesResponseBodyPoliciesTags {
	s.Key = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesTags) SetValue(v string) *ListIntegrationPoliciesResponseBodyPoliciesTags {
	s.Value = &v
	return s
}

func (s *ListIntegrationPoliciesResponseBodyPoliciesTags) Validate() error {
	return dara.Validate(s)
}
