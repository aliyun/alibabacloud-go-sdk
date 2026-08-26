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
	// The start offset of the video file. Unit: milliseconds.
	//
	// 	Notice:
	//
	// This parameter is valid only if the video source is a video file.
	//
	//
	//
	// > A value greater than 0 specifies the start time to read the file. The time is an offset from the first frame.
	//
	// example:
	//
	// 0
	BeginOffset *int32 `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
	// The ID of the production studio.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId returned in the response.
	//
	// - If you create a production studio in the console, find the ID on the **Cloud Production Studio*	- page. To go to this page, choose **LIVE Console*	- > **Production Studio**.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is its ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// This parameter is valid only for video files. Unit: milliseconds.
	//
	// - If the value is greater than **0**, it specifies the end time to read the file. The time is an offset from the first frame.
	//
	// - If the value is less than **0**, it specifies the end time to read the file. The time is an offset from the last frame.
	//
	// example:
	//
	// 10000
	EndOffset *int32 `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	// The ID of the image material in the media asset library.
	//
	// > This parameter is required only if the video source is an image.
	//
	// example:
	//
	// a089175eb5f4427684fc0715159a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The URL of the image material.
	//
	// > This parameter is available only if the video source is an image that has not been imported to the material library. The image must be in JPG or PNG format, and its size cannot exceed 10 MB.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The URL of the live stream.
	//
	// 	Notice:
	//
	// This parameter is required only if the video source is a live stream.
	//
	// example:
	//
	// rtmp://guide.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****
	LiveStreamUrl *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	// The material ID.
	//
	// example:
	//
	// f080575eb5f4427684fc0715159a****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Presentation Time Stamp (PTS) callback interval. Unit: milliseconds. This parameter is valid only for VOD materials.
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
	// This parameter is valid only for video files. It specifies the number of times to loop the video after playback is complete.
	//
	// - **0*	- (default): The video does not loop.
	//
	// - **-1**: The video loops indefinitely.
	//
	// example:
	//
	// 0
	RepeatNum *int32 `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty"`
	// The resource ID. If you add a video source to the production studio by calling the [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html) operation, use the ResourceId returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the video source.
	//
	// example:
	//
	// test001
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The URL of the video on demand (VOD) file.
	//
	// 	Notice:
	//
	// This parameter is available only if the video source is a video file that has not been imported to the material library.
	//
	//
	//
	// > VOD files must be in MP4, FLV, or TS format.
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
