// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromotionOfferRecordsForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *ListPromotionOfferRecordsForPartnerRequest
	GetActivityCode() *string
	SetBelongId(v string) *ListPromotionOfferRecordsForPartnerRequest
	GetBelongId() *string
	SetMaxResults(v int32) *ListPromotionOfferRecordsForPartnerRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPromotionOfferRecordsForPartnerRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListPromotionOfferRecordsForPartnerRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListPromotionOfferRecordsForPartnerRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListPromotionOfferRecordsForPartnerRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListPromotionOfferRecordsForPartnerRequest
	GetPageSize() *int32
}

type ListPromotionOfferRecordsForPartnerRequest struct {
	// The activity code.
	//
	// example:
	//
	// acwfradoj5u
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// The belonging ID.
	//
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// test
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field used for sorting.
	//
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort type. Valid values: ASC and DESC.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPromotionOfferRecordsForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionOfferRecordsForPartnerRequest) GoString() string {
	return s.String()
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPromotionOfferRecordsForPartnerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetActivityCode(v string) *ListPromotionOfferRecordsForPartnerRequest {
	s.ActivityCode = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetBelongId(v string) *ListPromotionOfferRecordsForPartnerRequest {
	s.BelongId = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetMaxResults(v int32) *ListPromotionOfferRecordsForPartnerRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetNextToken(v string) *ListPromotionOfferRecordsForPartnerRequest {
	s.NextToken = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetOrderColumn(v string) *ListPromotionOfferRecordsForPartnerRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetOrderType(v string) *ListPromotionOfferRecordsForPartnerRequest {
	s.OrderType = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetPageNum(v int32) *ListPromotionOfferRecordsForPartnerRequest {
	s.PageNum = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) SetPageSize(v int32) *ListPromotionOfferRecordsForPartnerRequest {
	s.PageSize = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
