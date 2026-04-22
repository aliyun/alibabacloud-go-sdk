// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetError(v interface{}) *PromptAgentSessionResponseBody
	GetError() interface{}
	SetId(v string) *PromptAgentSessionResponseBody
	GetId() *string
	SetJsonrpc(v string) *PromptAgentSessionResponseBody
	GetJsonrpc() *string
	SetMethod(v string) *PromptAgentSessionResponseBody
	GetMethod() *string
	SetParams(v interface{}) *PromptAgentSessionResponseBody
	GetParams() interface{}
	SetRequestId(v string) *PromptAgentSessionResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *PromptAgentSessionResponseBody
	GetResult() interface{}
	SetTimestamp(v int64) *PromptAgentSessionResponseBody
	GetTimestamp() *int64
}

type PromptAgentSessionResponseBody struct {
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string     `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Method  *string     `json:"Method,omitempty" xml:"Method,omitempty"`
	Params  interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D5D70885-7CC7-594A-80C7-2EF1B00FFB4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// end_turn
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 1747447032
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s PromptAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionResponseBody) GetError() interface{} {
	return s.Error
}

func (s *PromptAgentSessionResponseBody) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionResponseBody) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *PromptAgentSessionResponseBody) GetMethod() *string {
	return s.Method
}

func (s *PromptAgentSessionResponseBody) GetParams() interface{} {
	return s.Params
}

func (s *PromptAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PromptAgentSessionResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *PromptAgentSessionResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *PromptAgentSessionResponseBody) SetError(v interface{}) *PromptAgentSessionResponseBody {
	s.Error = v
	return s
}

func (s *PromptAgentSessionResponseBody) SetId(v string) *PromptAgentSessionResponseBody {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetJsonrpc(v string) *PromptAgentSessionResponseBody {
	s.Jsonrpc = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetMethod(v string) *PromptAgentSessionResponseBody {
	s.Method = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetParams(v interface{}) *PromptAgentSessionResponseBody {
	s.Params = v
	return s
}

func (s *PromptAgentSessionResponseBody) SetRequestId(v string) *PromptAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetResult(v interface{}) *PromptAgentSessionResponseBody {
	s.Result = v
	return s
}

func (s *PromptAgentSessionResponseBody) SetTimestamp(v int64) *PromptAgentSessionResponseBody {
	s.Timestamp = &v
	return s
}

func (s *PromptAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
