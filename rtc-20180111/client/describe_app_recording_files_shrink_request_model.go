// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordingFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppRecordingFilesShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeAppRecordingFilesShrinkRequest
	GetChannelId() *string
	SetEndTs(v int64) *DescribeAppRecordingFilesShrinkRequest
	GetEndTs() *int64
	SetPageNo(v int32) *DescribeAppRecordingFilesShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAppRecordingFilesShrinkRequest
	GetPageSize() *int32
	SetStartTs(v int64) *DescribeAppRecordingFilesShrinkRequest
	GetStartTs() *int64
	SetTaskIdsShrink(v string) *DescribeAppRecordingFilesShrinkRequest
	GetTaskIdsShrink() *string
}

type DescribeAppRecordingFilesShrinkRequest struct {
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
	StartTs       *int64  `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	TaskIdsShrink *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DescribeAppRecordingFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordingFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeAppRecordingFilesShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetAppId(v string) *DescribeAppRecordingFilesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetChannelId(v string) *DescribeAppRecordingFilesShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetEndTs(v int64) *DescribeAppRecordingFilesShrinkRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetPageNo(v int32) *DescribeAppRecordingFilesShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetPageSize(v int32) *DescribeAppRecordingFilesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetStartTs(v int64) *DescribeAppRecordingFilesShrinkRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) SetTaskIdsShrink(v string) *DescribeAppRecordingFilesShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *DescribeAppRecordingFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
