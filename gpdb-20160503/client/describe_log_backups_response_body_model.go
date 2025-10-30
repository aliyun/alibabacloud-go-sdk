// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeLogBackupsResponseBodyItems) *DescribeLogBackupsResponseBody
	GetItems() []*DescribeLogBackupsResponseBodyItems
	SetPageNumber(v int32) *DescribeLogBackupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLogBackupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLogBackupsResponseBody
	GetTotalCount() *int32
	SetTotalLogSize(v int64) *DescribeLogBackupsResponseBody
	GetTotalLogSize() *int64
}

type DescribeLogBackupsResponseBody struct {
	// Details of the backup sets.
	Items []*DescribeLogBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of backup sets on the current page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 24B9FCAF-2CBC-51C3-B563-F1C70D750187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total size of logs in the time range. Unit: bytes.
	//
	// example:
	//
	// 386748
	TotalLogSize *int64 `json:"TotalLogSize,omitempty" xml:"TotalLogSize,omitempty"`
}

func (s DescribeLogBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBody) GetItems() []*DescribeLogBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeLogBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLogBackupsResponseBody) GetTotalLogSize() *int64 {
	return s.TotalLogSize
}

func (s *DescribeLogBackupsResponseBody) SetItems(v []*DescribeLogBackupsResponseBodyItems) *DescribeLogBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageNumber(v int32) *DescribeLogBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetPageSize(v int32) *DescribeLogBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetRequestId(v string) *DescribeLogBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetTotalCount(v int32) *DescribeLogBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) SetTotalLogSize(v int64) *DescribeLogBackupsResponseBody {
	s.TotalLogSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBody) Validate() error {
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

type DescribeLogBackupsResponseBodyItems struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 12413721782
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the coordinator node.
	//
	// example:
	//
	// gp-bp12ga6v69h86****-master
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the log backup set that is stored in Object Storage Service (OSS).
	//
	// example:
	//
	// 000000010000000400000012
	LogFileName *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
	// The size of the log backup set. Unit: bytes.
	//
	// example:
	//
	// 77350
	LogFileSize *int64 `json:"LogFileSize,omitempty" xml:"LogFileSize,omitempty"`
	// The timestamp of the log.
	//
	// example:
	//
	// 2022-12-12T02:14:26Z
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The name of the compute node.
	//
	// example:
	//
	// segment-0
	SegmentName *string `json:"SegmentName,omitempty" xml:"SegmentName,omitempty"`
}

func (s DescribeLogBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponseBodyItems) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeLogBackupsResponseBodyItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeLogBackupsResponseBodyItems) GetLogFileName() *string {
	return s.LogFileName
}

func (s *DescribeLogBackupsResponseBodyItems) GetLogFileSize() *int64 {
	return s.LogFileSize
}

func (s *DescribeLogBackupsResponseBodyItems) GetLogTime() *string {
	return s.LogTime
}

func (s *DescribeLogBackupsResponseBodyItems) GetSegmentName() *string {
	return s.SegmentName
}

func (s *DescribeLogBackupsResponseBodyItems) SetBackupId(v string) *DescribeLogBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeLogBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileName(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogFileName = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogFileSize(v int64) *DescribeLogBackupsResponseBodyItems {
	s.LogFileSize = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetLogTime(v string) *DescribeLogBackupsResponseBodyItems {
	s.LogTime = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) SetSegmentName(v string) *DescribeLogBackupsResponseBodyItems {
	s.SegmentName = &v
	return s
}

func (s *DescribeLogBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
