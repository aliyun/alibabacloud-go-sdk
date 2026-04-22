// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetError(v interface{}) *LoadAgentSessionResponseBody
	GetError() interface{}
	SetId(v string) *LoadAgentSessionResponseBody
	GetId() *string
	SetJsonrpc(v string) *LoadAgentSessionResponseBody
	GetJsonrpc() *string
	SetMethod(v string) *LoadAgentSessionResponseBody
	GetMethod() *string
	SetParams(v interface{}) *LoadAgentSessionResponseBody
	GetParams() interface{}
	SetRequestId(v string) *LoadAgentSessionResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *LoadAgentSessionResponseBody
	GetResult() interface{}
	SetTimestamp(v int64) *LoadAgentSessionResponseBody
	GetTimestamp() *int64
}

type LoadAgentSessionResponseBody struct {
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 676303114031776
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
	// 0D41C608-0C60-5EB0-B986-1460909CF642
	RequestId *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 1769479322828
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s LoadAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionResponseBody) GetError() interface{} {
	return s.Error
}

func (s *LoadAgentSessionResponseBody) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionResponseBody) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *LoadAgentSessionResponseBody) GetMethod() *string {
	return s.Method
}

func (s *LoadAgentSessionResponseBody) GetParams() interface{} {
	return s.Params
}

func (s *LoadAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoadAgentSessionResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *LoadAgentSessionResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *LoadAgentSessionResponseBody) SetError(v interface{}) *LoadAgentSessionResponseBody {
	s.Error = v
	return s
}

func (s *LoadAgentSessionResponseBody) SetId(v string) *LoadAgentSessionResponseBody {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetJsonrpc(v string) *LoadAgentSessionResponseBody {
	s.Jsonrpc = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetMethod(v string) *LoadAgentSessionResponseBody {
	s.Method = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetParams(v interface{}) *LoadAgentSessionResponseBody {
	s.Params = v
	return s
}

func (s *LoadAgentSessionResponseBody) SetRequestId(v string) *LoadAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetResult(v interface{}) *LoadAgentSessionResponseBody {
	s.Result = v
	return s
}

func (s *LoadAgentSessionResponseBody) SetTimestamp(v int64) *LoadAgentSessionResponseBody {
	s.Timestamp = &v
	return s
}

func (s *LoadAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
