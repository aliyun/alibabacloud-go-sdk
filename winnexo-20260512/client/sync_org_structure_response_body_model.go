// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncOrgStructureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SyncOrgStructureResponseBody
	GetCode() *string
	SetMessage(v string) *SyncOrgStructureResponseBody
	GetMessage() *string
	SetRequestId(v string) *SyncOrgStructureResponseBody
	GetRequestId() *string
	SetStatus(v string) *SyncOrgStructureResponseBody
	GetStatus() *string
	SetTaskId(v int64) *SyncOrgStructureResponseBody
	GetTaskId() *int64
}

type SyncOrgStructureResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This value is empty when the request is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The initial status of the task.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The asynchronous synchronization task ID, used for polling the status through querySyncResult.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s SyncOrgStructureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncOrgStructureResponseBody) GoString() string {
	return s.String()
}

func (s *SyncOrgStructureResponseBody) GetCode() *string {
	return s.Code
}

func (s *SyncOrgStructureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SyncOrgStructureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncOrgStructureResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SyncOrgStructureResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SyncOrgStructureResponseBody) SetCode(v string) *SyncOrgStructureResponseBody {
	s.Code = &v
	return s
}

func (s *SyncOrgStructureResponseBody) SetMessage(v string) *SyncOrgStructureResponseBody {
	s.Message = &v
	return s
}

func (s *SyncOrgStructureResponseBody) SetRequestId(v string) *SyncOrgStructureResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncOrgStructureResponseBody) SetStatus(v string) *SyncOrgStructureResponseBody {
	s.Status = &v
	return s
}

func (s *SyncOrgStructureResponseBody) SetTaskId(v int64) *SyncOrgStructureResponseBody {
	s.TaskId = &v
	return s
}

func (s *SyncOrgStructureResponseBody) Validate() error {
	return dara.Validate(s)
}
