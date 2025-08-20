// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsLifecycleStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListApsLifecycleStrategyRequest
	GetDBClusterId() *string
	SetEndTime(v string) *ListApsLifecycleStrategyRequest
	GetEndTime() *string
	SetPageNumber(v string) *ListApsLifecycleStrategyRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListApsLifecycleStrategyRequest
	GetPageSize() *string
	SetRegionId(v string) *ListApsLifecycleStrategyRequest
	GetRegionId() *string
	SetStartTime(v string) *ListApsLifecycleStrategyRequest
	GetStartTime() *string
}

type ListApsLifecycleStrategyRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-01-02T11:22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2024-01-01T11:22Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListApsLifecycleStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApsLifecycleStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListApsLifecycleStrategyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsLifecycleStrategyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListApsLifecycleStrategyRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListApsLifecycleStrategyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListApsLifecycleStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApsLifecycleStrategyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListApsLifecycleStrategyRequest) SetDBClusterId(v string) *ListApsLifecycleStrategyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListApsLifecycleStrategyRequest) SetEndTime(v string) *ListApsLifecycleStrategyRequest {
	s.EndTime = &v
	return s
}

func (s *ListApsLifecycleStrategyRequest) SetPageNumber(v string) *ListApsLifecycleStrategyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApsLifecycleStrategyRequest) SetPageSize(v string) *ListApsLifecycleStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListApsLifecycleStrategyRequest) SetRegionId(v string) *ListApsLifecycleStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ListApsLifecycleStrategyRequest) SetStartTime(v string) *ListApsLifecycleStrategyRequest {
	s.StartTime = &v
	return s
}

func (s *ListApsLifecycleStrategyRequest) Validate() error {
	return dara.Validate(s)
}
