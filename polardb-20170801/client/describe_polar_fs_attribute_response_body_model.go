// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateType(v string) *DescribePolarFsAttributeResponseBody
	GetAccelerateType() *string
	SetAcceleratedStorageSpace(v float64) *DescribePolarFsAttributeResponseBody
	GetAcceleratedStorageSpace() *float64
	SetAcceleratingEnable(v string) *DescribePolarFsAttributeResponseBody
	GetAcceleratingEnable() *string
	SetBandwidth(v float64) *DescribePolarFsAttributeResponseBody
	GetBandwidth() *float64
	SetBandwidthBaseLine(v float64) *DescribePolarFsAttributeResponseBody
	GetBandwidthBaseLine() *float64
	SetBucketId(v string) *DescribePolarFsAttributeResponseBody
	GetBucketId() *string
	SetCategory(v string) *DescribePolarFsAttributeResponseBody
	GetCategory() *string
	SetClientDownloadPath(v string) *DescribePolarFsAttributeResponseBody
	GetClientDownloadPath() *string
	SetCreateTime(v string) *DescribePolarFsAttributeResponseBody
	GetCreateTime() *string
	SetCustomBucketPath(v string) *DescribePolarFsAttributeResponseBody
	GetCustomBucketPath() *string
	SetCustomBucketPathList(v []*DescribePolarFsAttributeResponseBodyCustomBucketPathList) *DescribePolarFsAttributeResponseBody
	GetCustomBucketPathList() []*DescribePolarFsAttributeResponseBodyCustomBucketPathList
	SetDBEndpointId(v string) *DescribePolarFsAttributeResponseBody
	GetDBEndpointId() *string
	SetDBType(v string) *DescribePolarFsAttributeResponseBody
	GetDBType() *string
	SetEndpointItems(v []*DescribePolarFsAttributeResponseBodyEndpointItems) *DescribePolarFsAttributeResponseBody
	GetEndpointItems() []*DescribePolarFsAttributeResponseBodyEndpointItems
	SetExpireTime(v string) *DescribePolarFsAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v string) *DescribePolarFsAttributeResponseBody
	GetExpired() *string
	SetFileSystemId(v string) *DescribePolarFsAttributeResponseBody
	GetFileSystemId() *string
	SetLockMode(v string) *DescribePolarFsAttributeResponseBody
	GetLockMode() *string
	SetMaxscaleEndpointId(v string) *DescribePolarFsAttributeResponseBody
	GetMaxscaleEndpointId() *string
	SetMetaConnString(v string) *DescribePolarFsAttributeResponseBody
	GetMetaConnString() *string
	SetMetaInstanceName(v string) *DescribePolarFsAttributeResponseBody
	GetMetaInstanceName() *string
	SetMetaMxsConnString(v string) *DescribePolarFsAttributeResponseBody
	GetMetaMxsConnString() *string
	SetMetaUrl(v string) *DescribePolarFsAttributeResponseBody
	GetMetaUrl() *string
	SetMinorVersion(v string) *DescribePolarFsAttributeResponseBody
	GetMinorVersion() *string
	SetMountInfo(v *DescribePolarFsAttributeResponseBodyMountInfo) *DescribePolarFsAttributeResponseBody
	GetMountInfo() *DescribePolarFsAttributeResponseBodyMountInfo
	SetPayType(v string) *DescribePolarFsAttributeResponseBody
	GetPayType() *string
	SetPolarFsInstanceDescription(v string) *DescribePolarFsAttributeResponseBody
	GetPolarFsInstanceDescription() *string
	SetPolarFsInstanceId(v string) *DescribePolarFsAttributeResponseBody
	GetPolarFsInstanceId() *string
	SetPolarFsStatus(v string) *DescribePolarFsAttributeResponseBody
	GetPolarFsStatus() *string
	SetPolarFsType(v string) *DescribePolarFsAttributeResponseBody
	GetPolarFsType() *string
	SetPolarFsVersion(v string) *DescribePolarFsAttributeResponseBody
	GetPolarFsVersion() *string
	SetRegionId(v string) *DescribePolarFsAttributeResponseBody
	GetRegionId() *string
	SetRelativeDbClusterId(v string) *DescribePolarFsAttributeResponseBody
	GetRelativeDbClusterId() *string
	SetRelativePfsClusterId(v string) *DescribePolarFsAttributeResponseBody
	GetRelativePfsClusterId() *string
	SetRequestId(v string) *DescribePolarFsAttributeResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *DescribePolarFsAttributeResponseBody
	GetSecurityGroupId() *string
	SetStorageSpace(v float64) *DescribePolarFsAttributeResponseBody
	GetStorageSpace() *float64
	SetStorageType(v string) *DescribePolarFsAttributeResponseBody
	GetStorageType() *string
	SetStorageUsed(v float64) *DescribePolarFsAttributeResponseBody
	GetStorageUsed() *float64
	SetUserDefaultAccName(v string) *DescribePolarFsAttributeResponseBody
	GetUserDefaultAccName() *string
	SetUserDefaultAccSk(v string) *DescribePolarFsAttributeResponseBody
	GetUserDefaultAccSk() *string
	SetVPCId(v string) *DescribePolarFsAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribePolarFsAttributeResponseBody
	GetVSwitchId() *string
	SetZoneId(v string) *DescribePolarFsAttributeResponseBody
	GetZoneId() *string
}

