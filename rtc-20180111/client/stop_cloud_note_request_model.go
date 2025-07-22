// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudNoteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopCloudNoteRequest
	GetAppId() *string
	SetChannelId(v string) *StopCloudNoteRequest
	GetChannelId() *string
	SetClientToken(v string) *StopCloudNoteRequest
	GetClientToken() *string
	SetTaskId(v string) *StopCloudNoteRequest
	GetTaskId() *string
}

type StopCloudNoteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qwsz1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopCloudNoteRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCloudNoteRequest) GoString() string {
	return s.String()
}

func (s *StopCloudNoteRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopCloudNoteRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StopCloudNoteRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopCloudNoteRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopCloudNoteRequest) SetAppId(v string) *StopCloudNoteRequest {
	s.AppId = &v
	return s
}

func (s *StopCloudNoteRequest) SetChannelId(v string) *StopCloudNoteRequest {
	s.ChannelId = &v
	return s
}

func (s *StopCloudNoteRequest) SetClientToken(v string) *StopCloudNoteRequest {
	s.ClientToken = &v
	return s
}

func (s *StopCloudNoteRequest) SetTaskId(v string) *StopCloudNoteRequest {
	s.TaskId = &v
	return s
}

func (s *StopCloudNoteRequest) Validate() error {
	return dara.Validate(s)
}
