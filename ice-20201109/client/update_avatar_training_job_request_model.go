// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAvatarTrainingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarDescription(v string) *UpdateAvatarTrainingJobRequest
	GetAvatarDescription() *string
	SetAvatarName(v string) *UpdateAvatarTrainingJobRequest
	GetAvatarName() *string
	SetJobId(v string) *UpdateAvatarTrainingJobRequest
	GetJobId() *string
	SetPortrait(v string) *UpdateAvatarTrainingJobRequest
	GetPortrait() *string
	SetThumbnail(v string) *UpdateAvatarTrainingJobRequest
	GetThumbnail() *string
	SetTransparent(v bool) *UpdateAvatarTrainingJobRequest
	GetTransparent() *bool
	SetVideo(v string) *UpdateAvatarTrainingJobRequest
	GetVideo() *string
}

type UpdateAvatarTrainingJobRequest struct {
	// 	- The description of the digital human.
	//
	// 	- The description can be up to 1,000 characters in length.
	AvatarDescription *string `json:"AvatarDescription,omitempty" xml:"AvatarDescription,omitempty"`
	// 	- The name of the digital human.
	//
	// 	- The name can be up to seven characters in length.
	AvatarName *string `json:"AvatarName,omitempty" xml:"AvatarName,omitempty"`
	// The ID of the digital human training job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
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
	// 	- The URL cannot be updated after the digital human is trained.
	//
	// example:
	//
	// https://your-bucket.oss-cn-hangzhou.aliyuncs.com/thumbnail.png
	Thumbnail *string `json:"Thumbnail,omitempty" xml:"Thumbnail,omitempty"`
	// 	- Indicates whether the input video supports alpha channels.
	//
	// 	- You can modify this parameter only if the job is in the Init or Fail state.
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

func (s UpdateAvatarTrainingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAvatarTrainingJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateAvatarTrainingJobRequest) GetAvatarDescription() *string {
	return s.AvatarDescription
}

func (s *UpdateAvatarTrainingJobRequest) GetAvatarName() *string {
	return s.AvatarName
}

func (s *UpdateAvatarTrainingJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateAvatarTrainingJobRequest) GetPortrait() *string {
	return s.Portrait
}

func (s *UpdateAvatarTrainingJobRequest) GetThumbnail() *string {
	return s.Thumbnail
}

func (s *UpdateAvatarTrainingJobRequest) GetTransparent() *bool {
	return s.Transparent
}

func (s *UpdateAvatarTrainingJobRequest) GetVideo() *string {
	return s.Video
}

func (s *UpdateAvatarTrainingJobRequest) SetAvatarDescription(v string) *UpdateAvatarTrainingJobRequest {
	s.AvatarDescription = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) SetAvatarName(v string) *UpdateAvatarTrainingJobRequest {
	s.AvatarName = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) SetJobId(v string) *UpdateAvatarTrainingJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) SetPortrait(v string) *UpdateAvatarTrainingJobRequest {
	s.Portrait = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) SetThumbnail(v string) *UpdateAvatarTrainingJobRequest {
	s.Thumbnail = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) SetTransparent(v bool) *UpdateAvatarTrainingJobRequest {
	s.Transparent = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) SetVideo(v string) *UpdateAvatarTrainingJobRequest {
	s.Video = &v
	return s
}

func (s *UpdateAvatarTrainingJobRequest) Validate() error {
	return dara.Validate(s)
}
