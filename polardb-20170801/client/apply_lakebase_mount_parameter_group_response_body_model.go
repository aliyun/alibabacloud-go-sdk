// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyLakebaseMountParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyLakebaseMountParameterGroupResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *ApplyLakebaseMountParameterGroupResponseBody
	GetTaskId() *int64
}

type ApplyLakebaseMountParameterGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The asynchronous task ID.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApplyLakebaseMountParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyLakebaseMountParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyLakebaseMountParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyLakebaseMountParameterGroupResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ApplyLakebaseMountParameterGroupResponseBody) SetRequestId(v string) *ApplyLakebaseMountParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyLakebaseMountParameterGroupResponseBody) SetTaskId(v int64) *ApplyLakebaseMountParameterGroupResponseBody {
	s.TaskId = &v
	return s
}

func (s *ApplyLakebaseMountParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
