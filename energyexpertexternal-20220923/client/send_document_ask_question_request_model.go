// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDocumentAskQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderId(v string) *SendDocumentAskQuestionRequest
	GetFolderId() *string
	SetPrompt(v string) *SendDocumentAskQuestionRequest
	GetPrompt() *string
	SetSessionId(v string) *SendDocumentAskQuestionRequest
	GetSessionId() *string
}

type SendDocumentAskQuestionRequest struct {
	// Folder ID, used to specify the range of documents for the query. If it is empty, it indicates that all documents under the default folder will be queried.
	//
	// example:
	//
	// 1a851c4a-1d65-11ef-99a7-ssfsfdd
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// The question queried by the user
	//
	// This parameter is required.
	//
	// example:
	//
	// Total carbon emissions in 2023
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// Q&A session ID, used to record multiple Q&A interactions of the same user. If it is empty, it indicates that sessions are not distinguished.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s SendDocumentAskQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s SendDocumentAskQuestionRequest) GoString() string {
	return s.String()
}

func (s *SendDocumentAskQuestionRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SendDocumentAskQuestionRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *SendDocumentAskQuestionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendDocumentAskQuestionRequest) SetFolderId(v string) *SendDocumentAskQuestionRequest {
	s.FolderId = &v
	return s
}

func (s *SendDocumentAskQuestionRequest) SetPrompt(v string) *SendDocumentAskQuestionRequest {
	s.Prompt = &v
	return s
}

func (s *SendDocumentAskQuestionRequest) SetSessionId(v string) *SendDocumentAskQuestionRequest {
	s.SessionId = &v
	return s
}

func (s *SendDocumentAskQuestionRequest) Validate() error {
	return dara.Validate(s)
}
