// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAgentRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBaseMeAgentId(v int64) *ExecuteAgentRequest
  GetBaseMeAgentId() *int64 
  SetJsonStr(v string) *ExecuteAgentRequest
  GetJsonStr() *string 
  SetStream(v bool) *ExecuteAgentRequest
  GetStream() *bool 
}

type ExecuteAgentRequest struct {
  // The ID of the business workspace.
  // 
  // example:
  // 
  // 123456
  BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
  // The complete JSON string. For more information, see the following detailed description.
  // 
  // example:
  // 
  // ""
  JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
  // Specifies whether to enable Server-Sent Events (SSE) responses. Set to true to enable SSE responses. Default value: false.
  // 
  // example:
  // 
  // false
  Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s ExecuteAgentRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAgentRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAgentRequest) GetBaseMeAgentId() *int64  {
  return s.BaseMeAgentId
}

func (s *ExecuteAgentRequest) GetJsonStr() *string  {
  return s.JsonStr
}

func (s *ExecuteAgentRequest) GetStream() *bool  {
  return s.Stream
}

func (s *ExecuteAgentRequest) SetBaseMeAgentId(v int64) *ExecuteAgentRequest {
  s.BaseMeAgentId = &v
  return s
}

func (s *ExecuteAgentRequest) SetJsonStr(v string) *ExecuteAgentRequest {
  s.JsonStr = &v
  return s
}

func (s *ExecuteAgentRequest) SetStream(v bool) *ExecuteAgentRequest {
  s.Stream = &v
  return s
}

func (s *ExecuteAgentRequest) Validate() error {
  return dara.Validate(s)
}

