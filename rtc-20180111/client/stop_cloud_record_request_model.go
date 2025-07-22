// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopCloudRecordRequest
	GetAppId() *string
	SetChannelId(v string) *StopCloudRecordRequest
	GetChannelId() *string
	SetTaskId(v string) *StopCloudRecordRequest
	GetTaskId() *string
}

type StopCloudRecordRequest struct {
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

func (s StopCloudRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StopCloudRecordRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopCloudRecordRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopCloudRecordRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopCloudRecordRequest) SetAppId(v string) *StopCloudRecordRequest {
	s.AppId = &v
	return s
}

func (s *StopCloudRecordRequest) SetChannelId(v string) *StopCloudRecordRequest {
	s.ChannelId = &v
	return s
}

func (s *StopCloudRecordRequest) SetTaskId(v string) *StopCloudRecordRequest {
	s.TaskId = &v
	return s
}

func (s *StopCloudRecordRequest) Validate() error {
	return dara.Validate(s)
}
