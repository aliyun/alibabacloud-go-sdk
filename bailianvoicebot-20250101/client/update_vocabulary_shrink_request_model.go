// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabularyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *UpdateVocabularyShrinkRequest
	GetBusinessUnitId() *string
	SetDescription(v string) *UpdateVocabularyShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdateVocabularyShrinkRequest
	GetName() *string
	SetVocabularyId(v string) *UpdateVocabularyShrinkRequest
	GetVocabularyId() *string
	SetWordsShrink(v string) *UpdateVocabularyShrinkRequest
	GetWordsShrink() *string
}

type UpdateVocabularyShrinkRequest struct {
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4061
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	WordsShrink  *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s UpdateVocabularyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabularyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateVocabularyShrinkRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateVocabularyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVocabularyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateVocabularyShrinkRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *UpdateVocabularyShrinkRequest) GetWordsShrink() *string {
	return s.WordsShrink
}

func (s *UpdateVocabularyShrinkRequest) SetBusinessUnitId(v string) *UpdateVocabularyShrinkRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateVocabularyShrinkRequest) SetDescription(v string) *UpdateVocabularyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateVocabularyShrinkRequest) SetName(v string) *UpdateVocabularyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateVocabularyShrinkRequest) SetVocabularyId(v string) *UpdateVocabularyShrinkRequest {
	s.VocabularyId = &v
	return s
}

func (s *UpdateVocabularyShrinkRequest) SetWordsShrink(v string) *UpdateVocabularyShrinkRequest {
	s.WordsShrink = &v
	return s
}

func (s *UpdateVocabularyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
