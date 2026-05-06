// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *UpdateVocabularyRequest
	GetBusinessUnitId() *string
	SetDescription(v string) *UpdateVocabularyRequest
	GetDescription() *string
	SetName(v string) *UpdateVocabularyRequest
	GetName() *string
	SetVocabularyId(v string) *UpdateVocabularyRequest
	GetVocabularyId() *string
	SetWords(v map[string]*string) *UpdateVocabularyRequest
	GetWords() map[string]*string
}

type UpdateVocabularyRequest struct {
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// d74d6290-7cbe-4436-b5d7-014ebb0f4061
	VocabularyId *string            `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	Words        map[string]*string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s UpdateVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabularyRequest) GoString() string {
	return s.String()
}

func (s *UpdateVocabularyRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *UpdateVocabularyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVocabularyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateVocabularyRequest) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *UpdateVocabularyRequest) GetWords() map[string]*string {
	return s.Words
}

func (s *UpdateVocabularyRequest) SetBusinessUnitId(v string) *UpdateVocabularyRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *UpdateVocabularyRequest) SetDescription(v string) *UpdateVocabularyRequest {
	s.Description = &v
	return s
}

func (s *UpdateVocabularyRequest) SetName(v string) *UpdateVocabularyRequest {
	s.Name = &v
	return s
}

func (s *UpdateVocabularyRequest) SetVocabularyId(v string) *UpdateVocabularyRequest {
	s.VocabularyId = &v
	return s
}

func (s *UpdateVocabularyRequest) SetWords(v map[string]*string) *UpdateVocabularyRequest {
	s.Words = v
	return s
}

func (s *UpdateVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
