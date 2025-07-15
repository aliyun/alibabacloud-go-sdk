// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceImageResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateInstanceImageResponseBody
	GetTaskId() *string
}

type UpdateInstanceImageResponseBody struct {
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateInstanceImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateInstanceImageResponseBody) SetRequestId(v string) *UpdateInstanceImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceImageResponseBody) SetTaskId(v string) *UpdateInstanceImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateInstanceImageResponseBody) Validate() error {
	return dara.Validate(s)
}
