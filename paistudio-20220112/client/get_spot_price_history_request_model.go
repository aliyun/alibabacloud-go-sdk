// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpotPriceHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetSpotPriceHistoryRequest
	GetEndTime() *string
	SetOrder(v string) *GetSpotPriceHistoryRequest
	GetOrder() *string
	SetPageNumber(v int32) *GetSpotPriceHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetSpotPriceHistoryRequest
	GetPageSize() *int32
	SetSortBy(v string) *GetSpotPriceHistoryRequest
	GetSortBy() *string
	SetSpotDuration(v int32) *GetSpotPriceHistoryRequest
	GetSpotDuration() *int32
	SetStartTime(v string) *GetSpotPriceHistoryRequest
	GetStartTime() *string
}

type GetSpotPriceHistoryRequest struct {
	// example:
	//
	// 2024-12-30T09:36:46Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 0
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// example:
	//
	// 2024-12-26T09:36:46Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetSpotPriceHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpotPriceHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetSpotPriceHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSpotPriceHistoryRequest) GetOrder() *string {
	return s.Order
}

func (s *GetSpotPriceHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetSpotPriceHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSpotPriceHistoryRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetSpotPriceHistoryRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *GetSpotPriceHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSpotPriceHistoryRequest) SetEndTime(v string) *GetSpotPriceHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) SetOrder(v string) *GetSpotPriceHistoryRequest {
	s.Order = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) SetPageNumber(v int32) *GetSpotPriceHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) SetPageSize(v int32) *GetSpotPriceHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) SetSortBy(v string) *GetSpotPriceHistoryRequest {
	s.SortBy = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) SetSpotDuration(v int32) *GetSpotPriceHistoryRequest {
	s.SpotDuration = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) SetStartTime(v string) *GetSpotPriceHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *GetSpotPriceHistoryRequest) Validate() error {
	return dara.Validate(s)
}
