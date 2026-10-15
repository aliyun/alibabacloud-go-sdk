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
	// The business ID.
	//
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order. Valid values: ASC and DESC.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The organization role ID, including predefined roles and custom roles:
	//
	// - Organization administrator (predefined role): 111111111
	//
	// - Permission management administrator (predefined role): 111111112
	//
	// - Regular user (predefined role): 111111113
	//
	// - Custom role: the role ID of the custom role
	//
	// example:
	//
	// 75b2f16f-35a5-4e90-949f-295ea14a4dc8
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
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
