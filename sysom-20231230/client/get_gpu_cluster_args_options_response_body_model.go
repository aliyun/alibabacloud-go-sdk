// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterArgsOptionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGpuClusterArgsOptionsResponseBody
	GetCode() *string
	SetData(v interface{}) *GetGpuClusterArgsOptionsResponseBody
	GetData() interface{}
	SetMessage(v string) *GetGpuClusterArgsOptionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGpuClusterArgsOptionsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *GetGpuClusterArgsOptionsResponseBody
	GetTotal() *int32
}

type GetGpuClusterArgsOptionsResponseBody struct {
	// The response code. Valid values: Success and Sysom.ServerError.
	//
	// example:
	//
	// Success  Fail
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result of the parameter retrieval.
	//
	// example:
	//
	// ["instance1", "instance2"]
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The error message returned when the code is not Success.
	//
	// example:
	//
	// some error msg while fail，“” if success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// return real request id
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetGpuClusterArgsOptionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterArgsOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGpuClusterArgsOptionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGpuClusterArgsOptionsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetGpuClusterArgsOptionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGpuClusterArgsOptionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGpuClusterArgsOptionsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetGpuClusterArgsOptionsResponseBody) SetCode(v string) *GetGpuClusterArgsOptionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetGpuClusterArgsOptionsResponseBody) SetData(v interface{}) *GetGpuClusterArgsOptionsResponseBody {
	s.Data = v
	return s
}

func (s *GetGpuClusterArgsOptionsResponseBody) SetMessage(v string) *GetGpuClusterArgsOptionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetGpuClusterArgsOptionsResponseBody) SetRequestId(v string) *GetGpuClusterArgsOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGpuClusterArgsOptionsResponseBody) SetTotal(v int32) *GetGpuClusterArgsOptionsResponseBody {
	s.Total = &v
	return s
}

func (s *GetGpuClusterArgsOptionsResponseBody) Validate() error {
	return dara.Validate(s)
}
