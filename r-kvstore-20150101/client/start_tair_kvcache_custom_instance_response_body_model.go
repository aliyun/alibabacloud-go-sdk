// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTairKVCacheCustomInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartTairKVCacheCustomInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StartTairKVCacheCustomInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartTairKVCacheCustomInstanceResponseBody
	GetTaskId() *string
}

type StartTairKVCacheCustomInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartTairKVCacheCustomInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTairKVCacheCustomInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartTairKVCacheCustomInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartTairKVCacheCustomInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTairKVCacheCustomInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartTairKVCacheCustomInstanceResponseBody) SetInstanceId(v string) *StartTairKVCacheCustomInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceResponseBody) SetRequestId(v string) *StartTairKVCacheCustomInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceResponseBody) SetTaskId(v string) *StartTairKVCacheCustomInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartTairKVCacheCustomInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
