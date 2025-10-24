// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudio(v string) *AddTemplateRequest
	GetAudio() *string
	SetContainer(v string) *AddTemplateRequest
	GetContainer() *string
	SetMuxConfig(v string) *AddTemplateRequest
	GetMuxConfig() *string
	SetName(v string) *AddTemplateRequest
	GetName() *string
	SetOwnerAccount(v string) *AddTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTemplateRequest
	GetResourceOwnerId() *int64
	SetTransConfig(v string) *AddTemplateRequest
	GetTransConfig() *string
	SetVideo(v string) *AddTemplateRequest
	GetVideo() *string
}

type AddTemplateRequest struct {
	// The audio stream settings. The value must be a JSON object. For more information, see [Audio](https://help.aliyun.com/document_detail/29253.html).
	//
	// > If you do not specify this parameter, output files do not contain audio streams. This parameter is required if you want to retain the audio streams.
	//
	// example:
	//
	// {"Codec":"H.264","Samplerate":"44100","Bitrate":"500","Channels":"2"}
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The container format. The value must be a JSON object that contains the Format parameter. If you do not specify this parameter, the transcoded media file is in MP4 format by default. This parameter is required if you want to use the transcoding template to generate media files in other formats. For more information, see [Container](https://help.aliyun.com/document_detail/29253.html).
	//
	// 	- Default value: MP4.
	//
	// 	- Video transcoding supports the following formats: FLV, MP4, HLS (M3U8 + TS), and MPEG-DASH (MPD + fMP4).
	//
	// > If the container format is FLV, the video codec cannot be set to H.265.
	//
	// 	- Audio transcoding supports the following formats: MP3, MP4, OGG, FLAC, and M4A.
	//
	// 	- Image transcoding supports the GIF and WebP formats.
	//
	// >
	//
	// 	- If the container format is GIF, the video codec must be set to GIF.
	//
	// 	- If the container format is WebP, the video codec must be set to WebP.
	//
	// example:
	//
	// {"Format":"mp4"}
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The segment settings. The value must be a JSON object. For more information, see [MuxConfig](https://help.aliyun.com/document_detail/29253.html). If you do not specify this parameter, media segment files are not generated. This parameter is required if you want to generate media segment files.
	//
	// example:
	//
	// {"Segment":{"Duration":"10"}}
	MuxConfig *string `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty"`
	// The name of the transcoding template. The name can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// mps-example
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The general transcoding settings. The value must be a JSON object. For more information, see [TransConfig](https://help.aliyun.com/document_detail/29253.html). If you do not specify this parameter, the default settings are used. This parameter is required if the default settings cannot meet your business requirements.
	//
	// example:
	//
	// {"TransMode":"onepass"}
	TransConfig *string `json:"TransConfig,omitempty" xml:"TransConfig,omitempty"`
	// The video stream settings. The value must be a JSON object. For more information, see [Video](https://help.aliyun.com/document_detail/29253.html).
	//
	// > If you do not specify this parameter, output files do not contain video streams. This parameter is required if you want to retain the video streams.
	//
	// example:
	//
	// {"Codec":"H.264","Profile":"high","Bitrate":"500","Crf":"15","Width":"256","Height":"800","Fps":"25","Gop":"10s"}
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s AddTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddTemplateRequest) GetAudio() *string {
	return s.Audio
}

func (s *AddTemplateRequest) GetContainer() *string {
	return s.Container
}

func (s *AddTemplateRequest) GetMuxConfig() *string {
	return s.MuxConfig
}

func (s *AddTemplateRequest) GetName() *string {
	return s.Name
}

func (s *AddTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTemplateRequest) GetTransConfig() *string {
	return s.TransConfig
}

func (s *AddTemplateRequest) GetVideo() *string {
	return s.Video
}

func (s *AddTemplateRequest) SetAudio(v string) *AddTemplateRequest {
	s.Audio = &v
	return s
}

func (s *AddTemplateRequest) SetContainer(v string) *AddTemplateRequest {
	s.Container = &v
	return s
}

func (s *AddTemplateRequest) SetMuxConfig(v string) *AddTemplateRequest {
	s.MuxConfig = &v
	return s
}

func (s *AddTemplateRequest) SetName(v string) *AddTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddTemplateRequest) SetOwnerAccount(v string) *AddTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddTemplateRequest) SetOwnerId(v int64) *AddTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTemplateRequest) SetResourceOwnerAccount(v string) *AddTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTemplateRequest) SetResourceOwnerId(v int64) *AddTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTemplateRequest) SetTransConfig(v string) *AddTemplateRequest {
	s.TransConfig = &v
	return s
}

func (s *AddTemplateRequest) SetVideo(v string) *AddTemplateRequest {
	s.Video = &v
	return s
}

func (s *AddTemplateRequest) Validate() error {
	return dara.Validate(s)
}
