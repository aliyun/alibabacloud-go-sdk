// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobResponseBody
	GetCode() *int32
	SetDetails(v string) *GetJobResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetJobResponseBody
	GetErrorCode() *string
	SetJob(v *Job) *GetJobResponseBody
	GetJob() *Job
	SetMessage(v string) *GetJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobResponseBody
	GetSuccess() *bool
}

type GetJobResponseBody struct {
	// Total amount of data under the conditions of this request. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Job
	Job *Job `json:"Job,omitempty" xml:"Job,omitempty"`
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobResponseBody) GetJob() *Job {
	return s.Job
}

func (s *GetJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobResponseBody) SetCode(v int32) *GetJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobResponseBody) SetDetails(v string) *GetJobResponseBody {
	s.Details = &v
	return s
}

func (s *GetJobResponseBody) SetErrorCode(v string) *GetJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResponseBody) SetJob(v *Job) *GetJobResponseBody {
	s.Job = v
	return s
}

func (s *GetJobResponseBody) SetMessage(v string) *GetJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v bool) *GetJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}
