// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualAnswerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*ContextualFile) *ContextualAnswerRequest
	GetFiles() []*ContextualFile
	SetMessages(v []*ContextualMessage) *ContextualAnswerRequest
	GetMessages() []*ContextualMessage
	SetProjectName(v string) *ContextualAnswerRequest
	GetProjectName() *string
}

type ContextualAnswerRequest struct {
	Files []*ContextualFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// This parameter is required.
	Messages []*ContextualMessage `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ContextualAnswerRequest) String() string {
	return dara.Prettify(s)
}

func (s ContextualAnswerRequest) GoString() string {
	return s.String()
}

func (s *ContextualAnswerRequest) GetFiles() []*ContextualFile {
	return s.Files
}

func (s *ContextualAnswerRequest) GetMessages() []*ContextualMessage {
	return s.Messages
}

func (s *ContextualAnswerRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ContextualAnswerRequest) SetFiles(v []*ContextualFile) *ContextualAnswerRequest {
	s.Files = v
	return s
}

func (s *ContextualAnswerRequest) SetMessages(v []*ContextualMessage) *ContextualAnswerRequest {
	s.Messages = v
	return s
}

func (s *ContextualAnswerRequest) SetProjectName(v string) *ContextualAnswerRequest {
	s.ProjectName = &v
	return s
}

func (s *ContextualAnswerRequest) Validate() error {
	return dara.Validate(s)
}
