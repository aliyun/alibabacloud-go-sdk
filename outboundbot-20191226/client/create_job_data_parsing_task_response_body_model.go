// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobDataParsingTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateJobDataParsingTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateJobDataParsingTaskResponseBody
	GetHttpStatusCode() *int32
	SetJobDataParsingTaskId(v string) *CreateJobDataParsingTaskResponseBody
	GetJobDataParsingTaskId() *string
	SetMessage(v string) *CreateJobDataParsingTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateJobDataParsingTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateJobDataParsingTaskResponseBody
	GetSuccess() *bool
}

type CreateJobDataParsingTaskResponseBody struct {
	// API status code
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
	// The ID of the uploaded job.
	//
	// > This parameter will be used as the JobDataParsingTaskId parameter in the DescribeJobDataParsingTaskProgress API.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// API message
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
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobDataParsingTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobDataParsingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobDataParsingTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateJobDataParsingTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateJobDataParsingTaskResponseBody) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *CreateJobDataParsingTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateJobDataParsingTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobDataParsingTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateJobDataParsingTaskResponseBody) SetCode(v string) *CreateJobDataParsingTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobDataParsingTaskResponseBody) SetHttpStatusCode(v int32) *CreateJobDataParsingTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateJobDataParsingTaskResponseBody) SetJobDataParsingTaskId(v string) *CreateJobDataParsingTaskResponseBody {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *CreateJobDataParsingTaskResponseBody) SetMessage(v string) *CreateJobDataParsingTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobDataParsingTaskResponseBody) SetRequestId(v string) *CreateJobDataParsingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobDataParsingTaskResponseBody) SetSuccess(v bool) *CreateJobDataParsingTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobDataParsingTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
