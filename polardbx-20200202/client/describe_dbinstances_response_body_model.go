// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstances(v []*DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody
	GetDBInstances() []*DescribeDBInstancesResponseBodyDBInstances
	SetPageNumber(v int32) *DescribeDBInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBInstancesResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *DescribeDBInstancesResponseBody
	GetTotalNumber() *int32
}

type DescribeDBInstancesResponseBody struct {
	DBInstances []*DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeDBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBody) GetDBInstances() []*DescribeDBInstancesResponseBodyDBInstances {
	return s.DBInstances
}

func (s *DescribeDBInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancesResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *DescribeDBInstancesResponseBody) SetDBInstances(v []*DescribeDBInstancesResponseBodyDBInstances) *DescribeDBInstancesResponseBody {
	s.DBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageNumber(v int32) *DescribeDBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetPageSize(v int32) *DescribeDBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetRequestId(v string) *DescribeDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) SetTotalNumber(v int32) *DescribeDBInstancesResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *DescribeDBInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyDBInstances struct {
	// example:
	//
	// pxc-c-dmlgit****
	CdcInstanceName *string `json:"CdcInstanceName,omitempty" xml:"CdcInstanceName,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount             *int32    `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	ColumnarInstanceName    *string   `json:"ColumnarInstanceName,omitempty" xml:"ColumnarInstanceName,omitempty"`
	ColumnarReadDBInstances []*string `json:"ColumnarReadDBInstances,omitempty" xml:"ColumnarReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// drds_polarxpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// true
	ContainBinlogX *bool   `json:"ContainBinlogX,omitempty" xml:"ContainBinlogX,omitempty"`
	CpuType        *string `json:"CpuType,omitempty" xml:"CpuType,omitempty"`
	// example:
	//
	// 2021-11-01T03:49:50.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// pxc-xxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.7
	DBVersion   *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mysql.n4.medium.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2021-12-01T16:00:00.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// pxc-hzr2yeov9jmg3z
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Unlock
	LockMode   *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.12-16349923_xcluster-20210926
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// example:
	//
	// 5
	NodeCount *int32                                             `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	Nodes     []*DescribeDBInstancesResponseBodyDBInstancesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// Prepaid
	PayType           *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// 主可用区。
	//
	// This parameter is required.
	PrimaryZone     *string   `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 次可用区。
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// Running
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// 40658534400
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// true
	SupportBinlogX *bool                                               `json:"SupportBinlogX,omitempty" xml:"SupportBinlogX,omitempty"`
	TagSet         []*DescribeDBInstancesResponseBodyDBInstancesTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// 第三可用区。
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// 拓扑类型：
	//
	// - **3azones**：三可用区；
	//
	// - **1azone**：单可用区。
	//
	// This parameter is required.
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// example:
	//
	// ReadWrite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// VPCID
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// cn-hangzhou-g
	ZoneId  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	GdnRole *string `json:"gdnRole,omitempty" xml:"gdnRole,omitempty"`
	IsInGdn *bool   `json:"isInGdn,omitempty" xml:"isInGdn,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetCdcInstanceName() *string {
	return s.CdcInstanceName
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetCnNodeClassCode() *string {
	return s.CnNodeClassCode
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetCnNodeCount() *int32 {
	return s.CnNodeCount
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetColumnarInstanceName() *string {
	return s.ColumnarInstanceName
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetColumnarReadDBInstances() []*string {
	return s.ColumnarReadDBInstances
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetContainBinlogX() *bool {
	return s.ContainBinlogX
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetCpuType() *string {
	return s.CpuType
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDnNodeClassCode() *string {
	return s.DnNodeClassCode
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetDnNodeCount() *int32 {
	return s.DnNodeCount
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetId() *string {
	return s.Id
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetNetwork() *string {
	return s.Network
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetNodes() []*DescribeDBInstancesResponseBodyDBInstancesNodes {
	return s.Nodes
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetPrimaryInstanceId() *string {
	return s.PrimaryInstanceId
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetReadDBInstances() []*string {
	return s.ReadDBInstances
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetSeries() *string {
	return s.Series
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetSupportBinlogX() *bool {
	return s.SupportBinlogX
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetTagSet() []*DescribeDBInstancesResponseBodyDBInstancesTagSet {
	return s.TagSet
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetTopologyType() *string {
	return s.TopologyType
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetGdnRole() *string {
	return s.GdnRole
}

func (s *DescribeDBInstancesResponseBodyDBInstances) GetIsInGdn() *bool {
	return s.IsInGdn
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCdcInstanceName(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CdcInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCnNodeClassCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCnNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetColumnarInstanceName(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ColumnarInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetColumnarReadDBInstances(v []*string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ColumnarReadDBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCommodityCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetContainBinlogX(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.ContainBinlogX = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCpuType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CpuType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetCreateTime(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBInstanceName(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDBVersion(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDescription(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDnNodeClassCode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetDnNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.DnNodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetEngine(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetExpireTime(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetExpired(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Id = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetLockMode(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetLockReason(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetMinorVersion(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNetwork(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Network = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodeClass(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodeCount(v int32) *DescribeDBInstancesResponseBodyDBInstances {
	s.NodeCount = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetNodes(v []*DescribeDBInstancesResponseBodyDBInstancesNodes) *DescribeDBInstancesResponseBodyDBInstances {
	s.Nodes = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPayType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPrimaryInstanceId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PrimaryInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetPrimaryZone(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetReadDBInstances(v []*string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ReadDBInstances = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetResourceGroupId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetSecondaryZone(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.SecondaryZone = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetSeries(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Series = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStatus(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Status = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStorageType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetStorageUsed(v int64) *DescribeDBInstancesResponseBodyDBInstances {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetSupportBinlogX(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.SupportBinlogX = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetTagSet(v []*DescribeDBInstancesResponseBodyDBInstancesTagSet) *DescribeDBInstancesResponseBodyDBInstances {
	s.TagSet = v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetTertiaryZone(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.TertiaryZone = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetTopologyType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.TopologyType = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetType(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.Type = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetVPCId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetGdnRole(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.GdnRole = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) SetIsInGdn(v bool) *DescribeDBInstancesResponseBodyDBInstances {
	s.IsInGdn = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyDBInstancesNodes struct {
	// example:
	//
	// polarx.x4.large.2e
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cn-hangzhou-g-aliyun
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) GetId() *string {
	return s.Id
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetClassCode(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.ClassCode = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetRegionId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetZoneId(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesResponseBodyDBInstancesTagSet struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesResponseBodyDBInstancesTagSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesResponseBodyDBInstancesTagSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) SetKey(v string) *DescribeDBInstancesResponseBodyDBInstancesTagSet {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) SetValue(v string) *DescribeDBInstancesResponseBodyDBInstancesTagSet {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesResponseBodyDBInstancesTagSet) Validate() error {
	return dara.Validate(s)
}
