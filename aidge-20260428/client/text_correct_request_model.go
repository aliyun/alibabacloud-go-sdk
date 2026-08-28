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
	// The source language code. This parameter is required. You can set this parameter to auto for automatic language detection. 14 languages are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// de
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The text to correct. This parameter is required.
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
