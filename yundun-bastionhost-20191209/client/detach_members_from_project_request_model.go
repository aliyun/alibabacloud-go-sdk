// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachMembersFromProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DetachMembersFromProjectRequest
	GetInstanceId() *string
	SetMemberIds(v string) *DetachMembersFromProjectRequest
	GetMemberIds() *string
	SetProjectId(v int64) *DetachMembersFromProjectRequest
	GetProjectId() *int64
	SetRegionId(v string) *DetachMembersFromProjectRequest
	GetRegionId() *string
}

type DetachMembersFromProjectRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UIDs of the member accounts to remove.
	//
	// > You can call the GetProject operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 161****6121
	MemberIds *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
	// The project ID.
	//
	// > You can call the ListProjects operation to query this parameter.
	//
	// This parameter is required.
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
}

func (s DetachMembersFromProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachMembersFromProjectRequest) GoString() string {
	return s.String()
}

func (s *DetachMembersFromProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachMembersFromProjectRequest) GetMemberIds() *string {
	return s.MemberIds
}

func (s *DetachMembersFromProjectRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DetachMembersFromProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachMembersFromProjectRequest) SetInstanceId(v string) *DetachMembersFromProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachMembersFromProjectRequest) SetMemberIds(v string) *DetachMembersFromProjectRequest {
	s.MemberIds = &v
	return s
}

func (s *DetachMembersFromProjectRequest) SetProjectId(v int64) *DetachMembersFromProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DetachMembersFromProjectRequest) SetRegionId(v string) *DetachMembersFromProjectRequest {
	s.RegionId = &v
	return s
}

func (s *DetachMembersFromProjectRequest) Validate() error {
	return dara.Validate(s)
}
