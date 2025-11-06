// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyStreamingPropertyRequest
	GetAppId() *string
	SetChannelId(v string) *ModifyStreamingPropertyRequest
	GetChannelId() *string
	SetTaskId(v string) *ModifyStreamingPropertyRequest
	GetTaskId() *string
	SetViewContent(v string) *ModifyStreamingPropertyRequest
	GetViewContent() *string
	SetViewSubscribers(v []*string) *ModifyStreamingPropertyRequest
	GetViewSubscribers() []*string
}

type ModifyStreamingPropertyRequest struct {
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
	// testid
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// main
	ViewContent *string `json:"ViewContent,omitempty" xml:"ViewContent,omitempty"`
	// ViewSubscribers。
	ViewSubscribers []*string `json:"ViewSubscribers,omitempty" xml:"ViewSubscribers,omitempty" type:"Repeated"`
}

func (s ModifyStreamingPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyStreamingPropertyRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyStreamingPropertyRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *ModifyStreamingPropertyRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyStreamingPropertyRequest) GetViewContent() *string {
	return s.ViewContent
}

func (s *ModifyStreamingPropertyRequest) GetViewSubscribers() []*string {
	return s.ViewSubscribers
}

func (s *ModifyStreamingPropertyRequest) SetAppId(v string) *ModifyStreamingPropertyRequest {
	s.AppId = &v
	return s
}

func (s *ModifyStreamingPropertyRequest) SetChannelId(v string) *ModifyStreamingPropertyRequest {
	s.ChannelId = &v
	return s
}

func (s *ModifyStreamingPropertyRequest) SetTaskId(v string) *ModifyStreamingPropertyRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyStreamingPropertyRequest) SetViewContent(v string) *ModifyStreamingPropertyRequest {
	s.ViewContent = &v
	return s
}

func (s *ModifyStreamingPropertyRequest) SetViewSubscribers(v []*string) *ModifyStreamingPropertyRequest {
	s.ViewSubscribers = v
	return s
}

func (s *ModifyStreamingPropertyRequest) Validate() error {
	return dara.Validate(s)
}
