// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateTaskRequest
	GetAppId() *string
	SetProcessId(v string) *CreateTaskRequest
	GetProcessId() *string
	SetTaskName(v string) *CreateTaskRequest
	GetTaskName() *string
	SetVariables(v map[string]interface{}) *CreateTaskRequest
	GetVariables() map[string]interface{}
}

type CreateTaskRequest struct {
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
	TaskName  *string                `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Variables map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateTaskRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *CreateTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTaskRequest) GetVariables() map[string]interface{} {
	return s.Variables
}

func (s *CreateTaskRequest) SetAppId(v string) *CreateTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateTaskRequest) SetProcessId(v string) *CreateTaskRequest {
	s.ProcessId = &v
	return s
}

func (s *CreateTaskRequest) SetTaskName(v string) *CreateTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTaskRequest) SetVariables(v map[string]interface{}) *CreateTaskRequest {
	s.Variables = v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	return dara.Validate(s)
}
