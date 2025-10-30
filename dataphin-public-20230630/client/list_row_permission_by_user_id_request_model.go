// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionByUserIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListRowPermissionByUserIdQuery(v *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) *ListRowPermissionByUserIdRequest
	GetListRowPermissionByUserIdQuery() *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery
	SetOpTenantId(v int64) *ListRowPermissionByUserIdRequest
	GetOpTenantId() *int64
}

type ListRowPermissionByUserIdRequest struct {
	// This parameter is required.
	ListRowPermissionByUserIdQuery *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery `json:"ListRowPermissionByUserIdQuery,omitempty" xml:"ListRowPermissionByUserIdQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListRowPermissionByUserIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdRequest) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdRequest) GetListRowPermissionByUserIdQuery() *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery {
	return s.ListRowPermissionByUserIdQuery
}

func (s *ListRowPermissionByUserIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListRowPermissionByUserIdRequest) SetListRowPermissionByUserIdQuery(v *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) *ListRowPermissionByUserIdRequest {
	s.ListRowPermissionByUserIdQuery = v
	return s
}

func (s *ListRowPermissionByUserIdRequest) SetOpTenantId(v int64) *ListRowPermissionByUserIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListRowPermissionByUserIdRequest) Validate() error {
	if s.ListRowPermissionByUserIdQuery != nil {
		if err := s.ListRowPermissionByUserIdQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 30008888
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) GoString() string {
	return s.String()
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) GetOperator() *string {
	return s.Operator
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) SetOperator(v string) *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery {
	s.Operator = &v
	return s
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) SetPageNum(v int32) *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery {
	s.PageNum = &v
	return s
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) SetPageSize(v int32) *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery {
	s.PageSize = &v
	return s
}

func (s *ListRowPermissionByUserIdRequestListRowPermissionByUserIdQuery) Validate() error {
	return dara.Validate(s)
}
