// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ResumeInstanceResponseBody
	GetData() *bool
	SetErrorCode(v string) *ResumeInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ResumeInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *ResumeInstanceResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ResumeInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeInstanceResponseBody
	GetSuccess() *bool
}

type ResumeInstanceResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) GetData() *bool {
	return s.Data
}

func (s *ResumeInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResumeInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResumeInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ResumeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeInstanceResponseBody) SetData(v bool) *ResumeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorCode(v string) *ResumeInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorMessage(v string) *ResumeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetHttpStatusCode(v string) *ResumeInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetSuccess(v bool) *ResumeInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
