// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterVideoResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginOffset(v int32) *ModifyCasterVideoResourceRequest
	GetBeginOffset() *int32
	SetCasterId(v string) *ModifyCasterVideoResourceRequest
	GetCasterId() *string
	SetEndOffset(v int32) *ModifyCasterVideoResourceRequest
	GetEndOffset() *int32
	SetImageId(v string) *ModifyCasterVideoResourceRequest
	GetImageId() *string
	SetImageUrl(v string) *ModifyCasterVideoResourceRequest
	GetImageUrl() *string
	SetLiveStreamUrl(v string) *ModifyCasterVideoResourceRequest
	GetLiveStreamUrl() *string
	SetMaterialId(v string) *ModifyCasterVideoResourceRequest
	GetMaterialId() *string
	SetOwnerId(v int64) *ModifyCasterVideoResourceRequest
	GetOwnerId() *int64
	SetPtsCallbackInterval(v int32) *ModifyCasterVideoResourceRequest
	GetPtsCallbackInterval() *int32
	SetRegionId(v string) *ModifyCasterVideoResourceRequest
	GetRegionId() *string
	SetRepeatNum(v int32) *ModifyCasterVideoResourceRequest
	GetRepeatNum() *int32
	SetResourceId(v string) *ModifyCasterVideoResourceRequest
	GetResourceId() *string
	SetResourceName(v string) *ModifyCasterVideoResourceRequest
	GetResourceName() *string
	SetVodUrl(v string) *ModifyCasterVideoResourceRequest
	GetVodUrl() *string
}

type ModifyCasterVideoResourceRequest struct {
	// The offset of the position where the system starts to read the video resource.
	//
	// This parameter takes effect only when the video resource is a video file. Unit: milliseconds.
	//
	// >  A value greater than 0 indicates an offset from the first frame.
	//
	// example:
	//
	// 0
	BeginOffset *int32 `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
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
	// This parameter takes effect only when the video resource is a video file. Unit: milliseconds.
	//
	// 	- A value greater than **0*	- indicates an offset from the first frame.
	//
	// 	- A value smaller than **0*	- indicates an offset from the last frame.
	//
	// example:
	//
	// 10000
	EndOffset *int32 `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	// ID of the media library image material.
	//
	// > This parameter is only available and must be provided when the video source type is an image.
	//
	// example:
	//
	// a089175eb5f4427684fc0715159a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Image material URL.
	//
	// >This parameter is only available when the video source type is an image and the image file has not been imported into the material library. Supports uploading images in jpg, png formats, with a maximum file size of 10MB.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The URL of the live stream.
	//
	// This parameter takes effect and is required only when the video resource is a live stream.
	//
	// example:
	//
	// rtmp://guide.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****
	LiveStreamUrl *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	// The ID of the material.
	//
	// example:
	//
	// f080575eb5f4427684fc0715159a****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval between presentation timestamp (PTS) callbacks. Unit: milliseconds. This parameter takes effect only when the video resource is a VOD file.
	//
	// example:
	//
	// 2000
	PtsCallbackInterval *int32  `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of playback times after the first playback is complete. This parameter takes effect only when the video resource is a file. Valid values:
	//
	// 	- **0**: indicates that the video is played only once. This is the default value.
	//
	// 	- **-1**: indicates that the video is played in loop mode.
	//
	// example:
	//
	// 0
	RepeatNum *int32 `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty"`
	// The ID of the video resource. It is reflected in the ResourceId parameter when you call the [AddCasterVideoResource](~~60250#doc-api-live-AddCasterVideoResource~~ "Adds a video resource to a production studio.") operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the video resource.
	//
	// example:
	//
	// test001
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The URL of the video-on-demand (VOD) file. This parameter takes effect only when the video resource is a video file that is not from the media library.
	//
	// >  The VOD file must be in the MP4, FLV, or TS format.
	//
	// example:
	//
	// http://developer.aliyundoc.com/AppName/StreamName.flv
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s ModifyCasterVideoResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterVideoResourceRequest) GetBeginOffset() *int32 {
	return s.BeginOffset
}

func (s *ModifyCasterVideoResourceRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterVideoResourceRequest) GetEndOffset() *int32 {
	return s.EndOffset
}

func (s *ModifyCasterVideoResourceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyCasterVideoResourceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ModifyCasterVideoResourceRequest) GetLiveStreamUrl() *string {
	return s.LiveStreamUrl
}

func (s *ModifyCasterVideoResourceRequest) GetMaterialId() *string {
	return s.MaterialId
}

func (s *ModifyCasterVideoResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCasterVideoResourceRequest) GetPtsCallbackInterval() *int32 {
	return s.PtsCallbackInterval
}

func (s *ModifyCasterVideoResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCasterVideoResourceRequest) GetRepeatNum() *int32 {
	return s.RepeatNum
}

func (s *ModifyCasterVideoResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ModifyCasterVideoResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ModifyCasterVideoResourceRequest) GetVodUrl() *string {
	return s.VodUrl
}

func (s *ModifyCasterVideoResourceRequest) SetBeginOffset(v int32) *ModifyCasterVideoResourceRequest {
	s.BeginOffset = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetCasterId(v string) *ModifyCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetEndOffset(v int32) *ModifyCasterVideoResourceRequest {
	s.EndOffset = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetImageId(v string) *ModifyCasterVideoResourceRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetImageUrl(v string) *ModifyCasterVideoResourceRequest {
	s.ImageUrl = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetLiveStreamUrl(v string) *ModifyCasterVideoResourceRequest {
	s.LiveStreamUrl = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetMaterialId(v string) *ModifyCasterVideoResourceRequest {
	s.MaterialId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetOwnerId(v int64) *ModifyCasterVideoResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetPtsCallbackInterval(v int32) *ModifyCasterVideoResourceRequest {
	s.PtsCallbackInterval = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetRegionId(v string) *ModifyCasterVideoResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetRepeatNum(v int32) *ModifyCasterVideoResourceRequest {
	s.RepeatNum = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetResourceId(v string) *ModifyCasterVideoResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetResourceName(v string) *ModifyCasterVideoResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) SetVodUrl(v string) *ModifyCasterVideoResourceRequest {
	s.VodUrl = &v
	return s
}

func (s *ModifyCasterVideoResourceRequest) Validate() error {
	return dara.Validate(s)
}
