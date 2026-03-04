// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranspositionResult interface {
	dara.Model
	String() string
	GoString() string
	SetTargetLanguage(v string) *TranspositionResult
	GetTargetLanguage() *string
	SetTranslatedText(v string) *TranspositionResult
	GetTranslatedText() *string
}

type TranspositionResult struct {
	// 	- The target language of the translation.
	//
	// 	- This field is only used in translation-related scenarios.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// 	- The translated text corresponding to the matched hotwords. Maximum length: 100 characters.
	//
	// 	- This field is only used in translation-related scenarios.
	TranslatedText *string `json:"TranslatedText,omitempty" xml:"TranslatedText,omitempty"`
}

func (s TranspositionResult) String() string {
	return dara.Prettify(s)
}

func (s TranspositionResult) GoString() string {
	return s.String()
}

func (s *TranspositionResult) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranspositionResult) GetTranslatedText() *string {
	return s.TranslatedText
}

func (s *TranspositionResult) SetTargetLanguage(v string) *TranspositionResult {
	s.TargetLanguage = &v
	return s
}

func (s *TranspositionResult) SetTranslatedText(v string) *TranspositionResult {
	s.TranslatedText = &v
	return s
}

func (s *TranspositionResult) Validate() error {
	return dara.Validate(s)
}
