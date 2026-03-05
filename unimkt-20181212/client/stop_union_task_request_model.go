// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopUnionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *StopUnionTaskRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *StopUnionTaskRequest
	GetProxyUserId() *int64
	SetTaskId(v int64) *StopUnionTaskRequest
	GetTaskId() *int64
}

type StopUnionTaskRequest struct {
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ProxyUserId *int64  `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopUnionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopUnionTaskRequest) GoString() string {
	return s.String()
}

func (s *StopUnionTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopUnionTaskRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *StopUnionTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StopUnionTaskRequest) SetChannelId(v string) *StopUnionTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StopUnionTaskRequest) SetProxyUserId(v int64) *StopUnionTaskRequest {
	s.ProxyUserId = &v
	return s
}

func (s *StopUnionTaskRequest) SetTaskId(v int64) *StopUnionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopUnionTaskRequest) Validate() error {
	return dara.Validate(s)
}
