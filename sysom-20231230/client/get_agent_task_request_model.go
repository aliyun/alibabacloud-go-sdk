// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetAgentTaskRequest
	GetXDebugId() *string
	SetTaskId(v string) *GetAgentTaskRequest
	GetTaskId() *string
	SetXSysomInvokeSource(v string) *GetAgentTaskRequest
	GetXSysomInvokeSource() *string
}

type GetAgentTaskRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 42172120177e4b3abd6fabb3a6b5e2dd
	TaskId             *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAgentTaskRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetAgentTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgentTaskRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetAgentTaskRequest) SetXDebugId(v string) *GetAgentTaskRequest {
	s.XDebugId = &v
	return s
}

func (s *GetAgentTaskRequest) SetTaskId(v string) *GetAgentTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAgentTaskRequest) SetXSysomInvokeSource(v string) *GetAgentTaskRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
