// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryRuleResponseBody
	GetCode() *string
	SetData(v *GetMemoryRuleResponseBodyData) *GetMemoryRuleResponseBody
	GetData() *GetMemoryRuleResponseBodyData
	SetHttpStatusCode(v int32) *GetMemoryRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMemoryRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMemoryRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMemoryRuleResponseBody
	GetSuccess() *bool
}

type GetMemoryRuleResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *GetMemoryRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMemoryRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryRuleResponseBody) GetData() *GetMemoryRuleResponseBodyData {
	return s.Data
}

func (s *GetMemoryRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMemoryRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMemoryRuleResponseBody) SetCode(v string) *GetMemoryRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryRuleResponseBody) SetData(v *GetMemoryRuleResponseBodyData) *GetMemoryRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetMemoryRuleResponseBody) SetHttpStatusCode(v int32) *GetMemoryRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMemoryRuleResponseBody) SetMessage(v string) *GetMemoryRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryRuleResponseBody) SetRequestId(v string) *GetMemoryRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryRuleResponseBody) SetSuccess(v bool) *GetMemoryRuleResponseBody {
	s.Success = &v
	return s
}

func (s *GetMemoryRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryRuleResponseBodyData struct {
	// createdAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// description
	//
	// example:
	//
	// Extract long-term preferences explicitly expressed by the user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// enabled
	//
	// This parameter is required.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// failedTaskCount
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	FailedTaskCount *int64 `json:"failedTaskCount,omitempty" xml:"failedTaskCount,omitempty"`
	// The time when the last execution started, in RFC 3339 format.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	LastExecutedAt *string `json:"lastExecutedAt,omitempty" xml:"lastExecutedAt,omitempty"`
	// lastTaskId
	//
	// example:
	//
	// task-2f5f9c5a8d0e4c3b
	LastTaskId *string `json:"lastTaskId,omitempty" xml:"lastTaskId,omitempty"`
	// lastTaskStatus
	//
	// example:
	//
	// SUCCEEDED
	LastTaskStatus *string `json:"lastTaskStatus,omitempty" xml:"lastTaskStatus,omitempty"`
	// memoryStoreId
	//
	// This parameter is required.
	//
	// example:
	//
	// ms-2f5f9c5a8d0e4c3b
	MemoryStoreId *string `json:"memoryStoreId,omitempty" xml:"memoryStoreId,omitempty"`
	// name
	//
	// This parameter is required.
	//
	// example:
	//
	// Long-term preferences
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// ruleId
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-2f5f9c5a8d0e4c3b
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// strategy
	//
	// This parameter is required.
	Strategy *GetMemoryRuleResponseBodyDataStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
	// succeededTaskCount
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	SucceededTaskCount *int64 `json:"succeededTaskCount,omitempty" xml:"succeededTaskCount,omitempty"`
	// The total number of distinct task IDs that have started execution. Retries are not counted separately.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	TaskCount *int64 `json:"taskCount,omitempty" xml:"taskCount,omitempty"`
	// updatedAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetMemoryRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemoryRuleResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMemoryRuleResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetMemoryRuleResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMemoryRuleResponseBodyData) GetFailedTaskCount() *int64 {
	return s.FailedTaskCount
}

func (s *GetMemoryRuleResponseBodyData) GetLastExecutedAt() *string {
	return s.LastExecutedAt
}

func (s *GetMemoryRuleResponseBodyData) GetLastTaskId() *string {
	return s.LastTaskId
}

func (s *GetMemoryRuleResponseBodyData) GetLastTaskStatus() *string {
	return s.LastTaskStatus
}

func (s *GetMemoryRuleResponseBodyData) GetMemoryStoreId() *string {
	return s.MemoryStoreId
}

func (s *GetMemoryRuleResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMemoryRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *GetMemoryRuleResponseBodyData) GetStrategy() *GetMemoryRuleResponseBodyDataStrategy {
	return s.Strategy
}

func (s *GetMemoryRuleResponseBodyData) GetSucceededTaskCount() *int64 {
	return s.SucceededTaskCount
}

func (s *GetMemoryRuleResponseBodyData) GetTaskCount() *int64 {
	return s.TaskCount
}

func (s *GetMemoryRuleResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMemoryRuleResponseBodyData) SetCreatedAt(v string) *GetMemoryRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetDescription(v string) *GetMemoryRuleResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetEnabled(v bool) *GetMemoryRuleResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetFailedTaskCount(v int64) *GetMemoryRuleResponseBodyData {
	s.FailedTaskCount = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetLastExecutedAt(v string) *GetMemoryRuleResponseBodyData {
	s.LastExecutedAt = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetLastTaskId(v string) *GetMemoryRuleResponseBodyData {
	s.LastTaskId = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetLastTaskStatus(v string) *GetMemoryRuleResponseBodyData {
	s.LastTaskStatus = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetMemoryStoreId(v string) *GetMemoryRuleResponseBodyData {
	s.MemoryStoreId = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetName(v string) *GetMemoryRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetRuleId(v string) *GetMemoryRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetStrategy(v *GetMemoryRuleResponseBodyDataStrategy) *GetMemoryRuleResponseBodyData {
	s.Strategy = v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetSucceededTaskCount(v int64) *GetMemoryRuleResponseBodyData {
	s.SucceededTaskCount = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetTaskCount(v int64) *GetMemoryRuleResponseBodyData {
	s.TaskCount = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) SetUpdatedAt(v string) *GetMemoryRuleResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetMemoryRuleResponseBodyData) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryRuleResponseBodyDataStrategy struct {
	// The strategy-level instruction. It takes effect together with the store-level instruction. If a conflict occurs, this field takes precedence.
	//
	// example:
	//
	// Retain only long-term preferences explicitly expressed by the user
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// type
	//
	// This parameter is required.
	//
	// example:
	//
	// auto
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMemoryRuleResponseBodyDataStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryRuleResponseBodyDataStrategy) GoString() string {
	return s.String()
}

func (s *GetMemoryRuleResponseBodyDataStrategy) GetInstruction() *string {
	return s.Instruction
}

func (s *GetMemoryRuleResponseBodyDataStrategy) GetType() *string {
	return s.Type
}

func (s *GetMemoryRuleResponseBodyDataStrategy) SetInstruction(v string) *GetMemoryRuleResponseBodyDataStrategy {
	s.Instruction = &v
	return s
}

func (s *GetMemoryRuleResponseBodyDataStrategy) SetType(v string) *GetMemoryRuleResponseBodyDataStrategy {
	s.Type = &v
	return s
}

func (s *GetMemoryRuleResponseBodyDataStrategy) Validate() error {
	return dara.Validate(s)
}
