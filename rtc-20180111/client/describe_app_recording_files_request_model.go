// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordingFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppRecordingFilesRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeAppRecordingFilesRequest
	GetChannelId() *string
	SetEndTs(v int64) *DescribeAppRecordingFilesRequest
	GetEndTs() *int64
	SetPageNo(v int32) *DescribeAppRecordingFilesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAppRecordingFilesRequest
	GetPageSize() *int32
	SetStartTs(v int64) *DescribeAppRecordingFilesRequest
	GetStartTs() *int64
	SetTaskIds(v []*string) *DescribeAppRecordingFilesRequest
	GetTaskIds() []*string
}

type DescribeAppRecordingFilesRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1712376532000
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1712376032000
	StartTs *int64    `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeAppRecordingFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordingFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordingFilesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppRecordingFilesRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeAppRecordingFilesRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeAppRecordingFilesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAppRecordingFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppRecordingFilesRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeAppRecordingFilesRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DescribeAppRecordingFilesRequest) SetAppId(v string) *DescribeAppRecordingFilesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppRecordingFilesRequest) SetChannelId(v string) *DescribeAppRecordingFilesRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeAppRecordingFilesRequest) SetEndTs(v int64) *DescribeAppRecordingFilesRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeAppRecordingFilesRequest) SetPageNo(v int32) *DescribeAppRecordingFilesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAppRecordingFilesRequest) SetPageSize(v int32) *DescribeAppRecordingFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppRecordingFilesRequest) SetStartTs(v int64) *DescribeAppRecordingFilesRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeAppRecordingFilesRequest) SetTaskIds(v []*string) *DescribeAppRecordingFilesRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeAppRecordingFilesRequest) Validate() error {
	return dara.Validate(s)
}
