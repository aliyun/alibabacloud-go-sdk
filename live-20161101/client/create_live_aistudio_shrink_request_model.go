// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveAIStudioShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundResourceId(v string) *CreateLiveAIStudioShrinkRequest
	GetBackgroundResourceId() *string
	SetBackgroundResourceUrl(v string) *CreateLiveAIStudioShrinkRequest
	GetBackgroundResourceUrl() *string
	SetBackgroundType(v string) *CreateLiveAIStudioShrinkRequest
	GetBackgroundType() *string
	SetDescription(v string) *CreateLiveAIStudioShrinkRequest
	GetDescription() *string
	SetHeight(v int32) *CreateLiveAIStudioShrinkRequest
	GetHeight() *int32
	SetMattingLayoutShrink(v string) *CreateLiveAIStudioShrinkRequest
	GetMattingLayoutShrink() *string
	SetMattingType(v string) *CreateLiveAIStudioShrinkRequest
	GetMattingType() *string
	SetMediaLayoutShrink(v string) *CreateLiveAIStudioShrinkRequest
	GetMediaLayoutShrink() *string
	SetMediaResourceId(v string) *CreateLiveAIStudioShrinkRequest
	GetMediaResourceId() *string
	SetMediaResourceUrl(v string) *CreateLiveAIStudioShrinkRequest
	GetMediaResourceUrl() *string
	SetMediaType(v string) *CreateLiveAIStudioShrinkRequest
	GetMediaType() *string
	SetOwnerId(v int64) *CreateLiveAIStudioShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLiveAIStudioShrinkRequest
	GetRegionId() *string
	SetStudioName(v string) *CreateLiveAIStudioShrinkRequest
	GetStudioName() *string
	SetWidth(v int32) *CreateLiveAIStudioShrinkRequest
	GetWidth() *int32
}

type CreateLiveAIStudioShrinkRequest struct {
	// The ID of the background material in ApsaraVideo VOD. You can obtain the ID from the ApsaraVideo VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	BackgroundResourceId *string `json:"BackgroundResourceId,omitempty" xml:"BackgroundResourceId,omitempty"`
	// The URL of the background material. Specify either this parameter or the BackgroundResourceId parameter.
	//
	// example:
	//
	// https://xxx.com/1.mp4
	BackgroundResourceUrl *string `json:"BackgroundResourceUrl,omitempty" xml:"BackgroundResourceUrl,omitempty"`
	// The type of the background material. Valid values:
	//
	// 	- VOD: a video in ApsaraVideo VOD
	//
	// 	- PIC: an image
	//
	// 	- LIVE: a live stream
	//
	// example:
	//
	// VOD
	BackgroundType *string `json:"BackgroundType,omitempty" xml:"BackgroundType,omitempty"`
	// The custom description.
	//
	// example:
	//
	// template 1080
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The preview height. Unit: pixels.
	//
	// The following preview specifications (width × height) are supported:
	//
	// 	- Landscape low definition 360p (640×360)
	//
	// 	- Portrait low definition 360p (360×640)
	//
	// 	- Landscape standard definition 480p (854×480)
	//
	// 	- Portrait standard definition 480p (480×854)
	//
	// 	- Landscape high definition 720p (1280×720)
	//
	// 	- Portrait high definition 720p (720×1280)
	//
	// 	- Landscape ultra-high definition 1080p (1920×1080)
	//
	// 	- Portrait ultra-high definition 1080p (1080×1920)
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The layout information of the chroma-keyed material.
	//
	// This parameter is required.
	MattingLayoutShrink *string `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty"`
	// The type of chroma key. Valid values:
	//
	// 	- green: green-screen chroma key
	//
	// 	- blue: blue-screen chroma key
	//
	// 	- complex: background replacement
	//
	// This parameter is required.
	//
	// example:
	//
	// complex
	MattingType *string `json:"MattingType,omitempty" xml:"MattingType,omitempty"`
	// The layout information of the multimedia material.
	MediaLayoutShrink *string `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty"`
	// The ID of the multimedia material in ApsaraVideo VOD. You can obtain the ID from the ApsaraVideo VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	MediaResourceId *string `json:"MediaResourceId,omitempty" xml:"MediaResourceId,omitempty"`
	// The URL of the multimedia material. Specify either this parameter or the MediaResourceId parameter.
	//
	// example:
	//
	// https://xxx.com/2.mp4
	MediaResourceUrl *string `json:"MediaResourceUrl,omitempty" xml:"MediaResourceUrl,omitempty"`
	// The type of the multimedia material. Valid values:
	//
	// 	- VOD: a video in ApsaraVideo VOD
	//
	// 	- PIC: an image
	//
	// 	- LIVE: a live stream
	//
	// example:
	//
	// VOD
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the virtual studio template. The name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// stu02
	StudioName *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
	// The preview width. Unit: pixels.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateLiveAIStudioShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveAIStudioShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveAIStudioShrinkRequest) GetBackgroundResourceId() *string {
	return s.BackgroundResourceId
}

