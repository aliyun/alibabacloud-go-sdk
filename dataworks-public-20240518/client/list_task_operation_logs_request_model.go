// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskOperationLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v int64) *ListTaskOperationLogsRequest
	GetDate() *int64
	SetId(v int64) *ListTaskOperationLogsRequest
	GetId() *int64
	SetPageNumber(v int32) *ListTaskOperationLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTaskOperationLogsRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListTaskOperationLogsRequest
	GetProjectEnv() *string
}

type ListTaskOperationLogsRequest struct {
	// The operation date, accurate to the day. The default value is the current day. You can query only the operation logs generated within the previous 31 days.
	//
	// example:
	//
	// 1710239005403
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s ListTaskOperationLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskOperationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskOperationLogsRequest) GetDate() *int64 {
	return s.Date
}

func (s *ListTaskOperationLogsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListTaskOperationLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskOperationLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskOperationLogsRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTaskOperationLogsRequest) SetDate(v int64) *ListTaskOperationLogsRequest {
	s.Date = &v
	return s
}

func (s *ListTaskOperationLogsRequest) SetId(v int64) *ListTaskOperationLogsRequest {
	s.Id = &v
	return s
}

func (s *ListTaskOperationLogsRequest) SetPageNumber(v int32) *ListTaskOperationLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskOperationLogsRequest) SetPageSize(v int32) *ListTaskOperationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskOperationLogsRequest) SetProjectEnv(v string) *ListTaskOperationLogsRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListTaskOperationLogsRequest) Validate() error {
	return dara.Validate(s)
}
