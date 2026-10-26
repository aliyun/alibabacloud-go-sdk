// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyUserScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetPolicyUserScopeRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyUserScopeRequest
	GetPolicyId() *string
	SetProjectId(v int64) *SetPolicyUserScopeRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyUserScopeRequest
	GetRegionId() *string
	SetScopeType(v string) *SetPolicyUserScopeRequest
	GetScopeType() *string
	SetUserGroupIds(v []*string) *SetPolicyUserScopeRequest
	GetUserGroupIds() []*string
	SetUserIds(v []*string) *SetPolicyUserScopeRequest
	GetUserIds() []*string
}

type SetPolicyUserScopeRequest struct {
	// The ID of the bastion host instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user scope for the control policy. Valid values:
	//
	// - **All**: The policy applies to all users.
	//
	// - **User**: The policy applies to selected users.
	//
	// - **UserGroup**: The policy applies to selected user groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The user group scope for the control policy.
	//
	// > This parameter is required when ScopeType is set to UserGroup. A maximum of 100 user groups can be specified.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The user scope for the control policy.
	//
	// > This parameter is required when ScopeType is set to User. A maximum of 500 users can be specified.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s SetPolicyUserScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyUserScopeRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyUserScopeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyUserScopeRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyUserScopeRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SetPolicyUserScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyUserScopeRequest) GetScopeType() *string {
	return s.ScopeType
}

func (s *SetPolicyUserScopeRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *SetPolicyUserScopeRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *SetPolicyUserScopeRequest) SetInstanceId(v string) *SetPolicyUserScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyUserScopeRequest) SetPolicyId(v string) *SetPolicyUserScopeRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyUserScopeRequest) SetProjectId(v int64) *SetPolicyUserScopeRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyUserScopeRequest) SetRegionId(v string) *SetPolicyUserScopeRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyUserScopeRequest) SetScopeType(v string) *SetPolicyUserScopeRequest {
	s.ScopeType = &v
	return s
}

func (s *SetPolicyUserScopeRequest) SetUserGroupIds(v []*string) *SetPolicyUserScopeRequest {
	s.UserGroupIds = v
	return s
}

func (s *SetPolicyUserScopeRequest) SetUserIds(v []*string) *SetPolicyUserScopeRequest {
	s.UserIds = v
	return s
}

func (s *SetPolicyUserScopeRequest) Validate() error {
	return dara.Validate(s)
}
