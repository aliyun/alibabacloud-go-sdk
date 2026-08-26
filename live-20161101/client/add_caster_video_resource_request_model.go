// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterVideoResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginOffset(v int32) *AddCasterVideoResourceRequest
	GetBeginOffset() *int32
	SetCasterId(v string) *AddCasterVideoResourceRequest
	GetCasterId() *string
	SetEndOffset(v int32) *AddCasterVideoResourceRequest
	GetEndOffset() *int32
	SetFixedDelayDuration(v int32) *AddCasterVideoResourceRequest
	GetFixedDelayDuration() *int32
	SetImageId(v string) *AddCasterVideoResourceRequest
	GetImageId() *string
	SetImageUrl(v string) *AddCasterVideoResourceRequest
	GetImageUrl() *string
	SetLiveStreamUrl(v string) *AddCasterVideoResourceRequest
	GetLiveStreamUrl() *string
	SetLocationId(v string) *AddCasterVideoResourceRequest
	GetLocationId() *string
	SetMaterialId(v string) *AddCasterVideoResourceRequest
	GetMaterialId() *string
	SetOwnerId(v int64) *AddCasterVideoResourceRequest
	GetOwnerId() *int64
	SetPtsCallbackInterval(v int32) *AddCasterVideoResourceRequest
	GetPtsCallbackInterval() *int32
	SetRegionId(v string) *AddCasterVideoResourceRequest
	GetRegionId() *string
	SetRepeatNum(v int32) *AddCasterVideoResourceRequest
	GetRepeatNum() *int32
	SetResourceName(v string) *AddCasterVideoResourceRequest
	GetResourceName() *string
	SetVodUrl(v string) *AddCasterVideoResourceRequest
	GetVodUrl() *string
}

type AddCasterVideoResourceRequest struct {
	// The start offset of the video file. Unit: milliseconds.
	//
	// 	Notice: This parameter takes effect only when the video source type is file video.
	//
	//
	// > A value greater than **0*	- indicates that reading starts from the offset time relative to the first frame.
	//
	// example:
	//
	// 1000
	BeginOffset *int32 `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, navigate to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the production studio name.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page of the ApsaraVideo Live console is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The end offset of the video file. Unit: milliseconds.
	//
	// 	Notice: This parameter takes effect only when the video source type is file video.
	//
	//
	//
	// - A value greater than **0**: reading ends at the offset time relative to the first frame.
	//
	// - A value less than **0**: reading ends at the offset time relative to the last frame.
	//
	// example:
	//
	// 10000
	EndOffset *int32 `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	// The fixed delay for the video, which can be used for subtitle synchronization. Unit: ms. Default value: 0. Value range: `[0-5000]`.
	//
	// example:
	//
	// 0
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// The media asset library image material ID.
	//
	// > This parameter is available and required only when the video source type is image.
	//
	// example:
	//
	// a089175eb5f4427684fc0715159a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image material URL.
	//
	// >This parameter is available only when the video source type is image and the image file has not been imported to the media asset library. JPG and PNG formats are supported. The maximum file size is 10 MB.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ApsaraVideo Live streaming URL.
	//
	// 	Notice:
	//
	//
	//
	// -  This parameter is required when the video source type is live stream.
	//
	//
	//
	// -  Do not include this parameter in the request when the video source type is not live stream.
	//
	// example:
	//
	// rtmp://guide.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****
	LiveStreamUrl *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	// The location identifier of the video source. This parameter is required.
	//
	// Defines the reference number of a scene in the layout. Each location can be associated with at most one resource. The format must match "RV01~RV12", which is RV + a number in the range of `[01~99]`.
	//
	// example:
	//
	// RV01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The media asset library material ID.
	//
	// 	Notice: This parameter is available and required only when the video source type is material.
	//
	//
	// If you call the [DescribeCasterConfig](https://help.aliyun.com/document_detail/2848011.html) operation to query the production studio configuration, check the UrgentMaterialId parameter value returned by the DescribeCasterConfig operation.
	//
	// > The UrgentMaterialId value is the media asset library material ID.
	//
	// example:
	//
	// f080575eb5f4427684fc0715159a****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The PTS callback interval. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	PtsCallbackInterval *int32 `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of times the video repeats after playback completes. Valid values:
	//
	// 	Notice: This parameter takes effect only when the video source type is file video.
	//
	//
	// - **0*	- (default): no repeat.
	//
	// - **-1**: loops indefinitely.
	//
	// example:
	//
	// 0
	RepeatNum *int32 `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty"`
	// The name of the video source.
	//
	// This parameter is required.
	//
	// example:
	//
	// test001
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The video-on-demand file URL.
	//
	// 	Notice: This parameter is available only when the video source type is file video and the video file has not been imported to the media asset library.
	//
	//
	// >Video-on-demand files are limited to MP4, FLV, and TS formats.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/StreamName.flv
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s AddCasterVideoResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *AddCasterVideoResourceRequest) GetBeginOffset() *int32 {
	return s.BeginOffset
}

func (s *AddCasterVideoResourceRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddCasterVideoResourceRequest) GetEndOffset() *int32 {
	return s.EndOffset
}

func (s *AddCasterVideoResourceRequest) GetFixedDelayDuration() *int32 {
	return s.FixedDelayDuration
}

func (s *AddCasterVideoResourceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *AddCasterVideoResourceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *AddCasterVideoResourceRequest) GetLiveStreamUrl() *string {
	return s.LiveStreamUrl
}

func (s *AddCasterVideoResourceRequest) GetLocationId() *string {
	return s.LocationId
}

func (s *AddCasterVideoResourceRequest) GetMaterialId() *string {
	return s.MaterialId
}

func (s *AddCasterVideoResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCasterVideoResourceRequest) GetPtsCallbackInterval() *int32 {
	return s.PtsCallbackInterval
}

func (s *AddCasterVideoResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCasterVideoResourceRequest) GetRepeatNum() *int32 {
	return s.RepeatNum
}

func (s *AddCasterVideoResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *AddCasterVideoResourceRequest) GetVodUrl() *string {
	return s.VodUrl
}

func (s *AddCasterVideoResourceRequest) SetBeginOffset(v int32) *AddCasterVideoResourceRequest {
	s.BeginOffset = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetCasterId(v string) *AddCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetEndOffset(v int32) *AddCasterVideoResourceRequest {
	s.EndOffset = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetFixedDelayDuration(v int32) *AddCasterVideoResourceRequest {
	s.FixedDelayDuration = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetImageId(v string) *AddCasterVideoResourceRequest {
	s.ImageId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetImageUrl(v string) *AddCasterVideoResourceRequest {
	s.ImageUrl = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetLiveStreamUrl(v string) *AddCasterVideoResourceRequest {
	s.LiveStreamUrl = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetLocationId(v string) *AddCasterVideoResourceRequest {
	s.LocationId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetMaterialId(v string) *AddCasterVideoResourceRequest {
	s.MaterialId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetOwnerId(v int64) *AddCasterVideoResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetPtsCallbackInterval(v int32) *AddCasterVideoResourceRequest {
	s.PtsCallbackInterval = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetRegionId(v string) *AddCasterVideoResourceRequest {
	s.RegionId = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetRepeatNum(v int32) *AddCasterVideoResourceRequest {
	s.RepeatNum = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetResourceName(v string) *AddCasterVideoResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *AddCasterVideoResourceRequest) SetVodUrl(v string) *AddCasterVideoResourceRequest {
	s.VodUrl = &v
	return s
}

func (s *AddCasterVideoResourceRequest) Validate() error {
	return dara.Validate(s)
}
