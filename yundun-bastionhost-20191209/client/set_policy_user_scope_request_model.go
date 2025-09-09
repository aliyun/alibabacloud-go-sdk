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
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// >  You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the control policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of users to whom the control policy applies. Valid values:
	//
	// 	- **All**: The control policy applies to all users.
	//
	// 	- **User**: The control policy applies to specified users.
	//
	// 	- **UserGroup**: The control policy applies to specified user groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
	// The user groups to which the control policy applies.
	//
	// > This parameter is required if ScopeType is set to UserGroup. You can specify up to 100 user group IDs.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The users to whom the control policy applies.
	//
	// > This parameter is required if ScopeType is set to User. You can specify up to 500 user IDs.
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
