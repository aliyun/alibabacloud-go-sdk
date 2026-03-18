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
	AggMethod *string `json:"aggMethod,omitempty" xml:"aggMethod,omitempty"`
	// This parameter is required.
	From             *int64    `json:"from,omitempty" xml:"from,omitempty"`
	PlotTypes        []*string `json:"plotTypes,omitempty" xml:"plotTypes,omitempty" type:"Repeated"`
	ProductId        *string   `json:"productId,omitempty" xml:"productId,omitempty"`
	Region           *string   `json:"region,omitempty" xml:"region,omitempty"`
	SubQuotaNickname *string   `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	TenantId         *string   `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// This parameter is required.
	To         *int64    `json:"to,omitempty" xml:"to,omitempty"`
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
