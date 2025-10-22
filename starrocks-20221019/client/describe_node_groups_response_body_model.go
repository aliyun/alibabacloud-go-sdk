// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeNodeGroupsResponseBody
	GetAccessDeniedDetail() *string
	SetData(v []*DescribeNodeGroupsResponseBodyData) *DescribeNodeGroupsResponseBody
	GetData() []*DescribeNodeGroupsResponseBodyData
	SetErrCode(v string) *DescribeNodeGroupsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeNodeGroupsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeNodeGroupsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeNodeGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNodeGroupsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeNodeGroupsResponseBody
	GetTotal() *int32
}

type DescribeNodeGroupsResponseBody struct {
	// example:
	//
	// {     "PolicyType": "AccountLevelIdentityBasedPolicy",     "AuthPrincipalOwnerId": "xxx",     "EncodedDiagnosticMessage": "xxx",     "AuthPrincipalType": "xxx",     "AuthPrincipalDisplayName": "xxx",     "NoPermissionType": "ImplicitDeny",     "AuthAction": "sr:xxx"   }
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*DescribeNodeGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// F285E0EC-E928-5E0B-B2C3-7F3FE42F2BE5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeNodeGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeNodeGroupsResponseBody) GetData() []*DescribeNodeGroupsResponseBodyData {
	return s.Data
}

