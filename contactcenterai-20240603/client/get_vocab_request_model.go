// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVocabularyId(v string) *GetVocabRequest
	GetVocabularyId() *string
	SetWorkspaceId(v string) *GetVocabRequest
	GetWorkspaceId() *string
}

type GetVocabRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dhbf***rbrdb
	VocabularyId *string `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-9864***1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVocabRequest) GoString() string {
	return s.String()
}

func (s *GetVocabRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetVocabRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetVocabRequest) SetVocabularyId(v string) *GetVocabRequest {
	s.VocabularyId = &v
	return s
}

func (s *GetVocabRequest) SetWorkspaceId(v string) *GetVocabRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetVocabRequest) Validate() error {
	return dara.Validate(s)
}
