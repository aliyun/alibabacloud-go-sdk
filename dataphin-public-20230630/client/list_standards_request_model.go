// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListStandardsRequestListQuery) *ListStandardsRequest
	GetListQuery() *ListStandardsRequestListQuery
	SetOpTenantId(v int64) *ListStandardsRequest
	GetOpTenantId() *int64
}

type ListStandardsRequest struct {
	// Search conditions.
	//
	// This parameter is required.
	ListQuery *ListStandardsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListStandardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsRequest) GoString() string {
	return s.String()
}

func (s *ListStandardsRequest) GetListQuery() *ListStandardsRequestListQuery {
	return s.ListQuery
}

func (s *ListStandardsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListStandardsRequest) SetListQuery(v *ListStandardsRequestListQuery) *ListStandardsRequest {
	s.ListQuery = v
	return s
}

func (s *ListStandardsRequest) SetOpTenantId(v int64) *ListStandardsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStandardsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStandardsRequestListQuery struct {
	// Directory of the standard.
	//
	// example:
	//
	// /dir1/dir2
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// Search keyword: fuzzy search by standard name, English name, or code. Case-insensitive, sorted by relevance.
	//
	// example:
	//
	// Test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Standard set ID list.
	StandardSetIdList []*int64 `json:"StandardSetIdList,omitempty" xml:"StandardSetIdList,omitempty" type:"Repeated"`
	// Stage of the standard: DEV or PROD.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	StandardStage *string `json:"StandardStage,omitempty" xml:"StandardStage,omitempty"`
	// Standard status list. Standard statuses under DEV stage: DRAFT, UNDER_REVISION, UNDER_REVIEW, REVIEW_PASSED, IN_PUBLISH. Standard statuses under PROD stage: NOT_ACTIVATED, ACTIVE, EXPIRED.
	StandardStatusList []*string `json:"StandardStatusList,omitempty" xml:"StandardStatusList,omitempty" type:"Repeated"`
	// Standard template ID list.
	StandardTemplateIdList []*int64 `json:"StandardTemplateIdList,omitempty" xml:"StandardTemplateIdList,omitempty" type:"Repeated"`
	// Standard type: Basic, EMPTY indicates the standard type is empty.
	StandardTypeList []*string `json:"StandardTypeList,omitempty" xml:"StandardTypeList,omitempty" type:"Repeated"`
	// User ID: only queries standards visible to this user ID. If empty, queries standards visible to the current user.
	//
	// example:
	//
	// 30012011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListStandardsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListStandardsRequestListQuery) GetDirectory() *string {
	return s.Directory
}

func (s *ListStandardsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListStandardsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListStandardsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStandardsRequestListQuery) GetStandardSetIdList() []*int64 {
	return s.StandardSetIdList
}

func (s *ListStandardsRequestListQuery) GetStandardStage() *string {
	return s.StandardStage
}

func (s *ListStandardsRequestListQuery) GetStandardStatusList() []*string {
	return s.StandardStatusList
}

func (s *ListStandardsRequestListQuery) GetStandardTemplateIdList() []*int64 {
	return s.StandardTemplateIdList
}

func (s *ListStandardsRequestListQuery) GetStandardTypeList() []*string {
	return s.StandardTypeList
}

func (s *ListStandardsRequestListQuery) GetUserId() *string {
	return s.UserId
}

func (s *ListStandardsRequestListQuery) SetDirectory(v string) *ListStandardsRequestListQuery {
	s.Directory = &v
	return s
}

func (s *ListStandardsRequestListQuery) SetKeyword(v string) *ListStandardsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListStandardsRequestListQuery) SetPageNo(v int32) *ListStandardsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListStandardsRequestListQuery) SetPageSize(v int32) *ListStandardsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListStandardsRequestListQuery) SetStandardSetIdList(v []*int64) *ListStandardsRequestListQuery {
	s.StandardSetIdList = v
	return s
}

func (s *ListStandardsRequestListQuery) SetStandardStage(v string) *ListStandardsRequestListQuery {
	s.StandardStage = &v
	return s
}

func (s *ListStandardsRequestListQuery) SetStandardStatusList(v []*string) *ListStandardsRequestListQuery {
	s.StandardStatusList = v
	return s
}

func (s *ListStandardsRequestListQuery) SetStandardTemplateIdList(v []*int64) *ListStandardsRequestListQuery {
	s.StandardTemplateIdList = v
	return s
}

func (s *ListStandardsRequestListQuery) SetStandardTypeList(v []*string) *ListStandardsRequestListQuery {
	s.StandardTypeList = v
	return s
}

func (s *ListStandardsRequestListQuery) SetUserId(v string) *ListStandardsRequestListQuery {
	s.UserId = &v
	return s
}

func (s *ListStandardsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
