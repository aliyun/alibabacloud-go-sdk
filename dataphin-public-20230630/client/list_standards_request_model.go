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
	SetOpUserId(v string) *ListStandardsRequest
	GetOpUserId() *string
}

type ListStandardsRequest struct {
	// The search conditions.
	//
	// This parameter is required.
	ListQuery *ListStandardsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the user who performs the operation.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
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

func (s *ListStandardsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListStandardsRequest) SetListQuery(v *ListStandardsRequestListQuery) *ListStandardsRequest {
	s.ListQuery = v
	return s
}

func (s *ListStandardsRequest) SetOpTenantId(v int64) *ListStandardsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListStandardsRequest) SetOpUserId(v string) *ListStandardsRequest {
	s.OpUserId = &v
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
	// The folder to which the standard belongs.
	//
	// example:
	//
	// /dir1/dir2
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The search keyword. Supports case-insensitive fuzzy search by standard name, English name, or code, and sorts results by relevance.
	//
	// example:
	//
	// Test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the standard sets to which the standards belong.
	StandardSetIdList []*int64 `json:"StandardSetIdList,omitempty" xml:"StandardSetIdList,omitempty" type:"Repeated"`
	// The stages of the standards. Valid values: DEV and PROD.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	StandardStage *string `json:"StandardStage,omitempty" xml:"StandardStage,omitempty"`
	// The list of standard statuses. For the DEV stage, the standard statuses are: DRAFT (Draft), UNDER_REVISION (Under Revision), UNDER_REVIEW (Under Review), REVIEW_PASSED (Approved), and IN_PUBLISH (Publishing). For the PROD stage, the standard statuses are: NOT_ACTIVATED (Not Activated), ACTIVE (Active), and EXPIRED (Expired).
	StandardStatusList []*string `json:"StandardStatusList,omitempty" xml:"StandardStatusList,omitempty" type:"Repeated"`
	// The IDs of the standard templates to which the standards belong.
	StandardTemplateIdList []*int64 `json:"StandardTemplateIdList,omitempty" xml:"StandardTemplateIdList,omitempty" type:"Repeated"`
	// The list of standard types. Valid values: Basic, and EMPTY (the standard type is empty).
	StandardTypeList []*string `json:"StandardTypeList,omitempty" xml:"StandardTypeList,omitempty" type:"Repeated"`
	// The user ID. If specified, only standards visible to this user are queried. If left empty, standards visible to the current user are queried.
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
