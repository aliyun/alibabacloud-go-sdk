// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerSSERequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*AnswerSSERequestMessages) *AnswerSSERequest
	GetMessages() []*AnswerSSERequestMessages
	SetParameters(v *AnswerSSERequestParameters) *AnswerSSERequest
	GetParameters() *AnswerSSERequestParameters
	SetWorkspaceId(v string) *AnswerSSERequest
	GetWorkspaceId() *string
}

type AnswerSSERequest struct {
	Messages   []*AnswerSSERequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	Parameters *AnswerSSERequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// llm-1ijrzuv3v0ivvls7
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s AnswerSSERequest) String() string {
	return dara.Prettify(s)
}

func (s AnswerSSERequest) GoString() string {
	return s.String()
}

func (s *AnswerSSERequest) GetMessages() []*AnswerSSERequestMessages {
	return s.Messages
}

func (s *AnswerSSERequest) GetParameters() *AnswerSSERequestParameters {
	return s.Parameters
}

func (s *AnswerSSERequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AnswerSSERequest) SetMessages(v []*AnswerSSERequestMessages) *AnswerSSERequest {
	s.Messages = v
	return s
}

func (s *AnswerSSERequest) SetParameters(v *AnswerSSERequestParameters) *AnswerSSERequest {
	s.Parameters = v
	return s
}

func (s *AnswerSSERequest) SetWorkspaceId(v string) *AnswerSSERequest {
	s.WorkspaceId = &v
	return s
}

func (s *AnswerSSERequest) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AnswerSSERequestMessages struct {
	Content []map[string]*string `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s AnswerSSERequestMessages) String() string {
	return dara.Prettify(s)
}

func (s AnswerSSERequestMessages) GoString() string {
	return s.String()
}

func (s *AnswerSSERequestMessages) GetContent() []map[string]*string {
	return s.Content
}

func (s *AnswerSSERequestMessages) GetRole() *string {
	return s.Role
}

func (s *AnswerSSERequestMessages) SetContent(v []map[string]*string) *AnswerSSERequestMessages {
	s.Content = v
	return s
}

func (s *AnswerSSERequestMessages) SetRole(v string) *AnswerSSERequestMessages {
	s.Role = &v
	return s
}

func (s *AnswerSSERequestMessages) Validate() error {
	return dara.Validate(s)
}

type AnswerSSERequestParameters struct {
	// example:
	//
	// 6
	Grade *int32 `json:"grade,omitempty" xml:"grade,omitempty"`
	// example:
	//
	// other
	Stage *string `json:"stage,omitempty" xml:"stage,omitempty"`
	// example:
	//
	// other
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s AnswerSSERequestParameters) String() string {
	return dara.Prettify(s)
}

func (s AnswerSSERequestParameters) GoString() string {
	return s.String()
}

func (s *AnswerSSERequestParameters) GetGrade() *int32 {
	return s.Grade
}

func (s *AnswerSSERequestParameters) GetStage() *string {
	return s.Stage
}

func (s *AnswerSSERequestParameters) GetSubject() *string {
	return s.Subject
}

func (s *AnswerSSERequestParameters) SetGrade(v int32) *AnswerSSERequestParameters {
	s.Grade = &v
	return s
}

func (s *AnswerSSERequestParameters) SetStage(v string) *AnswerSSERequestParameters {
	s.Stage = &v
	return s
}

func (s *AnswerSSERequestParameters) SetSubject(v string) *AnswerSSERequestParameters {
	s.Subject = &v
	return s
}

func (s *AnswerSSERequestParameters) Validate() error {
	return dara.Validate(s)
}
