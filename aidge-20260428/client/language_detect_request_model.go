// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLanguageDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceText(v string) *LanguageDetectRequest
	GetSourceText() *string
}

type LanguageDetectRequest struct {
	// The source text for language detection. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试文本
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
}

func (s LanguageDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s LanguageDetectRequest) GoString() string {
	return s.String()
}

func (s *LanguageDetectRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *LanguageDetectRequest) SetSourceText(v string) *LanguageDetectRequest {
	s.SourceText = &v
	return s
}

func (s *LanguageDetectRequest) Validate() error {
	return dara.Validate(s)
}
