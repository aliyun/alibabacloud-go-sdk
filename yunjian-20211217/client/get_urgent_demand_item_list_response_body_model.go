// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandItemListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetUrgentDemandItemListResponseBody
	GetCode() *int64
	SetData(v *GetUrgentDemandItemListResponseBodyData) *GetUrgentDemandItemListResponseBody
	GetData() *GetUrgentDemandItemListResponseBodyData
	SetMessage(v string) *GetUrgentDemandItemListResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetUrgentDemandItemListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *GetUrgentDemandItemListResponseBody
	GetTraceId() *string
}

type GetUrgentDemandItemListResponseBody struct {
	Code    *int64                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetUrgentDemandItemListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetUrgentDemandItemListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBody) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetUrgentDemandItemListResponseBody) GetData() *GetUrgentDemandItemListResponseBodyData {
	return s.Data
}

func (s *GetUrgentDemandItemListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUrgentDemandItemListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUrgentDemandItemListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetUrgentDemandItemListResponseBody) SetCode(v int64) *GetUrgentDemandItemListResponseBody {
	s.Code = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetData(v *GetUrgentDemandItemListResponseBodyData) *GetUrgentDemandItemListResponseBody {
	s.Data = v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetMessage(v string) *GetUrgentDemandItemListResponseBody {
	s.Message = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetSuccess(v bool) *GetUrgentDemandItemListResponseBody {
	s.Success = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) SetTraceId(v string) *GetUrgentDemandItemListResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandItemListResponseBodyData struct {
	Current *int64                                            `json:"current,omitempty" xml:"current,omitempty"`
	Pages   *int64                                            `json:"pages,omitempty" xml:"pages,omitempty"`
	Records []*GetUrgentDemandItemListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	Size    *int64                                            `json:"size,omitempty" xml:"size,omitempty"`
	Total   *int64                                            `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetUrgentDemandItemListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyData) GetCurrent() *int64 {
	return s.Current
}

func (s *GetUrgentDemandItemListResponseBodyData) GetPages() *int64 {
	return s.Pages
}

func (s *GetUrgentDemandItemListResponseBodyData) GetRecords() []*GetUrgentDemandItemListResponseBodyDataRecords {
	return s.Records
}

func (s *GetUrgentDemandItemListResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetUrgentDemandItemListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetUrgentDemandItemListResponseBodyData) SetCurrent(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Current = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetPages(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Pages = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetRecords(v []*GetUrgentDemandItemListResponseBodyDataRecords) *GetUrgentDemandItemListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetSize(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) SetTotal(v int64) *GetUrgentDemandItemListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandItemListResponseBodyDataRecords struct {
	Zone                   *string                                                               `json:"Zone,omitempty" xml:"Zone,omitempty"`
	AccountId              *string                                                               `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CommodityCode          *string                                                               `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityNum           *int64                                                                `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	CommodityTypeCode      *string                                                               `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	Creator                *string                                                               `json:"creator,omitempty" xml:"creator,omitempty"`
	CreatorName            *string                                                               `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	EffectTime             *string                                                               `json:"effectTime,omitempty" xml:"effectTime,omitempty"`
	GmtModified            *string                                                               `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                     *int64                                                                `json:"id,omitempty" xml:"id,omitempty"`
	Modifier               *string                                                               `json:"modifier,omitempty" xml:"modifier,omitempty"`
	ModifierName           *string                                                               `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	NetworkType            *string                                                               `json:"networkType,omitempty" xml:"networkType,omitempty"`
	PayDuration            *int64                                                                `json:"payDuration,omitempty" xml:"payDuration,omitempty"`
	PayDurationUnit        *string                                                               `json:"payDurationUnit,omitempty" xml:"payDurationUnit,omitempty"`
	PayType                *string                                                               `json:"payType,omitempty" xml:"payType,omitempty"`
	PlanId                 *int64                                                                `json:"planId,omitempty" xml:"planId,omitempty"`
	Region                 *string                                                               `json:"region,omitempty" xml:"region,omitempty"`
	UrgentDemandEbsRequest *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest `json:"urgentDemandEbsRequest,omitempty" xml:"urgentDemandEbsRequest,omitempty" type:"Struct"`
	UrgentDemandEcsRequest *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest `json:"urgentDemandEcsRequest,omitempty" xml:"urgentDemandEcsRequest,omitempty" type:"Struct"`
}

func (s GetUrgentDemandItemListResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetZone() *string {
	return s.Zone
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetCommodityNum() *int64 {
	return s.CommodityNum
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetEffectTime() *string {
	return s.EffectTime
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetId() *int64 {
	return s.Id
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetModifier() *string {
	return s.Modifier
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetModifierName() *string {
	return s.ModifierName
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetPayDuration() *int64 {
	return s.PayDuration
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetPayDurationUnit() *string {
	return s.PayDurationUnit
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetPayType() *string {
	return s.PayType
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetPlanId() *int64 {
	return s.PlanId
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetRegion() *string {
	return s.Region
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetUrgentDemandEbsRequest() *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	return s.UrgentDemandEbsRequest
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) GetUrgentDemandEcsRequest() *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	return s.UrgentDemandEcsRequest
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetZone(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Zone = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetAccountId(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.AccountId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCommodityCode(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCommodityNum(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CommodityNum = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCommodityTypeCode(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCreator(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetCreatorName(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.CreatorName = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetEffectTime(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.EffectTime = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetGmtModified(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetId(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetModifier(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Modifier = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetModifierName(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.ModifierName = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetNetworkType(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.NetworkType = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPayDuration(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PayDuration = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPayDurationUnit(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PayDurationUnit = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPayType(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PayType = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetPlanId(v int64) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.PlanId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetRegion(v string) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.Region = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetUrgentDemandEbsRequest(v *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.UrgentDemandEbsRequest = v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) SetUrgentDemandEcsRequest(v *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) *GetUrgentDemandItemListResponseBodyDataRecords {
	s.UrgentDemandEcsRequest = v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest struct {
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
	// 1
	DataDiskSize *int64 `json:"dataDiskSize,omitempty" xml:"dataDiskSize,omitempty"`
	// example:
	//
	// 111222
	ItemId *int64 `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// example:
	//
	// 1
	PerformanceLevel *int64 `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GetCommodityNum() *int64 {
	return s.CommodityNum
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GetDataDiskSize() *int64 {
	return s.DataDiskSize
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) GetPerformanceLevel() *int64 {
	return s.PerformanceLevel
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetCommodityCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetCommodityNum(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.CommodityNum = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetCommodityTypeCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetDataDiskSize(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.DataDiskSize = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetItemId(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.ItemId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) SetPerformanceLevel(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEbsRequest) Validate() error {
	return dara.Validate(s)
}

type GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest struct {
	CommodityCode     *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	CommodityNum      *int64  `json:"commodityNum,omitempty" xml:"commodityNum,omitempty"`
	CommodityTypeCode *string `json:"commodityTypeCode,omitempty" xml:"commodityTypeCode,omitempty"`
	ItemId            *int64  `json:"itemId,omitempty" xml:"itemId,omitempty"`
	VcpuCount         *int64  `json:"vcpuCount,omitempty" xml:"vcpuCount,omitempty"`
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GetCommodityNum() *int64 {
	return s.CommodityNum
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GetCommodityTypeCode() *string {
	return s.CommodityTypeCode
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) GetVcpuCount() *int64 {
	return s.VcpuCount
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetCommodityCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetCommodityNum(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.CommodityNum = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetCommodityTypeCode(v string) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.CommodityTypeCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetItemId(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.ItemId = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) SetVcpuCount(v int64) *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest {
	s.VcpuCount = &v
	return s
}

func (s *GetUrgentDemandItemListResponseBodyDataRecordsUrgentDemandEcsRequest) Validate() error {
	return dara.Validate(s)
}
