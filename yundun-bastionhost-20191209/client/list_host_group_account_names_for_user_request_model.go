// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupAccountNamesForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *ListHostGroupAccountNamesForUserRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *ListHostGroupAccountNamesForUserRequest
	GetInstanceId() *string
	SetProjectId(v int64) *ListHostGroupAccountNamesForUserRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListHostGroupAccountNamesForUserRequest
	GetRegionId() *string
	SetUserId(v string) *ListHostGroupAccountNamesForUserRequest
	GetUserId() *string
}

type ListHostGroupAccountNamesForUserRequest struct {
	// The ID of the host group to query.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The instance ID of the Bastionhost instance to which the user belongs.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the user belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to query.
	//
	// > You can call the [ListUsers](https://help.aliyun.com/document_detail/204522.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListHostGroupAccountNamesForUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostGroupAccountNamesForUserRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListHostGroupAccountNamesForUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostGroupAccountNamesForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListHostGroupAccountNamesForUserRequest) SetHostGroupId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetInstanceId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetProjectId(v int64) *ListHostGroupAccountNamesForUserRequest {
	s.ProjectId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetRegionId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetUserId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) Validate() error {
	return dara.Validate(s)
}
