// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingPropertyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyStreamingPropertyShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *ModifyStreamingPropertyShrinkRequest
	GetChannelId() *string
	SetTaskId(v string) *ModifyStreamingPropertyShrinkRequest
	GetTaskId() *string
	SetViewContent(v string) *ModifyStreamingPropertyShrinkRequest
	GetViewContent() *string
	SetViewSubscribersShrink(v string) *ModifyStreamingPropertyShrinkRequest
	GetViewSubscribersShrink() *string
}

type ModifyStreamingPropertyShrinkRequest struct {
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
	ViewSubscribersShrink *string `json:"ViewSubscribers,omitempty" xml:"ViewSubscribers,omitempty"`
}

func (s ModifyStreamingPropertyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingPropertyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyStreamingPropertyShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyStreamingPropertyShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *ModifyStreamingPropertyShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyStreamingPropertyShrinkRequest) GetViewContent() *string {
	return s.ViewContent
}

func (s *ModifyStreamingPropertyShrinkRequest) GetViewSubscribersShrink() *string {
	return s.ViewSubscribersShrink
}

func (s *ModifyStreamingPropertyShrinkRequest) SetAppId(v string) *ModifyStreamingPropertyShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyStreamingPropertyShrinkRequest) SetChannelId(v string) *ModifyStreamingPropertyShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *ModifyStreamingPropertyShrinkRequest) SetTaskId(v string) *ModifyStreamingPropertyShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyStreamingPropertyShrinkRequest) SetViewContent(v string) *ModifyStreamingPropertyShrinkRequest {
	s.ViewContent = &v
	return s
}

func (s *ModifyStreamingPropertyShrinkRequest) SetViewSubscribersShrink(v string) *ModifyStreamingPropertyShrinkRequest {
	s.ViewSubscribersShrink = &v
	return s
}

func (s *ModifyStreamingPropertyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
