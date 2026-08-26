// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveAIStudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundResourceId(v string) *CreateLiveAIStudioRequest
	GetBackgroundResourceId() *string
	SetBackgroundResourceUrl(v string) *CreateLiveAIStudioRequest
	GetBackgroundResourceUrl() *string
	SetBackgroundType(v string) *CreateLiveAIStudioRequest
	GetBackgroundType() *string
	SetDescription(v string) *CreateLiveAIStudioRequest
	GetDescription() *string
	SetHeight(v int32) *CreateLiveAIStudioRequest
	GetHeight() *int32
	SetMattingLayout(v *CreateLiveAIStudioRequestMattingLayout) *CreateLiveAIStudioRequest
	GetMattingLayout() *CreateLiveAIStudioRequestMattingLayout
	SetMattingType(v string) *CreateLiveAIStudioRequest
	GetMattingType() *string
	SetMediaLayout(v *CreateLiveAIStudioRequestMediaLayout) *CreateLiveAIStudioRequest
	GetMediaLayout() *CreateLiveAIStudioRequestMediaLayout
	SetMediaResourceId(v string) *CreateLiveAIStudioRequest
	GetMediaResourceId() *string
	SetMediaResourceUrl(v string) *CreateLiveAIStudioRequest
	GetMediaResourceUrl() *string
	SetMediaType(v string) *CreateLiveAIStudioRequest
	GetMediaType() *string
	SetOwnerId(v int64) *CreateLiveAIStudioRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateLiveAIStudioRequest
	GetRegionId() *string
	SetStudioName(v string) *CreateLiveAIStudioRequest
	GetStudioName() *string
	SetWidth(v int32) *CreateLiveAIStudioRequest
	GetWidth() *int32
}

type CreateLiveAIStudioRequest struct {
	// The video-on-demand resource ID of the background material. Obtain this value from the ApsaraVideo VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	BackgroundResourceId *string `json:"BackgroundResourceId,omitempty" xml:"BackgroundResourceId,omitempty"`
	// The access URL of the background material. Specify either this parameter or the resource ID.
	//
	// example:
	//
	// https://xxx.com/1.mp4
	BackgroundResourceUrl *string `json:"BackgroundResourceUrl,omitempty" xml:"BackgroundResourceUrl,omitempty"`
	// The type of the background material. Valid values:
	//
	// - VOD: video-on-demand video.
	//
	// - PIC: image.
	//
	// - LIVE: live stream.
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
	// The height of the preview screen. Unit: px.
	//
	// The width × height of the preview screen supports only the following specifications:
	//
	// - Landscape low definition 360P: 640 × 360
	//
	// - Portrait low definition 360P: 360 × 640
	//
	// - Landscape standard definition 480P: 854 × 480
	//
	// - Portrait standard definition 480P: 480 × 854
	//
	// - Landscape high definition 720P: 1280 × 720
	//
	// - Portrait high definition 720P: 720 × 1280
	//
	// - Landscape ultra-high definition 1080P: 1920 × 1080
	//
	// - Portrait ultra-high definition 1080P: 1080 × 1920.
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The layout position information of the source stream after image matting.
	//
	// This parameter is required.
	MattingLayout *CreateLiveAIStudioRequestMattingLayout `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty" type:"Struct"`
	// The image matting type. Valid values:
	//
	// - green: green screen matting.
	//
	// - blue: blue screen matting.
	//
	// - complex: real-scene matting.
	//
	// This parameter is required.
	//
	// example:
	//
	// complex
	MattingType *string `json:"MattingType,omitempty" xml:"MattingType,omitempty"`
	// The layout position information of the multimedia material.
	MediaLayout *CreateLiveAIStudioRequestMediaLayout `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty" type:"Struct"`
	// The video-on-demand resource ID of the multimedia material. Obtain this value from the ApsaraVideo VOD console.
	//
	// example:
	//
	// d0eb493192c771efba644531858c0102
	MediaResourceId *string `json:"MediaResourceId,omitempty" xml:"MediaResourceId,omitempty"`
	// The access URL of the multimedia material. Specify either this parameter or the resource ID.
	//
	// example:
	//
	// https://xxx.com/2.mp4
	MediaResourceUrl *string `json:"MediaResourceUrl,omitempty" xml:"MediaResourceUrl,omitempty"`
	// The type of the multimedia material. Valid values:
	//
	// - VOD: video-on-demand video.
	//
	// - PIC: image.
	//
	// - LIVE: live stream.
	//
	// example:
	//
	// VOD
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the virtual studio template. The name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// stu02
	StudioName *string `json:"StudioName,omitempty" xml:"StudioName,omitempty"`
	// The width of the preview screen. Unit: px.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateLiveAIStudioRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveAIStudioRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveAIStudioRequest) GetBackgroundResourceId() *string {
	return s.BackgroundResourceId
}

func (s *CreateLiveAIStudioRequest) GetBackgroundResourceUrl() *string {
	return s.BackgroundResourceUrl
}

func (s *CreateLiveAIStudioRequest) GetBackgroundType() *string {
	return s.BackgroundType
}

func (s *CreateLiveAIStudioRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLiveAIStudioRequest) GetHeight() *int32 {
	return s.Height
}

func (s *CreateLiveAIStudioRequest) GetMattingLayout() *CreateLiveAIStudioRequestMattingLayout {
	return s.MattingLayout
}

func (s *CreateLiveAIStudioRequest) GetMattingType() *string {
	return s.MattingType
}

func (s *CreateLiveAIStudioRequest) GetMediaLayout() *CreateLiveAIStudioRequestMediaLayout {
	return s.MediaLayout
}

