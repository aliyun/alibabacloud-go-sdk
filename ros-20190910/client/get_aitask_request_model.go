// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputOption(v string) *GetAITaskRequest
	GetOutputOption() *string
	SetTaskId(v string) *GetAITaskRequest
	GetTaskId() *string
}

type GetAITaskRequest struct {
	// Specifies whether to return the TaskOutput parameter. The TaskOutput parameter specifies the outputs of the AI task. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled (default)
	//
	// >  The value of TaskOutput may be excessively long. If you do not require the outputs of the task, we recommend that you set OutputOption to Disabled to improve the response speed of the API operation.
	//
	// example:
	//
	// Disabled
	OutputOption *string `json:"OutputOption,omitempty" xml:"OutputOption,omitempty"`
	// The ID of the AI task.
	//
	// example:
	//
	// t-asasas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAITaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAITaskRequest) GoString() string {
	return s.String()
}

func (s *GetAITaskRequest) GetOutputOption() *string {
	return s.OutputOption
}

func (s *GetAITaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAITaskRequest) SetOutputOption(v string) *GetAITaskRequest {
	s.OutputOption = &v
	return s
}

func (s *GetAITaskRequest) SetTaskId(v string) *GetAITaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAITaskRequest) Validate() error {
	return dara.Validate(s)
}
