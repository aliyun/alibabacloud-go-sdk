// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDBTaskSQLJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *ListDBTaskSQLJobDetailRequest
	GetJobId() *int64
	SetPageNumber(v int64) *ListDBTaskSQLJobDetailRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDBTaskSQLJobDetailRequest
	GetPageSize() *int64
	SetTid(v int64) *ListDBTaskSQLJobDetailRequest
	GetTid() *int64
}

type ListDBTaskSQLJobDetailRequest struct {
	// The ID of the SQL task. You can call the [ListDBTaskSQLJob](https://help.aliyun.com/document_detail/207049.html) operation to query the SQL task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1276****
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListDBTaskSQLJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDBTaskSQLJobDetailRequest) GoString() string {
	return s.String()
}

func (s *ListDBTaskSQLJobDetailRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDBTaskSQLJobDetailRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDBTaskSQLJobDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDBTaskSQLJobDetailRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDBTaskSQLJobDetailRequest) SetJobId(v int64) *ListDBTaskSQLJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) SetPageNumber(v int64) *ListDBTaskSQLJobDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) SetPageSize(v int64) *ListDBTaskSQLJobDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) SetTid(v int64) *ListDBTaskSQLJobDetailRequest {
	s.Tid = &v
	return s
}

func (s *ListDBTaskSQLJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
