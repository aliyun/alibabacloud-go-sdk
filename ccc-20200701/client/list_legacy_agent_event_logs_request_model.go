// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLegacyAgentEventLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListLegacyAgentEventLogsRequest
	GetAgentId() *string
	SetEndTime(v int64) *ListLegacyAgentEventLogsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListLegacyAgentEventLogsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListLegacyAgentEventLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLegacyAgentEventLogsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListLegacyAgentEventLogsRequest
	GetStartTime() *int64
}

type ListLegacyAgentEventLogsRequest struct {
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1658026440011
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
	// 1657853640015
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListLegacyAgentEventLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLegacyAgentEventLogsRequest) GoString() string {
	return s.String()
}

func (s *ListLegacyAgentEventLogsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListLegacyAgentEventLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListLegacyAgentEventLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLegacyAgentEventLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLegacyAgentEventLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLegacyAgentEventLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListLegacyAgentEventLogsRequest) SetAgentId(v string) *ListLegacyAgentEventLogsRequest {
	s.AgentId = &v
	return s
}

func (s *ListLegacyAgentEventLogsRequest) SetEndTime(v int64) *ListLegacyAgentEventLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListLegacyAgentEventLogsRequest) SetInstanceId(v string) *ListLegacyAgentEventLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLegacyAgentEventLogsRequest) SetPageNumber(v int32) *ListLegacyAgentEventLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLegacyAgentEventLogsRequest) SetPageSize(v int32) *ListLegacyAgentEventLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLegacyAgentEventLogsRequest) SetStartTime(v int64) *ListLegacyAgentEventLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListLegacyAgentEventLogsRequest) Validate() error {
	return dara.Validate(s)
}