type DescribePolarFsAttributeResponseBody struct {
	// The acceleration type.
	//
	// example:
	//
	// alluxio
	AccelerateType *string `json:"AccelerateType,omitempty" xml:"AccelerateType,omitempty"`
	// The acceleration storage space. Unit: GB.
	//
	// example:
	//
	// 1000
	AcceleratedStorageSpace *float64 `json:"AcceleratedStorageSpace,omitempty" xml:"AcceleratedStorageSpace,omitempty"`
	// Specifies whether the acceleration cache is enabled. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: Disabled.
	//
	// example:
	//
	// ON
	AcceleratingEnable *string `json:"AcceleratingEnable,omitempty" xml:"AcceleratingEnable,omitempty"`
	// The bandwidth. Unit: MB/s.
	//
	// example:
	//
	// 100
	Bandwidth *float64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth baseline. Unit: MB/s/TiB.
	//
	// example:
	//
	// 100
	BandwidthBaseLine *float64 `json:"BandwidthBaseLine,omitempty" xml:"BandwidthBaseLine,omitempty"`
	// The bucket ID.
	//
	// example:
	//
	// xxx
	BucketId *string `json:"BucketId,omitempty" xml:"BucketId,omitempty"`
	// The PolarLakebase edition. Valid values:
	//
	// - **high_performance**: High-performance Edition.
	//
	// - **basic**: Basic Edition.
	//
	// - **cold**: Cold Storage Edition.
	//
	// example:
	//
	// high_performance
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client download URL.
	//
	// example:
	//
	// oss://*
	ClientDownloadPath *string `json:"ClientDownloadPath,omitempty" xml:"ClientDownloadPath,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-08-02T05:57:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The custom bucket path.
	//
	// example:
	//
	// xxxxxx-%d.oss-cn-beijing-internal.aliyuncs.com
	CustomBucketPath *string `json:"CustomBucketPath,omitempty" xml:"CustomBucketPath,omitempty"`
	// The list of custom storage paths.
	CustomBucketPathList []*DescribePolarFsAttributeResponseBodyCustomBucketPathList `json:"CustomBucketPathList,omitempty" xml:"CustomBucketPathList,omitempty" type:"Repeated"`
	DBEndpointId         *string                                                     `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The database ecosystem type. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// example:
	//
	// MySQL
	DBType        *string                                              `json:"DBType,omitempty" xml:"DBType,omitempty"`
	EndpointItems []*DescribePolarFsAttributeResponseBodyEndpointItems `json:"EndpointItems,omitempty" xml:"EndpointItems,omitempty" type:"Repeated"`
	// The expiration time of the cluster.
	//
	// > This parameter is returned only for clusters that use the **Prepaid*	- (subscription) billing method. An empty value is returned for **Postpaid*	- (pay-as-you-go) clusters.
	//
	// example:
	//
	// 2025-10-10T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired.
	//
	// > This parameter is returned only for clusters that use the **Prepaid*	- (subscription) billing method.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The file system ID.
	//
	// example:
	//
	// xxx
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The lock mode. Valid values:
	//
	// - **Unlock**: Not locked.
	//
	// - **ManualLock**: Manually locked.
	//
	// - **LockByExpiration**: Automatically locked due to cluster expiration.
	//
	// example:
	//
	// Unlock
	LockMode           *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MaxscaleEndpointId *string `json:"MaxscaleEndpointId,omitempty" xml:"MaxscaleEndpointId,omitempty"`
	MetaConnString     *string `json:"MetaConnString,omitempty" xml:"MetaConnString,omitempty"`
	// example:
	//
	// pc-xxxxxxxxxxxxxxxxx
	MetaInstanceName  *string `json:"MetaInstanceName,omitempty" xml:"MetaInstanceName,omitempty"`
	MetaMxsConnString *string `json:"MetaMxsConnString,omitempty" xml:"MetaMxsConnString,omitempty"`
	// The encrypted metadata URL for Fuse mounting.
	//
	// example:
	//
	// e6cc1d2e2a6fa292038d999fda6501*****
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// The minor version of the instance.
	//
	// example:
	//
	// v1.3.0-v1.1.1
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The mount configuration.
	MountInfo *DescribePolarFsAttributeResponseBodyMountInfo `json:"MountInfo,omitempty" xml:"MountInfo,omitempty" type:"Struct"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The description of the PolarLakebase instance.
	//
	// example:
	//
	// pfs-xxx
	PolarFsInstanceDescription *string `json:"PolarFsInstanceDescription,omitempty" xml:"PolarFsInstanceDescription,omitempty"`
	// The PolarLakebase instance ID.
	//
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The PolarLakebase instance status.
	//
	// example:
	//
	// Running
	PolarFsStatus *string `json:"PolarFsStatus,omitempty" xml:"PolarFsStatus,omitempty"`
	// The instance version. Valid values:
	//
	// - **PolarFS 2.0**: 2.0.
	//
	// - **PolarFS 1.0**: 1.0.
	//
	// example:
	//
	// PolarFS 2.0
	PolarFsType *string `json:"PolarFsType,omitempty" xml:"PolarFsType,omitempty"`
	// The version.
	//
	// example:
	//
	// 1.0.1-1.0.3
	PolarFsVersion *string `json:"PolarFsVersion,omitempty" xml:"PolarFsVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the associated PolarDB cluster.
	//
	// example:
	//
	// pc-2zejpr41d9xk3uk34
	RelativeDbClusterId *string `json:"RelativeDbClusterId,omitempty" xml:"RelativeDbClusterId,omitempty"`
	// The instance ID of the associated PolarLakebase instance.
	//
	// example:
	//
	// pfs-**********
	RelativePfsClusterId *string `json:"RelativePfsClusterId,omitempty" xml:"RelativePfsClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F9E6A3B-C13E-4064-A010-18582A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The managed security group ID.
	//
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The storage space. Unit: GB.
	//
	// example:
	//
	// 1000
	StorageSpace *float64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage type for the High-performance Edition. Valid values:
	//
	// 	- **ESSDPL1**
	//
	// 	- **ESSDPL0**
	//
	// The storage type for the Basic Edition. Valid values:
	//
	// 	- **city_redundancy**: zone-redundant storage.
	//
	// example:
	//
	// essdpl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The used storage space. Unit: bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *float64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// lakebase_acc
	UserDefaultAccName *string `json:"UserDefaultAccName,omitempty" xml:"UserDefaultAccName,omitempty"`
	// example:
	//
	// EncryptedSecretKey==
	UserDefaultAccSk *string `json:"UserDefaultAccSk,omitempty" xml:"UserDefaultAccSk,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-**************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the vSwitch.
	//
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribePolarFsAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeResponseBody) GetAccelerateType() *string {
	return s.AccelerateType
}

