// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveAIStudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundResourceId(v string) *ModifyLiveAIStudioRequest
	GetBackgroundResourceId() *string
	SetBackgroundResourceUrl(v string) *ModifyLiveAIStudioRequest
	GetBackgroundResourceUrl() *string
	SetBackgroundType(v string) *ModifyLiveAIStudioRequest
	GetBackgroundType() *string
	SetDescription(v string) *ModifyLiveAIStudioRequest
	GetDescription() *string
	SetHeight(v int32) *ModifyLiveAIStudioRequest
	GetHeight() *int32
	SetMattingLayout(v *ModifyLiveAIStudioRequestMattingLayout) *ModifyLiveAIStudioRequest
	GetMattingLayout() *ModifyLiveAIStudioRequestMattingLayout
	SetMattingType(v string) *ModifyLiveAIStudioRequest
	GetMattingType() *string
	SetMediaLayout(v *ModifyLiveAIStudioRequestMediaLayout) *ModifyLiveAIStudioRequest
	GetMediaLayout() *ModifyLiveAIStudioRequestMediaLayout
	SetMediaResourceId(v string) *ModifyLiveAIStudioRequest
	GetMediaResourceId() *string
	SetMediaResourceUrl(v string) *ModifyLiveAIStudioRequest
	GetMediaResourceUrl() *string
	SetMediaType(v string) *ModifyLiveAIStudioRequest
	GetMediaType() *string
	SetOwnerId(v int64) *ModifyLiveAIStudioRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLiveAIStudioRequest
	GetRegionId() *string
	SetStudioName(v string) *ModifyLiveAIStudioRequest
	GetStudioName() *string
	SetWidth(v int32) *ModifyLiveAIStudioRequest
	GetWidth() *int32
}

