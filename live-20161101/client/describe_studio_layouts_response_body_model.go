// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStudioLayoutsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeStudioLayoutsResponseBody
	GetRequestId() *string
	SetStudioLayouts(v []*DescribeStudioLayoutsResponseBodyStudioLayouts) *DescribeStudioLayoutsResponseBody
	GetStudioLayouts() []*DescribeStudioLayoutsResponseBodyStudioLayouts
	SetTotal(v int32) *DescribeStudioLayoutsResponseBody
	GetTotal() *int32
}

type DescribeStudioLayoutsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The layout information.
	StudioLayouts []*DescribeStudioLayoutsResponseBodyStudioLayouts `json:"StudioLayouts,omitempty" xml:"StudioLayouts,omitempty" type:"Repeated"`
	// The number of layouts.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeStudioLayoutsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStudioLayoutsResponseBody) GetStudioLayouts() []*DescribeStudioLayoutsResponseBodyStudioLayouts {
	return s.StudioLayouts
}

func (s *DescribeStudioLayoutsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeStudioLayoutsResponseBody) SetRequestId(v string) *DescribeStudioLayoutsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBody) SetStudioLayouts(v []*DescribeStudioLayoutsResponseBodyStudioLayouts) *DescribeStudioLayoutsResponseBody {
	s.StudioLayouts = v
	return s
}

