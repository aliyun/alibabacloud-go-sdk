// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRtcMPUTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListRtcMPUTaskDetailRequest
	GetAppId() *string
	SetPageNo(v int64) *ListRtcMPUTaskDetailRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListRtcMPUTaskDetailRequest
	GetPageSize() *int64
	SetTaskId(v string) *ListRtcMPUTaskDetailRequest
	GetTaskId() *string
}

type ListRtcMPUTaskDetailRequest struct {
	// The application ID.
	//
	// > The application ID consists of uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The page number.
	//
	// > If you do not specify a task ID, you must specify the PageSize and PageNo parameters. In this case, the paged query results of all stream mixing and relaying tasks under the specified application ID are returned.
	//
	// example:
	//
	// 20
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. Valid values: 1 to 100.
	//
	// > If you do not specify a task ID, you must specify the PageSize and PageNo parameters. In this case, the paged query results of all stream mixing and relaying tasks under the specified application ID are returned.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID.
	//
	//
	// >- The task ID consists of uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 55 characters.
	//
	// - If you specify a task ID, the query is performed based on the task ID first, and the result contains the parameter details of the stream mixing and relaying task with the specified task ID.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListRtcMPUTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListRtcMPUTaskDetailRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListRtcMPUTaskDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRtcMPUTaskDetailRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListRtcMPUTaskDetailRequest) SetAppId(v string) *ListRtcMPUTaskDetailRequest {
	s.AppId = &v
	return s
}

func (s *ListRtcMPUTaskDetailRequest) SetPageNo(v int64) *ListRtcMPUTaskDetailRequest {
	s.PageNo = &v
	return s
}

func (s *ListRtcMPUTaskDetailRequest) SetPageSize(v int64) *ListRtcMPUTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListRtcMPUTaskDetailRequest) SetTaskId(v string) *ListRtcMPUTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *ListRtcMPUTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