type ModifyLiveAIStudioRequest struct {
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
	MattingLayout *ModifyLiveAIStudioRequestMattingLayout `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty" type:"Struct"`
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
	MediaLayout *ModifyLiveAIStudioRequestMediaLayout `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty" type:"Struct"`
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

func (s ModifyLiveAIStudioRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveAIStudioRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveAIStudioRequest) GetBackgroundResourceId() *string {
	return s.BackgroundResourceId
}

func (s *ModifyLiveAIStudioRequest) GetBackgroundResourceUrl() *string {
	return s.BackgroundResourceUrl
}

func (s *ModifyLiveAIStudioRequest) GetBackgroundType() *string {
	return s.BackgroundType
}

func (s *ModifyLiveAIStudioRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyLiveAIStudioRequest) GetHeight() *int32 {
	return s.Height
}

func (s *ModifyLiveAIStudioRequest) GetMattingLayout() *ModifyLiveAIStudioRequestMattingLayout {
	return s.MattingLayout
}

func (s *ModifyLiveAIStudioRequest) GetMattingType() *string {
	return s.MattingType
}

func (s *ModifyLiveAIStudioRequest) GetMediaLayout() *ModifyLiveAIStudioRequestMediaLayout {
	return s.MediaLayout
}

func (s *ModifyLiveAIStudioRequest) GetMediaResourceId() *string {
	return s.MediaResourceId
}

func (s *ModifyLiveAIStudioRequest) GetMediaResourceUrl() *string {
	return s.MediaResourceUrl
}

func (s *ModifyLiveAIStudioRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ModifyLiveAIStudioRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLiveAIStudioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLiveAIStudioRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *ModifyLiveAIStudioRequest) GetWidth() *int32 {
	return s.Width
}

func (s *ModifyLiveAIStudioRequest) SetBackgroundResourceId(v string) *ModifyLiveAIStudioRequest {
	s.BackgroundResourceId = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetBackgroundResourceUrl(v string) *ModifyLiveAIStudioRequest {
	s.BackgroundResourceUrl = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetBackgroundType(v string) *ModifyLiveAIStudioRequest {
	s.BackgroundType = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetDescription(v string) *ModifyLiveAIStudioRequest {
	s.Description = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetHeight(v int32) *ModifyLiveAIStudioRequest {
	s.Height = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetMattingLayout(v *ModifyLiveAIStudioRequestMattingLayout) *ModifyLiveAIStudioRequest {
	s.MattingLayout = v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetMattingType(v string) *ModifyLiveAIStudioRequest {
	s.MattingType = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetMediaLayout(v *ModifyLiveAIStudioRequestMediaLayout) *ModifyLiveAIStudioRequest {
	s.MediaLayout = v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetMediaResourceId(v string) *ModifyLiveAIStudioRequest {
	s.MediaResourceId = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetMediaResourceUrl(v string) *ModifyLiveAIStudioRequest {
	s.MediaResourceUrl = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetMediaType(v string) *ModifyLiveAIStudioRequest {
	s.MediaType = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetOwnerId(v int64) *ModifyLiveAIStudioRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetRegionId(v string) *ModifyLiveAIStudioRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetStudioName(v string) *ModifyLiveAIStudioRequest {
	s.StudioName = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) SetWidth(v int32) *ModifyLiveAIStudioRequest {
	s.Width = &v
	return s
}

func (s *ModifyLiveAIStudioRequest) Validate() error {
	if s.MattingLayout != nil {
		if err := s.MattingLayout.Validate(); err != nil {
			return err
		}
	}
	if s.MediaLayout != nil {
		if err := s.MediaLayout.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyLiveAIStudioRequestMattingLayout struct {
	// Normalized height value, which is the height ratio of the matted portrait to the background. Value range: **0~1**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// Position parameter, X coordinate. Value range: **0~1**. The material position uses the upper-left corner as the reference point.
	//
	// example:
	//
	// 0
	PositionX *float32 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// Position parameter, Y coordinate. Value range: **0~1**. The material position uses the upper-left corner as the reference point.
	//
	// example:
	//
	// 0
	PositionY *float32 `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
}

func (s ModifyLiveAIStudioRequestMattingLayout) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveAIStudioRequestMattingLayout) GoString() string {
	return s.String()
}

func (s *ModifyLiveAIStudioRequestMattingLayout) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *ModifyLiveAIStudioRequestMattingLayout) GetPositionX() *float32 {
	return s.PositionX
}

func (s *ModifyLiveAIStudioRequestMattingLayout) GetPositionY() *float32 {
	return s.PositionY
}

func (s *ModifyLiveAIStudioRequestMattingLayout) SetHeightNormalized(v float32) *ModifyLiveAIStudioRequestMattingLayout {
	s.HeightNormalized = &v
	return s
}

func (s *ModifyLiveAIStudioRequestMattingLayout) SetPositionX(v float32) *ModifyLiveAIStudioRequestMattingLayout {
	s.PositionX = &v
	return s
}

func (s *ModifyLiveAIStudioRequestMattingLayout) SetPositionY(v float32) *ModifyLiveAIStudioRequestMattingLayout {
	s.PositionY = &v
	return s
}

func (s *ModifyLiveAIStudioRequestMattingLayout) Validate() error {
	return dara.Validate(s)
}

type ModifyLiveAIStudioRequestMediaLayout struct {
	// Normalized height value of the material, which is the height ratio of the material to the background. Value range: **0~1**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// Position parameter, X coordinate. Value range: **0~1**. The material position uses the upper-left corner as the reference point.
	//
	// example:
	//
	// 0
	PositionX *float32 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// Position parameter, Y coordinate. Value range: **0~1**. The material position uses the upper-left corner as the reference point.
	//
	// example:
	//
	// 0
	PositionY *float32 `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
}

func (s ModifyLiveAIStudioRequestMediaLayout) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveAIStudioRequestMediaLayout) GoString() string {
	return s.String()
}

func (s *ModifyLiveAIStudioRequestMediaLayout) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *ModifyLiveAIStudioRequestMediaLayout) GetPositionX() *float32 {
	return s.PositionX
}

func (s *ModifyLiveAIStudioRequestMediaLayout) GetPositionY() *float32 {
	return s.PositionY
}

func (s *ModifyLiveAIStudioRequestMediaLayout) SetHeightNormalized(v float32) *ModifyLiveAIStudioRequestMediaLayout {
	s.HeightNormalized = &v
	return s
}

func (s *ModifyLiveAIStudioRequestMediaLayout) SetPositionX(v float32) *ModifyLiveAIStudioRequestMediaLayout {
	s.PositionX = &v
	return s
}

func (s *ModifyLiveAIStudioRequestMediaLayout) SetPositionY(v float32) *ModifyLiveAIStudioRequestMediaLayout {
	s.PositionY = &v
	return s
}

func (s *ModifyLiveAIStudioRequestMediaLayout) Validate() error {
	return dara.Validate(s)
}
