// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreprocessJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreprocessJobs(v *SubmitPreprocessJobsResponseBodyPreprocessJobs) *SubmitPreprocessJobsResponseBody
	GetPreprocessJobs() *SubmitPreprocessJobsResponseBodyPreprocessJobs
	SetRequestId(v string) *SubmitPreprocessJobsResponseBody
	GetRequestId() *string
}

type SubmitPreprocessJobsResponseBody struct {
	// The information about the job.
	PreprocessJobs *SubmitPreprocessJobsResponseBodyPreprocessJobs `json:"PreprocessJobs,omitempty" xml:"PreprocessJobs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E4EBD2BF-5EB0-4476-8829-9D94E1B1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitPreprocessJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponseBody) GetPreprocessJobs() *SubmitPreprocessJobsResponseBodyPreprocessJobs {
	return s.PreprocessJobs
}

func (s *SubmitPreprocessJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPreprocessJobsResponseBody) SetPreprocessJobs(v *SubmitPreprocessJobsResponseBodyPreprocessJobs) *SubmitPreprocessJobsResponseBody {
	s.PreprocessJobs = v
	return s
}

func (s *SubmitPreprocessJobsResponseBody) SetRequestId(v string) *SubmitPreprocessJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPreprocessJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitPreprocessJobsResponseBodyPreprocessJobs struct {
	PreprocessJob []*SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob `json:"PreprocessJob,omitempty" xml:"PreprocessJob,omitempty" type:"Repeated"`
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobs) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobs) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobs) GetPreprocessJob() []*SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob {
	return s.PreprocessJob
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobs) SetPreprocessJob(v []*SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) *SubmitPreprocessJobsResponseBodyPreprocessJobs {
	s.PreprocessJob = v
	return s
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobs) Validate() error {
	return dara.Validate(s)
}

type SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob struct {
	// The ID of the job.
	//
	// example:
	//
	// bb396607fd11fee9effbb99c4****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) SetJobId(v string) *SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob {
	s.JobId = &v
	return s
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) Validate() error {
	return dara.Validate(s)
}
