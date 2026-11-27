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
	// The application ID. You can create and query it in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The hotword parameter.
	//
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
	// The description. The description must be encoded in UTF-8 and be 1 to 200 characters in length.
	//
	// example:
	//
	// Fruit description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the hotword table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1qasw23ezcsrfsawq
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the hotword table. The name must be encoded in UTF-8 and be 1 to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Fruit
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Each vocabulary can contain up to 300 hotwords. Each hotword can be up to 10 Chinese characters or 5 English words in length.
	//
	// Hotwords must be encoded in UTF-8 and cannot contain punctuation marks or special characters. Otherwise, the hotwords do not take effect.
	//
	// The weight of a hotword must be an integer in the range of [-6, 5].
	//
	// Numbers in the text must be replaced with the corresponding Chinese characters based on their pronunciation. For example, \"59.9元\" must be replaced with \"五十九点九元\".
	//
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
	// The weight. The value must be an integer in the range of [-6, 5].
	//
	// - A value greater than 0 increases the probability of recognizing the word, and a value less than 0 decreases the probability.
	//
	// - A value of -6 indicates that the word should not be recognized if possible.
	//
	// - A value of 2 is a common value.
	//
	// If the effect is not obvious, you can increase the weight. However, a large weight may cause negative effects and lead to inaccurate recognition of other words.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The hotword.
	//
	// This parameter is required.
	//
	// example:
	//
	// Apple
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
