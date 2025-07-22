// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStreamingOutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopStreamingOutRequest
	GetAppId() *string
	SetChannelId(v string) *StopStreamingOutRequest
	GetChannelId() *string
	SetTaskId(v string) *StopStreamingOutRequest
	GetTaskId() *string
}

type StopStreamingOutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopStreamingOutRequest) String() string {
	return dara.Prettify(s)
}

func (s StopStreamingOutRequest) GoString() string {
	return s.String()
}

func (s *StopStreamingOutRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopStreamingOutRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopStreamingOutRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopStreamingOutRequest) SetAppId(v string) *StopStreamingOutRequest {
	s.AppId = &v
	return s
}

func (s *StopStreamingOutRequest) SetChannelId(v string) *StopStreamingOutRequest {
	s.ChannelId = &v
	return s
}

func (s *StopStreamingOutRequest) SetTaskId(v string) *StopStreamingOutRequest {
	s.TaskId = &v
	return s
}

func (s *StopStreamingOutRequest) Validate() error {
	return dara.Validate(s)
}
