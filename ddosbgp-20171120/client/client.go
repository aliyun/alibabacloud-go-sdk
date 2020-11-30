// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeOnDemandInstanceRequest struct {
	PageNo   *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOnDemandInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceRequest) SetPageNo(v int) *DescribeOnDemandInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeOnDemandInstanceRequest) SetPageSize(v int) *DescribeOnDemandInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOnDemandInstanceRequest) SetRegionId(v string) *DescribeOnDemandInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeOnDemandInstanceResponse struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *string                                      `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Instances []*DescribeOnDemandInstanceResponseInstances `json:"Instances,omitempty" xml:"Instances,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOnDemandInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceResponse) SetRequestId(v string) *DescribeOnDemandInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOnDemandInstanceResponse) SetTotal(v string) *DescribeOnDemandInstanceResponse {
	s.Total = &v
	return s
}

func (s *DescribeOnDemandInstanceResponse) SetInstances(v []*DescribeOnDemandInstanceResponseInstances) *DescribeOnDemandInstanceResponse {
	s.Instances = v
	return s
}

type DescribeOnDemandInstanceResponseInstances struct {
	InstanceId    *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Remark        *string   `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
	DefenseStatus *string   `json:"DefenseStatus,omitempty" xml:"DefenseStatus,omitempty" require:"true"`
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Ipnet         []*string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOnDemandInstanceResponseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeOnDemandInstanceResponseInstances) GoString() string {
	return s.String()
}

func (s *DescribeOnDemandInstanceResponseInstances) SetInstanceId(v string) *DescribeOnDemandInstanceResponseInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseInstances) SetRemark(v string) *DescribeOnDemandInstanceResponseInstances {
	s.Remark = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseInstances) SetDefenseStatus(v string) *DescribeOnDemandInstanceResponseInstances {
	s.DefenseStatus = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseInstances) SetRegionId(v string) *DescribeOnDemandInstanceResponseInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeOnDemandInstanceResponseInstances) SetIpnet(v []*string) *DescribeOnDemandInstanceResponseInstances {
	s.Ipnet = v
	return s
}

type ModifyOnDemaondDefenseStatusRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	DefenseStatus *string `json:"DefenseStatus,omitempty" xml:"DefenseStatus,omitempty" require:"true"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOnDemaondDefenseStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnDemaondDefenseStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyOnDemaondDefenseStatusRequest) SetInstanceId(v string) *ModifyOnDemaondDefenseStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyOnDemaondDefenseStatusRequest) SetDefenseStatus(v string) *ModifyOnDemaondDefenseStatusRequest {
	s.DefenseStatus = &v
	return s
}

func (s *ModifyOnDemaondDefenseStatusRequest) SetRegionId(v string) *ModifyOnDemaondDefenseStatusRequest {
	s.RegionId = &v
	return s
}

type ModifyOnDemaondDefenseStatusResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyOnDemaondDefenseStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOnDemaondDefenseStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyOnDemaondDefenseStatusResponse) SetRequestId(v string) *ModifyOnDemaondDefenseStatusResponse {
	s.RequestId = &v
	return s
}

type DescribeOpEntitiesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeOpEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesRequest) SetSourceIp(v string) *DescribeOpEntitiesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetLang(v string) *DescribeOpEntitiesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetCurrentPage(v int) *DescribeOpEntitiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetPageSize(v int) *DescribeOpEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetStartTime(v int64) *DescribeOpEntitiesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOpEntitiesRequest) SetEndTime(v int64) *DescribeOpEntitiesRequest {
	s.EndTime = &v
	return s
}

type DescribeOpEntitiesResponse struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount *int                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	OpEntities []*DescribeOpEntitiesResponseOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeOpEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponse) SetRequestId(v string) *DescribeOpEntitiesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetTotalCount(v int) *DescribeOpEntitiesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeOpEntitiesResponse) SetOpEntities(v []*DescribeOpEntitiesResponseOpEntities) *DescribeOpEntitiesResponse {
	s.OpEntities = v
	return s
}

type DescribeOpEntitiesResponseOpEntities struct {
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty" require:"true"`
	EntityType   *int    `json:"EntityType,omitempty" xml:"EntityType,omitempty" require:"true"`
	OpDesc       *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty" require:"true"`
	OpAccount    *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty" require:"true"`
	OpAction     *int    `json:"OpAction,omitempty" xml:"OpAction,omitempty" require:"true"`
	GmtCreate    *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
}

func (s DescribeOpEntitiesResponseOpEntities) String() string {
	return tea.Prettify(s)
}

func (s DescribeOpEntitiesResponseOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetEntityType(v int) *DescribeOpEntitiesResponseOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseOpEntities {
	s.OpDesc = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetOpAction(v int) *DescribeOpEntitiesResponseOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseOpEntities {
	s.GmtCreate = &v
	return s
}

type DescribePackPaidTrafficRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribePackPaidTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficRequest) SetSourceIp(v string) *DescribePackPaidTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetInstanceId(v string) *DescribePackPaidTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetCurrentPage(v int) *DescribePackPaidTrafficRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetPageSize(v int) *DescribePackPaidTrafficRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetStartTime(v int64) *DescribePackPaidTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePackPaidTrafficRequest) SetEndTime(v int64) *DescribePackPaidTrafficRequest {
	s.EndTime = &v
	return s
}

type DescribePackPaidTrafficResponse struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount       *int                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PackPaidTraffics []*DescribePackPaidTrafficResponsePackPaidTraffics `json:"PackPaidTraffics,omitempty" xml:"PackPaidTraffics,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePackPaidTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponse) SetRequestId(v string) *DescribePackPaidTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetTotalCount(v int) *DescribePackPaidTrafficResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribePackPaidTrafficResponse) SetPackPaidTraffics(v []*DescribePackPaidTrafficResponsePackPaidTraffics) *DescribePackPaidTrafficResponse {
	s.PackPaidTraffics = v
	return s
}

type DescribePackPaidTrafficResponsePackPaidTraffics struct {
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	StartTime        *int64   `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	BaseBandwidth    *int     `json:"BaseBandwidth,omitempty" xml:"BaseBandwidth,omitempty" require:"true"`
	ElasticBandwidth *int     `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty" require:"true"`
	PaidCapacity     *float32 `json:"PaidCapacity,omitempty" xml:"PaidCapacity,omitempty" require:"true"`
	TotalCapacity    *float32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty" require:"true"`
	MaxAttack        *float32 `json:"MaxAttack,omitempty" xml:"MaxAttack,omitempty" require:"true"`
}

func (s DescribePackPaidTrafficResponsePackPaidTraffics) String() string {
	return tea.Prettify(s)
}

func (s DescribePackPaidTrafficResponsePackPaidTraffics) GoString() string {
	return s.String()
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetInstanceId(v string) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.InstanceId = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetStartTime(v int64) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.StartTime = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetBaseBandwidth(v int) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.BaseBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetElasticBandwidth(v int) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetPaidCapacity(v float32) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.PaidCapacity = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetTotalCapacity(v float32) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.TotalCapacity = &v
	return s
}

func (s *DescribePackPaidTrafficResponsePackPaidTraffics) SetMaxAttack(v float32) *DescribePackPaidTrafficResponsePackPaidTraffics {
	s.MaxAttack = &v
	return s
}

type DescribeResourcePackUsageRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
}

func (s DescribeResourcePackUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageRequest) SetSourceIp(v string) *DescribeResourcePackUsageRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetEndTime(v int64) *DescribeResourcePackUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageRequest) SetStartTime(v int64) *DescribeResourcePackUsageRequest {
	s.StartTime = &v
	return s
}

type DescribeResourcePackUsageResponse struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Interval   *int64                                         `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
	StartTime  *int64                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime    *int64                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PackUsages []*DescribeResourcePackUsageResponsePackUsages `json:"PackUsages,omitempty" xml:"PackUsages,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponse) SetRequestId(v string) *DescribeResourcePackUsageResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetInterval(v int64) *DescribeResourcePackUsageResponse {
	s.Interval = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetStartTime(v int64) *DescribeResourcePackUsageResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetEndTime(v int64) *DescribeResourcePackUsageResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackUsageResponse) SetPackUsages(v []*DescribeResourcePackUsageResponsePackUsages) *DescribeResourcePackUsageResponse {
	s.PackUsages = v
	return s
}

