// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushExpireKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *FlushExpireKeysResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *FlushExpireKeysResponseBody
	GetRequestId() *string
	SetTaskId(v string) *FlushExpireKeysResponseBody
	GetTaskId() *string
}

type FlushExpireKeysResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 82E30AB7-E3A4-46AC-88A0-3E4DCDC5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 21986****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FlushExpireKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlushExpireKeysResponseBody) GoString() string {
	return s.String()
}

func (s *FlushExpireKeysResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FlushExpireKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlushExpireKeysResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *FlushExpireKeysResponseBody) SetInstanceId(v string) *FlushExpireKeysResponseBody {
	s.InstanceId = &v
	return s
}

func (s *FlushExpireKeysResponseBody) SetRequestId(v string) *FlushExpireKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlushExpireKeysResponseBody) SetTaskId(v string) *FlushExpireKeysResponseBody {
	s.TaskId = &v
	return s
}

func (s *FlushExpireKeysResponseBody) Validate() error {
	return dara.Validate(s)
}
