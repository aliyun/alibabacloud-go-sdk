// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPipelineAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitPipelineAsyncResponseBody
	GetCode() *string
	SetData(v int64) *SubmitPipelineAsyncResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *SubmitPipelineAsyncResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitPipelineAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitPipelineAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitPipelineAsyncResponseBody
	GetSuccess() *bool
}

type SubmitPipelineAsyncResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the submission detail.
	//
	// example:
	//
	// 123
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

func (s SubmitPipelineAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPipelineAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPipelineAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitPipelineAsyncResponseBody) GetData() *int64 {
	return s.Data
}

func (s *SubmitPipelineAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitPipelineAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitPipelineAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPipelineAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitPipelineAsyncResponseBody) SetCode(v string) *SubmitPipelineAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitPipelineAsyncResponseBody) SetData(v int64) *SubmitPipelineAsyncResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitPipelineAsyncResponseBody) SetHttpStatusCode(v int32) *SubmitPipelineAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitPipelineAsyncResponseBody) SetMessage(v string) *SubmitPipelineAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitPipelineAsyncResponseBody) SetRequestId(v string) *SubmitPipelineAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPipelineAsyncResponseBody) SetSuccess(v bool) *SubmitPipelineAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitPipelineAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
