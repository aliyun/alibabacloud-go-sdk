// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAgentStatusLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListLegacyAgentStatusLogsRequest
	GetAgentId() *string
	SetEndTime(v int64) *ListLegacyAgentStatusLogsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListLegacyAgentStatusLogsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListLegacyAgentStatusLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLegacyAgentStatusLogsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListLegacyAgentStatusLogsRequest
	GetStartTime() *int64
}

type ListLegacyAgentStatusLogsRequest struct {
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1657879880010
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1657778840011
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLegacyAgentStatusLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentStatusLogsRequest) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentStatusLogsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListLegacyAgentStatusLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLegacyAgentStatusLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLegacyAgentStatusLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyAgentStatusLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyAgentStatusLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListLegacyAgentStatusLogsRequest) SetAgentId(v string) *ListLegacyAgentStatusLogsRequest {
	s.AgentId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsRequest) SetEndTime(v int64) *ListLegacyAgentStatusLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLegacyAgentStatusLogsRequest) SetInstanceId(v string) *ListLegacyAgentStatusLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLegacyAgentStatusLogsRequest) SetPageNumber(v int32) *ListLegacyAgentStatusLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyAgentStatusLogsRequest) SetPageSize(v int32) *ListLegacyAgentStatusLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLegacyAgentStatusLogsRequest) SetStartTime(v int64) *ListLegacyAgentStatusLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLegacyAgentStatusLogsRequest) Validate() error {
	return dara.Validate(s)
}
