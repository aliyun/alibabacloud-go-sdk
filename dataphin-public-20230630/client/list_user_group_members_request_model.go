// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserGroupMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListUserGroupMembersRequestListQuery) *ListUserGroupMembersRequest
	GetListQuery() *ListUserGroupMembersRequestListQuery
	SetOpTenantId(v int64) *ListUserGroupMembersRequest
	GetOpTenantId() *int64
}

type ListUserGroupMembersRequest struct {
	// This parameter is required.
	ListQuery *ListUserGroupMembersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListUserGroupMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersRequest) GetListQuery() *ListUserGroupMembersRequestListQuery {
	return s.ListQuery
}

func (s *ListUserGroupMembersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUserGroupMembersRequest) SetListQuery(v *ListUserGroupMembersRequestListQuery) *ListUserGroupMembersRequest {
	s.ListQuery = v
	return s
}

func (s *ListUserGroupMembersRequest) SetOpTenantId(v int64) *ListUserGroupMembersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserGroupMembersRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserGroupMembersRequestListQuery struct {
	// example:
	//
	// a
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 232231
	UserGroupId *string   `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIdList  []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s ListUserGroupMembersRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListUserGroupMembersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListUserGroupMembersRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserGroupMembersRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListUserGroupMembersRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserGroupMembersRequestListQuery) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListUserGroupMembersRequestListQuery) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *ListUserGroupMembersRequestListQuery) SetKeyword(v string) *ListUserGroupMembersRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetPageNo(v int32) *ListUserGroupMembersRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetPageSize(v int32) *ListUserGroupMembersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetUserGroupId(v string) *ListUserGroupMembersRequestListQuery {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) SetUserIdList(v []*string) *ListUserGroupMembersRequestListQuery {
	s.UserIdList = v
	return s
}

func (s *ListUserGroupMembersRequestListQuery) Validate() error {
	return dara.Validate(s)
}
