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
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// DAY
	Granularity   *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	GroupByShrink *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// example:
	//
	// zh-CN
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId         *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	TimePeriodShrink *string `json:"timePeriod,omitempty" xml:"timePeriod,omitempty"`
	// example:
	//
	// 20
	TopNum *int32 `json:"topNum,omitempty" xml:"topNum,omitempty"`
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
