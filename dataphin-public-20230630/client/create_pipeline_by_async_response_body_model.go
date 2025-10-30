// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineByAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePipelineByAsyncResponseBody
	GetCode() *string
	SetData(v int64) *CreatePipelineByAsyncResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreatePipelineByAsyncResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePipelineByAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePipelineByAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePipelineByAsyncResponseBody
	GetSuccess() *bool
}

type CreatePipelineByAsyncResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 123
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

func (s CreatePipelineByAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineByAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineByAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePipelineByAsyncResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreatePipelineByAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePipelineByAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePipelineByAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineByAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePipelineByAsyncResponseBody) SetCode(v string) *CreatePipelineByAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePipelineByAsyncResponseBody) SetData(v int64) *CreatePipelineByAsyncResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePipelineByAsyncResponseBody) SetHttpStatusCode(v int32) *CreatePipelineByAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePipelineByAsyncResponseBody) SetMessage(v string) *CreatePipelineByAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePipelineByAsyncResponseBody) SetRequestId(v string) *CreatePipelineByAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineByAsyncResponseBody) SetSuccess(v bool) *CreatePipelineByAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePipelineByAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
