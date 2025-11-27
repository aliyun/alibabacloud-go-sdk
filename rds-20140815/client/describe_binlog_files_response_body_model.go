// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinlogFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeBinlogFilesResponseBodyItems) *DescribeBinlogFilesResponseBody
	GetItems() *DescribeBinlogFilesResponseBodyItems
	SetPageNumber(v int32) *DescribeBinlogFilesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeBinlogFilesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeBinlogFilesResponseBody
	GetRequestId() *string
	SetTotalFileSize(v int64) *DescribeBinlogFilesResponseBody
	GetTotalFileSize() *int64
	SetTotalRecordCount(v int32) *DescribeBinlogFilesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeBinlogFilesResponseBody struct {
	// The details of the log file.
	Items *DescribeBinlogFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log files on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ED169A3E-1657-4104-82AB-24EA8CD0DB75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total size of the log file.
	//
	// example:
	//
	// 2269410
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// The total number of log files.
	//
	// example:
	//
	// 100
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeBinlogFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesResponseBody) GetItems() *DescribeBinlogFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeBinlogFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBinlogFilesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeBinlogFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBinlogFilesResponseBody) GetTotalFileSize() *int64 {
	return s.TotalFileSize
}

func (s *DescribeBinlogFilesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeBinlogFilesResponseBody) SetItems(v *DescribeBinlogFilesResponseBodyItems) *DescribeBinlogFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetPageNumber(v int32) *DescribeBinlogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetPageRecordCount(v int32) *DescribeBinlogFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetRequestId(v string) *DescribeBinlogFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetTotalFileSize(v int64) *DescribeBinlogFilesResponseBody {
	s.TotalFileSize = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetTotalRecordCount(v int32) *DescribeBinlogFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBinlogFilesResponseBodyItems struct {
	BinLogFile []*DescribeBinlogFilesResponseBodyItemsBinLogFile `json:"BinLogFile,omitempty" xml:"BinLogFile,omitempty" type:"Repeated"`
}

func (s DescribeBinlogFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesResponseBodyItems) GetBinLogFile() []*DescribeBinlogFilesResponseBodyItemsBinLogFile {
	return s.BinLogFile
}

func (s *DescribeBinlogFilesResponseBodyItems) SetBinLogFile(v []*DescribeBinlogFilesResponseBodyItemsBinLogFile) *DescribeBinlogFilesResponseBodyItems {
	s.BinLogFile = v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItems) Validate() error {
	if s.BinLogFile != nil {
		for _, item := range s.BinLogFile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBinlogFilesResponseBodyItemsBinLogFile struct {
	// The checksum. The value of this parameter is calculated by using the CRC64 algorithm.
	//
	// example:
	//
	// 18358304393468701857
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The HTTP-based download URL of the log file. If the return value of this parameter is NULL, ApsaraDB RDS does not provide a download URL for the log file.
	//
	// example:
	//
	// http://rdsxxxxx.oss.aliyuncs.com/xxxxxx
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The size of the log file.
	//
	// Unit: bytes.
	//
	// example:
	//
	// 2269410
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The ID of the instance to which the log file belongs. This parameter helps determine whether the log file is generated on the primary instance or the secondary instance.
	//
	// >  You can log on to the ApsaraDB RDS console and go to the instance details page. In the left-side navigation pane, click **Service Availability*	- to view the values of **Primary Instance No.*	- and **Secondary Instance No.**.
	//
	// example:
	//
	// 5841973
	HostInstanceID *string `json:"HostInstanceID,omitempty" xml:"HostInstanceID,omitempty"`
	// The URL that is used to download files over an internal network.
	//
	// example:
	//
	// http://rdslog-hz-v3.oss-cn-hangzhou-internal.aliyuncs.com/xxxxxx
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// The expiration time of the URL.
	//
	// The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2013-06-09T18:00:00Z
	LinkExpiredTime *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
	// The beginning of the time range to query.
	//
	// The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-09T17:45:21Z
	LogBeginTime *string `json:"LogBeginTime,omitempty" xml:"LogBeginTime,omitempty"`
	// The end of the time range to query.
	//
	// The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-15T13:10:28Z
	LogEndTime *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	// The log file name.
	//
	// example:
	//
	// 000000040000000000000019
	LogFileName *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
	// The status of the log file that is stored in the Object Storage Service (OSS) bucket.
	//
	// Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **Completed**
	//
	// example:
	//
	// Completed
	RemoteStatus *string `json:"RemoteStatus,omitempty" xml:"RemoteStatus,omitempty"`
}

func (s DescribeBinlogFilesResponseBodyItemsBinLogFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesResponseBodyItemsBinLogFile) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetChecksum() *string {
	return s.Checksum
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetHostInstanceID() *string {
	return s.HostInstanceID
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetLogBeginTime() *string {
	return s.LogBeginTime
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetLogEndTime() *string {
	return s.LogEndTime
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetLogFileName() *string {
	return s.LogFileName
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) GetRemoteStatus() *string {
	return s.RemoteStatus
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetChecksum(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.Checksum = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetDownloadLink(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetFileSize(v int64) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.FileSize = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetHostInstanceID(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.HostInstanceID = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetIntranetDownloadLink(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetLinkExpiredTime(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetLogBeginTime(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.LogBeginTime = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetLogEndTime(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetLogFileName(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.LogFileName = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) SetRemoteStatus(v string) *DescribeBinlogFilesResponseBodyItemsBinLogFile {
	s.RemoteStatus = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsBinLogFile) Validate() error {
	return dara.Validate(s)
}
