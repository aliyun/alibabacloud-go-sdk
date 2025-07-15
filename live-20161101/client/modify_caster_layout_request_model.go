// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioLayer(v []*ModifyCasterLayoutRequestAudioLayer) *ModifyCasterLayoutRequest
	GetAudioLayer() []*ModifyCasterLayoutRequestAudioLayer
	SetBlendList(v []*string) *ModifyCasterLayoutRequest
	GetBlendList() []*string
	SetCasterId(v string) *ModifyCasterLayoutRequest
	GetCasterId() *string
	SetLayoutId(v string) *ModifyCasterLayoutRequest
	GetLayoutId() *string
	SetMixList(v []*string) *ModifyCasterLayoutRequest
	GetMixList() []*string
	SetOwnerId(v int64) *ModifyCasterLayoutRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCasterLayoutRequest
	GetRegionId() *string
	SetVideoLayer(v []*ModifyCasterLayoutRequestVideoLayer) *ModifyCasterLayoutRequest
	GetVideoLayer() []*ModifyCasterLayoutRequestVideoLayer
}

type ModifyCasterLayoutRequest struct {
	// The audio layers.
	//
	// This parameter is required.
	AudioLayer []*ModifyCasterLayoutRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
	// The location IDs of the video layers, which are in the same order as the video layers.
	//
	// For more information, see [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RV02
	BlendList []*string `json:"BlendList,omitempty" xml:"BlendList,omitempty" type:"Repeated"`
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the layout. If the layout was added by calling the [AddCasterLayout](https://help.aliyun.com/document_detail/2848025.html) operation, check the value of the response parameter LayoutId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The location IDs of the audio layers, which are in the same order as the audio layers.
	//
	// For more information, see [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// RV02
	MixList  []*string `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The video layers.
	//
	// This parameter is required.
	VideoLayer []*ModifyCasterLayoutRequestVideoLayer `json:"VideoLayer,omitempty" xml:"VideoLayer,omitempty" type:"Repeated"`
}

func (s ModifyCasterLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterLayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutRequest) GetAudioLayer() []*ModifyCasterLayoutRequestAudioLayer {
	return s.AudioLayer
}

func (s *ModifyCasterLayoutRequest) GetBlendList() []*string {
	return s.BlendList
}

func (s *ModifyCasterLayoutRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterLayoutRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *ModifyCasterLayoutRequest) GetMixList() []*string {
	return s.MixList
}

func (s *ModifyCasterLayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCasterLayoutRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCasterLayoutRequest) GetVideoLayer() []*ModifyCasterLayoutRequestVideoLayer {
	return s.VideoLayer
}

func (s *ModifyCasterLayoutRequest) SetAudioLayer(v []*ModifyCasterLayoutRequestAudioLayer) *ModifyCasterLayoutRequest {
	s.AudioLayer = v
	return s
}

func (s *ModifyCasterLayoutRequest) SetBlendList(v []*string) *ModifyCasterLayoutRequest {
	s.BlendList = v
	return s
}

func (s *ModifyCasterLayoutRequest) SetCasterId(v string) *ModifyCasterLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterLayoutRequest) SetLayoutId(v string) *ModifyCasterLayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *ModifyCasterLayoutRequest) SetMixList(v []*string) *ModifyCasterLayoutRequest {
	s.MixList = v
	return s
}

func (s *ModifyCasterLayoutRequest) SetOwnerId(v int64) *ModifyCasterLayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCasterLayoutRequest) SetRegionId(v string) *ModifyCasterLayoutRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCasterLayoutRequest) SetVideoLayer(v []*ModifyCasterLayoutRequestVideoLayer) *ModifyCasterLayoutRequest {
	s.VideoLayer = v
	return s
}

func (s *ModifyCasterLayoutRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyCasterLayoutRequestAudioLayer struct {
	// The fixed delay of the audio layer. This parameter is used to synchronize the audio with subtitles. Unit: milliseconds. Default value: **0**. Valid values: **0 to 5000**.
	//
	// example:
	//
	// 5000
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
	// The default value is **0**, which indicates that the audio layer is not scaled.
	//
	// example:
	//
	// 1
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s ModifyCasterLayoutRequestAudioLayer) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterLayoutRequestAudioLayer) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutRequestAudioLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *ModifyCasterLayoutRequestAudioLayer) GetValidChannel() *string {
	return s.ValidChannel
}

func (s *ModifyCasterLayoutRequestAudioLayer) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *ModifyCasterLayoutRequestAudioLayer) SetFixedDelayDuration(v int32) *ModifyCasterLayoutRequestAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *ModifyCasterLayoutRequestAudioLayer) SetValidChannel(v string) *ModifyCasterLayoutRequestAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *ModifyCasterLayoutRequestAudioLayer) SetVolumeRate(v float32) *ModifyCasterLayoutRequestAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *ModifyCasterLayoutRequestAudioLayer) Validate() error {
	return dara.Validate(s)
}

type ModifyCasterLayoutRequestVideoLayer struct {
	// The scaling mode of the video layer. Valid values:
	//
	// 	- **none*	- (default): indicates that the video layer is not scaled. The video layer is displayed based on its original size.
	//
	// 	- **fit**: indicates that the video layer is adapted to the fill area. In this case, the video layer is scaled proportionally, with its original aspect ratio retained. The video layer is placed in the center, with its longer sides aligned with the fill area. If the aspect ratio of the video layer is different from that of the fill area, the content of the lower layer is displayed alongside the shorter sides. If there is no lower layer, black bars are displayed instead.
	//
	// example:
	//
	// fit
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The fixed delay of the video layer. This parameter is used to synchronize the video with subtitles. Unit: milliseconds. Default value: **0**. Valid values: **0 to 5000**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The normalized value of the height of the video layer.
	//
	// 	- If the FillMode parameter of the video layer is set to none, the width of the video layer is proportionally scaled based on this parameter. The default value is **0**, which indicates that the video layer is not scaled.
	//
	// 	- If the FillMode parameter of the video layer is set to fit, the value of this parameter is greater than **0**.
	//
	// example:
	//
	// 1
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The normalized value of the position of the video layer, in the format of `[x,y]`. Default value: `[0,0]`.
	//
	// >  The values of x and y are normalized.
	//
	// example:
	//
	// 0.3
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
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
	// 	- If the FillMode parameter of the video layer is set to none, the height of the video layer is proportionally scaled based on this parameter. The default value is **0**, which indicates that the video layer is not scaled.
	//
	// 	- If the FillMode parameter of the video layer is set to fit, the value of this parameter is greater than **0**.
	//
	// example:
	//
	// 1
	WidthNormalized *float32 `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
}

func (s ModifyCasterLayoutRequestVideoLayer) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterLayoutRequestVideoLayer) GoString() string {
	return s.String()
}

