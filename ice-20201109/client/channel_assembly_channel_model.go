// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelAssemblyChannel interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPolicy(v bool) *ChannelAssemblyChannel
	GetAccessPolicy() *bool
	SetAccessToken(v string) *ChannelAssemblyChannel
	GetAccessToken() *string
	SetArn(v string) *ChannelAssemblyChannel
	GetArn() *string
	SetChannelName(v string) *ChannelAssemblyChannel
	GetChannelName() *string
	SetChannelTier(v string) *ChannelAssemblyChannel
	GetChannelTier() *string
	SetFillerSourceLocationName(v string) *ChannelAssemblyChannel
	GetFillerSourceLocationName() *string
	SetFillerSourceName(v string) *ChannelAssemblyChannel
	GetFillerSourceName() *string
	SetGmtCreate(v string) *ChannelAssemblyChannel
	GetGmtCreate() *string
	SetGmtModified(v string) *ChannelAssemblyChannel
	GetGmtModified() *string
	SetOutPutConfigList(v []*ChannelAssemblyChannelOutPutConfigList) *ChannelAssemblyChannel
	GetOutPutConfigList() []*ChannelAssemblyChannelOutPutConfigList
	SetPlaybackMode(v string) *ChannelAssemblyChannel
	GetPlaybackMode() *string
	SetState(v int32) *ChannelAssemblyChannel
	GetState() *int32
}

