// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerlessTopNAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListServerlessTopNAppsRequest
	GetEndTime() *int64
	SetLimit(v int32) *ListServerlessTopNAppsRequest
	GetLimit() *int32
	SetOrderBy(v string) *ListServerlessTopNAppsRequest
	GetOrderBy() *string
	SetRegionId(v string) *ListServerlessTopNAppsRequest
	GetRegionId() *string
	SetStartTime(v int64) *ListServerlessTopNAppsRequest
	GetStartTime() *int64
}

type ListServerlessTopNAppsRequest struct {
	// This parameter is required.
	EndTime *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Limit   *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListServerlessTopNAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServerlessTopNAppsRequest) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListServerlessTopNAppsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListServerlessTopNAppsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListServerlessTopNAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServerlessTopNAppsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListServerlessTopNAppsRequest) SetEndTime(v int64) *ListServerlessTopNAppsRequest {
	s.EndTime = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetLimit(v int32) *ListServerlessTopNAppsRequest {
	s.Limit = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetOrderBy(v string) *ListServerlessTopNAppsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetRegionId(v string) *ListServerlessTopNAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) SetStartTime(v int64) *ListServerlessTopNAppsRequest {
	s.StartTime = &v
	return s
}

func (s *ListServerlessTopNAppsRequest) Validate() error {
	return dara.Validate(s)
}
