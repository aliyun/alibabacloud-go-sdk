// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomScriptTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomScriptTemplateResponseBody
	GetCode() *string
	SetData(v string) *CreateCustomScriptTemplateResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateCustomScriptTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCustomScriptTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomScriptTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomScriptTemplateResponseBody
	GetSuccess() *bool
}

type CreateCustomScriptTemplateResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created custom scenario template.
	//
	// example:
	//
	// 8d6a6e41-8093-49af-a9d1-0281878758ac
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomScriptTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomScriptTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomScriptTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomScriptTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateCustomScriptTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCustomScriptTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomScriptTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomScriptTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomScriptTemplateResponseBody) SetCode(v string) *CreateCustomScriptTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomScriptTemplateResponseBody) SetData(v string) *CreateCustomScriptTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCustomScriptTemplateResponseBody) SetHttpStatusCode(v int32) *CreateCustomScriptTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomScriptTemplateResponseBody) SetMessage(v string) *CreateCustomScriptTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomScriptTemplateResponseBody) SetRequestId(v string) *CreateCustomScriptTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomScriptTemplateResponseBody) SetSuccess(v bool) *CreateCustomScriptTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomScriptTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
