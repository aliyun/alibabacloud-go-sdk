// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachMembersToProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachMembersToProjectRequest
	GetInstanceId() *string
	SetMemberIds(v string) *AttachMembersToProjectRequest
	GetMemberIds() *string
	SetProjectId(v int64) *AttachMembersToProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *AttachMembersToProjectRequest
	GetRegionId() *string
}

type AttachMembersToProjectRequest struct {
	// The Bastionhost instance ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UIDs of the member accounts to associate with the project.
	//
	// > You can call the ListMembersNotForProject operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 161****6121
	MemberIds *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
	// The ID of the project to which you want to associate member accounts.
	//
	// > You can call the ListProjects operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachMembersToProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachMembersToProjectRequest) GoString() string {
	return s.String()
}

func (s *AttachMembersToProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachMembersToProjectRequest) GetMemberIds() *string {
	return s.MemberIds
}

func (s *AttachMembersToProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AttachMembersToProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachMembersToProjectRequest) SetInstanceId(v string) *AttachMembersToProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachMembersToProjectRequest) SetMemberIds(v string) *AttachMembersToProjectRequest {
	s.MemberIds = &v
	return s
}

func (s *AttachMembersToProjectRequest) SetProjectId(v int64) *AttachMembersToProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *AttachMembersToProjectRequest) SetRegionId(v string) *AttachMembersToProjectRequest {
	s.RegionId = &v
	return s
}

func (s *AttachMembersToProjectRequest) Validate() error {
	return dara.Validate(s)
}
