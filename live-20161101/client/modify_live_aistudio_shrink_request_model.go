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
	// VOD resource ID of the background material, obtained from the VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	BackgroundResourceId *string `json:"BackgroundResourceId,omitempty" xml:"BackgroundResourceId,omitempty"`
	// Resource access URL of the background material.
	//
	// example:
	//
	// https://xxx.com/2.mp4
	BackgroundResourceUrl *string `json:"BackgroundResourceUrl,omitempty" xml:"BackgroundResourceUrl,omitempty"`
	// Background material type:
	//
	// - VOD: Video on demand
	//
	// - PIC: Image
	//
	// - LIVE: Live stream
	//
	// example:
	//
	// VOD
	BackgroundType *string `json:"BackgroundType,omitempty" xml:"BackgroundType,omitempty"`
	// Custom description.
	//
	// example:
	//
	// custom
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Preview screen height, unit: px.
	//
	// The preview screen width x height only supports the following specifications:
	//
	// - Landscape Smooth 360P 640x360
	//
	// - Portrait Smooth 360P 360x640
	//
	// - Landscape Standard Definition 480P 854x480
	//
	// - Portrait Standard Definition 480P 480x854
	//
	// - Landscape HD 720P 1280x720
	//
	// - Portrait HD 720P 720x1280
	//
	// - Landscape Full HD 1080P 1920x1080
	//
	// - Portrait Full HD 1080P 1080x1920
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Layout position information of the source stream after matting.
	//
	// This parameter is required.
	MattingLayoutShrink *string `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty"`
	// Matting type:
	//
	// - green: Green screen matting
	//
	// - blue: Blue screen matting
	//
	// - complex: Real-scene matting
	//
	// This parameter is required.
	//
	// example:
	//
	// complex
	MattingType *string `json:"MattingType,omitempty" xml:"MattingType,omitempty"`
	// Layout position information of the multimedia material.
	MediaLayoutShrink *string `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty"`
	// VOD resource ID of the multimedia material, obtained from the VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c01102
	MediaResourceId *string `json:"MediaResourceId,omitempty" xml:"MediaResourceId,omitempty"`
	// Resource access URL of the multimedia material. Either this or the resource ID should be provided.
	//
	// example:
	//
	// https://xxx.com/1.mp4
	MediaResourceUrl *string `json:"MediaResourceUrl,omitempty" xml:"MediaResourceUrl,omitempty"`
	// Multimedia material type:
	//
	// - VOD: Video on demand
	//
	// - PIC: Image
	//
	// - LIVE: Live stream
	//
	// example:
	//
	// VOD
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Virtual studio template name, same as the StudioName parameter in the create API.
	//
	// This parameter is required.
	//
	// example:
	//
	// stu02
	StudioName *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
	// Preview screen width, unit: px.
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
