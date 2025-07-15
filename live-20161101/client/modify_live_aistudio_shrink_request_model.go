// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveAIStudioShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundResourceId(v string) *ModifyLiveAIStudioShrinkRequest
	GetBackgroundResourceId() *string
	SetBackgroundResourceUrl(v string) *ModifyLiveAIStudioShrinkRequest
	GetBackgroundResourceUrl() *string
	SetBackgroundType(v string) *ModifyLiveAIStudioShrinkRequest
	GetBackgroundType() *string
	SetDescription(v string) *ModifyLiveAIStudioShrinkRequest
	GetDescription() *string
	SetHeight(v int32) *ModifyLiveAIStudioShrinkRequest
	GetHeight() *int32
	SetMattingLayoutShrink(v string) *ModifyLiveAIStudioShrinkRequest
	GetMattingLayoutShrink() *string
	SetMattingType(v string) *ModifyLiveAIStudioShrinkRequest
	GetMattingType() *string
	SetMediaLayoutShrink(v string) *ModifyLiveAIStudioShrinkRequest
	GetMediaLayoutShrink() *string
	SetMediaResourceId(v string) *ModifyLiveAIStudioShrinkRequest
	GetMediaResourceId() *string
	SetMediaResourceUrl(v string) *ModifyLiveAIStudioShrinkRequest
	GetMediaResourceUrl() *string
	SetMediaType(v string) *ModifyLiveAIStudioShrinkRequest
	GetMediaType() *string
	SetOwnerId(v int64) *ModifyLiveAIStudioShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLiveAIStudioShrinkRequest
	GetRegionId() *string
	SetStudioName(v string) *ModifyLiveAIStudioShrinkRequest
	GetStudioName() *string
	SetWidth(v int32) *ModifyLiveAIStudioShrinkRequest
	GetWidth() *int32
}

type ModifyLiveAIStudioShrinkRequest struct {
	// The ID of the background material in ApsaraVideo VOD. You can obtain the ID from the ApsaraVideo VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	BackgroundResourceId *string `json:"BackgroundResourceId,omitempty" xml:"BackgroundResourceId,omitempty"`
	// The URL of the background material.
	//
	// example:
	//
	// https://xxx.com/2.mp4
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
	// custom
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
	// d0eb493192c771efba644531858c01102
	MediaResourceId *string `json:"MediaResourceId,omitempty" xml:"MediaResourceId,omitempty"`
	// The URL of the multimedia material. Specify either this parameter or the MediaResourceId parameter.
	//
	// example:
	//
	// https://xxx.com/1.mp4
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
	// The name of the virtual studio template. The name is the same as the value of the StudioName parameter that was specified when you called the CreateLiveAIStudio operation to create the virtual studio template.
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

func (s ModifyLiveAIStudioShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveAIStudioShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveAIStudioShrinkRequest) GetBackgroundResourceId() *string {
	return s.BackgroundResourceId
}

func (s *ModifyLiveAIStudioShrinkRequest) GetBackgroundResourceUrl() *string {
	return s.BackgroundResourceUrl
}

func (s *ModifyLiveAIStudioShrinkRequest) GetBackgroundType() *string {
	return s.BackgroundType
}

func (s *ModifyLiveAIStudioShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyLiveAIStudioShrinkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *ModifyLiveAIStudioShrinkRequest) GetMattingLayoutShrink() *string {
	return s.MattingLayoutShrink
}

func (s *ModifyLiveAIStudioShrinkRequest) GetMattingType() *string {
	return s.MattingType
}

func (s *ModifyLiveAIStudioShrinkRequest) GetMediaLayoutShrink() *string {
	return s.MediaLayoutShrink
}

func (s *ModifyLiveAIStudioShrinkRequest) GetMediaResourceId() *string {
	return s.MediaResourceId
}

func (s *ModifyLiveAIStudioShrinkRequest) GetMediaResourceUrl() *string {
	return s.MediaResourceUrl
}

func (s *ModifyLiveAIStudioShrinkRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ModifyLiveAIStudioShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLiveAIStudioShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLiveAIStudioShrinkRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *ModifyLiveAIStudioShrinkRequest) GetWidth() *int32 {
	return s.Width
}

func (s *ModifyLiveAIStudioShrinkRequest) SetBackgroundResourceId(v string) *ModifyLiveAIStudioShrinkRequest {
	s.BackgroundResourceId = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetBackgroundResourceUrl(v string) *ModifyLiveAIStudioShrinkRequest {
	s.BackgroundResourceUrl = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetBackgroundType(v string) *ModifyLiveAIStudioShrinkRequest {
	s.BackgroundType = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetDescription(v string) *ModifyLiveAIStudioShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetHeight(v int32) *ModifyLiveAIStudioShrinkRequest {
	s.Height = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetMattingLayoutShrink(v string) *ModifyLiveAIStudioShrinkRequest {
	s.MattingLayoutShrink = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetMattingType(v string) *ModifyLiveAIStudioShrinkRequest {
	s.MattingType = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetMediaLayoutShrink(v string) *ModifyLiveAIStudioShrinkRequest {
	s.MediaLayoutShrink = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetMediaResourceId(v string) *ModifyLiveAIStudioShrinkRequest {
	s.MediaResourceId = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetMediaResourceUrl(v string) *ModifyLiveAIStudioShrinkRequest {
	s.MediaResourceUrl = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetMediaType(v string) *ModifyLiveAIStudioShrinkRequest {
	s.MediaType = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetOwnerId(v int64) *ModifyLiveAIStudioShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetRegionId(v string) *ModifyLiveAIStudioShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetStudioName(v string) *ModifyLiveAIStudioShrinkRequest {
	s.StudioName = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) SetWidth(v int32) *ModifyLiveAIStudioShrinkRequest {
	s.Width = &v
	return s
}

func (s *ModifyLiveAIStudioShrinkRequest) Validate() error {
	return dara.Validate(s)
}
