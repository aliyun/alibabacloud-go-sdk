// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakebaseMountParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLakebaseMountParameterGroupResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *ModifyLakebaseMountParameterGroupResponseBody
	GetTaskId() *int64
}

type ModifyLakebaseMountParameterGroupResponseBody struct {
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

func (s ModifyLakebaseMountParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakebaseMountParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLakebaseMountParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLakebaseMountParameterGroupResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyLakebaseMountParameterGroupResponseBody) SetRequestId(v string) *ModifyLakebaseMountParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupResponseBody) SetTaskId(v int64) *ModifyLakebaseMountParameterGroupResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyLakebaseMountParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
