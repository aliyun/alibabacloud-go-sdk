// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveUrgentDemandItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *SaveUrgentDemandItemRequest
	GetAccountId() *string
	SetCreator(v string) *SaveUrgentDemandItemRequest
	GetCreator() *string
	SetEffectTime(v string) *SaveUrgentDemandItemRequest
	GetEffectTime() *string
	SetModifier(v string) *SaveUrgentDemandItemRequest
	GetModifier() *string
	SetNetworkType(v string) *SaveUrgentDemandItemRequest
	GetNetworkType() *string
	SetPayDuration(v string) *SaveUrgentDemandItemRequest
	GetPayDuration() *string
	SetPayDurationUnit(v string) *SaveUrgentDemandItemRequest
	GetPayDurationUnit() *string
	SetPayType(v string) *SaveUrgentDemandItemRequest
	GetPayType() *string
	SetPlanId(v int64) *SaveUrgentDemandItemRequest
	GetPlanId() *int64
	SetRegion(v string) *SaveUrgentDemandItemRequest
	GetRegion() *string
	SetUrgentDemandEbsRequest(v *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) *SaveUrgentDemandItemRequest
	GetUrgentDemandEbsRequest() *SaveUrgentDemandItemRequestUrgentDemandEbsRequest
	SetUrgentDemandEcsRequest(v *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) *SaveUrgentDemandItemRequest
	GetUrgentDemandEcsRequest() *SaveUrgentDemandItemRequestUrgentDemandEcsRequest
	SetZone(v string) *SaveUrgentDemandItemRequest
	GetZone() *string
}

