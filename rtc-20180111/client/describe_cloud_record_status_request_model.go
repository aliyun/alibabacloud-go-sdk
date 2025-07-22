// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudRecordStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCloudRecordStatusRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeCloudRecordStatusRequest
	GetChannelId() *string
	SetTaskId(v string) *DescribeCloudRecordStatusRequest
	GetTaskId() *string
}

type DescribeCloudRecordStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1qaz***x
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// taskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeCloudRecordStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudRecordStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudRecordStatusRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeCloudRecordStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeCloudRecordStatusRequest) SetAppId(v string) *DescribeCloudRecordStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCloudRecordStatusRequest) SetChannelId(v string) *DescribeCloudRecordStatusRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCloudRecordStatusRequest) SetTaskId(v string) *DescribeCloudRecordStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeCloudRecordStatusRequest) Validate() error {
	return dara.Validate(s)
}
