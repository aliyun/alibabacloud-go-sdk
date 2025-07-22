// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingOutStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeStreamingOutStatusRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeStreamingOutStatusRequest
	GetChannelId() *string
	SetTaskId(v string) *DescribeStreamingOutStatusRequest
	GetTaskId() *string
}

type DescribeStreamingOutStatusRequest struct {
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

func (s DescribeStreamingOutStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingOutStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamingOutStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeStreamingOutStatusRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeStreamingOutStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeStreamingOutStatusRequest) SetAppId(v string) *DescribeStreamingOutStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeStreamingOutStatusRequest) SetChannelId(v string) *DescribeStreamingOutStatusRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeStreamingOutStatusRequest) SetTaskId(v string) *DescribeStreamingOutStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeStreamingOutStatusRequest) Validate() error {
	return dara.Validate(s)
}
