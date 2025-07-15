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
	// The layouts.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeCasterLayoutsResponseBodyLayoutsLayout struct {
	// The configurations of the audio layers.
	AudioLayers *DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayers `json:"AudioLayers,omitempty" xml:"AudioLayers,omitempty" type:"Struct"`
	// The location IDs of the video layers, which are in the same order as the video layers.
	BlendList *DescribeCasterLayoutsResponseBodyLayoutsLayoutBlendList `json:"BlendList,omitempty" xml:"BlendList,omitempty" type:"Struct"`
	// The ID of the layout.
	//
	// example:
	//
	// 72d2ec7a-4cd7-4a01-974b-7cd53947****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The location IDs of the audio layers, which are in the same order as the audio layers.
	MixList *DescribeCasterLayoutsResponseBodyLayoutsLayoutMixList `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Struct"`
	// The configurations of the video layers, which are in the default array sequence.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutAudioLayersAudioLayer struct {
	// The fixed delay of the audio layer. This parameter is used to synchronize the audio with subtitles.
	//
	// Unit: milliseconds. Default value: **0**. Valid values: **0 to 5000**.
	//
	// example:
	//
	// 20
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The sound channels that are used for volume input in the audio layer. Valid values:
	//
	// 	- **leftChannel**: the left channel
	//
	// 	- **rightChannel**: the right channel
	//
	// 	- **all*	- (default): both the left and right channels
	//
	// example:
	//
	// all
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The normalized value of the height of the audio layer. The width of the audio layer is proportionally scaled based on this parameter.
	//
	// >  The default value is **0**, which indicates that the audio layer is not scaled.
	//
	// example:
	//
	// 1
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
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
	return dara.Validate(s)
}

type DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayer struct {
	// The scaling mode of the video layer. Valid values:
	//
	// 	- **none*	- (default): specifies that the video layer is not scaled. The video layer is displayed based on its original size.
	//
	// 	- **fit**: specifies that the video layer is adapted to the fill area. The video layer is displayed based on the fill area. In this case, the video layer is scaled proportionally, with its original aspect ratio retained. The video layer is placed in the center, with its longer sides aligned with the fill area. If the aspect ratio of the video layer is different from that of the fill area, the content of the lower layer is displayed alongside the shorter sides. If there is no lower layer, black bars are displayed instead.
	//
	// example:
	//
	// fit
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The fixed delay of the video layer. This parameter is used to synchronize the video with subtitles.
	//
	// Unit: milliseconds. Default value: **0**. Valid values: **0 to 5000**.
	//
	// example:
	//
	// 20
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The normalized value of the height of the video layer.
	//
	// 	- If the FillMode parameter of the video layer is set to none, the width of the video layer is proportionally scaled based on this parameter. The default value is **0**, which indicates that the video layer is not scaled.
	//
	// 	- If the FillMode parameter of the video layer is set to fit, the value of this parameter is greater than **0**.
	//
	// example:
	//
	// 0.5
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The normalized value of the position of the video layer, in the format of `[x,y]`. Default value: `[0,0]`.
	//
	// >  The values of x and y are normalized.
	PositionNormalizeds *DescribeCasterLayoutsResponseBodyLayoutsLayoutVideoLayersVideoLayerPositionNormalizeds `json:"PositionNormalizeds,omitempty" xml:"PositionNormalizeds,omitempty" type:"Struct"`
	// The reference coordinates of the video layer. Valid values:
	//
	// 	- **topLeft*	- (default): the upper-left corner
	//
	// 	- **topRight**: the upper-right corner
	//
	// 	- **bottomLeft**: the lower-left corner
	//
	// 	- **bottomRight**: the lower-right corner
	//
	// 	- **center**: the center
	//
	// 	- **topCenter**: the upper center
	//
	// 	- **bottomCenter**: the lower center
	//
	// 	- **leftCenter**: the left center
	//
	// 	- **rightCenter**: the right center
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The normalized value of the width of the video layer.
	//
	// 	- If the FillMode parameter of the video layer is set to none, the height of the video layer is scaled based on this parameter. The default value is **0**, which indicates that the video layer is not scaled.
	//
	// 	- If the FillMode parameter of the video layer is set to fit, the value of this parameter is greater than **0**.
	//
	// example:
	//
	// 0.5
	WidthNormalized *float32 `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
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
	return dara.Validate(s)
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