func (s *DescribePolarFsAttributeResponseBody) GetAcceleratedStorageSpace() *float64 {
	return s.AcceleratedStorageSpace
}

func (s *DescribePolarFsAttributeResponseBody) GetAcceleratingEnable() *string {
	return s.AcceleratingEnable
}

func (s *DescribePolarFsAttributeResponseBody) GetBandwidth() *float64 {
	return s.Bandwidth
}

func (s *DescribePolarFsAttributeResponseBody) GetBandwidthBaseLine() *float64 {
	return s.BandwidthBaseLine
}

func (s *DescribePolarFsAttributeResponseBody) GetBucketId() *string {
	return s.BucketId
}

func (s *DescribePolarFsAttributeResponseBody) GetCategory() *string {
	return s.Category
}

func (s *DescribePolarFsAttributeResponseBody) GetClientDownloadPath() *string {
	return s.ClientDownloadPath
}

func (s *DescribePolarFsAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePolarFsAttributeResponseBody) GetCustomBucketPath() *string {
	return s.CustomBucketPath
}

func (s *DescribePolarFsAttributeResponseBody) GetCustomBucketPathList() []*DescribePolarFsAttributeResponseBodyCustomBucketPathList {
	return s.CustomBucketPathList
}

func (s *DescribePolarFsAttributeResponseBody) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribePolarFsAttributeResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribePolarFsAttributeResponseBody) GetEndpointItems() []*DescribePolarFsAttributeResponseBodyEndpointItems {
	return s.EndpointItems
}

