// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *JoinClusterResponseBodyJob) *JoinClusterResponseBody
	GetJob() *JoinClusterResponseBodyJob
	SetRequestId(v string) *JoinClusterResponseBody
	GetRequestId() *string
}

type JoinClusterResponseBody struct {
	// The details of the task.
	Job *JoinClusterResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JoinClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s JoinClusterResponseBody) GoString() string {
	return s.String()
}

func (s *JoinClusterResponseBody) GetJob() *JoinClusterResponseBodyJob {
	return s.Job
}

func (s *JoinClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *JoinClusterResponseBody) SetJob(v *JoinClusterResponseBodyJob) *JoinClusterResponseBody {
	s.Job = v
	return s
}

func (s *JoinClusterResponseBody) SetRequestId(v string) *JoinClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinClusterResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type JoinClusterResponseBodyJob struct {
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
	// The time when the task was created. Unit: milliseconds. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1711764127000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// 86
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The response returned after the task succeeds.
	//
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- success
	//
	// 	- running
	//
	// 	- cancel
	//
	// 	- fail
	//
	// example:
	//
	// running
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

func (s JoinClusterResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s JoinClusterResponseBodyJob) GoString() string {
	return s.String()
}

func (s *JoinClusterResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *JoinClusterResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *JoinClusterResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *JoinClusterResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *JoinClusterResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *JoinClusterResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *JoinClusterResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *JoinClusterResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *JoinClusterResponseBodyJob) SetCompleted(v bool) *JoinClusterResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetCreateTime(v string) *JoinClusterResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetError(v string) *JoinClusterResponseBodyJob {
	s.Error = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetJobId(v string) *JoinClusterResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetProgress(v int32) *JoinClusterResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetResponse(v string) *JoinClusterResponseBodyJob {
	s.Response = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetStatus(v string) *JoinClusterResponseBodyJob {
	s.Status = &v
	return s
}

func (s *JoinClusterResponseBodyJob) SetType(v string) *JoinClusterResponseBodyJob {
	s.Type = &v
	return s
}

func (s *JoinClusterResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
