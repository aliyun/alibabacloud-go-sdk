// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAICoachTaskReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueList(v []*CreateAICoachTaskReportRequestDialogueList) *CreateAICoachTaskReportRequest
	GetDialogueList() []*CreateAICoachTaskReportRequestDialogueList
	SetIdempotentId(v string) *CreateAICoachTaskReportRequest
	GetIdempotentId() *string
	SetTaskId(v string) *CreateAICoachTaskReportRequest
	GetTaskId() *string
}

type CreateAICoachTaskReportRequest struct {
	DialogueList []*CreateAICoachTaskReportRequestDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 123456789
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	// example:
	//
	// 874890065171169280
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateAICoachTaskReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskReportRequest) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskReportRequest) GetDialogueList() []*CreateAICoachTaskReportRequestDialogueList {
	return s.DialogueList
}

func (s *CreateAICoachTaskReportRequest) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *CreateAICoachTaskReportRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAICoachTaskReportRequest) SetDialogueList(v []*CreateAICoachTaskReportRequestDialogueList) *CreateAICoachTaskReportRequest {
	s.DialogueList = v
	return s
}

func (s *CreateAICoachTaskReportRequest) SetIdempotentId(v string) *CreateAICoachTaskReportRequest {
	s.IdempotentId = &v
	return s
}

func (s *CreateAICoachTaskReportRequest) SetTaskId(v string) *CreateAICoachTaskReportRequest {
	s.TaskId = &v
	return s
}

func (s *CreateAICoachTaskReportRequest) Validate() error {
	if s.DialogueList != nil {
		for _, item := range s.DialogueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAICoachTaskReportRequestDialogueList struct {
	// example:
	//
	// hello
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// coach
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s CreateAICoachTaskReportRequestDialogueList) String() string {
	return dara.Prettify(s)
}

func (s CreateAICoachTaskReportRequestDialogueList) GoString() string {
	return s.String()
}

func (s *CreateAICoachTaskReportRequestDialogueList) GetMessage() *string {
	return s.Message
}

func (s *CreateAICoachTaskReportRequestDialogueList) GetRole() *string {
	return s.Role
}

func (s *CreateAICoachTaskReportRequestDialogueList) SetMessage(v string) *CreateAICoachTaskReportRequestDialogueList {
	s.Message = &v
	return s
}

func (s *CreateAICoachTaskReportRequestDialogueList) SetRole(v string) *CreateAICoachTaskReportRequestDialogueList {
	s.Role = &v
	return s
}

func (s *CreateAICoachTaskReportRequestDialogueList) Validate() error {
	return dara.Validate(s)
}
