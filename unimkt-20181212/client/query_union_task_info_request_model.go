// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnionTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryUnionTaskInfoRequest
	GetChannelId() *string
	SetProxyUserId(v int64) *QueryUnionTaskInfoRequest
	GetProxyUserId() *int64
	SetTaskId(v int64) *QueryUnionTaskInfoRequest
	GetTaskId() *int64
}

type QueryUnionTaskInfoRequest struct {
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryUnionTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUnionTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryUnionTaskInfoRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryUnionTaskInfoRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *QueryUnionTaskInfoRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryUnionTaskInfoRequest) SetChannelId(v string) *QueryUnionTaskInfoRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryUnionTaskInfoRequest) SetProxyUserId(v int64) *QueryUnionTaskInfoRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryUnionTaskInfoRequest) SetTaskId(v int64) *QueryUnionTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *QueryUnionTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
