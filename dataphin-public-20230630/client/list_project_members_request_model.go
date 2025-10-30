// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListProjectMembersRequest
	GetId() *int64
	SetListQuery(v *ListProjectMembersRequestListQuery) *ListProjectMembersRequest
	GetListQuery() *ListProjectMembersRequestListQuery
	SetOpTenantId(v int64) *ListProjectMembersRequest
	GetOpTenantId() *int64
}

type ListProjectMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	ListQuery *ListProjectMembersRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListProjectMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequest) GetId() *int64 {
	return s.Id
}

func (s *ListProjectMembersRequest) GetListQuery() *ListProjectMembersRequestListQuery {
	return s.ListQuery
}

func (s *ListProjectMembersRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListProjectMembersRequest) SetId(v int64) *ListProjectMembersRequest {
	s.Id = &v
	return s
}

func (s *ListProjectMembersRequest) SetListQuery(v *ListProjectMembersRequestListQuery) *ListProjectMembersRequest {
	s.ListQuery = v
	return s
}

func (s *ListProjectMembersRequest) SetOpTenantId(v int64) *ListProjectMembersRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListProjectMembersRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectMembersRequestListQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListProjectMembersRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequestListQuery) GetEnv() *string {
	return s.Env
}

func (s *ListProjectMembersRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListProjectMembersRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectMembersRequestListQuery) SetEnv(v string) *ListProjectMembersRequestListQuery {
	s.Env = &v
	return s
}

func (s *ListProjectMembersRequestListQuery) SetPageNo(v int32) *ListProjectMembersRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListProjectMembersRequestListQuery) SetPageSize(v int32) *ListProjectMembersRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersRequestListQuery) Validate() error {
	return dara.Validate(s)
}