type DescribeResourcePackUsageResponsePackUsages struct {
	Traffic *float32 `json:"Traffic,omitempty" xml:"Traffic,omitempty" require:"true"`
	Time    *int64   `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeResourcePackUsageResponsePackUsages) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackUsageResponsePackUsages) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackUsageResponsePackUsages) SetTraffic(v float32) *DescribeResourcePackUsageResponsePackUsages {
	s.Traffic = &v
	return s
}

func (s *DescribeResourcePackUsageResponsePackUsages) SetTime(v int64) *DescribeResourcePackUsageResponsePackUsages {
	s.Time = &v
	return s
}

type DescribeResourcePackStatisticsRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeResourcePackStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsRequest) SetSourceIp(v string) *DescribeResourcePackStatisticsRequest {
	s.SourceIp = &v
	return s
}

type DescribeResourcePackStatisticsResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AvailablePackNum  *int    `json:"AvailablePackNum,omitempty" xml:"AvailablePackNum,omitempty" require:"true"`
	TotalCurrCapacity *int64  `json:"TotalCurrCapacity,omitempty" xml:"TotalCurrCapacity,omitempty" require:"true"`
	TotalUsedCapacity *int64  `json:"TotalUsedCapacity,omitempty" xml:"TotalUsedCapacity,omitempty" require:"true"`
	TotalInitCapacity *int64  `json:"TotalInitCapacity,omitempty" xml:"TotalInitCapacity,omitempty" require:"true"`
}

func (s DescribeResourcePackStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackStatisticsResponse) SetRequestId(v string) *DescribeResourcePackStatisticsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetAvailablePackNum(v int) *DescribeResourcePackStatisticsResponse {
	s.AvailablePackNum = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetTotalCurrCapacity(v int64) *DescribeResourcePackStatisticsResponse {
	s.TotalCurrCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetTotalUsedCapacity(v int64) *DescribeResourcePackStatisticsResponse {
	s.TotalUsedCapacity = &v
	return s
}

func (s *DescribeResourcePackStatisticsResponse) SetTotalInitCapacity(v int64) *DescribeResourcePackStatisticsResponse {
	s.TotalInitCapacity = &v
	return s
}

type DescribeResourcePackInstancesRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageSize    *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	CurrentPage *int    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty" require:"true"`
}

func (s DescribeResourcePackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesRequest) SetSourceIp(v string) *DescribeResourcePackInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetPageSize(v int) *DescribeResourcePackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourcePackInstancesRequest) SetCurrentPage(v int) *DescribeResourcePackInstancesRequest {
	s.CurrentPage = &v
	return s
}

type DescribeResourcePackInstancesResponse struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount    *int                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	ResourcePacks []*DescribeResourcePackInstancesResponseResourcePacks `json:"ResourcePacks,omitempty" xml:"ResourcePacks,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeResourcePackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponse) SetRequestId(v string) *DescribeResourcePackInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetTotalCount(v int) *DescribeResourcePackInstancesResponse {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourcePackInstancesResponse) SetResourcePacks(v []*DescribeResourcePackInstancesResponseResourcePacks) *DescribeResourcePackInstancesResponse {
	s.ResourcePacks = v
	return s
}

type DescribeResourcePackInstancesResponseResourcePacks struct {
	ResourcePackId *string `json:"ResourcePackId,omitempty" xml:"ResourcePackId,omitempty" require:"true"`
	InitCapacity   *int64  `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty" require:"true"`
	CurrCapacity   *int64  `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty" require:"true"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeResourcePackInstancesResponseResourcePacks) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcePackInstancesResponseResourcePacks) GoString() string {
	return s.String()
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetResourcePackId(v string) *DescribeResourcePackInstancesResponseResourcePacks {
	s.ResourcePackId = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetInitCapacity(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.InitCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetCurrCapacity(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetStartTime(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.StartTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetEndTime(v int64) *DescribeResourcePackInstancesResponseResourcePacks {
	s.EndTime = &v
	return s
}

func (s *DescribeResourcePackInstancesResponseResourcePacks) SetStatus(v string) *DescribeResourcePackInstancesResponseResourcePacks {
	s.Status = &v
	return s
}

type DeleteBlackholeRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
}

func (s DeleteBlackholeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeRequest) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeRequest) SetSourceIp(v string) *DeleteBlackholeRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteBlackholeRequest) SetPackId(v string) *DeleteBlackholeRequest {
	s.PackId = &v
	return s
}

func (s *DeleteBlackholeRequest) SetIp(v string) *DeleteBlackholeRequest {
	s.Ip = &v
	return s
}

type DeleteBlackholeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteBlackholeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBlackholeResponse) GoString() string {
	return s.String()
}

func (s *DeleteBlackholeResponse) SetRequestId(v string) *DeleteBlackholeResponse {
	s.RequestId = &v
	return s
}

type CheckGrantRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) SetSourceIp(v string) *CheckGrantRequest {
	s.SourceIp = &v
	return s
}

type CheckGrantResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Status    *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s CheckGrantResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckGrantResponse) GoString() string {
	return s.String()
}

func (s *CheckGrantResponse) SetRequestId(v string) *CheckGrantResponse {
	s.RequestId = &v
	return s
}

func (s *CheckGrantResponse) SetStatus(v int) *CheckGrantResponse {
	s.Status = &v
	return s
}

type DeleteProductRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
}

func (s DeleteProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) SetSourceIp(v string) *DeleteProductRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteProductRequest) SetPackId(v string) *DeleteProductRequest {
	s.PackId = &v
	return s
}

type DeleteProductResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetRequestId(v string) *DeleteProductResponse {
	s.RequestId = &v
	return s
}

type AddProductRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty" require:"true"`
}

