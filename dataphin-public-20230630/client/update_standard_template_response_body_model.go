// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStandardTemplateResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateStandardTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateStandardTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateStandardTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStandardTemplateResponseBody
	GetSuccess() *bool
}

type UpdateStandardTemplateResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s UpdateStandardTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateStandardTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateStandardTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStandardTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStandardTemplateResponseBody) SetCode(v string) *UpdateStandardTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStandardTemplateResponseBody) SetHttpStatusCode(v int32) *UpdateStandardTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateStandardTemplateResponseBody) SetMessage(v string) *UpdateStandardTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStandardTemplateResponseBody) SetRequestId(v string) *UpdateStandardTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStandardTemplateResponseBody) SetSuccess(v bool) *UpdateStandardTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStandardTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
