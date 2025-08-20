// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatementInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StatementInfo
	GetCode() *string
	SetCompletedTimeInMills(v int64) *StatementInfo
	GetCompletedTimeInMills() *int64
	SetOutput(v string) *StatementInfo
	GetOutput() *string
	SetProcess(v float32) *StatementInfo
	GetProcess() *float32
	SetStartedTimeInMills(v int64) *StatementInfo
	GetStartedTimeInMills() *int64
	SetState(v string) *StatementInfo
	GetState() *string
	SetStatementId(v string) *StatementInfo
	GetStatementId() *string
}

type StatementInfo struct {
	Code                 *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	CompletedTimeInMills *int64   `json:"CompletedTimeInMills,omitempty" xml:"CompletedTimeInMills,omitempty"`
	Output               *string  `json:"Output,omitempty" xml:"Output,omitempty"`
	Process              *float32 `json:"Process,omitempty" xml:"Process,omitempty"`
	StartedTimeInMills   *int64   `json:"StartedTimeInMills,omitempty" xml:"StartedTimeInMills,omitempty"`
	State                *string  `json:"State,omitempty" xml:"State,omitempty"`
	StatementId          *string  `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s StatementInfo) String() string {
	return dara.Prettify(s)
}

func (s StatementInfo) GoString() string {
	return s.String()
}

func (s *StatementInfo) GetCode() *string {
	return s.Code
}

func (s *StatementInfo) GetCompletedTimeInMills() *int64 {
	return s.CompletedTimeInMills
}

func (s *StatementInfo) GetOutput() *string {
	return s.Output
}

func (s *StatementInfo) GetProcess() *float32 {
	return s.Process
}

func (s *StatementInfo) GetStartedTimeInMills() *int64 {
	return s.StartedTimeInMills
}

func (s *StatementInfo) GetState() *string {
	return s.State
}

func (s *StatementInfo) GetStatementId() *string {
	return s.StatementId
}

func (s *StatementInfo) SetCode(v string) *StatementInfo {
	s.Code = &v
	return s
}

func (s *StatementInfo) SetCompletedTimeInMills(v int64) *StatementInfo {
	s.CompletedTimeInMills = &v
	return s
}

func (s *StatementInfo) SetOutput(v string) *StatementInfo {
	s.Output = &v
	return s
}

func (s *StatementInfo) SetProcess(v float32) *StatementInfo {
	s.Process = &v
	return s
}

func (s *StatementInfo) SetStartedTimeInMills(v int64) *StatementInfo {
	s.StartedTimeInMills = &v
	return s
}

func (s *StatementInfo) SetState(v string) *StatementInfo {
	s.State = &v
	return s
}

func (s *StatementInfo) SetStatementId(v string) *StatementInfo {
	s.StatementId = &v
	return s
}

func (s *StatementInfo) Validate() error {
	return dara.Validate(s)
}
