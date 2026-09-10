// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteDataCheckTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteDataCheckTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteDataCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataCheckTaskResponseBody
	GetSuccess() *bool
}

type DeleteDataCheckTaskResponseBody struct {
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
	// The request ID, which is used to locate and troubleshoot issues of this call.
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDataCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDataCheckTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteDataCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataCheckTaskResponseBody) SetErrCode(v string) *DeleteDataCheckTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDataCheckTaskResponseBody) SetErrMessage(v string) *DeleteDataCheckTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDataCheckTaskResponseBody) SetRequestId(v string) *DeleteDataCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataCheckTaskResponseBody) SetSuccess(v bool) *DeleteDataCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
