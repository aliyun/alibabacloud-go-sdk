// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCloudNotesShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeCloudNotesShrinkRequest
	GetChannelId() *string
	SetEndTs(v int64) *DescribeCloudNotesShrinkRequest
	GetEndTs() *int64
	SetPageNo(v int32) *DescribeCloudNotesShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeCloudNotesShrinkRequest
	GetPageSize() *int32
	SetStartTs(v int64) *DescribeCloudNotesShrinkRequest
	GetStartTs() *int64
	SetTaskIdsShrink(v string) *DescribeCloudNotesShrinkRequest
	GetTaskIdsShrink() *string
}

type DescribeCloudNotesShrinkRequest struct {
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

func (s DescribeCloudNotesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudNotesShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCloudNotesShrinkRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeCloudNotesShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCloudNotesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudNotesShrinkRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeCloudNotesShrinkRequest) GetTaskIdsShrink() *string {
	return s.TaskIdsShrink
}

func (s *DescribeCloudNotesShrinkRequest) SetAppId(v string) *DescribeCloudNotesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) SetChannelId(v string) *DescribeCloudNotesShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) SetEndTs(v int64) *DescribeCloudNotesShrinkRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) SetPageNo(v int32) *DescribeCloudNotesShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) SetPageSize(v int32) *DescribeCloudNotesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) SetStartTs(v int64) *DescribeCloudNotesShrinkRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) SetTaskIdsShrink(v string) *DescribeCloudNotesShrinkRequest {
	s.TaskIdsShrink = &v
	return s
}

func (s *DescribeCloudNotesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
