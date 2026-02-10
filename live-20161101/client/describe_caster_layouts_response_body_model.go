// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterLayoutsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayouts(v *DescribeCasterLayoutsResponseBodyLayouts) *DescribeCasterLayoutsResponseBody
	GetLayouts() *DescribeCasterLayoutsResponseBodyLayouts
	SetRequestId(v string) *DescribeCasterLayoutsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCasterLayoutsResponseBody
	GetTotal() *int32
}

type DescribeCasterLayoutsResponseBody struct {
	Layouts *DescribeCasterLayoutsResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// The normalized value of the width of the video layer image.
	//
	// 	- If the FillMode parameter of the video layer is set to none, the height of the video image is scaled based on this parameter. The default value is **0**, which indicates that the video image is displayed in the original size.
	//
	// 	- If the FillMode parameter of the video layer is set to fit, the value of the parameter is greater than **0**.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCasterLayoutsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBody) GetLayouts() *DescribeCasterLayoutsResponseBodyLayouts {
	return s.Layouts
}

func (s *DescribeCasterLayoutsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterLayoutsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterLayoutsResponseBody) SetLayouts(v *DescribeCasterLayoutsResponseBodyLayouts) *DescribeCasterLayoutsResponseBody {
	s.Layouts = v
	return s
}

func (s *DescribeCasterLayoutsResponseBody) SetRequestId(v string) *DescribeCasterLayoutsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBody) SetTotal(v int32) *DescribeCasterLayoutsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBody) Validate() error {
	if s.Layouts != nil {
		if err := s.Layouts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterLayoutsResponseBodyLayouts struct {
	Layout []*DescribeCasterLayoutsResponseBodyLayoutsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseBodyLayouts) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayouts) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayouts) GetLayout() []*DescribeCasterLayoutsResponseBodyLayoutsLayout {
	return s.Layout
}

func (s *DescribeCasterLayoutsResponseBodyLayouts) SetLayout(v []*DescribeCasterLayoutsResponseBodyLayoutsLayout) *DescribeCasterLayoutsResponseBodyLayouts {
	s.Layout = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayouts) Validate() error {
	if s.Layout != nil {
		for _, item := range s.Layout {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterLayoutsResponseBodyLayoutsLayout struct {
	AudioLayers *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers `json:"AudioLayers,omitempty" xml:"AudioLayers,omitempty" type:"Struct"`
	BlendList   *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList   `json:"BlendList,omitempty" xml:"BlendList,omitempty" type:"Struct"`
	LayoutId    *string                                                    `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	MixList     *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList     `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Struct"`
	VideoLayers *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers `json:"VideoLayers,omitempty" xml:"VideoLayers,omitempty" type:"Struct"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayout) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayout) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) GetAudioLayers() *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers {
	return s.AudioLayers
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) GetBlendList() *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList {
	return s.BlendList
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) GetMixList() *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList {
	return s.MixList
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) GetVideoLayers() *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers {
	return s.VideoLayers
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) SetAudioLayers(v *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers) *DescribeCasterLayoutsResponseBodyLayoutsLayout {
	s.AudioLayers = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) SetBlendList(v *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList) *DescribeCasterLayoutsResponseBodyLayoutsLayout {
	s.BlendList = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) SetLayoutId(v string) *DescribeCasterLayoutsResponseBodyLayoutsLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) SetMixList(v *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList) *DescribeCasterLayoutsResponseBodyLayoutsLayout {
	s.MixList = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) SetVideoLayers(v *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers) *DescribeCasterLayoutsResponseBodyLayoutsLayout {
	s.VideoLayers = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayout) Validate() error {
	if s.AudioLayers != nil {
		if err := s.AudioLayers.Validate(); err != nil {
			return err
		}
	}
	if s.BlendList != nil {
		if err := s.BlendList.Validate(); err != nil {
			return err
		}
	}
	if s.MixList != nil {
		if err := s.MixList.Validate(); err != nil {
			return err
		}
	}
	if s.VideoLayers != nil {
		if err := s.VideoLayers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers struct {
	AudioLayer []*DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers) GetAudioLayer() []*DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer {
	return s.AudioLayer
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers) SetAudioLayer(v []*DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers {
	s.AudioLayer = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers) Validate() error {
	if s.AudioLayer != nil {
		for _, item := range s.AudioLayer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer struct {
	FixedDelayDuration *int32   `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	ValidChannel       *string  `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	VolumeRate         *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) GetValidChannel() *string {
	return s.ValidChannel
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) SetFixedDelayDuration(v int32) *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) SetValidChannel(v string) *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) SetVolumeRate(v float32) *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList struct {
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList) GetLocationId() []*string {
	return s.LocationId
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList) SetLocationId(v []*string) *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList {
	s.LocationId = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList struct {
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList) GetLocationId() []*string {
	return s.LocationId
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList) SetLocationId(v []*string) *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList {
	s.LocationId = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers struct {
	VideoLayer []*DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer `json:"VideoLayer,omitempty" xml:"VideoLayer,omitempty" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers) GetVideoLayer() []*DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	return s.VideoLayer
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers) SetVideoLayer(v []*DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers {
	s.VideoLayer = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayers) Validate() error {
	if s.VideoLayer != nil {
		for _, item := range s.VideoLayer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer struct {
	FillMode            *string                                                                                 `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	FixedDelayDuration  *int32                                                                                  `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	HeightNormalized    *float32                                                                                `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	PositionNormalizeds *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds `json:"PositionNormalizeds,omitempty" xml:"PositionNormalizeds,omitempty" type:"Struct"`
	PositionRefer       *string                                                                                 `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	WidthNormalized     *float32                                                                                `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GetFillMode() *string {
	return s.FillMode
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GetPositionNormalizeds() *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds {
	return s.PositionNormalizeds
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GetPositionRefer() *string {
	return s.PositionRefer
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) GetWidthNormalized() *float32 {
	return s.WidthNormalized
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) SetFillMode(v string) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	s.FillMode = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) SetFixedDelayDuration(v int32) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) SetHeightNormalized(v float32) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	s.HeightNormalized = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) SetPositionNormalizeds(v *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	s.PositionNormalizeds = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) SetPositionRefer(v string) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	s.PositionRefer = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) SetWidthNormalized(v float32) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer {
	s.WidthNormalized = &v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer) Validate() error {
	if s.PositionNormalizeds != nil {
		if err := s.PositionNormalizeds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds struct {
	Position []*float32 `json:"Position,omitempty" xml:"Position,omitempty" type:"Repeated"`
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) GetPosition() []*float32 {
	return s.Position
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) SetPosition(v []*float32) *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds {
	s.Position = v
	return s
}

func (s *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds) Validate() error {
	return dara.Validate(s)
}
