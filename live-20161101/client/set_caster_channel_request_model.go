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
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value returned in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The production studio name in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The channel ID.
	//
	// The reference ID for the layout scene. You can set a maximum of one resource for each channel. The total number of channels is determined when you create the production studio. The format is \\`RV01\\` to \\`RV12\\`.
	//
	// This parameter is required.
	//
	// example:
	//
	// RV01
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The facial retouching settings. Valid values: 0 (whole), 1 (skin smoothing), 2 (skin whitening), 3 (dark circle removal), and 4 (nasolabial fold removal).
	//
	// example:
	//
	// 0
	FaceBeauty *string `json:"FaceBeauty,omitempty" xml:"FaceBeauty,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The playback status. This parameter applies only to video files, not live streams. Valid values:
	//
	// - **1*	- (default): Playback.
	//
	// - **0**: Pause.
	//
	// example:
	//
	// 1
	PlayStatus *int32 `json:"PlayStatus,omitempty" xml:"PlayStatus,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the video source.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// This parameter applies only to video files, not live streams. The value must be greater than or equal to 0. It specifies the offset from the first frame at which to start reading the file. Unit: milliseconds (ms).
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
