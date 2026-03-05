// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndUnionTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetChannelId(v string) *EndUnionTaskRequest
  GetChannelId() *string 
  SetProxyUserId(v int64) *EndUnionTaskRequest
  GetProxyUserId() *int64 
  SetTaskId(v int64) *EndUnionTaskRequest
  GetTaskId() *int64 
}

type EndUnionTaskRequest struct {
  // This parameter is required.
  ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
  ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
  TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s EndUnionTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s EndUnionTaskRequest) GoString() string {
  return s.String()
}

func (s *EndUnionTaskRequest) GetChannelId() *string  {
  return s.ChannelId
}

func (s *EndUnionTaskRequest) GetProxyUserId() *int64  {
  return s.ProxyUserId
}

func (s *EndUnionTaskRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *EndUnionTaskRequest) SetChannelId(v string) *EndUnionTaskRequest {
  s.ChannelId = &v
  return s
}

func (s *EndUnionTaskRequest) SetProxyUserId(v int64) *EndUnionTaskRequest {
  s.ProxyUserId = &v
  return s
}

func (s *EndUnionTaskRequest) SetTaskId(v int64) *EndUnionTaskRequest {
  s.TaskId = &v
  return s
}

func (s *EndUnionTaskRequest) Validate() error {
  return dara.Validate(s)
}

