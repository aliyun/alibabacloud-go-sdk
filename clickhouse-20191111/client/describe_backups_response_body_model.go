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
	// The backup sets.
	Items []*DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	// The end time of the backup task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-22T18:28:41Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The backup task ID.
	//
	// example:
	//
	// 117403****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Valid values: Only **Physical*	- is returned. This value indicates that a physical backup was performed.
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The number of nodes in the cluster.
	//
	// 	- If the cluster is of Single-replica Edition, the value ranges from 1 to 48.
	//
	// 	- If the cluster is of Double-replica Edition, the value ranges from 1 to 24.
	//
	// example:
	//
	// {"shard_count":4}
	BackupSetInfo *string `json:"BackupSetInfo,omitempty" xml:"BackupSetInfo,omitempty"`
	// The size of the backup set. Unit: MB.
	//
	// example:
	//
	// 131072
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup task. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-22T18:28:22Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup status. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failure**
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementalBackup**
	//
	// example:
	//
	// IncrementalBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp1qx68m06981****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The time when the backup set expired. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
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

func (s *DescribeBackupsResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeBackupsResponseBodyItems) SetDBClusterId(v string) *DescribeBackupsResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) SetExpireDate(v string) *DescribeBackupsResponseBodyItems {
	s.ExpireDate = &v
	return s
}

func (s *DescribeBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
