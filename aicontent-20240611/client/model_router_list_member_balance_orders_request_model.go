// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListMemberBalanceOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterListMemberBalanceOrdersRequest
	GetBalanceType() *string
	SetDirection(v string) *ModelRouterListMemberBalanceOrdersRequest
	GetDirection() *string
	SetOrderType(v string) *ModelRouterListMemberBalanceOrdersRequest
	GetOrderType() *string
	SetPage(v int32) *ModelRouterListMemberBalanceOrdersRequest
	GetPage() *int32
	SetPageSize(v int32) *ModelRouterListMemberBalanceOrdersRequest
	GetPageSize() *int32
}

type ModelRouterListMemberBalanceOrdersRequest struct {
	// The balance type filter. Valid values: permanent and monthly.
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The change direction filter. Valid values: in and out.
	//
	// example:
	//
	// in
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The change type filter.
	//
	// example:
	//
	// recharge
	OrderType *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ModelRouterListMemberBalanceOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListMemberBalanceOrdersRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterListMemberBalanceOrdersRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterListMemberBalanceOrdersRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModelRouterListMemberBalanceOrdersRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModelRouterListMemberBalanceOrdersRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterListMemberBalanceOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterListMemberBalanceOrdersRequest) SetBalanceType(v string) *ModelRouterListMemberBalanceOrdersRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersRequest) SetDirection(v string) *ModelRouterListMemberBalanceOrdersRequest {
	s.Direction = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersRequest) SetOrderType(v string) *ModelRouterListMemberBalanceOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersRequest) SetPage(v int32) *ModelRouterListMemberBalanceOrdersRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersRequest) SetPageSize(v int32) *ModelRouterListMemberBalanceOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterListMemberBalanceOrdersRequest) Validate() error {
	return dara.Validate(s)
}
