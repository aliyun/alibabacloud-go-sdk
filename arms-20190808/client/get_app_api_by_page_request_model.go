// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppApiByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetAppApiByPageRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetAppApiByPageRequest
	GetEndTime() *int64
	SetIntervalMills(v int32) *GetAppApiByPageRequest
	GetIntervalMills() *int32
	SetPId(v string) *GetAppApiByPageRequest
	GetPId() *string
	SetPageSize(v int32) *GetAppApiByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetAppApiByPageRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetAppApiByPageRequest
	GetStartTime() *int64
}

type GetAppApiByPageRequest struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1600066800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between the data shards to be queried. Unit: milliseconds. Minimum value: 60000. Maximum value: 2147483647.
	//
	// example:
	//
	// 60000
	IntervalMills *int32 `json:"IntervalMills,omitempty" xml:"IntervalMills,omitempty"`
	// The process identifier (PID) of the application. For information about how to obtain a PID, see [Obtain the PID of an application](https://www.alibabacloud.com/help/zh/doc-detail/186100.htm?spm=a2cdw.13409063.0.0.7a72281f0bkTfx#title-imy-7gj-qhr).
	//
	// This parameter is required.
	//
	// example:
	//
	// a2n80plglh@745eddxxx
	PId *string `json:"PId,omitempty" xml:"PId,omitempty"`
	// The number of entries to return on each page. This parameter is no longer supported. The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1600063200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppApiByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppApiByPageRequest) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetAppApiByPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAppApiByPageRequest) GetIntervalMills() *int32 {
	return s.IntervalMills
}

func (s *GetAppApiByPageRequest) GetPId() *string {
	return s.PId
}

func (s *GetAppApiByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAppApiByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAppApiByPageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAppApiByPageRequest) SetCurrentPage(v int32) *GetAppApiByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAppApiByPageRequest) SetEndTime(v int64) *GetAppApiByPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppApiByPageRequest) SetIntervalMills(v int32) *GetAppApiByPageRequest {
	s.IntervalMills = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPId(v string) *GetAppApiByPageRequest {
	s.PId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPageSize(v int32) *GetAppApiByPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageRequest) SetRegionId(v string) *GetAppApiByPageRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetStartTime(v int64) *GetAppApiByPageRequest {
	s.StartTime = &v
	return s
}

func (s *GetAppApiByPageRequest) Validate() error {
	return dara.Validate(s)
}
