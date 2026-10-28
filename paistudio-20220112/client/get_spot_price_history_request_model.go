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
	// The end time for querying the historical prices of spot instances. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Default value: null, which indicates the current time.
	//
	// example:
	//
	// 2024-12-30T09:36:46Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sort order. Default value: asc. Valid values:
	//
	// - desc: descending order.
	//
	// - asc: ascending order.
	//
	// This parameter takes effect only when you query the historical prices of Lingjun instance types.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: ***1***. This parameter takes effect only when you query the historical prices of Lingjun instance types.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. This parameter takes effect only when you query the historical prices of Lingjun instance types.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field used for sorting. Default value: GmtCreatedTime. Valid values:
	//
	// - GmtCreatedTime
	//
	// This parameter takes effect only when you query the historical prices of Lingjun instance types.
	//
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The protection period of the spot instance. Unit: hours. Note that only ECS instance types support Settings for the protection period of spot instances. Default value: 0. Valid values:
	//
	// - 1: After the spot instance is created, Alibaba Cloud ensures that the instance runs for 1 hour without automatic release. After 1 hour, the system automatically compares the bid price with the marketplace price and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// - 0: After the spot instance is created, Alibaba Cloud does not guarantee that the instance runs for 1 hour. The system automatically compares the bid price with the marketplace price and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// example:
	//
	// 0
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The start time for querying the historical prices of spot instances. The maximum value cannot exceed 7 days before the specified end time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. Default value: null, which indicates 3 days before the end time.
	//
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
