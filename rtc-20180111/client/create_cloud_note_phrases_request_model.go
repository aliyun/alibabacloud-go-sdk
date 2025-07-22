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
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	return dara.Validate(s)
}

type CreateCloudNotePhrasesRequestPhrase struct {
	// example:
	//
	// 水果描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 水果
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	return dara.Validate(s)
}

type CreateCloudNotePhrasesRequestPhraseWordWeights struct {
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
