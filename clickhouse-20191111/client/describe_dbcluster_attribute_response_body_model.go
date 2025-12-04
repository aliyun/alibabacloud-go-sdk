// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBCluster(v *DescribeDBClusterAttributeResponseBodyDBCluster) *DescribeDBClusterAttributeResponseBody
	GetDBCluster() *DescribeDBClusterAttributeResponseBodyDBCluster
	SetRequestId(v string) *DescribeDBClusterAttributeResponseBody
	GetRequestId() *string
}

type DescribeDBClusterAttributeResponseBody struct {
	// The information about the cluster.
	DBCluster *DescribeDBClusterAttributeResponseBodyDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) GetDBCluster() *DescribeDBClusterAttributeResponseBodyDBCluster {
	return s.DBCluster
}

func (s *DescribeDBClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBCluster(v *DescribeDBClusterAttributeResponseBodyDBCluster) *DescribeDBClusterAttributeResponseBody {
	s.DBCluster = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) Validate() error {
	if s.DBCluster != nil {
		if err := s.DBCluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyDBCluster struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 140692647406****
	AliUid                                *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppointmentElectZookeeperDisableWrite *bool   `json:"AppointmentElectZookeeperDisableWrite,omitempty" xml:"AppointmentElectZookeeperDisableWrite,omitempty"`
	AppointmentElectZookeeperTime         *string `json:"AppointmentElectZookeeperTime,omitempty" xml:"AppointmentElectZookeeperTime,omitempty"`
	AppointmentRestartNodeList            *string `json:"AppointmentRestartNodeList,omitempty" xml:"AppointmentRestartNodeList,omitempty"`
	AppointmentRestartNodeTime            *string `json:"AppointmentRestartNodeTime,omitempty" xml:"AppointmentRestartNodeTime,omitempty"`
	// The scheduled restart time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// 2023-11-06T12:00:00Z
	AppointmentRestartTime *string `json:"AppointmentRestartTime,omitempty" xml:"AppointmentRestartTime,omitempty"`
	// The major engine versions available for upgrades.
	//
	// example:
	//
	// {"MajorVersion":"MinorVersion"}
	AvailableUpgradeMajorVersion map[string]interface{} `json:"AvailableUpgradeMajorVersion,omitempty" xml:"AvailableUpgradeMajorVersion,omitempty"`
	// The site ID. Valid values:
	//
	// 	- **26842**: the China site (aliyun.com)
	//
	// 	- **26888**: the international site (alibabacloud.com)
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The edition of the cluster. Valid values:
	//
	// 	- **Basic**: Single-replica Edition
	//
	// 	- **HighAvailability**: Double-replica Edition
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The commodity code of the cluster.
	//
	// example:
	//
	// clickhouse_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The VPC endpoint of the cluster.
	//
	// example:
	//
	// cc-bp1qx68m06981****.ads.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The version of the ApsaraDB for ClickHouse console that is used to manage the cluster. Valid values:
	//
	// 	- **v1**
	//
	// 	- **v2**
	//
	// example:
	//
	// v1
	ControlVersion *string `json:"ControlVersion,omitempty" xml:"ControlVersion,omitempty"`
	// The time when the cluster was created. The value is in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-13T11:33:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The network type of the cluster. Only VPC is supported.
	//
	// example:
	//
	// vpc
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	// The cluster state. Valid values:
	//
	// 	- **Preparing**: The cluster is being prepared.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **Running**: The cluster is running.
	//
	// 	- **Deleting**: The cluster is being deleted.
	//
	// 	- **SCALING_OUT**: The storage capacity of the cluster is being expanded.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- **Common**: a common cluster
	//
	// 	- **Readonly**: a read-only cluster
	//
	// 	- **Guard**: a disaster recovery cluster
	//
	// example:
	//
	// Common
	DBClusterType *string `json:"DBClusterType,omitempty" xml:"DBClusterType,omitempty"`
	// The specifications of the cluster.
	//
	// 	- Valid values when the cluster is of Single-replica Edition:
	//
	//     	- **S4-NEW**
	//
	//     	- **S8**
	//
	//     	- **S16**
	//
	//     	- **S32**
	//
	//     	- **S64**
	//
	//     	- **S104**
	//
	// 	- Valid values when the cluster is of Double-replica Edition:
	//
	//     	- **C4-NEW**
	//
	//     	- **C8**
	//
	//     	- **C16**
	//
	//     	- **C32**
	//
	//     	- **C64**
	//
	//     	- **C104**
	//
	// example:
	//
	// C8
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The number of nodes.
	//
	// 	- Valid values when the cluster is of Single-replica Edition: 1 to 48.
	//
	// 	- Valid values when the cluster is of Double-replica Edition: 1 to 24.
	//
	// example:
	//
	// 1
	DBNodeCount *int64 `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	// The storage capacity of a single node of the cluster. Unit: GB.
	//
	// Valid values: 100 to 32000.
	//
	// >  This value is a multiple of 100.
	//
	// example:
	//
	// 100
	DBNodeStorage *int64 `json:"DBNodeStorage,omitempty" xml:"DBNodeStorage,omitempty"`
	// The Key Management Service (KMS) key that is used to encrypt data.
	//
	// >  If the value of the EncryptionType parameter is off, an empty string is returned for this parameter.
	//
	// example:
	//
	// 685f416f-87c9-4554-8d3a-75b6ce25****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Valid values:
	//
	// 	- **CloudDisk**: Disk encryption is enabled.
	//
	// 	- **off**: Data is not encrypted.
	//
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// ClickHouse
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The latest minor version to which the cluster can be updated.
	//
	// example:
	//
	// 1.34.0
	EngineLatestMinorVersion *string `json:"EngineLatestMinorVersion,omitempty" xml:"EngineLatestMinorVersion,omitempty"`
	// The current minor version.
	//
	// example:
	//
	// 1.6.0
	EngineMinorVersion *string `json:"EngineMinorVersion,omitempty" xml:"EngineMinorVersion,omitempty"`
	// The engine version.
	//
	// example:
	//
	// 21.8.10.19
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the cluster expired. The time is in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// >  Pay-as-you-go clusters never expire. If the cluster is a pay-as-you-go cluster, an empty string is returned for this parameter.
	//
	// example:
	//
	// 2022-11-11T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The extended storage space. Unit: GB.
	//
	// example:
	//
	// 500
	ExtStorageSize *int32 `json:"ExtStorageSize,omitempty" xml:"ExtStorageSize,omitempty"`
	// The extended storage type. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	ExtStorageType *string `json:"ExtStorageType,omitempty" xml:"ExtStorageType,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsExpired *string `json:"IsExpired,omitempty" xml:"IsExpired,omitempty"`
	// The lock mode of the cluster. Valid values:
	//
	// 	- **Unlock**: The cluster is not locked.
	//
	// 	- **ManualLock**: The cluster is manually locked.
	//
	// 	- **LockByExpiration**: The cluster is automatically locked due to cluster expiration.
	//
	// 	- **LockByRestoration**: The cluster is automatically locked because the cluster is about to be rolled back.
	//
	// 	- **LockByDiskQuota**: The cluster is automatically locked because the disk space is exhausted.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The cause why the cluster was locked.
	//
	// >  If the value of the LockMode parameter is Unlock, an empty string is returned for this parameter.
	//
	// example:
	//
	// DISK_FULL
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The update type. If the value of the parameter is **false**, it indicates a manual update.
	//
	// example:
	//
	// false
	MaintainAutoType *bool `json:"MaintainAutoType,omitempty" xml:"MaintainAutoType,omitempty"`
	// The maintenance window of the cluster. The value is in the HH:mmZ-HH:mmZ format. The time is displayed in UTC.
	//
	// For example, if you set the maintenance window to 00:00Z-01:00Z, the cluster can be maintained from 08:00 (UTC+8) to 09:00 (UTC+8).
	//
	// example:
	//
	// 00:00Z-01:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: indicates the pay-as-you-go billing method.
	//
	// 	- **Prepaid**: indicates the subscription billing method.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The HTTP port number.
	//
	// example:
	//
	// 8123
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The public endpoint.
	//
	// example:
	//
	// cc-bp1199ya710s7****.public.clickhouse.ads.aliyuncs.com
	PublicConnectionString *string `json:"PublicConnectionString,omitempty" xml:"PublicConnectionString,omitempty"`
	// The IP address that is used to connect to the cluster over the Internet.
	//
	// example:
	//
	// 121.40.xx.xx
	PublicIpAddr *string `json:"PublicIpAddr,omitempty" xml:"PublicIpAddr,omitempty"`
	// The TCP port number in the public endpoint.
	//
	// example:
	//
	// 3306
	PublicPort *string `json:"PublicPort,omitempty" xml:"PublicPort,omitempty"`
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
	// rg-acfmyf65je6****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the data migration task.
	ScaleOutStatus *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus `json:"ScaleOutStatus,omitempty" xml:"ScaleOutStatus,omitempty" type:"Struct"`
	// The storage type of the cluster. Valid values:
	//
	// 	- **CloudESSD**: The cluster uses an enhanced SSD (ESSD) of performance level (PL) 1.
	//
	// 	- **CloudESSD_PL2**: The cluster uses an ESSD of PL 2.
	//
	// 	- **CloudESSD_PL3**: The cluster uses an ESSD of PL 3.
	//
	// 	- **CloudEfficiency**: The cluster uses an ultra disk.
	//
	// example:
	//
	// CloudESSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Indicates whether data backup is supported. Valid values:
	//
	// 	- **1**: Data backup is supported.
	//
	// 	- **2**: Data backup is not supported.
	//
	// example:
	//
	// 1
	SupportBackup *int32 `json:"SupportBackup,omitempty" xml:"SupportBackup,omitempty"`
	// Indicates whether HTTPS ports are supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	SupportHttpsPort *bool `json:"SupportHttpsPort,omitempty" xml:"SupportHttpsPort,omitempty"`
	// Indicates whether the cluster supports a MySQL port. Valid values:
	//
	// 	- **true**: A MySQL port is supported.
	//
	// 	- **false**: A MySQL port is not supported.
	//
	// example:
	//
	// false
	SupportMysqlPort *bool `json:"SupportMysqlPort,omitempty" xml:"SupportMysqlPort,omitempty"`
	// Indicates whether tiered storage of hot data and cold data is supported. Valid values:
	//
	// 	- **1**: Tiered storage of hot data and cold data is supported.
	//
	// 	- **2**: Tiered storage of hot data and cold data is not supported.
	//
	// example:
	//
	// 1
	SupportOss *int32 `json:"SupportOss,omitempty" xml:"SupportOss,omitempty"`
	// The tags.
	Tags *DescribeDBClusterAttributeResponseBodyDBClusterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1n874li1t5y57wi****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which the cluster is deployed.
	//
	// example:
	//
	// vpc-bp10tr8k9qasioaty****
	VpcCloudInstanceId *string `json:"VpcCloudInstanceId,omitempty" xml:"VpcCloudInstanceId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp10tr8k9qasioaty****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IP address that is used to connect to the cluster over the VPC.
	//
	// example:
	//
	// 192.168.xx.xx
	VpcIpAddr *string `json:"VpcIpAddr,omitempty" xml:"VpcIpAddr,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The list of vSwitch IDs in multi-zone clusters.
	//
	// example:
	//
	// cn-shanghai-f: vsw-zm0n42d5vvuo****
	ZoneIdVswitchMap map[string]interface{} `json:"ZoneIdVswitchMap,omitempty" xml:"ZoneIdVswitchMap,omitempty"`
	// The ZooKeeper specifications.
	//
	// example:
	//
	// 4 Core 8 GB
	ZookeeperClass *string `json:"ZookeeperClass,omitempty" xml:"ZookeeperClass,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAppointmentElectZookeeperDisableWrite() *bool {
	return s.AppointmentElectZookeeperDisableWrite
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAppointmentElectZookeeperTime() *string {
	return s.AppointmentElectZookeeperTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAppointmentRestartNodeList() *string {
	return s.AppointmentRestartNodeList
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAppointmentRestartNodeTime() *string {
	return s.AppointmentRestartNodeTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAppointmentRestartTime() *string {
	return s.AppointmentRestartTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetAvailableUpgradeMajorVersion() map[string]interface{} {
	return s.AvailableUpgradeMajorVersion
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetBid() *string {
	return s.Bid
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetControlVersion() *string {
	return s.ControlVersion
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBClusterNetworkType() *string {
	return s.DBClusterNetworkType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBClusterType() *string {
	return s.DBClusterType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBNodeCount() *int64 {
	return s.DBNodeCount
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetDBNodeStorage() *int64 {
	return s.DBNodeStorage
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetEngineLatestMinorVersion() *string {
	return s.EngineLatestMinorVersion
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetEngineMinorVersion() *string {
	return s.EngineMinorVersion
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetExtStorageSize() *int32 {
	return s.ExtStorageSize
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetExtStorageType() *string {
	return s.ExtStorageType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetIsExpired() *string {
	return s.IsExpired
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetMaintainAutoType() *bool {
	return s.MaintainAutoType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetPublicConnectionString() *string {
	return s.PublicConnectionString
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetPublicIpAddr() *string {
	return s.PublicIpAddr
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetPublicPort() *string {
	return s.PublicPort
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetScaleOutStatus() *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus {
	return s.ScaleOutStatus
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetSupportBackup() *int32 {
	return s.SupportBackup
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetSupportHttpsPort() *bool {
	return s.SupportHttpsPort
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetSupportMysqlPort() *bool {
	return s.SupportMysqlPort
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetSupportOss() *int32 {
	return s.SupportOss
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetTags() *DescribeDBClusterAttributeResponseBodyDBClusterTags {
	return s.Tags
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetVpcCloudInstanceId() *string {
	return s.VpcCloudInstanceId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetVpcIpAddr() *string {
	return s.VpcIpAddr
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetZoneIdVswitchMap() map[string]interface{} {
	return s.ZoneIdVswitchMap
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) GetZookeeperClass() *string {
	return s.ZookeeperClass
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAliUid(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AliUid = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAppointmentElectZookeeperDisableWrite(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AppointmentElectZookeeperDisableWrite = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAppointmentElectZookeeperTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AppointmentElectZookeeperTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAppointmentRestartNodeList(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AppointmentRestartNodeList = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAppointmentRestartNodeTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AppointmentRestartNodeTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAppointmentRestartTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AppointmentRestartTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetAvailableUpgradeMajorVersion(v map[string]interface{}) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.AvailableUpgradeMajorVersion = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetBid(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Bid = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetCategory(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetCommodityCode(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetConnectionString(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetControlVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ControlVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetCreateTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBClusterType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBClusterType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBNodeCount(v int64) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetDBNodeStorage(v int64) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.DBNodeStorage = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEncryptionKey(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEncryptionType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngine(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngineLatestMinorVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EngineLatestMinorVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngineMinorVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EngineMinorVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetEngineVersion(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetExtStorageSize(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ExtStorageSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetExtStorageType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ExtStorageType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetIsExpired(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.IsExpired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetLockMode(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetLockReason(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.LockReason = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetMaintainAutoType(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.MaintainAutoType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPayType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPort(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPublicConnectionString(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PublicConnectionString = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPublicIpAddr(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PublicIpAddr = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetPublicPort(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.PublicPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetRegionId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetScaleOutStatus(v *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ScaleOutStatus = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetStorageType(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportBackup(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportBackup = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportHttpsPort(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportHttpsPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportMysqlPort(v bool) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportMysqlPort = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetSupportOss(v int32) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.SupportOss = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetTags(v *DescribeDBClusterAttributeResponseBodyDBClusterTags) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVpcCloudInstanceId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VpcCloudInstanceId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVpcId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetVpcIpAddr(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.VpcIpAddr = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZoneIdVswitchMap(v map[string]interface{}) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZoneIdVswitchMap = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) SetZookeeperClass(v string) *DescribeDBClusterAttributeResponseBodyDBCluster {
	s.ZookeeperClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBCluster) Validate() error {
	if s.ScaleOutStatus != nil {
		if err := s.ScaleOutStatus.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus struct {
	// The progress of the data migration task in percentage.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The progress of the data migration task. This value is displayed in the following format: Data volume that has been migrated/Total data volume.
	//
	// >  This parameter is returned only when the cluster is in the SCALING_OUT state.
	//
	// example:
	//
	// 0MB/60469MB
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) SetProgress(v string) *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) SetRatio(v string) *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus {
	s.Ratio = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterScaleOutStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterAttributeResponseBodyDBClusterTags struct {
	Tag []*DescribeDBClusterAttributeResponseBodyDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTags) GetTag() []*DescribeDBClusterAttributeResponseBodyDBClusterTagsTag {
	return s.Tag
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTags) SetTag(v []*DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) *DescribeDBClusterAttributeResponseBodyDBClusterTags {
	s.Tag = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterAttributeResponseBodyDBClusterTagsTag struct {
	// The tag name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// it
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) SetKey(v string) *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) SetValue(v string) *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBClusterTagsTag) Validate() error {
	return dara.Validate(s)
}
