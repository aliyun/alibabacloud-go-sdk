// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCircuitBreakerRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteCircuitBreakerRulesResponseBody
	GetCode() *int32
	SetData(v []*int64) *DeleteCircuitBreakerRulesResponseBody
	GetData() []*int64
	SetHttpStatusCode(v int32) *DeleteCircuitBreakerRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCircuitBreakerRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCircuitBreakerRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCircuitBreakerRulesResponseBody
	GetSuccess() *bool
}

type DeleteCircuitBreakerRulesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of the rules that were deleted.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FAF577DD-8E8E-5BE6-80CC-380E19******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCircuitBreakerRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCircuitBreakerRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCircuitBreakerRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteCircuitBreakerRulesResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *DeleteCircuitBreakerRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCircuitBreakerRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCircuitBreakerRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCircuitBreakerRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCircuitBreakerRulesResponseBody) SetCode(v int32) *DeleteCircuitBreakerRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCircuitBreakerRulesResponseBody) SetData(v []*int64) *DeleteCircuitBreakerRulesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCircuitBreakerRulesResponseBody) SetHttpStatusCode(v int32) *DeleteCircuitBreakerRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCircuitBreakerRulesResponseBody) SetMessage(v string) *DeleteCircuitBreakerRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCircuitBreakerRulesResponseBody) SetRequestId(v string) *DeleteCircuitBreakerRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCircuitBreakerRulesResponseBody) SetSuccess(v bool) *DeleteCircuitBreakerRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCircuitBreakerRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
