// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateClusterManagedCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *RotateClusterManagedCertResponseBodyJob) *RotateClusterManagedCertResponseBody
	GetJob() *RotateClusterManagedCertResponseBodyJob
	SetRequestId(v string) *RotateClusterManagedCertResponseBody
	GetRequestId() *string
}

type RotateClusterManagedCertResponseBody struct {
	Job *RotateClusterManagedCertResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RotateClusterManagedCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RotateClusterManagedCertResponseBody) GoString() string {
	return s.String()
}

func (s *RotateClusterManagedCertResponseBody) GetJob() *RotateClusterManagedCertResponseBodyJob {
	return s.Job
}

func (s *RotateClusterManagedCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RotateClusterManagedCertResponseBody) SetJob(v *RotateClusterManagedCertResponseBodyJob) *RotateClusterManagedCertResponseBody {
	s.Job = v
	return s
}

func (s *RotateClusterManagedCertResponseBody) SetRequestId(v string) *RotateClusterManagedCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *RotateClusterManagedCertResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RotateClusterManagedCertResponseBodyJob struct {
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
	// 100
	Process *int32 `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// success
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RotateClusterManagedCertResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s RotateClusterManagedCertResponseBodyJob) GoString() string {
	return s.String()
}

func (s *RotateClusterManagedCertResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *RotateClusterManagedCertResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *RotateClusterManagedCertResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *RotateClusterManagedCertResponseBodyJob) GetJobId() *string {
	return s.JobId
}

func (s *RotateClusterManagedCertResponseBodyJob) GetProcess() *int32 {
	return s.Process
}

func (s *RotateClusterManagedCertResponseBodyJob) GetResponse() *string {
	return s.Response
}

func (s *RotateClusterManagedCertResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *RotateClusterManagedCertResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *RotateClusterManagedCertResponseBodyJob) SetCompleted(v bool) *RotateClusterManagedCertResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetCreateTime(v string) *RotateClusterManagedCertResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetError(v string) *RotateClusterManagedCertResponseBodyJob {
	s.Error = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetJobId(v string) *RotateClusterManagedCertResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetProcess(v int32) *RotateClusterManagedCertResponseBodyJob {
	s.Process = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetResponse(v string) *RotateClusterManagedCertResponseBodyJob {
	s.Response = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetStatus(v string) *RotateClusterManagedCertResponseBodyJob {
	s.Status = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) SetType(v string) *RotateClusterManagedCertResponseBodyJob {
	s.Type = &v
	return s
}

func (s *RotateClusterManagedCertResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
