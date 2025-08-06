// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreprocessJobsConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitPreprocessJobsConsoleResponseBody
	GetRequestId() *string
	SetTranscodeJobs(v *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) *SubmitPreprocessJobsConsoleResponseBody
	GetTranscodeJobs() *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs
}

type SubmitPreprocessJobsConsoleResponseBody struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeJobs *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs `json:"TranscodeJobs,omitempty" xml:"TranscodeJobs,omitempty" type:"Struct"`
}

func (s SubmitPreprocessJobsConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPreprocessJobsConsoleResponseBody) GetTranscodeJobs() *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs {
	return s.TranscodeJobs
}

func (s *SubmitPreprocessJobsConsoleResponseBody) SetRequestId(v string) *SubmitPreprocessJobsConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponseBody) SetTranscodeJobs(v *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) *SubmitPreprocessJobsConsoleResponseBody {
	s.TranscodeJobs = v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs struct {
	TranscodeJob []*SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob `json:"TranscodeJob,omitempty" xml:"TranscodeJob,omitempty" type:"Repeated"`
}

func (s SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) GetTranscodeJob() []*SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob {
	return s.TranscodeJob
}

func (s *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) SetTranscodeJob(v []*SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob) *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs {
	s.TranscodeJob = v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobs) Validate() error {
	return dara.Validate(s)
}

type SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob) GetJobId() *string {
	return s.JobId
}

func (s *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob) SetJobId(v string) *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob {
	s.JobId = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleResponseBodyTranscodeJobsTranscodeJob) Validate() error {
	return dara.Validate(s)
}
