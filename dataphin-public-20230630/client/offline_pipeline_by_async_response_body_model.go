// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineByAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OfflinePipelineByAsyncResponseBody
	GetCode() *string
	SetData(v int64) *OfflinePipelineByAsyncResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *OfflinePipelineByAsyncResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OfflinePipelineByAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflinePipelineByAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflinePipelineByAsyncResponseBody
	GetSuccess() *bool
}

type OfflinePipelineByAsyncResponseBody struct {
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

func (s OfflinePipelineByAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineByAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *OfflinePipelineByAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflinePipelineByAsyncResponseBody) GetData() *int64 {
	return s.Data
}

func (s *OfflinePipelineByAsyncResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OfflinePipelineByAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflinePipelineByAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflinePipelineByAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflinePipelineByAsyncResponseBody) SetCode(v string) *OfflinePipelineByAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *OfflinePipelineByAsyncResponseBody) SetData(v int64) *OfflinePipelineByAsyncResponseBody {
	s.Data = &v
	return s
}

func (s *OfflinePipelineByAsyncResponseBody) SetHttpStatusCode(v int32) *OfflinePipelineByAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflinePipelineByAsyncResponseBody) SetMessage(v string) *OfflinePipelineByAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *OfflinePipelineByAsyncResponseBody) SetRequestId(v string) *OfflinePipelineByAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflinePipelineByAsyncResponseBody) SetSuccess(v bool) *OfflinePipelineByAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *OfflinePipelineByAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
