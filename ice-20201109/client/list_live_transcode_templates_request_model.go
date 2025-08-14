// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTranscodeTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListLiveTranscodeTemplatesRequest
	GetCategory() *string
	SetKeyWord(v string) *ListLiveTranscodeTemplatesRequest
	GetKeyWord() *string
	SetPageNo(v int32) *ListLiveTranscodeTemplatesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveTranscodeTemplatesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListLiveTranscodeTemplatesRequest
	GetSortBy() *string
	SetType(v string) *ListLiveTranscodeTemplatesRequest
	GetType() *string
	SetVideoCodec(v string) *ListLiveTranscodeTemplatesRequest
	GetVideoCodec() *string
}

type ListLiveTranscodeTemplatesRequest struct {
	// The category of the template. Valid values:
	//
	// 	- system
	//
	// 	- customized
	//
	// example:
	//
	// customized
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The search keyword. You can use the template ID or name as the keyword to search for templates. If you search for templates by name, fuzzy match is supported.
	//
	// example:
	//
	// my_template
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order. Valid values:
	//
	// 	- asc
	//
	// 	- desc
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the template. Valid values:
	//
	// 	- normal
	//
	// 	- narrow-band
	//
	// 	- audio-only
	//
	// 	- origin
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The video codec. Valid values:
	//
	// 	- H.264
	//
	// 	- H.265
	//
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
}

func (s ListLiveTranscodeTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTranscodeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveTranscodeTemplatesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListLiveTranscodeTemplatesRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListLiveTranscodeTemplatesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveTranscodeTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveTranscodeTemplatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveTranscodeTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListLiveTranscodeTemplatesRequest) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *ListLiveTranscodeTemplatesRequest) SetCategory(v string) *ListLiveTranscodeTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) SetKeyWord(v string) *ListLiveTranscodeTemplatesRequest {
	s.KeyWord = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) SetPageNo(v int32) *ListLiveTranscodeTemplatesRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) SetPageSize(v int32) *ListLiveTranscodeTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) SetSortBy(v string) *ListLiveTranscodeTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) SetType(v string) *ListLiveTranscodeTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) SetVideoCodec(v string) *ListLiveTranscodeTemplatesRequest {
	s.VideoCodec = &v
	return s
}

func (s *ListLiveTranscodeTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
