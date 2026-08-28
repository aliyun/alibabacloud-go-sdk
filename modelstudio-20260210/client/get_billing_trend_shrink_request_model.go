// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingTrendShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *GetBillingTrendShrinkRequest
	GetFilterShrink() *string
	SetGranularity(v string) *GetBillingTrendShrinkRequest
	GetGranularity() *string
	SetGroupByShrink(v string) *GetBillingTrendShrinkRequest
	GetGroupByShrink() *string
	SetLocale(v string) *GetBillingTrendShrinkRequest
	GetLocale() *string
	SetRegionId(v string) *GetBillingTrendShrinkRequest
	GetRegionId() *string
	SetTimePeriodShrink(v string) *GetBillingTrendShrinkRequest
	GetTimePeriodShrink() *string
	SetTopNum(v int32) *GetBillingTrendShrinkRequest
	GetTopNum() *int32
	SetZeroFilter(v bool) *GetBillingTrendShrinkRequest
	GetZeroFilter() *bool
}

type GetBillingTrendShrinkRequest struct {
	// The dimension filter conditions.
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The query granularity. This parameter is required.
	//
	// example:
	//
	// DAY
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The grouping conditions. This parameter must contain one and only one element.
	GroupByShrink *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The response language. Default value: en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The query time range, including the start time and end time. This parameter is required.
	TimePeriodShrink *string `json:"timePeriod,omitempty" xml:"timePeriod,omitempty"`
	// The number of groups to return. Valid values: 1 to 20. Default value: 20. The remaining groups are merged into "Others".
	//
	// example:
	//
	// 20
	TopNum *int32 `json:"topNum,omitempty" xml:"topNum,omitempty"`
	// Specifies whether to filter out groups with a zero amount. Default value: true.
	//
	// example:
	//
	// true
	ZeroFilter *bool `json:"zeroFilter,omitempty" xml:"zeroFilter,omitempty"`
}

func (s GetBillingTrendShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetBillingTrendShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetBillingTrendShrinkRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetBillingTrendShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *GetBillingTrendShrinkRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetBillingTrendShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBillingTrendShrinkRequest) GetTimePeriodShrink() *string {
	return s.TimePeriodShrink
}

func (s *GetBillingTrendShrinkRequest) GetTopNum() *int32 {
	return s.TopNum
}

func (s *GetBillingTrendShrinkRequest) GetZeroFilter() *bool {
	return s.ZeroFilter
}

func (s *GetBillingTrendShrinkRequest) SetFilterShrink(v string) *GetBillingTrendShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetGranularity(v string) *GetBillingTrendShrinkRequest {
	s.Granularity = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetGroupByShrink(v string) *GetBillingTrendShrinkRequest {
	s.GroupByShrink = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetLocale(v string) *GetBillingTrendShrinkRequest {
	s.Locale = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetRegionId(v string) *GetBillingTrendShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetTimePeriodShrink(v string) *GetBillingTrendShrinkRequest {
	s.TimePeriodShrink = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetTopNum(v int32) *GetBillingTrendShrinkRequest {
	s.TopNum = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) SetZeroFilter(v bool) *GetBillingTrendShrinkRequest {
	s.ZeroFilter = &v
	return s
}

func (s *GetBillingTrendShrinkRequest) Validate() error {
	return dara.Validate(s)
}
