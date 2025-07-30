// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebuggerResult interface {
	dara.Model
	String() string
	GoString() string
	SetDebuggerConfigContent(v string) *DebuggerResult
	GetDebuggerConfigContent() *string
	SetDebuggerJobIssues(v string) *DebuggerResult
	GetDebuggerJobIssues() *string
	SetDebuggerJobStatus(v string) *DebuggerResult
	GetDebuggerJobStatus() *string
	SetDebuggerReportURL(v string) *DebuggerResult
	GetDebuggerReportURL() *string
	SetJobDisplayName(v string) *DebuggerResult
	GetJobDisplayName() *string
	SetJobId(v string) *DebuggerResult
	GetJobId() *string
	SetJobUserId(v string) *DebuggerResult
	GetJobUserId() *string
}

type DebuggerResult struct {
	// example:
	//
	// {\"description\":\"这是一个新的pytorchjob模板\"}
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// example:
	//
	// { "ProfileReport": { "Name": "CPUBottleneck","Triggered": 10,"Violations": 2,"Details": "{}"}, "LowCPU": { "Name": "CPUBottleneck","Triggered": 10,"Violations": 2,"Details": "{}"}}
	DebuggerJobIssues *string `json:"DebuggerJobIssues,omitempty" xml:"DebuggerJobIssues,omitempty"`
	// example:
	//
	// {"Running": 1, "Failed": 1, "Succeeded": 2}
	DebuggerJobStatus *string `json:"DebuggerJobStatus,omitempty" xml:"DebuggerJobStatus,omitempty"`
	// example:
	//
	// http://xxx.com/debug/report/download/new_xxxx.html
	DebuggerReportURL *string `json:"DebuggerReportURL,omitempty" xml:"DebuggerReportURL,omitempty"`
	// example:
	//
	// dlc debugger test
	JobDisplayName *string `json:"JobDisplayName,omitempty" xml:"JobDisplayName,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 12344556
	JobUserId *string `json:"JobUserId,omitempty" xml:"JobUserId,omitempty"`
}

func (s DebuggerResult) String() string {
	return dara.Prettify(s)
}

func (s DebuggerResult) GoString() string {
	return s.String()
}

func (s *DebuggerResult) GetDebuggerConfigContent() *string {
	return s.DebuggerConfigContent
}

func (s *DebuggerResult) GetDebuggerJobIssues() *string {
	return s.DebuggerJobIssues
}

func (s *DebuggerResult) GetDebuggerJobStatus() *string {
	return s.DebuggerJobStatus
}

func (s *DebuggerResult) GetDebuggerReportURL() *string {
	return s.DebuggerReportURL
}

func (s *DebuggerResult) GetJobDisplayName() *string {
	return s.JobDisplayName
}

func (s *DebuggerResult) GetJobId() *string {
	return s.JobId
}

func (s *DebuggerResult) GetJobUserId() *string {
	return s.JobUserId
}

func (s *DebuggerResult) SetDebuggerConfigContent(v string) *DebuggerResult {
	s.DebuggerConfigContent = &v
	return s
}

func (s *DebuggerResult) SetDebuggerJobIssues(v string) *DebuggerResult {
	s.DebuggerJobIssues = &v
	return s
}

func (s *DebuggerResult) SetDebuggerJobStatus(v string) *DebuggerResult {
	s.DebuggerJobStatus = &v
	return s
}

func (s *DebuggerResult) SetDebuggerReportURL(v string) *DebuggerResult {
	s.DebuggerReportURL = &v
	return s
}

func (s *DebuggerResult) SetJobDisplayName(v string) *DebuggerResult {
	s.JobDisplayName = &v
	return s
}

func (s *DebuggerResult) SetJobId(v string) *DebuggerResult {
	s.JobId = &v
	return s
}

func (s *DebuggerResult) SetJobUserId(v string) *DebuggerResult {
	s.JobUserId = &v
	return s
}

func (s *DebuggerResult) Validate() error {
	return dara.Validate(s)
}
