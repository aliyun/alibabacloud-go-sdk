// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioLayer(v []*AddCasterLayoutRequestAudioLayer) *AddCasterLayoutRequest
	GetAudioLayer() []*AddCasterLayoutRequestAudioLayer
	SetBlendList(v []*string) *AddCasterLayoutRequest
	GetBlendList() []*string
	SetCasterId(v string) *AddCasterLayoutRequest
	GetCasterId() *string
	SetMixList(v []*string) *AddCasterLayoutRequest
	GetMixList() []*string
	SetOwnerId(v int64) *AddCasterLayoutRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCasterLayoutRequest
	GetRegionId() *string
	SetVideoLayer(v []*AddCasterLayoutRequestVideoLayer) *AddCasterLayoutRequest
	GetVideoLayer() []*AddCasterLayoutRequestVideoLayer
}

type AddCasterLayoutRequest struct {
	// The audio layouts.
	//
	// This parameter is required.
	AudioLayer []*AddCasterLayoutRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
	// The location IDs of the video sources. The order of the location IDs corresponds to the order of the video layers specified in the **VideoLayer*	- parameter. For more information about location IDs, see [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html).
	//
	// For LocationId, see [Add a video source](https://help.aliyun.com/document_detail/2848020.html). This ID corresponds to the order of the VideoLayers elements.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV01
	BlendList []*string `json:"BlendList,omitempty" xml:"BlendList,omitempty" type:"Repeated"`
	// The ID of the production studio.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, the CasterId is returned in the response.
	//
	// - If you create a production studio in the LIVE console, go to **Production Studio*	- > **Cloud Production Studio*	- to view the name of the production studio.
	//
	// > The name of the production studio on the Cloud Production Studio page is the ID of the production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The location IDs of the audio sources. The order of the location IDs corresponds to the order of the audio layers specified in the **AudioLayer*	- parameter. For more information about location IDs, see [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html).
	//
	// For \\`LocationId\\`, see [Add a video source](https://help.aliyun.com/document_detail/2848020.html). It corresponds to the order of the \\`AudioLayers\\` elements.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV01
	MixList []*string `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The video layouts.
	//
	// This parameter is required.
	VideoLayer []*AddCasterLayoutRequestVideoLayer `json:"VideoLayer,omitempty" xml:"VideoLayer,omitempty" type:"Repeated"`
}

func (s AddCasterLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasterLayoutRequest) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutRequest) GetAudioLayer() []*AddCasterLayoutRequestAudioLayer {
	return s.AudioLayer
}

func (s *AddCasterLayoutRequest) GetBlendList() []*string {
	return s.BlendList
}

func (s *AddCasterLayoutRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddCasterLayoutRequest) GetMixList() []*string {
	return s.MixList
}

func (s *AddCasterLayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCasterLayoutRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCasterLayoutRequest) GetVideoLayer() []*AddCasterLayoutRequestVideoLayer {
	return s.VideoLayer
}

func (s *AddCasterLayoutRequest) SetAudioLayer(v []*AddCasterLayoutRequestAudioLayer) *AddCasterLayoutRequest {
	s.AudioLayer = v
	return s
}

func (s *AddCasterLayoutRequest) SetBlendList(v []*string) *AddCasterLayoutRequest {
	s.BlendList = v
	return s
}

func (s *AddCasterLayoutRequest) SetCasterId(v string) *AddCasterLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterLayoutRequest) SetMixList(v []*string) *AddCasterLayoutRequest {
	s.MixList = v
	return s
}

func (s *AddCasterLayoutRequest) SetOwnerId(v int64) *AddCasterLayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCasterLayoutRequest) SetRegionId(v string) *AddCasterLayoutRequest {
	s.RegionId = &v
	return s
}

func (s *AddCasterLayoutRequest) SetVideoLayer(v []*AddCasterLayoutRequestVideoLayer) *AddCasterLayoutRequest {
	s.VideoLayer = v
	return s
}

func (s *AddCasterLayoutRequest) Validate() error {
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

type AddCasterLayoutRequestAudioLayer struct {
	// The fixed latency for the audio layer. Use this parameter to synchronize the audio with captions. Unit: milliseconds. Default value: 0. Valid values: **0*	- to **5000**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The sound channels that are used for audio input. Valid values:
	//
	// - **leftChannel**: Left channel.
	//
	// - **rightChannel**: Right channel.
	//
	// - **all*	- (default): Both channels.
	//
	// example:
	//
	// all
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The volume multiplication factor for the audio stream. Valid values: 0 to **10.0**.
	//
	// - **1.0*	- (default): The original volume is used.
	//
	// - A value less than **1*	- decreases the volume.
	//
	// - A value greater than **1*	- increases the volume.
	//
	// example:
	//
	// 1.0
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s AddCasterLayoutRequestAudioLayer) String() string {
	return dara.Prettify(s)
}

func (s AddCasterLayoutRequestAudioLayer) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutRequestAudioLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *AddCasterLayoutRequestAudioLayer) GetValidChannel() *string {
	return s.ValidChannel
}

func (s *AddCasterLayoutRequestAudioLayer) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *AddCasterLayoutRequestAudioLayer) SetFixedDelayDuration(v int32) *AddCasterLayoutRequestAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *AddCasterLayoutRequestAudioLayer) SetValidChannel(v string) *AddCasterLayoutRequestAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *AddCasterLayoutRequestAudioLayer) SetVolumeRate(v float32) *AddCasterLayoutRequestAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *AddCasterLayoutRequestAudioLayer) Validate() error {
	return dara.Validate(s)
}

type AddCasterLayoutRequestVideoLayer struct {
	// The fill mode of the element. Valid values:
	//
	// - **none*	- (default): No scaling. The video is displayed in its original size.
	//
	// - **fit**: The video is scaled to fit the fill area while maintaining its aspect ratio. The video is centered in the fill area. If the aspect ratio of the fill area is different from that of the video, the area along the shorter edge is not filled. This area displays the video of the underlying layer. If no underlying layer is configured, this area is black.
	//
	// example:
	//
	// fit
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The fixed latency for the video layer. Use this parameter to synchronize the video with captions. Unit: milliseconds. Default value: 0. Valid values: **0*	- to **5000**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The normalized height of the layer.
	//
	// - If you set FillMode to none, the width of the layer is scaled in proportion to the height. The default value is **0**. A value of 0 indicates that the video is displayed in its original size.
	//
	// - If you set FillMode to fit, this parameter is required and its value must be greater than **0**. The value specifies the normalized height of the fill area.
	//
	// example:
	//
	// 1
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The position of the video layer. The value is a normalized coordinate `[x,y]`. Default value: `[0,0]`.
	//
	// Note: The x and y coordinates must be normalized.
	//
	// example:
	//
	// 0.3
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	// The reference point for the position of the layer. Valid values:
	//
	// - **topLeft*	- (default): Top-left.
	//
	// - **topRight**: Top-right.
	//
	// - **bottomLeft**: Bottom-left.
	//
	// - **bottomRight**: Bottom-right.
	//
	// - **center**: Center.
	//
	// - **topCenter**: Top-center.
	//
	// - **bottomCenter**: Bottom-center.
	//
	// - **leftCenter**: Left-center.
	//
	// - **rightCenter**: Right-center.
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The normalized width of the layer.
	//
	// - If you set FillMode to none, the height of the layer is scaled in proportion to the width. The default value is **0**. A value of 0 indicates that the video is displayed in its original size.
	//
	// - If you set FillMode to fit, this parameter is required and its value must be greater than **0**. The value specifies the normalized width of the fill area.
	//
	// example:
	//
	// 1
	WidthNormalized *float32 `json:"WidthNormalized,omitempty" xml:"WidthNormalized,omitempty"`
}

func (s AddCasterLayoutRequestVideoLayer) String() string {
	return dara.Prettify(s)
}

func (s AddCasterLayoutRequestVideoLayer) GoString() string {
	return s.String()
}

func (s *AddCasterLayoutRequestVideoLayer) GetFillMode() *string {
	return s.FillMode
}

func (s *AddCasterLayoutRequestVideoLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *AddCasterLayoutRequestVideoLayer) GetHeightNormalized() *float32 {
	return s.HeightNormalized
}

func (s *AddCasterLayoutRequestVideoLayer) GetPositionNormalized() []*float32 {
	return s.PositionNormalized
}

func (s *AddCasterLayoutRequestVideoLayer) GetPositionRefer() *string {
	return s.PositionRefer
}

func (s *AddCasterLayoutRequestVideoLayer) GetWidthNormalized() *float32 {
	return s.WidthNormalized
}

func (s *AddCasterLayoutRequestVideoLayer) SetFillMode(v string) *AddCasterLayoutRequestVideoLayer {
	s.FillMode = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetFixedDelayDuration(v int32) *AddCasterLayoutRequestVideoLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetHeightNormalized(v float32) *AddCasterLayoutRequestVideoLayer {
	s.HeightNormalized = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetPositionNormalized(v []*float32) *AddCasterLayoutRequestVideoLayer {
	s.PositionNormalized = v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetPositionRefer(v string) *AddCasterLayoutRequestVideoLayer {
	s.PositionRefer = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) SetWidthNormalized(v float32) *AddCasterLayoutRequestVideoLayer {
	s.WidthNormalized = &v
	return s
}

func (s *AddCasterLayoutRequestVideoLayer) Validate() error {
	return dara.Validate(s)
}
