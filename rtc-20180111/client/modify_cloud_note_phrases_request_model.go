// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudNotePhrasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyCloudNotePhrasesRequest
	GetAppId() *string
	SetPhrase(v *ModifyCloudNotePhrasesRequestPhrase) *ModifyCloudNotePhrasesRequest
	GetPhrase() *ModifyCloudNotePhrasesRequestPhrase
}

type ModifyCloudNotePhrasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Phrase *ModifyCloudNotePhrasesRequestPhrase `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s ModifyCloudNotePhrasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudNotePhrasesRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudNotePhrasesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyCloudNotePhrasesRequest) GetPhrase() *ModifyCloudNotePhrasesRequestPhrase {
	return s.Phrase
}

func (s *ModifyCloudNotePhrasesRequest) SetAppId(v string) *ModifyCloudNotePhrasesRequest {
	s.AppId = &v
	return s
}

func (s *ModifyCloudNotePhrasesRequest) SetPhrase(v *ModifyCloudNotePhrasesRequestPhrase) *ModifyCloudNotePhrasesRequest {
	s.Phrase = v
	return s
}

func (s *ModifyCloudNotePhrasesRequest) Validate() error {
	if s.Phrase != nil {
		if err := s.Phrase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCloudNotePhrasesRequestPhrase struct {
	// example:
	//
	// 水果描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1qasw23ezcsrfsawq
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 水果
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	WordWeights []*ModifyCloudNotePhrasesRequestPhraseWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s ModifyCloudNotePhrasesRequestPhrase) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudNotePhrasesRequestPhrase) GoString() string {
	return s.String()
}

func (s *ModifyCloudNotePhrasesRequestPhrase) GetDescription() *string {
	return s.Description
}

func (s *ModifyCloudNotePhrasesRequestPhrase) GetId() *string {
	return s.Id
}

func (s *ModifyCloudNotePhrasesRequestPhrase) GetName() *string {
	return s.Name
}

func (s *ModifyCloudNotePhrasesRequestPhrase) GetWordWeights() []*ModifyCloudNotePhrasesRequestPhraseWordWeights {
	return s.WordWeights
}

func (s *ModifyCloudNotePhrasesRequestPhrase) SetDescription(v string) *ModifyCloudNotePhrasesRequestPhrase {
	s.Description = &v
	return s
}

func (s *ModifyCloudNotePhrasesRequestPhrase) SetId(v string) *ModifyCloudNotePhrasesRequestPhrase {
	s.Id = &v
	return s
}

func (s *ModifyCloudNotePhrasesRequestPhrase) SetName(v string) *ModifyCloudNotePhrasesRequestPhrase {
	s.Name = &v
	return s
}

func (s *ModifyCloudNotePhrasesRequestPhrase) SetWordWeights(v []*ModifyCloudNotePhrasesRequestPhraseWordWeights) *ModifyCloudNotePhrasesRequestPhrase {
	s.WordWeights = v
	return s
}

func (s *ModifyCloudNotePhrasesRequestPhrase) Validate() error {
	if s.WordWeights != nil {
		for _, item := range s.WordWeights {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCloudNotePhrasesRequestPhraseWordWeights struct {
	// This parameter is required.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 苹果
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ModifyCloudNotePhrasesRequestPhraseWordWeights) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudNotePhrasesRequestPhraseWordWeights) GoString() string {
	return s.String()
}

func (s *ModifyCloudNotePhrasesRequestPhraseWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyCloudNotePhrasesRequestPhraseWordWeights) GetWord() *string {
	return s.Word
}

func (s *ModifyCloudNotePhrasesRequestPhraseWordWeights) SetWeight(v int32) *ModifyCloudNotePhrasesRequestPhraseWordWeights {
	s.Weight = &v
	return s
}

func (s *ModifyCloudNotePhrasesRequestPhraseWordWeights) SetWord(v string) *ModifyCloudNotePhrasesRequestPhraseWordWeights {
	s.Word = &v
	return s
}

func (s *ModifyCloudNotePhrasesRequestPhraseWordWeights) Validate() error {
	return dara.Validate(s)
}
