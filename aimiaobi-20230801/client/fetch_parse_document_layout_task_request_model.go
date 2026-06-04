// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchParseDocumentLayoutTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *FetchParseDocumentLayoutTaskRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *FetchParseDocumentLayoutTaskRequest
	GetWorkspaceId() *string
}

type FetchParseDocumentLayoutTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 29ae0ba84c1c4cc694d0f4f1aead8005
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FetchParseDocumentLayoutTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchParseDocumentLayoutTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchParseDocumentLayoutTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FetchParseDocumentLayoutTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *FetchParseDocumentLayoutTaskRequest) SetTaskId(v string) *FetchParseDocumentLayoutTaskRequest {
	s.TaskId = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskRequest) SetWorkspaceId(v string) *FetchParseDocumentLayoutTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *FetchParseDocumentLayoutTaskRequest) Validate() error {
	return dara.Validate(s)
}
