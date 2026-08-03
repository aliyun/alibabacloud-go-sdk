// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachScriptGenerateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetAICoachScriptGenerateTaskRequest
	GetTaskId() *string
}

type GetAICoachScriptGenerateTaskRequest struct {
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetAICoachScriptGenerateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachScriptGenerateTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAICoachScriptGenerateTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAICoachScriptGenerateTaskRequest) SetTaskId(v string) *GetAICoachScriptGenerateTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAICoachScriptGenerateTaskRequest) Validate() error {
	return dara.Validate(s)
}
