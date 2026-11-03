// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteJobResponseBody
	GetCode() *string
	SetData(v string) *DeleteJobResponseBody
	GetData() *string
	SetErrorCode(v string) *DeleteJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteJobResponseBody
	GetTraceId() *string
}

type DeleteJobResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {"msg": "","code": 200,"success": true}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code. Valid values:
	//
	// - Successful request: The **ErrorCode*	- field is not returned.
	//
	// - Failed request: The **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidAppId.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - **success**: The request is successful.
	//
	// - Specific error code: The request is abnormal.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the task template was deleted. Valid values:
	//
	// - **true**: The task template was deleted.
	//
	// - **false**: The task template failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteJobResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteJobResponseBody) SetCode(v string) *DeleteJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteJobResponseBody) SetData(v string) *DeleteJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteJobResponseBody) SetErrorCode(v string) *DeleteJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteJobResponseBody) SetMessage(v string) *DeleteJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobResponseBody) SetSuccess(v bool) *DeleteJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteJobResponseBody) SetTraceId(v string) *DeleteJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteJobResponseBody) Validate() error {
	return dara.Validate(s)
}
