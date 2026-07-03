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
	// The list of instances.
	DBInstances []*DescribeDBInstancesResponseBodyDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
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
	if s.DBInstances != nil {
		for _, item := range s.DBInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstances struct {
	// The name of the log node.
	//
	// example:
	//
	// pxc-c-dmlgit****
	CdcInstanceName *string `json:"CdcInstanceName,omitempty" xml:"CdcInstanceName,omitempty"`
	// The CN node specifications. Valid values:
	//
	// - **polarx.x4.medium.2e**: 2 cores, 8 GB
	//
	// - **polarx.x4.large.2e**: 4 cores, 16 GB
	//
	// - **polarx.x8.large.2e**: 4 cores, 32 GB
	//
	// - **polarx.x4.xlarge.2e**: 8 cores, 32 GB
	//
	// - **polarx.x8.xlarge.2e**: 8 cores, 64 GB
	//
	// - **polarx.x4.2xlarge.2e**: 16 cores, 64 GB
	//
	// - **polarx.x8.2xlarge.2e**: 16 cores, 128 GB
	//
	// - **polarx.x4.4xlarge.2e**: 32 cores, 128 GB
	//
	// - **polarx.x8.4xlarge.2e**: 32 cores, 256 GB
	//
	// - **polarx.st.8xlarge.2e**: 60 cores, 470 GB
	//
	// - **polarx.st.12xlarge.2e**: 90 cores, 720 GB.
	//
	// example:
	//
	// polarx.x4.large.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// The number of CN nodes.
	//
	// example:
	//
	// 2
	CnNodeCount *int32 `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	// The name of the column store engine instance.
	//
	// example:
	//
	// xxxxxxxx
	ColumnarInstanceName *string `json:"ColumnarInstanceName,omitempty" xml:"ColumnarInstanceName,omitempty"`
	// The column store read-only instance information.
	ColumnarReadDBInstances []*string `json:"ColumnarReadDBInstances,omitempty" xml:"ColumnarReadDBInstances,omitempty" type:"Repeated"`
	// The commodity code.
	//
	// example:
	//
	// drds_polarxpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Indicates whether the instance contains the multi-stream log service. Valid values:
	//
	// - **true**: The instance contains the multi-stream log service.
	//
	// - **false**: The instance does not contain the multi-stream log service.
	//
	// example:
	//
	// true
	ContainBinlogX *bool   `json:"ContainBinlogX,omitempty" xml:"ContainBinlogX,omitempty"`
	CpuType        *string `json:"CpuType,omitempty" xml:"CpuType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-11-01T03:49:50.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database instance name.
	//
	// example:
	//
	// pxc-xxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The database type.
	//
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database version.
	//
	// example:
	//
	// 5.7
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The database description.
	//
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The DN node specifications. Valid values:
	//
	// - **mysql.n2.medium.25**: 2 cores, 4 GB
	//
	// - **mysql.n4.medium.25**: 2 cores, 8 GB
	//
	// - **mysql.x8.medium.25**: 2 cores, 16 GB
	//
	// - **mysql.n2.large.25**: 4 cores, 8 GB
	//
	// - **mysql.n4.large.25**: 4 cores, 16 GB
	//
	// - **mysql.x8.large.25**: 4 cores, 32 GB
	//
	// - **mysql.n2.xlarge.25**: 8 cores, 16 GB
	//
	// - **mysql.n4.xlarge.25**: 8 cores, 32 GB
	//
	// - **mysql.x8.xlarge.25**: 8 cores, 64 GB
	//
	// - **mysql.n4.2xlarge.25**: 16 cores, 64 GB
	//
	// - **mysql.x8.2xlarge.25**: 16 cores, 128 GB
	//
	// - **mysql.x4.4xlarge.25**: 32 cores, 128 GB
	//
	// - **mysql.x8.4xlarge.25**: 32 cores, 256 GB
	//
	// - **mysql.st.8xlarge.25**: 60 cores, 470 GB
	//
	// - **mysql.st.12xlarge.25**: 90 cores, 720 GB.
	//
	// example:
	//
	// mysql.n4.medium.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// The number of DN nodes.
	//
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// The engine type.
	//
	// example:
	//
	// polarx
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2021-12-01T16:00:00.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the instance has expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-hzr2yeov******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the instance is locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The lock reason.
	//
	// example:
	//
	// 欠费
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The current Milvus version.
	//
	// example:
	//
	// polarx-kernel_5.4.12-16349923_xcluster-20210926
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The network type.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 5
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The list of nodes.
	Nodes []*DescribeDBInstancesResponseBodyDBInstancesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType           *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// The primary zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The list of read-only instances.
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-hangzhou-l
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// The instance edition. Valid values:
	//
	// - **enterprise**: Enterprise Edition.
	//
	// - **standard**: Standard Edition.
	//
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// The instance status. For more information, see [Instance status table](https://help.aliyun.com/document_detail/339826.html).
	//
	// example:
	//
	// Running
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The storage usage.
	//
	// example:
	//
	// 40658534400
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// Indicates whether the instance supports multi-stream. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	SupportBinlogX *bool `json:"SupportBinlogX,omitempty" xml:"SupportBinlogX,omitempty"`
	// The set of tags.
	TagSet []*DescribeDBInstancesResponseBodyDBInstancesTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// The third zone in the three-zone deployment.
	//
	// example:
	//
	// cn-hangzhou-k
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// The topology type. Valid values:
	//
	// - **3azones**: three-zone deployment.
	//
	// - **1azone**: single-zone deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The instance type. Valid values:
	//
	// - **ReadWrite**: primary instance.
	//
	// - **ReadOnly**: read-only instance.
	//
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
	// The zone ID.
	//
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

func (s *DescribeDBInstancesResponseBodyDBInstances) GetEngineVersion() *string {
	return s.EngineVersion
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

func (s *DescribeDBInstancesResponseBodyDBInstances) SetEngineVersion(v string) *DescribeDBInstancesResponseBodyDBInstances {
	s.EngineVersion = &v
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
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagSet != nil {
		for _, item := range s.TagSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesResponseBodyDBInstancesNodes struct {
	// The instance specifications.
	//
	// example:
	//
	// polarx.x4.large.2e
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pxi-zd89wrzqh******
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// if can be null:
	// true
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou-g-aliyun
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone ID.
	//
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

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) GetName() *string {
	return s.Name
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

func (s *DescribeDBInstancesResponseBodyDBInstancesNodes) SetName(v string) *DescribeDBInstancesResponseBodyDBInstancesNodes {
	s.Name = &v
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
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
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
