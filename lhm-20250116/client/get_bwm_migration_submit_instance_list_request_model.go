// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationSubmitInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *GetBwmMigrationSubmitInstanceListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetBwmMigrationSubmitInstanceListRequest
	GetPageSize() *int32
	SetStatus(v int32) *GetBwmMigrationSubmitInstanceListRequest
	GetStatus() *int32
	SetTaskId(v string) *GetBwmMigrationSubmitInstanceListRequest
	GetTaskId() *string
}

type GetBwmMigrationSubmitInstanceListRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The instance status.
	//
	// example:
	//
	// 7
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The conversion task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetBwmMigrationSubmitInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationSubmitInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationSubmitInstanceListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetBwmMigrationSubmitInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetBwmMigrationSubmitInstanceListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetBwmMigrationSubmitInstanceListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBwmMigrationSubmitInstanceListRequest) SetPageIndex(v int32) *GetBwmMigrationSubmitInstanceListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListRequest) SetPageSize(v int32) *GetBwmMigrationSubmitInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListRequest) SetStatus(v int32) *GetBwmMigrationSubmitInstanceListRequest {
	s.Status = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListRequest) SetTaskId(v string) *GetBwmMigrationSubmitInstanceListRequest {
	s.TaskId = &v
	return s
}

func (s *GetBwmMigrationSubmitInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
