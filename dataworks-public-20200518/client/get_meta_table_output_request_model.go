// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *GetMetaTableOutputRequest
	GetEndDate() *string
	SetPageNumber(v int32) *GetMetaTableOutputRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMetaTableOutputRequest
	GetPageSize() *int32
	SetStartDate(v string) *GetMetaTableOutputRequest
	GetStartDate() *string
	SetTableGuid(v string) *GetMetaTableOutputRequest
	GetTableGuid() *string
	SetTaskId(v string) *GetMetaTableOutputRequest
	GetTaskId() *string
}

type GetMetaTableOutputRequest struct {
	// The end date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-15
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number. Valid values: 1 to 30. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-02
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The GUID of the metatable.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.sample_project.sample_table
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1048576
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetMetaTableOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableOutputRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableOutputRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetMetaTableOutputRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTableOutputRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableOutputRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetMetaTableOutputRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableOutputRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMetaTableOutputRequest) SetEndDate(v string) *GetMetaTableOutputRequest {
	s.EndDate = &v
	return s
}

func (s *GetMetaTableOutputRequest) SetPageNumber(v int32) *GetMetaTableOutputRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTableOutputRequest) SetPageSize(v int32) *GetMetaTableOutputRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableOutputRequest) SetStartDate(v string) *GetMetaTableOutputRequest {
	s.StartDate = &v
	return s
}

func (s *GetMetaTableOutputRequest) SetTableGuid(v string) *GetMetaTableOutputRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableOutputRequest) SetTaskId(v string) *GetMetaTableOutputRequest {
	s.TaskId = &v
	return s
}

func (s *GetMetaTableOutputRequest) Validate() error {
	return dara.Validate(s)
}
