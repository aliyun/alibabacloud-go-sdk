// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListUserProjectsRequestListQuery) *ListUserProjectsRequest
	GetListQuery() *ListUserProjectsRequestListQuery
	SetOpTenantId(v int64) *ListUserProjectsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListUserProjectsRequest
	GetOpUserId() *string
}

type ListUserProjectsRequest struct {
	// The query conditions.
	//
	// This parameter is required.
	ListQuery *ListUserProjectsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListUserProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProjectsRequest) GetListQuery() *ListUserProjectsRequestListQuery {
	return s.ListQuery
}

func (s *ListUserProjectsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListUserProjectsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListUserProjectsRequest) SetListQuery(v *ListUserProjectsRequestListQuery) *ListUserProjectsRequest {
	s.ListQuery = v
	return s
}

func (s *ListUserProjectsRequest) SetOpTenantId(v int64) *ListUserProjectsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListUserProjectsRequest) SetOpUserId(v string) *ListUserProjectsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListUserProjectsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserProjectsRequestListQuery struct {
	// The environment identifier.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the target user to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100023
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s ListUserProjectsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListUserProjectsRequestListQuery) GetEnv() *string {
	return s.Env
}

func (s *ListUserProjectsRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListUserProjectsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserProjectsRequestListQuery) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *ListUserProjectsRequestListQuery) SetEnv(v string) *ListUserProjectsRequestListQuery {
	s.Env = &v
	return s
}

func (s *ListUserProjectsRequestListQuery) SetPageNo(v int32) *ListUserProjectsRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListUserProjectsRequestListQuery) SetPageSize(v int32) *ListUserProjectsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListUserProjectsRequestListQuery) SetTargetUserId(v string) *ListUserProjectsRequestListQuery {
	s.TargetUserId = &v
	return s
}

func (s *ListUserProjectsRequestListQuery) Validate() error {
	return dara.Validate(s)
}
