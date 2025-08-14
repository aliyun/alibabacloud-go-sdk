// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPolicy(v bool) *CreateChannelRequest
	GetAccessPolicy() *bool
	SetAccessToken(v string) *CreateChannelRequest
	GetAccessToken() *string
	SetChannelName(v string) *CreateChannelRequest
	GetChannelName() *string
	SetChannelTier(v string) *CreateChannelRequest
	GetChannelTier() *string
	SetFillerSourceLocationName(v string) *CreateChannelRequest
	GetFillerSourceLocationName() *string
	SetFillerSourceName(v string) *CreateChannelRequest
	GetFillerSourceName() *string
	SetOutPutConfigList(v string) *CreateChannelRequest
	GetOutPutConfigList() *string
	SetPlaybackMode(v string) *CreateChannelRequest
	GetPlaybackMode() *string
}

type CreateChannelRequest struct {
	// Specifies whether to enable access control.
	//
	// example:
	//
	// false
	AccessPolicy *bool `json:"AccessPolicy,omitempty" xml:"AccessPolicy,omitempty"`
	// The token for accessing the channel.
	//
	// example:
	//
	// xxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The tier of the channel. Valid values: basic and standard.
	//
	// This parameter is required.
	//
	// example:
	//
	// basic
	ChannelTier *string `json:"ChannelTier,omitempty" xml:"ChannelTier,omitempty"`
	// The source location of the filler slate.
	//
	// example:
	//
	// MySourceLocation
	FillerSourceLocationName *string `json:"FillerSourceLocationName,omitempty" xml:"FillerSourceLocationName,omitempty"`
	// The name of the filler slate.
	//
	// example:
	//
	// FillerSource
	FillerSourceName *string `json:"FillerSourceName,omitempty" xml:"FillerSourceName,omitempty"`
	// The channel output configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{
	//
	// 	"ManifestName": "manifest-1",
	//
	// 	"Format": "HLS",
	//
	// 	"SourceGroupName": "source-group-1",
	//
	// 	"ManifestSettings": {
	//
	// 		"WindowDuration": 60,
	//
	// 		"AdMarkType": "Daterange"
	//
	// 	}
	//
	// }]
	OutPutConfigList *string `json:"OutPutConfigList,omitempty" xml:"OutPutConfigList,omitempty"`
	// The playback mode. Valid values: loop and linear.
	//
	// This parameter is required.
	//
	// example:
	//
	// loop
	PlaybackMode *string `json:"PlaybackMode,omitempty" xml:"PlaybackMode,omitempty"`
}

func (s CreateChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateChannelRequest) GetAccessPolicy() *bool {
	return s.AccessPolicy
}

func (s *CreateChannelRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateChannelRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateChannelRequest) GetChannelTier() *string {
	return s.ChannelTier
}

func (s *CreateChannelRequest) GetFillerSourceLocationName() *string {
	return s.FillerSourceLocationName
}

func (s *CreateChannelRequest) GetFillerSourceName() *string {
	return s.FillerSourceName
}

func (s *CreateChannelRequest) GetOutPutConfigList() *string {
	return s.OutPutConfigList
}

func (s *CreateChannelRequest) GetPlaybackMode() *string {
	return s.PlaybackMode
}

func (s *CreateChannelRequest) SetAccessPolicy(v bool) *CreateChannelRequest {
	s.AccessPolicy = &v
	return s
}

func (s *CreateChannelRequest) SetAccessToken(v string) *CreateChannelRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateChannelRequest) SetChannelName(v string) *CreateChannelRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateChannelRequest) SetChannelTier(v string) *CreateChannelRequest {
	s.ChannelTier = &v
	return s
}

func (s *CreateChannelRequest) SetFillerSourceLocationName(v string) *CreateChannelRequest {
	s.FillerSourceLocationName = &v
	return s
}

func (s *CreateChannelRequest) SetFillerSourceName(v string) *CreateChannelRequest {
	s.FillerSourceName = &v
	return s
}

func (s *CreateChannelRequest) SetOutPutConfigList(v string) *CreateChannelRequest {
	s.OutPutConfigList = &v
	return s
}

func (s *CreateChannelRequest) SetPlaybackMode(v string) *CreateChannelRequest {
	s.PlaybackMode = &v
	return s
}

func (s *CreateChannelRequest) Validate() error {
	return dara.Validate(s)
}
