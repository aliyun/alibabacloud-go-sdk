// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineByAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePipelineByAsyncResponseBody
	GetCode() *string
	SetData(v int64) *UpdatePipelineByAsyncResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdatePipelineByAsyncResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdatePipelineByAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePipelineByAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePipelineByAsyncResponseBody
	GetSuccess() *bool
}

type UpdatePipelineByAsyncResponseBody struct {
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

func (s UpdatePipelineByAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineByAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineByAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePipelineByAsyncResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdatePipelineByAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePipelineByAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePipelineByAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineByAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePipelineByAsyncResponseBody) SetCode(v string) *UpdatePipelineByAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePipelineByAsyncResponseBody) SetData(v int64) *UpdatePipelineByAsyncResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePipelineByAsyncResponseBody) SetHttpStatusCode(v int32) *UpdatePipelineByAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePipelineByAsyncResponseBody) SetMessage(v string) *UpdatePipelineByAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePipelineByAsyncResponseBody) SetRequestId(v string) *UpdatePipelineByAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineByAsyncResponseBody) SetSuccess(v bool) *UpdatePipelineByAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePipelineByAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
