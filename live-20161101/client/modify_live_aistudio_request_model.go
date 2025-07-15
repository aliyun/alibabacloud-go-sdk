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
	MattingLayout *ModifyLiveAIStudioRequestMattingLayout `json:"MattingLayout,omitempty" xml:"MattingLayout,omitempty" type:"Struct"`
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
	MediaLayout *ModifyLiveAIStudioRequestMediaLayout `json:"MediaLayout,omitempty" xml:"MediaLayout,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type ModifyLiveAIStudioRequestMattingLayout struct {
	// The normalized value of the height. The value indicates the ratio of the material height to the height of the background. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The x-coordinate of the material. Valid values: **0 to 1**. The upper-left corner is used as the coordinate origin for the material.
	//
	// example:
	//
	// 0
	PositionX *float32 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// The y-coordinate of the material. Valid values: **0 to 1**. The upper-left corner is used as the coordinate origin for the material.
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
	// The normalized value of the material height. The value indicates the ratio of the material height to the height of the background. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The x-coordinate of the material. Valid values: **0 to 1**. The upper-left corner is used as the coordinate origin for the material.
	//
	// example:
	//
	// 0
	PositionX *float32 `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// The y-coordinate of the material. Valid values: **0 to 1**. The upper-left corner is used as the coordinate origin for the material.
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
