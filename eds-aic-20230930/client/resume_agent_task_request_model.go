// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalPrompt(v string) *ResumeAgentTaskRequest
	GetAdditionalPrompt() *string
	SetTaskIds(v []*string) *ResumeAgentTaskRequest
	GetTaskIds() []*string
}

type ResumeAgentTaskRequest struct {
	AdditionalPrompt *string `json:"AdditionalPrompt,omitempty" xml:"AdditionalPrompt,omitempty"`
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s ResumeAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *ResumeAgentTaskRequest) GetAdditionalPrompt() *string {
	return s.AdditionalPrompt
}

func (s *ResumeAgentTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ResumeAgentTaskRequest) SetAdditionalPrompt(v string) *ResumeAgentTaskRequest {
	s.AdditionalPrompt = &v
	return s
}

func (s *ResumeAgentTaskRequest) SetTaskIds(v []*string) *ResumeAgentTaskRequest {
	s.TaskIds = v
	return s
}

func (s *ResumeAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
