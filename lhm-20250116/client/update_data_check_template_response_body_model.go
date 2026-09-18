// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataCheckTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *UpdateDataCheckTemplateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateDataCheckTemplateResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *UpdateDataCheckTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataCheckTemplateResponseBody
	GetSuccess() *bool
}

type UpdateDataCheckTemplateResponseBody struct {
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
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateDataCheckTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataCheckTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataCheckTemplateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateDataCheckTemplateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateDataCheckTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataCheckTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataCheckTemplateResponseBody) SetErrCode(v string) *UpdateDataCheckTemplateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDataCheckTemplateResponseBody) SetErrMessage(v string) *UpdateDataCheckTemplateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateDataCheckTemplateResponseBody) SetRequestId(v string) *UpdateDataCheckTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataCheckTemplateResponseBody) SetSuccess(v bool) *UpdateDataCheckTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataCheckTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
