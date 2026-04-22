// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *CancelAgentSessionResponseBodyJsonRpcResponse) *CancelAgentSessionResponseBody
	GetJsonRpcResponse() *CancelAgentSessionResponseBodyJsonRpcResponse
	SetRequestId(v string) *CancelAgentSessionResponseBody
	GetRequestId() *string
}

type CancelAgentSessionResponseBody struct {
	JsonRpcResponse *CancelAgentSessionResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 7CD3D216-5876-5DB1-A34A-396806F4A413
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionResponseBody) GetJsonRpcResponse() *CancelAgentSessionResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *CancelAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAgentSessionResponseBody) SetJsonRpcResponse(v *CancelAgentSessionResponseBodyJsonRpcResponse) *CancelAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *CancelAgentSessionResponseBody) SetRequestId(v string) *CancelAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAgentSessionResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelAgentSessionResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 7675839888324361477
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                              `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *CancelAgentSessionResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CancelAgentSessionResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) GetResult() *CancelAgentSessionResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) SetId(v string) *CancelAgentSessionResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) SetJsonrpc(v string) *CancelAgentSessionResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) SetResult(v *CancelAgentSessionResponseBodyJsonRpcResponseResult) *CancelAgentSessionResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelAgentSessionResponseBodyJsonRpcResponseResult struct {
	// example:
	//
	// session-d5d549fe4c2c4180a9814fb74190f502
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CancelAgentSessionResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponseResult) GetSessionId() *string {
	return s.SessionId
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponseResult) SetSessionId(v string) *CancelAgentSessionResponseBodyJsonRpcResponseResult {
	s.SessionId = &v
	return s
}

func (s *CancelAgentSessionResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
