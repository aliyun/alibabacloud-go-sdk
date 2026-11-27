// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordingFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAppRecordingFilesResponseBodyItems) *DescribeAppRecordingFilesResponseBody
	GetItems() []*DescribeAppRecordingFilesResponseBodyItems
	SetPageNo(v int32) *DescribeAppRecordingFilesResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAppRecordingFilesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAppRecordingFilesResponseBody
	GetRequestId() *string
	SetTotalCnt(v int32) *DescribeAppRecordingFilesResponseBody
	GetTotalCnt() *int32
}

type DescribeAppRecordingFilesResponseBody struct {
	// List of recording files.
	Items []*DescribeAppRecordingFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count.
	//
	// example:
	//
	// 10
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeAppRecordingFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordingFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordingFilesResponseBody) GetItems() []*DescribeAppRecordingFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeAppRecordingFilesResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAppRecordingFilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppRecordingFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppRecordingFilesResponseBody) GetTotalCnt() *int32 {
	return s.TotalCnt
}

func (s *DescribeAppRecordingFilesResponseBody) SetItems(v []*DescribeAppRecordingFilesResponseBodyItems) *DescribeAppRecordingFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAppRecordingFilesResponseBody) SetPageNo(v int32) *DescribeAppRecordingFilesResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBody) SetPageSize(v int32) *DescribeAppRecordingFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBody) SetRequestId(v string) *DescribeAppRecordingFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBody) SetTotalCnt(v int32) *DescribeAppRecordingFilesResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppRecordingFilesResponseBodyItems struct {
	// User-defined bucket.
	//
	// example:
	//
	// rtc-bucket
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Channel ID.
	//
	// example:
	//
	// testchannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Recording file creation timestamp, in milliseconds.
	//
	// example:
	//
	// 1712376032000
	FileCreateTs *int64 `json:"FileCreateTs,omitempty" xml:"FileCreateTs,omitempty"`
	// Recording file duration, in milliseconds.
	//
	// example:
	//
	// 200
	FileDuration *int32 `json:"FileDuration,omitempty" xml:"FileDuration,omitempty"`
	// Path where the recording file is stored. Use the SDK of your cloud provider\"s Object Storage Service to access the file. For Alibaba Cloud OSS, see [here](https://help.aliyun.com/zh/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects).
	//
	// example:
	//
	// record/appid/12_task_local1/1712279809158_1712279844691/playlist.mp4
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// Recording file size, in bytes.
	//
	// example:
	//
	// 10000
	FileSize *int32 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// See region used in the recording task.
	//
	// example:
	//
	// 1
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// Recording task start timestamp, in milliseconds.
	//
	// example:
	//
	// 1712376012000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// Recording task ID.
	//
	// example:
	//
	// test001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// See vendor used when starting the recording task.
	//
	// example:
	//
	// 1
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeAppRecordingFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordingFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetFileCreateTs() *int64 {
	return s.FileCreateTs
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetFileDuration() *int32 {
	return s.FileDuration
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetFileSize() *int32 {
	return s.FileSize
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetRegion() *int32 {
	return s.Region
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAppRecordingFilesResponseBodyItems) GetVendor() *int32 {
	return s.Vendor
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetBucket(v string) *DescribeAppRecordingFilesResponseBodyItems {
	s.Bucket = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetChannelId(v string) *DescribeAppRecordingFilesResponseBodyItems {
	s.ChannelId = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetFileCreateTs(v int64) *DescribeAppRecordingFilesResponseBodyItems {
	s.FileCreateTs = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetFileDuration(v int32) *DescribeAppRecordingFilesResponseBodyItems {
	s.FileDuration = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetFilePath(v string) *DescribeAppRecordingFilesResponseBodyItems {
	s.FilePath = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetFileSize(v int32) *DescribeAppRecordingFilesResponseBodyItems {
	s.FileSize = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetRegion(v int32) *DescribeAppRecordingFilesResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetStartTs(v int64) *DescribeAppRecordingFilesResponseBodyItems {
	s.StartTs = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetTaskId(v string) *DescribeAppRecordingFilesResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) SetVendor(v int32) *DescribeAppRecordingFilesResponseBodyItems {
	s.Vendor = &v
	return s
}

func (s *DescribeAppRecordingFilesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
