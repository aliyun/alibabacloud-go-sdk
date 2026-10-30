// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBillingOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillMonth(v string) *GetBillingOverviewRequest
	GetBillMonth() *string
	SetFilter(v *GetBillingOverviewRequestFilter) *GetBillingOverviewRequest
	GetFilter() *GetBillingOverviewRequestFilter
	SetGranularity(v string) *GetBillingOverviewRequest
	GetGranularity() *string
	SetGroupBy(v []*GetBillingOverviewRequestGroupBy) *GetBillingOverviewRequest
	GetGroupBy() []*GetBillingOverviewRequestGroupBy
	SetLocale(v string) *GetBillingOverviewRequest
	GetLocale() *string
	SetRegionId(v string) *GetBillingOverviewRequest
	GetRegionId() *string
	SetTimePeriod(v *GetBillingOverviewRequestTimePeriod) *GetBillingOverviewRequest
	GetTimePeriod() *GetBillingOverviewRequestTimePeriod
	SetTopNum(v int32) *GetBillingOverviewRequest
	GetTopNum() *int32
	SetZeroFilter(v bool) *GetBillingOverviewRequest
	GetZeroFilter() *bool
}

type GetBillingOverviewRequest struct {
	// The billing month.
	//
	// example:
	//
	// 2026-08
	BillMonth *string `json:"billMonth,omitempty" xml:"billMonth,omitempty"`
	// The filter conditions.
	Filter *GetBillingOverviewRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The time granularity.
	//
	// example:
	//
	// MONTH
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// The list of grouping conditions. Currently, you must specify exactly one grouping dimension.
	GroupBy []*GetBillingOverviewRequestGroupBy `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
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
	TimePeriod *GetBillingOverviewRequestTimePeriod `json:"timePeriod,omitempty" xml:"timePeriod,omitempty" type:"Struct"`
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

func (s GetBillingOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewRequest) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *GetBillingOverviewRequest) GetFilter() *GetBillingOverviewRequestFilter {
	return s.Filter
}

func (s *GetBillingOverviewRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *GetBillingOverviewRequest) GetGroupBy() []*GetBillingOverviewRequestGroupBy {
	return s.GroupBy
}

func (s *GetBillingOverviewRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetBillingOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBillingOverviewRequest) GetTimePeriod() *GetBillingOverviewRequestTimePeriod {
	return s.TimePeriod
}

func (s *GetBillingOverviewRequest) GetTopNum() *int32 {
	return s.TopNum
}

func (s *GetBillingOverviewRequest) GetZeroFilter() *bool {
	return s.ZeroFilter
}

func (s *GetBillingOverviewRequest) SetBillMonth(v string) *GetBillingOverviewRequest {
	s.BillMonth = &v
	return s
}

func (s *GetBillingOverviewRequest) SetFilter(v *GetBillingOverviewRequestFilter) *GetBillingOverviewRequest {
	s.Filter = v
	return s
}

func (s *GetBillingOverviewRequest) SetGranularity(v string) *GetBillingOverviewRequest {
	s.Granularity = &v
	return s
}

func (s *GetBillingOverviewRequest) SetGroupBy(v []*GetBillingOverviewRequestGroupBy) *GetBillingOverviewRequest {
	s.GroupBy = v
	return s
}

func (s *GetBillingOverviewRequest) SetLocale(v string) *GetBillingOverviewRequest {
	s.Locale = &v
	return s
}

func (s *GetBillingOverviewRequest) SetRegionId(v string) *GetBillingOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *GetBillingOverviewRequest) SetTimePeriod(v *GetBillingOverviewRequestTimePeriod) *GetBillingOverviewRequest {
	s.TimePeriod = v
	return s
}

func (s *GetBillingOverviewRequest) SetTopNum(v int32) *GetBillingOverviewRequest {
	s.TopNum = &v
	return s
}

func (s *GetBillingOverviewRequest) SetZeroFilter(v bool) *GetBillingOverviewRequest {
	s.ZeroFilter = &v
	return s
}

func (s *GetBillingOverviewRequest) Validate() error {
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

type GetBillingOverviewRequestFilter struct {
	// The list of dimension filters.
	Dimensions []*GetBillingOverviewRequestFilterDimensions `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
}

