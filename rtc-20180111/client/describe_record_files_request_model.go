// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRecordFilesRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeRecordFilesRequest
	GetChannelId() *string
	SetEndTime(v string) *DescribeRecordFilesRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeRecordFilesRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeRecordFilesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeRecordFilesRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeRecordFilesRequest
	GetStartTime() *string
	SetTaskIds(v []*string) *DescribeRecordFilesRequest
	GetTaskIds() []*string
}

type DescribeRecordFilesRequest struct {
	// This parameter is required.
	//
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
	// 2020-11-02T17:36:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2020-11-01T17:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// yourTaskId
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRecordFilesRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRecordFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRecordFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRecordFilesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeRecordFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecordFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordFilesRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DescribeRecordFilesRequest) SetAppId(v string) *DescribeRecordFilesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetChannelId(v string) *DescribeRecordFilesRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetEndTime(v string) *DescribeRecordFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetOwnerId(v int64) *DescribeRecordFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetPageNum(v int32) *DescribeRecordFilesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetPageSize(v int32) *DescribeRecordFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetStartTime(v string) *DescribeRecordFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetTaskIds(v []*string) *DescribeRecordFilesRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeRecordFilesRequest) Validate() error {
	return dara.Validate(s)
}
