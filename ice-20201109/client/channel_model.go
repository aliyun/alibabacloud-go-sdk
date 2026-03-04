// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPolicy(v bool) *Channel
	GetAccessPolicy() *bool
	SetAccessToken(v string) *Channel
	GetAccessToken() *string
	SetArn(v string) *Channel
	GetArn() *string
	SetChannelName(v string) *Channel
	GetChannelName() *string
	SetChannelTier(v string) *Channel
	GetChannelTier() *string
	SetFillerSourceLocationName(v string) *Channel
	GetFillerSourceLocationName() *string
	SetFillerSourceName(v string) *Channel
	GetFillerSourceName() *string
	SetGmtCreate(v string) *Channel
	GetGmtCreate() *string
	SetGmtModified(v string) *Channel
	GetGmtModified() *string
	SetOutPutConfigList(v []*ChannelOutPutConfigList) *Channel
	GetOutPutConfigList() []*ChannelOutPutConfigList
	SetPlaybackMode(v string) *Channel
	GetPlaybackMode() *string
	SetState(v int32) *Channel
	GetState() *int32
}

type Channel struct {
	// Specifies whether to enable access control.
	//
	// example:
	//
	// true
	AccessPolicy *bool `json:"AccessPolicy,omitempty" xml:"AccessPolicy,omitempty"`
	// The token for accessing the channel.
	//
	// example:
	//
	// xxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The ARN of the channel.
	//
	// example:
	//
	// acs:ims:mediaweaver:<regionId>:<userId>:channel/myChannel
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The name of the channel.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The tier of the channel.
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
	// The source name of the filler slate.
	//
	// example:
	//
	// MySource
	FillerSourceName *string `json:"FillerSourceName,omitempty" xml:"FillerSourceName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-01-15T03:44:16Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2024-09-01T10:09Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The output configurations.
	OutPutConfigList []*ChannelOutPutConfigList `json:"OutPutConfigList,omitempty" xml:"OutPutConfigList,omitempty" type:"Repeated"`
	// The playback mode.
	//
	// example:
	//
	// loop
	PlaybackMode *string `json:"PlaybackMode,omitempty" xml:"PlaybackMode,omitempty"`
	// The state of the channel.
	//
	// example:
	//
	// 0
	State *int32 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s Channel) String() string {
	return dara.Prettify(s)
}

func (s Channel) GoString() string {
	return s.String()
}

func (s *Channel) GetAccessPolicy() *bool {
	return s.AccessPolicy
}

func (s *Channel) GetAccessToken() *string {
	return s.AccessToken
}

func (s *Channel) GetArn() *string {
	return s.Arn
}

func (s *Channel) GetChannelName() *string {
	return s.ChannelName
}

func (s *Channel) GetChannelTier() *string {
	return s.ChannelTier
}

func (s *Channel) GetFillerSourceLocationName() *string {
	return s.FillerSourceLocationName
}

func (s *Channel) GetFillerSourceName() *string {
	return s.FillerSourceName
}

func (s *Channel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Channel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Channel) GetOutPutConfigList() []*ChannelOutPutConfigList {
	return s.OutPutConfigList
}

func (s *Channel) GetPlaybackMode() *string {
	return s.PlaybackMode
}

func (s *Channel) GetState() *int32 {
	return s.State
}

func (s *Channel) SetAccessPolicy(v bool) *Channel {
	s.AccessPolicy = &v
	return s
}

func (s *Channel) SetAccessToken(v string) *Channel {
	s.AccessToken = &v
	return s
}

func (s *Channel) SetArn(v string) *Channel {
	s.Arn = &v
	return s
}

func (s *Channel) SetChannelName(v string) *Channel {
	s.ChannelName = &v
	return s
}

func (s *Channel) SetChannelTier(v string) *Channel {
	s.ChannelTier = &v
	return s
}

func (s *Channel) SetFillerSourceLocationName(v string) *Channel {
	s.FillerSourceLocationName = &v
	return s
}

func (s *Channel) SetFillerSourceName(v string) *Channel {
	s.FillerSourceName = &v
	return s
}

func (s *Channel) SetGmtCreate(v string) *Channel {
	s.GmtCreate = &v
	return s
}

func (s *Channel) SetGmtModified(v string) *Channel {
	s.GmtModified = &v
	return s
}

func (s *Channel) SetOutPutConfigList(v []*ChannelOutPutConfigList) *Channel {
	s.OutPutConfigList = v
	return s
}

func (s *Channel) SetPlaybackMode(v string) *Channel {
	s.PlaybackMode = &v
	return s
}

func (s *Channel) SetState(v int32) *Channel {
	s.State = &v
	return s
}

func (s *Channel) Validate() error {
	if s.OutPutConfigList != nil {
		for _, item := range s.OutPutConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChannelOutPutConfigList struct {
	// The name of the channel.
	//
	// example:
	//
	// MyChannel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The format.
	//
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The manifest name.
	//
	// example:
	//
	// index1
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The manifest settings.
	//
	// example:
	//
	// {"WindowDuration": 60,"AdMarkType":"Daterange"}
	ManifestSettings *string `json:"ManifestSettings,omitempty" xml:"ManifestSettings,omitempty"`
	// The playback URL.
	//
	// example:
	//
	// https://xxxxx.com/xxx.m3u8
	PlaybackUrl *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	// The name of the source group.
	//
	// example:
	//
	// group1
	SourceGroupName *string `json:"SourceGroupName,omitempty" xml:"SourceGroupName,omitempty"`
}

func (s ChannelOutPutConfigList) String() string {
	return dara.Prettify(s)
}

func (s ChannelOutPutConfigList) GoString() string {
	return s.String()
}

func (s *ChannelOutPutConfigList) GetChannelName() *string {
	return s.ChannelName
}

func (s *ChannelOutPutConfigList) GetFormat() *string {
	return s.Format
}

func (s *ChannelOutPutConfigList) GetManifestName() *string {
	return s.ManifestName
}

func (s *ChannelOutPutConfigList) GetManifestSettings() *string {
	return s.ManifestSettings
}

func (s *ChannelOutPutConfigList) GetPlaybackUrl() *string {
	return s.PlaybackUrl
}

func (s *ChannelOutPutConfigList) GetSourceGroupName() *string {
	return s.SourceGroupName
}

func (s *ChannelOutPutConfigList) SetChannelName(v string) *ChannelOutPutConfigList {
	s.ChannelName = &v
	return s
}

func (s *ChannelOutPutConfigList) SetFormat(v string) *ChannelOutPutConfigList {
	s.Format = &v
	return s
}

func (s *ChannelOutPutConfigList) SetManifestName(v string) *ChannelOutPutConfigList {
	s.ManifestName = &v
	return s
}

func (s *ChannelOutPutConfigList) SetManifestSettings(v string) *ChannelOutPutConfigList {
	s.ManifestSettings = &v
	return s
}

func (s *ChannelOutPutConfigList) SetPlaybackUrl(v string) *ChannelOutPutConfigList {
	s.PlaybackUrl = &v
	return s
}

func (s *ChannelOutPutConfigList) SetSourceGroupName(v string) *ChannelOutPutConfigList {
	s.SourceGroupName = &v
	return s
}

func (s *ChannelOutPutConfigList) Validate() error {
	return dara.Validate(s)
}
