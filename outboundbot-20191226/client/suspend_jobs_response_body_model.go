// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SuspendJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SuspendJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendJobsResponseBody
	GetSuccess() *bool
}

type SuspendJobsResponseBody struct {
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

func (s SuspendJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SuspendJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendJobsResponseBody) SetCode(v string) *SuspendJobsResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendJobsResponseBody) SetHttpStatusCode(v int32) *SuspendJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendJobsResponseBody) SetMessage(v string) *SuspendJobsResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendJobsResponseBody) SetRequestId(v string) *SuspendJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendJobsResponseBody) SetSuccess(v bool) *SuspendJobsResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
