// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacRolePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryRbacRolePermissionsRequest
	GetBizId() *string
	SetOrderColumn(v string) *QueryRbacRolePermissionsRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryRbacRolePermissionsRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryRbacRolePermissionsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryRbacRolePermissionsRequest
	GetPageSize() *int32
	SetRoleId(v string) *QueryRbacRolePermissionsRequest
	GetRoleId() *string
}

type QueryRbacRolePermissionsRequest struct {
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	OrderType   *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RoleId      *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s QueryRbacRolePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRolePermissionsRequest) GoString() string {
	return s.String()
}

func (s *QueryRbacRolePermissionsRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryRbacRolePermissionsRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryRbacRolePermissionsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryRbacRolePermissionsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRbacRolePermissionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRbacRolePermissionsRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *QueryRbacRolePermissionsRequest) SetBizId(v string) *QueryRbacRolePermissionsRequest {
	s.BizId = &v
	return s
}

func (s *QueryRbacRolePermissionsRequest) SetOrderColumn(v string) *QueryRbacRolePermissionsRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryRbacRolePermissionsRequest) SetOrderType(v string) *QueryRbacRolePermissionsRequest {
	s.OrderType = &v
	return s
}

func (s *QueryRbacRolePermissionsRequest) SetPageNum(v int32) *QueryRbacRolePermissionsRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRbacRolePermissionsRequest) SetPageSize(v int32) *QueryRbacRolePermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRbacRolePermissionsRequest) SetRoleId(v string) *QueryRbacRolePermissionsRequest {
	s.RoleId = &v
	return s
}

func (s *QueryRbacRolePermissionsRequest) Validate() error {
	return dara.Validate(s)
}
