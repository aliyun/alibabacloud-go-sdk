// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseInstanceInfoForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetBizId() *string
	SetEnv(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetEnv() *string
	SetOrderColumn(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetOrderType() *string
	SetPageNum(v int32) *QuerySupabaseInstanceInfoForAdminRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySupabaseInstanceInfoForAdminRequest
	GetPageSize() *int32
	SetUserId(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetUserId() *string
}

type QuerySupabaseInstanceInfoForAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staging
	Env         *string `json:"Env,omitempty" xml:"Env,omitempty"`
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	OrderType   *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySupabaseInstanceInfoForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseInstanceInfoForAdminRequest) GoString() string {
	return s.String()
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetEnv() *string {
	return s.Env
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetUserId() *string {
	return s.UserId
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetBizId(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.BizId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetEnv(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.Env = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetOrderColumn(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.OrderColumn = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetOrderType(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.OrderType = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetPageNum(v int32) *QuerySupabaseInstanceInfoForAdminRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetPageSize(v int32) *QuerySupabaseInstanceInfoForAdminRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetUserId(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.UserId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) Validate() error {
	return dara.Validate(s)
}
