// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateTaskShrinkRequest
	GetAppId() *string
	SetProcessId(v string) *CreateTaskShrinkRequest
	GetProcessId() *string
	SetTaskName(v string) *CreateTaskShrinkRequest
	GetTaskName() *string
	SetVariablesShrink(v string) *CreateTaskShrinkRequest
	GetVariablesShrink() *string
}

type CreateTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// JXT3MLZJ56**65D6
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// P3KTG43SW7**6A17I
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateTaskShrinkRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *CreateTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTaskShrinkRequest) GetVariablesShrink() *string {
	return s.VariablesShrink
}

func (s *CreateTaskShrinkRequest) SetAppId(v string) *CreateTaskShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateTaskShrinkRequest) SetProcessId(v string) *CreateTaskShrinkRequest {
	s.ProcessId = &v
	return s
}

func (s *CreateTaskShrinkRequest) SetTaskName(v string) *CreateTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTaskShrinkRequest) SetVariablesShrink(v string) *CreateTaskShrinkRequest {
	s.VariablesShrink = &v
	return s
}

func (s *CreateTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
