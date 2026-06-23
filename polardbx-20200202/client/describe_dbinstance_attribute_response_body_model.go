// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstance(v *DescribeDBInstanceAttributeResponseBodyDBInstance) *DescribeDBInstanceAttributeResponseBody
	GetDBInstance() *DescribeDBInstanceAttributeResponseBodyDBInstance
	SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceAttributeResponseBody struct {
	// The database instance information.
	DBInstance *DescribeDBInstanceAttributeResponseBodyDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBody) GetDBInstance() *DescribeDBInstanceAttributeResponseBodyDBInstance {
	return s.DBInstance
}

func (s *DescribeDBInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceAttributeResponseBody) SetDBInstance(v *DescribeDBInstanceAttributeResponseBodyDBInstance) *DescribeDBInstanceAttributeResponseBody {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) SetRequestId(v string) *DescribeDBInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBody) Validate() error {
	if s.DBInstance != nil {
		if err := s.DBInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceAttributeResponseBodyDBInstance struct {
	AiGatewayEnabled *string `json:"AiGatewayEnabled,omitempty" xml:"AiGatewayEnabled,omitempty"`
	// Indicates whether the In-Memory Column Index feature is supported.
	//
	// example:
	//
	// true
	CanNotCreateColumnar *bool `json:"CanNotCreateColumnar,omitempty" xml:"CanNotCreateColumnar,omitempty"`
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
	// polarx.x4.xlarge.2e
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
	// xxxx
	ColumnarInstanceName *string `json:"ColumnarInstanceName,omitempty" xml:"ColumnarInstanceName,omitempty"`
	// The column store read-only instance information.
	ColumnarReadDBInstances []*string `json:"ColumnarReadDBInstances,omitempty" xml:"ColumnarReadDBInstances,omitempty" type:"Repeated"`
	// The commodity code of the instance. The value is fixed as drds_polarxpost_public_cn.
	//
	// example:
	//
	// drds_polarxpost_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The endpoint information.
	ConnAddrs []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// The internal network connection string.
	//
	// example:
	//
	// pxc-sprpx766vo****.polarx.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CpuType          *string `json:"CpuType,omitempty" xml:"CpuType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-08-31T08:56:25.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance type. Valid values:
	//
	// - **ReadWrite**: primary instance.
	//
	// - **ReadOnly**: read-only instance.
	//
	// example:
	//
	// ReadWrite
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The node specifications of the instance.
	//
	// example:
	//
	// polarx.x4.large.2e
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of instance nodes.
	//
	// example:
	//
	// 2
	DBNodeCount *int32 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The database node information.
	DBNodes []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// The database type. The value is fixed as polarx.
	//
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database version.
	//
	// example:
	//
	// 5.5
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The instance description.
	//
	// example:
	//
	// test instance
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the DN nodes of the instance have different specifications. Valid values:
	//
	// - true: The specifications are different.
	//
	// - false: The specifications are the same.
	//
	// example:
	//
	// false
	DifferentDNSpec *bool `json:"DifferentDNSpec,omitempty" xml:"DifferentDNSpec,omitempty"`
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
	// mysql.x8.large.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// The number of DN nodes.
	//
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// The disk space of the DN data node, in GB.
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// The database type. The value is fixed as polarx.
	//
	// example:
	//
	// polarx
	Engine        *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration time. Format: yyyy-MM-ddTHH:mm:ss.sss+0000 (UTC).
	//
	// example:
	//
	// 2022-08-31T16:00:00.000+0000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// Indicates whether the instance has expired. Valid values:
	//
	// - **true**: Expired.
	//
	// - **false**: Not expired.
	//
	// example:
	//
	// false
	Expired         *string                                                           `json:"Expired,omitempty" xml:"Expired,omitempty"`
	GdnInstanceName *string                                                           `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	GdnMemberList   []*DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList `json:"GdnMemberList,omitempty" xml:"GdnMemberList,omitempty" type:"Repeated"`
	GdnRole         *string                                                           `json:"GdnRole,omitempty" xml:"GdnRole,omitempty"`
	// The ID of the primary instance. If this parameter is not returned, the instance is a primary instance.
	//
	// example:
	//
	// pxc-zkralxpc5d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The engine version of the instance. This is an internal parameter.
	//
	// example:
	//
	// 18
	KindCode *int32 `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// This parameter is required.
	LTSVersions []*string `json:"LTSVersions,omitempty" xml:"LTSVersions,omitempty" type:"Repeated"`
	// The latest minor engine version supported by the instance.
	//
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	LatestMinorVersion *string `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// - **Unlock**: Normal.
	//
	// - **ManualLock**: Manually locked.
	//
	// - **LockByExpiration**: Automatically locked due to instance expiration.
	//
	// - **LockByRestoration**: Automatically locked before instance rollback.
	//
	// - **LockByDiskQuota**: Automatically locked due to insufficient disk space.
	//
	// - **LockReadInstanceByDiskQuota**: Read-only instance automatically locked due to insufficient disk space.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The end time of the maintenance window. The time is in UTC. Add 8 hours to obtain the maintenance window displayed in the console.
	//
	// example:
	//
	// 06:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is in UTC. Add 8 hours to obtain the maintenance window displayed in the console.
	//
	// example:
	//
	// 06:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The current minor engine version.
	//
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The network type of the instance. Only VPC is supported, which indicates Virtual Private Cloud.
	//
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The internal network connection port.
	//
	// example:
	//
	// 3306
	Port              *string `json:"Port,omitempty" xml:"Port,omitempty"`
	PrimaryInstanceId *string `json:"PrimaryInstanceId,omitempty" xml:"PrimaryInstanceId,omitempty"`
	// The primary zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-e
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// The list of read-only instance names.
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-*********
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The three-role mode status. Valid values:
	//
	// - **false**: Disabled.
	//
	// - **true**: Enabled.
	//
	// example:
	//
	// false
	RightsSeparationEnabled *bool `json:"RightsSeparationEnabled,omitempty" xml:"RightsSeparationEnabled,omitempty"`
	// The three-role mode status. Valid values:
	//
	// - **disabled**: Disabled.
	//
	// - **enabled**: Enabled.
	//
	// - **processing**: Being processed.
	//
	// - **unknown**: Unknown. This may be caused by the instance being unreachable.
	//
	// example:
	//
	// disabled
	RightsSeparationStatus *string `json:"RightsSeparationStatus,omitempty" xml:"RightsSeparationStatus,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-shenzhen-a
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
	Series       *string `json:"Series,omitempty" xml:"Series,omitempty"`
	SpecCategory *string `json:"SpecCategory,omitempty" xml:"SpecCategory,omitempty"`
	// The instance status. For more information, see [Instance status table](https://help.aliyun.com/document_detail/339826.html).
	//
	// example:
	//
	// Running
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The used storage space, in bytes.
	//
	// example:
	//
	// 17042505728
	StorageUsed *int64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// The tag set.
	TagSet []*DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// The tertiary zone for Three-zone deployment.
	//
	// example:
	//
	// cn-shenzhen-e
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
	// vpc-xxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetAiGatewayEnabled() *string {
	return s.AiGatewayEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetCanNotCreateColumnar() *bool {
	return s.CanNotCreateColumnar
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetCnNodeClassCode() *string {
	return s.CnNodeClassCode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetCnNodeCount() *int32 {
	return s.CnNodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetColumnarInstanceName() *string {
	return s.ColumnarInstanceName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetColumnarReadDBInstances() []*string {
	return s.ColumnarReadDBInstances
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetConnAddrs() []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetCpuType() *string {
	return s.CpuType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDBNodes() []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	return s.DBNodes
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDifferentDNSpec() *bool {
	return s.DifferentDNSpec
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDnNodeClassCode() *string {
	return s.DnNodeClassCode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDnNodeCount() *int32 {
	return s.DnNodeCount
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetDnStorageSpace() *string {
	return s.DnStorageSpace
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetGdnInstanceName() *string {
	return s.GdnInstanceName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetGdnMemberList() []*DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	return s.GdnMemberList
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetGdnRole() *string {
	return s.GdnRole
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetId() *string {
	return s.Id
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetKindCode() *int32 {
	return s.KindCode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetLTSVersions() []*string {
	return s.LTSVersions
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetLatestMinorVersion() *string {
	return s.LatestMinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetNetwork() *string {
	return s.Network
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetPrimaryInstanceId() *string {
	return s.PrimaryInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetReadDBInstances() []*string {
	return s.ReadDBInstances
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetRightsSeparationEnabled() *bool {
	return s.RightsSeparationEnabled
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetRightsSeparationStatus() *string {
	return s.RightsSeparationStatus
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetSeries() *string {
	return s.Series
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetSpecCategory() *string {
	return s.SpecCategory
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetTagSet() []*DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet {
	return s.TagSet
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetTopologyType() *string {
	return s.TopologyType
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetAiGatewayEnabled(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.AiGatewayEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCanNotCreateColumnar(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CanNotCreateColumnar = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCnNodeClassCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCnNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetColumnarInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ColumnarInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetColumnarReadDBInstances(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ColumnarReadDBInstances = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCommodityCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetConnAddrs(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCpuType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CpuType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetCreateTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBInstanceType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBNodes(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDBVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDescription(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDifferentDNSpec(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DifferentDNSpec = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDnNodeClassCode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDnNodeCount(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetDnStorageSpace(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.DnStorageSpace = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetEngine(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetExpireDate(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetExpired(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetGdnInstanceName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.GdnInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetGdnMemberList(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.GdnMemberList = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetGdnRole(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.GdnRole = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetKindCode(v int32) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLTSVersions(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LTSVersions = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLatestMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetLockMode(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetMinorVersion(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetNetwork(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Network = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPayType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPort(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPrimaryInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.PrimaryInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetPrimaryZone(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetReadDBInstances(v []*string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetResourceGroupId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRightsSeparationEnabled(v bool) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RightsSeparationEnabled = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetRightsSeparationStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.RightsSeparationStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetSecondaryZone(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.SecondaryZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetSeries(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Series = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetSpecCategory(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.SpecCategory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStorageType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetStorageUsed(v int64) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetTagSet(v []*DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.TagSet = v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetTertiaryZone(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.TertiaryZone = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetTopologyType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.TopologyType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstance) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GdnMemberList != nil {
		for _, item := range s.GdnMemberList {
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

type DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs struct {
	// The endpoint.
	//
	// example:
	//
	// polardbx-xxx.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The connection port number.
	//
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The connection type. **VPC*	- indicates an internal network connection. **PUBLIC*	- indicates a public network connection.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The internal CloudInstanceId within the VPC. You can ignore this parameter.
	//
	// example:
	//
	// pxc-zkralxpc5d****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GetPort() *int64 {
	return s.Port
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetPort(v int64) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetType(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes struct {
	// The name of the compute node.
	//
	// example:
	//
	// pxc-i-********
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" xml:"ComputeNodeId,omitempty"`
	// The name of the storage node.
	//
	// example:
	//
	// pxc-xdb-xxxxxx
	DataNodeId *string `json:"DataNodeId,omitempty" xml:"DataNodeId,omitempty"`
	// The logical node ID.
	//
	// example:
	//
	// pxi-*********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node specifications.
	//
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// The region ID of the node.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The zone in which the node resides.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GetComputeNodeId() *string {
	return s.ComputeNodeId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GetDataNodeId() *string {
	return s.DataNodeId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GetId() *string {
	return s.Id
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetComputeNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.ComputeNodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetDataNodeId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.DataNodeId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetRegionId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) SetZoneId(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList struct {
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) GetMemberName() *string {
	return s.MemberName
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) GetRole() *string {
	return s.Role
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) SetMemberName(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	s.MemberName = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) SetRole(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) SetStatus(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceGdnMemberList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet struct {
	// The tag key.
	//
	// example:
	//
	// key2
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) SetKey(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) SetValue(v string) *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseBodyDBInstanceTagSet) Validate() error {
	return dara.Validate(s)
}
