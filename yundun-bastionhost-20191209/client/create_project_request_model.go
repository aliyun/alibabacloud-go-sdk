// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProjectRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateProjectRequest
	GetInstanceId() *string
	SetMemberIds(v string) *CreateProjectRequest
	GetMemberIds() *string
	SetName(v string) *CreateProjectRequest
	GetName() *string
	SetRegionId(v string) *CreateProjectRequest
	GetRegionId() *string
}

type CreateProjectRequest struct {
	// The description of the project.
	//
	// example:
	//
	// comment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zvp282aly06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The UIDs of the member accounts to associate with the project.
	//
	// > You can call the ListMembersNotForProject operation to obtain this parameter.
	//
	// example:
	//
	// 161****6121
	MemberIds *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateProjectRequest) GetMemberIds() *string {
	return s.MemberIds
}

func (s *CreateProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetInstanceId(v string) *CreateProjectRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateProjectRequest) SetMemberIds(v string) *CreateProjectRequest {
	s.MemberIds = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetRegionId(v string) *CreateProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	return dara.Validate(s)
}
