// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeCrossRegionBackupDBInstanceResponseBodyItems) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetItems() *DescribeCrossRegionBackupDBInstanceResponseBodyItems
	SetItemsNumbers(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetItemsNumbers() *int32
	SetPageNumber(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetRequestId() *string
	SetTotalRecords(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody
	GetTotalRecords() *int32
}

type DescribeCrossRegionBackupDBInstanceResponseBody struct {
	// The cross-region backup settings.
	Items *DescribeCrossRegionBackupDBInstanceResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The total number of items returned for cross-region backup settings.
	//
	// example:
	//
	// 1
	ItemsNumbers *int32 `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 33517002-182D-40BE-93EC-610BD3381045
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalRecords *int32 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeCrossRegionBackupDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetItems() *DescribeCrossRegionBackupDBInstanceResponseBodyItems {
	return s.Items
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetItemsNumbers() *int32 {
	return s.ItemsNumbers
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) GetTotalRecords() *int32 {
	return s.TotalRecords
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetItems(v *DescribeCrossRegionBackupDBInstanceResponseBodyItems) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetItemsNumbers(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetPageNumber(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetPageSize(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetRegionId(v string) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetRequestId(v string) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) SetTotalRecords(v int32) *DescribeCrossRegionBackupDBInstanceResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCrossRegionBackupDBInstanceResponseBodyItems struct {
	Item []*DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeCrossRegionBackupDBInstanceResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupDBInstanceResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItems) GetItem() []*DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItems) SetItem(v []*DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) *DescribeCrossRegionBackupDBInstanceResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItems) Validate() error {
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

type DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem struct {
	// The status of the cross-region backup feature on the instance. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	BackupEnabled *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	// The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-12T05:44:21Z
	BackupEnabledTime *string `json:"BackupEnabledTime,omitempty" xml:"BackupEnabledTime,omitempty"`
	// The ID of the destination region within which the cross-region backup file is stored.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The policy that is used to save the cross-region backup files of the instance. Default value: **1**. The value 1 indicates that all cross-region backup files are saved.
	//
	// example:
	//
	// 1
	CrossBackupType *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	// The name of the instance. It must be 2 to 256 characters in length. The value can contain letters, digits, underscores (_), and hyphens (-), and must start with a letter.
	//
	// >  The value cannot start with http:// or https://.
	//
	// example:
	//
	// Test database
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance status. For more information, see [Instance statuses](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The lock status of the instance. Valid values:
	//
	// 	- **Unlock**: The instance is not locked.
	//
	// 	- **ManualLock**: The instance is manually locked.
	//
	// 	- **LockByExpiration**: The instance is automatically locked after it expires.
	//
	// 	- **LockByRestoration**: The instance is automatically locked before it is rolled back.
	//
	// 	- **LockByDiskQuota**: The instance is automatically locked because its storage capacity is exhausted and the instance is inaccessible.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The status of the cross-region log backup feature on the instance. Valid values:
	//
	// 	- **Disable**
	//
	// 	- **Enable**
	//
	// example:
	//
	// Enable
	LogBackupEnabled *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	// The time when the cross-region log backup feature was enabled on the instance. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-12T05:44:21Z
	LogBackupEnabledTime *string `json:"LogBackupEnabledTime,omitempty" xml:"LogBackupEnabledTime,omitempty"`
	// The policy that is used to retain the cross-region backup files of the instance. Cross-region backups can be retained only based on the specified retention period. Default value: **1**.
	//
	// example:
	//
	// 1
	RetentType *int32 `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	// The number of days for which the cross-region backup files of the instance are retained. Valid values: **7 to 1825**.
	//
	// example:
	//
	// 15
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
}

func (s DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetBackupEnabled() *string {
	return s.BackupEnabled
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetBackupEnabledTime() *string {
	return s.BackupEnabledTime
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetCrossBackupType() *string {
	return s.CrossBackupType
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetEngine() *string {
	return s.Engine
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetLogBackupEnabled() *string {
	return s.LogBackupEnabled
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetLogBackupEnabledTime() *string {
	return s.LogBackupEnabledTime
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetRetentType() *int32 {
	return s.RetentType
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) GetRetention() *int32 {
	return s.Retention
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetBackupEnabled(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.BackupEnabled = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetBackupEnabledTime(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.BackupEnabledTime = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetCrossBackupRegion(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetCrossBackupType(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.CrossBackupType = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetDBInstanceDescription(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetDBInstanceId(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetDBInstanceStatus(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetEngine(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.Engine = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetEngineVersion(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.EngineVersion = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetLockMode(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.LockMode = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetLogBackupEnabled(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.LogBackupEnabled = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetLogBackupEnabledTime(v string) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.LogBackupEnabledTime = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetRetentType(v int32) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.RetentType = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) SetRetention(v int32) *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem {
	s.Retention = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponseBodyItemsItem) Validate() error {
	return dara.Validate(s)
}
