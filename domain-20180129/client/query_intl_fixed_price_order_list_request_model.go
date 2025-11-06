// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIntlFixedPriceOrderListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryIntlFixedPriceOrderListRequest
	GetBizId() *string
	SetCurrentPage(v int64) *QueryIntlFixedPriceOrderListRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *QueryIntlFixedPriceOrderListRequest
	GetPageSize() *int64
	SetStatus(v int64) *QueryIntlFixedPriceOrderListRequest
	GetStatus() *int64
}

type QueryIntlFixedPriceOrderListRequest struct {
	// example:
	//
	// T2024061115213700****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryIntlFixedPriceOrderListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIntlFixedPriceOrderListRequest) GoString() string {
	return s.String()
}

func (s *QueryIntlFixedPriceOrderListRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryIntlFixedPriceOrderListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QueryIntlFixedPriceOrderListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryIntlFixedPriceOrderListRequest) GetStatus() *int64 {
	return s.Status
}

func (s *QueryIntlFixedPriceOrderListRequest) SetBizId(v string) *QueryIntlFixedPriceOrderListRequest {
	s.BizId = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListRequest) SetCurrentPage(v int64) *QueryIntlFixedPriceOrderListRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListRequest) SetPageSize(v int64) *QueryIntlFixedPriceOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListRequest) SetStatus(v int64) *QueryIntlFixedPriceOrderListRequest {
	s.Status = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListRequest) Validate() error {
	return dara.Validate(s)
}
