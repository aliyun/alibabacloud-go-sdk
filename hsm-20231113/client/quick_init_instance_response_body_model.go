// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickInitInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *QuickInitInstanceResponseBodyJob) *QuickInitInstanceResponseBody
	GetJob() *QuickInitInstanceResponseBodyJob
	SetRequestId(v string) *QuickInitInstanceResponseBody
	GetRequestId() *string
}

type QuickInitInstanceResponseBody struct {
	// The details of the task.
	Job *QuickInitInstanceResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuickInitInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuickInitInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceResponseBody) GetJob() *QuickInitInstanceResponseBodyJob {
	return s.Job
}

func (s *QuickInitInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuickInitInstanceResponseBody) SetJob(v *QuickInitInstanceResponseBodyJob) *QuickInitInstanceResponseBody {
	s.Job = v
	return s
}

func (s *QuickInitInstanceResponseBody) SetRequestId(v string) *QuickInitInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuickInitInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuickInitInstanceResponseBodyJob struct {
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
	// 1699515963000
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
	// job-000fi9k1v2hclo321sal
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The progress of the task. Unit: percent (%).
	//
	// example:
	//
	// 100
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

func (s QuickInitInstanceResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s QuickInitInstanceResponseBodyJob) GoString() string {
	return s.String()
}

func (s *QuickInitInstanceResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *QuickInitInstanceResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QuickInitInstanceResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *QuickInitInstanceResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *QuickInitInstanceResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *QuickInitInstanceResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *QuickInitInstanceResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *QuickInitInstanceResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *QuickInitInstanceResponseBodyJob) SetCompleted(v bool) *QuickInitInstanceResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetCreateTime(v string) *QuickInitInstanceResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetError(v string) *QuickInitInstanceResponseBodyJob {
	s.Error = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetJobId(v string) *QuickInitInstanceResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetProgress(v int32) *QuickInitInstanceResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetResponse(v string) *QuickInitInstanceResponseBodyJob {
	s.Response = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetStatus(v string) *QuickInitInstanceResponseBodyJob {
	s.Status = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) SetType(v string) *QuickInitInstanceResponseBodyJob {
	s.Type = &v
	return s
}

func (s *QuickInitInstanceResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
