// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocSmartCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *RunDocSmartCardRequest
	GetDocId() *string
	SetModelName(v string) *RunDocSmartCardRequest
	GetModelName() *string
	SetPrompt(v string) *RunDocSmartCardRequest
	GetPrompt() *string
	SetSessionId(v string) *RunDocSmartCardRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunDocSmartCardRequest
	GetWorkspaceId() *string
}

type RunDocSmartCardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 84ufBYEeLMZOjRFo84HJ7ySL3Efr55
	DocId     *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Prompt    *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunDocSmartCardRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardRequest) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunDocSmartCardRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunDocSmartCardRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunDocSmartCardRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocSmartCardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocSmartCardRequest) SetDocId(v string) *RunDocSmartCardRequest {
	s.DocId = &v
	return s
}

func (s *RunDocSmartCardRequest) SetModelName(v string) *RunDocSmartCardRequest {
	s.ModelName = &v
	return s
}

func (s *RunDocSmartCardRequest) SetPrompt(v string) *RunDocSmartCardRequest {
	s.Prompt = &v
	return s
}

func (s *RunDocSmartCardRequest) SetSessionId(v string) *RunDocSmartCardRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocSmartCardRequest) SetWorkspaceId(v string) *RunDocSmartCardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocSmartCardRequest) Validate() error {
	return dara.Validate(s)
}
