// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iA2aRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *A2aRequest
	GetEnv() *string
	SetId(v string) *A2aRequest
	GetId() *string
	SetJsonrpc(v string) *A2aRequest
	GetJsonrpc() *string
	SetMethod(v string) *A2aRequest
	GetMethod() *string
	SetParams(v interface{}) *A2aRequest
	GetParams() interface{}
}

type A2aRequest struct {
	// example:
	//
	// a2a
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 791
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"jsonrpc,omitempty" xml:"jsonrpc,omitempty"`
	// example:
	//
	// message/stream
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// example:
	//
	// {
	//
	//     "message": {
	//
	//       "role": "user",
	//
	//       "parts": [
	//
	//         {
	//
	//           "kind": "text",
	//
	//           "text": "你好“
	//
	//         }
	//
	//       ]
	//
	//       "kind": "message"
	//
	//     }
	//
	//   }
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
}

func (s A2aRequest) String() string {
	return dara.Prettify(s)
}

func (s A2aRequest) GoString() string {
	return s.String()
}

func (s *A2aRequest) GetEnv() *string {
	return s.Env
}

func (s *A2aRequest) GetId() *string {
	return s.Id
}

func (s *A2aRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *A2aRequest) GetMethod() *string {
	return s.Method
}

func (s *A2aRequest) GetParams() interface{} {
	return s.Params
}

func (s *A2aRequest) SetEnv(v string) *A2aRequest {
	s.Env = &v
	return s
}

func (s *A2aRequest) SetId(v string) *A2aRequest {
	s.Id = &v
	return s
}

func (s *A2aRequest) SetJsonrpc(v string) *A2aRequest {
	s.Jsonrpc = &v
	return s
}

func (s *A2aRequest) SetMethod(v string) *A2aRequest {
	s.Method = &v
	return s
}

func (s *A2aRequest) SetParams(v interface{}) *A2aRequest {
	s.Params = v
	return s
}

func (s *A2aRequest) Validate() error {
	return dara.Validate(s)
}
