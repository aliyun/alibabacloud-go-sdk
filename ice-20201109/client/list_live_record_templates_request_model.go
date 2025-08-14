// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRecordTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListLiveRecordTemplatesRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListLiveRecordTemplatesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLiveRecordTemplatesRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListLiveRecordTemplatesRequest
	GetSortBy() *string
	SetTemplateIds(v []*string) *ListLiveRecordTemplatesRequest
	GetTemplateIds() []*string
	SetType(v string) *ListLiveRecordTemplatesRequest
	GetType() *string
}

type ListLiveRecordTemplatesRequest struct {
	// The search keyword. You can use the template ID or name as the keyword to search for templates. If you search for templates by name, fuzzy match is supported.
	//
	// example:
	//
	// test template
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- asc: sorts the query results in ascending order.
	//
	// 	- desc: sorts the query results in descending order.
	//
	// example:
	//
	// desc
	SortBy      *string   `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
	// The type of the template.
	//
	// Valid values:
	//
	// 	- system
	//
	// 	- custom
	//
	// example:
	//
	// custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLiveRecordTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRecordTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRecordTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLiveRecordTemplatesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLiveRecordTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLiveRecordTemplatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLiveRecordTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *ListLiveRecordTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListLiveRecordTemplatesRequest) SetKeyword(v string) *ListLiveRecordTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListLiveRecordTemplatesRequest) SetPageNo(v int64) *ListLiveRecordTemplatesRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordTemplatesRequest) SetPageSize(v int64) *ListLiveRecordTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordTemplatesRequest) SetSortBy(v string) *ListLiveRecordTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordTemplatesRequest) SetTemplateIds(v []*string) *ListLiveRecordTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *ListLiveRecordTemplatesRequest) SetType(v string) *ListLiveRecordTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListLiveRecordTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
