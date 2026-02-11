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
	CurrentPage   *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime       *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IntervalMills *int32 `json:"IntervalMills,omitempty" xml:"IntervalMills,omitempty"`
	// This parameter is required.
	PId      *string `json:"PId,omitempty" xml:"PId,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
