// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetJobResponseBodyJob) *GetJobResponseBody
	GetJob() *GetJobResponseBodyJob
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
}

type GetJobResponseBody struct {
	// The details of the task.
	Job *GetJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetJob() *GetJobResponseBodyJob {
	return s.Job
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) SetJob(v *GetJobResponseBodyJob) *GetJobResponseBody {
	s.Job = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
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

type GetJobResponseBodyJob struct {
	// Indicates whether the task is complete. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The error message returned if the task fails.
	//
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// job-202401250936hze747fd7e0007005
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The progress of the task. Unit: percent (%).
	//
	// example:
	//
	// 95
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The response returned after the task succeeds.
	//
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The task status. Valid values:
	//
	// 	- success
	//
	// 	- running
	//
	// 	- fail
	//
	// 	- cancel
	//
	// example:
	//
	// fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The operation type. Valid values:
	//
	// 	- create
	//
	// 	- cancel
	//
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *GetJobResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *GetJobResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBodyJob) GetProgress() *int64 {
	return s.Progress
}

func (s *GetJobResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *GetJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *GetJobResponseBodyJob) SetCompleted(v bool) *GetJobResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *GetJobResponseBodyJob) SetError(v string) *GetJobResponseBodyJob {
	s.Error = &v
	return s
}

func (s *GetJobResponseBodyJob) SetJobId(v string) *GetJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJob) SetProgress(v int64) *GetJobResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *GetJobResponseBodyJob) SetResponse(v string) *GetJobResponseBodyJob {
	s.Response = &v
	return s
}

func (s *GetJobResponseBodyJob) SetStatus(v string) *GetJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJob) SetType(v string) *GetJobResponseBodyJob {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
