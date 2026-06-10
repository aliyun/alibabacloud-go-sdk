// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseConfigsForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QuerySupabaseConfigsForAdminRequest
	GetBizId() *string
	SetEnv(v string) *QuerySupabaseConfigsForAdminRequest
	GetEnv() *string
	SetOrderColumn(v string) *QuerySupabaseConfigsForAdminRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QuerySupabaseConfigsForAdminRequest
	GetOrderType() *string
	SetPageNum(v int32) *QuerySupabaseConfigsForAdminRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySupabaseConfigsForAdminRequest
	GetPageSize() *int32
	SetUserId(v string) *QuerySupabaseConfigsForAdminRequest
	GetUserId() *string
}

type QuerySupabaseConfigsForAdminRequest struct {
	// Business ID
	//
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Env   *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// Sorting field
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// Sorting type: ASC or DESC
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Page number, default is 1
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Page size, default is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// User ID
	//
	// example:
	//
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySupabaseConfigsForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseConfigsForAdminRequest) GoString() string {
	return s.String()
}

func (s *QuerySupabaseConfigsForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *QuerySupabaseConfigsForAdminRequest) GetEnv() *string {
	return s.Env
}

func (s *QuerySupabaseConfigsForAdminRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QuerySupabaseConfigsForAdminRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QuerySupabaseConfigsForAdminRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySupabaseConfigsForAdminRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySupabaseConfigsForAdminRequest) GetUserId() *string {
	return s.UserId
}

func (s *QuerySupabaseConfigsForAdminRequest) SetBizId(v string) *QuerySupabaseConfigsForAdminRequest {
	s.BizId = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) SetEnv(v string) *QuerySupabaseConfigsForAdminRequest {
	s.Env = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) SetOrderColumn(v string) *QuerySupabaseConfigsForAdminRequest {
	s.OrderColumn = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) SetOrderType(v string) *QuerySupabaseConfigsForAdminRequest {
	s.OrderType = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) SetPageNum(v int32) *QuerySupabaseConfigsForAdminRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) SetPageSize(v int32) *QuerySupabaseConfigsForAdminRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) SetUserId(v string) *QuerySupabaseConfigsForAdminRequest {
	s.UserId = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminRequest) Validate() error {
	return dara.Validate(s)
}
