// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *Template) *GetTemplateResponseBody
	GetData() *Template
	SetErrorCode(v string) *GetTemplateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTemplateResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetTemplateResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTemplateResponseBody
	GetSuccess() *bool
}

type GetTemplateResponseBody struct {
	// The returned data.
	Data *Template `json:"data,omitempty" xml:"data,omitempty"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// 040003
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// InvalidUser.NotFound
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 484D9DDA-300D-525E-AF7A-0CCCA5C64A7A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetData() *Template {
	return s.Data
}

func (s *GetTemplateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTemplateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTemplateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTemplateResponseBody) SetData(v *Template) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetErrorCode(v string) *GetTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetErrorMessage(v string) *GetTemplateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTemplateResponseBody) SetHttpStatusCode(v string) *GetTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSuccess(v bool) *GetTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
