// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeSharedBackupsResponseBodyItems) *DescribeSharedBackupsResponseBody
	GetItems() []*DescribeSharedBackupsResponseBodyItems
	SetPageNumber(v string) *DescribeSharedBackupsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeSharedBackupsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeSharedBackupsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeSharedBackupsResponseBody
	GetTotalRecordCount() *string
}

type DescribeSharedBackupsResponseBody struct {
	Items []*DescribeSharedBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSharedBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSharedBackupsResponseBody) GetItems() []*DescribeSharedBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeSharedBackupsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeSharedBackupsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeSharedBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSharedBackupsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeSharedBackupsResponseBody) SetItems(v []*DescribeSharedBackupsResponseBodyItems) *DescribeSharedBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSharedBackupsResponseBody) SetPageNumber(v string) *DescribeSharedBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSharedBackupsResponseBody) SetPageRecordCount(v string) *DescribeSharedBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSharedBackupsResponseBody) SetRequestId(v string) *DescribeSharedBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSharedBackupsResponseBody) SetTotalRecordCount(v string) *DescribeSharedBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSharedBackupsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSharedBackupsResponseBodyItems struct {
	// example:
	//
	// 2020-05-12T03:25:55Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// Manual
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// example:
	//
	// 4639948800
	BackupSetSize *string `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// example:
	//
	// 2020-11-15T07:30:05Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// Level-2
	BackupsLevel *string `json:"BackupsLevel,omitempty" xml:"BackupsLevel,omitempty"`
	// example:
	//
	// 1589253947
	ConsistentTime *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hongzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// example:
	//
	// ShareIncoming
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// example:
	//
	// 170*************
	SharerUID *string `json:"SharerUID,omitempty" xml:"SharerUID,omitempty"`
}

func (s DescribeSharedBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupSetSize() *string {
	return s.BackupSetSize
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeSharedBackupsResponseBodyItems) GetBackupsLevel() *string {
	return s.BackupsLevel
}

func (s *DescribeSharedBackupsResponseBodyItems) GetConsistentTime() *string {
	return s.ConsistentTime
}

func (s *DescribeSharedBackupsResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSharedBackupsResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribeSharedBackupsResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeSharedBackupsResponseBodyItems) GetPayType() *string {
	return s.PayType
}

func (s *DescribeSharedBackupsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSharedBackupsResponseBodyItems) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeSharedBackupsResponseBodyItems) GetShareType() *string {
	return s.ShareType
}

func (s *DescribeSharedBackupsResponseBodyItems) GetSharerUID() *string {
	return s.SharerUID
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupId(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupMethod(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupMethod = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupMode(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupSetSize(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupType(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupType = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetBackupsLevel(v string) *DescribeSharedBackupsResponseBodyItems {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetConsistentTime(v string) *DescribeSharedBackupsResponseBodyItems {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetDBClusterId(v string) *DescribeSharedBackupsResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetDBType(v string) *DescribeSharedBackupsResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetDBVersion(v string) *DescribeSharedBackupsResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetPayType(v string) *DescribeSharedBackupsResponseBodyItems {
	s.PayType = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetRegionId(v string) *DescribeSharedBackupsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetServerlessType(v string) *DescribeSharedBackupsResponseBodyItems {
	s.ServerlessType = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetShareType(v string) *DescribeSharedBackupsResponseBodyItems {
	s.ShareType = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) SetSharerUID(v string) *DescribeSharedBackupsResponseBodyItems {
	s.SharerUID = &v
	return s
}

func (s *DescribeSharedBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
