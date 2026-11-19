// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoryTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMemoryTasksResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMemoryTasksResponseBodyItems) *ListMemoryTasksResponseBody
	GetItems() []*ListMemoryTasksResponseBodyItems
	SetMaxResults(v int32) *ListMemoryTasksResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMemoryTasksResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMemoryTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoryTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemoryTasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMemoryTasksResponseBody
	GetTotalCount() *int64
}

type ListMemoryTasksResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// items
	Items []*ListMemoryTasksResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// maxResults
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// nextToken
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// totalCount
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMemoryTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemoryTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemoryTasksResponseBody) GetItems() []*ListMemoryTasksResponseBodyItems {
	return s.Items
}

func (s *ListMemoryTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemoryTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemoryTasksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMemoryTasksResponseBody) SetCode(v string) *ListMemoryTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetHttpStatusCode(v int32) *ListMemoryTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetItems(v []*ListMemoryTasksResponseBodyItems) *ListMemoryTasksResponseBody {
	s.Items = v
	return s
}

func (s *ListMemoryTasksResponseBody) SetMaxResults(v int32) *ListMemoryTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetMessage(v string) *ListMemoryTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetNextToken(v string) *ListMemoryTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetRequestId(v string) *ListMemoryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetSuccess(v bool) *ListMemoryTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemoryTasksResponseBody) SetTotalCount(v int64) *ListMemoryTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMemoryTasksResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemoryTasksResponseBodyItems struct {
	// changes
	Changes []*ListMemoryTasksResponseBodyItemsChanges `json:"changes,omitempty" xml:"changes,omitempty" type:"Repeated"`
	// createdAt
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// error
	Error *ListMemoryTasksResponseBodyItemsError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// finishedAt
	//
	// example:
	//
	// 2026-09-23T10:00:01Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// inputMessageCount
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	InputMessageCount *int64 `json:"inputMessageCount,omitempty" xml:"inputMessageCount,omitempty"`
	// inputTokenCount
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	InputTokenCount *int64 `json:"inputTokenCount,omitempty" xml:"inputTokenCount,omitempty"`
	// memories
	//
	// This parameter is required.
	Memories []*ListMemoryTasksResponseBodyItemsMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
	// startedAt
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// status
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// strategyConfig
	//
	// This parameter is required.
	StrategyConfig *ListMemoryTasksResponseBodyItemsStrategyConfig `json:"strategyConfig,omitempty" xml:"strategyConfig,omitempty" type:"Struct"`
	// taskId
	//
	// This parameter is required.
	//
	// example:
	//
	// task-2f5f9c5a8d0e4c3b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// triggerType
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListMemoryTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponseBodyItems) GetChanges() []*ListMemoryTasksResponseBodyItemsChanges {
	return s.Changes
}

func (s *ListMemoryTasksResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMemoryTasksResponseBodyItems) GetError() *ListMemoryTasksResponseBodyItemsError {
	return s.Error
}

func (s *ListMemoryTasksResponseBodyItems) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListMemoryTasksResponseBodyItems) GetInputMessageCount() *int64 {
	return s.InputMessageCount
}

func (s *ListMemoryTasksResponseBodyItems) GetInputTokenCount() *int64 {
	return s.InputTokenCount
}

func (s *ListMemoryTasksResponseBodyItems) GetMemories() []*ListMemoryTasksResponseBodyItemsMemories {
	return s.Memories
}

func (s *ListMemoryTasksResponseBodyItems) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListMemoryTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListMemoryTasksResponseBodyItems) GetStrategyConfig() *ListMemoryTasksResponseBodyItemsStrategyConfig {
	return s.StrategyConfig
}

