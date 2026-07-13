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
	SetClarificationAnswers(v []*ResumeAgentTaskRequestClarificationAnswers) *ResumeAgentTaskRequest
	GetClarificationAnswers() []*ResumeAgentTaskRequestClarificationAnswers
	SetTaskIds(v []*string) *ResumeAgentTaskRequest
	GetTaskIds() []*string
	SetToolCallId(v string) *ResumeAgentTaskRequest
	GetToolCallId() *string
}

type ResumeAgentTaskRequest struct {
	// The additional prompt to append. This parameter takes effect only when the task is passively paused, such as when the task is paused and waiting for user confirmation.
	//
	// example:
	//
	// 验证码为***。
	AdditionalPrompt     *string                                       `json:"AdditionalPrompt,omitempty" xml:"AdditionalPrompt,omitempty"`
	ClarificationAnswers []*ResumeAgentTaskRequestClarificationAnswers `json:"ClarificationAnswers,omitempty" xml:"ClarificationAnswers,omitempty" type:"Repeated"`
	// The list of task IDs.
	//
	// This parameter is required.
	TaskIds    []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	ToolCallId *string   `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
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

func (s *ResumeAgentTaskRequest) GetClarificationAnswers() []*ResumeAgentTaskRequestClarificationAnswers {
	return s.ClarificationAnswers
}

func (s *ResumeAgentTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ResumeAgentTaskRequest) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *ResumeAgentTaskRequest) SetAdditionalPrompt(v string) *ResumeAgentTaskRequest {
	s.AdditionalPrompt = &v
	return s
}

func (s *ResumeAgentTaskRequest) SetClarificationAnswers(v []*ResumeAgentTaskRequestClarificationAnswers) *ResumeAgentTaskRequest {
	s.ClarificationAnswers = v
	return s
}

func (s *ResumeAgentTaskRequest) SetTaskIds(v []*string) *ResumeAgentTaskRequest {
	s.TaskIds = v
	return s
}

func (s *ResumeAgentTaskRequest) SetToolCallId(v string) *ResumeAgentTaskRequest {
	s.ToolCallId = &v
	return s
}

func (s *ResumeAgentTaskRequest) Validate() error {
	if s.ClarificationAnswers != nil {
		for _, item := range s.ClarificationAnswers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ResumeAgentTaskRequestClarificationAnswers struct {
	CustomValue *string   `json:"CustomValue,omitempty" xml:"CustomValue,omitempty"`
	Id          *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	SelectedIds []*string `json:"SelectedIds,omitempty" xml:"SelectedIds,omitempty" type:"Repeated"`
}

func (s ResumeAgentTaskRequestClarificationAnswers) String() string {
	return dara.Prettify(s)
}

func (s ResumeAgentTaskRequestClarificationAnswers) GoString() string {
	return s.String()
}

func (s *ResumeAgentTaskRequestClarificationAnswers) GetCustomValue() *string {
	return s.CustomValue
}

func (s *ResumeAgentTaskRequestClarificationAnswers) GetId() *string {
	return s.Id
}

func (s *ResumeAgentTaskRequestClarificationAnswers) GetSelectedIds() []*string {
	return s.SelectedIds
}

func (s *ResumeAgentTaskRequestClarificationAnswers) SetCustomValue(v string) *ResumeAgentTaskRequestClarificationAnswers {
	s.CustomValue = &v
	return s
}

func (s *ResumeAgentTaskRequestClarificationAnswers) SetId(v string) *ResumeAgentTaskRequestClarificationAnswers {
	s.Id = &v
	return s
}

func (s *ResumeAgentTaskRequestClarificationAnswers) SetSelectedIds(v []*string) *ResumeAgentTaskRequestClarificationAnswers {
	s.SelectedIds = v
	return s
}

func (s *ResumeAgentTaskRequestClarificationAnswers) Validate() error {
	return dara.Validate(s)
}
