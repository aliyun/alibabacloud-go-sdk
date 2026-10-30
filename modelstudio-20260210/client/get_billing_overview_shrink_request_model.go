// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingOverviewShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillMonth(v string) *GetBillingOverviewShrinkRequest
	GetBillMonth() *string
	SetFilterShrink(v string) *GetBillingOverviewShrinkRequest
	GetFilterShrink() *string
	SetGranularity(v string) *GetBillingOverviewShrinkRequest
	GetGranularity() *string
	SetGroupByShrink(v string) *GetBillingOverviewShrinkRequest
	GetGroupByShrink() *string
	SetLocale(v string) *GetBillingOverviewShrinkRequest
	GetLocale() *string
	SetRegionId(v string) *GetBillingOverviewShrinkRequest
	GetRegionId() *string
	SetTimePeriodShrink(v string) *GetBillingOverviewShrinkRequest
	GetTimePeriodShrink() *string
	SetTopNum(v int32) *GetBillingOverviewShrinkRequest
	GetTopNum() *int32
	SetZeroFilter(v bool) *GetBillingOverviewShrinkRequest
	GetZeroFilter() *bool
}

type GetBillingOverviewShrinkRequest struct {
	// The billing month.
	//
	// example:
	//
	// 2026-08
	BillMonth *string `json:"billMonth,omitempty" xml:"billMonth,omitempty"`
	// The filter conditions.
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The time granularity.
	//
	// example:
	//
	// MONTH
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The list of grouping conditions. Currently, you must specify exactly one grouping dimension.
	GroupByShrink *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The language of the response. Default value: en-US.
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
	// The query time range, which includes the start and end times. Both the start and end times are included in the query range. The time format must be consistent with the granularity, and the start time cannot be later than the end time.
	TimePeriodShrink *string `json:"timePeriod,omitempty" xml:"timePeriod,omitempty"`
	// The number of groups to return. Valid values: 1 to 20. Default value: 20.
	//
	// example:
	//
	// 20
	TopNum *int32 `json:"topNum,omitempty" xml:"topNum,omitempty"`
	// Specifies whether to filter out groups with an amount of 0. Default value: true.
	//
	// example:
	//
	// true
	ZeroFilter *bool `json:"zeroFilter,omitempty" xml:"zeroFilter,omitempty"`
}

func (s GetBillingOverviewShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewShrinkRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetBillingOverviewShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetBillingOverviewShrinkRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetBillingOverviewShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *GetBillingOverviewShrinkRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetBillingOverviewShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBillingOverviewShrinkRequest) GetTimePeriodShrink() *string {
	return s.TimePeriodShrink
}

func (s *GetBillingOverviewShrinkRequest) GetTopNum() *int32 {
	return s.TopNum
}

func (s *GetBillingOverviewShrinkRequest) GetZeroFilter() *bool {
	return s.ZeroFilter
}

func (s *GetBillingOverviewShrinkRequest) SetBillMonth(v string) *GetBillingOverviewShrinkRequest {
	s.BillMonth = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetFilterShrink(v string) *GetBillingOverviewShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetGranularity(v string) *GetBillingOverviewShrinkRequest {
	s.Granularity = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetGroupByShrink(v string) *GetBillingOverviewShrinkRequest {
	s.GroupByShrink = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetLocale(v string) *GetBillingOverviewShrinkRequest {
	s.Locale = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetRegionId(v string) *GetBillingOverviewShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetTimePeriodShrink(v string) *GetBillingOverviewShrinkRequest {
	s.TimePeriodShrink = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetTopNum(v int32) *GetBillingOverviewShrinkRequest {
	s.TopNum = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) SetZeroFilter(v bool) *GetBillingOverviewShrinkRequest {
	s.ZeroFilter = &v
	return s
}

func (s *GetBillingOverviewShrinkRequest) Validate() error {
	return dara.Validate(s)
}
