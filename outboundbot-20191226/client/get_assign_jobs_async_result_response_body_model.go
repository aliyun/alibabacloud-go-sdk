// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssignJobsAsyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAssignJobsAsyncResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAssignJobsAsyncResultResponseBody
	GetHttpStatusCode() *int32
	SetJobGroupId(v string) *GetAssignJobsAsyncResultResponseBody
	GetJobGroupId() *string
	SetJobsId(v []*string) *GetAssignJobsAsyncResultResponseBody
	GetJobsId() []*string
	SetMessage(v string) *GetAssignJobsAsyncResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAssignJobsAsyncResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAssignJobsAsyncResultResponseBody
	GetSuccess() *bool
	SetTimeout(v bool) *GetAssignJobsAsyncResultResponseBody
	GetTimeout() *bool
	SetValid(v bool) *GetAssignJobsAsyncResultResponseBody
	GetValid() *bool
}

type GetAssignJobsAsyncResultResponseBody struct {
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
	// d3550dd1-360d-4fe5-b4b2-667a4a664dab
	JobGroupId *string   `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobsId     []*string `json:"JobsId,omitempty" xml:"JobsId,omitempty" type:"Repeated"`
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
	// example:
	//
	// false
	Timeout *bool `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetAssignJobsAsyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssignJobsAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssignJobsAsyncResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAssignJobsAsyncResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAssignJobsAsyncResultResponseBody) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *GetAssignJobsAsyncResultResponseBody) GetJobsId() []*string {
	return s.JobsId
}

func (s *GetAssignJobsAsyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAssignJobsAsyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssignJobsAsyncResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAssignJobsAsyncResultResponseBody) GetTimeout() *bool {
	return s.Timeout
}

func (s *GetAssignJobsAsyncResultResponseBody) GetValid() *bool {
	return s.Valid
}

func (s *GetAssignJobsAsyncResultResponseBody) SetCode(v string) *GetAssignJobsAsyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetHttpStatusCode(v int32) *GetAssignJobsAsyncResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetJobGroupId(v string) *GetAssignJobsAsyncResultResponseBody {
	s.JobGroupId = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetJobsId(v []*string) *GetAssignJobsAsyncResultResponseBody {
	s.JobsId = v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetMessage(v string) *GetAssignJobsAsyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetRequestId(v string) *GetAssignJobsAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetSuccess(v bool) *GetAssignJobsAsyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetTimeout(v bool) *GetAssignJobsAsyncResultResponseBody {
	s.Timeout = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) SetValid(v bool) *GetAssignJobsAsyncResultResponseBody {
	s.Valid = &v
	return s
}

func (s *GetAssignJobsAsyncResultResponseBody) Validate() error {
	return dara.Validate(s)
}
