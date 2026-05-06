// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *CreateVocabularyRequest
	GetBusinessUnitId() *string
	SetDescription(v string) *CreateVocabularyRequest
	GetDescription() *string
	SetName(v string) *CreateVocabularyRequest
	GetName() *string
	SetWords(v map[string]*string) *CreateVocabularyRequest
	GetWords() map[string]*string
}

type CreateVocabularyRequest struct {
	// example:
	//
	// llm-baployoyopf22m2r
	BusinessUnitId *string            `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	Description    *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name           *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Words          map[string]*string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s CreateVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabularyRequest) GoString() string {
	return s.String()
}

func (s *CreateVocabularyRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *CreateVocabularyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVocabularyRequest) GetName() *string {
	return s.Name
}

func (s *CreateVocabularyRequest) GetWords() map[string]*string {
	return s.Words
}

func (s *CreateVocabularyRequest) SetBusinessUnitId(v string) *CreateVocabularyRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *CreateVocabularyRequest) SetDescription(v string) *CreateVocabularyRequest {
	s.Description = &v
	return s
}

func (s *CreateVocabularyRequest) SetName(v string) *CreateVocabularyRequest {
	s.Name = &v
	return s
}

func (s *CreateVocabularyRequest) SetWords(v map[string]*string) *CreateVocabularyRequest {
	s.Words = v
	return s
}

func (s *CreateVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