func (s AddProductRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProductRequest) GoString() string {
	return s.String()
}

func (s *AddProductRequest) SetSourceIp(v string) *AddProductRequest {
	s.SourceIp = &v
	return s
}

func (s *AddProductRequest) SetPackId(v string) *AddProductRequest {
	s.PackId = &v
	return s
}

func (s *AddProductRequest) SetProduct(v string) *AddProductRequest {
	s.Product = &v
	return s
}

type AddProductResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddProductResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProductResponse) GoString() string {
	return s.String()
}

func (s *AddProductResponse) SetRequestId(v string) *AddProductResponse {
	s.RequestId = &v
	return s
}

type AddIpRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	IpList   *string `json:"IpList,omitempty" xml:"IpList,omitempty" require:"true"`
}

func (s AddIpRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpRequest) GoString() string {
	return s.String()
}

func (s *AddIpRequest) SetSourceIp(v string) *AddIpRequest {
	s.SourceIp = &v
	return s
}

func (s *AddIpRequest) SetPackId(v string) *AddIpRequest {
	s.PackId = &v
	return s
}

func (s *AddIpRequest) SetIpList(v string) *AddIpRequest {
	s.IpList = &v
	return s
}

type AddIpResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddIpResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpResponse) GoString() string {
	return s.String()
}

func (s *AddIpResponse) SetRequestId(v string) *AddIpResponse {
	s.RequestId = &v
	return s
}

type DescribeInstanceListRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackIdList     *string `json:"PackIdList,omitempty" xml:"PackIdList,omitempty"`
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	PageNo         *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
}

func (s DescribeInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListRequest) SetSourceIp(v string) *DescribeInstanceListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPackIdList(v string) *DescribeInstanceListRequest {
	s.PackIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetInstanceIdList(v string) *DescribeInstanceListRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageNo(v int) *DescribeInstanceListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstanceListRequest) SetPageSize(v int) *DescribeInstanceListRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceListResponse struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total        *int64                                      `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	InstanceList []*DescribeInstanceListResponseInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponse) SetRequestId(v string) *DescribeInstanceListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceListResponse) SetTotal(v int64) *DescribeInstanceListResponse {
	s.Total = &v
	return s
}

func (s *DescribeInstanceListResponse) SetInstanceList(v []*DescribeInstanceListResponseInstanceList) *DescribeInstanceListResponse {
	s.InstanceList = v
	return s
}

type DescribeInstanceListResponseInstanceList struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Product    *string `json:"Product,omitempty" xml:"Product,omitempty" require:"true"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty" require:"true"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	PackId     *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	GmtCreate  *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
}

func (s DescribeInstanceListResponseInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceListResponseInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceListResponseInstanceList) SetExpireTime(v int64) *DescribeInstanceListResponseInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetInstanceId(v string) *DescribeInstanceListResponseInstanceList {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetProduct(v string) *DescribeInstanceListResponseInstanceList {
	s.Product = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetRemark(v string) *DescribeInstanceListResponseInstanceList {
	s.Remark = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetStatus(v string) *DescribeInstanceListResponseInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetPackId(v string) *DescribeInstanceListResponseInstanceList {
	s.PackId = &v
	return s
}

func (s *DescribeInstanceListResponseInstanceList) SetGmtCreate(v int64) *DescribeInstanceListResponseInstanceList {
	s.GmtCreate = &v
	return s
}

type DescribeTopTrafficRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" require:"true"`
	Ipnet           *string `json:"Ipnet,omitempty" xml:"Ipnet,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Rn              *int    `json:"Rn,omitempty" xml:"Rn,omitempty"`
	PageNo          *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTopTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficRequest) SetInstanceId(v string) *DescribeTopTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetIpnet(v string) *DescribeTopTrafficRequest {
	s.Ipnet = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetStartTime(v string) *DescribeTopTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetEndTime(v string) *DescribeTopTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetRn(v int) *DescribeTopTrafficRequest {
	s.Rn = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetPageNo(v int) *DescribeTopTrafficRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetPageSize(v int) *DescribeTopTrafficRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetResourceGroupId(v string) *DescribeTopTrafficRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTopTrafficRequest) SetRegionId(v string) *DescribeTopTrafficRequest {
	s.RegionId = &v
	return s
}

type DescribeTopTrafficResponse struct {
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total       *int64                                   `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	TrafficList []*DescribeTopTrafficResponseTrafficList `json:"TrafficList,omitempty" xml:"TrafficList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeTopTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficResponse) SetRequestId(v string) *DescribeTopTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeTopTrafficResponse) SetTotal(v int64) *DescribeTopTrafficResponse {
	s.Total = &v
	return s
}

func (s *DescribeTopTrafficResponse) SetTrafficList(v []*DescribeTopTrafficResponseTrafficList) *DescribeTopTrafficResponse {
	s.TrafficList = v
	return s
}

type DescribeTopTrafficResponseTrafficList struct {
	Pps       *int    `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
	Bps       *int    `json:"Bps,omitempty" xml:"Bps,omitempty" require:"true"`
	AttackBps *int    `json:"AttackBps,omitempty" xml:"AttackBps,omitempty" require:"true"`
	AttackPps *int    `json:"AttackPps,omitempty" xml:"AttackPps,omitempty" require:"true"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
}

func (s DescribeTopTrafficResponseTrafficList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopTrafficResponseTrafficList) GoString() string {
	return s.String()
}

func (s *DescribeTopTrafficResponseTrafficList) SetPps(v int) *DescribeTopTrafficResponseTrafficList {
	s.Pps = &v
	return s
}

func (s *DescribeTopTrafficResponseTrafficList) SetBps(v int) *DescribeTopTrafficResponseTrafficList {
	s.Bps = &v
	return s
}

func (s *DescribeTopTrafficResponseTrafficList) SetAttackBps(v int) *DescribeTopTrafficResponseTrafficList {
	s.AttackBps = &v
	return s
}

func (s *DescribeTopTrafficResponseTrafficList) SetAttackPps(v int) *DescribeTopTrafficResponseTrafficList {
	s.AttackPps = &v
	return s
}

func (s *DescribeTopTrafficResponseTrafficList) SetIp(v string) *DescribeTopTrafficResponseTrafficList {
	s.Ip = &v
	return s
}

type DescribeDdosEventRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId    *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	StartTime *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageSize  *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	PageNo    *int    `json:"PageNo,omitempty" xml:"PageNo,omitempty" require:"true"`
}

func (s DescribeDdosEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventRequest) SetSourceIp(v string) *DescribeDdosEventRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPackId(v string) *DescribeDdosEventRequest {
	s.PackId = &v
	return s
}

func (s *DescribeDdosEventRequest) SetStartTime(v int) *DescribeDdosEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetEndTime(v int) *DescribeDdosEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageSize(v int) *DescribeDdosEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosEventRequest) SetPageNo(v int) *DescribeDdosEventRequest {
	s.PageNo = &v
	return s
}

type DescribeDdosEventResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Total     *int64                             `json:"Total,omitempty" xml:"Total,omitempty" require:"true"`
	Events    []*DescribeDdosEventResponseEvents `json:"Events,omitempty" xml:"Events,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDdosEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponse) SetRequestId(v string) *DescribeDdosEventResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventResponse) SetTotal(v int64) *DescribeDdosEventResponse {
	s.Total = &v
	return s
}

func (s *DescribeDdosEventResponse) SetEvents(v []*DescribeDdosEventResponseEvents) *DescribeDdosEventResponse {
	s.Events = v
	return s
}

type DescribeDdosEventResponseEvents struct {
	StartTime *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Pps       *int    `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
	Ip        *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	Mbps      *int    `json:"Mbps,omitempty" xml:"Mbps,omitempty" require:"true"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s DescribeDdosEventResponseEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeDdosEventResponseEvents) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventResponseEvents) SetStartTime(v int) *DescribeDdosEventResponseEvents {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetEndTime(v int) *DescribeDdosEventResponseEvents {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetPps(v int) *DescribeDdosEventResponseEvents {
	s.Pps = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetIp(v string) *DescribeDdosEventResponseEvents {
	s.Ip = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetMbps(v int) *DescribeDdosEventResponseEvents {
	s.Mbps = &v
	return s
}

func (s *DescribeDdosEventResponseEvents) SetStatus(v string) *DescribeDdosEventResponseEvents {
	s.Status = &v
	return s
}

type DescribeTrafficRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	StartTime *int    `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime   *int    `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Interval  *int    `json:"Interval,omitempty" xml:"Interval,omitempty" require:"true"`
}

func (s DescribeTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficRequest) SetSourceIp(v string) *DescribeTrafficRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTrafficRequest) SetName(v string) *DescribeTrafficRequest {
	s.Name = &v
	return s
}

func (s *DescribeTrafficRequest) SetStartTime(v int) *DescribeTrafficRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetEndTime(v int) *DescribeTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTrafficRequest) SetInterval(v int) *DescribeTrafficRequest {
	s.Interval = &v
	return s
}

type DescribeTrafficResponse struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	FlowList  []*DescribeTrafficResponseFlowList `json:"FlowList,omitempty" xml:"FlowList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponse) SetRequestId(v string) *DescribeTrafficResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficResponse) SetFlowList(v []*DescribeTrafficResponseFlowList) *DescribeTrafficResponse {
	s.FlowList = v
	return s
}

type DescribeTrafficResponseFlowList struct {
	Pps      *int    `json:"Pps,omitempty" xml:"Pps,omitempty" require:"true"`
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty" require:"true"`
	Kbps     *int    `json:"Kbps,omitempty" xml:"Kbps,omitempty" require:"true"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Time     *int    `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
}

func (s DescribeTrafficResponseFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficResponseFlowList) GoString() string {
	return s.String()
}

func (s *DescribeTrafficResponseFlowList) SetPps(v int) *DescribeTrafficResponseFlowList {
	s.Pps = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetFlowType(v string) *DescribeTrafficResponseFlowList {
	s.FlowType = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetKbps(v int) *DescribeTrafficResponseFlowList {
	s.Kbps = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetName(v string) *DescribeTrafficResponseFlowList {
	s.Name = &v
	return s
}

func (s *DescribeTrafficResponseFlowList) SetTime(v int) *DescribeTrafficResponseFlowList {
	s.Time = &v
	return s
}

type DeleteIpRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	IpList   *string `json:"IpList,omitempty" xml:"IpList,omitempty" require:"true"`
}

func (s DeleteIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpRequest) SetSourceIp(v string) *DeleteIpRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteIpRequest) SetPackId(v string) *DeleteIpRequest {
	s.PackId = &v
	return s
}

func (s *DeleteIpRequest) SetIpList(v string) *DeleteIpRequest {
	s.IpList = &v
	return s
}

type DeleteIpResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpResponse) SetRequestId(v string) *DeleteIpResponse {
	s.RequestId = &v
	return s
}

type DescribePackRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PackId   *string `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
}

func (s DescribePackRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackRequest) GoString() string {
	return s.String()
}

func (s *DescribePackRequest) SetSourceIp(v string) *DescribePackRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePackRequest) SetPackId(v string) *DescribePackRequest {
	s.PackId = &v
	return s
}

type DescribePackResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PackInfo  *DescribePackResponsePackInfo `json:"PackInfo,omitempty" xml:"PackInfo,omitempty" require:"true" type:"Struct"`
}

func (s DescribePackResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponse) GoString() string {
	return s.String()
}

func (s *DescribePackResponse) SetRequestId(v string) *DescribePackResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePackResponse) SetPackInfo(v *DescribePackResponsePackInfo) *DescribePackResponse {
	s.PackInfo = v
	return s
}

type DescribePackResponsePackInfo struct {
	Region                        *string                                 `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	AvailableDeleteBlackholeCount *int                                    `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty" require:"true"`
	IpList                        []*DescribePackResponsePackInfoIpList   `json:"IpList,omitempty" xml:"IpList,omitempty" require:"true" type:"Repeated"`
	PackConfig                    *DescribePackResponsePackInfoPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribePackResponsePackInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponsePackInfo) GoString() string {
	return s.String()
}

