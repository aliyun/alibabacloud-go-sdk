// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCrossRegionBackupsResponseBody
	GetEndTime() *string
	SetItems(v *DescribeCrossRegionBackupsResponseBodyItems) *DescribeCrossRegionBackupsResponseBody
	GetItems() *DescribeCrossRegionBackupsResponseBodyItems
	SetPageNumber(v int32) *DescribeCrossRegionBackupsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeCrossRegionBackupsResponseBody
	GetPageRecordCount() *int32
	SetRegionId(v string) *DescribeCrossRegionBackupsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeCrossRegionBackupsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeCrossRegionBackupsResponseBody
	GetStartTime() *string
	SetTotalRecordCount(v int32) *DescribeCrossRegionBackupsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCrossRegionBackupsResponseBody struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2019-06-15T12:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cross-region data backup files.
	Items *DescribeCrossRegionBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number. Pages start from page 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of cross-region data backup files on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60912B41-7579-4B5D-B289-8856030F0A6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2019-05-30T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCrossRegionBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCrossRegionBackupsResponseBody) GetItems() *DescribeCrossRegionBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeCrossRegionBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionBackupsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeCrossRegionBackupsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossRegionBackupsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCrossRegionBackupsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCrossRegionBackupsResponseBody) SetEndTime(v string) *DescribeCrossRegionBackupsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetItems(v *DescribeCrossRegionBackupsResponseBodyItems) *DescribeCrossRegionBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetPageNumber(v int32) *DescribeCrossRegionBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetPageRecordCount(v int32) *DescribeCrossRegionBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetRegionId(v string) *DescribeCrossRegionBackupsResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetRequestId(v string) *DescribeCrossRegionBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetStartTime(v string) *DescribeCrossRegionBackupsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) SetTotalRecordCount(v int32) *DescribeCrossRegionBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCrossRegionBackupsResponseBodyItems struct {
	Item []*DescribeCrossRegionBackupsResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeCrossRegionBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsResponseBodyItems) GetItem() []*DescribeCrossRegionBackupsResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeCrossRegionBackupsResponseBodyItems) SetItem(v []*DescribeCrossRegionBackupsResponseBodyItemsItem) *DescribeCrossRegionBackupsResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCrossRegionBackupsResponseBodyItemsItem struct {
	// The time when the cross-region data backup file was generated.
	//
	// example:
	//
	// 2019-06-15T12:10:00Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The method that is used to generate the cross-region data backup file. Valid values:
	//
	// 	- **L**: logical backup
	//
	// 	- **P**: physical backup
	//
	// example:
	//
	// P
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The level at which the cross-region data backup file is generated.
	//
	// 	- **0**: instance-level backup
	//
	// 	- **1**: database-level backup
	//
	// example:
	//
	// 0
	BackupSetScale *int32 `json:"BackupSetScale,omitempty" xml:"BackupSetScale,omitempty"`
	// The status of the cross-region data backup. Valid values:
	//
	// 	- **0**: The cross-region data backup is successful.
	//
	// 	- **1**: The cross-region data backup failed.
	//
	// example:
	//
	// 0
	BackupSetStatus *int32 `json:"BackupSetStatus,omitempty" xml:"BackupSetStatus,omitempty"`
	// The time when the cross-region data backup started.
	//
	// example:
	//
	// 2019-05-30T12:10:00Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The type of the cross-region data backup. Valid values:
	//
	// 	- **F**: full data backup
	//
	// 	- **I**: incremental data backup
	//
	// example:
	//
	// F
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition.
	//
	// 	- **HighAvailability**: RDS High-availability Edition.
	//
	// 	- **Finance**: RDS Enterprise Edition. This edition is available only for the China site (aliyun.com).
	//
	// example:
	//
	// HighAvailability
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The point in time that is indicated by the data in the cross-region data backup file.
	//
	// example:
	//
	// 2019-06-12T05:44:46Z
	ConsistentTime *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The external URL from which you can download the cross-region data backup file.
	//
	// example:
	//
	// http://rdsddrbak-shanghai.oss-cn-shanghai.aliyuncs.com/xxxxx
	CrossBackupDownloadLink *string `json:"CrossBackupDownloadLink,omitempty" xml:"CrossBackupDownloadLink,omitempty"`
	// The ID of the cross-region data backup file.
	//
	// example:
	//
	// 14377
	CrossBackupId *int32 `json:"CrossBackupId,omitempty" xml:"CrossBackupId,omitempty"`
	// The ID of the region in which the cross-region backup files of the instance are stored.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The name of the compressed package that contains the cross-region data backup file.
	//
	// example:
	//
	// cn-hangzhou_rm-xxxxx_hins81xxx_data_20190612134426_qp.xb
	CrossBackupSetFile *string `json:"CrossBackupSetFile,omitempty" xml:"CrossBackupSetFile,omitempty"`
	// The location where the cross-region data backup file is stored.
	//
	// example:
	//
	// oss
	CrossBackupSetLocation *string `json:"CrossBackupSetLocation,omitempty" xml:"CrossBackupSetLocation,omitempty"`
	// The size of the cross-region data backup file. Unit: bytes.
	//
	// example:
	//
	// 5312836
	CrossBackupSetSize *int64 `json:"CrossBackupSetSize,omitempty" xml:"CrossBackupSetSize,omitempty"`
	// The storage type. Valid values:
	//
	// 	- **local_ssd**: local SSDs. This is the recommended storage type.
	//
	// 	- **cloud_ssd**: standard SSD.
	//
	// 	- **cloud_essd**: enhanced SSD (ESSD).
	//
	// example:
	//
	// ssd
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance ID. This parameter is used to determine whether the instance that generates the cross-region data backup file is a primary or secondary instance.
	//
	// example:
	//
	// 8161055
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The regions to which the cross-region data backup file can be restored.
	RestoreRegions *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions `json:"RestoreRegions,omitempty" xml:"RestoreRegions,omitempty" type:"Struct"`
}

func (s DescribeCrossRegionBackupsResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetBackupSetScale() *int32 {
	return s.BackupSetScale
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetBackupSetStatus() *int32 {
	return s.BackupSetStatus
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCategory() *string {
	return s.Category
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetConsistentTime() *string {
	return s.ConsistentTime
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCrossBackupDownloadLink() *string {
	return s.CrossBackupDownloadLink
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCrossBackupId() *int32 {
	return s.CrossBackupId
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCrossBackupSetFile() *string {
	return s.CrossBackupSetFile
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCrossBackupSetLocation() *string {
	return s.CrossBackupSetLocation
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetCrossBackupSetSize() *int64 {
	return s.CrossBackupSetSize
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetEngine() *string {
	return s.Engine
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) GetRestoreRegions() *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions {
	return s.RestoreRegions
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetBackupEndTime(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetBackupMethod(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.BackupMethod = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetBackupSetScale(v int32) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.BackupSetScale = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetBackupSetStatus(v int32) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.BackupSetStatus = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetBackupStartTime(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetBackupType(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.BackupType = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCategory(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.Category = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetConsistentTime(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCrossBackupDownloadLink(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.CrossBackupDownloadLink = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCrossBackupId(v int32) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.CrossBackupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCrossBackupRegion(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCrossBackupSetFile(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.CrossBackupSetFile = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCrossBackupSetLocation(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.CrossBackupSetLocation = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetCrossBackupSetSize(v int64) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.CrossBackupSetSize = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetDBInstanceStorageType(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.DBInstanceStorageType = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetEngine(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.Engine = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetEngineVersion(v string) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.EngineVersion = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetInstanceId(v int32) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.InstanceId = &v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) SetRestoreRegions(v *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions) *DescribeCrossRegionBackupsResponseBodyItemsItem {
	s.RestoreRegions = v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItem) Validate() error {
	if s.RestoreRegions != nil {
		if err := s.RestoreRegions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions struct {
	RestoreRegion []*string `json:"RestoreRegion,omitempty" xml:"RestoreRegion,omitempty" type:"Repeated"`
}

func (s DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions) GetRestoreRegion() []*string {
	return s.RestoreRegion
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions) SetRestoreRegion(v []*string) *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions {
	s.RestoreRegion = v
	return s
}

func (s *DescribeCrossRegionBackupsResponseBodyItemsItemRestoreRegions) Validate() error {
	return dara.Validate(s)
}
