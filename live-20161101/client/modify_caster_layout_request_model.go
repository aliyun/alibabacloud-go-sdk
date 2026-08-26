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
	// The audio information.
	//
	// This parameter is required.
	AudioLayer []*ModifyCasterLayoutRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
	// The location ID (LocationId) of the video resource element.
	//
	// For the LocationId, see [Add a video source](https://help.aliyun.com/document_detail/2848020.html). The elements correspond to the VideoLayers elements in order.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV02
	BlendList []*string `json:"BlendList,omitempty" xml:"BlendList,omitempty" type:"Repeated"`
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster operation](https://help.aliyun.com/document_detail/2848009.html), check the CasterId parameter returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page of the ApsaraVideo Live console is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The layout ID. If you added the production studio layout by calling the [AddCasterLayout operation](https://help.aliyun.com/document_detail/2848025.html), check the LayoutId parameter returned by the AddCasterLayout operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The location ID (LocationId) of the audio resource element.
	//
	// For the LocationId, see [Add a video source](https://help.aliyun.com/document_detail/2848020.html). The elements correspond to the AudioLayers elements in order.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV02
	MixList []*string `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The video information.
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
	if s.AudioLayer != nil {
		for _, item := range s.AudioLayer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ModifyCasterLayoutRequestAudioLayer struct {
	// The fixed delay for the audio. This can be used for subtitle synchronization. Unit: milliseconds. Default value: **0**. Valid values: **0 to 5000**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The audio channels that can be used as volume input. Valid values:
	//
	// - **leftChannel**: left channel.
	//
	// - **rightChannel**: right channel.
	//
	// - **all*	- (default): both channels.
	//
	// example:
	//
	// all
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The normalized height ratio of the Layer element. The width of the element is proportionally scaled based on this height.
	//
	// Default value: **0**, which indicates that the element is displayed at its original size.
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
	// The element fill mode.
	//
	// - **none*	- (default): no fill. The Layer settings are configured with the image as the target.
	//
	// - **fit**: adaptive. The Layer settings are configured with the fill area (box) as the target. The image is scaled based on the original aspect ratio and centered within the fill area (box) using a long-edge alignment method. If the aspect ratio of the fill area does not match the image, the short edges are not filled (the lower Layer image is displayed. If no lower Layer is configured, the default black background is displayed).
	//
	// example:
	//
	// fit
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The fixed delay for the video. This can be used for subtitle synchronization. Unit: milliseconds. Default value: **0**. Valid values: **0 to 5000**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The normalized height ratio of the Layer element.
	//
	//
	//
	// - If the no-fill mode is used, the width of the element is proportionally scaled based on this height. Default value: **0**, which indicates that the image is displayed at its original size.
	//
	// - If the adaptive mode is used, this field is required and must be greater than **0**. It specifies the normalized height ratio of the fill area (box).
	//
	// example:
	//
	// 1
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The normalized position values `[x,y]` of the Layer element. Default value: `[0,0]`.
	//
	// 	Note: The x and y values must be normalized.
	//
	// example:
	//
	// 0.3
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	// The reference coordinate for the position of the element. Valid values:
	//
	// - **topLeft*	- (default): top-left.
	//
	// - **topRight**: top-right.
	//
	// - **bottomLeft**: bottom-left.
	//
	// - **bottomRight**: bottom-right.
	//
	// - **center**: center.
	//
	// - **topCenter**: top-center.
	//
	// - **bottomCenter**: bottom-center.
	//
	// - **leftCenter**: left-center.
	//
	// - **rightCenter**: right-center.
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The normalized width ratio of the Layer element.
	//
	// - If the no-fill mode is used, the height of the element is proportionally scaled based on this width. Default value: **0**, which indicates that the image is displayed at its original size.
	//
	// - If the adaptive mode is used, this field is required and must be greater than **0**. It specifies the normalized width ratio of the fill area (box).
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
