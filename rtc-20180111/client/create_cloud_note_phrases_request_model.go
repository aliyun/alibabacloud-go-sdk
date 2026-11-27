// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudNotePhrasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateCloudNotePhrasesRequest
	GetAppId() *string
	SetPhrase(v *CreateCloudNotePhrasesRequestPhrase) *CreateCloudNotePhrasesRequest
	GetPhrase() *CreateCloudNotePhrasesRequestPhrase
}

type CreateCloudNotePhrasesRequest struct {
	// The application ID. You can create and query application IDs in the console. Only a single ID is supported.
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
	Phrase *CreateCloudNotePhrasesRequestPhrase `json:"Phrase,omitempty" xml:"Phrase,omitempty" type:"Struct"`
}

func (s CreateCloudNotePhrasesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudNotePhrasesRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudNotePhrasesRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateCloudNotePhrasesRequest) GetPhrase() *CreateCloudNotePhrasesRequestPhrase {
	return s.Phrase
}

func (s *CreateCloudNotePhrasesRequest) SetAppId(v string) *CreateCloudNotePhrasesRequest {
	s.AppId = &v
	return s
}

func (s *CreateCloudNotePhrasesRequest) SetPhrase(v *CreateCloudNotePhrasesRequestPhrase) *CreateCloudNotePhrasesRequest {
	s.Phrase = v
	return s
}

func (s *CreateCloudNotePhrasesRequest) Validate() error {
	if s.Phrase != nil {
		if err := s.Phrase.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCloudNotePhrasesRequestPhrase struct {
	// The description. The description must be encoded in UTF-8 and be 1 to 200 characters in length.
	//
	// example:
	//
	// Fruit description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the hotword table. The name must be encoded in UTF-8 and be 1 to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Fruits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Each hotword table can contain up to 300 hotwords, and each hotword can be up to 10 Chinese characters in length.
	//
	// Hotwords must be encoded in UTF-8 and cannot contain punctuation marks or special characters. Otherwise, the hotwords do not take effect.
	//
	// The weight of a hotword must be an integer in the range of [-6, 5].
	//
	// Numbers in the text must be replaced with the corresponding Chinese characters based on their pronunciation. For example, "59.9元" must be replaced with "五十九点九元".
	//
	// This parameter is required.
	WordWeights []*CreateCloudNotePhrasesRequestPhraseWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s CreateCloudNotePhrasesRequestPhrase) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudNotePhrasesRequestPhrase) GoString() string {
	return s.String()
}

func (s *CreateCloudNotePhrasesRequestPhrase) GetDescription() *string {
	return s.Description
}

func (s *CreateCloudNotePhrasesRequestPhrase) GetName() *string {
	return s.Name
}

func (s *CreateCloudNotePhrasesRequestPhrase) GetWordWeights() []*CreateCloudNotePhrasesRequestPhraseWordWeights {
	return s.WordWeights
}

func (s *CreateCloudNotePhrasesRequestPhrase) SetDescription(v string) *CreateCloudNotePhrasesRequestPhrase {
	s.Description = &v
	return s
}

func (s *CreateCloudNotePhrasesRequestPhrase) SetName(v string) *CreateCloudNotePhrasesRequestPhrase {
	s.Name = &v
	return s
}

func (s *CreateCloudNotePhrasesRequestPhrase) SetWordWeights(v []*CreateCloudNotePhrasesRequestPhraseWordWeights) *CreateCloudNotePhrasesRequestPhrase {
	s.WordWeights = v
	return s
}

func (s *CreateCloudNotePhrasesRequestPhrase) Validate() error {
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

type CreateCloudNotePhrasesRequestPhraseWordWeights struct {
	// The weight must be an integer in the range of [-6, 5].
	//
	// - A value greater than 0 increases the probability of the word being recognized, and a value less than 0 decreases the probability.
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

func (s CreateCloudNotePhrasesRequestPhraseWordWeights) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudNotePhrasesRequestPhraseWordWeights) GoString() string {
	return s.String()
}

func (s *CreateCloudNotePhrasesRequestPhraseWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateCloudNotePhrasesRequestPhraseWordWeights) GetWord() *string {
	return s.Word
}

func (s *CreateCloudNotePhrasesRequestPhraseWordWeights) SetWeight(v int32) *CreateCloudNotePhrasesRequestPhraseWordWeights {
	s.Weight = &v
	return s
}

func (s *CreateCloudNotePhrasesRequestPhraseWordWeights) SetWord(v string) *CreateCloudNotePhrasesRequestPhraseWordWeights {
	s.Word = &v
	return s
}

func (s *CreateCloudNotePhrasesRequestPhraseWordWeights) Validate() error {
	return dara.Validate(s)
}
