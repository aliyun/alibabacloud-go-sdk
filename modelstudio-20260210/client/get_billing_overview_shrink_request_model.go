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
	SetGroupByShrink(v string) *GetBillingOverviewShrinkRequest
	GetGroupByShrink() *string
	SetLocale(v string) *GetBillingOverviewShrinkRequest
	GetLocale() *string
	SetRegionId(v string) *GetBillingOverviewShrinkRequest
	GetRegionId() *string
	SetTopNum(v int32) *GetBillingOverviewShrinkRequest
	GetTopNum() *int32
	SetZeroFilter(v bool) *GetBillingOverviewShrinkRequest
	GetZeroFilter() *bool
}

type GetBillingOverviewShrinkRequest struct {
	// example:
	//
	// 2026-08
	BillMonth     *string `json:"billMonth,omitempty" xml:"billMonth,omitempty"`
	FilterShrink  *string `json:"filter,omitempty" xml:"filter,omitempty"`
	GroupByShrink *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// example:
	//
	// zh-CN
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 20
	TopNum *int32 `json:"topNum,omitempty" xml:"topNum,omitempty"`
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

func (s *GetBillingOverviewShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *GetBillingOverviewShrinkRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetBillingOverviewShrinkRequest) GetRegionId() *string {
	return s.RegionId
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
