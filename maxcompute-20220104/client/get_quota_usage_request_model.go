// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggMethod(v string) *GetQuotaUsageRequest
	GetAggMethod() *string
	SetFrom(v int64) *GetQuotaUsageRequest
	GetFrom() *int64
	SetPlotTypes(v []*string) *GetQuotaUsageRequest
	GetPlotTypes() []*string
	SetProductId(v string) *GetQuotaUsageRequest
	GetProductId() *string
	SetRegion(v string) *GetQuotaUsageRequest
	GetRegion() *string
	SetSubQuotaNickname(v string) *GetQuotaUsageRequest
	GetSubQuotaNickname() *string
	SetTenantId(v string) *GetQuotaUsageRequest
	GetTenantId() *string
	SetTo(v int64) *GetQuotaUsageRequest
	GetTo() *int64
	SetYAxisTypes(v []*string) *GetQuotaUsageRequest
	GetYAxisTypes() []*string
}

type GetQuotaUsageRequest struct {
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
	PlotTypes []*string `json:"plotTypes,omitempty" xml:"plotTypes,omitempty" type:"Repeated"`
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
	YAxisTypes []*string `json:"yAxisTypes,omitempty" xml:"yAxisTypes,omitempty" type:"Repeated"`
}

func (s GetQuotaUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUsageRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageRequest) GetAggMethod() *string {
	return s.AggMethod
}

func (s *GetQuotaUsageRequest) GetFrom() *int64 {
	return s.From
}

func (s *GetQuotaUsageRequest) GetPlotTypes() []*string {
	return s.PlotTypes
}

func (s *GetQuotaUsageRequest) GetProductId() *string {
	return s.ProductId
}

func (s *GetQuotaUsageRequest) GetRegion() *string {
	return s.Region
}

func (s *GetQuotaUsageRequest) GetSubQuotaNickname() *string {
	return s.SubQuotaNickname
}

func (s *GetQuotaUsageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetQuotaUsageRequest) GetTo() *int64 {
	return s.To
}

func (s *GetQuotaUsageRequest) GetYAxisTypes() []*string {
	return s.YAxisTypes
}

func (s *GetQuotaUsageRequest) SetAggMethod(v string) *GetQuotaUsageRequest {
	s.AggMethod = &v
	return s
}

func (s *GetQuotaUsageRequest) SetFrom(v int64) *GetQuotaUsageRequest {
	s.From = &v
	return s
}

func (s *GetQuotaUsageRequest) SetPlotTypes(v []*string) *GetQuotaUsageRequest {
	s.PlotTypes = v
	return s
}

func (s *GetQuotaUsageRequest) SetProductId(v string) *GetQuotaUsageRequest {
	s.ProductId = &v
	return s
}

func (s *GetQuotaUsageRequest) SetRegion(v string) *GetQuotaUsageRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaUsageRequest) SetSubQuotaNickname(v string) *GetQuotaUsageRequest {
	s.SubQuotaNickname = &v
	return s
}

func (s *GetQuotaUsageRequest) SetTenantId(v string) *GetQuotaUsageRequest {
	s.TenantId = &v
	return s
}

func (s *GetQuotaUsageRequest) SetTo(v int64) *GetQuotaUsageRequest {
	s.To = &v
	return s
}

func (s *GetQuotaUsageRequest) SetYAxisTypes(v []*string) *GetQuotaUsageRequest {
	s.YAxisTypes = v
	return s
}

func (s *GetQuotaUsageRequest) Validate() error {
	return dara.Validate(s)
}
