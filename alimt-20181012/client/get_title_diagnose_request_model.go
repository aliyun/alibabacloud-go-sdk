// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleDiagnoseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *GetTitleDiagnoseRequest
	GetCategoryId() *string
	SetExtra(v string) *GetTitleDiagnoseRequest
	GetExtra() *string
	SetLanguage(v string) *GetTitleDiagnoseRequest
	GetLanguage() *string
	SetPlatform(v string) *GetTitleDiagnoseRequest
	GetPlatform() *string
	SetTitle(v string) *GetTitleDiagnoseRequest
	GetTitle() *string
}

type GetTitleDiagnoseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {   "product_id": "1",   "platform": "ae" }
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
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
	// Amroe Japan Comic Theme Uzumaki Naruto Namikaze Minato 3D Visual Cartoon 7 Color USB Touch
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTitleDiagnoseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTitleDiagnoseRequest) GoString() string {
	return s.String()
}

func (s *GetTitleDiagnoseRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetTitleDiagnoseRequest) GetExtra() *string {
	return s.Extra
}

func (s *GetTitleDiagnoseRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetTitleDiagnoseRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetTitleDiagnoseRequest) GetTitle() *string {
	return s.Title
}

func (s *GetTitleDiagnoseRequest) SetCategoryId(v string) *GetTitleDiagnoseRequest {
	s.CategoryId = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetExtra(v string) *GetTitleDiagnoseRequest {
	s.Extra = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetLanguage(v string) *GetTitleDiagnoseRequest {
	s.Language = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetPlatform(v string) *GetTitleDiagnoseRequest {
	s.Platform = &v
	return s
}

func (s *GetTitleDiagnoseRequest) SetTitle(v string) *GetTitleDiagnoseRequest {
	s.Title = &v
	return s
}

func (s *GetTitleDiagnoseRequest) Validate() error {
	return dara.Validate(s)
}