func (s *DescribeStudioLayoutsResponseBody) SetTotal(v int32) *DescribeStudioLayoutsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBody) Validate() error {
	if s.StudioLayouts != nil {
		for _, item := range s.StudioLayouts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStudioLayoutsResponseBodyStudioLayouts struct {
	// The background resource configuration.
	BgImageConfig *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig `json:"BgImageConfig,omitempty" xml:"BgImageConfig,omitempty" type:"Struct"`
	// The common layout information. This field is returned when the layout is a common layout.
	CommonConfig *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig `json:"CommonConfig,omitempty" xml:"CommonConfig,omitempty" type:"Struct"`
	// The layer order configuration.
	LayerOrderConfigList []*DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList `json:"LayerOrderConfigList,omitempty" xml:"LayerOrderConfigList,omitempty" type:"Repeated"`
	// The studio layout ID.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The studio layout name.
	//
	// example:
	//
	// 测试布局
	LayoutName *string `json:"LayoutName,omitempty" xml:"LayoutName,omitempty"`
	// The studio layout type. Valid values:
	//
	// - **common**: common layout.
	//
	// - **studio**: studio layout.
	//
	// example:
	//
	// studio
	LayoutType *string `json:"LayoutType,omitempty" xml:"LayoutType,omitempty"`
	// The multimedia input resource configuration.
	MediaInputConfigList []*DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList `json:"MediaInputConfigList,omitempty" xml:"MediaInputConfigList,omitempty" type:"Repeated"`
	// The chroma key input configuration.
	ScreenInputConfigList []*DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList `json:"ScreenInputConfigList,omitempty" xml:"ScreenInputConfigList,omitempty" type:"Repeated"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayouts) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayouts) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetBgImageConfig() *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig {
	return s.BgImageConfig
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetCommonConfig() *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig {
	return s.CommonConfig
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetLayerOrderConfigList() []*DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList {
	return s.LayerOrderConfigList
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetLayoutName() *string {
	return s.LayoutName
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetLayoutType() *string {
	return s.LayoutType
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetMediaInputConfigList() []*DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	return s.MediaInputConfigList
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) GetScreenInputConfigList() []*DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	return s.ScreenInputConfigList
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetBgImageConfig(v *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.BgImageConfig = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetCommonConfig(v *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.CommonConfig = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetLayerOrderConfigList(v []*DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.LayerOrderConfigList = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetLayoutId(v string) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.LayoutId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetLayoutName(v string) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.LayoutName = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetLayoutType(v string) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.LayoutType = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetMediaInputConfigList(v []*DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.MediaInputConfigList = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) SetScreenInputConfigList(v []*DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) *DescribeStudioLayoutsResponseBodyStudioLayouts {
	s.ScreenInputConfigList = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayouts) Validate() error {
	if s.BgImageConfig != nil {
		if err := s.BgImageConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CommonConfig != nil {
		if err := s.CommonConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LayerOrderConfigList != nil {
		for _, item := range s.LayerOrderConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MediaInputConfigList != nil {
		for _, item := range s.MediaInputConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScreenInputConfigList != nil {
		for _, item := range s.ScreenInputConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig struct {
	// The unique ID of the background material.
	//
	// example:
	//
	// k12kj31****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The URL of the material.
	//
	// example:
	//
	// http://example.org
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The location ID.
	//
	// example:
	//
	// RV01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The video-on-demand material ID.
	//
	// example:
	//
	// asdfas9df89asd8f9****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) GetId() *string {
	return s.Id
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) GetLocationId() *string {
	return s.LocationId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) GetMaterialId() *string {
	return s.MaterialId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) SetId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig {
	s.Id = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) SetImageUrl(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig {
	s.ImageUrl = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) SetLocationId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig {
	s.LocationId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) SetMaterialId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig {
	s.MaterialId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsBgImageConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig struct {
	// The channel location ID to which the video resource is bound.
	//
	// example:
	//
	// RV01
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The video resource ID.
	//
	// example:
	//
	// asdfasdfasdfasdfa****
	VideoResourceId *string `json:"VideoResourceId,omitempty" xml:"VideoResourceId,omitempty"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) GetVideoResourceId() *string {
	return s.VideoResourceId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) SetChannelId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig {
	s.ChannelId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) SetVideoResourceId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig {
	s.VideoResourceId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsCommonConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList struct {
	// The unique ID of the resource.
	//
	// example:
	//
	// k12kj31****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the resource configuration. Valid values:
	//
	// - **background**: background material.
	//
	// - **media**: multimedia material.
	//
	// example:
	//
	// media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) GetId() *string {
	return s.Id
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) GetType() *string {
	return s.Type
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) SetId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList {
	s.Id = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) SetType(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList {
	s.Type = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsLayerOrderConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList struct {
	// The channel location ID to which the video resource is bound.
	//
	// example:
	//
	// RV01
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The fill type. Default value: none.
	//
	// example:
	//
	// none
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The normalized height of the material. This is the height ratio of the material to the background. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.4
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The unique ID of the multimedia material.
	//
	// example:
	//
	// k12kj31****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The video-on-demand image material ID.
	//
	// example:
	//
	// lkajsdfsa8fd89asd8****
	ImageMaterialId *string `json:"ImageMaterialId,omitempty" xml:"ImageMaterialId,omitempty"`
	// The multimedia material number. Used for frontend display only and has no logical function.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The normalized position of the material fill area [x,y]. The values of x and y range from **0 to 1**. For example, [0.1,0.2] indicates a horizontal offset of 10% and a vertical offset of 20% from the upper-left corner.
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	// The position reference coordinate of the material. Default value: topLeft, which indicates that the position is set based on the upper-left corner.
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The video resource ID.
	//
	// example:
	//
	// asdfasdfasdfasdfa****
	VideoResourceId *string `json:"VideoResourceId,omitempty" xml:"VideoResourceId,omitempty"`
	// The normalized width of the material. This is the width ratio of the material to the background. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.4
	WidthNormalized *float32 `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetFillMode() *string {
	return s.FillMode
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetId() *string {
	return s.Id
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetImageMaterialId() *string {
	return s.ImageMaterialId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetPositionNormalized() []*float32 {
	return s.PositionNormalized
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetPositionRefer() *string {
	return s.PositionRefer
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetVideoResourceId() *string {
	return s.VideoResourceId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) GetWidthNormalized() *float32 {
	return s.WidthNormalized
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetChannelId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.ChannelId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetFillMode(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.FillMode = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetHeightNormalized(v float32) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.HeightNormalized = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.Id = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetImageMaterialId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.ImageMaterialId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetIndex(v int32) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.Index = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetPositionNormalized(v []*float32) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.PositionNormalized = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetPositionRefer(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.PositionRefer = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetVideoResourceId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.VideoResourceId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) SetWidthNormalized(v float32) *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList {
	s.WidthNormalized = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsMediaInputConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList struct {
	// The audio configuration information.
	AudioConfig *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig `json:"AudioConfig,omitempty" xml:"AudioConfig,omitempty" type:"Struct"`
	// The channel location ID to which the video resource is bound.
	//
	// example:
	//
	// RV01
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The chroma key color gamut. Valid values:
	//
	// - **blue**: blue screen background.
	//
	// - **green**: green screen background.
	//
	// - **auto**: automatic detection.
	//
	// - **complex**: real-scene chroma keying.
	//
	// example:
	//
	// green
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The normalized height. This is the height ratio of the extracted portrait to the background. Valid values: **0 to 1**.
	//
	// example:
	//
	// 0.4
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The unique ID of the chroma key source material.
	//
	// example:
	//
	// k12kj31****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The chroma key source number. Used for frontend display only and has no logical function.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Indicates whether only audio is used.
	//
	// example:
	//
	// true
	OnlyAudio *bool `json:"OnlyAudio,omitempty" xml:"OnlyAudio,omitempty"`
	// The portrait type. Valid values:
	//
	// - **0**: half-body.
	//
	// - **1**: full-body.
	//
	// example:
	//
	// 0
	PortraitType *int32 `json:"PortraitType,omitempty" xml:"PortraitType,omitempty"`
	// The position parameter, x coordinate. Valid values: **0 to 1**. The material position is based on the upper-left corner.
	//
	// example:
	//
	// 0.1
	PositionX *string `json:"PositionX,omitempty" xml:"PositionX,omitempty"`
	// The position parameter, y coordinate. Valid values: **0 to 1**. The material position is based on the upper-left corner.
	//
	// example:
	//
	// 0.2
	PositionY *string `json:"PositionY,omitempty" xml:"PositionY,omitempty"`
	// The video resource ID.
	//
	// example:
	//
	// asdfasdfasdfasdfa****
	VideoResourceId *string `json:"VideoResourceId,omitempty" xml:"VideoResourceId,omitempty"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetAudioConfig() *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig {
	return s.AudioConfig
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetColor() *string {
	return s.Color
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetId() *string {
	return s.Id
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetOnlyAudio() *bool {
	return s.OnlyAudio
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetPortraitType() *int32 {
	return s.PortraitType
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetPositionX() *string {
	return s.PositionX
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetPositionY() *string {
	return s.PositionY
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) GetVideoResourceId() *string {
	return s.VideoResourceId
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetAudioConfig(v *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.AudioConfig = v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetChannelId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.ChannelId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetColor(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.Color = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetHeightNormalized(v float32) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.HeightNormalized = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.Id = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetIndex(v int32) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.Index = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetOnlyAudio(v bool) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.OnlyAudio = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetPortraitType(v int32) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.PortraitType = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetPositionX(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.PositionX = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetPositionY(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.PositionY = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) SetVideoResourceId(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList {
	s.VideoResourceId = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigList) Validate() error {
	if s.AudioConfig != nil {
		if err := s.AudioConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig struct {
	// The corresponding channel.
	//
	// example:
	//
	// 1
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The volume.
	//
	// example:
	//
	// 1.0
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) GoString() string {
	return s.String()
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) GetValidChannel() *string {
	return s.ValidChannel
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) SetValidChannel(v string) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig {
	s.ValidChannel = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) SetVolumeRate(v float32) *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig {
	s.VolumeRate = &v
	return s
}

func (s *DescribeStudioLayoutsResponseBodyStudioLayoutsScreenInputConfigListAudioConfig) Validate() error {
	return dara.Validate(s)
}
