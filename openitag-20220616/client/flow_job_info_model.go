// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowJobInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDisplay(v bool) *FlowJobInfo
	GetDisplay() *bool
	SetJobId(v string) *FlowJobInfo
	GetJobId() *string
	SetJobType(v string) *FlowJobInfo
	GetJobType() *string
	SetMessageId(v string) *FlowJobInfo
	GetMessageId() *string
	SetProcessType(v string) *FlowJobInfo
	GetProcessType() *string
	SetTaskId(v string) *FlowJobInfo
	GetTaskId() *string
}

type FlowJobInfo struct {
	// Whether to display. Possible values are:
	//
	// - true: Display.
	//
	// - false: Do not display.
	//
	// example:
	//
	// true
	Display *bool `json:"Display,omitempty" xml:"Display,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 1475***441221943296
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Task Type. Currently, only DOWNLOWD_MARKRESULT_FLOW is supported.
	//
	// example:
	//
	// DOWNLOWD_MARKRESULT_FLOW
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// Message ID.
	//
	// example:
	//
	// 792B76F4000E681A95155146A002D5F8
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Processing information.
	//
	// example:
	//
	// NEW_INIT
	ProcessType *string `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// Job ID.
	//
	// example:
	//
	// 1543***518306500608
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FlowJobInfo) String() string {
	return dara.Prettify(s)
}

func (s FlowJobInfo) GoString() string {
	return s.String()
}

func (s *FlowJobInfo) GetDisplay() *bool {
	return s.Display
}

func (s *FlowJobInfo) GetJobId() *string {
	return s.JobId
}

func (s *FlowJobInfo) GetJobType() *string {
	return s.JobType
}

func (s *FlowJobInfo) GetMessageId() *string {
	return s.MessageId
}

func (s *FlowJobInfo) GetProcessType() *string {
	return s.ProcessType
}

func (s *FlowJobInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *FlowJobInfo) SetDisplay(v bool) *FlowJobInfo {
	s.Display = &v
	return s
}

func (s *FlowJobInfo) SetJobId(v string) *FlowJobInfo {
	s.JobId = &v
	return s
}

func (s *FlowJobInfo) SetJobType(v string) *FlowJobInfo {
	s.JobType = &v
	return s
}

func (s *FlowJobInfo) SetMessageId(v string) *FlowJobInfo {
	s.MessageId = &v
	return s
}

func (s *FlowJobInfo) SetProcessType(v string) *FlowJobInfo {
	s.ProcessType = &v
	return s
}

func (s *FlowJobInfo) SetTaskId(v string) *FlowJobInfo {
	s.TaskId = &v
	return s
}

func (s *FlowJobInfo) Validate() error {
	return dara.Validate(s)
}
