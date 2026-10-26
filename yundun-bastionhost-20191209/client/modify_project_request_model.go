// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyProjectRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyProjectRequest
	GetInstanceId() *string
	SetMemberIds(v string) *ModifyProjectRequest
	GetMemberIds() *string
	SetName(v string) *ModifyProjectRequest
	GetName() *string
	SetProjectId(v int64) *ModifyProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyProjectRequest
	GetRegionId() *string
}

type ModifyProjectRequest struct {
	// The new description of the project. The description can be up to 500 characters in length.
	//
	// example:
	//
	// Bastionhost demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID of the bastion host for which you want to modify project information.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2v120f0y
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UIDs of the member accounts to associate with the project after modification.
	//
	// > You can call the ListMembersNotForProject operation to obtain this parameter. If you do not specify this parameter, no changes are made.
	//
	// example:
	//
	// 160****5312
	MemberIds *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
	// The new name of the project. The name must be 1 to 128 characters in length, cannot start with a special character, and can contain only the following special characters: periods (.), underscores (_), hyphens (-), backslashes (\\), and spaces.
	//
	// example:
	//
	// projectname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the project that you want to modify.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyProjectRequest) GetMemberIds() *string {
	return s.MemberIds
}

func (s *ModifyProjectRequest) GetName() *string {
	return s.Name
}

func (s *ModifyProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ModifyProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyProjectRequest) SetDescription(v string) *ModifyProjectRequest {
	s.Description = &v
	return s
}

func (s *ModifyProjectRequest) SetInstanceId(v string) *ModifyProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProjectRequest) SetMemberIds(v string) *ModifyProjectRequest {
	s.MemberIds = &v
	return s
}

func (s *ModifyProjectRequest) SetName(v string) *ModifyProjectRequest {
	s.Name = &v
	return s
}

func (s *ModifyProjectRequest) SetProjectId(v int64) *ModifyProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyProjectRequest) SetRegionId(v string) *ModifyProjectRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProjectRequest) Validate() error {
	return dara.Validate(s)
}
