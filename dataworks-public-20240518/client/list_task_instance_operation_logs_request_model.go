// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstanceOperationLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v int64) *ListTaskInstanceOperationLogsRequest
	GetDate() *int64
	SetId(v int64) *ListTaskInstanceOperationLogsRequest
	GetId() *int64
	SetPageNumber(v int32) *ListTaskInstanceOperationLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTaskInstanceOperationLogsRequest
	GetPageSize() *int32
}

type ListTaskInstanceOperationLogsRequest struct {
	// The operation date, accurate to the day. The default value is the current day. You can query only the operation logs generated within the previous 31 days. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1710239005403
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The instance ID.
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
}

func (s ListTaskInstanceOperationLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstanceOperationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskInstanceOperationLogsRequest) GetDate() *int64 {
	return s.Date
}

func (s *ListTaskInstanceOperationLogsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListTaskInstanceOperationLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTaskInstanceOperationLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskInstanceOperationLogsRequest) SetDate(v int64) *ListTaskInstanceOperationLogsRequest {
	s.Date = &v
	return s
}

func (s *ListTaskInstanceOperationLogsRequest) SetId(v int64) *ListTaskInstanceOperationLogsRequest {
	s.Id = &v
	return s
}

func (s *ListTaskInstanceOperationLogsRequest) SetPageNumber(v int32) *ListTaskInstanceOperationLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTaskInstanceOperationLogsRequest) SetPageSize(v int32) *ListTaskInstanceOperationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskInstanceOperationLogsRequest) Validate() error {
	return dara.Validate(s)
}
