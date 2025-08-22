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
}

type CreateJobResponseBody struct {
	// example:
	//
	// job-518855d9a058c32798c319561f
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
