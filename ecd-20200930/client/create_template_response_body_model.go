// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTemplateResponseBody
	GetCode() *string
	SetData(v string) *CreateTemplateResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTemplateResponseBody
	GetSuccess() *bool
}

type CreateTemplateResponseBody struct {
	// The execution result of the operation. If the request was successful, `success` is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The template ID.
	//
	// example:
	//
	// b-0cc7rx533*****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The creation result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 791CC0D3-1A38-573B-8F5F-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTemplateResponseBody) SetCode(v string) *CreateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTemplateResponseBody) SetData(v string) *CreateTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTemplateResponseBody) SetHttpStatusCode(v int32) *CreateTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTemplateResponseBody) SetMessage(v string) *CreateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetSuccess(v bool) *CreateTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
