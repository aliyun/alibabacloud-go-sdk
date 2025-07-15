// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UninstallAppResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UninstallAppResponseBody
	GetTaskId() *string
}

type UninstallAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E5138F7E-46B5-526A-8C99-82DEAE6B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallAppResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallAppResponseBody) SetRequestId(v string) *UninstallAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAppResponseBody) SetTaskId(v string) *UninstallAppResponseBody {
	s.TaskId = &v
	return s
}

func (s *UninstallAppResponseBody) Validate() error {
	return dara.Validate(s)
}