func (s *ListMemoryTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListMemoryTasksResponseBodyItems) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListMemoryTasksResponseBodyItems) SetChanges(v []*ListMemoryTasksResponseBodyItemsChanges) *ListMemoryTasksResponseBodyItems {
	s.Changes = v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetCreatedAt(v string) *ListMemoryTasksResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetError(v *ListMemoryTasksResponseBodyItemsError) *ListMemoryTasksResponseBodyItems {
	s.Error = v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetFinishedAt(v string) *ListMemoryTasksResponseBodyItems {
	s.FinishedAt = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetInputMessageCount(v int64) *ListMemoryTasksResponseBodyItems {
	s.InputMessageCount = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetInputTokenCount(v int64) *ListMemoryTasksResponseBodyItems {
	s.InputTokenCount = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetMemories(v []*ListMemoryTasksResponseBodyItemsMemories) *ListMemoryTasksResponseBodyItems {
	s.Memories = v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetStartedAt(v string) *ListMemoryTasksResponseBodyItems {
	s.StartedAt = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetStatus(v string) *ListMemoryTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetStrategyConfig(v *ListMemoryTasksResponseBodyItemsStrategyConfig) *ListMemoryTasksResponseBodyItems {
	s.StrategyConfig = v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetTaskId(v string) *ListMemoryTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) SetTriggerType(v string) *ListMemoryTasksResponseBodyItems {
	s.TriggerType = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItems) Validate() error {
	if s.Changes != nil {
		for _, item := range s.Changes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	if s.Memories != nil {
		for _, item := range s.Memories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StrategyConfig != nil {
		if err := s.StrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemoryTasksResponseBodyItemsChanges struct {
	// memoryId
	//
	// This parameter is required.
	//
	// example:
	//
	// mem-2f5f9c5a8d0e4c3b
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// operation
	//
	// This parameter is required.
	//
	// example:
	//
	// ADD
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
}

func (s ListMemoryTasksResponseBodyItemsChanges) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponseBodyItemsChanges) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponseBodyItemsChanges) GetMemoryId() *string {
	return s.MemoryId
}

func (s *ListMemoryTasksResponseBodyItemsChanges) GetOperation() *string {
	return s.Operation
}

func (s *ListMemoryTasksResponseBodyItemsChanges) SetMemoryId(v string) *ListMemoryTasksResponseBodyItemsChanges {
	s.MemoryId = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsChanges) SetOperation(v string) *ListMemoryTasksResponseBodyItemsChanges {
	s.Operation = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsChanges) Validate() error {
	return dara.Validate(s)
}

type ListMemoryTasksResponseBodyItemsError struct {
	// code
	//
	// This parameter is required.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	//
	// This parameter is required.
	//
	// example:
	//
	// The event content format is incorrect
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListMemoryTasksResponseBodyItemsError) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponseBodyItemsError) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponseBodyItemsError) GetCode() *string {
	return s.Code
}

func (s *ListMemoryTasksResponseBodyItemsError) GetMessage() *string {
	return s.Message
}

func (s *ListMemoryTasksResponseBodyItemsError) SetCode(v string) *ListMemoryTasksResponseBodyItemsError {
	s.Code = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsError) SetMessage(v string) *ListMemoryTasksResponseBodyItemsError {
	s.Message = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsError) Validate() error {
	return dara.Validate(s)
}

type ListMemoryTasksResponseBodyItemsMemories struct {
	// memoryId
	//
	// This parameter is required.
	//
	// example:
	//
	// mem-2f5f9c5a8d0e4c3b
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
}

func (s ListMemoryTasksResponseBodyItemsMemories) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponseBodyItemsMemories) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponseBodyItemsMemories) GetMemoryId() *string {
	return s.MemoryId
}

func (s *ListMemoryTasksResponseBodyItemsMemories) SetMemoryId(v string) *ListMemoryTasksResponseBodyItemsMemories {
	s.MemoryId = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsMemories) Validate() error {
	return dara.Validate(s)
}

type ListMemoryTasksResponseBodyItemsStrategyConfig struct {
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

func (s ListMemoryTasksResponseBodyItemsStrategyConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryTasksResponseBodyItemsStrategyConfig) GoString() string {
	return s.String()
}

func (s *ListMemoryTasksResponseBodyItemsStrategyConfig) GetInstruction() *string {
	return s.Instruction
}

func (s *ListMemoryTasksResponseBodyItemsStrategyConfig) GetType() *string {
	return s.Type
}

func (s *ListMemoryTasksResponseBodyItemsStrategyConfig) SetInstruction(v string) *ListMemoryTasksResponseBodyItemsStrategyConfig {
	s.Instruction = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsStrategyConfig) SetType(v string) *ListMemoryTasksResponseBodyItemsStrategyConfig {
	s.Type = &v
	return s
}

func (s *ListMemoryTasksResponseBodyItemsStrategyConfig) Validate() error {
	return dara.Validate(s)
}
