// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterSceneAudioResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAudioLayers(v *DescribeCasterSceneAudioResponseBodyAudioLayers) *DescribeCasterSceneAudioResponseBody
	GetAudioLayers() *DescribeCasterSceneAudioResponseBodyAudioLayers
	SetCasterId(v string) *DescribeCasterSceneAudioResponseBody
	GetCasterId() *string
	SetFollowEnable(v int32) *DescribeCasterSceneAudioResponseBody
	GetFollowEnable() *int32
	SetMixList(v *DescribeCasterSceneAudioResponseBodyMixList) *DescribeCasterSceneAudioResponseBody
	GetMixList() *DescribeCasterSceneAudioResponseBodyMixList
	SetRequestId(v string) *DescribeCasterSceneAudioResponseBody
	GetRequestId() *string
}

type DescribeCasterSceneAudioResponseBody struct {
	// The configurations of the audio layers.
	AudioLayers *DescribeCasterSceneAudioResponseBodyAudioLayers `json:"AudioLayers,omitempty" xml:"AudioLayers,omitempty" type:"Struct"`
	// The ID of the production studio. You can specify the ID in a request to start a scene in the production studio.
	//
	// example:
	//
	// 97df6b7f-3490-47d2-ac50-88338765****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The audio mode. By default, the audio follows video (AFV) mode is used. Valid values:
	//
	// 	- **0**: the audio mixing mode
	//
	// 	- **1**: the AFV mode
	//
	// example:
	//
	// 1
	FollowEnable *int32                                       `json:"FollowEnable,omitempty" xml:"FollowEnable,omitempty"`
	MixList      *DescribeCasterSceneAudioResponseBodyMixList `json:"MixList,omitempty" xml:"MixList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 98745637-3490-47d2-ac50-883387567098
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCasterSceneAudioResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseBody) GetAudioLayers() *DescribeCasterSceneAudioResponseBodyAudioLayers {
	return s.AudioLayers
}

func (s *DescribeCasterSceneAudioResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterSceneAudioResponseBody) GetFollowEnable() *int32 {
	return s.FollowEnable
}

func (s *DescribeCasterSceneAudioResponseBody) GetMixList() *DescribeCasterSceneAudioResponseBodyMixList {
	return s.MixList
}

func (s *DescribeCasterSceneAudioResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterSceneAudioResponseBody) SetAudioLayers(v *DescribeCasterSceneAudioResponseBodyAudioLayers) *DescribeCasterSceneAudioResponseBody {
	s.AudioLayers = v
	return s
}

func (s *DescribeCasterSceneAudioResponseBody) SetCasterId(v string) *DescribeCasterSceneAudioResponseBody {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseBody) SetFollowEnable(v int32) *DescribeCasterSceneAudioResponseBody {
	s.FollowEnable = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseBody) SetMixList(v *DescribeCasterSceneAudioResponseBodyMixList) *DescribeCasterSceneAudioResponseBody {
	s.MixList = v
	return s
}

func (s *DescribeCasterSceneAudioResponseBody) SetRequestId(v string) *DescribeCasterSceneAudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseBody) Validate() error {
	if s.AudioLayers != nil {
		if err := s.AudioLayers.Validate(); err != nil {
			return err
		}
	}
	if s.MixList != nil {
		if err := s.MixList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterSceneAudioResponseBodyAudioLayers struct {
	AudioLayer []*DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer `json:"AudioLayer,omitempty" xml:"AudioLayer,omitempty" type:"Repeated"`
}

func (s DescribeCasterSceneAudioResponseBodyAudioLayers) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseBodyAudioLayers) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayers) GetAudioLayer() []*DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer {
	return s.AudioLayer
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayers) SetAudioLayer(v []*DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) *DescribeCasterSceneAudioResponseBodyAudioLayers {
	s.AudioLayer = v
	return s
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayers) Validate() error {
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

type DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer struct {
	// The fixed delay of the audio layer. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The sound channel type of the audio layer. Valid values:
	//
	// 	- **left**: the left channel
	//
	// 	- **right**: the right channel
	//
	// 	- **all*	- (default): both the left and right channels
	//
	// example:
	//
	// all
	ValidChannel *string `json:"ValidChannel,omitempty" xml:"ValidChannel,omitempty"`
	// The volume of the audio layer.
	//
	// example:
	//
	// 1
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) GetValidChannel() *string {
	return s.ValidChannel
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) SetFixedDelayDuration(v int32) *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer {
	s.FixedDelayDuration = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) SetValidChannel(v string) *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer {
	s.ValidChannel = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) SetVolumeRate(v float32) *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer {
	s.VolumeRate = &v
	return s
}

func (s *DescribeCasterSceneAudioResponseBodyAudioLayersAudioLayer) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterSceneAudioResponseBodyMixList struct {
	LocationId []*string `json:"LocationId,omitempty" xml:"LocationId,omitempty" type:"Repeated"`
}

func (s DescribeCasterSceneAudioResponseBodyMixList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterSceneAudioResponseBodyMixList) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioResponseBodyMixList) GetLocationId() []*string {
	return s.LocationId
}

func (s *DescribeCasterSceneAudioResponseBodyMixList) SetLocationId(v []*string) *DescribeCasterSceneAudioResponseBodyMixList {
	s.LocationId = v
	return s
}

func (s *DescribeCasterSceneAudioResponseBodyMixList) Validate() error {
	return dara.Validate(s)
}
