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
	SetGroupBy(v []*GetBillingOverviewRequestGroupBy) *GetBillingOverviewRequest
	GetGroupBy() []*GetBillingOverviewRequestGroupBy
	SetLocale(v string) *GetBillingOverviewRequest
	GetLocale() *string
	SetRegionId(v string) *GetBillingOverviewRequest
	GetRegionId() *string
	SetTopNum(v int32) *GetBillingOverviewRequest
	GetTopNum() *int32
	SetZeroFilter(v bool) *GetBillingOverviewRequest
	GetZeroFilter() *bool
}

type GetBillingOverviewRequest struct {
	// The billing month. This parameter is required.
	//
	// example:
	//
	// 2026-08
	BillMonth *string `json:"billMonth,omitempty" xml:"billMonth,omitempty"`
	// The filter condition.
	Filter *GetBillingOverviewRequestFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	// The list of grouping conditions. Currently, you must specify exactly one grouping dimension.
	GroupBy []*GetBillingOverviewRequestGroupBy `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
	// The response language. Default value: en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The number of groups to return. Valid values: 1 to 20. Default value: 20.
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

func (s *GetBillingOverviewRequest) GetGroupBy() []*GetBillingOverviewRequestGroupBy {
	return s.GroupBy
}

func (s *GetBillingOverviewRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetBillingOverviewRequest) GetRegionId() *string {
	return s.RegionId
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
	// The filter field. For more information, see the "Additional information" section below.
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
	// The grouping dimension code. For more information, see the "Additional information" section below.
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
