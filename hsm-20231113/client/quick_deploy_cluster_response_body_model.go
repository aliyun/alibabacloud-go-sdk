// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuickDeployClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *QuickDeployClusterResponseBodyJob) *QuickDeployClusterResponseBody
	GetJob() *QuickDeployClusterResponseBodyJob
	SetRequestId(v string) *QuickDeployClusterResponseBody
	GetRequestId() *string
}

type QuickDeployClusterResponseBody struct {
	Job *QuickDeployClusterResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QuickDeployClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuickDeployClusterResponseBody) GoString() string {
	return s.String()
}

func (s *QuickDeployClusterResponseBody) GetJob() *QuickDeployClusterResponseBodyJob {
	return s.Job
}

func (s *QuickDeployClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuickDeployClusterResponseBody) SetJob(v *QuickDeployClusterResponseBodyJob) *QuickDeployClusterResponseBody {
	s.Job = v
	return s
}

func (s *QuickDeployClusterResponseBody) SetRequestId(v string) *QuickDeployClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuickDeployClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuickDeployClusterResponseBodyJob struct {
	// example:
	//
	// true
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1653274407000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// OperationTimeout
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// job-202401250936hze747fd7e0007005
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 86
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuickDeployClusterResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s QuickDeployClusterResponseBodyJob) GoString() string {
	return s.String()
}

func (s *QuickDeployClusterResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *QuickDeployClusterResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QuickDeployClusterResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *QuickDeployClusterResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *QuickDeployClusterResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *QuickDeployClusterResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *QuickDeployClusterResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *QuickDeployClusterResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *QuickDeployClusterResponseBodyJob) SetCompleted(v bool) *QuickDeployClusterResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetCreateTime(v string) *QuickDeployClusterResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetError(v string) *QuickDeployClusterResponseBodyJob {
	s.Error = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetJobId(v string) *QuickDeployClusterResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetProgress(v int32) *QuickDeployClusterResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetResponse(v string) *QuickDeployClusterResponseBodyJob {
	s.Response = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetStatus(v string) *QuickDeployClusterResponseBodyJob {
	s.Status = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) SetType(v string) *QuickDeployClusterResponseBodyJob {
	s.Type = &v
	return s
}

func (s *QuickDeployClusterResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
