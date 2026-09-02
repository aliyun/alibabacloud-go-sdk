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
	// The detection scenario. Default value: common. If you are using a search phrase scenario, set this parameter to query. If an incorrect value is passed or the parameter is not specified, the common general identification is used. Note: pass query in lowercase.
	//
	// example:
	//
	// query
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The source text to be identified. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample text
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
