// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeDownloadRecordsResponseBody
	GetAccessDeniedDetail() *string
	SetRecords(v []*DescribeDownloadRecordsResponseBodyRecords) *DescribeDownloadRecordsResponseBody
	GetRecords() []*DescribeDownloadRecordsResponseBodyRecords
	SetRequestId(v string) *DescribeDownloadRecordsResponseBody
	GetRequestId() *string
}

type DescribeDownloadRecordsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The queried download tasks.
	Records []*DescribeDownloadRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D761DA51-12F8-5457-AAA9-F52B9F436D2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDownloadRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadRecordsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeDownloadRecordsResponseBody) GetRecords() []*DescribeDownloadRecordsResponseBodyRecords {
	return s.Records
}

func (s *DescribeDownloadRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDownloadRecordsResponseBody) SetAccessDeniedDetail(v string) *DescribeDownloadRecordsResponseBody {
	s.AccessDeniedDetail = &v
	return s
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
	// The download job ID.
	//
	// example:
	//
	// 636890
	DownloadId *int64 `json:"DownloadId,omitempty" xml:"DownloadId,omitempty"`
	// The error message returned if the download job failed.
	//
	// example:
	//
	// The query result is empty.
	ExceptionMsg *string `json:"ExceptionMsg,omitempty" xml:"ExceptionMsg,omitempty"`
	// The name of the downloaded file.
	//
	// example:
	//
	// 20210806094635-20210806095135
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The status of the download job. Valid values:
	//
	// 	- **running**
	//
	// 	- **finished**
	//
	// 	- **failed**
	//
	// example:
	//
	// finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The download URL of the file.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

func (s *DescribeDownloadRecordsResponseBodyRecords) GetExceptionMsg() *string {
	return s.ExceptionMsg
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetFileName() *string {
	return s.FileName
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetStatus() *string {
	return s.Status
}

func (s *DescribeDownloadRecordsResponseBodyRecords) GetUrl() *string {
	return s.Url
}

func (s *DescribeDownloadRecordsResponseBodyRecords) SetDownloadId(v int64) *DescribeDownloadRecordsResponseBodyRecords {
	s.DownloadId = &v
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

func (s *DescribeDownloadRecordsResponseBodyRecords) SetUrl(v string) *DescribeDownloadRecordsResponseBodyRecords {
	s.Url = &v
	return s
}

func (s *DescribeDownloadRecordsResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
