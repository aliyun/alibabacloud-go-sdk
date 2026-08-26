// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterSceneAudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioLayer(v []*UpdateCasterSceneAudioRequestAudioLayer) *UpdateCasterSceneAudioRequest
	GetAudioLayer() []*UpdateCasterSceneAudioRequestAudioLayer
	SetCasterId(v string) *UpdateCasterSceneAudioRequest
	GetCasterId() *string
	SetFollowEnable(v int32) *UpdateCasterSceneAudioRequest
	GetFollowEnable() *int32
	SetMixList(v []*string) *UpdateCasterSceneAudioRequest
	GetMixList() []*string
	SetOwnerId(v int64) *UpdateCasterSceneAudioRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateCasterSceneAudioRequest
	GetRegionId() *string
	SetSceneId(v string) *UpdateCasterSceneAudioRequest
	GetSceneId() *string
}

type UpdateCasterSceneAudioRequest struct {
	// The audio configurations.
	AudioLayer []*UpdateCasterSceneAudioRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
	// The ID of the production studio.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, you can obtain the ID from the CasterId parameter in the response.
	//
	// - If you create a production studio in the LIVE console, go to the **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the ID of the production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// Specifies whether to enable the AFV mode. If you leave this parameter empty, the last configuration is retained. Valid values:
	//
	// - **0**: audio mixing mode.
	//
	// - **1**: audio-follows-video mode.
	//
	// example:
	//
	// 1
	FollowEnable *int32 `json:"FollowEnable,omitempty" xml:"FollowEnable,omitempty"`
	// The list of associated location IDs. The order of the location IDs must be the same as the order of the audio layers.
	//
	// example:
	//
	// RV01
	MixList []*string `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
	OwnerId *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene. If you query the list of scenes in a production studio by calling the [DescribeCasterScenes](https://help.aliyun.com/document_detail/2848039.html) operation, you can obtain the ID from the ComponentId parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e1****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdateCasterSceneAudioRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneAudioRequest) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioRequest) GetAudioLayer() []*UpdateCasterSceneAudioRequestAudioLayer {
	return s.AudioLayer
}

func (s *UpdateCasterSceneAudioRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *UpdateCasterSceneAudioRequest) GetFollowEnable() *int32 {
	return s.FollowEnable
}

func (s *UpdateCasterSceneAudioRequest) GetMixList() []*string {
	return s.MixList
}

func (s *UpdateCasterSceneAudioRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCasterSceneAudioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCasterSceneAudioRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateCasterSceneAudioRequest) SetAudioLayer(v []*UpdateCasterSceneAudioRequestAudioLayer) *UpdateCasterSceneAudioRequest {
	s.AudioLayer = v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetCasterId(v string) *UpdateCasterSceneAudioRequest {
	s.CasterId = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetFollowEnable(v int32) *UpdateCasterSceneAudioRequest {
	s.FollowEnable = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetMixList(v []*string) *UpdateCasterSceneAudioRequest {
	s.MixList = v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetOwnerId(v int64) *UpdateCasterSceneAudioRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetRegionId(v string) *UpdateCasterSceneAudioRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) SetSceneId(v string) *UpdateCasterSceneAudioRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateCasterSceneAudioRequest) Validate() error {
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

type UpdateCasterSceneAudioRequestAudioLayer struct {
	// Specifies whether to enable the features provided by the 3A audio algorithm. This parameter consists of the following fields:
	//
	// - **enableAgc**: (Optional) Specifies whether to enable the automatic gain control (AGC) feature of the 3A algorithm. Valid values: **0*	- (disabled, default) and **1*	- (enabled).
	//
	// - **enableAns**: (Optional) Specifies whether to enable the intelligent noise reduction feature of the 3A algorithm. Valid values: **0*	- (disabled, default) and **1*	- (enabled).
	//
	// - **ansMode**: (Optional) The mode of the intelligent noise reduction feature. This field is active only when **enableAns*	- is set to **1**. Valid values: **0*	- (speech noise reduction, default) and **1*	- (music noise reduction).
	//
	// > For better noise reduction, set ansMode to 1.
	//
	// - **enableBeautify**: (Optional) Specifies whether to enable voice beautification. Valid values: **0*	- (disabled, default) and **1*	- (enabled).
	//
	// - **voiceBeautifyMode**: (Optional) The voice beautification mode. This field is active only when **enableBeautify*	- is set to **1**. Valid values: **0*	- (magnetic male voice, default) and **1*	- (fresh female voice).
	//
	// example:
	//
	// {   "enableAgc":0,   "enableAns":1 }
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The fixed latency of the audio layer. This parameter is used to synchronize the audio with captions.
	//
	// Unit: milliseconds. Valid values: 0 to **5000**. Default value: **0**.
	//
	// example:
	//
	// 0
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The sound channels that are used for volume input. Valid values:
	//
	// - **leftChannel**: the left sound channel.
	//
	// - **rightChannel**: the right sound channel.
	//
	// - **all*	- (default): both sound channels.
	//
	// example:
	//
	// all
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The volume multiplier for the audio stream. Valid values: 0 to **10.0**. Default value: **1.0**.
	//
	// - **1.0**: The original volume is used.
	//
	// - A value less than **1*	- decreases the volume.
	//
	// - A value greater than **1*	- increases the volume.
	//
	// example:
	//
	// 1
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s UpdateCasterSceneAudioRequestAudioLayer) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneAudioRequestAudioLayer) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) GetFilter() *string {
	return s.Filter
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) GetValidChannel() *string {
	return s.ValidChannel
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetFilter(v string) *UpdateCasterSceneAudioRequestAudioLayer {
	s.Filter = &v
	return s
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetFixedDelayDuration(v int32) *UpdateCasterSceneAudioRequestAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetValidChannel(v string) *UpdateCasterSceneAudioRequestAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) SetVolumeRate(v float32) *UpdateCasterSceneAudioRequestAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *UpdateCasterSceneAudioRequestAudioLayer) Validate() error {
	return dara.Validate(s)
}
