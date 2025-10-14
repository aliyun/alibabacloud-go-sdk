// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinaryLogListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogList(v []*DescribeBinaryLogListResponseBodyLogList) *DescribeBinaryLogListResponseBody
	GetLogList() []*DescribeBinaryLogListResponseBodyLogList
	SetPageNumber(v int32) *DescribeBinaryLogListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBinaryLogListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBinaryLogListResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *DescribeBinaryLogListResponseBody
	GetTotalNumber() *int32
}

type DescribeBinaryLogListResponseBody struct {
	LogList []*DescribeBinaryLogListResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2DFF784E-DC31-5BBC-9B25-9261CD9E0AA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeBinaryLogListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinaryLogListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponseBody) GetLogList() []*DescribeBinaryLogListResponseBodyLogList {
	return s.LogList
}

func (s *DescribeBinaryLogListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBinaryLogListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBinaryLogListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBinaryLogListResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *DescribeBinaryLogListResponseBody) SetLogList(v []*DescribeBinaryLogListResponseBodyLogList) *DescribeBinaryLogListResponseBody {
	s.LogList = v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetPageNumber(v int32) *DescribeBinaryLogListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetPageSize(v int32) *DescribeBinaryLogListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetRequestId(v string) *DescribeBinaryLogListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) SetTotalNumber(v int32) *DescribeBinaryLogListResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *DescribeBinaryLogListResponseBody) Validate() error {
	if s.LogList != nil {
		for _, item := range s.LogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBinaryLogListResponseBodyLogList struct {
	// example:
	//
	// 2021-09-09 10:27:46
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 2021-09-09 10:27:46
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// http://polarx-cdc-binlog-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/polardbx_cdc/pxc-hzfd132143sfds1/binlog.000001?Expires=1636469502&OSSAccessKeyId=LT13fds12dsafddsf&Signature=fdpm%bdsfadsa%2F%bdsafdsaf%3D
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// example:
	//
	// 2021-11-09 10:27:46
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// binlog.000001
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 536870912
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// example:
	//
	// 2021-11-09 10:27:46
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 0
	PurgeStatus *int32 `json:"PurgeStatus,omitempty" xml:"PurgeStatus,omitempty"`
	// example:
	//
	// 10.110.88.9
	UploadHost *string `json:"UploadHost,omitempty" xml:"UploadHost,omitempty"`
	// example:
	//
	// 2
	UploadStatus *int32 `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
}

func (s DescribeBinaryLogListResponseBodyLogList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinaryLogListResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetFileName() *string {
	return s.FileName
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetId() *int64 {
	return s.Id
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetPurgeStatus() *int32 {
	return s.PurgeStatus
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetUploadHost() *string {
	return s.UploadHost
}

func (s *DescribeBinaryLogListResponseBodyLogList) GetUploadStatus() *int32 {
	return s.UploadStatus
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetBeginTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.BeginTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetCreatedTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetDownloadLink(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetEndTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.EndTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetFileName(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.FileName = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetId(v int64) *DescribeBinaryLogListResponseBodyLogList {
	s.Id = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetLogSize(v int64) *DescribeBinaryLogListResponseBodyLogList {
	s.LogSize = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetModifiedTime(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetPurgeStatus(v int32) *DescribeBinaryLogListResponseBodyLogList {
	s.PurgeStatus = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetUploadHost(v string) *DescribeBinaryLogListResponseBodyLogList {
	s.UploadHost = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) SetUploadStatus(v int32) *DescribeBinaryLogListResponseBodyLogList {
	s.UploadStatus = &v
	return s
}

func (s *DescribeBinaryLogListResponseBodyLogList) Validate() error {
	return dara.Validate(s)
}
