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
	// The content of the files involved in the current Q&A. It is recommended to use the return value of the ContextualRetrieval interface as input.
	Files []*ContextualFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// Yes, the history of conversations and tool invocations. The latest message is at the end (index n-1), and the oldest message is at the beginning (index 0).
	//
	// It must be in the form of user-assistant pairs, with a total count of 2*n+1, and the length of the latest question should not exceed 1000 characters.
	//
	// The length of the historical conversation is limited to 100.
	//
	// This parameter is required.
	Messages []*ContextualMessage `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// Project name. For how to obtain it, see [Creating a Project](https://help.aliyun.com/zh/imm/getting-started/create-a-project-1?spm=a2c4g.11186623.help-menu-search-62354.d_0).
	//
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
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
