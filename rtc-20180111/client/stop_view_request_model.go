// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopViewRequest
	GetAppId() *string
	SetChannelId(v string) *StopViewRequest
	GetChannelId() *string
	SetTaskId(v string) *StopViewRequest
	GetTaskId() *string
}

type StopViewRequest struct {
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

func (s StopViewRequest) String() string {
	return dara.Prettify(s)
}

func (s StopViewRequest) GoString() string {
	return s.String()
}

func (s *StopViewRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopViewRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopViewRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopViewRequest) SetAppId(v string) *StopViewRequest {
	s.AppId = &v
	return s
}

func (s *StopViewRequest) SetChannelId(v string) *StopViewRequest {
	s.ChannelId = &v
	return s
}

func (s *StopViewRequest) SetTaskId(v string) *StopViewRequest {
	s.TaskId = &v
	return s
}

func (s *StopViewRequest) Validate() error {
	return dara.Validate(s)
}