func (s *DescribePackResponsePackInfo) SetRegion(v string) *DescribePackResponsePackInfo {
	s.Region = &v
	return s
}

func (s *DescribePackResponsePackInfo) SetAvailableDeleteBlackholeCount(v int) *DescribePackResponsePackInfo {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribePackResponsePackInfo) SetIpList(v []*DescribePackResponsePackInfoIpList) *DescribePackResponsePackInfo {
	s.IpList = v
	return s
}

func (s *DescribePackResponsePackInfo) SetPackConfig(v *DescribePackResponsePackInfoPackConfig) *DescribePackResponsePackInfo {
	s.PackConfig = v
	return s
}

type DescribePackResponsePackInfoIpList struct {
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
}

func (s DescribePackResponsePackInfoIpList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponsePackInfoIpList) GoString() string {
	return s.String()
}

func (s *DescribePackResponsePackInfoIpList) SetIp(v string) *DescribePackResponsePackInfoIpList {
	s.Ip = &v
	return s
}

type DescribePackResponsePackInfoPackConfig struct {
	PackAdvThre   *int `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty" require:"true"`
	IpAdvanceThre *int `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty" require:"true"`
	IpBasicThre   *int `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty" require:"true"`
	PackBasicThre *int `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty" require:"true"`
	IpSpec        *int `json:"IpSpec,omitempty" xml:"IpSpec,omitempty" require:"true"`
}

func (s DescribePackResponsePackInfoPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribePackResponsePackInfoPackConfig) GoString() string {
	return s.String()
}

func (s *DescribePackResponsePackInfoPackConfig) SetPackAdvThre(v int) *DescribePackResponsePackInfoPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribePackResponsePackInfoPackConfig) SetIpAdvanceThre(v int) *DescribePackResponsePackInfoPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribePackResponsePackInfoPackConfig) SetIpBasicThre(v int) *DescribePackResponsePackInfoPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribePackResponsePackInfoPackConfig) SetPackBasicThre(v int) *DescribePackResponsePackInfoPackConfig {
	s.PackBasicThre = &v
	return s
}

func (s *DescribePackResponsePackInfoPackConfig) SetIpSpec(v int) *DescribePackResponsePackInfoPackConfig {
	s.IpSpec = &v
	return s
}

type DescribePackListRequest struct {
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePackListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListRequest) GoString() string {
	return s.String()
}

func (s *DescribePackListRequest) SetResourceGroupId(v string) *DescribePackListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePackListRequest) SetRegionId(v string) *DescribePackListRequest {
	s.RegionId = &v
	return s
}

type DescribePackListResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PackList  []*DescribePackListResponsePackList `json:"PackList,omitempty" xml:"PackList,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePackListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponse) GoString() string {
	return s.String()
}

func (s *DescribePackListResponse) SetRequestId(v string) *DescribePackListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePackListResponse) SetPackList(v []*DescribePackListResponsePackList) *DescribePackListResponse {
	s.PackList = v
	return s
}

type DescribePackListResponsePackList struct {
	AvailableDeleteBlackholeCount *int                                        `json:"AvailableDeleteBlackholeCount,omitempty" xml:"AvailableDeleteBlackholeCount,omitempty" require:"true"`
	Region                        *string                                     `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	PackId                        *string                                     `json:"PackId,omitempty" xml:"PackId,omitempty" require:"true"`
	PackConfig                    *DescribePackListResponsePackListPackConfig `json:"PackConfig,omitempty" xml:"PackConfig,omitempty" require:"true" type:"Struct"`
}

func (s DescribePackListResponsePackList) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponsePackList) GoString() string {
	return s.String()
}

func (s *DescribePackListResponsePackList) SetAvailableDeleteBlackholeCount(v int) *DescribePackListResponsePackList {
	s.AvailableDeleteBlackholeCount = &v
	return s
}

func (s *DescribePackListResponsePackList) SetRegion(v string) *DescribePackListResponsePackList {
	s.Region = &v
	return s
}

func (s *DescribePackListResponsePackList) SetPackId(v string) *DescribePackListResponsePackList {
	s.PackId = &v
	return s
}

func (s *DescribePackListResponsePackList) SetPackConfig(v *DescribePackListResponsePackListPackConfig) *DescribePackListResponsePackList {
	s.PackConfig = v
	return s
}

type DescribePackListResponsePackListPackConfig struct {
	PackAdvThre   *int `json:"PackAdvThre,omitempty" xml:"PackAdvThre,omitempty" require:"true"`
	IpAdvanceThre *int `json:"IpAdvanceThre,omitempty" xml:"IpAdvanceThre,omitempty" require:"true"`
	IpBasicThre   *int `json:"IpBasicThre,omitempty" xml:"IpBasicThre,omitempty" require:"true"`
	PackBasicThre *int `json:"PackBasicThre,omitempty" xml:"PackBasicThre,omitempty" require:"true"`
	IpSpec        *int `json:"IpSpec,omitempty" xml:"IpSpec,omitempty" require:"true"`
}

func (s DescribePackListResponsePackListPackConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribePackListResponsePackListPackConfig) GoString() string {
	return s.String()
}

func (s *DescribePackListResponsePackListPackConfig) SetPackAdvThre(v int) *DescribePackListResponsePackListPackConfig {
	s.PackAdvThre = &v
	return s
}

func (s *DescribePackListResponsePackListPackConfig) SetIpAdvanceThre(v int) *DescribePackListResponsePackListPackConfig {
	s.IpAdvanceThre = &v
	return s
}

func (s *DescribePackListResponsePackListPackConfig) SetIpBasicThre(v int) *DescribePackListResponsePackListPackConfig {
	s.IpBasicThre = &v
	return s
}

func (s *DescribePackListResponsePackListPackConfig) SetPackBasicThre(v int) *DescribePackListResponsePackListPackConfig {
	s.PackBasicThre = &v
	return s
}