func (s *DescribePolarFsAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribePolarFsAttributeResponseBody) GetExpired() *string {
	return s.Expired
}

func (s *DescribePolarFsAttributeResponseBody) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribePolarFsAttributeResponseBody) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribePolarFsAttributeResponseBody) GetMaxscaleEndpointId() *string {
	return s.MaxscaleEndpointId
}

func (s *DescribePolarFsAttributeResponseBody) GetMetaConnString() *string {
	return s.MetaConnString
}

func (s *DescribePolarFsAttributeResponseBody) GetMetaInstanceName() *string {
	return s.MetaInstanceName
}

func (s *DescribePolarFsAttributeResponseBody) GetMetaMxsConnString() *string {
	return s.MetaMxsConnString
}

func (s *DescribePolarFsAttributeResponseBody) GetMetaUrl() *string {
	return s.MetaUrl
}

func (s *DescribePolarFsAttributeResponseBody) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribePolarFsAttributeResponseBody) GetMountInfo() *DescribePolarFsAttributeResponseBodyMountInfo {
	return s.MountInfo
}

func (s *DescribePolarFsAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribePolarFsAttributeResponseBody) GetPolarFsInstanceDescription() *string {
	return s.PolarFsInstanceDescription
}

func (s *DescribePolarFsAttributeResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribePolarFsAttributeResponseBody) GetPolarFsStatus() *string {
	return s.PolarFsStatus
}

func (s *DescribePolarFsAttributeResponseBody) GetPolarFsType() *string {
	return s.PolarFsType
}

func (s *DescribePolarFsAttributeResponseBody) GetPolarFsVersion() *string {
	return s.PolarFsVersion
}

func (s *DescribePolarFsAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePolarFsAttributeResponseBody) GetRelativeDbClusterId() *string {
	return s.RelativeDbClusterId
}

func (s *DescribePolarFsAttributeResponseBody) GetRelativePfsClusterId() *string {
	return s.RelativePfsClusterId
}

func (s *DescribePolarFsAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarFsAttributeResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribePolarFsAttributeResponseBody) GetStorageSpace() *float64 {
	return s.StorageSpace
}

func (s *DescribePolarFsAttributeResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribePolarFsAttributeResponseBody) GetStorageUsed() *float64 {
	return s.StorageUsed
}

func (s *DescribePolarFsAttributeResponseBody) GetUserDefaultAccName() *string {
	return s.UserDefaultAccName
}

func (s *DescribePolarFsAttributeResponseBody) GetUserDefaultAccSk() *string {
	return s.UserDefaultAccSk
}

