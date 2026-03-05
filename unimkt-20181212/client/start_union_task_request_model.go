// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartUnionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *StartUnionTaskRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *StartUnionTaskRequest
	GetProxyUserId() *int64
	SetTaskId(v int64) *StartUnionTaskRequest
	GetTaskId() *int64
}

type StartUnionTaskRequest struct {
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ProxyUserId *int64  `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartUnionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartUnionTaskRequest) GoString() string {
	return s.String()
}

func (s *StartUnionTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartUnionTaskRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *StartUnionTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StartUnionTaskRequest) SetChannelId(v string) *StartUnionTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StartUnionTaskRequest) SetProxyUserId(v int64) *StartUnionTaskRequest {
	s.ProxyUserId = &v
	return s
}

func (s *StartUnionTaskRequest) SetTaskId(v int64) *StartUnionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartUnionTaskRequest) Validate() error {
	return dara.Validate(s)
}
