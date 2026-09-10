// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerReadAsyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetInnerReadAsyncResultResponseBody
	GetData() *string
	SetErrCode(v string) *GetInnerReadAsyncResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetInnerReadAsyncResultResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetInnerReadAsyncResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInnerReadAsyncResultResponseBody
	GetSuccess() *string
}

type GetInnerReadAsyncResultResponseBody struct {
	// The business data returned by the operation in string format. The content varies depending on the operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This value is an empty string if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. This value is an empty string if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID. Use this value to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Use errCode and errMessage to troubleshoot the issue.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInnerReadAsyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInnerReadAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerReadAsyncResultResponseBody) GetData() *string {
	return s.Data
}

func (s *GetInnerReadAsyncResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetInnerReadAsyncResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetInnerReadAsyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInnerReadAsyncResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInnerReadAsyncResultResponseBody) SetData(v string) *GetInnerReadAsyncResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetInnerReadAsyncResultResponseBody) SetErrCode(v string) *GetInnerReadAsyncResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetInnerReadAsyncResultResponseBody) SetErrMessage(v string) *GetInnerReadAsyncResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetInnerReadAsyncResultResponseBody) SetRequestId(v string) *GetInnerReadAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInnerReadAsyncResultResponseBody) SetSuccess(v string) *GetInnerReadAsyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetInnerReadAsyncResultResponseBody) Validate() error {
	return dara.Validate(s)
}
