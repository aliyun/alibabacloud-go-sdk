// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarFsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateStorageSize(v int64) *CreatePolarFsRequest
	GetAccelerateStorageSize() *int64
	SetAccelerateSwitch(v string) *CreatePolarFsRequest
	GetAccelerateSwitch() *string
	SetAccelerateType(v string) *CreatePolarFsRequest
	GetAccelerateType() *string
	SetAuthorizedUserIds(v string) *CreatePolarFsRequest
	GetAuthorizedUserIds() *string
	SetAutoRenew(v bool) *CreatePolarFsRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v bool) *CreatePolarFsRequest
	GetAutoUseCoupon() *bool
	SetCreationCategory(v string) *CreatePolarFsRequest
	GetCreationCategory() *string
	SetCustomBucketCount(v int32) *CreatePolarFsRequest
	GetCustomBucketCount() *int32
	SetCustomBucketPath(v string) *CreatePolarFsRequest
	GetCustomBucketPath() *string
	SetCustomBucketPathList(v []*CreatePolarFsRequestCustomBucketPathList) *CreatePolarFsRequest
	GetCustomBucketPathList() []*CreatePolarFsRequestCustomBucketPathList
	SetCustomOssAk(v string) *CreatePolarFsRequest
	GetCustomOssAk() *string
	SetCustomOssSk(v string) *CreatePolarFsRequest
	GetCustomOssSk() *string
	SetDBClusterId(v string) *CreatePolarFsRequest
	GetDBClusterId() *string
	SetDBType(v string) *CreatePolarFsRequest
	GetDBType() *string
	SetPayType(v string) *CreatePolarFsRequest
	GetPayType() *string
	SetPeriod(v string) *CreatePolarFsRequest
	GetPeriod() *string
	SetPromotionCode(v string) *CreatePolarFsRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreatePolarFsRequest
	GetRegionId() *string
	SetStorageSpace(v int64) *CreatePolarFsRequest
	GetStorageSpace() *int64
	SetStorageType(v string) *CreatePolarFsRequest
	GetStorageType() *string
	SetUsedTime(v string) *CreatePolarFsRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreatePolarFsRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreatePolarFsRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreatePolarFsRequest
	GetZoneId() *string
}

