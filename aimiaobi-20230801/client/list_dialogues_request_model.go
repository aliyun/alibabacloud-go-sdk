// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialoguesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListDialoguesRequest
	GetAgentKey() *string
	SetCurrent(v int32) *ListDialoguesRequest
	GetCurrent() *int32
	SetDialogueType(v int32) *ListDialoguesRequest
	GetDialogueType() *int32
	SetEndTime(v string) *ListDialoguesRequest
	GetEndTime() *string
	SetSize(v int32) *ListDialoguesRequest
	GetSize() *int32
	SetStartTime(v string) *ListDialoguesRequest
	GetStartTime() *string
	SetTaskId(v string) *ListDialoguesRequest
	GetTaskId() *string
}

type ListDialoguesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 2
	DialogueType *int32 `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 2024-01-04 11:46:07
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListDialoguesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDialoguesRequest) GoString() string {
	return s.String()
}

func (s *ListDialoguesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListDialoguesRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListDialoguesRequest) GetDialogueType() *int32 {
	return s.DialogueType
}

func (s *ListDialoguesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDialoguesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListDialoguesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListDialoguesRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListDialoguesRequest) SetAgentKey(v string) *ListDialoguesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDialoguesRequest) SetCurrent(v int32) *ListDialoguesRequest {
	s.Current = &v
	return s
}

func (s *ListDialoguesRequest) SetDialogueType(v int32) *ListDialoguesRequest {
	s.DialogueType = &v
	return s
}

func (s *ListDialoguesRequest) SetEndTime(v string) *ListDialoguesRequest {
	s.EndTime = &v
	return s
}

func (s *ListDialoguesRequest) SetSize(v int32) *ListDialoguesRequest {
	s.Size = &v
	return s
}

func (s *ListDialoguesRequest) SetStartTime(v string) *ListDialoguesRequest {
	s.StartTime = &v
	return s
}

func (s *ListDialoguesRequest) SetTaskId(v string) *ListDialoguesRequest {
	s.TaskId = &v
	return s
}

func (s *ListDialoguesRequest) Validate() error {
	return dara.Validate(s)
}
