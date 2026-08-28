// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *GetBillingTrendRequestFilter) *GetBillingTrendRequest
	GetFilter() *GetBillingTrendRequestFilter
	SetGranularity(v string) *GetBillingTrendRequest
	GetGranularity() *string
	SetGroupBy(v []*GetBillingTrendRequestGroupBy) *GetBillingTrendRequest
	GetGroupBy() []*GetBillingTrendRequestGroupBy
	SetLocale(v string) *GetBillingTrendRequest
	GetLocale() *string
	SetRegionId(v string) *GetBillingTrendRequest
	GetRegionId() *string
	SetTimePeriod(v *GetBillingTrendRequestTimePeriod) *GetBillingTrendRequest
	GetTimePeriod() *GetBillingTrendRequestTimePeriod
	SetTopNum(v int32) *GetBillingTrendRequest
	GetTopNum() *int32
	SetZeroFilter(v bool) *GetBillingTrendRequest
	GetZeroFilter() *bool
}

type GetBillingTrendRequest struct {
	// The dimension filter conditions.
	Filter *GetBillingTrendRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The query granularity. This parameter is required.
	//
	// example:
	//
	// DAY
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The grouping conditions. This parameter must contain one and only one element.
	GroupBy []*GetBillingTrendRequestGroupBy `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
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
	TimePeriod *GetBillingTrendRequestTimePeriod `json:"timePeriod,omitempty" xml:"timePeriod,omitempty" type:"Struct"`
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

func (s GetBillingTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendRequest) GoString() string {
	return s.String()
}

func (s *GetBillingTrendRequest) GetFilter() *GetBillingTrendRequestFilter {
	return s.Filter
}

func (s *GetBillingTrendRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetBillingTrendRequest) GetGroupBy() []*GetBillingTrendRequestGroupBy {
	return s.GroupBy
}

func (s *GetBillingTrendRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetBillingTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBillingTrendRequest) GetTimePeriod() *GetBillingTrendRequestTimePeriod {
	return s.TimePeriod
}

func (s *GetBillingTrendRequest) GetTopNum() *int32 {
	return s.TopNum
}

func (s *GetBillingTrendRequest) GetZeroFilter() *bool {
	return s.ZeroFilter
}

func (s *GetBillingTrendRequest) SetFilter(v *GetBillingTrendRequestFilter) *GetBillingTrendRequest {
	s.Filter = v
	return s
}

func (s *GetBillingTrendRequest) SetGranularity(v string) *GetBillingTrendRequest {
	s.Granularity = &v
	return s
}

func (s *GetBillingTrendRequest) SetGroupBy(v []*GetBillingTrendRequestGroupBy) *GetBillingTrendRequest {
	s.GroupBy = v
	return s
}

func (s *GetBillingTrendRequest) SetLocale(v string) *GetBillingTrendRequest {
	s.Locale = &v
	return s
}

func (s *GetBillingTrendRequest) SetRegionId(v string) *GetBillingTrendRequest {
	s.RegionId = &v
	return s
}

func (s *GetBillingTrendRequest) SetTimePeriod(v *GetBillingTrendRequestTimePeriod) *GetBillingTrendRequest {
	s.TimePeriod = v
	return s
}

func (s *GetBillingTrendRequest) SetTopNum(v int32) *GetBillingTrendRequest {
	s.TopNum = &v
	return s
}

func (s *GetBillingTrendRequest) SetZeroFilter(v bool) *GetBillingTrendRequest {
	s.ZeroFilter = &v
	return s
}

func (s *GetBillingTrendRequest) Validate() error {
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		for _, item := range s.GroupBy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TimePeriod != nil {
		if err := s.TimePeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBillingTrendRequestFilter struct {
	// The dimension filter list.
	Dimensions []*GetBillingTrendRequestFilterDimensions `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
}

func (s GetBillingTrendRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendRequestFilter) GoString() string {
	return s.String()
}

func (s *GetBillingTrendRequestFilter) GetDimensions() []*GetBillingTrendRequestFilterDimensions {
	return s.Dimensions
}

func (s *GetBillingTrendRequestFilter) SetDimensions(v []*GetBillingTrendRequestFilterDimensions) *GetBillingTrendRequestFilter {
	s.Dimensions = v
	return s
}

func (s *GetBillingTrendRequestFilter) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBillingTrendRequestFilterDimensions struct {
	// The filter dimension code. For more information, see the "Additional information" section below.
	//
	// example:
	//
	// CHARGE_TYPE
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The filter method.
	//
	// example:
	//
	// IN
	SelectType *string `json:"selectType,omitempty" xml:"selectType,omitempty"`
	// The filter value list.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetBillingTrendRequestFilterDimensions) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendRequestFilterDimensions) GoString() string {
	return s.String()
}

func (s *GetBillingTrendRequestFilterDimensions) GetCode() *string {
	return s.Code
}

func (s *GetBillingTrendRequestFilterDimensions) GetSelectType() *string {
	return s.SelectType
}

func (s *GetBillingTrendRequestFilterDimensions) GetValues() []*string {
	return s.Values
}

func (s *GetBillingTrendRequestFilterDimensions) SetCode(v string) *GetBillingTrendRequestFilterDimensions {
	s.Code = &v
	return s
}

func (s *GetBillingTrendRequestFilterDimensions) SetSelectType(v string) *GetBillingTrendRequestFilterDimensions {
	s.SelectType = &v
	return s
}

func (s *GetBillingTrendRequestFilterDimensions) SetValues(v []*string) *GetBillingTrendRequestFilterDimensions {
	s.Values = v
	return s
}

func (s *GetBillingTrendRequestFilterDimensions) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendRequestGroupBy struct {
	// The grouping dimension code. For more information, see the "Additional information" section below.
	//
	// example:
	//
	// BASE_MODEL
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetBillingTrendRequestGroupBy) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendRequestGroupBy) GoString() string {
	return s.String()
}

func (s *GetBillingTrendRequestGroupBy) GetCode() *string {
	return s.Code
}

func (s *GetBillingTrendRequestGroupBy) SetCode(v string) *GetBillingTrendRequestGroupBy {
	s.Code = &v
	return s
}

func (s *GetBillingTrendRequestGroupBy) Validate() error {
	return dara.Validate(s)
}

type GetBillingTrendRequestTimePeriod struct {
	// The end time.
	//
	// example:
	//
	// 2026-08-25
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-08-01
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetBillingTrendRequestTimePeriod) String() string {
	return dara.Prettify(s)
}

func (s GetBillingTrendRequestTimePeriod) GoString() string {
	return s.String()
}

func (s *GetBillingTrendRequestTimePeriod) GetEnd() *string {
	return s.End
}

func (s *GetBillingTrendRequestTimePeriod) GetStart() *string {
	return s.Start
}

func (s *GetBillingTrendRequestTimePeriod) SetEnd(v string) *GetBillingTrendRequestTimePeriod {
	s.End = &v
	return s
}

func (s *GetBillingTrendRequestTimePeriod) SetStart(v string) *GetBillingTrendRequestTimePeriod {
	s.Start = &v
	return s
}

func (s *GetBillingTrendRequestTimePeriod) Validate() error {
	return dara.Validate(s)
}
