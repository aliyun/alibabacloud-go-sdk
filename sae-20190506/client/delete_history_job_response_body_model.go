// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHistoryJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHistoryJobResponseBody
	GetCode() *string
	SetData(v string) *DeleteHistoryJobResponseBody
	GetData() *string
	SetErrorCode(v string) *DeleteHistoryJobResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteHistoryJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHistoryJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHistoryJobResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteHistoryJobResponseBody
	GetTraceId() *string
}

type DeleteHistoryJobResponseBody struct {
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
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidAppId.NotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
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
	// Indicates whether the job is deleted successfully. Valid values:
	//
	// - **true**: The job is deleted.
	//
	// - **false**: The deletion failed.
	//
	// [_single.error.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query invocation details.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteHistoryJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHistoryJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHistoryJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHistoryJobResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteHistoryJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteHistoryJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHistoryJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHistoryJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHistoryJobResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteHistoryJobResponseBody) SetCode(v string) *DeleteHistoryJobResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) SetData(v string) *DeleteHistoryJobResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) SetErrorCode(v string) *DeleteHistoryJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) SetMessage(v string) *DeleteHistoryJobResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) SetRequestId(v string) *DeleteHistoryJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) SetSuccess(v bool) *DeleteHistoryJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) SetTraceId(v string) *DeleteHistoryJobResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteHistoryJobResponseBody) Validate() error {
	return dara.Validate(s)
}
