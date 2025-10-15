// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *ListApplicationsForGroupRequest
	GetApplicationIds() []*string
	SetGroupId(v string) *ListApplicationsForGroupRequest
	GetGroupId() *string
	SetInstanceId(v string) *ListApplicationsForGroupRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListApplicationsForGroupRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListApplicationsForGroupRequest
	GetPageSize() *int64
}

type ListApplicationsForGroupRequest struct {
	// 非必填，如果填写则可以基于应用ID进行过滤，列表中最多包含100个元素。
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// 组的唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 当前查询的列表页码，默认为1。
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 当前查询的列表页码，默认为20。
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationsForGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForGroupRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *ListApplicationsForGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListApplicationsForGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForGroupRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListApplicationsForGroupRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationsForGroupRequest) SetApplicationIds(v []*string) *ListApplicationsForGroupRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsForGroupRequest) SetGroupId(v string) *ListApplicationsForGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ListApplicationsForGroupRequest) SetInstanceId(v string) *ListApplicationsForGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForGroupRequest) SetPageNumber(v int64) *ListApplicationsForGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsForGroupRequest) SetPageSize(v int64) *ListApplicationsForGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsForGroupRequest) Validate() error {
	return dara.Validate(s)
}
