// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCloudNotesRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeCloudNotesRequest
	GetChannelId() *string
	SetEndTs(v int64) *DescribeCloudNotesRequest
	GetEndTs() *int64
	SetPageNo(v int32) *DescribeCloudNotesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeCloudNotesRequest
	GetPageSize() *int32
	SetStartTs(v int64) *DescribeCloudNotesRequest
	GetStartTs() *int64
	SetTaskIds(v []*string) *DescribeCloudNotesRequest
	GetTaskIds() []*string
}

type DescribeCloudNotesRequest struct {
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

func (s DescribeCloudNotesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudNotesRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCloudNotesRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeCloudNotesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeCloudNotesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudNotesRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeCloudNotesRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *DescribeCloudNotesRequest) SetAppId(v string) *DescribeCloudNotesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCloudNotesRequest) SetChannelId(v string) *DescribeCloudNotesRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCloudNotesRequest) SetEndTs(v int64) *DescribeCloudNotesRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeCloudNotesRequest) SetPageNo(v int32) *DescribeCloudNotesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCloudNotesRequest) SetPageSize(v int32) *DescribeCloudNotesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudNotesRequest) SetStartTs(v int64) *DescribeCloudNotesRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeCloudNotesRequest) SetTaskIds(v []*string) *DescribeCloudNotesRequest {
	s.TaskIds = v
	return s
}

func (s *DescribeCloudNotesRequest) Validate() error {
	return dara.Validate(s)
}
