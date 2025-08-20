// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateVocabRequest
	GetDescription() *string
	SetName(v string) *UpdateVocabRequest
	GetName() *string
	SetVocabularyId(v string) *UpdateVocabRequest
	GetVocabularyId() *string
	SetWordWeightList(v []*UpdateVocabRequestWordWeightList) *UpdateVocabRequest
	GetWordWeightList() []*UpdateVocabRequestWordWeightList
	SetWorkspaceId(v string) *UpdateVocabRequest
	GetWorkspaceId() *string
}

type UpdateVocabRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dsvsv***dsvv
	VocabularyId   *string                             `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	WordWeightList []*UpdateVocabRequestWordWeightList `json:"wordWeightList,omitempty" xml:"wordWeightList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jhfr****w8v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateVocabRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabRequest) GoString() string {
	return s.String()
}

func (s *UpdateVocabRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVocabRequest) GetName() *string {
	return s.Name
}

func (s *UpdateVocabRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *UpdateVocabRequest) GetWordWeightList() []*UpdateVocabRequestWordWeightList {
	return s.WordWeightList
}

func (s *UpdateVocabRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateVocabRequest) SetDescription(v string) *UpdateVocabRequest {
	s.Description = &v
	return s
}

func (s *UpdateVocabRequest) SetName(v string) *UpdateVocabRequest {
	s.Name = &v
	return s
}

func (s *UpdateVocabRequest) SetVocabularyId(v string) *UpdateVocabRequest {
	s.VocabularyId = &v
	return s
}

func (s *UpdateVocabRequest) SetWordWeightList(v []*UpdateVocabRequestWordWeightList) *UpdateVocabRequest {
	s.WordWeightList = v
	return s
}

func (s *UpdateVocabRequest) SetWorkspaceId(v string) *UpdateVocabRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateVocabRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateVocabRequestWordWeightList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
	// This parameter is required.
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s UpdateVocabRequestWordWeightList) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabRequestWordWeightList) GoString() string {
	return s.String()
}

func (s *UpdateVocabRequestWordWeightList) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateVocabRequestWordWeightList) GetWord() *string {
	return s.Word
}

func (s *UpdateVocabRequestWordWeightList) SetWeight(v int32) *UpdateVocabRequestWordWeightList {
	s.Weight = &v
	return s
}

func (s *UpdateVocabRequestWordWeightList) SetWord(v string) *UpdateVocabRequestWordWeightList {
	s.Word = &v
	return s
}

func (s *UpdateVocabRequestWordWeightList) Validate() error {
	return dara.Validate(s)
}
