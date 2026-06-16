// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextCorrectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceLanguage(v string) *TextCorrectRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *TextCorrectRequest
	GetSourceText() *string
}

type TextCorrectRequest struct {
	// Source language code. Required. You can pass "auto" for automatic language detection. Supports 14 languages.
	//
	// This parameter is required.
	//
	// example:
	//
	// de
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// Text to be corrected. Required.
	//
	// This parameter is required.
	//
	// example:
	//
	// Empfelung
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s TextCorrectRequest) String() string {
	return dara.Prettify(s)
}

func (s TextCorrectRequest) GoString() string {
	return s.String()
}

func (s *TextCorrectRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextCorrectRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *TextCorrectRequest) SetSourceLanguage(v string) *TextCorrectRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextCorrectRequest) SetSourceText(v string) *TextCorrectRequest {
	s.SourceText = &v
	return s
}

func (s *TextCorrectRequest) Validate() error {
	return dara.Validate(s)
}
