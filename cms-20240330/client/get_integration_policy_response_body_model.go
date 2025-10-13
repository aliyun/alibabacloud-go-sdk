// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicy(v *GetIntegrationPolicyResponseBodyPolicy) *GetIntegrationPolicyResponseBody
	GetPolicy() *GetIntegrationPolicyResponseBodyPolicy
	SetRequestId(v string) *GetIntegrationPolicyResponseBody
	GetRequestId() *string
}

type GetIntegrationPolicyResponseBody struct {
	Policy *GetIntegrationPolicyResponseBodyPolicy `json:"policy,omitempty" xml:"policy,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetIntegrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBody) GetPolicy() *GetIntegrationPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetIntegrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntegrationPolicyResponseBody) SetPolicy(v *GetIntegrationPolicyResponseBodyPolicy) *GetIntegrationPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetIntegrationPolicyResponseBody) SetRequestId(v string) *GetIntegrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBody) Validate() error {
	if s.Policy != nil {
		if err := s.Policy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIntegrationPolicyResponseBodyPolicy struct {
	BindResource *GetIntegrationPolicyResponseBodyPolicyBindResource `json:"bindResource,omitempty" xml:"bindResource,omitempty" type:"Struct"`
	EntityGroup  *GetIntegrationPolicyResponseBodyPolicyEntityGroup  `json:"entityGroup,omitempty" xml:"entityGroup,omitempty" type:"Struct"`
	ManagedInfo  *GetIntegrationPolicyResponseBodyPolicyManagedInfo  `json:"managedInfo,omitempty" xml:"managedInfo,omitempty" type:"Struct"`
	// example:
	//
	// policy-c9efed2b99c348d49e589c5f780fc074
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// example:
	//
	// ControlPolicy4DetailVportInfo
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// example:
	//
	// CS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string                                       `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*GetIntegrationPolicyResponseBodyPolicyTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// u123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// prometheus
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetBindResource() *GetIntegrationPolicyResponseBodyPolicyBindResource {
	return s.BindResource
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetEntityGroup() *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	return s.EntityGroup
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetManagedInfo() *GetIntegrationPolicyResponseBodyPolicyManagedInfo {
	return s.ManagedInfo
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetTags() []*GetIntegrationPolicyResponseBodyPolicyTags {
	return s.Tags
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetUserId() *string {
	return s.UserId
}

func (s *GetIntegrationPolicyResponseBodyPolicy) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetBindResource(v *GetIntegrationPolicyResponseBodyPolicyBindResource) *GetIntegrationPolicyResponseBodyPolicy {
	s.BindResource = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetEntityGroup(v *GetIntegrationPolicyResponseBodyPolicyEntityGroup) *GetIntegrationPolicyResponseBodyPolicy {
	s.EntityGroup = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetManagedInfo(v *GetIntegrationPolicyResponseBodyPolicyManagedInfo) *GetIntegrationPolicyResponseBodyPolicy {
	s.ManagedInfo = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetPolicyId(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetPolicyName(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetPolicyType(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.PolicyType = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetRegionId(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.RegionId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetResourceGroupId(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.ResourceGroupId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetTags(v []*GetIntegrationPolicyResponseBodyPolicyTags) *GetIntegrationPolicyResponseBodyPolicy {
	s.Tags = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetUserId(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.UserId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) SetWorkspace(v string) *GetIntegrationPolicyResponseBodyPolicy {
	s.Workspace = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicy) Validate() error {
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

type GetIntegrationPolicyResponseBodyPolicyBindResource struct {
	// example:
	//
	// 00b1630f02814f95a9bce717d8d56bb2
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// example:
	//
	// 10.12.0.1/16
	VpcCidr *string `json:"vpcCidr,omitempty" xml:"vpcCidr,omitempty"`
	// example:
	//
	// vpc-2zegqpeyxplhtmdg70xnr
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicyBindResource) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyBindResource) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) GetVpcCidr() *string {
	return s.VpcCidr
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) GetVpcId() *string {
	return s.VpcId
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) SetClusterId(v string) *GetIntegrationPolicyResponseBodyPolicyBindResource {
	s.ClusterId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) SetClusterType(v string) *GetIntegrationPolicyResponseBodyPolicyBindResource {
	s.ClusterType = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) SetVpcCidr(v string) *GetIntegrationPolicyResponseBodyPolicyBindResource {
	s.VpcCidr = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) SetVpcId(v string) *GetIntegrationPolicyResponseBodyPolicyBindResource {
	s.VpcId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyBindResource) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyEntityGroup struct {
	// example:
	//
	// xxxxxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// eg-b79f65d11fb94e779867cf937c3a3002
	EntityGroupId *string `json:"entityGroupId,omitempty" xml:"entityGroupId,omitempty"`
	// example:
	//
	// prod-database
	EntityGroupName *string                                                       `json:"entityGroupName,omitempty" xml:"entityGroupName,omitempty"`
	EntityRules     *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules `json:"entityRules,omitempty" xml:"entityRules,omitempty" type:"Struct"`
	// example:
	//
	// status: 200 AND totalTime > 0.5
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// u123456
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// test-api
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroup) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroup) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetDescription() *string {
	return s.Description
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetEntityGroupId() *string {
	return s.EntityGroupId
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetEntityGroupName() *string {
	return s.EntityGroupName
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetEntityRules() *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	return s.EntityRules
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetQuery() *string {
	return s.Query
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetUserId() *string {
	return s.UserId
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetDescription(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.Description = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetEntityGroupId(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.EntityGroupId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetEntityGroupName(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.EntityGroupName = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetEntityRules(v *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.EntityRules = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetQuery(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.Query = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetRegionId(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.RegionId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetUserId(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.UserId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) SetWorkspace(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroup {
	s.Workspace = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroup) Validate() error {
	if s.EntityRules != nil {
		if err := s.EntityRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules struct {
	Annotations []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Repeated"`
	EntityTypes []*string                                                                  `json:"entityTypes,omitempty" xml:"entityTypes,omitempty" type:"Repeated"`
	FieldRules  []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules  `json:"fieldRules,omitempty" xml:"fieldRules,omitempty" type:"Repeated"`
	InstanceIds []*string                                                                  `json:"instanceIds,omitempty" xml:"instanceIds,omitempty" type:"Repeated"`
	IpMatchRule *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule   `json:"ipMatchRule,omitempty" xml:"ipMatchRule,omitempty" type:"Struct"`
	Labels      []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels      `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	RegionIds   []*string                                                                  `json:"regionIds,omitempty" xml:"regionIds,omitempty" type:"Repeated"`
	// example:
	//
	// rg-aek3aqsuvlv3yyq
	ResourceGroupId *string                                                             `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetAnnotations() []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations {
	return s.Annotations
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetFieldRules() []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules {
	return s.FieldRules
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetIpMatchRule() *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule {
	return s.IpMatchRule
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetLabels() []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels {
	return s.Labels
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) GetTags() []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags {
	return s.Tags
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetAnnotations(v []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.Annotations = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetEntityTypes(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.EntityTypes = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetFieldRules(v []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.FieldRules = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetInstanceIds(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.InstanceIds = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetIpMatchRule(v *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.IpMatchRule = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetLabels(v []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.Labels = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetRegionIds(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.RegionIds = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetResourceGroupId(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.ResourceGroupId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) SetTags(v []*GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules {
	s.Tags = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRules) Validate() error {
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

type GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations struct {
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// example:
	//
	// use
	TagKey    *string   `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) GetOp() *string {
	return s.Op
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) GetTagKey() *string {
	return s.TagKey
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) GetTagValues() []*string {
	return s.TagValues
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) SetOp(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations {
	s.Op = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) SetTagKey(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations {
	s.TagKey = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) SetTagValues(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations {
	s.TagValues = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesAnnotations) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules struct {
	// example:
	//
	// test
	FieldKey    *string   `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	FieldValues []*string `json:"fieldValues,omitempty" xml:"fieldValues,omitempty" type:"Repeated"`
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) GetFieldKey() *string {
	return s.FieldKey
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) GetFieldValues() []*string {
	return s.FieldValues
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) GetOp() *string {
	return s.Op
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) SetFieldKey(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules {
	s.FieldKey = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) SetFieldValues(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules {
	s.FieldValues = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) SetOp(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules {
	s.Op = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesFieldRules) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule struct {
	// example:
	//
	// 10.10.0.1/16
	IpCidr *string `json:"ipCidr,omitempty" xml:"ipCidr,omitempty"`
	// example:
	//
	// test-key
	IpFieldKey *string `json:"ipFieldKey,omitempty" xml:"ipFieldKey,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) GetIpCidr() *string {
	return s.IpCidr
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) GetIpFieldKey() *string {
	return s.IpFieldKey
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) SetIpCidr(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule {
	s.IpCidr = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) SetIpFieldKey(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule {
	s.IpFieldKey = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesIpMatchRule) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels struct {
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// example:
	//
	// key1
	TagKey    *string   `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) GetOp() *string {
	return s.Op
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) GetTagKey() *string {
	return s.TagKey
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) GetTagValues() []*string {
	return s.TagValues
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) SetOp(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels {
	s.Op = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) SetTagKey(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels {
	s.TagKey = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) SetTagValues(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels {
	s.TagValues = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesLabels) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags struct {
	// example:
	//
	// add
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// example:
	//
	// key2
	TagKey    *string   `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValues []*string `json:"tagValues,omitempty" xml:"tagValues,omitempty" type:"Repeated"`
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) GetOp() *string {
	return s.Op
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) GetTagValues() []*string {
	return s.TagValues
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) SetOp(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags {
	s.Op = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) SetTagKey(v string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags {
	s.TagKey = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) SetTagValues(v []*string) *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags {
	s.TagValues = v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyEntityGroupEntityRulesTags) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyManagedInfo struct {
	// example:
	//
	// sg-xxxxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// vsw-xxxxxxxxx
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicyManagedInfo) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyManagedInfo) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyManagedInfo) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetIntegrationPolicyResponseBodyPolicyManagedInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetIntegrationPolicyResponseBodyPolicyManagedInfo) SetSecurityGroupId(v string) *GetIntegrationPolicyResponseBodyPolicyManagedInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyManagedInfo) SetVswitchId(v string) *GetIntegrationPolicyResponseBodyPolicyManagedInfo {
	s.VswitchId = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyManagedInfo) Validate() error {
	return dara.Validate(s)
}

type GetIntegrationPolicyResponseBodyPolicyTags struct {
	// example:
	//
	// use
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// db
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetIntegrationPolicyResponseBodyPolicyTags) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyResponseBodyPolicyTags) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyResponseBodyPolicyTags) GetKey() *string {
	return s.Key
}

func (s *GetIntegrationPolicyResponseBodyPolicyTags) GetValue() *string {
	return s.Value
}

func (s *GetIntegrationPolicyResponseBodyPolicyTags) SetKey(v string) *GetIntegrationPolicyResponseBodyPolicyTags {
	s.Key = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyTags) SetValue(v string) *GetIntegrationPolicyResponseBodyPolicyTags {
	s.Value = &v
	return s
}

func (s *GetIntegrationPolicyResponseBodyPolicyTags) Validate() error {
	return dara.Validate(s)
}
