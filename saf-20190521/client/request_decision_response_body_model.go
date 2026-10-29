// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestDecisionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *RequestDecisionResponseBody
	GetCode() *int64
	SetData(v map[string]interface{}) *RequestDecisionResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *RequestDecisionResponseBody
	GetMessage() *string
	SetRequestId(v string) *RequestDecisionResponseBody
	GetRequestId() *string
}

type RequestDecisionResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data information, which includes fields such as tags, score, and extend.
	//
	// example:
	//
	// {"score": 80.4, "tags": "ra0601,rnk0002,rn0312"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The result message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F41DC6-2070-5CE7-8B2D-7A23D882932D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RequestDecisionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RequestDecisionResponseBody) GoString() string {
	return s.String()
}

func (s *RequestDecisionResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RequestDecisionResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *RequestDecisionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RequestDecisionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RequestDecisionResponseBody) SetCode(v int64) *RequestDecisionResponseBody {
	s.Code = &v
	return s
}

func (s *RequestDecisionResponseBody) SetData(v map[string]interface{}) *RequestDecisionResponseBody {
	s.Data = v
	return s
}

func (s *RequestDecisionResponseBody) SetMessage(v string) *RequestDecisionResponseBody {
	s.Message = &v
	return s
}

func (s *RequestDecisionResponseBody) SetRequestId(v string) *RequestDecisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RequestDecisionResponseBody) Validate() error {
	return dara.Validate(s)
}
