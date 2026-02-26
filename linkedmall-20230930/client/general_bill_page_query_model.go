// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralBillPageQuery interface {
	dara.Model
	String() string
	GoString() string
	SetAsc(v bool) *GeneralBillPageQuery
	GetAsc() *bool
	SetBillId(v string) *GeneralBillPageQuery
	GetBillId() *string
	SetBillPeriod(v string) *GeneralBillPageQuery
	GetBillPeriod() *string
	SetLimit(v int32) *GeneralBillPageQuery
	GetLimit() *int32
	SetOrderBy(v string) *GeneralBillPageQuery
	GetOrderBy() *string
	SetOrderDirection(v string) *GeneralBillPageQuery
	GetOrderDirection() *string
	SetPageNumber(v int32) *GeneralBillPageQuery
	GetPageNumber() *int32
	SetPageSize(v int32) *GeneralBillPageQuery
	GetPageSize() *int32
	SetShopId(v string) *GeneralBillPageQuery
	GetShopId() *string
	SetStart(v int32) *GeneralBillPageQuery
	GetStart() *int32
}

type GeneralBillPageQuery struct {
	// asc
	Asc *bool `json:"asc,omitempty" xml:"asc,omitempty"`
	// billId
	BillId *string `json:"billId,omitempty" xml:"billId,omitempty"`
	// billPeriod
	BillPeriod *string `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	// limit
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// orderBy
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// orderDirection
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// pageNumber
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// pageSize
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// shopId
	ShopId *string `json:"shopId,omitempty" xml:"shopId,omitempty"`
	// start
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GeneralBillPageQuery) String() string {
	return dara.Prettify(s)
}

func (s GeneralBillPageQuery) GoString() string {
	return s.String()
}

func (s *GeneralBillPageQuery) GetAsc() *bool {
	return s.Asc
}

func (s *GeneralBillPageQuery) GetBillId() *string {
	return s.BillId
}

func (s *GeneralBillPageQuery) GetBillPeriod() *string {
	return s.BillPeriod
}

func (s *GeneralBillPageQuery) GetLimit() *int32 {
	return s.Limit
}

func (s *GeneralBillPageQuery) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GeneralBillPageQuery) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *GeneralBillPageQuery) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GeneralBillPageQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GeneralBillPageQuery) GetShopId() *string {
	return s.ShopId
}

func (s *GeneralBillPageQuery) GetStart() *int32 {
	return s.Start
}

func (s *GeneralBillPageQuery) SetAsc(v bool) *GeneralBillPageQuery {
	s.Asc = &v
	return s
}

func (s *GeneralBillPageQuery) SetBillId(v string) *GeneralBillPageQuery {
	s.BillId = &v
	return s
}

func (s *GeneralBillPageQuery) SetBillPeriod(v string) *GeneralBillPageQuery {
	s.BillPeriod = &v
	return s
}

func (s *GeneralBillPageQuery) SetLimit(v int32) *GeneralBillPageQuery {
	s.Limit = &v
	return s
}

func (s *GeneralBillPageQuery) SetOrderBy(v string) *GeneralBillPageQuery {
	s.OrderBy = &v
	return s
}

func (s *GeneralBillPageQuery) SetOrderDirection(v string) *GeneralBillPageQuery {
	s.OrderDirection = &v
	return s
}

func (s *GeneralBillPageQuery) SetPageNumber(v int32) *GeneralBillPageQuery {
	s.PageNumber = &v
	return s
}

func (s *GeneralBillPageQuery) SetPageSize(v int32) *GeneralBillPageQuery {
	s.PageSize = &v
	return s
}

func (s *GeneralBillPageQuery) SetShopId(v string) *GeneralBillPageQuery {
	s.ShopId = &v
	return s
}

func (s *GeneralBillPageQuery) SetStart(v int32) *GeneralBillPageQuery {
	s.Start = &v
	return s
}

func (s *GeneralBillPageQuery) Validate() error {
	return dara.Validate(s)
}
