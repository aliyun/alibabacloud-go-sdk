// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *UpdateDataCheckTaskResponseBody
	GetData() *int64
	SetErrCode(v string) *UpdateDataCheckTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateDataCheckTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *UpdateDataCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataCheckTaskResponseBody
	GetSuccess() *bool
}

type UpdateDataCheckTaskResponseBody struct {
	// The business data value returned by the operation. The specific meaning varies by operation.
	//
	// example:
	//
	// 100
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
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
	// The request ID. This value is used to locate and troubleshoot issues with the call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDataCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTaskResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateDataCheckTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateDataCheckTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateDataCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataCheckTaskResponseBody) SetData(v int64) *UpdateDataCheckTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDataCheckTaskResponseBody) SetErrCode(v string) *UpdateDataCheckTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDataCheckTaskResponseBody) SetErrMessage(v string) *UpdateDataCheckTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateDataCheckTaskResponseBody) SetRequestId(v string) *UpdateDataCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataCheckTaskResponseBody) SetSuccess(v bool) *UpdateDataCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
