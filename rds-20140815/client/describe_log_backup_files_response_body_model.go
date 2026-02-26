// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeLogBackupFilesResponseBodyItems) *DescribeLogBackupFilesResponseBody
	GetItems() *DescribeLogBackupFilesResponseBodyItems
	SetPageNumber(v int32) *DescribeLogBackupFilesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeLogBackupFilesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeLogBackupFilesResponseBody
	GetRequestId() *string
	SetTotalFileSize(v int64) *DescribeLogBackupFilesResponseBody
	GetTotalFileSize() *int64
	SetTotalRecordCount(v int32) *DescribeLogBackupFilesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeLogBackupFilesResponseBody struct {
	Items *DescribeLogBackupFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log files on the current page.
	//
	// example:
	//
	// 100
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F8EC669C-FC85-43D7-AF06-C3641626B37E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total size of log files. Unit: bytes.
	//
	// example:
	//
	// 2300
	TotalFileSize *int64 `json:"TotalFileSize,omitempty" xml:"TotalFileSize,omitempty"`
	// The total number of log files.
	//
	// example:
	//
	// 17
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeLogBackupFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupFilesResponseBody) GetItems() *DescribeLogBackupFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeLogBackupFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogBackupFilesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeLogBackupFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogBackupFilesResponseBody) GetTotalFileSize() *int64 {
	return s.TotalFileSize
}

func (s *DescribeLogBackupFilesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeLogBackupFilesResponseBody) SetItems(v *DescribeLogBackupFilesResponseBodyItems) *DescribeLogBackupFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeLogBackupFilesResponseBody) SetPageNumber(v int32) *DescribeLogBackupFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBody) SetPageRecordCount(v int32) *DescribeLogBackupFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBody) SetRequestId(v string) *DescribeLogBackupFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBody) SetTotalFileSize(v int64) *DescribeLogBackupFilesResponseBody {
	s.TotalFileSize = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBody) SetTotalRecordCount(v int32) *DescribeLogBackupFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogBackupFilesResponseBodyItems struct {
	BinLogFile []*DescribeLogBackupFilesResponseBodyItemsBinLogFile `json:"BinLogFile,omitempty" xml:"BinLogFile,omitempty" type:"Repeated"`
}

func (s DescribeLogBackupFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupFilesResponseBodyItems) GetBinLogFile() []*DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	return s.BinLogFile
}

func (s *DescribeLogBackupFilesResponseBodyItems) SetBinLogFile(v []*DescribeLogBackupFilesResponseBodyItemsBinLogFile) *DescribeLogBackupFilesResponseBodyItems {
	s.BinLogFile = v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItems) Validate() error {
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

type DescribeLogBackupFilesResponseBodyItemsBinLogFile struct {
	DownloadLink         *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	FileSize             *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	LinkExpiredTime      *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
	LogBeginTime         *string `json:"LogBeginTime,omitempty" xml:"LogBeginTime,omitempty"`
	LogEndTime           *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty"`
}

func (s DescribeLogBackupFilesResponseBodyItemsBinLogFile) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupFilesResponseBodyItemsBinLogFile) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) GetLogBeginTime() *string {
	return s.LogBeginTime
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) GetLogEndTime() *string {
	return s.LogEndTime
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) SetDownloadLink(v string) *DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	s.DownloadLink = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) SetFileSize(v int64) *DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	s.FileSize = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) SetIntranetDownloadLink(v string) *DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) SetLinkExpiredTime(v string) *DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) SetLogBeginTime(v string) *DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	s.LogBeginTime = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) SetLogEndTime(v string) *DescribeLogBackupFilesResponseBodyItemsBinLogFile {
	s.LogEndTime = &v
	return s
}

func (s *DescribeLogBackupFilesResponseBodyItemsBinLogFile) Validate() error {
	return dara.Validate(s)
}
