// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICoachTaskSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAICoachTaskSessionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAICoachTaskSessionRequest
	GetPageSize() *int32
	SetSessionId(v string) *ListAICoachTaskSessionRequest
	GetSessionId() *string
	SetTaskId(v string) *ListAICoachTaskSessionRequest
	GetTaskId() *string
}

type ListAICoachTaskSessionRequest struct {
	PageNumber *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SessionId  *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	TaskId     *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListAICoachTaskSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAICoachTaskSessionRequest) GoString() string {
	return s.String()
}

func (s *ListAICoachTaskSessionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAICoachTaskSessionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAICoachTaskSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAICoachTaskSessionRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAICoachTaskSessionRequest) SetPageNumber(v int32) *ListAICoachTaskSessionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICoachTaskSessionRequest) SetPageSize(v int32) *ListAICoachTaskSessionRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICoachTaskSessionRequest) SetSessionId(v string) *ListAICoachTaskSessionRequest {
	s.SessionId = &v
	return s
}

func (s *ListAICoachTaskSessionRequest) SetTaskId(v string) *ListAICoachTaskSessionRequest {
	s.TaskId = &v
	return s
}

func (s *ListAICoachTaskSessionRequest) Validate() error {
	return dara.Validate(s)
}
