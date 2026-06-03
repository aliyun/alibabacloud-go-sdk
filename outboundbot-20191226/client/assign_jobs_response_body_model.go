// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssignJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AssignJobsResponseBody
	GetHttpStatusCode() *int32
	SetJobGroupId(v string) *AssignJobsResponseBody
	GetJobGroupId() *string
	SetJobsId(v []*string) *AssignJobsResponseBody
	GetJobsId() []*string
	SetMessage(v string) *AssignJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssignJobsResponseBody
	GetSuccess() *bool
}

type AssignJobsResponseBody struct {
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
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
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
}

func (s AssignJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsResponseBody) GoString() string {
	return s.String()
}

func (s *AssignJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssignJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AssignJobsResponseBody) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *AssignJobsResponseBody) GetJobsId() []*string {
	return s.JobsId
}

func (s *AssignJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssignJobsResponseBody) SetCode(v string) *AssignJobsResponseBody {
	s.Code = &v
	return s
}

func (s *AssignJobsResponseBody) SetHttpStatusCode(v int32) *AssignJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AssignJobsResponseBody) SetJobGroupId(v string) *AssignJobsResponseBody {
	s.JobGroupId = &v
	return s
}

func (s *AssignJobsResponseBody) SetJobsId(v []*string) *AssignJobsResponseBody {
	s.JobsId = v
	return s
}

func (s *AssignJobsResponseBody) SetMessage(v string) *AssignJobsResponseBody {
	s.Message = &v
	return s
}

func (s *AssignJobsResponseBody) SetRequestId(v string) *AssignJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignJobsResponseBody) SetSuccess(v bool) *AssignJobsResponseBody {
	s.Success = &v
	return s
}

func (s *AssignJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
