// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *CreateJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateJobResponseBody
	GetSuccess() *string
}

type CreateJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// Submitted batch job 10\\n
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A0A38A38-1565-555E-B597-E48A2E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) SetSuccess(v string) *CreateJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
