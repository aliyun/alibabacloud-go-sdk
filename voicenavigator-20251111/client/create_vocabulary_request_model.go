// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVocabularyRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateVocabularyRequest
	GetInstanceId() *string
	SetName(v string) *CreateVocabularyRequest
	GetName() *string
	SetWords(v map[string]*string) *CreateVocabularyRequest
	GetWords() map[string]*string
}

type CreateVocabularyRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// e48e45dd-e47a-4744-a063-f08cbebb1c5a
	InstanceId *string            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Words      map[string]*string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s CreateVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabularyRequest) GoString() string {
	return s.String()
}

func (s *CreateVocabularyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVocabularyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVocabularyRequest) GetName() *string {
	return s.Name
}

func (s *CreateVocabularyRequest) GetWords() map[string]*string {
	return s.Words
}

func (s *CreateVocabularyRequest) SetDescription(v string) *CreateVocabularyRequest {
	s.Description = &v
	return s
}

func (s *CreateVocabularyRequest) SetInstanceId(v string) *CreateVocabularyRequest {
	s.InstanceId = &v
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
