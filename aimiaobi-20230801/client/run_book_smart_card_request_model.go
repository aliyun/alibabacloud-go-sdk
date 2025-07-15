// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookSmartCardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *RunBookSmartCardRequest
	GetDocId() *string
	SetSessionId(v string) *RunBookSmartCardRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunBookSmartCardRequest
	GetWorkspaceId() *string
}

type RunBookSmartCardRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 84ufBYEeLMZOjRFo84HJ7ySL3Efr55
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
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
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunBookSmartCardRequest) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardRequest) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunBookSmartCardRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunBookSmartCardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunBookSmartCardRequest) SetDocId(v string) *RunBookSmartCardRequest {
	s.DocId = &v
	return s
}

func (s *RunBookSmartCardRequest) SetSessionId(v string) *RunBookSmartCardRequest {
	s.SessionId = &v
	return s
}

func (s *RunBookSmartCardRequest) SetWorkspaceId(v string) *RunBookSmartCardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunBookSmartCardRequest) Validate() error {
	return dara.Validate(s)
}
