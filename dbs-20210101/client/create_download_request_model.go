// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdminDatabase(v string) *CreateDownloadRequest
	GetAdminDatabase() *string
	SetBakSetId(v string) *CreateDownloadRequest
	GetBakSetId() *string
	SetBakSetSize(v string) *CreateDownloadRequest
	GetBakSetSize() *string
	SetBakSetType(v string) *CreateDownloadRequest
	GetBakSetType() *string
	SetClusterName(v string) *CreateDownloadRequest
	GetClusterName() *string
	SetDownloadPointInTime(v string) *CreateDownloadRequest
	GetDownloadPointInTime() *string
	SetFormatType(v string) *CreateDownloadRequest
	GetFormatType() *string
	SetInstanceName(v string) *CreateDownloadRequest
	GetInstanceName() *string
	SetIsCluster(v string) *CreateDownloadRequest
	GetIsCluster() *string
	SetIsPhysical(v bool) *CreateDownloadRequest
	GetIsPhysical() *bool
	SetPrimaryKeyTypeOnly(v string) *CreateDownloadRequest
	GetPrimaryKeyTypeOnly() *string
	SetRegionCode(v string) *CreateDownloadRequest
	GetRegionCode() *string
	SetTargetBucket(v string) *CreateDownloadRequest
	GetTargetBucket() *string
	SetTargetOssRegion(v string) *CreateDownloadRequest
	GetTargetOssRegion() *string
	SetTargetPath(v string) *CreateDownloadRequest
	GetTargetPath() *string
	SetTargetType(v string) *CreateDownloadRequest
	GetTargetType() *string
	SetUseZstd(v string) *CreateDownloadRequest
	GetUseZstd() *string
}

