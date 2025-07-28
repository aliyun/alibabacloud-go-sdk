// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandItemListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *GetUrgentDemandItemListRequest
	GetCommodityCode() *string
	SetCommodityTypeCode(v string) *GetUrgentDemandItemListRequest
	GetCommodityTypeCode() *string
	SetCurrent(v int64) *GetUrgentDemandItemListRequest
	GetCurrent() *int64
	SetPlanId(v int64) *GetUrgentDemandItemListRequest
	GetPlanId() *int64
	SetRegion(v string) *GetUrgentDemandItemListRequest
	GetRegion() *string
	SetSize(v int64) *GetUrgentDemandItemListRequest
	GetSize() *int64
	SetZone(v string) *GetUrgentDemandItemListRequest
	GetZone() *string
}

type GetUrgentDemandItemListRequest struct {
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// ecs/yundisk
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 1
	Current *int64  `json:"current,omitempty" xml:"current,omitempty"`
	PlanId  *int64  `json:"planId,omitempty" xml:"planId,omitempty"`
	Region  *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 10
	Size *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s GetUrgentDemandItemListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetUrgentDemandItemListRequest) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *GetUrgentDemandItemListRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *GetUrgentDemandItemListRequest) GetPlanId() *int64 {
	return s.PlanId
}

func (s *GetUrgentDemandItemListRequest) GetRegion() *string {
	return s.Region
}

func (s *GetUrgentDemandItemListRequest) GetSize() *int64 {
	return s.Size
}

func (s *GetUrgentDemandItemListRequest) GetZone() *string {
	return s.Zone
}

func (s *GetUrgentDemandItemListRequest) SetCommodityCode(v string) *GetUrgentDemandItemListRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetCommodityTypeCode(v string) *GetUrgentDemandItemListRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetCurrent(v int64) *GetUrgentDemandItemListRequest {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetPlanId(v int64) *GetUrgentDemandItemListRequest {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetRegion(v string) *GetUrgentDemandItemListRequest {
	s.Region = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetSize(v int64) *GetUrgentDemandItemListRequest {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) SetZone(v string) *GetUrgentDemandItemListRequest {
	s.Zone = &v
	return s
}

func (s *GetUrgentDemandItemListRequest) Validate() error {
	return dara.Validate(s)
}
