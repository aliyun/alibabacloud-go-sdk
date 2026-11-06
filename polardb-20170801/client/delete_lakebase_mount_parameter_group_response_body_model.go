// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakebaseMountParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLakebaseMountParameterGroupResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *DeleteLakebaseMountParameterGroupResponseBody
	GetTaskId() *int64
}

type DeleteLakebaseMountParameterGroupResponseBody struct {
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

func (s DeleteLakebaseMountParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakebaseMountParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLakebaseMountParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLakebaseMountParameterGroupResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DeleteLakebaseMountParameterGroupResponseBody) SetRequestId(v string) *DeleteLakebaseMountParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLakebaseMountParameterGroupResponseBody) SetTaskId(v int64) *DeleteLakebaseMountParameterGroupResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteLakebaseMountParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
