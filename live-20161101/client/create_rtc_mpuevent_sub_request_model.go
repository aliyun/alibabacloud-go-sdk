// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtcMPUEventSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateRtcMPUEventSubRequest
	GetAppId() *string
	SetCallbackUrl(v string) *CreateRtcMPUEventSubRequest
	GetCallbackUrl() *string
	SetChannelIds(v string) *CreateRtcMPUEventSubRequest
	GetChannelIds() *string
}

type CreateRtcMPUEventSubRequest struct {
	// The ID of the application.
	//
	// > The ID can be up to 64 characters in length and can contain letters, digits, underscores, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL.
	//
	// > The callback URL can be up to 2,083 characters in length. You can use headers such as HTTP and HTTPS in callback URLs. The URL can contain letters, digits, and the following special characters: - _ ? % = # . / +
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The ID of the channel to which you want to send mixed-stream relay event callbacks. Separate multiple channel IDs with commas (,).
	//
	// >
	//
	// 	- If you leave this parameter empty, you are subscribed to mixed-stream relay events of all channels in the application.
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

func (s CreateRtcMPUEventSubRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRtcMPUEventSubRequest) GoString() string {
	return s.String()
}

func (s *CreateRtcMPUEventSubRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateRtcMPUEventSubRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *CreateRtcMPUEventSubRequest) GetChannelIds() *string {
	return s.ChannelIds
}

func (s *CreateRtcMPUEventSubRequest) SetAppId(v string) *CreateRtcMPUEventSubRequest {
	s.AppId = &v
	return s
}

func (s *CreateRtcMPUEventSubRequest) SetCallbackUrl(v string) *CreateRtcMPUEventSubRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateRtcMPUEventSubRequest) SetChannelIds(v string) *CreateRtcMPUEventSubRequest {
	s.ChannelIds = &v
	return s
}

func (s *CreateRtcMPUEventSubRequest) Validate() error {
	return dara.Validate(s)
}
