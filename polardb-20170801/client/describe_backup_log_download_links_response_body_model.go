// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLogDownloadLinksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupLogDownloadLinkItem(v []*DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) *DescribeBackupLogDownloadLinksResponseBody
	GetBackupLogDownloadLinkItem() []*DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem
	SetRequestId(v string) *DescribeBackupLogDownloadLinksResponseBody
	GetRequestId() *string
}

type DescribeBackupLogDownloadLinksResponseBody struct {
	// The details of the redo log backups.
	BackupLogDownloadLinkItem []*DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem `json:"BackupLogDownloadLinkItem,omitempty" xml:"BackupLogDownloadLinkItem,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CD35F3-F3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupLogDownloadLinksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogDownloadLinksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogDownloadLinksResponseBody) GetBackupLogDownloadLinkItem() []*DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	return s.BackupLogDownloadLinkItem
}

func (s *DescribeBackupLogDownloadLinksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupLogDownloadLinksResponseBody) SetBackupLogDownloadLinkItem(v []*DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) *DescribeBackupLogDownloadLinksResponseBody {
	s.BackupLogDownloadLinkItem = v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBody) SetRequestId(v string) *DescribeBackupLogDownloadLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBody) Validate() error {
	if s.BackupLogDownloadLinkItem != nil {
		for _, item := range s.BackupLogDownloadLinkItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem struct {
	// The ID of the redo log backup.
	//
	// example:
	//
	// 78167430
	BackupLogId *string `json:"BackupLogId,omitempty" xml:"BackupLogId,omitempty"`
	// The public download link for the redo log backup.
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The error code for the individual record. Empty if the request is successful.
	//
	// example:
	//
	// BackupLogNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The internal network download link for the redo log backup.
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	// The expiration time of the download link in UTC.
	//
	// example:
	//
	// 2026-10-23T03:00:00Z
	LinkExpiredTime *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
	// The error message for the individual record. Empty if the request is successful.
	//
	// example:
	//
	// The specified archived log file does not exist or has been cleaned up.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GetBackupLogId() *string {
	return s.BackupLogId
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GetLinkExpiredTime() *string {
	return s.LinkExpiredTime
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) SetBackupLogId(v string) *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	s.BackupLogId = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) SetDownloadLink(v string) *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) SetErrorCode(v string) *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	s.ErrorCode = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) SetIntranetDownloadLink(v string) *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) SetLinkExpiredTime(v string) *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) SetMessage(v string) *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem {
	s.Message = &v
	return s
}

func (s *DescribeBackupLogDownloadLinksResponseBodyBackupLogDownloadLinkItem) Validate() error {
	return dara.Validate(s)
}
