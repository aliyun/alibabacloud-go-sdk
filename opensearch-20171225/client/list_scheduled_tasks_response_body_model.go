// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListScheduledTasksResponseBody
	GetRequestId() *string
	SetResult(v []map[string]interface{}) *ListScheduledTasksResponseBody
	GetResult() []map[string]interface{}
	SetTotalCount(v int64) *ListScheduledTasksResponseBody
	GetTotalCount() *int64
}

type ListScheduledTasksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListScheduledTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledTasksResponseBody) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *ListScheduledTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListScheduledTasksResponseBody) SetRequestId(v string) *ListScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetResult(v []map[string]interface{}) *ListScheduledTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListScheduledTasksResponseBody) SetTotalCount(v int64) *ListScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduledTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
