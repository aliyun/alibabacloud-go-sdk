// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicImageJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicImageJob(v *SubmitDynamicImageJobResponseBodyDynamicImageJob) *SubmitDynamicImageJobResponseBody
	GetDynamicImageJob() *SubmitDynamicImageJobResponseBodyDynamicImageJob
	SetRequestId(v string) *SubmitDynamicImageJobResponseBody
	GetRequestId() *string
}

type SubmitDynamicImageJobResponseBody struct {
	// The information about the animated image job.
	DynamicImageJob *SubmitDynamicImageJobResponseBodyDynamicImageJob `json:"DynamicImageJob,omitempty" xml:"DynamicImageJob,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-74A6-BEF6-D7393642****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDynamicImageJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponseBody) GetDynamicImageJob() *SubmitDynamicImageJobResponseBodyDynamicImageJob {
	return s.DynamicImageJob
}

func (s *SubmitDynamicImageJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDynamicImageJobResponseBody) SetDynamicImageJob(v *SubmitDynamicImageJobResponseBodyDynamicImageJob) *SubmitDynamicImageJobResponseBody {
	s.DynamicImageJob = v
	return s
}

func (s *SubmitDynamicImageJobResponseBody) SetRequestId(v string) *SubmitDynamicImageJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDynamicImageJobResponseBody) Validate() error {
	if s.DynamicImageJob != nil {
		if err := s.DynamicImageJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDynamicImageJobResponseBodyDynamicImageJob struct {
	// The ID of the animated image job.
	//
	// example:
	//
	// ad90a501b1bfb72374ad0050746****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitDynamicImageJobResponseBodyDynamicImageJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicImageJobResponseBodyDynamicImageJob) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponseBodyDynamicImageJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitDynamicImageJobResponseBodyDynamicImageJob) SetJobId(v string) *SubmitDynamicImageJobResponseBodyDynamicImageJob {
	s.JobId = &v
	return s
}

func (s *SubmitDynamicImageJobResponseBodyDynamicImageJob) Validate() error {
	return dara.Validate(s)
}
