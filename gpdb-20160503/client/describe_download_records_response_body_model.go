// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody
	GetRecords() []*DescribeDownloadRecordsResponseBodyRecords
	SetRequestId(v string) *DescribeDownloadRecordsResponseBody
	GetRequestId() *string
}

type DescribeDownloadRecordsResponseBody struct {
	// The URL that is used to download the file.
	Records []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) GetRecords() []*DescribeDownloadRecordsResponseBodyRecords {
	return s.Records
}

func (s *DescribeDownloadRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadRecordsResponseBody) SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) SetRequestId(v string) *DescribeDownloadRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDownloadRecordsResponseBodyRecords struct {
	// The ID of the download record.
	//
	// example:
	//
	// 1150
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The URL that can be used to download the file.
	//
	// example:
	//
	// https://perth-download-task.oss-cn-beijing.aliyuncs.com/*****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Error message
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// 20220509113448-20220509173448.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The state of the upload task. After you call the DownloadDiagnosisRecords operation, query diagnostic information is first uploaded to Object Storage Service (OSS). After the upload task is complete, the query diagnostic information can be downloaded. Valid values:
	//
	// 	- **running**: uploading
	//
	// 	- **finished**: uploaded
	//
	// 	- **failed**: failed
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDownloadRecordsResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetDownloadId() *int64 {
	return s.DownloadId
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetExceptionMsg() *string {
	return s.ExceptionMsg
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetFileName() *string {
	return s.FileName
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetStatus() *string {
	return s.Status
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetExceptionMsg(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.ExceptionMsg = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetFileName(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.FileName = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetStatus(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Status = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
