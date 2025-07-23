// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *RestoreInstanceResponseBodyJob) *RestoreInstanceResponseBody
	GetJob() *RestoreInstanceResponseBodyJob
	SetRequestId(v string) *RestoreInstanceResponseBody
	GetRequestId() *string
}

type RestoreInstanceResponseBody struct {
	// The details of the task.
	Job *RestoreInstanceResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBody) GetJob() *RestoreInstanceResponseBodyJob {
	return s.Job
}

func (s *RestoreInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreInstanceResponseBody) SetJob(v *RestoreInstanceResponseBodyJob) *RestoreInstanceResponseBody {
	s.Job = v
	return s
}

func (s *RestoreInstanceResponseBody) SetRequestId(v string) *RestoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RestoreInstanceResponseBodyJob struct {
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
	// job-540356379023708160
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The progress of the task. Unit: percent (%).
	//
	// example:
	//
	// 50
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The response returned after the task succeeds.
	//
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// The task status.
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

func (s RestoreInstanceResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceResponseBodyJob) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *RestoreInstanceResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RestoreInstanceResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *RestoreInstanceResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *RestoreInstanceResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *RestoreInstanceResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *RestoreInstanceResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *RestoreInstanceResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *RestoreInstanceResponseBodyJob) SetCompleted(v bool) *RestoreInstanceResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetCreateTime(v string) *RestoreInstanceResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetError(v string) *RestoreInstanceResponseBodyJob {
	s.Error = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetJobId(v string) *RestoreInstanceResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetProgress(v int32) *RestoreInstanceResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetResponse(v string) *RestoreInstanceResponseBodyJob {
	s.Response = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetStatus(v string) *RestoreInstanceResponseBodyJob {
	s.Status = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) SetType(v string) *RestoreInstanceResponseBodyJob {
	s.Type = &v
	return s
}

func (s *RestoreInstanceResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
