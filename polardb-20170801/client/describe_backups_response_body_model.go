// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody
	GetItems() *DescribeBackupsResponseBodyItems
	SetPageNumber(v string) *DescribeBackupsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeBackupsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetTotalLevel2BackupSize(v string) *DescribeBackupsResponseBody
	GetTotalLevel2BackupSize() *string
	SetTotalRecordCount(v string) *DescribeBackupsResponseBody
	GetTotalRecordCount() *string
}

type DescribeBackupsResponseBody struct {
	Items *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4639948800
	TotalLevel2BackupSize *string `json:"TotalLevel2BackupSize,omitempty" xml:"TotalLevel2BackupSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetItems() *DescribeBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeBackupsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetTotalLevel2BackupSize() *string {
	return s.TotalLevel2BackupSize
}

func (s *DescribeBackupsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageRecordCount(v string) *DescribeBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalLevel2BackupSize(v string) *DescribeBackupsResponseBody {
	s.TotalLevel2BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalRecordCount(v string) *DescribeBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupsResponseBodyItems struct {
	Backup []*DescribeBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) GetBackup() []*DescribeBackupsResponseBodyItemsBackup {
	return s.Backup
}

func (s *DescribeBackupsResponseBodyItems) SetBackup(v []*DescribeBackupsResponseBodyItemsBackup) *DescribeBackupsResponseBodyItems {
	s.Backup = v
	return s
}

func (s *DescribeBackupsResponseBodyItems) Validate() error {
	if s.Backup != nil {
		for _, item := range s.Backup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupsResponseBodyItemsBackup struct {
	BackupEndTime    *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId         *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMethod     *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupMode       *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupSetSize    *string `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	BackupStartTime  *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStatus     *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType       *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupsLevel     *string `json:"BackupsLevel,omitempty" xml:"BackupsLevel,omitempty"`
	ConsistentTime   *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	DBClusterId      *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	ExpectExpireType *string `json:"ExpectExpireType,omitempty" xml:"ExpectExpireType,omitempty"`
	IsAvail          *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupSetSize() *string {
	return s.BackupSetSize
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupsLevel() *string {
	return s.BackupsLevel
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetConsistentTime() *string {
	return s.ConsistentTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetExpectExpireType() *string {
	return s.ExpectExpireType
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetIsAvail() *string {
	return s.IsAvail
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSetSize(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupsLevel(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetConsistentTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetExpectExpireTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetExpectExpireType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ExpectExpireType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetIsAvail(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) Validate() error {
	return dara.Validate(s)
}
