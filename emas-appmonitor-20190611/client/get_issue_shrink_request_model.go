// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetIssueShrinkRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetIssueShrinkRequest
	GetBizModule() *string
	SetDigestHash(v string) *GetIssueShrinkRequest
	GetDigestHash() *string
	SetFilterShrink(v string) *GetIssueShrinkRequest
	GetFilterShrink() *string
	SetOs(v string) *GetIssueShrinkRequest
	GetOs() *string
	SetTimeRange(v *GetIssueShrinkRequestTimeRange) *GetIssueShrinkRequest
	GetTimeRange() *GetIssueShrinkRequestTimeRange
}

type GetIssueShrinkRequest struct {
	// AppKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// h5Resource
	BizModule *string `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	// example:
	//
	// 2963475858785631
	DigestHash   *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// h5
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	TimeRange *GetIssueShrinkRequestTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
}

func (s GetIssueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIssueShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetIssueShrinkRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetIssueShrinkRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetIssueShrinkRequest) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetIssueShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetIssueShrinkRequest) GetOs() *string {
	return s.Os
}

func (s *GetIssueShrinkRequest) GetTimeRange() *GetIssueShrinkRequestTimeRange {
	return s.TimeRange
}

func (s *GetIssueShrinkRequest) SetAppKey(v int64) *GetIssueShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *GetIssueShrinkRequest) SetBizModule(v string) *GetIssueShrinkRequest {
	s.BizModule = &v
	return s
}

func (s *GetIssueShrinkRequest) SetDigestHash(v string) *GetIssueShrinkRequest {
	s.DigestHash = &v
	return s
}

func (s *GetIssueShrinkRequest) SetFilterShrink(v string) *GetIssueShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetIssueShrinkRequest) SetOs(v string) *GetIssueShrinkRequest {
	s.Os = &v
	return s
}

func (s *GetIssueShrinkRequest) SetTimeRange(v *GetIssueShrinkRequestTimeRange) *GetIssueShrinkRequest {
	s.TimeRange = v
	return s
}

func (s *GetIssueShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type GetIssueShrinkRequestTimeRange struct {
	// example:
	//
	// 2024-08-23T02:12:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Host
	Granularity *int32 `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// minute
	GranularityUnit *string `json:"GranularityUnit,omitempty" xml:"GranularityUnit,omitempty"`
	// example:
	//
	// 2024-12-18 00:00:00
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetIssueShrinkRequestTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetIssueShrinkRequestTimeRange) GoString() string {
	return s.String()
}

func (s *GetIssueShrinkRequestTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetIssueShrinkRequestTimeRange) GetGranularity() *int32 {
	return s.Granularity
}

func (s *GetIssueShrinkRequestTimeRange) GetGranularityUnit() *string {
	return s.GranularityUnit
}

func (s *GetIssueShrinkRequestTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetIssueShrinkRequestTimeRange) SetEndTime(v int64) *GetIssueShrinkRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetIssueShrinkRequestTimeRange) SetGranularity(v int32) *GetIssueShrinkRequestTimeRange {
	s.Granularity = &v
	return s
}

func (s *GetIssueShrinkRequestTimeRange) SetGranularityUnit(v string) *GetIssueShrinkRequestTimeRange {
	s.GranularityUnit = &v
	return s
}

func (s *GetIssueShrinkRequestTimeRange) SetStartTime(v int64) *GetIssueShrinkRequestTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetIssueShrinkRequestTimeRange) Validate() error {
	return dara.Validate(s)
}
