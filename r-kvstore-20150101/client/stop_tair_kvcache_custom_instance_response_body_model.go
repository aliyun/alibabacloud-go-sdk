// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTairKVCacheCustomInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopTairKVCacheCustomInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StopTairKVCacheCustomInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopTairKVCacheCustomInstanceResponseBody
	GetTaskId() *string
}

type StopTairKVCacheCustomInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopTairKVCacheCustomInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopTairKVCacheCustomInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopTairKVCacheCustomInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopTairKVCacheCustomInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopTairKVCacheCustomInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopTairKVCacheCustomInstanceResponseBody) SetInstanceId(v string) *StopTairKVCacheCustomInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceResponseBody) SetRequestId(v string) *StopTairKVCacheCustomInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceResponseBody) SetTaskId(v string) *StopTairKVCacheCustomInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopTairKVCacheCustomInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
