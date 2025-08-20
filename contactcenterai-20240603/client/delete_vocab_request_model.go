// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVocabularyId(v string) *DeleteVocabRequest
	GetVocabularyId() *string
	SetWorkspaceId(v string) *DeleteVocabRequest
	GetWorkspaceId() *string
}

type DeleteVocabRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ern*******rve
	VocabularyId *string `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-0*****jlg8s
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVocabRequest) GoString() string {
	return s.String()
}

func (s *DeleteVocabRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *DeleteVocabRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteVocabRequest) SetVocabularyId(v string) *DeleteVocabRequest {
	s.VocabularyId = &v
	return s
}

func (s *DeleteVocabRequest) SetWorkspaceId(v string) *DeleteVocabRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteVocabRequest) Validate() error {
	return dara.Validate(s)
}
