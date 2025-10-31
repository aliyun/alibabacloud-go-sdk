// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *SyncClusterResponseBodyJob) *SyncClusterResponseBody
	GetJob() *SyncClusterResponseBodyJob
	SetRequestId(v string) *SyncClusterResponseBody
	GetRequestId() *string
}

type SyncClusterResponseBody struct {
	// The details of the task.
	Job *SyncClusterResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncClusterResponseBody) GoString() string {
	return s.String()
}

func (s *SyncClusterResponseBody) GetJob() *SyncClusterResponseBodyJob {
	return s.Job
}

func (s *SyncClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncClusterResponseBody) SetJob(v *SyncClusterResponseBodyJob) *SyncClusterResponseBody {
	s.Job = v
	return s
}

func (s *SyncClusterResponseBody) SetRequestId(v string) *SyncClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncClusterResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncClusterResponseBodyJob struct {
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
	// job-000bu7m5vjmyz9s7qz85
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The progress of the task. Unit: percent (%).
	//
	// example:
	//
	// 90
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

func (s SyncClusterResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s SyncClusterResponseBodyJob) GoString() string {
	return s.String()
}

func (s *SyncClusterResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *SyncClusterResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SyncClusterResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *SyncClusterResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *SyncClusterResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *SyncClusterResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *SyncClusterResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *SyncClusterResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *SyncClusterResponseBodyJob) SetCompleted(v bool) *SyncClusterResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetCreateTime(v string) *SyncClusterResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetError(v string) *SyncClusterResponseBodyJob {
	s.Error = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetJobId(v string) *SyncClusterResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetProgress(v int32) *SyncClusterResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetResponse(v string) *SyncClusterResponseBodyJob {
	s.Response = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetStatus(v string) *SyncClusterResponseBodyJob {
	s.Status = &v
	return s
}

func (s *SyncClusterResponseBodyJob) SetType(v string) *SyncClusterResponseBodyJob {
	s.Type = &v
	return s
}

func (s *SyncClusterResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
