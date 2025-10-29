// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobStatus(v *GetJobStatusResponseBodyJobStatus) *GetJobStatusResponseBody
	GetJobStatus() *GetJobStatusResponseBodyJobStatus
	SetRequestId(v string) *GetJobStatusResponseBody
	GetRequestId() *string
}

type GetJobStatusResponseBody struct {
	// The real-time status information of the asynchronous task.
	JobStatus *GetJobStatusResponseBodyJobStatus `json:"JobStatus,omitempty" xml:"JobStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5E2BFE96-C0E0-5A98-85C8-633EC803198D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) GetJobStatus() *GetJobStatusResponseBodyJobStatus {
	return s.JobStatus
}

func (s *GetJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobStatusResponseBody) SetJobStatus(v *GetJobStatusResponseBodyJobStatus) *GetJobStatusResponseBody {
	s.JobStatus = v
	return s
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobStatusResponseBody) Validate() error {
	if s.JobStatus != nil {
		if err := s.JobStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobStatusResponseBodyJobStatus struct {
	// Indicates whether the asynchronous task is complete. Valid values: True False
	//
	// example:
	//
	// False
	Completed *string `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The time when the asynchronous task was created.
	//
	// example:
	//
	// 1729063449802
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned if the asynchronous task fails.
	//
	// example:
	//
	// Not Found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// C664CDE3-9C0B-5792-B17F-6C543783BBBC
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the asynchronous task. Valid values:
	//
	// 	- **Create**: The asynchronous task is used to create an object.
	//
	// 	- **Update**: The asynchronous task is used to update an object.
	//
	// 	- **Cancel**: The asynchronous task is used to cancel an operation.
	//
	// example:
	//
	// Create
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Fail**
	//
	// 	- **Cancel**
	//
	// 	- **Running**
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobStatusResponseBodyJobStatus) String() string {
	return dara.Prettify(s)
}

func (s GetJobStatusResponseBodyJobStatus) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBodyJobStatus) GetCompleted() *string {
	return s.Completed
}

func (s *GetJobStatusResponseBodyJobStatus) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobStatusResponseBodyJobStatus) GetError() *string {
	return s.Error
}

func (s *GetJobStatusResponseBodyJobStatus) GetJobId() *string {
	return s.JobId
}

func (s *GetJobStatusResponseBodyJobStatus) GetJobType() *string {
	return s.JobType
}

func (s *GetJobStatusResponseBodyJobStatus) GetStatus() *string {
	return s.Status
}

func (s *GetJobStatusResponseBodyJobStatus) SetCompleted(v string) *GetJobStatusResponseBodyJobStatus {
	s.Completed = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetCreateTime(v string) *GetJobStatusResponseBodyJobStatus {
	s.CreateTime = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetError(v string) *GetJobStatusResponseBodyJobStatus {
	s.Error = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetJobId(v string) *GetJobStatusResponseBodyJobStatus {
	s.JobId = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetJobType(v string) *GetJobStatusResponseBodyJobStatus {
	s.JobType = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetStatus(v string) *GetJobStatusResponseBodyJobStatus {
	s.Status = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) Validate() error {
	return dara.Validate(s)
}
