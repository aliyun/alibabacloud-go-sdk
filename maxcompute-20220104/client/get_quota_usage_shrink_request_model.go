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
	AggMethod *string `json:"aggMethod,omitempty" xml:"aggMethod,omitempty"`
	// This parameter is required.
	From             *int64  `json:"from,omitempty" xml:"from,omitempty"`
	PlotTypesShrink  *string `json:"plotTypes,omitempty" xml:"plotTypes,omitempty"`
	ProductId        *string `json:"productId,omitempty" xml:"productId,omitempty"`
	Region           *string `json:"region,omitempty" xml:"region,omitempty"`
	SubQuotaNickname *string `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	TenantId         *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	To               *int64  `json:"to,omitempty" xml:"to,omitempty"`
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