func (s *DescribePackListResponsePackListPackConfig) SetIpSpec(v int) *DescribePackListResponsePackListPackConfig {
	s.IpSpec = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-beijing":            tea.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote":          tea.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai":           tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen":           tea.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-1":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("ddosbgp.aliyuncs.com"),
		"eu-central-1":          tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("ddosbgp.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("ddosbgp.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("ddosbgp.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddosbgp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DescribeOnDemandInstanceWithOptions(request *DescribeOnDemandInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeOnDemandInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOnDemandInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOnDemandInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOnDemandInstance(request *DescribeOnDemandInstanceRequest) (_result *DescribeOnDemandInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOnDemandInstanceResponse{}
	_body, _err := client.DescribeOnDemandInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyOnDemaondDefenseStatusWithOptions(request *ModifyOnDemaondDefenseStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyOnDemaondDefenseStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyOnDemaondDefenseStatusResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyOnDemaondDefenseStatus"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyOnDemaondDefenseStatus(request *ModifyOnDemaondDefenseStatusRequest) (_result *ModifyOnDemaondDefenseStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOnDemaondDefenseStatusResponse{}
	_body, _err := client.ModifyOnDemaondDefenseStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOpEntities"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackPaidTrafficWithOptions(request *DescribePackPaidTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribePackPaidTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePackPaidTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackPaidTraffic(request *DescribePackPaidTrafficRequest) (_result *DescribePackPaidTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackPaidTrafficResponse{}
	_body, _err := client.DescribePackPaidTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackUsageWithOptions(request *DescribeResourcePackUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourcePackUsage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackUsage(request *DescribeResourcePackUsageRequest) (_result *DescribeResourcePackUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackUsageResponse{}
	_body, _err := client.DescribeResourcePackUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackStatisticsWithOptions(request *DescribeResourcePackStatisticsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourcePackStatistics"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackStatistics(request *DescribeResourcePackStatisticsRequest) (_result *DescribeResourcePackStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackStatisticsResponse{}
	_body, _err := client.DescribeResourcePackStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourcePackInstancesWithOptions(request *DescribeResourcePackInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeResourcePackInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourcePackInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourcePackInstances(request *DescribeResourcePackInstancesRequest) (_result *DescribeResourcePackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourcePackInstancesResponse{}
	_body, _err := client.DescribeResourcePackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *util.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteBlackhole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBlackhole(request *DeleteBlackholeRequest) (_result *DeleteBlackholeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DeleteBlackholeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *util.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.DoRequest(tea.String("CheckGrant"), tea.String("HTTPS"), tea.String("GET"), tea.String("2017-11-20"), tea.String("AK"), tea.ToMap(request), nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckGrant(request *CheckGrantRequest) (_result *CheckGrantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckGrantResponse{}
	_body, _err := client.CheckGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductWithOptions(request *DeleteProductRequest, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteProduct"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProduct(request *DeleteProductRequest) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProductWithOptions(request *AddProductRequest, runtime *util.RuntimeOptions) (_result *AddProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddProductResponse{}
	_body, _err := client.DoRequest(tea.String("AddProduct"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProduct(request *AddProductRequest) (_result *AddProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddProductResponse{}
	_body, _err := client.AddProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *util.RuntimeOptions) (_result *AddIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddIpResponse{}
	_body, _err := client.DoRequest(tea.String("AddIp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIp(request *AddIpRequest) (_result *AddIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpResponse{}
	_body, _err := client.AddIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeInstanceList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (_result *DescribeInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DescribeInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopTrafficWithOptions(request *DescribeTopTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTopTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTopTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeTopTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopTraffic(request *DescribeTopTrafficRequest) (_result *DescribeTopTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopTrafficResponse{}
	_body, _err := client.DescribeTopTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *util.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDdosEvent"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (_result *DescribeDdosEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DescribeDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeTraffic"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraffic(request *DescribeTrafficRequest) (_result *DescribeTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DescribeTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *util.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteIp"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIp(request *DeleteIpRequest) (_result *DeleteIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpResponse{}
	_body, _err := client.DeleteIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackWithOptions(request *DescribePackRequest, runtime *util.RuntimeOptions) (_result *DescribePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePackResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePack"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePack(request *DescribePackRequest) (_result *DescribePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackResponse{}
	_body, _err := client.DescribePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackListWithOptions(request *DescribePackListRequest, runtime *util.RuntimeOptions) (_result *DescribePackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePackListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePackList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2017-11-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackList(request *DescribePackListRequest) (_result *DescribePackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackListResponse{}
	_body, _err := client.DescribePackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
