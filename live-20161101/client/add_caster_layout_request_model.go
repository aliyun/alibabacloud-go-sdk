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
	// Audio layout.
	//
	// This parameter is required.
	AudioLayer []*AddCasterLayoutRequestAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
	// The element represents the location ID of the video resource, i.e., LocationId. Refer to [Adding Video Source](https://help.aliyun.com/document_detail/60250.html) for LocationId, which corresponds in order with the VideoLayers elements.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV01
	BlendList []*string `json:"BlendList,omitempty" xml:"BlendList,omitempty" type:"Repeated"`
	// The ID of the production studio.
	//
	// If you create a production studio through the [CreateCaster](~~69338#doc-api-live-CreateCaster~~ "Creates a production studio.") interface, check the value of the CasterId parameter in the response.
	//
	// If you create a production studio through the ApsaraVideo Live Console, log in to the console, then check the ID of the production studio through the following path:
	//
	// Production Studios > Production Studio Management
	//
	// >  The CasterId is reflected in the Name column on the Production Studio Management page.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The element represents the location ID of the audio resource, i.e., LocationId.
	//
	// LocationId is referred to in [Adding Video Source](https://help.aliyun.com/document_detail/60250.html), and corresponds in order with the AudioLayers elements.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV01
	MixList  []*string `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Video layout.
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
	return dara.Validate(s)
}

type AddCasterLayoutRequestAudioLayer struct {
	// The fixed delay of audio layer N. You can use this parameter to synchronize the audio with subtitles. Unit: milliseconds. Valid values: **0 to 5000**. Default value: **0**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The valid voice channels of audio layer N. Valid values:
	//
	// 	- **leftChannel**: the left channel.
	//
	// 	- **rightChannel**: the right channel.
	//
	// 	- **all**: both the left and right channels. This is the default value.
	//
	// example:
	//
	// all
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The multiples of the original volume at which audio layer N plays audio streams. Valid values: **0 to 10.0**.
	//
	// 	- The default value **1.0*	- indicates that audio layer N plays audio streams at the original volume.
	//
	// 	- A value smaller than **1.0*	- indicates that audio layer N plays audio streams at a lower volume than the original one.
	//
	// 	- A value greater than **1.0*	- indicates that audio layer N plays audio streams at a higher volume than the original one.
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
	// The scaling mode of video layer N. Valid values:
	//
	// 	- **none**: The image is not scaled to fill in the specified layout section. Set video layer N based on the image size of the video resource. This is the default value.
	//
	// 	- **fit**: The image is scaled with the original aspect ratio to fill in the specified layout section. Set video layer N based on the section size. The image is centered in the layout section with the long side of the image equaling that of the section. If the aspect ratio of the image is inconsistent with that of the section, the short side of the image may be shorter than that of the section. The area outside the image displays the next video layer or the background if no next video layer exists. By default, the background color is black.
	//
	// example:
	//
	// fit
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The fixed delay of video layer N. You can use this parameter to synchronize the video with subtitles. Unit: milliseconds. Valid values: **0 to 5000**. Default value: **0**.
	//
	// example:
	//
	// 5000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The normalized value of the height of the image of video layer N.
	//
	// 	- If the FillMode parameter of video layer N is set to none, the width of the video image is scaled based on this parameter. The default value is **0**, which indicates that the video image is displayed in the original size.
	//
	// 	- If the FillMode parameter of video layer N is set to fit, you must set this parameter to a value greater than **0**. In this case, the video image is scaled with the original aspect ratio to fill in the specified layout section based on this parameter.
	//
	// example:
	//
	// 1
	HeightNormalized *float32 `json:"HeightNormalized,omitempty" xml:"HeightNormalized,omitempty"`
	// The normalized value of the `[x,y]` coordinates of video layer N in the production studio. The default coordinates are `[0,0]`.
	//
	// >  The coordinates indicate the location of video layer N in the production studio. Set this parameter to the normalized value of the coordinates.
	//
	// example:
	//
	// 0.3
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	// The reference coordinates of video layer N in the production studio. Valid values:
	//
	// 	- **topLeft**: the upper-left corner. This is the default value.
	//
	// 	- **topRight**: the upper-right corner.
	//
	// 	- **bottomLeft**: the lower-left corner.
	//
	// 	- **bottomRight**: the lower-right corner.
	//
	// 	- **center**: the center position.
	//
	// 	- **topCenter**: the upper center position.
	//
	// 	- **bottomCenter**: the lower center position.
	//
	// 	- **leftCenter**: the left center position.
	//
	// 	- **rightCenter**: the right center position.
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The normalized value of the width of the image of video layer N.
	//
	// 	- If the FillMode parameter of video layer N is set to none, the height of the video image is scaled based on this parameter. The default value is **0**, which indicates that the video image is displayed in the original size.
	//
	// 	- If the FillMode parameter of video layer N is set to fit, you must set this parameter to a value greater than **0**. In this case, the video image is scaled with the original aspect ratio to fill in the specified layout section based on this parameter.
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