func (s *CreateLiveAIStudioRequest) GetMediaResourceId() *string {
	return s.MediaResourceId
}

func (s *CreateLiveAIStudioRequest) GetMediaResourceUrl() *string {
	return s.MediaResourceUrl
}

func (s *CreateLiveAIStudioRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *CreateLiveAIStudioRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLiveAIStudioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLiveAIStudioRequest) GetStudioName() *string {
	return s.StudioName
}

func (s *CreateLiveAIStudioRequest) GetWidth() *int32 {
	return s.Width
}

func (s *CreateLiveAIStudioRequest) SetBackgroundResourceId(v string) *CreateLiveAIStudioRequest {
	s.BackgroundResourceId = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetBackgroundResourceUrl(v string) *CreateLiveAIStudioRequest {
	s.BackgroundResourceUrl = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetBackgroundType(v string) *CreateLiveAIStudioRequest {
	s.BackgroundType = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetDescription(v string) *CreateLiveAIStudioRequest {
	s.Description = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetHeight(v int32) *CreateLiveAIStudioRequest {
	s.Height = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetMattingLayout(v *CreateLiveAIStudioRequestMattingLayout) *CreateLiveAIStudioRequest {
	s.MattingLayout = v
	return s
}

func (s *CreateLiveAIStudioRequest) SetMattingType(v string) *CreateLiveAIStudioRequest {
	s.MattingType = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetMediaLayout(v *CreateLiveAIStudioRequestMediaLayout) *CreateLiveAIStudioRequest {
	s.MediaLayout = v
	return s
}

func (s *CreateLiveAIStudioRequest) SetMediaResourceId(v string) *CreateLiveAIStudioRequest {
	s.MediaResourceId = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetMediaResourceUrl(v string) *CreateLiveAIStudioRequest {
	s.MediaResourceUrl = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetMediaType(v string) *CreateLiveAIStudioRequest {
	s.MediaType = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetOwnerId(v int64) *CreateLiveAIStudioRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetRegionId(v string) *CreateLiveAIStudioRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetStudioName(v string) *CreateLiveAIStudioRequest {
	s.StudioName = &v
	return s
}

func (s *CreateLiveAIStudioRequest) SetWidth(v int32) *CreateLiveAIStudioRequest {
	s.Width = &v
	return s
}

func (s *CreateLiveAIStudioRequest) Validate() error {
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

type CreateLiveAIStudioRequestMattingLayout struct {
	// The normalized height of the material, which is the ratio of the material height to the background height. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The x-coordinate of the position. Valid values: **0 to 1**. The position of the material is based on the upper-left corner as the reference point.
	//
	// example:
	//
	// 0.3
	PositionX *float32 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// The y-coordinate of the position. Valid values: **0 to 1**. The position of the material is based on the upper-left corner as the reference point.
	//
	// example:
	//
	// 0.3
	PositionY *float32 `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
}

func (s CreateLiveAIStudioRequestMattingLayout) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveAIStudioRequestMattingLayout) GoString() string {
	return s.String()
}

func (s *CreateLiveAIStudioRequestMattingLayout) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *CreateLiveAIStudioRequestMattingLayout) GetPositionX() *float32 {
	return s.PositionX
}

func (s *CreateLiveAIStudioRequestMattingLayout) GetPositionY() *float32 {
	return s.PositionY
}

func (s *CreateLiveAIStudioRequestMattingLayout) SetHeightNormalized(v float32) *CreateLiveAIStudioRequestMattingLayout {
	s.HeightNormalized = &v
	return s
}

func (s *CreateLiveAIStudioRequestMattingLayout) SetPositionX(v float32) *CreateLiveAIStudioRequestMattingLayout {
	s.PositionX = &v
	return s
}

func (s *CreateLiveAIStudioRequestMattingLayout) SetPositionY(v float32) *CreateLiveAIStudioRequestMattingLayout {
	s.PositionY = &v
	return s
}

func (s *CreateLiveAIStudioRequestMattingLayout) Validate() error {
	return dara.Validate(s)
}

type CreateLiveAIStudioRequestMediaLayout struct {
	// The normalized height of the material, which is the ratio of the material height to the background height. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The x-coordinate of the position. Valid values: **0 to 1**. The position of the material is based on the upper-left corner as the reference point.
	//
	// example:
	//
	// 0
	PositionX *float32 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// The y-coordinate of the position. Valid values: **0 to 1**. The position of the material is based on the upper-left corner as the reference point.
	//
	// example:
	//
	// 0
	PositionY *float32 `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
}

func (s CreateLiveAIStudioRequestMediaLayout) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveAIStudioRequestMediaLayout) GoString() string {
	return s.String()
}

func (s *CreateLiveAIStudioRequestMediaLayout) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *CreateLiveAIStudioRequestMediaLayout) GetPositionX() *float32 {
	return s.PositionX
}

func (s *CreateLiveAIStudioRequestMediaLayout) GetPositionY() *float32 {
	return s.PositionY
}

func (s *CreateLiveAIStudioRequestMediaLayout) SetHeightNormalized(v float32) *CreateLiveAIStudioRequestMediaLayout {
	s.HeightNormalized = &v
	return s
}

func (s *CreateLiveAIStudioRequestMediaLayout) SetPositionX(v float32) *CreateLiveAIStudioRequestMediaLayout {
	s.PositionX = &v
	return s
}

func (s *CreateLiveAIStudioRequestMediaLayout) SetPositionY(v float32) *CreateLiveAIStudioRequestMediaLayout {
	s.PositionY = &v
	return s
}

func (s *CreateLiveAIStudioRequestMediaLayout) Validate() error {
	return dara.Validate(s)
}
