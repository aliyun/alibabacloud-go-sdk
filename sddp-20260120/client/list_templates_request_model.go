// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTemplatesRequest
	GetCurrentPage() *int32
	SetFeatureType(v int32) *ListTemplatesRequest
	GetFeatureType() *int32
	SetLang(v string) *ListTemplatesRequest
	GetLang() *string
	SetPageSize(v int32) *ListTemplatesRequest
	GetPageSize() *int32
	SetUsageScenario(v int32) *ListTemplatesRequest
	GetUsageScenario() *int32
}

type ListTemplatesRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	FeatureType   *int32  `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UsageScenario *int32  `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTemplatesRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *ListTemplatesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetUsageScenario() *int32 {
	return s.UsageScenario
}

func (s *ListTemplatesRequest) SetCurrentPage(v int32) *ListTemplatesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTemplatesRequest) SetFeatureType(v int32) *ListTemplatesRequest {
	s.FeatureType = &v
	return s
}

func (s *ListTemplatesRequest) SetLang(v string) *ListTemplatesRequest {
	s.Lang = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetUsageScenario(v int32) *ListTemplatesRequest {
	s.UsageScenario = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
