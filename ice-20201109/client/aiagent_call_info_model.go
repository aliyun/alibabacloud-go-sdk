// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIAgentCallInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCallDuration(v int32) *AIAgentCallInfo
	GetCallDuration() *int32
	SetCallEndTime(v string) *AIAgentCallInfo
	GetCallEndTime() *string
	SetCallStartTime(v string) *AIAgentCallInfo
	GetCallStartTime() *string
	SetCalleeNumber(v string) *AIAgentCallInfo
	GetCalleeNumber() *string
	SetCallerNumber(v string) *AIAgentCallInfo
	GetCallerNumber() *string
	SetHangupRole(v int32) *AIAgentCallInfo
	GetHangupRole() *int32
	SetStatus(v string) *AIAgentCallInfo
	GetStatus() *string
}

type AIAgentCallInfo struct {
	// The duration of the call, in seconds.
	//
	// example:
	//
	// 5
	CallDuration *int32 `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	// The time the call ended, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-01T16:53:58.875932+00:00
	CallEndTime *string `json:"CallEndTime,omitempty" xml:"CallEndTime,omitempty"`
	// The time the call started, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-01T16:53:53.184797+00:00
	CallStartTime *string `json:"CallStartTime,omitempty" xml:"CallStartTime,omitempty"`
	// The number of the called party.
	//
	// example:
	//
	// 136******794
	CalleeNumber *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	// The number of the calling party.
	//
	// example:
	//
	// 183*****333
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// Indicates which party ended the call.
	//
	// 0: The agent ended the call.
	//
	// 1: The user ended the call.
	//
	// 2: The system ended the call for a transfer.
	//
	// example:
	//
	// 0
	HangupRole *int32 `json:"HangupRole,omitempty" xml:"HangupRole,omitempty"`
	// The status of the call.
	//
	// example:
	//
	// 4
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AIAgentCallInfo) String() string {
	return dara.Prettify(s)
}

func (s AIAgentCallInfo) GoString() string {
	return s.String()
}

func (s *AIAgentCallInfo) GetCallDuration() *int32 {
	return s.CallDuration
}

func (s *AIAgentCallInfo) GetCallEndTime() *string {
	return s.CallEndTime
}

func (s *AIAgentCallInfo) GetCallStartTime() *string {
	return s.CallStartTime
}

func (s *AIAgentCallInfo) GetCalleeNumber() *string {
	return s.CalleeNumber
}

func (s *AIAgentCallInfo) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *AIAgentCallInfo) GetHangupRole() *int32 {
	return s.HangupRole
}

func (s *AIAgentCallInfo) GetStatus() *string {
	return s.Status
}

func (s *AIAgentCallInfo) SetCallDuration(v int32) *AIAgentCallInfo {
	s.CallDuration = &v
	return s
}

func (s *AIAgentCallInfo) SetCallEndTime(v string) *AIAgentCallInfo {
	s.CallEndTime = &v
	return s
}

func (s *AIAgentCallInfo) SetCallStartTime(v string) *AIAgentCallInfo {
	s.CallStartTime = &v
	return s
}

func (s *AIAgentCallInfo) SetCalleeNumber(v string) *AIAgentCallInfo {
	s.CalleeNumber = &v
	return s
}

func (s *AIAgentCallInfo) SetCallerNumber(v string) *AIAgentCallInfo {
	s.CallerNumber = &v
	return s
}

func (s *AIAgentCallInfo) SetHangupRole(v int32) *AIAgentCallInfo {
	s.HangupRole = &v
	return s
}

func (s *AIAgentCallInfo) SetStatus(v string) *AIAgentCallInfo {
	s.Status = &v
	return s
}

func (s *AIAgentCallInfo) Validate() error {
	return dara.Validate(s)
}