type CreatePolarFsRequest struct {
	// The acceleration storage space for Basic Edition with acceleration enabled. Unit: GB.
	//
	// example:
	//
	// 500
	AccelerateStorageSize *int64 `json:"AccelerateStorageSize,omitempty" xml:"AccelerateStorageSize,omitempty"`
	// The acceleration mode. Valid values:
	//
	// - **ONLY**: enables acceleration only.
	//
	// - **ON**: enables cold data storage and acceleration.
	//
	// example:
	//
	// ONLY
	AccelerateSwitch *string `json:"AccelerateSwitch,omitempty" xml:"AccelerateSwitch,omitempty"`
	// The acceleration type. Valid values:
	//
	// - **juice**: file system acceleration.
	//
	// - **alluxio**: transparent acceleration.
	//
	// example:
	//
	// alluxio
	AccelerateType *string `json:"AccelerateType,omitempty" xml:"AccelerateType,omitempty"`
	// The list of authorized account IDs for Cold Storage Edition instances, separated by commas (,).
	//
	// example:
	//
	// 128***********，198***********
	AuthorizedUserIds *string `json:"AuthorizedUserIds,omitempty" xml:"AuthorizedUserIds,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is disabled.
	//
	// Default value: **false**.
	//
	// > This parameter takes effect only when **PayType*	- is set to **Prepaid**.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to automatically use coupons. Valid values:
	//
	// - **true**: Coupons are used (default).
	//
	// - **false**: Coupons are not used.
	//
	// example:
	//
	// true
	AutoUseCoupon *bool `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The edition. Valid values:
	//
	// - **basic**: Basic Edition (default).
	//
	// - **cold**: Cold Storage Edition.
	//
	// - **high_performance**: High-performance Edition.
	//
	// example:
	//
	// basic
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// The number of buckets.
	//
	// > This parameter is required only when acceleration (file system acceleration) is enabled.
	//
	// example:
	//
	// 1
	CustomBucketCount *int32 `json:"CustomBucketCount,omitempty" xml:"CustomBucketCount,omitempty"`
	// The bucket path.
	//
	// > This parameter is required only when acceleration (file system acceleration) is enabled.
	//
	// example:
	//
	// /test
	CustomBucketPath *string `json:"CustomBucketPath,omitempty" xml:"CustomBucketPath,omitempty"`
	// The bucket and path information.
	//
	// > This parameter is required for transparent acceleration scenarios.
	CustomBucketPathList []*CreatePolarFsRequestCustomBucketPathList `json:"CustomBucketPathList,omitempty" xml:"CustomBucketPathList,omitempty" type:"Repeated"`
	// The custom AccessKey ID.
	//
	// example:
	//
	// xxx
	CustomOssAk *string `json:"CustomOssAk,omitempty" xml:"CustomOssAk,omitempty"`
	// The custom AccessKey secret.
	//
	// example:
	//
	// xxx
	CustomOssSk *string `json:"CustomOssSk,omitempty" xml:"CustomOssSk,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine. Valid values:
	//
	// - **MySQL**
	//
	// - **PostgreSQL**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required when **PayType*	- is set to **Prepaid**. Specifies whether the subscription cluster uses a yearly or monthly billing cycle. You must pass this parameter when the billing method is subscription.
	//
	// - **Year**: The subscription period is measured in years.
	//
	// - **Month**: The subscription period is measured in months.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The coupon code. If this parameter is not specified, the default coupon is used.
	//
	// example:
	//
	// 727xxxxxx934
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID.
	//
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The storage space. Unit: GB.
	//
	// example:
	//
	// 50
	StorageSpace *int64 `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// Valid values for high-performance storage type:
	//
	// - **ESSDPL0**
	//
	// - **ESSDPL1**
	//
	// Valid values for Basic Edition storage type:
	//
	// - **city_redundancy (zone-redundant)**
	//
	// Valid values for Cold Storage Edition storage type:
	//
	// - **city_redundancy (zone-redundant)**
	//
	// - **local_redundancy (locally redundant)**
	//
	// example:
	//
	// local_redundancy
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// This parameter is required when **PayType*	- is set to **Prepaid**.
	//
	// - When **Period*	- is set to **Month**, the valid values of **UsedTime*	- are integers in the range of `[1-9]`.
	//
	// - When **Period*	- is set to **Year**, the valid values of **UsedTime*	- are integers in the range of `[1-3]`.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreatePolarFsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarFsRequest) GetAccelerateStorageSize() *int64 {
	return s.AccelerateStorageSize
}

func (s *CreatePolarFsRequest) GetAccelerateSwitch() *string {
	return s.AccelerateSwitch
}

func (s *CreatePolarFsRequest) GetAccelerateType() *string {
	return s.AccelerateType
}

func (s *CreatePolarFsRequest) GetAuthorizedUserIds() *string {
	return s.AuthorizedUserIds
}

func (s *CreatePolarFsRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreatePolarFsRequest) GetAutoUseCoupon() *bool {
	return s.AutoUseCoupon
}

func (s *CreatePolarFsRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *CreatePolarFsRequest) GetCustomBucketCount() *int32 {
	return s.CustomBucketCount
}

func (s *CreatePolarFsRequest) GetCustomBucketPath() *string {
	return s.CustomBucketPath
}

func (s *CreatePolarFsRequest) GetCustomBucketPathList() []*CreatePolarFsRequestCustomBucketPathList {
	return s.CustomBucketPathList
}

func (s *CreatePolarFsRequest) GetCustomOssAk() *string {
	return s.CustomOssAk
}

func (s *CreatePolarFsRequest) GetCustomOssSk() *string {
	return s.CustomOssSk
}

func (s *CreatePolarFsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreatePolarFsRequest) GetDBType() *string {
	return s.DBType
}

func (s *CreatePolarFsRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreatePolarFsRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreatePolarFsRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreatePolarFsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePolarFsRequest) GetStorageSpace() *int64 {
	return s.StorageSpace
}

func (s *CreatePolarFsRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreatePolarFsRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreatePolarFsRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreatePolarFsRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreatePolarFsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreatePolarFsRequest) SetAccelerateStorageSize(v int64) *CreatePolarFsRequest {
	s.AccelerateStorageSize = &v
	return s
}

