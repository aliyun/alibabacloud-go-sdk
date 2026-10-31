// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBatchTemplateResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBatchTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBatchTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBatchTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBatchTemplateResponseBody
	GetSuccess() *bool
}

type UpdateBatchTemplateResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
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
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBatchTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBatchTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBatchTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBatchTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBatchTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBatchTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBatchTemplateResponseBody) SetCode(v string) *UpdateBatchTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBatchTemplateResponseBody) SetHttpStatusCode(v int32) *UpdateBatchTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBatchTemplateResponseBody) SetMessage(v string) *UpdateBatchTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBatchTemplateResponseBody) SetRequestId(v string) *UpdateBatchTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBatchTemplateResponseBody) SetSuccess(v bool) *UpdateBatchTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBatchTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
