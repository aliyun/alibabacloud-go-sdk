// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAcpCompletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *GenerateAcpCompletionRequest
	GetAgentName() *string
	SetId(v string) *GenerateAcpCompletionRequest
	GetId() *string
	SetJsonrpc(v string) *GenerateAcpCompletionRequest
	GetJsonrpc() *string
	SetMethod(v string) *GenerateAcpCompletionRequest
	GetMethod() *string
	SetParams(v map[string]interface{}) *GenerateAcpCompletionRequest
	GetParams() map[string]interface{}
}

type GenerateAcpCompletionRequest struct {
	AgentName *string                `json:"agentName,omitempty" xml:"agentName,omitempty"`
	Id        *string                `json:"id,omitempty" xml:"id,omitempty"`
	Jsonrpc   *string                `json:"jsonrpc,omitempty" xml:"jsonrpc,omitempty"`
	Method    *string                `json:"method,omitempty" xml:"method,omitempty"`
	Params    map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
}

func (s GenerateAcpCompletionRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAcpCompletionRequest) GoString() string {
	return s.String()
}

func (s *GenerateAcpCompletionRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *GenerateAcpCompletionRequest) GetId() *string {
	return s.Id
}

func (s *GenerateAcpCompletionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GenerateAcpCompletionRequest) GetMethod() *string {
	return s.Method
}

func (s *GenerateAcpCompletionRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *GenerateAcpCompletionRequest) SetAgentName(v string) *GenerateAcpCompletionRequest {
	s.AgentName = &v
	return s
}

func (s *GenerateAcpCompletionRequest) SetId(v string) *GenerateAcpCompletionRequest {
	s.Id = &v
	return s
}

func (s *GenerateAcpCompletionRequest) SetJsonrpc(v string) *GenerateAcpCompletionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GenerateAcpCompletionRequest) SetMethod(v string) *GenerateAcpCompletionRequest {
	s.Method = &v
	return s
}

func (s *GenerateAcpCompletionRequest) SetParams(v map[string]interface{}) *GenerateAcpCompletionRequest {
	s.Params = v
	return s
}

func (s *GenerateAcpCompletionRequest) Validate() error {
	return dara.Validate(s)
}
