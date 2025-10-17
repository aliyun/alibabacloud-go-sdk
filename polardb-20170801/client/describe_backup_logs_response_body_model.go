// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeBackupLogsResponseBodyItems) *DescribeBackupLogsResponseBody
	GetItems() *DescribeBackupLogsResponseBodyItems
	SetPageNumber(v string) *DescribeBackupLogsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeBackupLogsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeBackupLogsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeBackupLogsResponseBody
	GetTotalRecordCount() *string
}

type DescribeBackupLogsResponseBody struct {
	// The details of the backup logs.
	Items *DescribeBackupLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ADAC63DB-0691-4ECE-949A-FAEA68******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeBackupLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseBody) GetItems() *DescribeBackupLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupLogsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeBackupLogsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeBackupLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupLogsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeBackupLogsResponseBody) SetItems(v *DescribeBackupLogsResponseBodyItems) *DescribeBackupLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetPageNumber(v string) *DescribeBackupLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetPageRecordCount(v string) *DescribeBackupLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetRequestId(v string) *DescribeBackupLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetTotalRecordCount(v string) *DescribeBackupLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupLogsResponseBodyItems struct {
	BackupLog []*DescribeBackupLogsResponseBodyItemsBackupLog `json:"BackupLog,omitempty" xml:"BackupLog,omitempty" type:"Repeated"`
}

func (s DescribeBackupLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseBodyItems) GetBackupLog() []*DescribeBackupLogsResponseBodyItemsBackupLog {
	return s.BackupLog
}

func (s *DescribeBackupLogsResponseBodyItems) SetBackupLog(v []*DescribeBackupLogsResponseBodyItemsBackupLog) *DescribeBackupLogsResponseBodyItems {
	s.BackupLog = v
	return s
}

func (s *DescribeBackupLogsResponseBodyItems) Validate() error {
	if s.BackupLog != nil {
		for _, item := range s.BackupLog {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupLogsResponseBodyItemsBackupLog struct {
	// The time when the backup task ended. The time follows the ISO 8601 standard in the `YYYY-MM-DD\\"T\\"HH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-02-12T03:55:31Z
	BackupLogEndTime *string `json:"BackupLogEndTime,omitempty" xml:"BackupLogEndTime,omitempty"`
	// The ID of the backup log.
	//
	// example:
	//
	// 1111111111
	BackupLogId *string `json:"BackupLogId,omitempty" xml:"BackupLogId,omitempty"`
	// The name of the backup log.
	//
	// example:
	//
	// ib_logfile1
	BackupLogName *string `json:"BackupLogName,omitempty" xml:"BackupLogName,omitempty"`
	// The size of the backup log. Unit: bytes.
	//
	// example:
	//
	// 1073741824
	BackupLogSize *string `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	// The time when the backup task started. The time follows the ISO 8601 standard in the `YYYY-MM-DD\\"T\\"HH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-02-12T03:54:43Z
	BackupLogStartTime *string `json:"BackupLogStartTime,omitempty" xml:"BackupLogStartTime,omitempty"`
	// The public URL used to download the backup log.
	//
	// example:
	//
	// http://***********.oss-cn-hangzhou.aliyuncs.com
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The internal URL used to download the backup log.
	//
	// example:
	//
	// http://***********.oss-cn-hangzhou-internal.aliyuncs.com
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// The time when the download URL expires.
	//
	// example:
	//
	// 2020-02-14T08:40:50Z
	LinkExpiredTime *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
}

func (s DescribeBackupLogsResponseBodyItemsBackupLog) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogsResponseBodyItemsBackupLog) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetBackupLogEndTime() *string {
	return s.BackupLogEndTime
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetBackupLogId() *string {
	return s.BackupLogId
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetBackupLogName() *string {
	return s.BackupLogName
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetBackupLogSize() *string {
	return s.BackupLogSize
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetBackupLogStartTime() *string {
	return s.BackupLogStartTime
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogEndTime(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogEndTime = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogId(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogId = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogName(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogName = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogSize(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogStartTime(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogStartTime = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetDownloadLink(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetIntranetDownloadLink(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetLinkExpiredTime(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) Validate() error {
	return dara.Validate(s)
}
