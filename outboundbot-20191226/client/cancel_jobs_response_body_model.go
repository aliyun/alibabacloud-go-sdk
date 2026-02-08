// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CancelJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelJobsResponseBody
	GetSuccess() *bool
}

type CancelJobsResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CancelJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelJobsResponseBody) SetCode(v string) *CancelJobsResponseBody {
	s.Code = &v
	return s
}

func (s *CancelJobsResponseBody) SetHttpStatusCode(v int32) *CancelJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelJobsResponseBody) SetMessage(v string) *CancelJobsResponseBody {
	s.Message = &v
	return s
}

func (s *CancelJobsResponseBody) SetRequestId(v string) *CancelJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelJobsResponseBody) SetSuccess(v bool) *CancelJobsResponseBody {
	s.Success = &v
	return s
}

func (s *CancelJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
