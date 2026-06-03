// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAcpCommandRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAgentName(v string) *ExecuteAcpCommandRequest
  GetAgentName() *string 
  SetId(v string) *ExecuteAcpCommandRequest
  GetId() *string 
  SetJsonrpc(v string) *ExecuteAcpCommandRequest
  GetJsonrpc() *string 
  SetMethod(v string) *ExecuteAcpCommandRequest
  GetMethod() *string 
  SetParams(v map[string]interface{}) *ExecuteAcpCommandRequest
  GetParams() map[string]interface{} 
}

type ExecuteAcpCommandRequest struct {
  AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
  Id *string `json:"id,omitempty" xml:"id,omitempty"`
  Jsonrpc *string `json:"jsonrpc,omitempty" xml:"jsonrpc,omitempty"`
  Method *string `json:"method,omitempty" xml:"method,omitempty"`
  Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ExecuteAcpCommandRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAcpCommandRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAcpCommandRequest) GetAgentName() *string  {
  return s.AgentName
}

func (s *ExecuteAcpCommandRequest) GetId() *string  {
  return s.Id
}

func (s *ExecuteAcpCommandRequest) GetJsonrpc() *string  {
  return s.Jsonrpc
}

func (s *ExecuteAcpCommandRequest) GetMethod() *string  {
  return s.Method
}

func (s *ExecuteAcpCommandRequest) GetParams() map[string]interface{}  {
  return s.Params
}

func (s *ExecuteAcpCommandRequest) SetAgentName(v string) *ExecuteAcpCommandRequest {
  s.AgentName = &v
  return s
}

func (s *ExecuteAcpCommandRequest) SetId(v string) *ExecuteAcpCommandRequest {
  s.Id = &v
  return s
}

func (s *ExecuteAcpCommandRequest) SetJsonrpc(v string) *ExecuteAcpCommandRequest {
  s.Jsonrpc = &v
  return s
}

func (s *ExecuteAcpCommandRequest) SetMethod(v string) *ExecuteAcpCommandRequest {
  s.Method = &v
  return s
}

func (s *ExecuteAcpCommandRequest) SetParams(v map[string]interface{}) *ExecuteAcpCommandRequest {
  s.Params = v
  return s
}

func (s *ExecuteAcpCommandRequest) Validate() error {
  return dara.Validate(s)
}

