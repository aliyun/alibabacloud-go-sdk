// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ResumeJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResumeJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResumeJobsResponseBody
	GetSuccess() *bool
}

type ResumeJobsResponseBody struct {
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

func (s ResumeJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResumeJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeJobsResponseBody) SetCode(v string) *ResumeJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeJobsResponseBody) SetHttpStatusCode(v int32) *ResumeJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeJobsResponseBody) SetMessage(v string) *ResumeJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeJobsResponseBody) SetRequestId(v string) *ResumeJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeJobsResponseBody) SetSuccess(v bool) *ResumeJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