func (s *ModifyCasterLayoutRequestVideoLayer) GetFillMode() *string {
	return s.FillMode
}

func (s *ModifyCasterLayoutRequestVideoLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *ModifyCasterLayoutRequestVideoLayer) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *ModifyCasterLayoutRequestVideoLayer) GetPositionNormalized() []*float32 {
	return s.PositionNormalized
}

func (s *ModifyCasterLayoutRequestVideoLayer) GetPositionRefer() *string {
	return s.PositionRefer
}

func (s *ModifyCasterLayoutRequestVideoLayer) GetWidthNormalized() *float32 {
	return s.WidthNormalized
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetFillMode(v string) *ModifyCasterLayoutRequestVideoLayer {
	s.FillMode = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetFixedDelayDuration(v int32) *ModifyCasterLayoutRequestVideoLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetHeightNormalized(v float32) *ModifyCasterLayoutRequestVideoLayer {
	s.HeightNormalized = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetPositionNormalized(v []*float32) *ModifyCasterLayoutRequestVideoLayer {
	s.PositionNormalized = v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetPositionRefer(v string) *ModifyCasterLayoutRequestVideoLayer {
	s.PositionRefer = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) SetWidthNormalized(v float32) *ModifyCasterLayoutRequestVideoLayer {
	s.WidthNormalized = &v
	return s
}

func (s *ModifyCasterLayoutRequestVideoLayer) Validate() error {
	return dara.Validate(s)
}