func (s *CreateLiveAIStudioShrinkRequest) GetBackgroundResourceUrl() *string {
	return s.BackgroundResourceUrl
}

func (s *CreateLiveAIStudioShrinkRequest) GetBackgroundType() *string {
	return s.BackgroundType
}

func (s *CreateLiveAIStudioShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLiveAIStudioShrinkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *CreateLiveAIStudioShrinkRequest) GetMattingLayoutShrink() *string {
	return s.MattingLayoutShrink
}

func (s *CreateLiveAIStudioShrinkRequest) GetMattingType() *string {
	return s.MattingType
}

func (s *CreateLiveAIStudioShrinkRequest) GetMediaLayoutShrink() *string {
	return s.MediaLayoutShrink
}

func (s *CreateLiveAIStudioShrinkRequest) GetMediaResourceId() *string {
	return s.MediaResourceId
}

func (s *CreateLiveAIStudioShrinkRequest) GetMediaResourceUrl() *string {
	return s.MediaResourceUrl
}

func (s *CreateLiveAIStudioShrinkRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *CreateLiveAIStudioShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLiveAIStudioShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLiveAIStudioShrinkRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *CreateLiveAIStudioShrinkRequest) GetWidth() *int32 {
	return s.Width
}

func (s *CreateLiveAIStudioShrinkRequest) SetBackgroundResourceId(v string) *CreateLiveAIStudioShrinkRequest {
	s.BackgroundResourceId = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetBackgroundResourceUrl(v string) *CreateLiveAIStudioShrinkRequest {
	s.BackgroundResourceUrl = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetBackgroundType(v string) *CreateLiveAIStudioShrinkRequest {
	s.BackgroundType = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetDescription(v string) *CreateLiveAIStudioShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetHeight(v int32) *CreateLiveAIStudioShrinkRequest {
	s.Height = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetMattingLayoutShrink(v string) *CreateLiveAIStudioShrinkRequest {
	s.MattingLayoutShrink = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetMattingType(v string) *CreateLiveAIStudioShrinkRequest {
	s.MattingType = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetMediaLayoutShrink(v string) *CreateLiveAIStudioShrinkRequest {
	s.MediaLayoutShrink = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetMediaResourceId(v string) *CreateLiveAIStudioShrinkRequest {
	s.MediaResourceId = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetMediaResourceUrl(v string) *CreateLiveAIStudioShrinkRequest {
	s.MediaResourceUrl = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetMediaType(v string) *CreateLiveAIStudioShrinkRequest {
	s.MediaType = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetOwnerId(v int64) *CreateLiveAIStudioShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetRegionId(v string) *CreateLiveAIStudioShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetStudioName(v string) *CreateLiveAIStudioShrinkRequest {
	s.StudioName = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) SetWidth(v int32) *CreateLiveAIStudioShrinkRequest {
	s.Width = &v
	return s
}

func (s *CreateLiveAIStudioShrinkRequest) Validate() error {
	return dara.Validate(s)
}
