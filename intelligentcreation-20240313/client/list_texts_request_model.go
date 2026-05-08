// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGenerationSource(v string) *ListTextsRequest
	GetGenerationSource() *string
	SetIndustry(v string) *ListTextsRequest
	GetIndustry() *string
	SetKeyword(v string) *ListTextsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListTextsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTextsRequest
	GetPageSize() *int32
	SetPublishStatus(v string) *ListTextsRequest
	GetPublishStatus() *string
	SetTextStyleType(v string) *ListTextsRequest
	GetTextStyleType() *string
	SetTextTheme(v string) *ListTextsRequest
	GetTextTheme() *string
}

type ListTextsRequest struct {
	// example:
	//
	// API
	GenerationSource *string `json:"generationSource,omitempty" xml:"generationSource,omitempty"`
	// example:
	//
	// Common
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Keyword  *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// PUBLISH
	PublishStatus *string `json:"publishStatus,omitempty" xml:"publishStatus,omitempty"`
	// example:
	//
	// WECHAT_MOMENT
	TextStyleType *string `json:"textStyleType,omitempty" xml:"textStyleType,omitempty"`
	// example:
	//
	// xxx
	TextTheme *string `json:"textTheme,omitempty" xml:"textTheme,omitempty"`
}

func (s ListTextsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTextsRequest) GoString() string {
	return s.String()
}

func (s *ListTextsRequest) GetGenerationSource() *string {
	return s.GenerationSource
}

func (s *ListTextsRequest) GetIndustry() *string {
	return s.Industry
}

func (s *ListTextsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTextsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTextsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTextsRequest) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *ListTextsRequest) GetTextStyleType() *string {
	return s.TextStyleType
}

func (s *ListTextsRequest) GetTextTheme() *string {
	return s.TextTheme
}

func (s *ListTextsRequest) SetGenerationSource(v string) *ListTextsRequest {
	s.GenerationSource = &v
	return s
}

func (s *ListTextsRequest) SetIndustry(v string) *ListTextsRequest {
	s.Industry = &v
	return s
}

func (s *ListTextsRequest) SetKeyword(v string) *ListTextsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTextsRequest) SetPageNumber(v int32) *ListTextsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTextsRequest) SetPageSize(v int32) *ListTextsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTextsRequest) SetPublishStatus(v string) *ListTextsRequest {
	s.PublishStatus = &v
	return s
}

func (s *ListTextsRequest) SetTextStyleType(v string) *ListTextsRequest {
	s.TextStyleType = &v
	return s
}

func (s *ListTextsRequest) SetTextTheme(v string) *ListTextsRequest {
	s.TextTheme = &v
	return s
}

func (s *ListTextsRequest) Validate() error {
	return dara.Validate(s)
}
