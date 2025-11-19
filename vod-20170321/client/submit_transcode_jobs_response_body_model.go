// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTranscodeJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitTranscodeJobsResponseBody
	GetRequestId() *string
	SetTranscodeJobs(v *SubmitTranscodeJobsResponseBodyTranscodeJobs) *SubmitTranscodeJobsResponseBody
	GetTranscodeJobs() *SubmitTranscodeJobsResponseBodyTranscodeJobs
	SetTranscodeTaskId(v string) *SubmitTranscodeJobsResponseBody
	GetTranscodeTaskId() *string
}

type SubmitTranscodeJobsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the transcoding job.
	//
	// >  This parameter is not returned for HLS packaging tasks. You must asynchronously receive the transcoding result.
	TranscodeJobs *SubmitTranscodeJobsResponseBodyTranscodeJobs `json:"TranscodeJobs,omitempty" xml:"TranscodeJobs,omitempty" type:"Struct"`
	// The ID of the transcoding task that was submitted.
	//
	// example:
	//
	// 9f4a0df7da2c8a81c8c0408c84****
	TranscodeTaskId *string `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
}

func (s SubmitTranscodeJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTranscodeJobsResponseBody) GetTranscodeJobs() *SubmitTranscodeJobsResponseBodyTranscodeJobs {
	return s.TranscodeJobs
}

func (s *SubmitTranscodeJobsResponseBody) GetTranscodeTaskId() *string {
	return s.TranscodeTaskId
}

func (s *SubmitTranscodeJobsResponseBody) SetRequestId(v string) *SubmitTranscodeJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTranscodeJobsResponseBody) SetTranscodeJobs(v *SubmitTranscodeJobsResponseBodyTranscodeJobs) *SubmitTranscodeJobsResponseBody {
	s.TranscodeJobs = v
	return s
}

func (s *SubmitTranscodeJobsResponseBody) SetTranscodeTaskId(v string) *SubmitTranscodeJobsResponseBody {
	s.TranscodeTaskId = &v
	return s
}

func (s *SubmitTranscodeJobsResponseBody) Validate() error {
	if s.TranscodeJobs != nil {
		if err := s.TranscodeJobs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTranscodeJobsResponseBodyTranscodeJobs struct {
	TranscodeJob []*SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob `json:"TranscodeJob,omitempty" xml:"TranscodeJob,omitempty" type:"Repeated"`
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobs) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobs) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobs) GetTranscodeJob() []*SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob {
	return s.TranscodeJob
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobs) SetTranscodeJob(v []*SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) *SubmitTranscodeJobsResponseBodyTranscodeJobs {
	s.TranscodeJob = v
	return s
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobs) Validate() error {
	if s.TranscodeJob != nil {
		for _, item := range s.TranscodeJob {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob struct {
	// The ID of the transcoding job.
	//
	// >  This parameter is not returned for HLS packaging tasks. You must asynchronously receive the transcoding result.
	//
	// example:
	//
	// d8921ce8505716cfe86fb112c4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) SetJobId(v string) *SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob {
	s.JobId = &v
	return s
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) Validate() error {
	return dara.Validate(s)
}