type CreateDownloadRequest struct {
	AdminDatabase *string `json:"AdminDatabase,omitempty" xml:"AdminDatabase,omitempty"`
	// The ID of the backup set. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to query the ID of the backup set.
	//
	// > This parameter is required if the BakSetType parameter is set to full.
	//
	// example:
	//
	// 146005****
	BakSetId *string `json:"BakSetId,omitempty" xml:"BakSetId,omitempty"`
	// The size of the full backup set. Unit: bytes. You can call the [DescribeBackups](https://help.aliyun.com/document_detail/26273.html) operation to query the size of the full backup set.
	//
	// example:
	//
	// 216****
	BakSetSize *string `json:"BakSetSize,omitempty" xml:"BakSetSize,omitempty"`
	// The type of the download task. Valid values:
	//
	// 	- **full**: downloads a full backup set.
	//
	// 	- **pitr**: downloads a backup set at a specific point in time.
	//
	// example:
	//
	// full
	BakSetType  *string `json:"BakSetType,omitempty" xml:"BakSetType,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The point in time at which the backup set is downloaded. Specify a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > This parameter is required if the BakSetType parameter is set to pitr.
	//
	// example:
	//
	// 1661331864000
	DownloadPointInTime *string `json:"DownloadPointInTime,omitempty" xml:"DownloadPointInTime,omitempty"`
	// The format to which the downloaded backup set is converted. Valid values:
	//
	// 	- **CSV**
	//
	// 	- **SQL**
	//
	// 	- **Parquet**
	//
	// > This parameter is required.
	//
	// example:
	//
	// CSV
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-wz994c1t1****
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsCluster          *string `json:"IsCluster,omitempty" xml:"IsCluster,omitempty"`
	IsPhysical         *bool   `json:"IsPhysical,omitempty" xml:"IsPhysical,omitempty"`
	PrimaryKeyTypeOnly *string `json:"PrimaryKeyTypeOnly,omitempty" xml:"PrimaryKeyTypeOnly,omitempty"`
	// The ID of the region in which the instance resides. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/26231.html) operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The name of the OSS bucket that is used to store the backup set.
	//
	// 	- This parameter is required if the TargetType parameter is set to OSS.
	//
	// 	- Make sure that your account is granted the **AliyunDBSDefaultRole*	- permission. For more information, see [Use RAM for resource authorization](https://help.aliyun.com/document_detail/26307.html). You can also grant permissions based on the operation instructions in the Resource Access Management (RAM) console.
	//
	// example:
	//
	// test123
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// The region in which the OSS bucket resides.
	//
	// > This parameter is required if the TargetType parameter is set to OSS.
	//
	// example:
	//
	// cn-beijing
	TargetOssRegion *string `json:"TargetOssRegion,omitempty" xml:"TargetOssRegion,omitempty"`
	// The destination path to which the backup set is downloaded.
	//
	// > This parameter is required if the TargetType parameter is set to OSS.
	//
	// example:
	//
	// test_db/path
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination to which the backup set is downloaded. Valid values:
	//
	// 	- **OSS**
	//
	// 	- **URL**
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UseZstd    *string `json:"UseZstd,omitempty" xml:"UseZstd,omitempty"`
}

func (s CreateDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDownloadRequest) GoString() string {
	return s.String()
}

func (s *CreateDownloadRequest) GetAdminDatabase() *string {
	return s.AdminDatabase
}

func (s *CreateDownloadRequest) GetBakSetId() *string {
	return s.BakSetId
}

func (s *CreateDownloadRequest) GetBakSetSize() *string {
	return s.BakSetSize
}

func (s *CreateDownloadRequest) GetBakSetType() *string {
	return s.BakSetType
}

func (s *CreateDownloadRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateDownloadRequest) GetDownloadPointInTime() *string {
	return s.DownloadPointInTime
}

func (s *CreateDownloadRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *CreateDownloadRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateDownloadRequest) GetIsCluster() *string {
	return s.IsCluster
}

func (s *CreateDownloadRequest) GetIsPhysical() *bool {
	return s.IsPhysical
}

func (s *CreateDownloadRequest) GetPrimaryKeyTypeOnly() *string {
	return s.PrimaryKeyTypeOnly
}

func (s *CreateDownloadRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *CreateDownloadRequest) GetTargetBucket() *string {
	return s.TargetBucket
}

func (s *CreateDownloadRequest) GetTargetOssRegion() *string {
	return s.TargetOssRegion
}

func (s *CreateDownloadRequest) GetTargetPath() *string {
	return s.TargetPath
}

func (s *CreateDownloadRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateDownloadRequest) GetUseZstd() *string {
	return s.UseZstd
}

func (s *CreateDownloadRequest) SetAdminDatabase(v string) *CreateDownloadRequest {
	s.AdminDatabase = &v
	return s
}

func (s *CreateDownloadRequest) SetBakSetId(v string) *CreateDownloadRequest {
	s.BakSetId = &v
	return s
}

func (s *CreateDownloadRequest) SetBakSetSize(v string) *CreateDownloadRequest {
	s.BakSetSize = &v
	return s
}

func (s *CreateDownloadRequest) SetBakSetType(v string) *CreateDownloadRequest {
	s.BakSetType = &v
	return s
}

func (s *CreateDownloadRequest) SetClusterName(v string) *CreateDownloadRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateDownloadRequest) SetDownloadPointInTime(v string) *CreateDownloadRequest {
	s.DownloadPointInTime = &v
	return s
}

func (s *CreateDownloadRequest) SetFormatType(v string) *CreateDownloadRequest {
	s.FormatType = &v
	return s
}

func (s *CreateDownloadRequest) SetInstanceName(v string) *CreateDownloadRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateDownloadRequest) SetIsCluster(v string) *CreateDownloadRequest {
	s.IsCluster = &v
	return s
}

func (s *CreateDownloadRequest) SetIsPhysical(v bool) *CreateDownloadRequest {
	s.IsPhysical = &v
	return s
}

func (s *CreateDownloadRequest) SetPrimaryKeyTypeOnly(v string) *CreateDownloadRequest {
	s.PrimaryKeyTypeOnly = &v
	return s
}

func (s *CreateDownloadRequest) SetRegionCode(v string) *CreateDownloadRequest {
	s.RegionCode = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetBucket(v string) *CreateDownloadRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetOssRegion(v string) *CreateDownloadRequest {
	s.TargetOssRegion = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetPath(v string) *CreateDownloadRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateDownloadRequest) SetTargetType(v string) *CreateDownloadRequest {
	s.TargetType = &v
	return s
}

func (s *CreateDownloadRequest) SetUseZstd(v string) *CreateDownloadRequest {
	s.UseZstd = &v
	return s
}

func (s *CreateDownloadRequest) Validate() error {
	return dara.Validate(s)
}
