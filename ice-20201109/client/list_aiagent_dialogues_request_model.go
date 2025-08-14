// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentDialoguesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListAIAgentDialoguesRequest
	GetEndTime() *int64
	SetOrder(v string) *ListAIAgentDialoguesRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListAIAgentDialoguesRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *ListAIAgentDialoguesRequest
	GetPageSize() *int32
	SetRoundLimit(v string) *ListAIAgentDialoguesRequest
	GetRoundLimit() *string
	SetSessionId(v string) *ListAIAgentDialoguesRequest
	GetSessionId() *string
	SetStartTime(v int64) *ListAIAgentDialoguesRequest
	GetStartTime() *int64
}

type ListAIAgentDialoguesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17358082464030
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RoundLimit *string `json:"RoundLimit,omitempty" xml:"RoundLimit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f27f9b9be28642a88e18****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAIAgentDialoguesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentDialoguesRequest) GoString() string {
	return s.String()
}

func (s *ListAIAgentDialoguesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAIAgentDialoguesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListAIAgentDialoguesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAIAgentDialoguesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAIAgentDialoguesRequest) GetRoundLimit() *string {
	return s.RoundLimit
}

func (s *ListAIAgentDialoguesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListAIAgentDialoguesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAIAgentDialoguesRequest) SetEndTime(v int64) *ListAIAgentDialoguesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) SetOrder(v string) *ListAIAgentDialoguesRequest {
	s.Order = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) SetPageNumber(v int64) *ListAIAgentDialoguesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) SetPageSize(v int32) *ListAIAgentDialoguesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) SetRoundLimit(v string) *ListAIAgentDialoguesRequest {
	s.RoundLimit = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) SetSessionId(v string) *ListAIAgentDialoguesRequest {
	s.SessionId = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) SetStartTime(v int64) *ListAIAgentDialoguesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAIAgentDialoguesRequest) Validate() error {
	return dara.Validate(s)
}