func (s GetBillingOverviewRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewRequestFilter) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewRequestFilter) GetDimensions() []*GetBillingOverviewRequestFilterDimensions {
	return s.Dimensions
}

func (s *GetBillingOverviewRequestFilter) SetDimensions(v []*GetBillingOverviewRequestFilterDimensions) *GetBillingOverviewRequestFilter {
	s.Dimensions = v
	return s
}

func (s *GetBillingOverviewRequestFilter) Validate() error {
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

type GetBillingOverviewRequestFilterDimensions struct {
	// The filter field. For more information, see the "Additional notes" section below.
	//
	// example:
	//
	// CHARGE_TYPE
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The filter type.
	//
	// example:
	//
	// IN
	SelectType *string `json:"selectType,omitempty" xml:"selectType,omitempty"`
	// The list of filter values.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetBillingOverviewRequestFilterDimensions) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewRequestFilterDimensions) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewRequestFilterDimensions) GetCode() *string {
	return s.Code
}

func (s *GetBillingOverviewRequestFilterDimensions) GetSelectType() *string {
	return s.SelectType
}

func (s *GetBillingOverviewRequestFilterDimensions) GetValues() []*string {
	return s.Values
}

func (s *GetBillingOverviewRequestFilterDimensions) SetCode(v string) *GetBillingOverviewRequestFilterDimensions {
	s.Code = &v
	return s
}

func (s *GetBillingOverviewRequestFilterDimensions) SetSelectType(v string) *GetBillingOverviewRequestFilterDimensions {
	s.SelectType = &v
	return s
}

func (s *GetBillingOverviewRequestFilterDimensions) SetValues(v []*string) *GetBillingOverviewRequestFilterDimensions {
	s.Values = v
	return s
}

func (s *GetBillingOverviewRequestFilterDimensions) Validate() error {
	return dara.Validate(s)
}

type GetBillingOverviewRequestGroupBy struct {
	// The code of the grouping dimension. For more information, see the "Additional notes" section below.
	//
	// example:
	//
	// BASE_MODEL
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetBillingOverviewRequestGroupBy) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewRequestGroupBy) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewRequestGroupBy) GetCode() *string {
	return s.Code
}

func (s *GetBillingOverviewRequestGroupBy) SetCode(v string) *GetBillingOverviewRequestGroupBy {
	s.Code = &v
	return s
}

func (s *GetBillingOverviewRequestGroupBy) Validate() error {
	return dara.Validate(s)
}

type GetBillingOverviewRequestTimePeriod struct {
	// The end time of the query. This time point is included in the query range. The format is determined by granularity.
	//
	// example:
	//
	// 202609
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// The start time of the query. This time point is included in the query range. The format is determined by granularity.
	//
	// example:
	//
	// 202609
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetBillingOverviewRequestTimePeriod) String() string {
	return dara.Prettify(s)
}

func (s GetBillingOverviewRequestTimePeriod) GoString() string {
	return s.String()
}

func (s *GetBillingOverviewRequestTimePeriod) GetEnd() *string {
	return s.End
}

func (s *GetBillingOverviewRequestTimePeriod) GetStart() *string {
	return s.Start
}

func (s *GetBillingOverviewRequestTimePeriod) SetEnd(v string) *GetBillingOverviewRequestTimePeriod {
	s.End = &v
	return s
}

func (s *GetBillingOverviewRequestTimePeriod) SetStart(v string) *GetBillingOverviewRequestTimePeriod {
	s.Start = &v
	return s
}

func (s *GetBillingOverviewRequestTimePeriod) Validate() error {
	return dara.Validate(s)
}
