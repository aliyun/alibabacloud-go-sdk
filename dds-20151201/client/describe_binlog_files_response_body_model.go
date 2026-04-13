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
	SetMaxRecordsPerPage(v int32) *DescribeBinlogFilesResponseBody
	GetMaxRecordsPerPage() *int32
	SetPageNumber(v int32) *DescribeBinlogFilesResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeBinlogFilesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeBinlogFilesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeBinlogFilesResponseBody struct {
	Items *DescribeBinlogFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 30
	MaxRecordsPerPage *int32 `json:"MaxRecordsPerPage,omitempty" xml:"MaxRecordsPerPage,omitempty"`
	// example:
	//
	// 3
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// F8CA8312-530A-413A-9129-F2BB32A8D404
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 240
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

func (s *DescribeBinlogFilesResponseBody) GetMaxRecordsPerPage() *int32 {
	return s.MaxRecordsPerPage
}

func (s *DescribeBinlogFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBinlogFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBinlogFilesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeBinlogFilesResponseBody) SetItems(v *DescribeBinlogFilesResponseBodyItems) *DescribeBinlogFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetMaxRecordsPerPage(v int32) *DescribeBinlogFilesResponseBody {
	s.MaxRecordsPerPage = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetPageNumber(v int32) *DescribeBinlogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinlogFilesResponseBody) SetRequestId(v string) *DescribeBinlogFilesResponseBody {
	s.RequestId = &v
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
	LogFile []*DescribeBinlogFilesResponseBodyItemsLogFile `json:"LogFile,omitempty" xml:"LogFile,omitempty" type:"Repeated"`
}

func (s DescribeBinlogFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesResponseBodyItems) GetLogFile() []*DescribeBinlogFilesResponseBodyItemsLogFile {
	return s.LogFile
}

func (s *DescribeBinlogFilesResponseBodyItems) SetLogFile(v []*DescribeBinlogFilesResponseBodyItemsLogFile) *DescribeBinlogFilesResponseBodyItems {
	s.LogFile = v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItems) Validate() error {
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

type DescribeBinlogFilesResponseBodyItemsLogFile struct {
	BinLogId        *string `json:"BinLogId,omitempty" xml:"BinLogId,omitempty"`
	DownloadLink    *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	DumpBucket      *string `json:"DumpBucket,omitempty" xml:"DumpBucket,omitempty"`
	DumpDownloadURL *string `json:"DumpDownloadURL,omitempty" xml:"DumpDownloadURL,omitempty"`
	DumpState       *int32  `json:"DumpState,omitempty" xml:"DumpState,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	FileSize        *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	LinkExpiredTime *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
	LogBeginTime    *string `json:"LogBeginTime,omitempty" xml:"LogBeginTime,omitempty"`
	LogEndTime      *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
	LogFileName     *string `json:"LogFileName,omitempty" xml:"LogFileName,omitempty"`
	LogStatus       *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty"`
	OSSEndpoint     *string `json:"OSSEndpoint,omitempty" xml:"OSSEndpoint,omitempty"`
}

func (s DescribeBinlogFilesResponseBodyItemsLogFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesResponseBodyItemsLogFile) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetBinLogId() *string {
	return s.BinLogId
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetDumpBucket() *string {
	return s.DumpBucket
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetDumpDownloadURL() *string {
	return s.DumpDownloadURL
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetDumpState() *int32 {
	return s.DumpState
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetFileId() *string {
	return s.FileId
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetFileSize() *string {
	return s.FileSize
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetLogBeginTime() *string {
	return s.LogBeginTime
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetLogEndTime() *string {
	return s.LogEndTime
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetLogFileName() *string {
	return s.LogFileName
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetLogStatus() *string {
	return s.LogStatus
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) GetOSSEndpoint() *string {
	return s.OSSEndpoint
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetBinLogId(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.BinLogId = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetDownloadLink(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetDumpBucket(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.DumpBucket = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetDumpDownloadURL(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.DumpDownloadURL = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetDumpState(v int32) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.DumpState = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetFileId(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.FileId = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetFileSize(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.FileSize = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetLinkExpiredTime(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetLogBeginTime(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.LogBeginTime = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetLogEndTime(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetLogFileName(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.LogFileName = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetLogStatus(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.LogStatus = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) SetOSSEndpoint(v string) *DescribeBinlogFilesResponseBodyItemsLogFile {
	s.OSSEndpoint = &v
	return s
}

func (s *DescribeBinlogFilesResponseBodyItemsLogFile) Validate() error {
	return dara.Validate(s)
}