func (s *DescribeNodeGroupsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeNodeGroupsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeNodeGroupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeNodeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNodeGroupsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeNodeGroupsResponseBody) SetAccessDeniedDetail(v string) *DescribeNodeGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetData(v []*DescribeNodeGroupsResponseBodyData) *DescribeNodeGroupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetErrCode(v string) *DescribeNodeGroupsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetErrMessage(v string) *DescribeNodeGroupsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetHttpStatusCode(v int32) *DescribeNodeGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetRequestId(v string) *DescribeNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetSuccess(v bool) *DescribeNodeGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) SetTotal(v int32) *DescribeNodeGroupsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeNodeGroupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeGroupsResponseBodyData struct {
	// example:
	//
	// OVERDUE
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// example:
	//
	// onEci
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// 1742179008000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// c-96f3bc7f04b2****
	BillingInstanceId *string `json:"BillingInstanceId,omitempty" xml:"BillingInstanceId,omitempty"`
	// example:
	//
	// emapreduce_starrockspost_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// FE
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// 1
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// true
	DefaultGroup *bool   `json:"DefaultGroup,omitempty" xml:"DefaultGroup,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	DiskNumber *int32 `json:"DiskNumber,omitempty" xml:"DiskNumber,omitempty"`
	// example:
	//
	// 3
	ElasticNodeNumber *int32 `json:"ElasticNodeNumber,omitempty" xml:"ElasticNodeNumber,omitempty"`
	// example:
	//
	// false
	EnablePublicNetwork *bool `json:"EnablePublicNetwork,omitempty" xml:"EnablePublicNetwork,omitempty"`
	// example:
	//
	// be-c-79b0bb9f58ef****-internal.starrocks.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 1742179008000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 8030
	HttpPort *int32 `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// example:
	//
	// c-96f3bc7f04b2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// local_ssd_4_4xlarge
	LocalStorageInstanceType *string `json:"LocalStorageInstanceType,omitempty" xml:"LocalStorageInstanceType,omitempty"`
	// example:
	//
	// 4
	MemoryCpuRatio *int32 `json:"MemoryCpuRatio,omitempty" xml:"MemoryCpuRatio,omitempty"`
	// example:
	//
	// ng-5e2ba600fee3****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// ng_1
	NodeGroupName *string                                       `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	NodeInfo      []*DescribeNodeGroupsResponseBodyDataNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
	// example:
	//
	// prePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// fe-c-79b0bb9f58ef****.starrocks.aliyuncs.com
	PublicAddress *string `json:"PublicAddress,omitempty" xml:"PublicAddress,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3
	ResidentNodeNumber *int32 `json:"ResidentNodeNumber,omitempty" xml:"ResidentNodeNumber,omitempty"`
	// example:
	//
	// 10000
	RunningTime *int64 `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"StoragePerformanceLevel,omitempty" xml:"StoragePerformanceLevel,omitempty"`
	// example:
	//
	// 100
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// 3
	TargetElasticNodeNumber *int32 `json:"TargetElasticNodeNumber,omitempty" xml:"TargetElasticNodeNumber,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNodeGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupsResponseBodyData) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeNodeGroupsResponseBodyData) GetArchitecture() *string {
	return s.Architecture
}

func (s *DescribeNodeGroupsResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeNodeGroupsResponseBodyData) GetBillingInstanceId() *string {
	return s.BillingInstanceId
}

func (s *DescribeNodeGroupsResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeNodeGroupsResponseBodyData) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeNodeGroupsResponseBodyData) GetCu() *int32 {
	return s.Cu
}

func (s *DescribeNodeGroupsResponseBodyData) GetDefaultGroup() *bool {
	return s.DefaultGroup
}

func (s *DescribeNodeGroupsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeNodeGroupsResponseBodyData) GetDiskNumber() *int32 {
	return s.DiskNumber
}

func (s *DescribeNodeGroupsResponseBodyData) GetElasticNodeNumber() *int32 {
	return s.ElasticNodeNumber
}

func (s *DescribeNodeGroupsResponseBodyData) GetEnablePublicNetwork() *bool {
	return s.EnablePublicNetwork
}

func (s *DescribeNodeGroupsResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeNodeGroupsResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeNodeGroupsResponseBodyData) GetHttpPort() *int32 {
	return s.HttpPort
}

func (s *DescribeNodeGroupsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNodeGroupsResponseBodyData) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *DescribeNodeGroupsResponseBodyData) GetMemoryCpuRatio() *int32 {
	return s.MemoryCpuRatio
}

func (s *DescribeNodeGroupsResponseBodyData) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeGroupsResponseBodyData) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeNodeGroupsResponseBodyData) GetNodeInfo() []*DescribeNodeGroupsResponseBodyDataNodeInfo {
	return s.NodeInfo
}

func (s *DescribeNodeGroupsResponseBodyData) GetPayType() *string {
	return s.PayType
}

func (s *DescribeNodeGroupsResponseBodyData) GetPublicAddress() *string {
	return s.PublicAddress
}

func (s *DescribeNodeGroupsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNodeGroupsResponseBodyData) GetResidentNodeNumber() *int32 {
	return s.ResidentNodeNumber
}

func (s *DescribeNodeGroupsResponseBodyData) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *DescribeNodeGroupsResponseBodyData) GetSpecType() *string {
	return s.SpecType
}

func (s *DescribeNodeGroupsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeNodeGroupsResponseBodyData) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *DescribeNodeGroupsResponseBodyData) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *DescribeNodeGroupsResponseBodyData) GetTargetElasticNodeNumber() *int32 {
	return s.TargetElasticNodeNumber
}

func (s *DescribeNodeGroupsResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeNodeGroupsResponseBodyData) SetAccountStatus(v string) *DescribeNodeGroupsResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetArchitecture(v string) *DescribeNodeGroupsResponseBodyData {
	s.Architecture = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetBeginTime(v int64) *DescribeNodeGroupsResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetBillingInstanceId(v string) *DescribeNodeGroupsResponseBodyData {
	s.BillingInstanceId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetCommodityCode(v string) *DescribeNodeGroupsResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetComponentType(v string) *DescribeNodeGroupsResponseBodyData {
	s.ComponentType = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetCu(v int32) *DescribeNodeGroupsResponseBodyData {
	s.Cu = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetDefaultGroup(v bool) *DescribeNodeGroupsResponseBodyData {
	s.DefaultGroup = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetDescription(v string) *DescribeNodeGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetDiskNumber(v int32) *DescribeNodeGroupsResponseBodyData {
	s.DiskNumber = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetElasticNodeNumber(v int32) *DescribeNodeGroupsResponseBodyData {
	s.ElasticNodeNumber = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetEnablePublicNetwork(v bool) *DescribeNodeGroupsResponseBodyData {
	s.EnablePublicNetwork = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetEndpoint(v string) *DescribeNodeGroupsResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetExpireTime(v int64) *DescribeNodeGroupsResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetHttpPort(v int32) *DescribeNodeGroupsResponseBodyData {
	s.HttpPort = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetInstanceId(v string) *DescribeNodeGroupsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetLocalStorageInstanceType(v string) *DescribeNodeGroupsResponseBodyData {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetMemoryCpuRatio(v int32) *DescribeNodeGroupsResponseBodyData {
	s.MemoryCpuRatio = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetNodeGroupId(v string) *DescribeNodeGroupsResponseBodyData {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetNodeGroupName(v string) *DescribeNodeGroupsResponseBodyData {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetNodeInfo(v []*DescribeNodeGroupsResponseBodyDataNodeInfo) *DescribeNodeGroupsResponseBodyData {
	s.NodeInfo = v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetPayType(v string) *DescribeNodeGroupsResponseBodyData {
	s.PayType = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetPublicAddress(v string) *DescribeNodeGroupsResponseBodyData {
	s.PublicAddress = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetRegionId(v string) *DescribeNodeGroupsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetResidentNodeNumber(v int32) *DescribeNodeGroupsResponseBodyData {
	s.ResidentNodeNumber = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetRunningTime(v int64) *DescribeNodeGroupsResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetSpecType(v string) *DescribeNodeGroupsResponseBodyData {
	s.SpecType = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetStatus(v string) *DescribeNodeGroupsResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetStoragePerformanceLevel(v string) *DescribeNodeGroupsResponseBodyData {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetStorageSize(v int32) *DescribeNodeGroupsResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetTargetElasticNodeNumber(v int32) *DescribeNodeGroupsResponseBodyData {
	s.TargetElasticNodeNumber = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) SetZoneId(v string) *DescribeNodeGroupsResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyData) Validate() error {
	if s.NodeInfo != nil {
		for _, item := range s.NodeInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeGroupsResponseBodyDataNodeInfo struct {
	// example:
	//
	// cn-e1733d44-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeNodeGroupsResponseBodyDataNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupsResponseBodyDataNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupsResponseBodyDataNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeNodeGroupsResponseBodyDataNodeInfo) SetNodeId(v string) *DescribeNodeGroupsResponseBodyDataNodeInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeGroupsResponseBodyDataNodeInfo) Validate() error {
	return dara.Validate(s)
}
