// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStateLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListAgentStateLogsRequest
	GetAgentId() *string
	SetEndTime(v int64) *ListAgentStateLogsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListAgentStateLogsRequest
	GetInstanceId() *string
	SetStartTime(v int64) *ListAgentStateLogsRequest
	GetStartTime() *int64
}

type ListAgentStateLogsRequest struct {
	// Agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// End UNIX timestamp. The default value is the current time. The time difference between EndTime and StartTime must not exceed 7 days. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1620273600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Start UNIX timestamp. The default value is the start time of the current day. The earliest allowed value is 180 days before the current date. The format is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1620230400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAgentStateLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStateLogsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentStateLogsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentStateLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAgentStateLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentStateLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAgentStateLogsRequest) SetAgentId(v string) *ListAgentStateLogsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAgentStateLogsRequest) SetEndTime(v int64) *ListAgentStateLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAgentStateLogsRequest) SetInstanceId(v string) *ListAgentStateLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentStateLogsRequest) SetStartTime(v int64) *ListAgentStateLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAgentStateLogsRequest) Validate() error {
	return dara.Validate(s)
}
