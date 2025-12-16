// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFunctionInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFunctionInstanceResponseBody
	GetCode() *string
	SetHttpCode(v int64) *DeleteFunctionInstanceResponseBody
	GetHttpCode() *int64
	SetLatency(v int64) *DeleteFunctionInstanceResponseBody
	GetLatency() *int64
	SetMessage(v string) *DeleteFunctionInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFunctionInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteFunctionInstanceResponseBody
	GetStatus() *string
}

type DeleteFunctionInstanceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// "Instance.NotExist"
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
	// 10
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// "instance not exist."
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// "1081EB05-473C-5BF4-95BE-6D7CAD5E2213"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// 	- OK: The request is successful.
	//
	// 	- FAIL: The request fails.
	//
	// example:
	//
	// "OK"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFunctionInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFunctionInstanceResponseBody) GetHttpCode() *int64 {
	return s.HttpCode
}

func (s *DeleteFunctionInstanceResponseBody) GetLatency() *int64 {
	return s.Latency
}

func (s *DeleteFunctionInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFunctionInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFunctionInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteFunctionInstanceResponseBody) SetCode(v string) *DeleteFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetHttpCode(v int64) *DeleteFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetLatency(v int64) *DeleteFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetMessage(v string) *DeleteFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetRequestId(v string) *DeleteFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetStatus(v string) *DeleteFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
