// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody
	GetItems() []*DescribeBackupsResponseBodyItems
	SetPageNumber(v string) *DescribeBackupsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeBackupsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeBackupsResponseBody
	GetTotalCount() *string
}

type DescribeBackupsResponseBody struct {
	Items []*DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetItems() []*DescribeBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeBackupsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeBackupsResponseBody) SetItems(v []*DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v string) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v string) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
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

type DescribeBackupsResponseBodyItems struct {
	// example:
	//
	// 2021-11-22T18:28:41Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// example:
	//
	// 117403****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// example:
	//
	// {"shard_count"：4}
	BackupSetInfo *string `json:"BackupSetInfo,omitempty" xml:"BackupSetInfo,omitempty"`
	// example:
	//
	// 131072
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// example:
	//
	// 2021-11-22T18:28:22Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// IncrementalBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// example:
	//
	// cc-bp179i5956tih2m93
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 2022-07-22T18:28:41Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupsResponseBodyItems) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyItems) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyItems) GetBackupSetInfo() *string {
	return s.BackupSetInfo
}

func (s *DescribeBackupsResponseBodyItems) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeBackupsResponseBodyItems) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupsResponseBodyItems) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupsResponseBodyItems) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupsResponseBodyItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupsResponseBodyItems) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupId(v string) *DescribeBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupMethod(v string) *DescribeBackupsResponseBodyItems {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupSetInfo(v string) *DescribeBackupsResponseBodyItems {
	s.BackupSetInfo = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupSize(v int64) *DescribeBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetBackupType(v string) *DescribeBackupsResponseBodyItems {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetExpireDate(v string) *DescribeBackupsResponseBodyItems {
	s.ExpireDate = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
