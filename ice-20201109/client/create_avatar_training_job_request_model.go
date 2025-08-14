// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAvatarTrainingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarDescription(v string) *CreateAvatarTrainingJobRequest
	GetAvatarDescription() *string
	SetAvatarName(v string) *CreateAvatarTrainingJobRequest
	GetAvatarName() *string
	SetAvatarType(v string) *CreateAvatarTrainingJobRequest
	GetAvatarType() *string
	SetPortrait(v string) *CreateAvatarTrainingJobRequest
	GetPortrait() *string
	SetThumbnail(v string) *CreateAvatarTrainingJobRequest
	GetThumbnail() *string
	SetTransparent(v bool) *CreateAvatarTrainingJobRequest
	GetTransparent() *bool
	SetVideo(v string) *CreateAvatarTrainingJobRequest
	GetVideo() *string
}

type CreateAvatarTrainingJobRequest struct {
	// 	- The description of the digital human.
	//
	// 	- The description can be up to 1,000 characters in length.
	AvatarDescription *string `json:"AvatarDescription,omitempty" xml:"AvatarDescription,omitempty"`
	// 	- The name of the digital human.
	//
	// 	- The name can be up to seven characters in length.
	//
	// This parameter is required.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The type of the digital human.
	//
	// example:
	//
	// 2DAvatar
	AvatarType *string `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	// 	- The media asset ID of the portrait image.
	//
	// 	- The value must be 32 characters in length.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Portrait *string `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	// 	- The thumbnail URL.
	//
	// 	- After the digital human is trained, the thumbnail is uploaded to this URL.
	//
	// 	- The URL must be a valid public Object Storage Service (OSS) URL.
	//
	// 	- The URL can be up to 512 characters in length.
	//
	// example:
	//
	// https://your-bucket.oss-cn-hangzhou.aliyuncs.com/thumbnail.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// 	- Specifies whether the training video supports alpha channels.
	//
	//     **
	//
	//     **Note**: Make sure that the current settings are consistent with those of the submitted training video. Otherwise, the digital human may malfunction.
	//
	// example:
	//
	// True
	Transparent *bool `json:"Transparent,omitempty" xml:"Transparent,omitempty"`
	// 	- The ID of the video used for training.
	//
	// 	- The value must be 32 characters in length.
	//
	// 	- Supported formats: MP4, MOV, and WebM.
	//
	// 	- The duration of the video must be 5 to 15 minutes.
	//
	// 	- The resolution of the video must be 1920×1080 or 1080×1920.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CreateAvatarTrainingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAvatarTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *CreateAvatarTrainingJobRequest) GetAvatarDescription() *string {
	return s.AvatarDescription
}

func (s *CreateAvatarTrainingJobRequest) GetAvatarName() *string {
	return s.AvatarName
}

func (s *CreateAvatarTrainingJobRequest) GetAvatarType() *string {
	return s.AvatarType
}

func (s *CreateAvatarTrainingJobRequest) GetPortrait() *string {
	return s.Portrait
}

func (s *CreateAvatarTrainingJobRequest) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *CreateAvatarTrainingJobRequest) GetTransparent() *bool {
	return s.Transparent
}

func (s *CreateAvatarTrainingJobRequest) GetVideo() *string {
	return s.Video
}

func (s *CreateAvatarTrainingJobRequest) SetAvatarDescription(v string) *CreateAvatarTrainingJobRequest {
	s.AvatarDescription = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) SetAvatarName(v string) *CreateAvatarTrainingJobRequest {
	s.AvatarName = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) SetAvatarType(v string) *CreateAvatarTrainingJobRequest {
	s.AvatarType = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) SetPortrait(v string) *CreateAvatarTrainingJobRequest {
	s.Portrait = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) SetThumbnail(v string) *CreateAvatarTrainingJobRequest {
	s.Thumbnail = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) SetTransparent(v bool) *CreateAvatarTrainingJobRequest {
	s.Transparent = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) SetVideo(v string) *CreateAvatarTrainingJobRequest {
	s.Video = &v
	return s
}

func (s *CreateAvatarTrainingJobRequest) Validate() error {
	return dara.Validate(s)
}
