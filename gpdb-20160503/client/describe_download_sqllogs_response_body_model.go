// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadSQLLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*DescribeDownloadSQLLogsResponseBodyRecords) *DescribeDownloadSQLLogsResponseBody
	GetRecords() []*DescribeDownloadSQLLogsResponseBodyRecords
	SetRequestId(v string) *DescribeDownloadSQLLogsResponseBody
	GetRequestId() *string
}

type DescribeDownloadSQLLogsResponseBody struct {
	// List of download records.
	Records []*DescribeDownloadSQLLogsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// FDE9942A-D919-527B-B559-5D0F6F20CCEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadSQLLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSQLLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsResponseBody) GetRecords() []*DescribeDownloadSQLLogsResponseBodyRecords {
	return s.Records
}

func (s *DescribeDownloadSQLLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadSQLLogsResponseBody) SetRecords(v []*DescribeDownloadSQLLogsResponseBodyRecords) *DescribeDownloadSQLLogsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBody) SetRequestId(v string) *DescribeDownloadSQLLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDownloadSQLLogsResponseBodyRecords struct {
	// Download record ID.
	//
	// example:
	//
	// 1150
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// Download link.
	//
	// example:
	//
	// https://perth-download-task.oss-cn-beijing.aliyuncs.com/*****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// Error message.
	//
	// example:
	//
	// Error message
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// File name.
	//
	// example:
	//
	// 20220509113448-20220509173448.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Task status, with possible values being:
	//
	// - **running**: Downloading.
	//
	// - **finished**: Completed.
	//
	// - **failed**: Download failed.
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDownloadSQLLogsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadSQLLogsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) GetDownloadId() *int64 {
	return s.DownloadId
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) GetExceptionMsg() *string {
	return s.ExceptionMsg
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) GetFileName() *string {
	return s.FileName
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) GetStatus() *string {
	return s.Status
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetDownloadUrl(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetFileName(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) SetStatus(v string) *DescribeDownloadSQLLogsResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *DescribeDownloadSQLLogsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
