// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTaskDialoguesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSearchTaskDialoguesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSearchTaskDialoguesRequest
	GetPageSize() *int32
	SetTaskId(v string) *ListSearchTaskDialoguesRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *ListSearchTaskDialoguesRequest
	GetWorkspaceId() *string
}

type ListSearchTaskDialoguesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListSearchTaskDialoguesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTaskDialoguesRequest) GoString() string {
	return s.String()
}

func (s *ListSearchTaskDialoguesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSearchTaskDialoguesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSearchTaskDialoguesRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListSearchTaskDialoguesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSearchTaskDialoguesRequest) SetPageNumber(v int32) *ListSearchTaskDialoguesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSearchTaskDialoguesRequest) SetPageSize(v int32) *ListSearchTaskDialoguesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSearchTaskDialoguesRequest) SetTaskId(v string) *ListSearchTaskDialoguesRequest {
	s.TaskId = &v
	return s
}

func (s *ListSearchTaskDialoguesRequest) SetWorkspaceId(v string) *ListSearchTaskDialoguesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListSearchTaskDialoguesRequest) Validate() error {
	return dara.Validate(s)
}
