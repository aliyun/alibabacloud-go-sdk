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
	// The end time for querying historical spot instance prices. Specify the time in ISO 8601 format using UTC+0, as yyyy-MM-ddTHH:mm:ssZ. Default value: empty. An empty value means the current time.
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
	// This parameter applies only when you query historical prices for Lingjun instance types.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the current page. Default value: ***1***. This parameter applies only when you query historical prices for Lingjun instance types.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. This parameter applies only when you query historical prices for Lingjun instance types.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field to sort by. Default value: GmtCreatedTime. Valid values:
	//
	// - GmtCreatedTime
	//
	// This parameter applies only when you query historical prices for Lingjun instance types.
	//
	// example:
	//
	// GmtCreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The retention period for the spot instance, in hours. Note that only ECS instance types support this parameter. Default value: 0. Valid values:
	//
	// - 1: Alibaba Cloud guarantees that the instance runs for at least one hour after creation. After one hour, the system compares your bid price with the current market price and checks resource inventory to decide whether to retain or revoke the instance.
	//
	// - 0: Alibaba Cloud does not guarantee one-hour runtime. The system compares your bid price with the current market price and checks resource inventory to decide whether to retain or revoke the instance.
	//
	// example:
	//
	// 0
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The start time for querying historical spot instance prices. This time must be no more than seven days before the end time. Specify the time in ISO 8601 format using UTC+0, as yyyy-MM-ddTHH:mm:ssZ. Default value: empty. An empty value means three days before the end time.
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
