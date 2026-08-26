// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMonitorListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorId(v string) *DescribeLiveStreamMonitorListRequest
	GetMonitorId() *string
	SetOrderRule(v int32) *DescribeLiveStreamMonitorListRequest
	GetOrderRule() *int32
	SetOwnerId(v int64) *DescribeLiveStreamMonitorListRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeLiveStreamMonitorListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeLiveStreamMonitorListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLiveStreamMonitorListRequest
	GetRegionId() *string
	SetStatus(v int32) *DescribeLiveStreamMonitorListRequest
	GetStatus() *int32
}

type DescribeLiveStreamMonitorListRequest struct {
	// The ID of the monitoring session.
	//
	// > Obtain the MonitorId value from the response parameters of the [CreateLiveStreamMonitor](https://help.aliyun.com/document_detail/2848129.html) operation. If you leave this parameter empty, the data of all monitoring sessions is returned.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	MonitorId *string `json:"MonitorId,omitempty" xml:"MonitorId,omitempty"`
	// The order in which to sort the monitoring sessions. Valid values:
	//
	// - 0: Default. The monitoring sessions are sorted by monitoring status in descending order (active sessions are listed first). The start time is not used for sorting.
	//
	// - 1: The monitoring sessions are sorted by start time in descending order.
	//
	// - 2: The monitoring sessions are sorted by start time in ascending order.
	//
	// example:
	//
	// 1
	OrderRule *int32 `json:"OrderRule,omitempty" xml:"OrderRule,omitempty"`
	OwnerId   *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 2
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of monitoring sessions to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the monitoring session. Valid values:
	//
	// - 1: The session is being monitored.
	//
	// - 0: The session is not being monitored.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLiveStreamMonitorListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListRequest) GetMonitorId() *string {
	return s.MonitorId
}

func (s *DescribeLiveStreamMonitorListRequest) GetOrderRule() *int32 {
	return s.OrderRule
}

func (s *DescribeLiveStreamMonitorListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamMonitorListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeLiveStreamMonitorListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamMonitorListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamMonitorListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeLiveStreamMonitorListRequest) SetMonitorId(v string) *DescribeLiveStreamMonitorListRequest {
	s.MonitorId = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) SetOrderRule(v int32) *DescribeLiveStreamMonitorListRequest {
	s.OrderRule = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) SetOwnerId(v int64) *DescribeLiveStreamMonitorListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) SetPageNum(v int32) *DescribeLiveStreamMonitorListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) SetPageSize(v int32) *DescribeLiveStreamMonitorListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) SetRegionId(v string) *DescribeLiveStreamMonitorListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) SetStatus(v int32) *DescribeLiveStreamMonitorListRequest {
	s.Status = &v
	return s
}

func (s *DescribeLiveStreamMonitorListRequest) Validate() error {
	return dara.Validate(s)
}
