// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntegrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityGroup(v *CreateIntegrationPolicyRequestEntityGroup) *CreateIntegrationPolicyRequest
	GetEntityGroup() *CreateIntegrationPolicyRequestEntityGroup
	SetPolicyName(v string) *CreateIntegrationPolicyRequest
	GetPolicyName() *string
	SetPolicyType(v string) *CreateIntegrationPolicyRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *CreateIntegrationPolicyRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateIntegrationPolicyRequestTags) *CreateIntegrationPolicyRequest
	GetTags() []*CreateIntegrationPolicyRequestTags
	SetWorkspace(v string) *CreateIntegrationPolicyRequest
	GetWorkspace() *string
}

type CreateIntegrationPolicyRequest struct {
	EntityGroup *CreateIntegrationPolicyRequestEntityGroup `json:"entityGroup,omitempty" xml:"entityGroup,omitempty" type:"Struct"`
	// example:
	//
	// prod-database
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CS
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// example:
	//
	// rg-aekz2km4kmhtbii
	ResourceGroupId *string                               `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*CreateIntegrationPolicyRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// prometheus
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateIntegrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateIntegrationPolicyRequest) GetEntityGroup() *CreateIntegrationPolicyRequestEntityGroup {
	return s.EntityGroup
}

func (s *CreateIntegrationPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateIntegrationPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateIntegrationPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIntegrationPolicyRequest) GetTags() []*CreateIntegrationPolicyRequestTags {
	return s.Tags
}

func (s *CreateIntegrationPolicyRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateIntegrationPolicyRequest) SetEntityGroup(v *CreateIntegrationPolicyRequestEntityGroup) *CreateIntegrationPolicyRequest {
	s.EntityGroup = v
	return s
}

func (s *CreateIntegrationPolicyRequest) SetPolicyName(v string) *CreateIntegrationPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateIntegrationPolicyRequest) SetPolicyType(v string) *CreateIntegrationPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateIntegrationPolicyRequest) SetResourceGroupId(v string) *CreateIntegrationPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIntegrationPolicyRequest) SetTags(v []*CreateIntegrationPolicyRequestTags) *CreateIntegrationPolicyRequest {
	s.Tags = v
	return s
}

func (s *CreateIntegrationPolicyRequest) SetWorkspace(v string) *CreateIntegrationPolicyRequest {
	s.Workspace = &v
	return s
}

func (s *CreateIntegrationPolicyRequest) Validate() error {
	if s.EntityGroup != nil {
		if err := s.EntityGroup.Validate(); err != nil {
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

type CreateIntegrationPolicyRequestEntityGroup struct {
	// example:
	//
	// acs.ack.cluster
	ClusterEntityType *string `json:"clusterEntityType,omitempty" xml:"clusterEntityType,omitempty"`
	// example:
	//
	// na61prod3-na61cloudhdfsssd
	ClusterId          *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	DisablePolicyShare *bool   `json:"disablePolicyShare,omitempty" xml:"disablePolicyShare,omitempty"`
	// example:
	//
	// eg-b79f65d11fb94e779867cf937c3a3002
	EntityGroupId *string `json:"entityGroupId,omitempty" xml:"entityGroupId,omitempty"`
	// example:
	//
	// vpc-bp18fgg3ffxa9czna40xt
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateIntegrationPolicyRequestEntityGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationPolicyRequestEntityGroup) GoString() string {
	return s.String()
}

func (s *CreateIntegrationPolicyRequestEntityGroup) GetClusterEntityType() *string {
	return s.ClusterEntityType
}

func (s *CreateIntegrationPolicyRequestEntityGroup) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateIntegrationPolicyRequestEntityGroup) GetDisablePolicyShare() *bool {
	return s.DisablePolicyShare
}

func (s *CreateIntegrationPolicyRequestEntityGroup) GetEntityGroupId() *string {
	return s.EntityGroupId
}

func (s *CreateIntegrationPolicyRequestEntityGroup) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateIntegrationPolicyRequestEntityGroup) SetClusterEntityType(v string) *CreateIntegrationPolicyRequestEntityGroup {
	s.ClusterEntityType = &v
	return s
}

func (s *CreateIntegrationPolicyRequestEntityGroup) SetClusterId(v string) *CreateIntegrationPolicyRequestEntityGroup {
	s.ClusterId = &v
	return s
}

func (s *CreateIntegrationPolicyRequestEntityGroup) SetDisablePolicyShare(v bool) *CreateIntegrationPolicyRequestEntityGroup {
	s.DisablePolicyShare = &v
	return s
}

func (s *CreateIntegrationPolicyRequestEntityGroup) SetEntityGroupId(v string) *CreateIntegrationPolicyRequestEntityGroup {
	s.EntityGroupId = &v
	return s
}

func (s *CreateIntegrationPolicyRequestEntityGroup) SetVpcId(v string) *CreateIntegrationPolicyRequestEntityGroup {
	s.VpcId = &v
	return s
}

func (s *CreateIntegrationPolicyRequestEntityGroup) Validate() error {
	return dara.Validate(s)
}

type CreateIntegrationPolicyRequestTags struct {
	// example:
	//
	// use
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// database
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateIntegrationPolicyRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationPolicyRequestTags) GoString() string {
	return s.String()
}

func (s *CreateIntegrationPolicyRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateIntegrationPolicyRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateIntegrationPolicyRequestTags) SetKey(v string) *CreateIntegrationPolicyRequestTags {
	s.Key = &v
	return s
}

func (s *CreateIntegrationPolicyRequestTags) SetValue(v string) *CreateIntegrationPolicyRequestTags {
	s.Value = &v
	return s
}

func (s *CreateIntegrationPolicyRequestTags) Validate() error {
	return dara.Validate(s)
}