func (s *DescribePolarFsAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePolarFsAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePolarFsAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribePolarFsAttributeResponseBody) SetAccelerateType(v string) *DescribePolarFsAttributeResponseBody {
	s.AccelerateType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetAcceleratedStorageSpace(v float64) *DescribePolarFsAttributeResponseBody {
	s.AcceleratedStorageSpace = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetAcceleratingEnable(v string) *DescribePolarFsAttributeResponseBody {
	s.AcceleratingEnable = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetBandwidth(v float64) *DescribePolarFsAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetBandwidthBaseLine(v float64) *DescribePolarFsAttributeResponseBody {
	s.BandwidthBaseLine = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetBucketId(v string) *DescribePolarFsAttributeResponseBody {
	s.BucketId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetCategory(v string) *DescribePolarFsAttributeResponseBody {
	s.Category = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetClientDownloadPath(v string) *DescribePolarFsAttributeResponseBody {
	s.ClientDownloadPath = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetCreateTime(v string) *DescribePolarFsAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetCustomBucketPath(v string) *DescribePolarFsAttributeResponseBody {
	s.CustomBucketPath = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetCustomBucketPathList(v []*DescribePolarFsAttributeResponseBodyCustomBucketPathList) *DescribePolarFsAttributeResponseBody {
	s.CustomBucketPathList = v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetDBEndpointId(v string) *DescribePolarFsAttributeResponseBody {
	s.DBEndpointId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetDBType(v string) *DescribePolarFsAttributeResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetEndpointItems(v []*DescribePolarFsAttributeResponseBodyEndpointItems) *DescribePolarFsAttributeResponseBody {
	s.EndpointItems = v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetExpireTime(v string) *DescribePolarFsAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetExpired(v string) *DescribePolarFsAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetFileSystemId(v string) *DescribePolarFsAttributeResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetLockMode(v string) *DescribePolarFsAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMaxscaleEndpointId(v string) *DescribePolarFsAttributeResponseBody {
	s.MaxscaleEndpointId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMetaConnString(v string) *DescribePolarFsAttributeResponseBody {
	s.MetaConnString = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMetaInstanceName(v string) *DescribePolarFsAttributeResponseBody {
	s.MetaInstanceName = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMetaMxsConnString(v string) *DescribePolarFsAttributeResponseBody {
	s.MetaMxsConnString = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMetaUrl(v string) *DescribePolarFsAttributeResponseBody {
	s.MetaUrl = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMinorVersion(v string) *DescribePolarFsAttributeResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetMountInfo(v *DescribePolarFsAttributeResponseBodyMountInfo) *DescribePolarFsAttributeResponseBody {
	s.MountInfo = v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetPayType(v string) *DescribePolarFsAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetPolarFsInstanceDescription(v string) *DescribePolarFsAttributeResponseBody {
	s.PolarFsInstanceDescription = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetPolarFsInstanceId(v string) *DescribePolarFsAttributeResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetPolarFsStatus(v string) *DescribePolarFsAttributeResponseBody {
	s.PolarFsStatus = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetPolarFsType(v string) *DescribePolarFsAttributeResponseBody {
	s.PolarFsType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetPolarFsVersion(v string) *DescribePolarFsAttributeResponseBody {
	s.PolarFsVersion = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetRegionId(v string) *DescribePolarFsAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetRelativeDbClusterId(v string) *DescribePolarFsAttributeResponseBody {
	s.RelativeDbClusterId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetRelativePfsClusterId(v string) *DescribePolarFsAttributeResponseBody {
	s.RelativePfsClusterId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetRequestId(v string) *DescribePolarFsAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetSecurityGroupId(v string) *DescribePolarFsAttributeResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetStorageSpace(v float64) *DescribePolarFsAttributeResponseBody {
	s.StorageSpace = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetStorageType(v string) *DescribePolarFsAttributeResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetStorageUsed(v float64) *DescribePolarFsAttributeResponseBody {
	s.StorageUsed = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetUserDefaultAccName(v string) *DescribePolarFsAttributeResponseBody {
	s.UserDefaultAccName = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetUserDefaultAccSk(v string) *DescribePolarFsAttributeResponseBody {
	s.UserDefaultAccSk = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetVPCId(v string) *DescribePolarFsAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetVSwitchId(v string) *DescribePolarFsAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) SetZoneId(v string) *DescribePolarFsAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBody) Validate() error {
	if s.CustomBucketPathList != nil {
		for _, item := range s.CustomBucketPathList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EndpointItems != nil {
		for _, item := range s.EndpointItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MountInfo != nil {
		if err := s.MountInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarFsAttributeResponseBodyCustomBucketPathList struct {
	// The custom storage bucket.
	//
	// example:
	//
	// pfs-xxx.oss-[regionId]-internal.aliyuncs.com
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The custom storage path.
	//
	// example:
	//
	// /data
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribePolarFsAttributeResponseBodyCustomBucketPathList) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeResponseBodyCustomBucketPathList) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeResponseBodyCustomBucketPathList) GetBucket() *string {
	return s.Bucket
}

func (s *DescribePolarFsAttributeResponseBodyCustomBucketPathList) GetPath() *string {
	return s.Path
}

func (s *DescribePolarFsAttributeResponseBodyCustomBucketPathList) SetBucket(v string) *DescribePolarFsAttributeResponseBodyCustomBucketPathList {
	s.Bucket = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyCustomBucketPathList) SetPath(v string) *DescribePolarFsAttributeResponseBodyCustomBucketPathList {
	s.Path = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyCustomBucketPathList) Validate() error {
	return dara.Validate(s)
}

type DescribePolarFsAttributeResponseBodyEndpointItems struct {
	AddressItems []*DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// example:
	//
	// ep-xxxxxxxxx
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// example:
	//
	// S3Gateway
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s DescribePolarFsAttributeResponseBodyEndpointItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeResponseBodyEndpointItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) GetAddressItems() []*DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	return s.AddressItems
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) SetAddressItems(v []*DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) *DescribePolarFsAttributeResponseBodyEndpointItems {
	s.AddressItems = v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) SetDBEndpointId(v string) *DescribePolarFsAttributeResponseBodyEndpointItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) SetEndpointType(v string) *DescribePolarFsAttributeResponseBodyEndpointItems {
	s.EndpointType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItems) Validate() error {
	if s.AddressItems != nil {
		for _, item := range s.AddressItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GetPort() *string {
	return s.Port
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) SetConnectionString(v string) *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) SetIPAddress(v string) *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) SetNetType(v string) *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) SetPort(v string) *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	s.Port = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) SetVPCId(v string) *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) SetVSwitchId(v string) *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyEndpointItemsAddressItems) Validate() error {
	return dara.Validate(s)
}

type DescribePolarFsAttributeResponseBodyMountInfo struct {
	// The cluster management endpoint.
	//
	// example:
	//
	// TCP://**.**.**.**:3000,TCP://**.**.**.**:3000,TCP://**.**.**.**:3000
	PolarDbProxy *string `json:"PolarDbProxy,omitempty" xml:"PolarDbProxy,omitempty"`
	// The file system name.
	//
	// example:
	//
	// pfs-**********
	PolarFsCluster *string `json:"PolarFsCluster,omitempty" xml:"PolarFsCluster,omitempty"`
	// The token value.
	//
	// example:
	//
	// a734298c391cb9ebd05e2ee85feb624
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribePolarFsAttributeResponseBodyMountInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarFsAttributeResponseBodyMountInfo) GoString() string {
	return s.String()
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) GetPolarDbProxy() *string {
	return s.PolarDbProxy
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) GetPolarFsCluster() *string {
	return s.PolarFsCluster
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) GetToken() *string {
	return s.Token
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) SetPolarDbProxy(v string) *DescribePolarFsAttributeResponseBodyMountInfo {
	s.PolarDbProxy = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) SetPolarFsCluster(v string) *DescribePolarFsAttributeResponseBodyMountInfo {
	s.PolarFsCluster = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) SetToken(v string) *DescribePolarFsAttributeResponseBodyMountInfo {
	s.Token = &v
	return s
}

func (s *DescribePolarFsAttributeResponseBodyMountInfo) Validate() error {
	return dara.Validate(s)
}
