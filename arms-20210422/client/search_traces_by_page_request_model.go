// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *SearchTracesByPageRequest
	GetEndTime() *int64
	SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest
	GetExclusionFilters() []*SearchTracesByPageRequestExclusionFilters
	SetMinDuration(v int64) *SearchTracesByPageRequest
	GetMinDuration() *int64
	SetOperationName(v string) *SearchTracesByPageRequest
	GetOperationName() *string
	SetPageNumber(v int32) *SearchTracesByPageRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchTracesByPageRequest
	GetPageSize() *int32
	SetRegionId(v string) *SearchTracesByPageRequest
	GetRegionId() *string
	SetReverse(v bool) *SearchTracesByPageRequest
	GetReverse() *bool
	SetServiceIp(v string) *SearchTracesByPageRequest
	GetServiceIp() *string
	SetServiceName(v string) *SearchTracesByPageRequest
	GetServiceName() *string
	SetStartTime(v int64) *SearchTracesByPageRequest
	GetStartTime() *int64
}

type SearchTracesByPageRequest struct {
	// This parameter is required.
	EndTime          *int64                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExclusionFilters []*SearchTracesByPageRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	MinDuration      *int64                                       `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName    *string                                      `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse     *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp   *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchTracesByPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchTracesByPageRequest) GetExclusionFilters() []*SearchTracesByPageRequestExclusionFilters {
	return s.ExclusionFilters
}

func (s *SearchTracesByPageRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *SearchTracesByPageRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesByPageRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTracesByPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTracesByPageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTracesByPageRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *SearchTracesByPageRequest) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesByPageRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesByPageRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchTracesByPageRequest) SetEndTime(v int64) *SearchTracesByPageRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesByPageRequest) SetMinDuration(v int64) *SearchTracesByPageRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOperationName(v string) *SearchTracesByPageRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageNumber(v int32) *SearchTracesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageSize(v int32) *SearchTracesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageRequest) SetRegionId(v string) *SearchTracesByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesByPageRequest) SetReverse(v bool) *SearchTracesByPageRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceIp(v string) *SearchTracesByPageRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceName(v string) *SearchTracesByPageRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetStartTime(v int64) *SearchTracesByPageRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesByPageRequest) Validate() error {
	if s.ExclusionFilters != nil {
		for _, item := range s.ExclusionFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTracesByPageRequestExclusionFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesByPageRequestExclusionFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequestExclusionFilters) GetKey() *string {
	return s.Key
}

func (s *SearchTracesByPageRequestExclusionFilters) GetValue() *string {
	return s.Value
}

func (s *SearchTracesByPageRequestExclusionFilters) SetKey(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesByPageRequestExclusionFilters) SetValue(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Value = &v
	return s
}

func (s *SearchTracesByPageRequestExclusionFilters) Validate() error {
	return dara.Validate(s)
}
