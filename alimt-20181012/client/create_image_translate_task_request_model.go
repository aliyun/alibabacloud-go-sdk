// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateImageTranslateTaskRequest
	GetClientToken() *string
	SetExtra(v string) *CreateImageTranslateTaskRequest
	GetExtra() *string
	SetSourceLanguage(v string) *CreateImageTranslateTaskRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *CreateImageTranslateTaskRequest
	GetTargetLanguage() *string
	SetUrlList(v string) *CreateImageTranslateTaskRequest
	GetUrlList() *string
}

type CreateImageTranslateTaskRequest struct {
	// example:
	//
	// 1
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {"have_ocr":"false","without_text":"false","have_psd":"true","ignore_entity":"false"}
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
	// http://xxx,http://yyy
	UrlList *string `json:"UrlList,omitempty" xml:"UrlList,omitempty"`
}

func (s CreateImageTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageTranslateTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageTranslateTaskRequest) GetExtra() *string {
	return s.Extra
}

func (s *CreateImageTranslateTaskRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *CreateImageTranslateTaskRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *CreateImageTranslateTaskRequest) GetUrlList() *string {
	return s.UrlList
}

func (s *CreateImageTranslateTaskRequest) SetClientToken(v string) *CreateImageTranslateTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetExtra(v string) *CreateImageTranslateTaskRequest {
	s.Extra = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetSourceLanguage(v string) *CreateImageTranslateTaskRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetTargetLanguage(v string) *CreateImageTranslateTaskRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) SetUrlList(v string) *CreateImageTranslateTaskRequest {
	s.UrlList = &v
	return s
}

func (s *CreateImageTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}
