// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateVocabularyRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateVocabularyRequest
	GetInstanceId() *string
	SetName(v string) *UpdateVocabularyRequest
	GetName() *string
	SetVocabularyId(v string) *UpdateVocabularyRequest
	GetVocabularyId() *string
	SetWords(v map[string]*string) *UpdateVocabularyRequest
	GetWords() map[string]*string
}

type UpdateVocabularyRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *UpdateVocabularyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateVocabularyRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *UpdateVocabularyRequest) SetDescription(v string) *UpdateVocabularyRequest {
	s.Description = &v
	return s
}

func (s *UpdateVocabularyRequest) SetInstanceId(v string) *UpdateVocabularyRequest {
	s.InstanceId = &v
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
