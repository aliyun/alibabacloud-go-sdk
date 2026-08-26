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
	// The ID of the application to subscribe to. You can view your application IDs by navigating to **ApsaraVideo Live > Live+ > ApsaraVideo Real-time Communication > Application Management**. If no application exists, create one by clicking **Create Application**.
	//
	// > The application ID consists of uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The callback URL. For the URL format, refer to the callback content specifications below.
	//
	// > The callback URL protocol must be HTTP or HTTPS. The URL can contain only the following characters: a-z, A-Z, 0-9, -, _, ?, %, =, #, ., /, and +. The URL cannot exceed 2083 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The channel IDs of the stream mixing tasks for which you want to receive callbacks. You can specify multiple channel IDs separated by commas (,).
	//
	// >- If you leave this parameter empty, callbacks for all stream mixing and relaying tasks under the specified AppId are received by default.
	//
	// - When specifying multiple channel IDs, do not include duplicates. You can specify up to 20 channel IDs at a time.
	//
	// - Each channel ID consists of uppercase and lowercase letters, digits, underscores, and hyphens (-), with a maximum of 64 characters.
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