type ChannelAssemblyChannel struct {
	AccessPolicy             *bool                                     `json:"AccessPolicy,omitempty" xml:"AccessPolicy,omitempty"`
	AccessToken              *string                                   `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	Arn                      *string                                   `json:"Arn,omitempty" xml:"Arn,omitempty"`
	ChannelName              *string                                   `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	ChannelTier              *string                                   `json:"ChannelTier,omitempty" xml:"ChannelTier,omitempty"`
	FillerSourceLocationName *string                                   `json:"FillerSourceLocationName,omitempty" xml:"FillerSourceLocationName,omitempty"`
	FillerSourceName         *string                                   `json:"FillerSourceName,omitempty" xml:"FillerSourceName,omitempty"`
	GmtCreate                *string                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified              *string                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OutPutConfigList         []*ChannelAssemblyChannelOutPutConfigList `json:"OutPutConfigList,omitempty" xml:"OutPutConfigList,omitempty" type:"Repeated"`
	PlaybackMode             *string                                   `json:"PlaybackMode,omitempty" xml:"PlaybackMode,omitempty"`
	State                    *int32                                    `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ChannelAssemblyChannel) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblyChannel) GoString() string {
	return s.String()
}

func (s *ChannelAssemblyChannel) GetAccessPolicy() *bool {
	return s.AccessPolicy
}

func (s *ChannelAssemblyChannel) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ChannelAssemblyChannel) GetArn() *string {
	return s.Arn
}

func (s *ChannelAssemblyChannel) GetChannelName() *string {
	return s.ChannelName
}

func (s *ChannelAssemblyChannel) GetChannelTier() *string {
	return s.ChannelTier
}

func (s *ChannelAssemblyChannel) GetFillerSourceLocationName() *string {
	return s.FillerSourceLocationName
}

func (s *ChannelAssemblyChannel) GetFillerSourceName() *string {
	return s.FillerSourceName
}

func (s *ChannelAssemblyChannel) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ChannelAssemblyChannel) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ChannelAssemblyChannel) GetOutPutConfigList() []*ChannelAssemblyChannelOutPutConfigList {
	return s.OutPutConfigList
}

func (s *ChannelAssemblyChannel) GetPlaybackMode() *string {
	return s.PlaybackMode
}

func (s *ChannelAssemblyChannel) GetState() *int32 {
	return s.State
}

func (s *ChannelAssemblyChannel) SetAccessPolicy(v bool) *ChannelAssemblyChannel {
	s.AccessPolicy = &v
	return s
}

func (s *ChannelAssemblyChannel) SetAccessToken(v string) *ChannelAssemblyChannel {
	s.AccessToken = &v
	return s
}

func (s *ChannelAssemblyChannel) SetArn(v string) *ChannelAssemblyChannel {
	s.Arn = &v
	return s
}

func (s *ChannelAssemblyChannel) SetChannelName(v string) *ChannelAssemblyChannel {
	s.ChannelName = &v
	return s
}

func (s *ChannelAssemblyChannel) SetChannelTier(v string) *ChannelAssemblyChannel {
	s.ChannelTier = &v
	return s
}

func (s *ChannelAssemblyChannel) SetFillerSourceLocationName(v string) *ChannelAssemblyChannel {
	s.FillerSourceLocationName = &v
	return s
}

func (s *ChannelAssemblyChannel) SetFillerSourceName(v string) *ChannelAssemblyChannel {
	s.FillerSourceName = &v
	return s
}

func (s *ChannelAssemblyChannel) SetGmtCreate(v string) *ChannelAssemblyChannel {
	s.GmtCreate = &v
	return s
}

func (s *ChannelAssemblyChannel) SetGmtModified(v string) *ChannelAssemblyChannel {
	s.GmtModified = &v
	return s
}

func (s *ChannelAssemblyChannel) SetOutPutConfigList(v []*ChannelAssemblyChannelOutPutConfigList) *ChannelAssemblyChannel {
	s.OutPutConfigList = v
	return s
}

func (s *ChannelAssemblyChannel) SetPlaybackMode(v string) *ChannelAssemblyChannel {
	s.PlaybackMode = &v
	return s
}

func (s *ChannelAssemblyChannel) SetState(v int32) *ChannelAssemblyChannel {
	s.State = &v
	return s
}

func (s *ChannelAssemblyChannel) Validate() error {
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

type ChannelAssemblyChannelOutPutConfigList struct {
	ChannelName      *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	Format           *string `json:"Format,omitempty" xml:"Format,omitempty"`
	ManifestName     *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	ManifestSettings *string `json:"ManifestSettings,omitempty" xml:"ManifestSettings,omitempty"`
	PlaybackUrl      *string `json:"PlaybackUrl,omitempty" xml:"PlaybackUrl,omitempty"`
	SourceGroupName  *string `json:"SourceGroupName,omitempty" xml:"SourceGroupName,omitempty"`
}

func (s ChannelAssemblyChannelOutPutConfigList) String() string {
	return dara.Prettify(s)
}

func (s ChannelAssemblyChannelOutPutConfigList) GoString() string {
	return s.String()
}

func (s *ChannelAssemblyChannelOutPutConfigList) GetChannelName() *string {
	return s.ChannelName
}

func (s *ChannelAssemblyChannelOutPutConfigList) GetFormat() *string {
	return s.Format
}

func (s *ChannelAssemblyChannelOutPutConfigList) GetManifestName() *string {
	return s.ManifestName
}

func (s *ChannelAssemblyChannelOutPutConfigList) GetManifestSettings() *string {
	return s.ManifestSettings
}

func (s *ChannelAssemblyChannelOutPutConfigList) GetPlaybackUrl() *string {
	return s.PlaybackUrl
}

func (s *ChannelAssemblyChannelOutPutConfigList) GetSourceGroupName() *string {
	return s.SourceGroupName
}

func (s *ChannelAssemblyChannelOutPutConfigList) SetChannelName(v string) *ChannelAssemblyChannelOutPutConfigList {
	s.ChannelName = &v
	return s
}

func (s *ChannelAssemblyChannelOutPutConfigList) SetFormat(v string) *ChannelAssemblyChannelOutPutConfigList {
	s.Format = &v
	return s
}

func (s *ChannelAssemblyChannelOutPutConfigList) SetManifestName(v string) *ChannelAssemblyChannelOutPutConfigList {
	s.ManifestName = &v
	return s
}

func (s *ChannelAssemblyChannelOutPutConfigList) SetManifestSettings(v string) *ChannelAssemblyChannelOutPutConfigList {
	s.ManifestSettings = &v
	return s
}

func (s *ChannelAssemblyChannelOutPutConfigList) SetPlaybackUrl(v string) *ChannelAssemblyChannelOutPutConfigList {
	s.PlaybackUrl = &v
	return s
}

func (s *ChannelAssemblyChannelOutPutConfigList) SetSourceGroupName(v string) *ChannelAssemblyChannelOutPutConfigList {
	s.SourceGroupName = &v
	return s
}

func (s *ChannelAssemblyChannelOutPutConfigList) Validate() error {
	return dara.Validate(s)
}
