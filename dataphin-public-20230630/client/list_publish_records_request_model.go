// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListPublishRecordsRequestListQuery) *ListPublishRecordsRequest
	GetListQuery() *ListPublishRecordsRequestListQuery
	SetOpTenantId(v int64) *ListPublishRecordsRequest
	GetOpTenantId() *int64
}

type ListPublishRecordsRequest struct {
	// This parameter is required.
	ListQuery *ListPublishRecordsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListPublishRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsRequest) GetListQuery() *ListPublishRecordsRequestListQuery {
	return s.ListQuery
}

func (s *ListPublishRecordsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListPublishRecordsRequest) SetListQuery(v *ListPublishRecordsRequestListQuery) *ListPublishRecordsRequest {
	s.ListQuery = v
	return s
}

func (s *ListPublishRecordsRequest) SetOpTenantId(v int64) *ListPublishRecordsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListPublishRecordsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishRecordsRequestListQuery struct {
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	SearchFilter *ListPublishRecordsRequestListQuerySearchFilter `json:"SearchFilter,omitempty" xml:"SearchFilter,omitempty" type:"Struct"`
}

func (s ListPublishRecordsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPublishRecordsRequestListQuery) GetSearchFilter() *ListPublishRecordsRequestListQuerySearchFilter {
	return s.SearchFilter
}

func (s *ListPublishRecordsRequestListQuery) SetKeyword(v string) *ListPublishRecordsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListPublishRecordsRequestListQuery) SetSearchFilter(v *ListPublishRecordsRequestListQuerySearchFilter) *ListPublishRecordsRequestListQuery {
	s.SearchFilter = v
	return s
}

func (s *ListPublishRecordsRequestListQuery) Validate() error {
	if s.SearchFilter != nil {
		if err := s.SearchFilter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPublishRecordsRequestListQuerySearchFilter struct {
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
	PublishEndTime *string `json:"PublishEndTime,omitempty" xml:"PublishEndTime,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	PublishStartTime  *string   `json:"PublishStartTime,omitempty" xml:"PublishStartTime,omitempty"`
	PublishStatusList []*int32  `json:"PublishStatusList,omitempty" xml:"PublishStatusList,omitempty" type:"Repeated"`
	SubmitterList     []*string `json:"SubmitterList,omitempty" xml:"SubmitterList,omitempty" type:"Repeated"`
}

func (s ListPublishRecordsRequestListQuerySearchFilter) String() string {
	return dara.Prettify(s)
}

func (s ListPublishRecordsRequestListQuerySearchFilter) GoString() string {
	return s.String()
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetChangeTypeList() []*int32 {
	return s.ChangeTypeList
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetPage() *int32 {
	return s.Page
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetProjectIdList() []*int64 {
	return s.ProjectIdList
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetPublishEndTime() *string {
	return s.PublishEndTime
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetPublishStartTime() *string {
	return s.PublishStartTime
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetPublishStatusList() []*int32 {
	return s.PublishStatusList
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) GetSubmitterList() []*string {
	return s.SubmitterList
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetChangeTypeList(v []*int32) *ListPublishRecordsRequestListQuerySearchFilter {
	s.ChangeTypeList = v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetPage(v int32) *ListPublishRecordsRequestListQuerySearchFilter {
	s.Page = &v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetPageSize(v int32) *ListPublishRecordsRequestListQuerySearchFilter {
	s.PageSize = &v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetProjectIdList(v []*int64) *ListPublishRecordsRequestListQuerySearchFilter {
	s.ProjectIdList = v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetPublishEndTime(v string) *ListPublishRecordsRequestListQuerySearchFilter {
	s.PublishEndTime = &v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetPublishStartTime(v string) *ListPublishRecordsRequestListQuerySearchFilter {
	s.PublishStartTime = &v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetPublishStatusList(v []*int32) *ListPublishRecordsRequestListQuerySearchFilter {
	s.PublishStatusList = v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) SetSubmitterList(v []*string) *ListPublishRecordsRequestListQuerySearchFilter {
	s.SubmitterList = v
	return s
}

func (s *ListPublishRecordsRequestListQuerySearchFilter) Validate() error {
	return dara.Validate(s)
}
