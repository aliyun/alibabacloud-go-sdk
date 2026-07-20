// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacUserRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationUserId(v string) *QueryRbacUserRolesRequest
	GetApplicationUserId() *string
	SetBizId(v string) *QueryRbacUserRolesRequest
	GetBizId() *string
	SetOrderColumn(v string) *QueryRbacUserRolesRequest
	GetOrderColumn() *string
	SetOrderType(v string) *QueryRbacUserRolesRequest
	GetOrderType() *string
	SetPageNum(v int32) *QueryRbacUserRolesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryRbacUserRolesRequest
	GetPageSize() *int32
}

type QueryRbacUserRolesRequest struct {
	ApplicationUserId *string `json:"ApplicationUserId,omitempty" xml:"ApplicationUserId,omitempty"`
	BizId             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OrderColumn       *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	OrderType         *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryRbacUserRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacUserRolesRequest) GoString() string {
	return s.String()
}

func (s *QueryRbacUserRolesRequest) GetApplicationUserId() *string {
	return s.ApplicationUserId
}

func (s *QueryRbacUserRolesRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryRbacUserRolesRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *QueryRbacUserRolesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryRbacUserRolesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryRbacUserRolesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRbacUserRolesRequest) SetApplicationUserId(v string) *QueryRbacUserRolesRequest {
	s.ApplicationUserId = &v
	return s
}

func (s *QueryRbacUserRolesRequest) SetBizId(v string) *QueryRbacUserRolesRequest {
	s.BizId = &v
	return s
}

func (s *QueryRbacUserRolesRequest) SetOrderColumn(v string) *QueryRbacUserRolesRequest {
	s.OrderColumn = &v
	return s
}

func (s *QueryRbacUserRolesRequest) SetOrderType(v string) *QueryRbacUserRolesRequest {
	s.OrderType = &v
	return s
}

func (s *QueryRbacUserRolesRequest) SetPageNum(v int32) *QueryRbacUserRolesRequest {
	s.PageNum = &v
	return s
}

func (s *QueryRbacUserRolesRequest) SetPageSize(v int32) *QueryRbacUserRolesRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRbacUserRolesRequest) Validate() error {
	return dara.Validate(s)
}
