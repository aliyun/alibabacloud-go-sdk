// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStandardTemplateResponseBody
	GetCode() *string
	SetData(v int64) *CreateStandardTemplateResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateStandardTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateStandardTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStandardTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStandardTemplateResponseBody
	GetSuccess() *bool
}

type CreateStandardTemplateResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 22
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStandardTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStandardTemplateResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateStandardTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateStandardTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStandardTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStandardTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStandardTemplateResponseBody) SetCode(v string) *CreateStandardTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetData(v int64) *CreateStandardTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetHttpStatusCode(v int32) *CreateStandardTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetMessage(v string) *CreateStandardTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetRequestId(v string) *CreateStandardTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) SetSuccess(v bool) *CreateStandardTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStandardTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
