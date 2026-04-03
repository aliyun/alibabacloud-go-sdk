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
	CallDuration  *int32  `json:"CallDuration,omitempty" xml:"CallDuration,omitempty"`
	CallEndTime   *string `json:"CallEndTime,omitempty" xml:"CallEndTime,omitempty"`
	CallStartTime *string `json:"CallStartTime,omitempty" xml:"CallStartTime,omitempty"`
	CalleeNumber  *string `json:"CalleeNumber,omitempty" xml:"CalleeNumber,omitempty"`
	CallerNumber  *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	HangupRole    *int32  `json:"HangupRole,omitempty" xml:"HangupRole,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