type SaveUrgentDemandItemRequest struct {
	// example:
	//
	// 12321312
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 111222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 2021-12-27 00:00:00
	EffectTime *string `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	// example:
	//
	// 111222
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 网络类型 vpc（私有网络）/classic（经典网络）
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// example:
	//
	// 10
	PayDuration *string `json:"payDuration,omitempty" xml:"payDuration,omitempty"`
	// example:
	//
	// 购买时长单位(month(月)，week(周)，day(天))
	PayDurationUnit *string `json:"payDurationUnit,omitempty" xml:"payDurationUnit,omitempty"`
	// example:
	//
	// 付费类型 prepay(预付费)/postpay（后付费）
	PayType *string `json:"payType,omitempty" xml:"payType,omitempty"`
	// example:
	//
	// 111222
	PlanId *int64 `json:"planId,omitempty" xml:"planId,omitempty"`
	// example:
	//
	// cn-beijing
	Region                 *string                                            `json:"region,omitempty" xml:"region,omitempty"`
	UrgentDemandEbsRequest *SaveUrgentDemandItemRequestUrgentDemandEbsRequest `json:"urgentDemandEbsRequest,omitempty" xml:"urgentDemandEbsRequest,omitempty" type:"Struct"`
	UrgentDemandEcsRequest *SaveUrgentDemandItemRequestUrgentDemandEcsRequest `json:"urgentDemandEcsRequest,omitempty" xml:"urgentDemandEcsRequest,omitempty" type:"Struct"`
	// example:
	//
	// cn-beijing-a
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s SaveUrgentDemandItemRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveUrgentDemandItemRequest) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *SaveUrgentDemandItemRequest) GetCreator() *string {
	return s.Creator
}

func (s *SaveUrgentDemandItemRequest) GetEffectTime() *string {
	return s.EffectTime
}

func (s *SaveUrgentDemandItemRequest) GetModifier() *string {
	return s.Modifier
}

func (s *SaveUrgentDemandItemRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SaveUrgentDemandItemRequest) GetPayDuration() *string {
	return s.PayDuration
}

func (s *SaveUrgentDemandItemRequest) GetPayDurationUnit() *string {
	return s.PayDurationUnit
}

func (s *SaveUrgentDemandItemRequest) GetPayType() *string {
	return s.PayType
}

func (s *SaveUrgentDemandItemRequest) GetPlanId() *int64 {
	return s.PlanId
}

func (s *SaveUrgentDemandItemRequest) GetRegion() *string {
	return s.Region
}

func (s *SaveUrgentDemandItemRequest) GetUrgentDemandEbsRequest() *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	return s.UrgentDemandEbsRequest
}

func (s *SaveUrgentDemandItemRequest) GetUrgentDemandEcsRequest() *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	return s.UrgentDemandEcsRequest
}

func (s *SaveUrgentDemandItemRequest) GetZone() *string {
	return s.Zone
}

func (s *SaveUrgentDemandItemRequest) SetAccountId(v string) *SaveUrgentDemandItemRequest {
	s.AccountId = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetCreator(v string) *SaveUrgentDemandItemRequest {
	s.Creator = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetEffectTime(v string) *SaveUrgentDemandItemRequest {
	s.EffectTime = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetModifier(v string) *SaveUrgentDemandItemRequest {
	s.Modifier = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetNetworkType(v string) *SaveUrgentDemandItemRequest {
	s.NetworkType = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPayDuration(v string) *SaveUrgentDemandItemRequest {
	s.PayDuration = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPayDurationUnit(v string) *SaveUrgentDemandItemRequest {
	s.PayDurationUnit = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPayType(v string) *SaveUrgentDemandItemRequest {
	s.PayType = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetPlanId(v int64) *SaveUrgentDemandItemRequest {
	s.PlanId = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetRegion(v string) *SaveUrgentDemandItemRequest {
	s.Region = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetUrgentDemandEbsRequest(v *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) *SaveUrgentDemandItemRequest {
	s.UrgentDemandEbsRequest = v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetUrgentDemandEcsRequest(v *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) *SaveUrgentDemandItemRequest {
	s.UrgentDemandEcsRequest = v
	return s
}

func (s *SaveUrgentDemandItemRequest) SetZone(v string) *SaveUrgentDemandItemRequest {
	s.Zone = &v
	return s
}

func (s *SaveUrgentDemandItemRequest) Validate() error {
	return dara.Validate(s)
}

type SaveUrgentDemandItemRequestUrgentDemandEbsRequest struct {
	// example:
	//
	// cloud_essd
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 1
	CommodityNum *int64 `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	// example:
	//
	// yundisk
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 111222
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 1
	PerformanceLevel *int64 `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s SaveUrgentDemandItemRequestUrgentDemandEbsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GetCommodityNum() *int64 {
	return s.CommodityNum
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) GetPerformanceLevel() *int64 {
	return s.PerformanceLevel
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetCommodityCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.CommodityCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetCommodityNum(v int64) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.CommodityNum = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetCommodityTypeCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetItemId(v int64) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.ItemId = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) SetPerformanceLevel(v int64) *SaveUrgentDemandItemRequestUrgentDemandEbsRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEbsRequest) Validate() error {
	return dara.Validate(s)
}

type SaveUrgentDemandItemRequestUrgentDemandEcsRequest struct {
	// example:
	//
	// ecs.sn2ne.6xlarge
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// example:
	//
	// 2
	CommodityNum *int64 `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	// example:
	//
	// ecs
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	// example:
	//
	// 111222
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 2
	VCpuCount *int64 `json:"vCpuCount,omitempty" xml:"vCpuCount,omitempty"`
}

func (s SaveUrgentDemandItemRequestUrgentDemandEcsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GoString() string {
	return s.String()
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GetCommodityNum() *int64 {
	return s.CommodityNum
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) GetVCpuCount() *int64 {
	return s.VCpuCount
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetCommodityCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.CommodityCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetCommodityNum(v int64) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.CommodityNum = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetCommodityTypeCode(v string) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetItemId(v int64) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.ItemId = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) SetVCpuCount(v int64) *SaveUrgentDemandItemRequestUrgentDemandEcsRequest {
	s.VCpuCount = &v
	return s
}

func (s *SaveUrgentDemandItemRequestUrgentDemandEcsRequest) Validate() error {
	return dara.Validate(s)
}
