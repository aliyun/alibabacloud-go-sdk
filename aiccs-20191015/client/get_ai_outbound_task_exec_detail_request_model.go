// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskExecDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchVersion(v int32) *GetAiOutboundTaskExecDetailRequest
	GetBatchVersion() *int32
	SetCaseId(v int64) *GetAiOutboundTaskExecDetailRequest
	GetCaseId() *int64
	SetCaseStatus(v int32) *GetAiOutboundTaskExecDetailRequest
	GetCaseStatus() *int32
	SetCreateTimeEnd(v int64) *GetAiOutboundTaskExecDetailRequest
	GetCreateTimeEnd() *int64
	SetCreateTimeStart(v int64) *GetAiOutboundTaskExecDetailRequest
	GetCreateTimeStart() *int64
	SetCurrentPage(v int32) *GetAiOutboundTaskExecDetailRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *GetAiOutboundTaskExecDetailRequest
	GetInstanceId() *string
	SetPageSize(v int32) *GetAiOutboundTaskExecDetailRequest
	GetPageSize() *int32
	SetPhoneNum(v string) *GetAiOutboundTaskExecDetailRequest
	GetPhoneNum() *string
	SetTaskId(v int64) *GetAiOutboundTaskExecDetailRequest
	GetTaskId() *int64
}

type GetAiOutboundTaskExecDetailRequest struct {
	// example:
	//
	// 1
	BatchVersion *int32 `json:"BatchVersion,omitempty" xml:"BatchVersion,omitempty"`
	// example:
	//
	// 123456
	CaseId *int64 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	// example:
	//
	// 1
	CaseStatus *int32 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	// example:
	//
	// 1632290119000
	CreateTimeEnd *int64 `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 1632289999000
	CreateTimeStart *int64 `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agent_***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 150****000
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAiOutboundTaskExecDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskExecDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskExecDetailRequest) GetBatchVersion() *int32 {
	return s.BatchVersion
}

func (s *GetAiOutboundTaskExecDetailRequest) GetCaseId() *int64 {
	return s.CaseId
}

func (s *GetAiOutboundTaskExecDetailRequest) GetCaseStatus() *int32 {
	return s.CaseStatus
}

func (s *GetAiOutboundTaskExecDetailRequest) GetCreateTimeEnd() *int64 {
	return s.CreateTimeEnd
}

func (s *GetAiOutboundTaskExecDetailRequest) GetCreateTimeStart() *int64 {
	return s.CreateTimeStart
}

func (s *GetAiOutboundTaskExecDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAiOutboundTaskExecDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAiOutboundTaskExecDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAiOutboundTaskExecDetailRequest) GetPhoneNum() *string {
	return s.PhoneNum
}

func (s *GetAiOutboundTaskExecDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetAiOutboundTaskExecDetailRequest) SetBatchVersion(v int32) *GetAiOutboundTaskExecDetailRequest {
	s.BatchVersion = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetCaseId(v int64) *GetAiOutboundTaskExecDetailRequest {
	s.CaseId = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetCaseStatus(v int32) *GetAiOutboundTaskExecDetailRequest {
	s.CaseStatus = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetCreateTimeEnd(v int64) *GetAiOutboundTaskExecDetailRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetCreateTimeStart(v int64) *GetAiOutboundTaskExecDetailRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetCurrentPage(v int32) *GetAiOutboundTaskExecDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetInstanceId(v string) *GetAiOutboundTaskExecDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetPageSize(v int32) *GetAiOutboundTaskExecDetailRequest {
	s.PageSize = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetPhoneNum(v string) *GetAiOutboundTaskExecDetailRequest {
	s.PhoneNum = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) SetTaskId(v int64) *GetAiOutboundTaskExecDetailRequest {
	s.TaskId = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailRequest) Validate() error {
	return dara.Validate(s)
}
