// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompleted(v bool) *CopyImageResponseBody
	GetCompleted() *bool
	SetCreateTime(v string) *CopyImageResponseBody
	GetCreateTime() *string
	SetError(v string) *CopyImageResponseBody
	GetError() *string
	SetJobId(v string) *CopyImageResponseBody
	GetJobId() *string
	SetProgress(v int32) *CopyImageResponseBody
	GetProgress() *int32
	SetRequestId(v string) *CopyImageResponseBody
	GetRequestId() *string
	SetResponse(v string) *CopyImageResponseBody
	GetResponse() *string
	SetStatus(v string) *CopyImageResponseBody
	GetStatus() *string
	SetType(v string) *CopyImageResponseBody
	GetType() *string
}

type CopyImageResponseBody struct {
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
	// 1724379766191
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
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s CopyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CopyImageResponseBody) GetCompleted() *bool {
	return s.Completed
}

func (s *CopyImageResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CopyImageResponseBody) GetError() *string {
	return s.Error
}

func (s *CopyImageResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CopyImageResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *CopyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyImageResponseBody) GetResponse() *string {
	return s.Response
}

func (s *CopyImageResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CopyImageResponseBody) GetType() *string {
	return s.Type
}

func (s *CopyImageResponseBody) SetCompleted(v bool) *CopyImageResponseBody {
	s.Completed = &v
	return s
}

func (s *CopyImageResponseBody) SetCreateTime(v string) *CopyImageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CopyImageResponseBody) SetError(v string) *CopyImageResponseBody {
	s.Error = &v
	return s
}

func (s *CopyImageResponseBody) SetJobId(v string) *CopyImageResponseBody {
	s.JobId = &v
	return s
}

func (s *CopyImageResponseBody) SetProgress(v int32) *CopyImageResponseBody {
	s.Progress = &v
	return s
}

func (s *CopyImageResponseBody) SetRequestId(v string) *CopyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyImageResponseBody) SetResponse(v string) *CopyImageResponseBody {
	s.Response = &v
	return s
}

func (s *CopyImageResponseBody) SetStatus(v string) *CopyImageResponseBody {
	s.Status = &v
	return s
}

func (s *CopyImageResponseBody) SetType(v string) *CopyImageResponseBody {
	s.Type = &v
	return s
}

func (s *CopyImageResponseBody) Validate() error {
	return dara.Validate(s)
}
