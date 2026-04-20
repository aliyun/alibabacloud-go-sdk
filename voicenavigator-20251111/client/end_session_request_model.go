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
  // example:
  // 
  // 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // 0a88e269-01f5-49ac-97af-5830f0ccb271
  ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 07d72ea0-6e38-48d4-8371-14deaaba996f
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

