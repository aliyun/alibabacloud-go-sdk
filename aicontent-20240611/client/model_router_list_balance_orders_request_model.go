// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListBalanceOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterListBalanceOrdersRequest
	GetBalanceType() *string
	SetDirection(v string) *ModelRouterListBalanceOrdersRequest
	GetDirection() *string
	SetMaxResults(v int32) *ModelRouterListBalanceOrdersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterListBalanceOrdersRequest
	GetNextToken() *string
	SetOrderType(v string) *ModelRouterListBalanceOrdersRequest
	GetOrderType() *string
	SetPage(v int32) *ModelRouterListBalanceOrdersRequest
	GetPage() *int32
	SetPageSize(v int32) *ModelRouterListBalanceOrdersRequest
	GetPageSize() *int32
}

type ModelRouterListBalanceOrdersRequest struct {
	// The balance type filter. Valid values: permanent, monthly. If this parameter is left empty, all types are queried.
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The direction filter. Valid values: in (income), out (expenditure). If this parameter is left empty, all directions are queried.
	//
	// example:
	//
	// in
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// The maximum number of results.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The change type filter. Valid values: recharge, periodic_recharge, manual_deduct, transfer_out, transfer_in, return_out, return_in, write_off, monthly_expire, and deficit_writeoff. If this parameter is left empty, all types are queried.
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

func (s ModelRouterListBalanceOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListBalanceOrdersRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterListBalanceOrdersRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterListBalanceOrdersRequest) GetDirection() *string {
	return s.Direction
}

func (s *ModelRouterListBalanceOrdersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterListBalanceOrdersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterListBalanceOrdersRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ModelRouterListBalanceOrdersRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterListBalanceOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterListBalanceOrdersRequest) SetBalanceType(v string) *ModelRouterListBalanceOrdersRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) SetDirection(v string) *ModelRouterListBalanceOrdersRequest {
	s.Direction = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) SetMaxResults(v int32) *ModelRouterListBalanceOrdersRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) SetNextToken(v string) *ModelRouterListBalanceOrdersRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) SetOrderType(v string) *ModelRouterListBalanceOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) SetPage(v int32) *ModelRouterListBalanceOrdersRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) SetPageSize(v int32) *ModelRouterListBalanceOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *ModelRouterListBalanceOrdersRequest) Validate() error {
	return dara.Validate(s)
}
