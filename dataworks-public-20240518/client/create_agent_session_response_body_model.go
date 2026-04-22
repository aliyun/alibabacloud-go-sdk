// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *CreateAgentSessionResponseBodyJsonRpcResponse) *CreateAgentSessionResponseBody
	GetJsonRpcResponse() *CreateAgentSessionResponseBodyJsonRpcResponse
	SetRequestId(v string) *CreateAgentSessionResponseBody
	GetRequestId() *string
}

type CreateAgentSessionResponseBody struct {
	JsonRpcResponse *CreateAgentSessionResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 8A9D5E6C-5817-5837-9715-6E3967EC6123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBody) GetJsonRpcResponse() *CreateAgentSessionResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *CreateAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSessionResponseBody) SetJsonRpcResponse(v *CreateAgentSessionResponseBodyJsonRpcResponse) *CreateAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *CreateAgentSessionResponseBody) SetRequestId(v string) *CreateAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSessionResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSessionResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                              `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *CreateAgentSessionResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CreateAgentSessionResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetResult() *CreateAgentSessionResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetId(v string) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetJsonrpc(v string) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetResult(v *CreateAgentSessionResponseBodyJsonRpcResponseResult) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSessionResponseBodyJsonRpcResponseResult struct {
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CreateAgentSessionResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseResult) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseResult) SetSessionId(v string) *CreateAgentSessionResponseBodyJsonRpcResponseResult {
	s.SessionId = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
