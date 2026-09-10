// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteDataCheckTemplateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteDataCheckTemplateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteDataCheckTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataCheckTemplateResponseBody
	GetSuccess() *bool
}

type DeleteDataCheckTemplateResponseBody struct {
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
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, use errCode and errMessage to troubleshoot the issue.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteDataCheckTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckTemplateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteDataCheckTemplateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteDataCheckTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataCheckTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataCheckTemplateResponseBody) SetErrCode(v string) *DeleteDataCheckTemplateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDataCheckTemplateResponseBody) SetErrMessage(v string) *DeleteDataCheckTemplateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDataCheckTemplateResponseBody) SetRequestId(v string) *DeleteDataCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataCheckTemplateResponseBody) SetSuccess(v bool) *DeleteDataCheckTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataCheckTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
