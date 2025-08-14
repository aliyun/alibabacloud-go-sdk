// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveChannelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioSettingsShrink(v string) *UpdateMediaLiveChannelShrinkRequest
	GetAudioSettingsShrink() *string
	SetChannelId(v string) *UpdateMediaLiveChannelShrinkRequest
	GetChannelId() *string
	SetInputAttachmentsShrink(v string) *UpdateMediaLiveChannelShrinkRequest
	GetInputAttachmentsShrink() *string
	SetName(v string) *UpdateMediaLiveChannelShrinkRequest
	GetName() *string
	SetOutputGroupsShrink(v string) *UpdateMediaLiveChannelShrinkRequest
	GetOutputGroupsShrink() *string
	SetVideoSettingsShrink(v string) *UpdateMediaLiveChannelShrinkRequest
	GetVideoSettingsShrink() *string
}

type UpdateMediaLiveChannelShrinkRequest struct {
	// The audio settings.
	AudioSettingsShrink *string `json:"AudioSettings,omitempty" xml:"AudioSettings,omitempty"`
	// The ID of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The inputs associated with the channel.
	//
	// This parameter is required.
	InputAttachmentsShrink *string `json:"InputAttachments,omitempty" xml:"InputAttachments,omitempty"`
	// The name of the channel. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// mych
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output groups.
	//
	// This parameter is required.
	OutputGroupsShrink *string `json:"OutputGroups,omitempty" xml:"OutputGroups,omitempty"`
	// The video settings.
	VideoSettingsShrink *string `json:"VideoSettings,omitempty" xml:"VideoSettings,omitempty"`
}

func (s UpdateMediaLiveChannelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelShrinkRequest) GetAudioSettingsShrink() *string {
	return s.AudioSettingsShrink
}

func (s *UpdateMediaLiveChannelShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateMediaLiveChannelShrinkRequest) GetInputAttachmentsShrink() *string {
	return s.InputAttachmentsShrink
}

func (s *UpdateMediaLiveChannelShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveChannelShrinkRequest) GetOutputGroupsShrink() *string {
	return s.OutputGroupsShrink
}

func (s *UpdateMediaLiveChannelShrinkRequest) GetVideoSettingsShrink() *string {
	return s.VideoSettingsShrink
}

func (s *UpdateMediaLiveChannelShrinkRequest) SetAudioSettingsShrink(v string) *UpdateMediaLiveChannelShrinkRequest {
	s.AudioSettingsShrink = &v
	return s
}

func (s *UpdateMediaLiveChannelShrinkRequest) SetChannelId(v string) *UpdateMediaLiveChannelShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateMediaLiveChannelShrinkRequest) SetInputAttachmentsShrink(v string) *UpdateMediaLiveChannelShrinkRequest {
	s.InputAttachmentsShrink = &v
	return s
}

func (s *UpdateMediaLiveChannelShrinkRequest) SetName(v string) *UpdateMediaLiveChannelShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveChannelShrinkRequest) SetOutputGroupsShrink(v string) *UpdateMediaLiveChannelShrinkRequest {
	s.OutputGroupsShrink = &v
	return s
}

func (s *UpdateMediaLiveChannelShrinkRequest) SetVideoSettingsShrink(v string) *UpdateMediaLiveChannelShrinkRequest {
	s.VideoSettingsShrink = &v
	return s
}

func (s *UpdateMediaLiveChannelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
