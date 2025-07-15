// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcMPUEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateRtcMPUEventSubRequest
	GetAppId() *string
	SetCallbackUrl(v string) *UpdateRtcMPUEventSubRequest
	GetCallbackUrl() *string
	SetChannelIds(v string) *UpdateRtcMPUEventSubRequest
	GetChannelIds() *string
}

type UpdateRtcMPUEventSubRequest struct {
	// The ID of the application.
	//
	// >  The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL.
	//
	// >  You can use headers such as HTTP and HTTPS in callback URLs. The URL can be up to 2,083 characters and contain letters, digits, and the following special characters: - _ ? % = # . / +
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The ID of the channel to which you want to send mixed-stream relay event callbacks. Separate multiple channel IDs with commas (,).
	//
	// >
	//
	// 	- If you leave this parameter empty, you are subscribed to all mixed-stream relay events submitted in the application.
	//
	// 	- You cannot specify duplicate channel IDs. You can specify up to 20 channel IDs in each call.
	//
	// 	- The ID can be up to 64 characters in length and contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// yourChannelIds
	ChannelIds *string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty"`
}

func (s UpdateRtcMPUEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcMPUEventSubRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtcMPUEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateRtcMPUEventSubRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *UpdateRtcMPUEventSubRequest) GetChannelIds() *string {
	return s.ChannelIds
}

func (s *UpdateRtcMPUEventSubRequest) SetAppId(v string) *UpdateRtcMPUEventSubRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRtcMPUEventSubRequest) SetCallbackUrl(v string) *UpdateRtcMPUEventSubRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UpdateRtcMPUEventSubRequest) SetChannelIds(v string) *UpdateRtcMPUEventSubRequest {
	s.ChannelIds = &v
	return s
}

func (s *UpdateRtcMPUEventSubRequest) Validate() error {
	return dara.Validate(s)
}
