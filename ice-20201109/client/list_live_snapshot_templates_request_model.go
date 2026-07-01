// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListLiveSnapshotTemplatesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListLiveSnapshotTemplatesRequest
	GetPageSize() *int32
	SetSearchKeyWord(v string) *ListLiveSnapshotTemplatesRequest
	GetSearchKeyWord() *string
	SetSortBy(v string) *ListLiveSnapshotTemplatesRequest
	GetSortBy() *string
	SetTemplateIds(v []*string) *ListLiveSnapshotTemplatesRequest
	GetTemplateIds() []*string
	SetType(v string) *ListLiveSnapshotTemplatesRequest
	GetType() *string
}

type ListLiveSnapshotTemplatesRequest struct {
	// The page number. The value must be greater than or equal to 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword. You can perform a fuzzy search by template ID or template name.
	//
	// - Maximum length: 128 characters.
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	// The sorting method. By default, results are sorted by creation time in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The template IDs.
	//
	// - This parameter does not take effect if `SearchKeyWord` is specified.
	//
	// - You can specify a maximum of 200 template IDs.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
	// The type of the template. By default, templates of all types are queried.
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveSnapshotTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotTemplatesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListLiveSnapshotTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveSnapshotTemplatesRequest) GetSearchKeyWord() *string {
	return s.SearchKeyWord
}

func (s *ListLiveSnapshotTemplatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveSnapshotTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *ListLiveSnapshotTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListLiveSnapshotTemplatesRequest) SetPageNo(v int32) *ListLiveSnapshotTemplatesRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveSnapshotTemplatesRequest) SetPageSize(v int32) *ListLiveSnapshotTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveSnapshotTemplatesRequest) SetSearchKeyWord(v string) *ListLiveSnapshotTemplatesRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *ListLiveSnapshotTemplatesRequest) SetSortBy(v string) *ListLiveSnapshotTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveSnapshotTemplatesRequest) SetTemplateIds(v []*string) *ListLiveSnapshotTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *ListLiveSnapshotTemplatesRequest) SetType(v string) *ListLiveSnapshotTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListLiveSnapshotTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
