// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInnerConvertAsyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetInnerConvertAsyncResultResponseBody
	GetData() *string
	SetErrCode(v string) *GetInnerConvertAsyncResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetInnerConvertAsyncResultResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *GetInnerConvertAsyncResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetInnerConvertAsyncResultResponseBody
	GetSuccess() *string
}

type GetInnerConvertAsyncResultResponseBody struct {
	// The business data returned by the operation in string format. The specific content varies by operation.
	//
	// example:
	//
	// demo
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The request ID that uniquely identifies the call. Provide this value when troubleshooting issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Troubleshoot the issue based on errCode and errMessage.
	//
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInnerConvertAsyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInnerConvertAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetInnerConvertAsyncResultResponseBody) GetData() *string {
	return s.Data
}

func (s *GetInnerConvertAsyncResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetInnerConvertAsyncResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetInnerConvertAsyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInnerConvertAsyncResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetInnerConvertAsyncResultResponseBody) SetData(v string) *GetInnerConvertAsyncResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetInnerConvertAsyncResultResponseBody) SetErrCode(v string) *GetInnerConvertAsyncResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetInnerConvertAsyncResultResponseBody) SetErrMessage(v string) *GetInnerConvertAsyncResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetInnerConvertAsyncResultResponseBody) SetRequestId(v string) *GetInnerConvertAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInnerConvertAsyncResultResponseBody) SetSuccess(v string) *GetInnerConvertAsyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetInnerConvertAsyncResultResponseBody) Validate() error {
	return dara.Validate(s)
}
