// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudio(v string) *UpdateTemplateRequest
	GetAudio() *string
	SetContainer(v string) *UpdateTemplateRequest
	GetContainer() *string
	SetMuxConfig(v string) *UpdateTemplateRequest
	GetMuxConfig() *string
	SetName(v string) *UpdateTemplateRequest
	GetName() *string
	SetOwnerAccount(v string) *UpdateTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *UpdateTemplateRequest
	GetTemplateId() *string
	SetTransConfig(v string) *UpdateTemplateRequest
	GetTransConfig() *string
	SetVideo(v string) *UpdateTemplateRequest
	GetVideo() *string
}

type UpdateTemplateRequest struct {
	// The transmuxing configurations. The value is a JSON object. For more information, see the [MuxConfig](https://help.aliyun.com/document_detail/29253.html) parameter of the "Parameter details" topic.
	//
	// example:
	//
	// {"Codec":"aac","Samplerate":"44100","Bitrate":"500","Channels":"2"}
	Audio *string `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The configurations of the video stream. The value is a JSON object. For more information, see the [Video](https://help.aliyun.com/document_detail/29253.html) parameter of the "Parameter details" topic.
	//
	// example:
	//
	// {"Format":"mp4"}
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The general transcoding configurations. The value is a JSON object. For more information, see the [TransConfig](https://help.aliyun.com/document_detail/29253.html) parameter of the "Parameter details" topic.
	//
	// example:
	//
	// {"Segment":{"Duration":"10"}}
	MuxConfig *string `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty"`
	// The container format. The value is a JSON object. Default format: **MP4**.
	//
	// 	- Video formats: FLV, MP4, HLS (M3U8 + TS), and MPEG-DASH (MPD + fMP4)
	//
	// 	- Audio formats: MP3, MP4, OGG, FLAC, and M4A
	//
	// 	- Images formats: GIF and WebP
	//
	// For more information, see the [Container](https://help.aliyun.com/document_detail/29253.html) parameter of the "Parameter details" topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// MPS-example
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the template. The name can be up to 128 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16f01ad6175e4230ac42bb5182cd****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// {"TransMode":"onepass"}
	TransConfig *string `json:"TransConfig,omitempty" xml:"TransConfig,omitempty"`
	// The configurations of the audio stream. The value is a JSON object. For more information, see the [Audio](https://help.aliyun.com/document_detail/29253.html) parameter of the "Parameter details" topic.
	//
	// example:
	//
	// {"Codec":"H.264","Profile":"high","Bitrate":"500","Crf":"15","Width":"256","Height":"800","Fps":"25","Gop":"10"}
	Video *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetAudio() *string {
	return s.Audio
}

func (s *UpdateTemplateRequest) GetContainer() *string {
	return s.Container
}

func (s *UpdateTemplateRequest) GetMuxConfig() *string {
	return s.MuxConfig
}

func (s *UpdateTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateTemplateRequest) GetTransConfig() *string {
	return s.TransConfig
}

func (s *UpdateTemplateRequest) GetVideo() *string {
	return s.Video
}

func (s *UpdateTemplateRequest) SetAudio(v string) *UpdateTemplateRequest {
	s.Audio = &v
	return s
}

func (s *UpdateTemplateRequest) SetContainer(v string) *UpdateTemplateRequest {
	s.Container = &v
	return s
}

func (s *UpdateTemplateRequest) SetMuxConfig(v string) *UpdateTemplateRequest {
	s.MuxConfig = &v
	return s
}

func (s *UpdateTemplateRequest) SetName(v string) *UpdateTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTemplateRequest) SetOwnerAccount(v string) *UpdateTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTemplateRequest) SetOwnerId(v int64) *UpdateTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTemplateRequest) SetResourceOwnerAccount(v string) *UpdateTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTemplateRequest) SetResourceOwnerId(v int64) *UpdateTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTransConfig(v string) *UpdateTemplateRequest {
	s.TransConfig = &v
	return s
}

func (s *UpdateTemplateRequest) SetVideo(v string) *UpdateTemplateRequest {
	s.Video = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
