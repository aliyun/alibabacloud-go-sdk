// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFunctionTaskResponseBody
	GetCode() *string
	SetHttpCode(v int64) *DeleteFunctionTaskResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *DeleteFunctionTaskResponseBody
	GetLatency() *int64
	SetMessage(v string) *DeleteFunctionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFunctionTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteFunctionTaskResponseBody
	GetStatus() *string
}

type DeleteFunctionTaskResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// Task.UnableDelete
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// "1081EB05-473C-5BF4-95BE-6D7CAD5E2213"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFunctionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFunctionTaskResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *DeleteFunctionTaskResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *DeleteFunctionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFunctionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFunctionTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteFunctionTaskResponseBody) SetCode(v string) *DeleteFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetHttpCode(v int64) *DeleteFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetLatency(v int64) *DeleteFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetMessage(v string) *DeleteFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetRequestId(v string) *DeleteFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetStatus(v string) *DeleteFunctionTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
