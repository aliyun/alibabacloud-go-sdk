// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLLogFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeSQLLogFilesResponseBodyItems) *DescribeSQLLogFilesResponseBody
	GetItems() *DescribeSQLLogFilesResponseBodyItems
	SetPageNumber(v int32) *DescribeSQLLogFilesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSQLLogFilesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSQLLogFilesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeSQLLogFilesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeSQLLogFilesResponseBody struct {
	// An array that consists of the returned audit log files.
	Items *DescribeSQLLogFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSQLLogFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBody) GetItems() *DescribeSQLLogFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLLogFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSQLLogFilesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSQLLogFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLLogFilesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeSQLLogFilesResponseBody) SetItems(v *DescribeSQLLogFilesResponseBodyItems) *DescribeSQLLogFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetPageNumber(v int32) *DescribeSQLLogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetPageRecordCount(v int32) *DescribeSQLLogFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetRequestId(v string) *DescribeSQLLogFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) SetTotalRecordCount(v int32) *DescribeSQLLogFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLLogFilesResponseBodyItems struct {
	LogFile []*DescribeSQLLogFilesResponseBodyItemsLogFile `json:"LogFile,omitempty" xml:"LogFile,omitempty" type:"Repeated"`
}

func (s DescribeSQLLogFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBodyItems) GetLogFile() []*DescribeSQLLogFilesResponseBodyItemsLogFile {
	return s.LogFile
}

func (s *DescribeSQLLogFilesResponseBodyItems) SetLogFile(v []*DescribeSQLLogFilesResponseBodyItemsLogFile) *DescribeSQLLogFilesResponseBodyItems {
	s.LogFile = v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItems) Validate() error {
	if s.LogFile != nil {
		for _, item := range s.LogFile {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLLogFilesResponseBodyItemsLogFile struct {
	// The file name.
	//
	// example:
	//
	// custinsxxxxx.csv
	FileID *string `json:"FileID,omitempty" xml:"FileID,omitempty"`
	// The download URL of the file. If the audit log file cannot be downloaded, this parameter is null.
	//
	// example:
	//
	// http://rdslog-hz-v3.oss-cn-hangzhou.aliyuncs.com/xxxxx
	LogDownloadURL *string `json:"LogDownloadURL,omitempty" xml:"LogDownloadURL,omitempty"`
	// The time at which the last SQL statement recorded in the audit log file was executed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-05-24T07:00:00Z
	LogEndTime *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	// The size of the audit log file. Unit: bytes.
	//
	// example:
	//
	// 3000
	LogSize *string `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The time at which the first SQL statement recorded in the audit log file was executed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-05-23T07:00:00Z
	LogStartTime *string `json:"LogStartTime,omitempty" xml:"LogStartTime,omitempty"`
	// The status of the audit log file. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// 	- **Generating**
	//
	// example:
	//
	// Success
	LogStatus *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
}

func (s DescribeSQLLogFilesResponseBodyItemsLogFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLLogFilesResponseBodyItemsLogFile) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) GetFileID() *string {
	return s.FileID
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) GetLogDownloadURL() *string {
	return s.LogDownloadURL
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) GetLogEndTime() *string {
	return s.LogEndTime
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) GetLogSize() *string {
	return s.LogSize
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) GetLogStartTime() *string {
	return s.LogStartTime
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) GetLogStatus() *string {
	return s.LogStatus
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetFileID(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.FileID = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogDownloadURL(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogDownloadURL = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogEndTime(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogSize(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogSize = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogStartTime(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogStartTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) SetLogStatus(v string) *DescribeSQLLogFilesResponseBodyItemsLogFile {
	s.LogStatus = &v
	return s
}

func (s *DescribeSQLLogFilesResponseBodyItemsLogFile) Validate() error {
	return dara.Validate(s)
}
