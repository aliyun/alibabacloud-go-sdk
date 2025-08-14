// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveChannelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioSettingsShrink(v string) *CreateMediaLiveChannelShrinkRequest
	GetAudioSettingsShrink() *string
	SetInputAttachmentsShrink(v string) *CreateMediaLiveChannelShrinkRequest
	GetInputAttachmentsShrink() *string
	SetName(v string) *CreateMediaLiveChannelShrinkRequest
	GetName() *string
	SetOutputGroupsShrink(v string) *CreateMediaLiveChannelShrinkRequest
	GetOutputGroupsShrink() *string
	SetVideoSettingsShrink(v string) *CreateMediaLiveChannelShrinkRequest
	GetVideoSettingsShrink() *string
}

type CreateMediaLiveChannelShrinkRequest struct {
	// The audio settings.
	AudioSettingsShrink *string `json:"AudioSettings,omitempty" xml:"AudioSettings,omitempty"`
	// The associated inputs.
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

func (s CreateMediaLiveChannelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelShrinkRequest) GetAudioSettingsShrink() *string {
	return s.AudioSettingsShrink
}

func (s *CreateMediaLiveChannelShrinkRequest) GetInputAttachmentsShrink() *string {
	return s.InputAttachmentsShrink
}

func (s *CreateMediaLiveChannelShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveChannelShrinkRequest) GetOutputGroupsShrink() *string {
	return s.OutputGroupsShrink
}

func (s *CreateMediaLiveChannelShrinkRequest) GetVideoSettingsShrink() *string {
	return s.VideoSettingsShrink
}

func (s *CreateMediaLiveChannelShrinkRequest) SetAudioSettingsShrink(v string) *CreateMediaLiveChannelShrinkRequest {
	s.AudioSettingsShrink = &v
	return s
}

func (s *CreateMediaLiveChannelShrinkRequest) SetInputAttachmentsShrink(v string) *CreateMediaLiveChannelShrinkRequest {
	s.InputAttachmentsShrink = &v
	return s
}

func (s *CreateMediaLiveChannelShrinkRequest) SetName(v string) *CreateMediaLiveChannelShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveChannelShrinkRequest) SetOutputGroupsShrink(v string) *CreateMediaLiveChannelShrinkRequest {
	s.OutputGroupsShrink = &v
	return s
}

func (s *CreateMediaLiveChannelShrinkRequest) SetVideoSettingsShrink(v string) *CreateMediaLiveChannelShrinkRequest {
	s.VideoSettingsShrink = &v
	return s
}

func (s *CreateMediaLiveChannelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
