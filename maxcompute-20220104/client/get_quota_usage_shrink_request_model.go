// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggMethod(v string) *GetQuotaUsageShrinkRequest
	GetAggMethod() *string
	SetFrom(v int64) *GetQuotaUsageShrinkRequest
	GetFrom() *int64
	SetPlotTypesShrink(v string) *GetQuotaUsageShrinkRequest
	GetPlotTypesShrink() *string
	SetProductId(v string) *GetQuotaUsageShrinkRequest
	GetProductId() *string
	SetRegion(v string) *GetQuotaUsageShrinkRequest
	GetRegion() *string
	SetSubQuotaNickname(v string) *GetQuotaUsageShrinkRequest
	GetSubQuotaNickname() *string
	SetTenantId(v string) *GetQuotaUsageShrinkRequest
	GetTenantId() *string
	SetTo(v int64) *GetQuotaUsageShrinkRequest
	GetTo() *int64
	SetYAxisTypesShrink(v string) *GetQuotaUsageShrinkRequest
	GetYAxisTypesShrink() *string
}

type GetQuotaUsageShrinkRequest struct {
	// The aggregation algorithm. For a better page experience, up to 60 points can be displayed for each metric. If you select a time range longer than 1 hour, the chart uses the average value within the range (minutes of the selected time range/60) to aggregate data by default. You can change the aggregation algorithm based on your business requirements.
	//
	// example:
	//
	// max
	AggMethod *string `json:"aggMethod,omitempty" xml:"aggMethod,omitempty"`
	// The time when the query starts. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you set the **from*	- and **to*	- parameters to the same value, the time range is invalid and an error message is returned.
	//
	// 	- This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669081045
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The types of the charts.
	PlotTypesShrink *string `json:"plotTypes,omitempty" xml:"plotTypes,omitempty"`
	// The quota type. Default value: ODPS.
	//
	// 	- ODPS: computing quota
	//
	// 	- TUNNEL: Tunnel quota
	//
	// example:
	//
	// ODPS
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The alias of the level-2 quota.
	//
	// example:
	//
	// ot_tunnel_quota
	SubQuotaNickname *string `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	// The ID of the tenant. You can log on to the MaxCompute console, and choose Tenants > Tenant Property from the left-side navigation pane to view the tenant ID.
	//
	// example:
	//
	// 478403690625249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The time when the query ends. The value is the log time that is specified when log data is written.
	//
	// 	- The time range that is specified in this operation is a left-closed, right-open interval. The interval includes the start time specified by the **from*	- parameter, but does not include the end time specified by the **to*	- parameter. If you set the **from*	- and **to*	- parameters to the same value, the time range is invalid and an error message is returned.
	//
	// 	- This value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1669360870
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// The data metric fields.
	YAxisTypesShrink *string `json:"yAxisTypes,omitempty" xml:"yAxisTypes,omitempty"`
}

func (s GetQuotaUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageShrinkRequest) GetAggMethod() *string {
	return s.AggMethod
}

func (s *GetQuotaUsageShrinkRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetQuotaUsageShrinkRequest) GetPlotTypesShrink() *string {
	return s.PlotTypesShrink
}

func (s *GetQuotaUsageShrinkRequest) GetProductId() *string {
	return s.ProductId
}

func (s *GetQuotaUsageShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQuotaUsageShrinkRequest) GetSubQuotaNickname() *string {
	return s.SubQuotaNickname
}

func (s *GetQuotaUsageShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaUsageShrinkRequest) GetTo() *int64 {
	return s.To
}

func (s *GetQuotaUsageShrinkRequest) GetYAxisTypesShrink() *string {
	return s.YAxisTypesShrink
}

func (s *GetQuotaUsageShrinkRequest) SetAggMethod(v string) *GetQuotaUsageShrinkRequest {
	s.AggMethod = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetFrom(v int64) *GetQuotaUsageShrinkRequest {
	s.From = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetPlotTypesShrink(v string) *GetQuotaUsageShrinkRequest {
	s.PlotTypesShrink = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetProductId(v string) *GetQuotaUsageShrinkRequest {
	s.ProductId = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetRegion(v string) *GetQuotaUsageShrinkRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetSubQuotaNickname(v string) *GetQuotaUsageShrinkRequest {
	s.SubQuotaNickname = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetTenantId(v string) *GetQuotaUsageShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetTo(v int64) *GetQuotaUsageShrinkRequest {
	s.To = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) SetYAxisTypesShrink(v string) *GetQuotaUsageShrinkRequest {
	s.YAxisTypesShrink = &v
	return s
}

func (s *GetQuotaUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
