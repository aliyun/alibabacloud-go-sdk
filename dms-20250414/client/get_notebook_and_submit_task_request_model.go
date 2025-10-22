// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotebookAndSubmitTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v string) *GetNotebookAndSubmitTaskRequest
	GetParams() *string
	SetPath(v string) *GetNotebookAndSubmitTaskRequest
	GetPath() *string
	SetRetry(v int64) *GetNotebookAndSubmitTaskRequest
	GetRetry() *int64
	SetSessionId(v string) *GetNotebookAndSubmitTaskRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *GetNotebookAndSubmitTaskRequest
	GetWorkspaceId() *string
}

type GetNotebookAndSubmitTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {\\"dt\\": \\"2022-10-14\\"}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /Workspace/code/default/test.ipynb
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// true
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8vkixxxxx***
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 8630242382****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetNotebookAndSubmitTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNotebookAndSubmitTaskRequest) GoString() string {
	return s.String()
}

func (s *GetNotebookAndSubmitTaskRequest) GetParams() *string {
	return s.Params
}

func (s *GetNotebookAndSubmitTaskRequest) GetPath() *string {
	return s.Path
}

func (s *GetNotebookAndSubmitTaskRequest) GetRetry() *int64 {
	return s.Retry
}

func (s *GetNotebookAndSubmitTaskRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetNotebookAndSubmitTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetNotebookAndSubmitTaskRequest) SetParams(v string) *GetNotebookAndSubmitTaskRequest {
	s.Params = &v
	return s
}

func (s *GetNotebookAndSubmitTaskRequest) SetPath(v string) *GetNotebookAndSubmitTaskRequest {
	s.Path = &v
	return s
}

func (s *GetNotebookAndSubmitTaskRequest) SetRetry(v int64) *GetNotebookAndSubmitTaskRequest {
	s.Retry = &v
	return s
}

func (s *GetNotebookAndSubmitTaskRequest) SetSessionId(v string) *GetNotebookAndSubmitTaskRequest {
	s.SessionId = &v
	return s
}

func (s *GetNotebookAndSubmitTaskRequest) SetWorkspaceId(v string) *GetNotebookAndSubmitTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetNotebookAndSubmitTaskRequest) Validate() error {
	return dara.Validate(s)
}
