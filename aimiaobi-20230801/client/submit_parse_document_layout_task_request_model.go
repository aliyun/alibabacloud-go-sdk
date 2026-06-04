// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitParseDocumentLayoutTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SubmitParseDocumentLayoutTaskRequest
	GetContent() *string
	SetWorkspaceId(v string) *SubmitParseDocumentLayoutTaskRequest
	GetWorkspaceId() *string
}

type SubmitParseDocumentLayoutTaskRequest struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitParseDocumentLayoutTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitParseDocumentLayoutTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitParseDocumentLayoutTaskRequest) GetContent() *string {
	return s.Content
}

func (s *SubmitParseDocumentLayoutTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitParseDocumentLayoutTaskRequest) SetContent(v string) *SubmitParseDocumentLayoutTaskRequest {
	s.Content = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskRequest) SetWorkspaceId(v string) *SubmitParseDocumentLayoutTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitParseDocumentLayoutTaskRequest) Validate() error {
	return dara.Validate(s)
}
