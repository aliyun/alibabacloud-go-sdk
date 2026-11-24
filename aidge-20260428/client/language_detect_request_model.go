// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLanguageDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScene(v string) *LanguageDetectRequest
	GetScene() *string
	SetSourceText(v string) *LanguageDetectRequest
	GetSourceText() *string
}

type LanguageDetectRequest struct {
	// The detection scenario. Default value: common. For search phrase scenarios, set this parameter to query. If an invalid value is specified or this parameter is left empty, the common general detection scenario is used. Note: The value query must be in lowercase.
	//
	// example:
	//
	// query
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The source text to be detected. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test text
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s LanguageDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s LanguageDetectRequest) GoString() string {
	return s.String()
}

func (s *LanguageDetectRequest) GetScene() *string {
	return s.Scene
}

func (s *LanguageDetectRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *LanguageDetectRequest) SetScene(v string) *LanguageDetectRequest {
	s.Scene = &v
	return s
}

func (s *LanguageDetectRequest) SetSourceText(v string) *LanguageDetectRequest {
	s.SourceText = &v
	return s
}

func (s *LanguageDetectRequest) Validate() error {
	return dara.Validate(s)
}
