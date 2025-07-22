// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSqlLogTaskRequest
	GetInstanceId() *string
	SetPageNo(v int32) *DescribeSqlLogTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeSqlLogTaskRequest
	GetPageSize() *int32
	SetTaskId(v string) *DescribeSqlLogTaskRequest
	GetTaskId() *string
}

type DescribeSqlLogTaskRequest struct {
	// The ID of the database instance.
	//
	// example:
	//
	// r-bp1nti25tc7bq5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID.
	//
	// example:
	//
	// a4f5c4494dbd6713185d87a97aa53e8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSqlLogTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlLogTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeSqlLogTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSqlLogTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSqlLogTaskRequest) SetInstanceId(v string) *DescribeSqlLogTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlLogTaskRequest) SetPageNo(v int32) *DescribeSqlLogTaskRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeSqlLogTaskRequest) SetPageSize(v int32) *DescribeSqlLogTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlLogTaskRequest) SetTaskId(v string) *DescribeSqlLogTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSqlLogTaskRequest) Validate() error {
	return dara.Validate(s)
}
