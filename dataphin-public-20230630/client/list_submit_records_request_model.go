// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubmitRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListSubmitRecordsRequestListQuery) *ListSubmitRecordsRequest
	GetListQuery() *ListSubmitRecordsRequestListQuery
	SetOpTenantId(v int64) *ListSubmitRecordsRequest
	GetOpTenantId() *int64
}

type ListSubmitRecordsRequest struct {
	// This parameter is required.
	ListQuery *ListSubmitRecordsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSubmitRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsRequest) GetListQuery() *ListSubmitRecordsRequestListQuery {
	return s.ListQuery
}

func (s *ListSubmitRecordsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSubmitRecordsRequest) SetListQuery(v *ListSubmitRecordsRequestListQuery) *ListSubmitRecordsRequest {
	s.ListQuery = v
	return s
}

func (s *ListSubmitRecordsRequest) SetOpTenantId(v int64) *ListSubmitRecordsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSubmitRecordsRequest) Validate() error {
	return dara.Validate(s)
}

type ListSubmitRecordsRequestListQuery struct {
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	SearchFilter *ListSubmitRecordsRequestListQuerySearchFilter `json:"SearchFilter,omitempty" xml:"SearchFilter,omitempty" type:"Struct"`
}

func (s ListSubmitRecordsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSubmitRecordsRequestListQuery) GetSearchFilter() *ListSubmitRecordsRequestListQuerySearchFilter {
	return s.SearchFilter
}

func (s *ListSubmitRecordsRequestListQuery) SetKeyword(v string) *ListSubmitRecordsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListSubmitRecordsRequestListQuery) SetSearchFilter(v *ListSubmitRecordsRequestListQuerySearchFilter) *ListSubmitRecordsRequestListQuery {
	s.SearchFilter = v
	return s
}

func (s *ListSubmitRecordsRequestListQuery) Validate() error {
	return dara.Validate(s)
}

type ListSubmitRecordsRequestListQuerySearchFilter struct {
	ChangeTypeList []*int32 `json:"ChangeTypeList,omitempty" xml:"ChangeTypeList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	ProjectIdList []*int64 `json:"ProjectIdList,omitempty" xml:"ProjectIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-10-10 10:00:00
	SubmitEndTime *string `json:"SubmitEndTime,omitempty" xml:"SubmitEndTime,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	SubmitStartTime *string   `json:"SubmitStartTime,omitempty" xml:"SubmitStartTime,omitempty"`
	SubmitterList   []*string `json:"SubmitterList,omitempty" xml:"SubmitterList,omitempty" type:"Repeated"`
}

func (s ListSubmitRecordsRequestListQuerySearchFilter) String() string {
	return dara.Prettify(s)
}

func (s ListSubmitRecordsRequestListQuerySearchFilter) GoString() string {
	return s.String()
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetChangeTypeList() []*int32 {
	return s.ChangeTypeList
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetPage() *int32 {
	return s.Page
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetProjectIdList() []*int64 {
	return s.ProjectIdList
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetSubmitEndTime() *string {
	return s.SubmitEndTime
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetSubmitStartTime() *string {
	return s.SubmitStartTime
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) GetSubmitterList() []*string {
	return s.SubmitterList
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetChangeTypeList(v []*int32) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.ChangeTypeList = v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetPage(v int32) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.Page = &v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetPageSize(v int32) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.PageSize = &v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetProjectIdList(v []*int64) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.ProjectIdList = v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetSubmitEndTime(v string) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.SubmitEndTime = &v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetSubmitStartTime(v string) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.SubmitStartTime = &v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) SetSubmitterList(v []*string) *ListSubmitRecordsRequestListQuerySearchFilter {
	s.SubmitterList = v
	return s
}

func (s *ListSubmitRecordsRequestListQuerySearchFilter) Validate() error {
	return dara.Validate(s)
}
