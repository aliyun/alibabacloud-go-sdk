// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunScriptChatRequest
	GetPrompt() *string
	SetTaskId(v string) *RunScriptChatRequest
	GetTaskId() *string
}

type RunScriptChatRequest struct {
	// This parameter is required.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RunScriptChatRequest) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatRequest) GoString() string {
	return s.String()
}

func (s *RunScriptChatRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunScriptChatRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunScriptChatRequest) SetPrompt(v string) *RunScriptChatRequest {
	s.Prompt = &v
	return s
}

func (s *RunScriptChatRequest) SetTaskId(v string) *RunScriptChatRequest {
	s.TaskId = &v
	return s
}

func (s *RunScriptChatRequest) Validate() error {
	return dara.Validate(s)
}
