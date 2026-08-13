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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 任务初始状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 异步同步任务 ID，用于 querySyncResult 轮询状态
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
