// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateSupabaseForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *AllocateSupabaseForAdminRequest
	GetBizId() *string
	SetOrderColumn(v string) *AllocateSupabaseForAdminRequest
	GetOrderColumn() *string
	SetOrderType(v string) *AllocateSupabaseForAdminRequest
	GetOrderType() *string
	SetPageNum(v int32) *AllocateSupabaseForAdminRequest
	GetPageNum() *int32
	SetPageSize(v int32) *AllocateSupabaseForAdminRequest
	GetPageSize() *int32
	SetUserId(v string) *AllocateSupabaseForAdminRequest
	GetUserId() *string
}

type AllocateSupabaseForAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// BUY
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
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AllocateSupabaseForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateSupabaseForAdminRequest) GoString() string {
	return s.String()
}

func (s *AllocateSupabaseForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *AllocateSupabaseForAdminRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *AllocateSupabaseForAdminRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *AllocateSupabaseForAdminRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *AllocateSupabaseForAdminRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AllocateSupabaseForAdminRequest) GetUserId() *string {
	return s.UserId
}

func (s *AllocateSupabaseForAdminRequest) SetBizId(v string) *AllocateSupabaseForAdminRequest {
	s.BizId = &v
	return s
}

func (s *AllocateSupabaseForAdminRequest) SetOrderColumn(v string) *AllocateSupabaseForAdminRequest {
	s.OrderColumn = &v
	return s
}

func (s *AllocateSupabaseForAdminRequest) SetOrderType(v string) *AllocateSupabaseForAdminRequest {
	s.OrderType = &v
	return s
}

func (s *AllocateSupabaseForAdminRequest) SetPageNum(v int32) *AllocateSupabaseForAdminRequest {
	s.PageNum = &v
	return s
}

func (s *AllocateSupabaseForAdminRequest) SetPageSize(v int32) *AllocateSupabaseForAdminRequest {
	s.PageSize = &v
	return s
}

func (s *AllocateSupabaseForAdminRequest) SetUserId(v string) *AllocateSupabaseForAdminRequest {
	s.UserId = &v
	return s
}

func (s *AllocateSupabaseForAdminRequest) Validate() error {
	return dara.Validate(s)
}
