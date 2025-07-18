// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpsertCollectionDataJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetUpsertCollectionDataJobResponseBodyJob) *GetUpsertCollectionDataJobResponseBody
	GetJob() *GetUpsertCollectionDataJobResponseBodyJob
	SetMessage(v string) *GetUpsertCollectionDataJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUpsertCollectionDataJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetUpsertCollectionDataJobResponseBody
	GetStatus() *string
}

type GetUpsertCollectionDataJobResponseBody struct {
	// The information about the vector data upload job.
	Job *GetUpsertCollectionDataJobResponseBodyJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetUpsertCollectionDataJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUpsertCollectionDataJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetUpsertCollectionDataJobResponseBody) GetJob() *GetUpsertCollectionDataJobResponseBodyJob {
	return s.Job
}

func (s *GetUpsertCollectionDataJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUpsertCollectionDataJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUpsertCollectionDataJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetUpsertCollectionDataJobResponseBody) SetJob(v *GetUpsertCollectionDataJobResponseBodyJob) *GetUpsertCollectionDataJobResponseBody {
	s.Job = v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBody) SetMessage(v string) *GetUpsertCollectionDataJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBody) SetRequestId(v string) *GetUpsertCollectionDataJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBody) SetStatus(v string) *GetUpsertCollectionDataJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUpsertCollectionDataJobResponseBodyJob struct {
	// Indicates whether the operation is complete.
	//
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2024-01-08 16:52:04.864664
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// Failed to connect database.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The progress of the vector data upload job. The value of this parameter indicates the number of data entries that have been uploaded.
	//
	// example:
	//
	// 1600
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The status of the job.
	//
	// >  Valid values:
	//
	// 	- Success
	//
	// 	- Failed (See the Error parameter for failure reasons.)
	//
	// 	- Cancelling
	//
	// 	- Cancelled
	//
	// 	- Start
	//
	// 	- Running
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the job was updated.
	//
	// example:
	//
	// 2024-01-08 16:53:04.864664
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetUpsertCollectionDataJobResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetUpsertCollectionDataJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetCompleted() *bool {
	return s.Completed
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetError() *string {
	return s.Error
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetId() *string {
	return s.Id
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetProgress() *int32 {
	return s.Progress
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetCompleted(v bool) *GetUpsertCollectionDataJobResponseBodyJob {
	s.Completed = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetCreateTime(v string) *GetUpsertCollectionDataJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetError(v string) *GetUpsertCollectionDataJobResponseBodyJob {
	s.Error = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetId(v string) *GetUpsertCollectionDataJobResponseBodyJob {
	s.Id = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetProgress(v int32) *GetUpsertCollectionDataJobResponseBodyJob {
	s.Progress = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetStatus(v string) *GetUpsertCollectionDataJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) SetUpdateTime(v string) *GetUpsertCollectionDataJobResponseBodyJob {
	s.UpdateTime = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponseBodyJob) Validate() error {
	return dara.Validate(s)
}
