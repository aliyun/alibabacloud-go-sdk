// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarFsAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
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
	SetDBType(v string) *DescribePolarFsAttributeResponseBody
	GetDBType() *string
	SetExpireTime(v string) *DescribePolarFsAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v string) *DescribePolarFsAttributeResponseBody
	GetExpired() *string
	SetFileSystemId(v string) *DescribePolarFsAttributeResponseBody
	GetFileSystemId() *string
	SetLockMode(v string) *DescribePolarFsAttributeResponseBody
	GetLockMode() *string
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
	SetVPCId(v string) *DescribePolarFsAttributeResponseBody
	GetVPCId() *string
	SetVSwitchId(v string) *DescribePolarFsAttributeResponseBody
	GetVSwitchId() *string
	SetZoneId(v string) *DescribePolarFsAttributeResponseBody
	GetZoneId() *string
}

type DescribePolarFsAttributeResponseBody struct {
	// example:
	//
	// 1000
	AcceleratedStorageSpace *float64 `json:"AcceleratedStorageSpace,omitempty" xml:"AcceleratedStorageSpace,omitempty"`
	// example:
	//
	// ON
	AcceleratingEnable *string `json:"AcceleratingEnable,omitempty" xml:"AcceleratingEnable,omitempty"`
	// example:
	//
	// 100
	Bandwidth *float64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// 100
	BandwidthBaseLine *float64 `json:"BandwidthBaseLine,omitempty" xml:"BandwidthBaseLine,omitempty"`
	BucketId          *string  `json:"BucketId,omitempty" xml:"BucketId,omitempty"`
	// example:
	//
	// high_performance
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientDownloadPath *string `json:"ClientDownloadPath,omitempty" xml:"ClientDownloadPath,omitempty"`
	// example:
	//
	// 2021-08-02T05:57:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 2025-10-10T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// false
	Expired      *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// Unlock
	LockMode     *string                                        `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	MetaUrl      *string                                        `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	MinorVersion *string                                        `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	MountInfo    *DescribePolarFsAttributeResponseBodyMountInfo `json:"MountInfo,omitempty" xml:"MountInfo,omitempty" type:"Struct"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// pfs-xxx
	PolarFsInstanceDescription *string `json:"PolarFsInstanceDescription,omitempty" xml:"PolarFsInstanceDescription,omitempty"`
	// example:
	//
	// pfs-2ze0i74ka607*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// example:
	//
	// Running
	PolarFsStatus *string `json:"PolarFsStatus,omitempty" xml:"PolarFsStatus,omitempty"`
	// example:
	//
	// PolarFS 2.0
	PolarFsType *string `json:"PolarFsType,omitempty" xml:"PolarFsType,omitempty"`
	// example:
	//
	// 1.0.1-1.0.3
	PolarFsVersion *string `json:"PolarFsVersion,omitempty" xml:"PolarFsVersion,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// pc-2zejpr41d9xk3uk34
	RelativeDbClusterId  *string `json:"RelativeDbClusterId,omitempty" xml:"RelativeDbClusterId,omitempty"`
	RelativePfsClusterId *string `json:"RelativePfsClusterId,omitempty" xml:"RelativePfsClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F9E6A3B-C13E-4064-A010-18582A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// 1000
	StorageSpace *float64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// example:
	//
	// essdpl1
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// example:
	//
	// 3012558848
	StorageUsed *float64 `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-**************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *DescribePolarFsAttributeResponseBody) GetDBType() *string {
	return s.DBType
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

func (s *DescribePolarFsAttributeResponseBody) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribePolarFsAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribePolarFsAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
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

func (s *DescribePolarFsAttributeResponseBody) SetDBType(v string) *DescribePolarFsAttributeResponseBody {
	s.DBType = &v
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
	if s.MountInfo != nil {
		if err := s.MountInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarFsAttributeResponseBodyMountInfo struct {
	// example:
	//
	// TCP://**.**.**.**:3000,TCP://**.**.**.**:3000,TCP://**.**.**.**:3000
	PolarDbProxy *string `json:"PolarDbProxy,omitempty" xml:"PolarDbProxy,omitempty"`
	// example:
	//
	// pfs-**********
	PolarFsCluster *string `json:"PolarFsCluster,omitempty" xml:"PolarFsCluster,omitempty"`
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