func (s *CreatePolarFsRequest) SetAccelerateSwitch(v string) *CreatePolarFsRequest {
	s.AccelerateSwitch = &v
	return s
}

func (s *CreatePolarFsRequest) SetAccelerateType(v string) *CreatePolarFsRequest {
	s.AccelerateType = &v
	return s
}

func (s *CreatePolarFsRequest) SetAuthorizedUserIds(v string) *CreatePolarFsRequest {
	s.AuthorizedUserIds = &v
	return s
}

func (s *CreatePolarFsRequest) SetAutoRenew(v bool) *CreatePolarFsRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreatePolarFsRequest) SetAutoUseCoupon(v bool) *CreatePolarFsRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *CreatePolarFsRequest) SetCreationCategory(v string) *CreatePolarFsRequest {
	s.CreationCategory = &v
	return s
}

func (s *CreatePolarFsRequest) SetCustomBucketCount(v int32) *CreatePolarFsRequest {
	s.CustomBucketCount = &v
	return s
}

func (s *CreatePolarFsRequest) SetCustomBucketPath(v string) *CreatePolarFsRequest {
	s.CustomBucketPath = &v
	return s
}

func (s *CreatePolarFsRequest) SetCustomBucketPathList(v []*CreatePolarFsRequestCustomBucketPathList) *CreatePolarFsRequest {
	s.CustomBucketPathList = v
	return s
}

func (s *CreatePolarFsRequest) SetCustomOssAk(v string) *CreatePolarFsRequest {
	s.CustomOssAk = &v
	return s
}

func (s *CreatePolarFsRequest) SetCustomOssSk(v string) *CreatePolarFsRequest {
	s.CustomOssSk = &v
	return s
}

func (s *CreatePolarFsRequest) SetDBClusterId(v string) *CreatePolarFsRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreatePolarFsRequest) SetDBType(v string) *CreatePolarFsRequest {
	s.DBType = &v
	return s
}

func (s *CreatePolarFsRequest) SetPayType(v string) *CreatePolarFsRequest {
	s.PayType = &v
	return s
}

func (s *CreatePolarFsRequest) SetPeriod(v string) *CreatePolarFsRequest {
	s.Period = &v
	return s
}

func (s *CreatePolarFsRequest) SetPromotionCode(v string) *CreatePolarFsRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreatePolarFsRequest) SetRegionId(v string) *CreatePolarFsRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolarFsRequest) SetStorageSpace(v int64) *CreatePolarFsRequest {
	s.StorageSpace = &v
	return s
}

func (s *CreatePolarFsRequest) SetStorageType(v string) *CreatePolarFsRequest {
	s.StorageType = &v
	return s
}

func (s *CreatePolarFsRequest) SetUsedTime(v string) *CreatePolarFsRequest {
	s.UsedTime = &v
	return s
}

func (s *CreatePolarFsRequest) SetVPCId(v string) *CreatePolarFsRequest {
	s.VPCId = &v
	return s
}

func (s *CreatePolarFsRequest) SetVSwitchId(v string) *CreatePolarFsRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreatePolarFsRequest) SetZoneId(v string) *CreatePolarFsRequest {
	s.ZoneId = &v
	return s
}

func (s *CreatePolarFsRequest) Validate() error {
	if s.CustomBucketPathList != nil {
		for _, item := range s.CustomBucketPathList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePolarFsRequestCustomBucketPathList struct {
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

func (s CreatePolarFsRequestCustomBucketPathList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarFsRequestCustomBucketPathList) GoString() string {
	return s.String()
}

func (s *CreatePolarFsRequestCustomBucketPathList) GetBucket() *string {
	return s.Bucket
}

func (s *CreatePolarFsRequestCustomBucketPathList) GetPath() *string {
	return s.Path
}

func (s *CreatePolarFsRequestCustomBucketPathList) SetBucket(v string) *CreatePolarFsRequestCustomBucketPathList {
	s.Bucket = &v
	return s
}

func (s *CreatePolarFsRequestCustomBucketPathList) SetPath(v string) *CreatePolarFsRequestCustomBucketPathList {
	s.Path = &v
	return s
}

func (s *CreatePolarFsRequestCustomBucketPathList) Validate() error {
	return dara.Validate(s)
}
