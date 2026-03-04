// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotword interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *Hotword
	GetLanguage() *string
	SetText(v string) *Hotword
	GetText() *string
	SetTranspositionResultList(v []*TranspositionResult) *Hotword
	GetTranspositionResultList() []*TranspositionResult
	SetWeight(v int32) *Hotword
	GetWeight() *int32
}

type Hotword struct {
	// The language of the hotword text. Valid values:
	//
	// 	- For structured media analysis and ASR: zh: Chinese en: English
	//
	// 	- For video translation: Supports 53 languages.
	//
	// *
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The text of the hotword.
	//
	// 	- For structured media analysis and ASR:
	//
	// 	- 	- Chinese: Up to 15 characters.
	//
	// 	- 	- English: Up to 10 words, separated by spaces.
	//
	// 	- 	- Mixed: Each letter is counted as one character (following the Chinese limit).
	//
	// 	- For video translation: Up to 100 characters.
	//
	// *
	//
	// example:
	//
	// hello
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 	- Predefined translation results for the hotword.
	//
	// 	- This field is only used in translation-related scenarios.
	TranspositionResultList []*TranspositionResult `json:"TranspositionResultList,omitempty" xml:"TranspositionResultList,omitempty" type:"Repeated"`
	// The weight of the hotword.
	//
	// 1.  Valid values: [-6,5].
	//
	// 2.  A positive value increases the likelihood of the word being recognized, while a negative value decreases the likelihood.
	//
	// 3.  A value of -6 specifies that recognition of this word should be minimized.
	//
	// 4.  Recommended value: 2.
	//
	// 5.  If the desired effect is not achieved, you can increase the weight. However, excessively high weights may negatively impact the recognition accuracy of other words.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s Hotword) String() string {
	return dara.Prettify(s)
}

func (s Hotword) GoString() string {
	return s.String()
}

func (s *Hotword) GetLanguage() *string {
	return s.Language
}

func (s *Hotword) GetText() *string {
	return s.Text
}

func (s *Hotword) GetTranspositionResultList() []*TranspositionResult {
	return s.TranspositionResultList
}

func (s *Hotword) GetWeight() *int32 {
	return s.Weight
}

func (s *Hotword) SetLanguage(v string) *Hotword {
	s.Language = &v
	return s
}

func (s *Hotword) SetText(v string) *Hotword {
	s.Text = &v
	return s
}

func (s *Hotword) SetTranspositionResultList(v []*TranspositionResult) *Hotword {
	s.TranspositionResultList = v
	return s
}

func (s *Hotword) SetWeight(v int32) *Hotword {
	s.Weight = &v
	return s
}

func (s *Hotword) Validate() error {
	if s.TranspositionResultList != nil {
		for _, item := range s.TranspositionResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
