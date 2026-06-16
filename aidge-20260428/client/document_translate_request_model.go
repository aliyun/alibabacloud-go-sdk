// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *DocumentTranslateRequest
	GetFileType() *string
	SetGlossary(v string) *DocumentTranslateRequest
	GetGlossary() *string
	SetTargetLanguage(v string) *DocumentTranslateRequest
	GetTargetLanguage() *string
	SetUrl(v string) *DocumentTranslateRequest
	GetUrl() *string
}

type DocumentTranslateRequest struct {
	// The type of the document. Valid values: PDF and Word. Size limits: Word 200 MB/300 pages, PDF 200 MB/300 pages. The maximum size of a single file is 200 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// PDF
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The glossary ID to use when the glossary feature is required. Supports custom translation results, including do-not-translate (ABC-ABC), specified translation (ABC-DEF), and skip translation (ABC-empty value). This is commonly used for brand name protection.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The target language. The language code uses the two-letter ISO 639-1 standard.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The OSS URL of the document to be translated.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aib-innovation-oss.oss-accelerate.aliyuncs.com/AI_Business/38dao/testdemo.pdf?Expires=3356578313&OSSAccessKeyId=LTAI5tE8X3gEy66SRU1V8dig&Signature=8niQY2HtMQY7h05zmSUdyORML9E%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DocumentTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s DocumentTranslateRequest) GoString() string {
	return s.String()
}

func (s *DocumentTranslateRequest) GetFileType() *string {
	return s.FileType
}

func (s *DocumentTranslateRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *DocumentTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *DocumentTranslateRequest) GetUrl() *string {
	return s.Url
}

func (s *DocumentTranslateRequest) SetFileType(v string) *DocumentTranslateRequest {
	s.FileType = &v
	return s
}

func (s *DocumentTranslateRequest) SetGlossary(v string) *DocumentTranslateRequest {
	s.Glossary = &v
	return s
}

func (s *DocumentTranslateRequest) SetTargetLanguage(v string) *DocumentTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *DocumentTranslateRequest) SetUrl(v string) *DocumentTranslateRequest {
	s.Url = &v
	return s
}

func (s *DocumentTranslateRequest) Validate() error {
	return dara.Validate(s)
}
