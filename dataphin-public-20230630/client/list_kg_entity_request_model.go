// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKgEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *ListKgEntityRequest
	GetEntityType() *string
	SetListQuery(v *ListKgEntityRequestListQuery) *ListKgEntityRequest
	GetListQuery() *ListKgEntityRequestListQuery
	SetOpTenantId(v int64) *ListKgEntityRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListKgEntityRequest
	GetOpUserId() *string
	SetWorkspaceId(v string) *ListKgEntityRequest
	GetWorkspaceId() *string
}

type ListKgEntityRequest struct {
	// The entity type code.
	//
	// example:
	//
	// Company
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The paged query filter conditions.
	ListQuery *ListKgEntityRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListKgEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityRequest) GoString() string {
	return s.String()
}

func (s *ListKgEntityRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListKgEntityRequest) GetListQuery() *ListKgEntityRequestListQuery {
	return s.ListQuery
}

func (s *ListKgEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListKgEntityRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListKgEntityRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListKgEntityRequest) SetEntityType(v string) *ListKgEntityRequest {
	s.EntityType = &v
	return s
}

func (s *ListKgEntityRequest) SetListQuery(v *ListKgEntityRequestListQuery) *ListKgEntityRequest {
	s.ListQuery = v
	return s
}

func (s *ListKgEntityRequest) SetOpTenantId(v int64) *ListKgEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListKgEntityRequest) SetOpUserId(v string) *ListKgEntityRequest {
	s.OpUserId = &v
	return s
}

func (s *ListKgEntityRequest) SetWorkspaceId(v string) *ListKgEntityRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListKgEntityRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKgEntityRequestListQuery struct {
	// The property filter conditions.
	FilterList []*ListKgEntityRequestListQueryFilterList `json:"FilterList,omitempty" xml:"FilterList,omitempty" type:"Repeated"`
	// The keyword for searching display properties.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKgEntityRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListKgEntityRequestListQuery) GetFilterList() []*ListKgEntityRequestListQueryFilterList {
	return s.FilterList
}

func (s *ListKgEntityRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListKgEntityRequestListQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListKgEntityRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKgEntityRequestListQuery) SetFilterList(v []*ListKgEntityRequestListQueryFilterList) *ListKgEntityRequestListQuery {
	s.FilterList = v
	return s
}

func (s *ListKgEntityRequestListQuery) SetKeyword(v string) *ListKgEntityRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListKgEntityRequestListQuery) SetPageNum(v int32) *ListKgEntityRequestListQuery {
	s.PageNum = &v
	return s
}

func (s *ListKgEntityRequestListQuery) SetPageSize(v int32) *ListKgEntityRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListKgEntityRequestListQuery) Validate() error {
	if s.FilterList != nil {
		for _, item := range s.FilterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKgEntityRequestListQueryFilterList struct {
	// The operator. Valid values:
	//
	// - eq: equal to.
	//
	// - neq: not equal to.
	//
	// - contains: contains.
	//
	// - gt: greater than.
	//
	// - gte: greater than or equal to.
	//
	// - lt: less than.
	//
	// - lte: less than or equal to.
	//
	// - like: fuzzy match.
	//
	// This parameter is required.
	//
	// example:
	//
	// eq
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// The property code.
	//
	// This parameter is required.
	//
	// example:
	//
	// company_name
	PropertyCode *string `json:"PropertyCode,omitempty" xml:"PropertyCode,omitempty"`
	// The property match value.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListKgEntityRequestListQueryFilterList) String() string {
	return dara.Prettify(s)
}

func (s ListKgEntityRequestListQueryFilterList) GoString() string {
	return s.String()
}

func (s *ListKgEntityRequestListQueryFilterList) GetOp() *string {
	return s.Op
}

func (s *ListKgEntityRequestListQueryFilterList) GetPropertyCode() *string {
	return s.PropertyCode
}

func (s *ListKgEntityRequestListQueryFilterList) GetValue() *string {
	return s.Value
}

func (s *ListKgEntityRequestListQueryFilterList) SetOp(v string) *ListKgEntityRequestListQueryFilterList {
	s.Op = &v
	return s
}

func (s *ListKgEntityRequestListQueryFilterList) SetPropertyCode(v string) *ListKgEntityRequestListQueryFilterList {
	s.PropertyCode = &v
	return s
}

func (s *ListKgEntityRequestListQueryFilterList) SetValue(v string) *ListKgEntityRequestListQueryFilterList {
	s.Value = &v
	return s
}

func (s *ListKgEntityRequestListQueryFilterList) Validate() error {
	return dara.Validate(s)
}
