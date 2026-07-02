// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartTairKVCacheCustomInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RestartTairKVCacheCustomInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *RestartTairKVCacheCustomInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RestartTairKVCacheCustomInstanceResponseBody
	GetTaskId() *string
}

type RestartTairKVCacheCustomInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartTairKVCacheCustomInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartTairKVCacheCustomInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) SetInstanceId(v string) *RestartTairKVCacheCustomInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) SetRequestId(v string) *RestartTairKVCacheCustomInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) SetTaskId(v string) *RestartTairKVCacheCustomInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartTairKVCacheCustomInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
