// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetTraceAppConfigResponseBody
	GetCode() *int64
	SetData(v string) *GetTraceAppConfigResponseBody
	GetData() *string
	SetMessage(v string) *GetTraceAppConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTraceAppConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTraceAppConfigResponseBody
	GetSuccess() *bool
}

type GetTraceAppConfigResponseBody struct {
	// The HTTP status code.
	//
	// Valid values:
	//
	// 	- 2xx: The request was successful.
	//
	// 	- 3xx: The request was redirected.
	//
	// 	- 4xx: A request error occurred.
	//
	// 	- 5xx: A server error occurred.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The custom settings of the application, which is a JSON string. For more information, see **Additional description of response parameters**.
	//
	// example:
	//
	// {"profiler":{"enable":true}}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTraceAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAppConfigResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetTraceAppConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *GetTraceAppConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTraceAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceAppConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTraceAppConfigResponseBody) SetCode(v int64) *GetTraceAppConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetTraceAppConfigResponseBody) SetData(v string) *GetTraceAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetTraceAppConfigResponseBody) SetMessage(v string) *GetTraceAppConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetTraceAppConfigResponseBody) SetRequestId(v string) *GetTraceAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceAppConfigResponseBody) SetSuccess(v bool) *GetTraceAppConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetTraceAppConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
