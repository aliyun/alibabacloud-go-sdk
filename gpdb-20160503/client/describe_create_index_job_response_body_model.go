// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreateIndexJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *DescribeCreateIndexJobResponseBodyJob) *DescribeCreateIndexJobResponseBody
	GetJob() *DescribeCreateIndexJobResponseBodyJob
	SetMessage(v string) *DescribeCreateIndexJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCreateIndexJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeCreateIndexJobResponseBody
	GetStatus() *string
}

type DescribeCreateIndexJobResponseBody struct {
	Job *DescribeCreateIndexJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCreateIndexJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreateIndexJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreateIndexJobResponseBody) GetJob() *DescribeCreateIndexJobResponseBodyJob {
	return s.Job
}

func (s *DescribeCreateIndexJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCreateIndexJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreateIndexJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeCreateIndexJobResponseBody) SetJob(v *DescribeCreateIndexJobResponseBodyJob) *DescribeCreateIndexJobResponseBody {
	s.Job = v
	return s
}

func (s *DescribeCreateIndexJobResponseBody) SetMessage(v string) *DescribeCreateIndexJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBody) SetRequestId(v string) *DescribeCreateIndexJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBody) SetStatus(v string) *DescribeCreateIndexJobResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCreateIndexJobResponseBodyJob struct {
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-01-08 16:52:04.864664
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Failed to connect database.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// Job ID。
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 20
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2024-01-08 16:53:04.864664
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeCreateIndexJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreateIndexJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetId() *string {
	return s.Id
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeCreateIndexJobResponseBodyJob) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetCompleted(v bool) *DescribeCreateIndexJobResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetCreateTime(v string) *DescribeCreateIndexJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetError(v string) *DescribeCreateIndexJobResponseBodyJob {
	s.Error = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetId(v string) *DescribeCreateIndexJobResponseBodyJob {
	s.Id = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetProgress(v int32) *DescribeCreateIndexJobResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetStatus(v string) *DescribeCreateIndexJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) SetUpdateTime(v string) *DescribeCreateIndexJobResponseBodyJob {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCreateIndexJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
