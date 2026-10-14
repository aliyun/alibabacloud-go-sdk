// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMedia(v []*CreateVideoTaskRequestMedia) *CreateVideoTaskRequest
	GetMedia() []*CreateVideoTaskRequestMedia
	SetAudio(v bool) *CreateVideoTaskRequest
	GetAudio() *bool
	SetClientToken(v string) *CreateVideoTaskRequest
	GetClientToken() *string
	SetDuration(v int32) *CreateVideoTaskRequest
	GetDuration() *int32
	SetModel(v string) *CreateVideoTaskRequest
	GetModel() *string
	SetPrompt(v string) *CreateVideoTaskRequest
	GetPrompt() *string
	SetPromptExtend(v bool) *CreateVideoTaskRequest
	GetPromptExtend() *bool
	SetRatio(v string) *CreateVideoTaskRequest
	GetRatio() *string
	SetResolution(v string) *CreateVideoTaskRequest
	GetResolution() *string
	SetSeed(v int32) *CreateVideoTaskRequest
	GetSeed() *int32
	SetWatermark(v bool) *CreateVideoTaskRequest
	GetWatermark() *bool
}

type CreateVideoTaskRequest struct {
	Media []*CreateVideoTaskRequestMedia `json:"Media,omitempty" xml:"Media,omitempty" type:"Repeated"`
	// Specifies whether to include an audio track. Default value: true.
	//
	// example:
	//
	// true
	Audio *bool `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The idempotency parameter.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The video duration in seconds. Valid values: an integer from 2 to 30, -1 (intelligent duration), and 0 (server default of 5 seconds).
	//
	// example:
	//
	// 2
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The model name. Valid values: wan3.0-video (default) and wan3.0-video-prime (high-speed version).
	//
	// example:
	//
	// wan3.0-video
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The video description prompt. You must specify either this parameter or Media. For text-to-video generation, only Prompt is required.
	//
	// example:
	//
	// A meeting room full of cats
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Specifies whether to enable intelligent rewriting (prompt_extend). Default value: true.
	//
	// example:
	//
	// true
	PromptExtend *bool `json:"PromptExtend,omitempty" xml:"PromptExtend,omitempty"`
	// The aspect ratio. Valid values: adaptive (default), 16:9, 4:3, 1:1, 3:4, and 9:16.
	//
	// example:
	//
	// adaptive
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The resolution. Valid values: 1080P (default), 720P, and 480P.
	//
	// example:
	//
	// 1080P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The random seed. Valid values: -1 (random), an integer from 0 to 2147483647, and 0 (server default of random).
	//
	// example:
	//
	// 42
	Seed *int32 `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// Specifies whether to add a Model Studio watermark. Default value: false.
	//
	// example:
	//
	// false
	Watermark *bool `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s CreateVideoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoTaskRequest) GetMedia() []*CreateVideoTaskRequestMedia {
	return s.Media
}

func (s *CreateVideoTaskRequest) GetAudio() *bool {
	return s.Audio
}

func (s *CreateVideoTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVideoTaskRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateVideoTaskRequest) GetModel() *string {
	return s.Model
}

func (s *CreateVideoTaskRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateVideoTaskRequest) GetPromptExtend() *bool {
	return s.PromptExtend
}

func (s *CreateVideoTaskRequest) GetRatio() *string {
	return s.Ratio
}

func (s *CreateVideoTaskRequest) GetResolution() *string {
	return s.Resolution
}

func (s *CreateVideoTaskRequest) GetSeed() *int32 {
	return s.Seed
}

func (s *CreateVideoTaskRequest) GetWatermark() *bool {
	return s.Watermark
}

func (s *CreateVideoTaskRequest) SetMedia(v []*CreateVideoTaskRequestMedia) *CreateVideoTaskRequest {
	s.Media = v
	return s
}

func (s *CreateVideoTaskRequest) SetAudio(v bool) *CreateVideoTaskRequest {
	s.Audio = &v
	return s
}

func (s *CreateVideoTaskRequest) SetClientToken(v string) *CreateVideoTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVideoTaskRequest) SetDuration(v int32) *CreateVideoTaskRequest {
	s.Duration = &v
	return s
}

func (s *CreateVideoTaskRequest) SetModel(v string) *CreateVideoTaskRequest {
	s.Model = &v
	return s
}

func (s *CreateVideoTaskRequest) SetPrompt(v string) *CreateVideoTaskRequest {
	s.Prompt = &v
	return s
}

func (s *CreateVideoTaskRequest) SetPromptExtend(v bool) *CreateVideoTaskRequest {
	s.PromptExtend = &v
	return s
}

func (s *CreateVideoTaskRequest) SetRatio(v string) *CreateVideoTaskRequest {
	s.Ratio = &v
	return s
}

func (s *CreateVideoTaskRequest) SetResolution(v string) *CreateVideoTaskRequest {
	s.Resolution = &v
	return s
}

func (s *CreateVideoTaskRequest) SetSeed(v int32) *CreateVideoTaskRequest {
	s.Seed = &v
	return s
}

func (s *CreateVideoTaskRequest) SetWatermark(v bool) *CreateVideoTaskRequest {
	s.Watermark = &v
	return s
}

func (s *CreateVideoTaskRequest) Validate() error {
	if s.Media != nil {
		for _, item := range s.Media {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVideoTaskRequestMedia struct {
	// The media type for the first item. Valid values: first_frame, last_frame, reference_image, reference_video, reference_audio, file, and link.
	//
	// example:
	//
	// JPEG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The media URL for the first item. You can specify a public HTTP or HTTPS URL, a temporary oss:// URL, or inline data in the format of data:{mime};base64,....
	//
	// example:
	//
	// Image URL
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateVideoTaskRequestMedia) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoTaskRequestMedia) GoString() string {
	return s.String()
}

func (s *CreateVideoTaskRequestMedia) GetType() *string {
	return s.Type
}

func (s *CreateVideoTaskRequestMedia) GetUrl() *string {
	return s.Url
}

func (s *CreateVideoTaskRequestMedia) SetType(v string) *CreateVideoTaskRequestMedia {
	s.Type = &v
	return s
}

func (s *CreateVideoTaskRequestMedia) SetUrl(v string) *CreateVideoTaskRequestMedia {
	s.Url = &v
	return s
}

func (s *CreateVideoTaskRequestMedia) Validate() error {
	return dara.Validate(s)
}
