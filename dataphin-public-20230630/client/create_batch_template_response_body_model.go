// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateBatchTemplateResponseBody
	GetCode() *string
	SetData(v int64) *CreateBatchTemplateResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateBatchTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBatchTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBatchTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBatchTemplateResponseBody
	GetSuccess() *bool
}

type CreateBatchTemplateResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the offline computing template.
	//
	// example:
	//
	// 1001
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateBatchTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBatchTemplateResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateBatchTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBatchTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBatchTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBatchTemplateResponseBody) SetCode(v string) *CreateBatchTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBatchTemplateResponseBody) SetData(v int64) *CreateBatchTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *CreateBatchTemplateResponseBody) SetHttpStatusCode(v int32) *CreateBatchTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBatchTemplateResponseBody) SetMessage(v string) *CreateBatchTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBatchTemplateResponseBody) SetRequestId(v string) *CreateBatchTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchTemplateResponseBody) SetSuccess(v bool) *CreateBatchTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBatchTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
