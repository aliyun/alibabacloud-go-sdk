// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitBatchJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitBatchJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitBatchJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitBatchJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitBatchJobsResponseBody
	GetSuccess() *bool
}

type SubmitBatchJobsResponseBody struct {
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
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitBatchJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBatchJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitBatchJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitBatchJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitBatchJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBatchJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitBatchJobsResponseBody) SetCode(v string) *SubmitBatchJobsResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitBatchJobsResponseBody) SetHttpStatusCode(v int32) *SubmitBatchJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitBatchJobsResponseBody) SetMessage(v string) *SubmitBatchJobsResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitBatchJobsResponseBody) SetRequestId(v string) *SubmitBatchJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBatchJobsResponseBody) SetSuccess(v bool) *SubmitBatchJobsResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitBatchJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
