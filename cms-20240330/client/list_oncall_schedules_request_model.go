// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOncallSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleIds(v []*string) *ListOncallSchedulesRequest
	GetOncallScheduleIds() []*string
	SetOncallScheduleName(v string) *ListOncallSchedulesRequest
	GetOncallScheduleName() *string
	SetPageNumber(v int32) *ListOncallSchedulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOncallSchedulesRequest
	GetPageSize() *int32
	SetWorkspace(v string) *ListOncallSchedulesRequest
	GetWorkspace() *string
}

type ListOncallSchedulesRequest struct {
	// The IDs of the on-call schedules.
	OncallScheduleIds []*string `json:"oncallScheduleIds,omitempty" xml:"oncallScheduleIds,omitempty" type:"Repeated"`
	// The name of the on-call schedule.
	//
	// example:
	//
	// testOncallScheduleName
	OncallScheduleName *string `json:"oncallScheduleName,omitempty" xml:"oncallScheduleName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListOncallSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesRequest) GetOncallScheduleIds() []*string {
	return s.OncallScheduleIds
}

func (s *ListOncallSchedulesRequest) GetOncallScheduleName() *string {
	return s.OncallScheduleName
}

func (s *ListOncallSchedulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOncallSchedulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOncallSchedulesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListOncallSchedulesRequest) SetOncallScheduleIds(v []*string) *ListOncallSchedulesRequest {
	s.OncallScheduleIds = v
	return s
}

func (s *ListOncallSchedulesRequest) SetOncallScheduleName(v string) *ListOncallSchedulesRequest {
	s.OncallScheduleName = &v
	return s
}

func (s *ListOncallSchedulesRequest) SetPageNumber(v int32) *ListOncallSchedulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOncallSchedulesRequest) SetPageSize(v int32) *ListOncallSchedulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOncallSchedulesRequest) SetWorkspace(v string) *ListOncallSchedulesRequest {
	s.Workspace = &v
	return s
}

func (s *ListOncallSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
