// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSupabaseForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *OperateSupabaseForAdminRequest
	GetBizId() *string
	SetExecuteSql(v string) *OperateSupabaseForAdminRequest
	GetExecuteSql() *string
	SetOperateType(v string) *OperateSupabaseForAdminRequest
	GetOperateType() *string
	SetOrderByClause(v string) *OperateSupabaseForAdminRequest
	GetOrderByClause() *string
	SetOrderColumn(v string) *OperateSupabaseForAdminRequest
	GetOrderColumn() *string
	SetOrderType(v string) *OperateSupabaseForAdminRequest
	GetOrderType() *string
	SetPageNum(v int32) *OperateSupabaseForAdminRequest
	GetPageNum() *int32
	SetPageSize(v int32) *OperateSupabaseForAdminRequest
	GetPageSize() *int32
	SetTableName(v string) *OperateSupabaseForAdminRequest
	GetTableName() *string
	SetUserId(v string) *OperateSupabaseForAdminRequest
	GetUserId() *string
	SetWhereClause(v string) *OperateSupabaseForAdminRequest
	GetWhereClause() *string
}

type OperateSupabaseForAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// select 	- from profiles
	ExecuteSql *string `json:"ExecuteSql,omitempty" xml:"ExecuteSql,omitempty"`
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// null
	OrderByClause *string `json:"OrderByClause,omitempty" xml:"OrderByClause,omitempty"`
	// example:
	//
	// CreationTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// default.ai_advertising_material_rec_train_v1103
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// null
	WhereClause *string `json:"WhereClause,omitempty" xml:"WhereClause,omitempty"`
}

func (s OperateSupabaseForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateSupabaseForAdminRequest) GoString() string {
	return s.String()
}

func (s *OperateSupabaseForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *OperateSupabaseForAdminRequest) GetExecuteSql() *string {
	return s.ExecuteSql
}

func (s *OperateSupabaseForAdminRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *OperateSupabaseForAdminRequest) GetOrderByClause() *string {
	return s.OrderByClause
}

func (s *OperateSupabaseForAdminRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *OperateSupabaseForAdminRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *OperateSupabaseForAdminRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *OperateSupabaseForAdminRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OperateSupabaseForAdminRequest) GetTableName() *string {
	return s.TableName
}

func (s *OperateSupabaseForAdminRequest) GetUserId() *string {
	return s.UserId
}

func (s *OperateSupabaseForAdminRequest) GetWhereClause() *string {
	return s.WhereClause
}

func (s *OperateSupabaseForAdminRequest) SetBizId(v string) *OperateSupabaseForAdminRequest {
	s.BizId = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetExecuteSql(v string) *OperateSupabaseForAdminRequest {
	s.ExecuteSql = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetOperateType(v string) *OperateSupabaseForAdminRequest {
	s.OperateType = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetOrderByClause(v string) *OperateSupabaseForAdminRequest {
	s.OrderByClause = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetOrderColumn(v string) *OperateSupabaseForAdminRequest {
	s.OrderColumn = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetOrderType(v string) *OperateSupabaseForAdminRequest {
	s.OrderType = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetPageNum(v int32) *OperateSupabaseForAdminRequest {
	s.PageNum = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetPageSize(v int32) *OperateSupabaseForAdminRequest {
	s.PageSize = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetTableName(v string) *OperateSupabaseForAdminRequest {
	s.TableName = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetUserId(v string) *OperateSupabaseForAdminRequest {
	s.UserId = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) SetWhereClause(v string) *OperateSupabaseForAdminRequest {
	s.WhereClause = &v
	return s
}

func (s *OperateSupabaseForAdminRequest) Validate() error {
	return dara.Validate(s)
}
