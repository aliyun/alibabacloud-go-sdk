// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotfixTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotfixTaskResponseBody
	GetCode() *string
	SetData(v *GetHotfixTaskResponseBodyData) *GetHotfixTaskResponseBody
	GetData() *GetHotfixTaskResponseBodyData
	SetMessage(v string) *GetHotfixTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotfixTaskResponseBody
	GetRequestId() *string
}

type GetHotfixTaskResponseBody struct {
	// The status code.
	//
	// - If `code` is `Success`, the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. If the authorization fails, check the `message` field for the detailed error message of the fault.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetHotfixTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code` is `Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotfixTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotfixTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotfixTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotfixTaskResponseBody) GetData() *GetHotfixTaskResponseBodyData {
	return s.Data
}

func (s *GetHotfixTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotfixTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotfixTaskResponseBody) SetCode(v string) *GetHotfixTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotfixTaskResponseBody) SetData(v *GetHotfixTaskResponseBodyData) *GetHotfixTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetHotfixTaskResponseBody) SetMessage(v string) *GetHotfixTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotfixTaskResponseBody) SetRequestId(v string) *GetHotfixTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotfixTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotfixTaskResponseBodyData struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2026-10-13T08:14:40
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// CVE ID
	//
	// example:
	//
	// CVE-2026-31431
	CveId *string `json:"cveId,omitempty" xml:"cveId,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2026-10-13T08:14:40
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The list of instance tasks.
	Jobs []*GetHotfixTaskResponseBodyDataJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// The operation type of the task.
	//
	// example:
	//
	// Install
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the asynchronous task.
	//
	// example:
	//
	// hotfix-da3dcc0b-5019-4b59-b12f-2e500da9083c
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetHotfixTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHotfixTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotfixTaskResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetHotfixTaskResponseBodyData) GetCveId() *string {
	return s.CveId
}

func (s *GetHotfixTaskResponseBodyData) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetHotfixTaskResponseBodyData) GetJobs() []*GetHotfixTaskResponseBodyDataJobs {
	return s.Jobs
}

func (s *GetHotfixTaskResponseBodyData) GetOperation() *string {
	return s.Operation
}

func (s *GetHotfixTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetHotfixTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetHotfixTaskResponseBodyData) SetCreatedAt(v string) *GetHotfixTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetHotfixTaskResponseBodyData) SetCveId(v string) *GetHotfixTaskResponseBodyData {
	s.CveId = &v
	return s
}

func (s *GetHotfixTaskResponseBodyData) SetFinishedAt(v string) *GetHotfixTaskResponseBodyData {
	s.FinishedAt = &v
	return s
}

func (s *GetHotfixTaskResponseBodyData) SetJobs(v []*GetHotfixTaskResponseBodyDataJobs) *GetHotfixTaskResponseBodyData {
	s.Jobs = v
	return s
}

func (s *GetHotfixTaskResponseBodyData) SetOperation(v string) *GetHotfixTaskResponseBodyData {
	s.Operation = &v
	return s
}

func (s *GetHotfixTaskResponseBodyData) SetStatus(v string) *GetHotfixTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetHotfixTaskResponseBodyData) SetTaskId(v string) *GetHotfixTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetHotfixTaskResponseBodyData) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotfixTaskResponseBodyDataJobs struct {
	// The unified error code.
	//
	// example:
	//
	// INSTANCE_NOT_EXISTS
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified instance does not exist
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The hotfix status.
	//
	// example:
	//
	// CONFLICT
	HotfixStatus *string `json:"hotfixStatus,omitempty" xml:"hotfixStatus,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-2zehme0rs1tc090fdl3n
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The status of the instance task.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetHotfixTaskResponseBodyDataJobs) String() string {
	return dara.Prettify(s)
}

func (s GetHotfixTaskResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *GetHotfixTaskResponseBodyDataJobs) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetHotfixTaskResponseBodyDataJobs) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetHotfixTaskResponseBodyDataJobs) GetHotfixStatus() *string {
	return s.HotfixStatus
}

func (s *GetHotfixTaskResponseBodyDataJobs) GetInstance() *string {
	return s.Instance
}

func (s *GetHotfixTaskResponseBodyDataJobs) GetRegion() *string {
	return s.Region
}

func (s *GetHotfixTaskResponseBodyDataJobs) GetStatus() *string {
	return s.Status
}

func (s *GetHotfixTaskResponseBodyDataJobs) SetErrorCode(v string) *GetHotfixTaskResponseBodyDataJobs {
	s.ErrorCode = &v
	return s
}

func (s *GetHotfixTaskResponseBodyDataJobs) SetErrorMessage(v string) *GetHotfixTaskResponseBodyDataJobs {
	s.ErrorMessage = &v
	return s
}

func (s *GetHotfixTaskResponseBodyDataJobs) SetHotfixStatus(v string) *GetHotfixTaskResponseBodyDataJobs {
	s.HotfixStatus = &v
	return s
}

func (s *GetHotfixTaskResponseBodyDataJobs) SetInstance(v string) *GetHotfixTaskResponseBodyDataJobs {
	s.Instance = &v
	return s
}

func (s *GetHotfixTaskResponseBodyDataJobs) SetRegion(v string) *GetHotfixTaskResponseBodyDataJobs {
	s.Region = &v
	return s
}

func (s *GetHotfixTaskResponseBodyDataJobs) SetStatus(v string) *GetHotfixTaskResponseBodyDataJobs {
	s.Status = &v
	return s
}

func (s *GetHotfixTaskResponseBodyDataJobs) Validate() error {
	return dara.Validate(s)
}
