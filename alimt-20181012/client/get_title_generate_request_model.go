// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleGenerateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v string) *GetTitleGenerateRequest
	GetAttributes() *string
	SetCategoryId(v string) *GetTitleGenerateRequest
	GetCategoryId() *string
	SetExtra(v string) *GetTitleGenerateRequest
	GetExtra() *string
	SetHotWords(v string) *GetTitleGenerateRequest
	GetHotWords() *string
	SetLanguage(v string) *GetTitleGenerateRequest
	GetLanguage() *string
	SetPlatform(v string) *GetTitleGenerateRequest
	GetPlatform() *string
	SetTitle(v string) *GetTitleGenerateRequest
	GetTitle() *string
}

type GetTitleGenerateRequest struct {
	// example:
	//
	// {2:"None",10:"Plastic"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127896018
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {   "product_id": "1",   "platform": "ae" }
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// None,Plastic
	HotWords *string `json:"HotWords,omitempty" xml:"HotWords,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ae
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10pcs 80ml Kitchen Disposable
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTitleGenerateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTitleGenerateRequest) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateRequest) GetAttributes() *string {
	return s.Attributes
}

func (s *GetTitleGenerateRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetTitleGenerateRequest) GetExtra() *string {
	return s.Extra
}

func (s *GetTitleGenerateRequest) GetHotWords() *string {
	return s.HotWords
}

func (s *GetTitleGenerateRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetTitleGenerateRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetTitleGenerateRequest) GetTitle() *string {
	return s.Title
}

func (s *GetTitleGenerateRequest) SetAttributes(v string) *GetTitleGenerateRequest {
	s.Attributes = &v
	return s
}

func (s *GetTitleGenerateRequest) SetCategoryId(v string) *GetTitleGenerateRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTitleGenerateRequest) SetExtra(v string) *GetTitleGenerateRequest {
	s.Extra = &v
	return s
}

func (s *GetTitleGenerateRequest) SetHotWords(v string) *GetTitleGenerateRequest {
	s.HotWords = &v
	return s
}

func (s *GetTitleGenerateRequest) SetLanguage(v string) *GetTitleGenerateRequest {
	s.Language = &v
	return s
}

func (s *GetTitleGenerateRequest) SetPlatform(v string) *GetTitleGenerateRequest {
	s.Platform = &v
	return s
}

func (s *GetTitleGenerateRequest) SetTitle(v string) *GetTitleGenerateRequest {
	s.Title = &v
	return s
}

func (s *GetTitleGenerateRequest) Validate() error {
	return dara.Validate(s)
}
