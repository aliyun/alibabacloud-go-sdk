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
	// The audio mode. By default, the AFV mode is used. If you do not specify this parameter, the scene retains the last configuration. Valid values:
	//
	// 	- **0**: the audio mixing mode.
	//
	// 	- **1**: the AFV mode.
	//
	// example:
	//
	// 1
	FollowEnable *int32 `json:"FollowEnable,omitempty" xml:"FollowEnable,omitempty"`
	// The location IDs of the audio layers, which are in the same order as the audio layers.
	//
	// example:
	//
	// RV01
	MixList  []*string `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene. If you call the [DescribeCasterScenes](https://help.aliyun.com/document_detail/2848039.html) operation to query scenes of the production studio, check the value of the response parameter ComponentId to obtain the ID.
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
	return dara.Validate(s)
}

type UpdateCasterSceneAudioRequestAudioLayer struct {
	// Specifies whether to enable the features provided by the audio 3A algorithms. This parameter consists of the following fields:
	//
	// 	- **enableAgc**: specifies whether to enable automatic gain control (AGC). This field is optional. Valid values: 0 and 1. **0*	- is the default value, which specifies that AGC is disabled. **1*	- specifies that AGC is enabled.
	//
	// 	- **enableAns**: specifies whether to enable active noise suppression (ANS). This field is optional. Valid values: 0 and 1. **0*	- is the default value, which specifies that ANS is disabled. **1*	- specifies that ANS is enabled.
	//
	// 	- **ansMode**: specifies the mode for ANS. This field is optional and takes effect only if you set **enableAns*	- to **1**. Valid values: 0 and 1. **0*	- is the default value, which specifies the speech noise reduction mode. **1*	- specifies the music noise reduction mode.
	//
	// >  To ensure a better noise reduction effect, we recommend that you set ansMode to 1.
	//
	// 	- **enableBeautify**: specifies whether to enable voice change. This field is optional. Valid values: 0 and 1. **0*	- is the default value, which specifies that voice change is disabled. **1*	- specifies that voice change is enabled.
	//
	// 	- **voiceBeautifyMode**: specifies the mode for voice change. This field is optional and takes effect only if you set **enableBeautify*	- to **1**. Valid values: 0 and 1. **0*	- is the default value, which specifies the magnetic male voice mode. **1*	- specifies the fresh female voice mode.
	//
	// example:
	//
	// {   "enableAgc":0,   "enableAns":1 }
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The fixed delay of the audio layer. This parameter is used to synchronize the audio with subtitles.
	//
	// Unit: milliseconds. Valid values: **0 to 5000**. Default value: **0**.
	//
	// example:
	//
	// 0
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
	// The multiple of the original volume at which the audio layer plays audio. Valid values: **0 to 10.0**. Default value: **1.0**.
	//
	// 	- **1.0**: specifies that the audio layer plays audio at the original volume.
	//
	// 	- A value smaller than **1**: specifies that the audio layer plays audio at a volume that is less than the original volume.
	//
	// 	- A value greater than **1**: specifies that the audio layer plays audio at a volume that is more than the original volume.
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
