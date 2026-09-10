// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteDataCheckConfigResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteDataCheckConfigResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteDataCheckConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataCheckConfigResponseBody
	GetSuccess() *bool
}

type DeleteDataCheckConfigResponseBody struct {
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
	// The request ID. You can use this ID to locate and troubleshoot issues.
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

func (s DeleteDataCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDataCheckConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteDataCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataCheckConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataCheckConfigResponseBody) SetErrCode(v string) *DeleteDataCheckConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDataCheckConfigResponseBody) SetErrMessage(v string) *DeleteDataCheckConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDataCheckConfigResponseBody) SetRequestId(v string) *DeleteDataCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataCheckConfigResponseBody) SetSuccess(v bool) *DeleteDataCheckConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataCheckConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
