// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebuggerJobIssue interface {
	dara.Model
	String() string
	GoString() string
	SetDebuggerJobIssue(v string) *DebuggerJobIssue
	GetDebuggerJobIssue() *string
	SetGmtCreateTime(v string) *DebuggerJobIssue
	GetGmtCreateTime() *string
	SetJobDebuggerIssueId(v string) *DebuggerJobIssue
	GetJobDebuggerIssueId() *string
	SetJobId(v string) *DebuggerJobIssue
	GetJobId() *string
	SetReasonCode(v string) *DebuggerJobIssue
	GetReasonCode() *string
	SetReasonMessage(v string) *DebuggerJobIssue
	GetReasonMessage() *string
	SetRuleName(v string) *DebuggerJobIssue
	GetRuleName() *string
}

type DebuggerJobIssue struct {
	// example:
	//
	// {"Name": "CPUBottleneck",  "Triggered": 10, "Violations": 2,  "Details": "{}"}
	DebuggerJobIssue *string `json:"DebuggerJobIssue,omitempty" xml:"DebuggerJobIssue,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// de-826ca1bcfba30
	JobDebuggerIssueId *string `json:"JobDebuggerIssueId,omitempty" xml:"JobDebuggerIssueId,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1002300
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// GPU利用率低
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// ProfileReport
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DebuggerJobIssue) String() string {
	return dara.Prettify(s)
}

func (s DebuggerJobIssue) GoString() string {
	return s.String()
}

func (s *DebuggerJobIssue) GetDebuggerJobIssue() *string {
	return s.DebuggerJobIssue
}

func (s *DebuggerJobIssue) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DebuggerJobIssue) GetJobDebuggerIssueId() *string {
	return s.JobDebuggerIssueId
}

func (s *DebuggerJobIssue) GetJobId() *string {
	return s.JobId
}

func (s *DebuggerJobIssue) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DebuggerJobIssue) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *DebuggerJobIssue) GetRuleName() *string {
	return s.RuleName
}

func (s *DebuggerJobIssue) SetDebuggerJobIssue(v string) *DebuggerJobIssue {
	s.DebuggerJobIssue = &v
	return s
}

func (s *DebuggerJobIssue) SetGmtCreateTime(v string) *DebuggerJobIssue {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerJobIssue) SetJobDebuggerIssueId(v string) *DebuggerJobIssue {
	s.JobDebuggerIssueId = &v
	return s
}

func (s *DebuggerJobIssue) SetJobId(v string) *DebuggerJobIssue {
	s.JobId = &v
	return s
}

func (s *DebuggerJobIssue) SetReasonCode(v string) *DebuggerJobIssue {
	s.ReasonCode = &v
	return s
}

func (s *DebuggerJobIssue) SetReasonMessage(v string) *DebuggerJobIssue {
	s.ReasonMessage = &v
	return s
}

func (s *DebuggerJobIssue) SetRuleName(v string) *DebuggerJobIssue {
	s.RuleName = &v
	return s
}

func (s *DebuggerJobIssue) Validate() error {
	return dara.Validate(s)
}
