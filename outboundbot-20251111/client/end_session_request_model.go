// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndSessionRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EndSessionRequest
  GetInstanceId() *string 
  SetScriptId(v string) *EndSessionRequest
  GetScriptId() *string 
  SetSessionId(v string) *EndSessionRequest
  GetSessionId() *string 
}

type EndSessionRequest struct {
  // The instance identity of the associated product. You must use an existing instance that the invoker is authorized to access. The instance must match the instance to which the scenario, channel, and session belong. The maximum length is subject to the current POP parameter definitions.
  // 
  // example:
  // 
  // instance-example-001
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The ID of the dialog scenario under the instance. You must specify an existing scenario. Subsequent dialog and session termination requests for the same text session must use the same scenario that was used to start the session.
  // 
  // example:
  // 
  // script-example-001
  ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
  // The session ID used to uniquely identify the current interaction. For a text session, use the text- prefixed ID returned when the session is started. For web voice, use the ws- prefixed ID. For outbound call tests with real numbers, use the test- prefixed ID. Subsequent calls must use the actual returned values.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // text-11111111-2222-4333-8444-555555555555
  SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EndSessionRequest) String() string {
  return dara.Prettify(s)
}

func (s EndSessionRequest) GoString() string {
  return s.String()
}

func (s *EndSessionRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EndSessionRequest) GetScriptId() *string  {
  return s.ScriptId
}

func (s *EndSessionRequest) GetSessionId() *string  {
  return s.SessionId
}

func (s *EndSessionRequest) SetInstanceId(v string) *EndSessionRequest {
  s.InstanceId = &v
  return s
}

func (s *EndSessionRequest) SetScriptId(v string) *EndSessionRequest {
  s.ScriptId = &v
  return s
}

func (s *EndSessionRequest) SetSessionId(v string) *EndSessionRequest {
  s.SessionId = &v
  return s
}

func (s *EndSessionRequest) Validate() error {
  return dara.Validate(s)
}

