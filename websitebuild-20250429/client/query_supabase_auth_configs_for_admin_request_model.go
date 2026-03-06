// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseAuthConfigsForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *QuerySupabaseAuthConfigsForAdminRequest
	GetAuthType() *string
	SetBizId(v string) *QuerySupabaseAuthConfigsForAdminRequest
	GetBizId() *string
	SetOrderColumn(v string) *QuerySupabaseAuthConfigsForAdminRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QuerySupabaseAuthConfigsForAdminRequest
	GetOrderType() *string
	SetPageNum(v int32) *QuerySupabaseAuthConfigsForAdminRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySupabaseAuthConfigsForAdminRequest
	GetPageSize() *int32
	SetUserId(v string) *QuerySupabaseAuthConfigsForAdminRequest
	GetUserId() *string
}

type QuerySupabaseAuthConfigsForAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// key
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// gmtCreated
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
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySupabaseAuthConfigsForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseAuthConfigsForAdminRequest) GoString() string {
	return s.String()
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) GetUserId() *string {
	return s.UserId
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetAuthType(v string) *QuerySupabaseAuthConfigsForAdminRequest {
	s.AuthType = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetBizId(v string) *QuerySupabaseAuthConfigsForAdminRequest {
	s.BizId = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetOrderColumn(v string) *QuerySupabaseAuthConfigsForAdminRequest {
	s.OrderColumn = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetOrderType(v string) *QuerySupabaseAuthConfigsForAdminRequest {
	s.OrderType = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetPageNum(v int32) *QuerySupabaseAuthConfigsForAdminRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetPageSize(v int32) *QuerySupabaseAuthConfigsForAdminRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) SetUserId(v string) *QuerySupabaseAuthConfigsForAdminRequest {
	s.UserId = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminRequest) Validate() error {
	return dara.Validate(s)
}
