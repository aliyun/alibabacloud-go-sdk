// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskInvocation interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceID(v string) *TaskInvocation
	GetInstanceID() *string
	SetInvocationID(v string) *TaskInvocation
	GetInvocationID() *string
	SetInvocationTarget(v string) *TaskInvocation
	GetInvocationTarget() *string
	SetOutput(v string) *TaskInvocation
	GetOutput() *string
	SetRequestID(v string) *TaskInvocation
	GetRequestID() *string
	SetSlsLogStore(v string) *TaskInvocation
	GetSlsLogStore() *string
	SetSlsProject(v string) *TaskInvocation
	GetSlsProject() *string
	SetStatus(v string) *TaskInvocation
	GetStatus() *string
}

type TaskInvocation struct {
	// example:
	//
	// c-nkj8shz7xxxx
	InstanceID *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	// example:
	//
	// E099843B-10A2-4936-9964-4E0EE263D564
	InvocationID *string `json:"invocationID,omitempty" xml:"invocationID,omitempty"`
	// example:
	//
	// acs:fc:cn-hangzhou:143xxxx:services/xxx.LATEST/functions/xxx
	InvocationTarget *string `json:"invocationTarget,omitempty" xml:"invocationTarget,omitempty"`
	// example:
	//
	// {"key1":"value1","key2":"value2"}
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// example:
	//
	// 1B3058B1-F1C9-457C-B95C-2C250A4B3118
	RequestID *string `json:"requestID,omitempty" xml:"requestID,omitempty"`
	// example:
	//
	// my-sls-logstore
	SlsLogStore *string `json:"slsLogStore,omitempty" xml:"slsLogStore,omitempty"`
	// example:
	//
	// my-sls-project
	SlsProject *string `json:"slsProject,omitempty" xml:"slsProject,omitempty"`
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TaskInvocation) String() string {
	return dara.Prettify(s)
}

func (s TaskInvocation) GoString() string {
	return s.String()
}

func (s *TaskInvocation) GetInstanceID() *string {
	return s.InstanceID
}

func (s *TaskInvocation) GetInvocationID() *string {
	return s.InvocationID
}

func (s *TaskInvocation) GetInvocationTarget() *string {
	return s.InvocationTarget
}

func (s *TaskInvocation) GetOutput() *string {
	return s.Output
}

func (s *TaskInvocation) GetRequestID() *string {
	return s.RequestID
}

func (s *TaskInvocation) GetSlsLogStore() *string {
	return s.SlsLogStore
}

func (s *TaskInvocation) GetSlsProject() *string {
	return s.SlsProject
}

func (s *TaskInvocation) GetStatus() *string {
	return s.Status
}

func (s *TaskInvocation) SetInstanceID(v string) *TaskInvocation {
	s.InstanceID = &v
	return s
}

func (s *TaskInvocation) SetInvocationID(v string) *TaskInvocation {
	s.InvocationID = &v
	return s
}

func (s *TaskInvocation) SetInvocationTarget(v string) *TaskInvocation {
	s.InvocationTarget = &v
	return s
}

func (s *TaskInvocation) SetOutput(v string) *TaskInvocation {
	s.Output = &v
	return s
}

func (s *TaskInvocation) SetRequestID(v string) *TaskInvocation {
	s.RequestID = &v
	return s
}

func (s *TaskInvocation) SetSlsLogStore(v string) *TaskInvocation {
	s.SlsLogStore = &v
	return s
}

func (s *TaskInvocation) SetSlsProject(v string) *TaskInvocation {
	s.SlsProject = &v
	return s
}

func (s *TaskInvocation) SetStatus(v string) *TaskInvocation {
	s.Status = &v
	return s
}

func (s *TaskInvocation) Validate() error {
	return dara.Validate(s)
}
