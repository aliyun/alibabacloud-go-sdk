// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtra(v string) *GetImageTranslateRequest
	GetExtra() *string
	SetSourceLanguage(v string) *GetImageTranslateRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *GetImageTranslateRequest
	GetTargetLanguage() *string
	SetUrl(v string) *GetImageTranslateRequest
	GetUrl() *string
}

type GetImageTranslateRequest struct {
	// example:
	//
	// {"have_ocr": "false", "without_text":"true", "have_psd": "false", "ignore_entity": "false"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxxxxx.oss-cn-shenzhen.aliyuncs.com/xxxxxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetImageTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetImageTranslateRequest) GetExtra() *string {
	return s.Extra
}

func (s *GetImageTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *GetImageTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *GetImageTranslateRequest) GetUrl() *string {
	return s.Url
}

func (s *GetImageTranslateRequest) SetExtra(v string) *GetImageTranslateRequest {
	s.Extra = &v
	return s
}

func (s *GetImageTranslateRequest) SetSourceLanguage(v string) *GetImageTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GetImageTranslateRequest) SetTargetLanguage(v string) *GetImageTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *GetImageTranslateRequest) SetUrl(v string) *GetImageTranslateRequest {
	s.Url = &v
	return s
}

func (s *GetImageTranslateRequest) Validate() error {
	return dara.Validate(s)
}
