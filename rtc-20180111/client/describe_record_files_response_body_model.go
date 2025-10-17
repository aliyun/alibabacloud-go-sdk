// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordFiles(v []*DescribeRecordFilesResponseBodyRecordFiles) *DescribeRecordFilesResponseBody
	GetRecordFiles() []*DescribeRecordFilesResponseBodyRecordFiles
	SetRequestId(v string) *DescribeRecordFilesResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *DescribeRecordFilesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeRecordFilesResponseBody
	GetTotalPage() *int64
}

type DescribeRecordFilesResponseBody struct {
	RecordFiles []*DescribeRecordFilesResponseBodyRecordFiles `json:"RecordFiles,omitempty" xml:"RecordFiles,omitempty" type:"Repeated"`
	// example:
	//
	// 760bad53276431c499e30dc36f6b****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeRecordFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponseBody) GetRecordFiles() []*DescribeRecordFilesResponseBodyRecordFiles {
	return s.RecordFiles
}

func (s *DescribeRecordFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordFilesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeRecordFilesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeRecordFilesResponseBody) SetRecordFiles(v []*DescribeRecordFilesResponseBodyRecordFiles) *DescribeRecordFilesResponseBody {
	s.RecordFiles = v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetRequestId(v string) *DescribeRecordFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetTotalNum(v int64) *DescribeRecordFilesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetTotalPage(v int64) *DescribeRecordFilesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) Validate() error {
	if s.RecordFiles != nil {
		for _, item := range s.RecordFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecordFilesResponseBodyRecordFiles struct {
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2020-10-02T17:36:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1800
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2020-11-01T17:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2020-11-02T17:36:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeRecordFilesResponseBodyRecordFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFilesResponseBodyRecordFiles) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetDuration() *float32 {
	return s.Duration
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) GetUrl() *string {
	return s.Url
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetAppId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.AppId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetChannelId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetCreateTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetDuration(v float32) *DescribeRecordFilesResponseBodyRecordFiles {
	s.Duration = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetStartTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetStopTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.StopTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetTaskId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.TaskId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetUrl(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.Url = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) Validate() error {
	return dara.Validate(s)
}
