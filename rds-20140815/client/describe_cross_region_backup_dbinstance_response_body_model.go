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
	BackupEnabled         *string `json:"BackupEnabled,omitempty" xml:"BackupEnabled,omitempty"`
	BackupEnabledTime     *string `json:"BackupEnabledTime,omitempty" xml:"BackupEnabledTime,omitempty"`
	CrossBackupRegion     *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	CrossBackupType       *string `json:"CrossBackupType,omitempty" xml:"CrossBackupType,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceStatus      *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	LockMode              *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LogBackupEnabled      *string `json:"LogBackupEnabled,omitempty" xml:"LogBackupEnabled,omitempty"`
	LogBackupEnabledTime  *string `json:"LogBackupEnabledTime,omitempty" xml:"LogBackupEnabledTime,omitempty"`
	RetentType            *int32  `json:"RetentType,omitempty" xml:"RetentType,omitempty"`
	Retention             *int32  `json:"Retention,omitempty" xml:"Retention,omitempty"`
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
