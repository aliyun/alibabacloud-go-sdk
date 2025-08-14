// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *ListAIAgentInstanceRequest
	GetAIAgentId() *string
	SetEndTime(v string) *ListAIAgentInstanceRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListAIAgentInstanceRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAIAgentInstanceRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListAIAgentInstanceRequest
	GetStartTime() *string
}

type ListAIAgentInstanceRequest struct {
	// The ID of the AI agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4***
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. This parameter does not have a default value.
	//
	// example:
	//
	// 2023-01-02T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: 1. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. This parameter does not have a default value.
	//
	// example:
	//
	// 2023-01-01T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAIAgentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListAIAgentInstanceRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *ListAIAgentInstanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListAIAgentInstanceRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAIAgentInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAIAgentInstanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListAIAgentInstanceRequest) SetAIAgentId(v string) *ListAIAgentInstanceRequest {
	s.AIAgentId = &v
	return s
}

func (s *ListAIAgentInstanceRequest) SetEndTime(v string) *ListAIAgentInstanceRequest {
	s.EndTime = &v
	return s
}

func (s *ListAIAgentInstanceRequest) SetPageNumber(v int64) *ListAIAgentInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAIAgentInstanceRequest) SetPageSize(v int64) *ListAIAgentInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentInstanceRequest) SetStartTime(v string) *ListAIAgentInstanceRequest {
	s.StartTime = &v
	return s
}

func (s *ListAIAgentInstanceRequest) Validate() error {
	return dara.Validate(s)
}
