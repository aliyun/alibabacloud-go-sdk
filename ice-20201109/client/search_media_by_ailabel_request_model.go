// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByAILabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFilters(v string) *SearchMediaByAILabelRequest
	GetCustomFilters() *string
	SetMatchingMode(v string) *SearchMediaByAILabelRequest
	GetMatchingMode() *string
	SetMediaId(v string) *SearchMediaByAILabelRequest
	GetMediaId() *string
	SetMediaType(v string) *SearchMediaByAILabelRequest
	GetMediaType() *string
	SetMultimodalSearchType(v string) *SearchMediaByAILabelRequest
	GetMultimodalSearchType() *string
	SetNamespace(v string) *SearchMediaByAILabelRequest
	GetNamespace() *string
	SetPageNo(v int32) *SearchMediaByAILabelRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaByAILabelRequest
	GetPageSize() *int32
	SetSearchLibName(v string) *SearchMediaByAILabelRequest
	GetSearchLibName() *string
	SetSortBy(v string) *SearchMediaByAILabelRequest
	GetSortBy() *string
	SetSpecificSearch(v bool) *SearchMediaByAILabelRequest
	GetSpecificSearch() *bool
	SetText(v string) *SearchMediaByAILabelRequest
	GetText() *string
	SetUtcCreate(v string) *SearchMediaByAILabelRequest
	GetUtcCreate() *string
}

type SearchMediaByAILabelRequest struct {
	// example:
	//
	// {}
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	MatchingMode  *string `json:"MatchingMode,omitempty" xml:"MatchingMode,omitempty"`
	// The ID of the media asset. This parameter is required if you want to query media asset clips.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of the media assets. Valid values:
	//
	// 	- image
	//
	// 	- video
	//
	// 	- audio
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The type of query. Valid values:
	//
	// 	- PersonName: queries media assets based on character names.
	//
	// 	- Ocr: queries media assets based on subtitles.
	//
	// 	- AiCategory: queries media assets based on AI categories.
	//
	// 	- FullSearch (default): queries all media assets.
	//
	// example:
	//
	// Ocr
	MultimodalSearchType *string `json:"MultimodalSearchType,omitempty" xml:"MultimodalSearchType,omitempty"`
	Namespace            *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the search library.
	//
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The sorting method of the results. Valid values:
	//
	// 	- CreationTime:Desc (default): sorts results in reverse chronological order.
	//
	// 	- CreationTime:Asc: sorts results in chronological order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Specifies whether to query media asset clips. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SpecificSearch *bool `json:"SpecificSearch,omitempty" xml:"SpecificSearch,omitempty"`
	// The content that you want to query.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// {}
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
}

func (s SearchMediaByAILabelRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByAILabelRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaByAILabelRequest) GetCustomFilters() *string {
	return s.CustomFilters
}

func (s *SearchMediaByAILabelRequest) GetMatchingMode() *string {
	return s.MatchingMode
}

func (s *SearchMediaByAILabelRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaByAILabelRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaByAILabelRequest) GetMultimodalSearchType() *string {
	return s.MultimodalSearchType
}

func (s *SearchMediaByAILabelRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchMediaByAILabelRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchMediaByAILabelRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMediaByAILabelRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchMediaByAILabelRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchMediaByAILabelRequest) GetSpecificSearch() *bool {
	return s.SpecificSearch
}

func (s *SearchMediaByAILabelRequest) GetText() *string {
	return s.Text
}

func (s *SearchMediaByAILabelRequest) GetUtcCreate() *string {
	return s.UtcCreate
}

func (s *SearchMediaByAILabelRequest) SetCustomFilters(v string) *SearchMediaByAILabelRequest {
	s.CustomFilters = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetMatchingMode(v string) *SearchMediaByAILabelRequest {
	s.MatchingMode = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetMediaId(v string) *SearchMediaByAILabelRequest {
	s.MediaId = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetMediaType(v string) *SearchMediaByAILabelRequest {
	s.MediaType = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetMultimodalSearchType(v string) *SearchMediaByAILabelRequest {
	s.MultimodalSearchType = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetNamespace(v string) *SearchMediaByAILabelRequest {
	s.Namespace = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetPageNo(v int32) *SearchMediaByAILabelRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetPageSize(v int32) *SearchMediaByAILabelRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetSearchLibName(v string) *SearchMediaByAILabelRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetSortBy(v string) *SearchMediaByAILabelRequest {
	s.SortBy = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetSpecificSearch(v bool) *SearchMediaByAILabelRequest {
	s.SpecificSearch = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetText(v string) *SearchMediaByAILabelRequest {
	s.Text = &v
	return s
}

func (s *SearchMediaByAILabelRequest) SetUtcCreate(v string) *SearchMediaByAILabelRequest {
	s.UtcCreate = &v
	return s
}

func (s *SearchMediaByAILabelRequest) Validate() error {
	return dara.Validate(s)
}
