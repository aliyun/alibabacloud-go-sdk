// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakebaseMountParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterGroupId(v string) *CreateLakebaseMountParameterGroupResponseBody
	GetParameterGroupId() *string
	SetRequestId(v string) *CreateLakebaseMountParameterGroupResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *CreateLakebaseMountParameterGroupResponseBody
	GetTaskId() *int64
}

type CreateLakebaseMountParameterGroupResponseBody struct {
	// The ID of the created parameter template.
	//
	// example:
	//
	// pcpg-lakebase-xxx
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
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

func (s CreateLakebaseMountParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLakebaseMountParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLakebaseMountParameterGroupResponseBody) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *CreateLakebaseMountParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLakebaseMountParameterGroupResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateLakebaseMountParameterGroupResponseBody) SetParameterGroupId(v string) *CreateLakebaseMountParameterGroupResponseBody {
	s.ParameterGroupId = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupResponseBody) SetRequestId(v string) *CreateLakebaseMountParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupResponseBody) SetTaskId(v int64) *CreateLakebaseMountParameterGroupResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateLakebaseMountParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
