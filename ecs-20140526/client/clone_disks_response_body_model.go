// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloneDisksResponseBody
	GetRequestId() *string
	SetTaskGroupId(v string) *CloneDisksResponseBody
	GetTaskGroupId() *string
}

type CloneDisksResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task group ID of the disk cloning operation. You can call [DescribeTasks](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ecs-2014-05-26-describetasks) to query the task execution result.
	//
	// example:
	//
	// g-2ze2op2grqpclwu7****
	TaskGroupId *string `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
}

func (s CloneDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksResponseBody) GoString() string {
	return s.String()
}

func (s *CloneDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneDisksResponseBody) GetTaskGroupId() *string {
	return s.TaskGroupId
}

func (s *CloneDisksResponseBody) SetRequestId(v string) *CloneDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneDisksResponseBody) SetTaskGroupId(v string) *CloneDisksResponseBody {
	s.TaskGroupId = &v
	return s
}

func (s *CloneDisksResponseBody) Validate() error {
	return dara.Validate(s)
}
