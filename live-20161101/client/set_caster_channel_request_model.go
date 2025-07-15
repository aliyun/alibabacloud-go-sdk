// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *SetCasterChannelRequest
	GetCasterId() *string
	SetChannelId(v string) *SetCasterChannelRequest
	GetChannelId() *string
	SetFaceBeauty(v string) *SetCasterChannelRequest
	GetFaceBeauty() *string
	SetOwnerId(v int64) *SetCasterChannelRequest
	GetOwnerId() *int64
	SetPlayStatus(v int32) *SetCasterChannelRequest
	GetPlayStatus() *int32
	SetRegionId(v string) *SetCasterChannelRequest
	GetRegionId() *string
	SetResourceId(v string) *SetCasterChannelRequest
	GetResourceId() *string
	SetSeekOffset(v int32) *SetCasterChannelRequest
	GetSeekOffset() *int32
}

type SetCasterChannelRequest struct {
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
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the channel.
	//
	// When channels are enabled, the layout references the channel IDs. You can specify up to one resource for a channel. The number of resources is limited by the number of the channels of the production studio. The value must be in the RV[Number] format, such as RV01 and RV12.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV01
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The face retouching effect. Valid values: 0 (all effects), 1 (skin smoothing), 2 (skin whitening), 3 (dark circles removal), and 4 (nasolabial folds removal).
	//
	// example:
	//
	// 0
	FaceBeauty *string `json:"FaceBeauty,omitempty" xml:"FaceBeauty,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The playback status. This parameter take effects for video files, but not for live streams. Valid values:
	//
	// 	- **1**: specifies that the video source is playing. This is the default value.
	//
	// 	- **0**: specifies that the playback of the video source is paused.
	//
	// example:
	//
	// 1
	PlayStatus *int32  `json:"PlayStatus,omitempty" xml:"PlayStatus,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the video source.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The offset of the position where the production studio starts reading the video source. The value must be greater than or equal to 0, indicating an offset from the first frame. This parameter take effects for video files, but not for live streams. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	SeekOffset *int32 `json:"SeekOffset,omitempty" xml:"SeekOffset,omitempty"`
}

func (s SetCasterChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCasterChannelRequest) GoString() string {
	return s.String()
}

func (s *SetCasterChannelRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *SetCasterChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *SetCasterChannelRequest) GetFaceBeauty() *string {
	return s.FaceBeauty
}

func (s *SetCasterChannelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCasterChannelRequest) GetPlayStatus() *int32 {
	return s.PlayStatus
}

func (s *SetCasterChannelRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetCasterChannelRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *SetCasterChannelRequest) GetSeekOffset() *int32 {
	return s.SeekOffset
}

func (s *SetCasterChannelRequest) SetCasterId(v string) *SetCasterChannelRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterChannelRequest) SetChannelId(v string) *SetCasterChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *SetCasterChannelRequest) SetFaceBeauty(v string) *SetCasterChannelRequest {
	s.FaceBeauty = &v
	return s
}

func (s *SetCasterChannelRequest) SetOwnerId(v int64) *SetCasterChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCasterChannelRequest) SetPlayStatus(v int32) *SetCasterChannelRequest {
	s.PlayStatus = &v
	return s
}

func (s *SetCasterChannelRequest) SetRegionId(v string) *SetCasterChannelRequest {
	s.RegionId = &v
	return s
}

func (s *SetCasterChannelRequest) SetResourceId(v string) *SetCasterChannelRequest {
	s.ResourceId = &v
	return s
}

func (s *SetCasterChannelRequest) SetSeekOffset(v int32) *SetCasterChannelRequest {
	s.SeekOffset = &v
	return s
}

func (s *SetCasterChannelRequest) Validate() error {
	return dara.Validate(s)
}
