// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *SearchTracesRequest
	GetEndTime() *int64
	SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest
	GetExclusionFilters() []*SearchTracesRequestExclusionFilters
	SetMinDuration(v int64) *SearchTracesRequest
	GetMinDuration() *int64
	SetOperationName(v string) *SearchTracesRequest
	GetOperationName() *string
	SetRegionId(v string) *SearchTracesRequest
	GetRegionId() *string
	SetReverse(v bool) *SearchTracesRequest
	GetReverse() *bool
	SetServiceIp(v string) *SearchTracesRequest
	GetServiceIp() *string
	SetServiceName(v string) *SearchTracesRequest
	GetServiceName() *string
	SetStartTime(v int64) *SearchTracesRequest
	GetStartTime() *int64
	SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest
	GetTag() []*SearchTracesRequestTag
}

type SearchTracesRequest struct {
	// This parameter is required.
	EndTime          *int64                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExclusionFilters []*SearchTracesRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	MinDuration      *int64                                 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName    *string                                `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse     *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp   *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// This parameter is required.
	StartTime *int64                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag       []*SearchTracesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchTracesRequest) GetExclusionFilters() []*SearchTracesRequestExclusionFilters {
	return s.ExclusionFilters
}

func (s *SearchTracesRequest) GetMinDuration() *int64 {
	return s.MinDuration
}

func (s *SearchTracesRequest) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchTracesRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *SearchTracesRequest) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchTracesRequest) GetTag() []*SearchTracesRequestTag {
	return s.Tag
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

func (s *SearchTracesRequest) Validate() error {
	if s.ExclusionFilters != nil {
		for _, item := range s.ExclusionFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchTracesRequestExclusionFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestExclusionFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestExclusionFilters) GetKey() *string {
	return s.Key
}

func (s *SearchTracesRequestExclusionFilters) GetValue() *string {
	return s.Value
}

func (s *SearchTracesRequestExclusionFilters) SetKey(v string) *SearchTracesRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestExclusionFilters) SetValue(v string) *SearchTracesRequestExclusionFilters {
	s.Value = &v
	return s
}

func (s *SearchTracesRequestExclusionFilters) Validate() error {
	return dara.Validate(s)
}

type SearchTracesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) GetKey() *string {
	return s.Key
}

func (s *SearchTracesRequestTag) GetValue() *string {
	return s.Value
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

func (s *SearchTracesRequestTag) Validate() error {
	return dara.Validate(s)
}
