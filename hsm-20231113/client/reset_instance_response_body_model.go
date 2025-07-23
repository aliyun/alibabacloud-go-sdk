// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *ResetInstanceResponseBodyJob) *ResetInstanceResponseBody
	GetJob() *ResetInstanceResponseBodyJob
	SetRequestId(v string) *ResetInstanceResponseBody
	GetRequestId() *string
}

type ResetInstanceResponseBody struct {
	// The details of the task.
	Job *ResetInstanceResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponseBody) GetJob() *ResetInstanceResponseBodyJob {
	return s.Job
}

func (s *ResetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetInstanceResponseBody) SetJob(v *ResetInstanceResponseBodyJob) *ResetInstanceResponseBody {
	s.Job = v
	return s
}

func (s *ResetInstanceResponseBody) SetRequestId(v string) *ResetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ResetInstanceResponseBodyJob struct {
	// Indicates whether the task is complete.
	//
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The time when the task is created. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1653274407000
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
	// job-0007bl8oev0u3jqyfu6a
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The progress of the task. Unit: percent (%).
	//
	// example:
	//
	// 80
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
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
	// 	- cancel
	//
	// 	- fail
	//
	// example:
	//
	// success
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

func (s ResetInstanceResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceResponseBodyJob) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *ResetInstanceResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ResetInstanceResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *ResetInstanceResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *ResetInstanceResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *ResetInstanceResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *ResetInstanceResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *ResetInstanceResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *ResetInstanceResponseBodyJob) SetCompleted(v bool) *ResetInstanceResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetCreateTime(v string) *ResetInstanceResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetError(v string) *ResetInstanceResponseBodyJob {
	s.Error = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetJobId(v string) *ResetInstanceResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetProgress(v int32) *ResetInstanceResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetResponse(v string) *ResetInstanceResponseBodyJob {
	s.Response = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetStatus(v string) *ResetInstanceResponseBodyJob {
	s.Status = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) SetType(v string) *ResetInstanceResponseBodyJob {
	s.Type = &v
	return s
}

func (s *ResetInstanceResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
