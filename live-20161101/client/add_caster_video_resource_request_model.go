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
	// The offset of the position where the system starts to read the video source. Unit: milliseconds.
	//
	// **
	//
	// **Important*	- This parameter takes effect only if the video source is a file.
	//
	// > A value greater than **0*	- specifies an offset from the first frame.
	//
	// example:
	//
	// 1000
	BeginOffset *int32 `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// > You can find the ID of the production studio in the Instance Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The offset of the position where the system stops reading the video source. Unit: milliseconds.
	//
	// **
	//
	// **Important*	- This parameter takes effect only if the video source is a file.
	//
	// 	- A value greater than **0*	- specifies an offset from the first frame.
	//
	// 	- A value less than **0*	- specifies an offset from the last frame.
	//
	// example:
	//
	// 10000
	EndOffset *int32 `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	// The fixed delay of the video layer. This parameter is used to synchronize the video with subtitles. Unit: milliseconds. Default value: 0. Valid values: `0 to 5000`.
	//
	// example:
	//
	// 0
	FixedDelayDuration *int32 `json:"FixedDelayDuration,omitempty" xml:"FixedDelayDuration,omitempty"`
	// ID of the media library image material.
	//
	// >This parameter is only available and must be provided when the video source type is an image.
	//
	// example:
	//
	// a089175eb5f4427684fc0715159a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Image material URL.
	//
	// >This parameter is available only when the video source type is an image and the image file has not been imported into the material library. Supports uploading images in jpg, png formats, with a maximum file size of 10MB.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The streaming URL.
	//
	// **
	//
	// **Important*	- This parameter is required if the video source is a live stream.
	//
	// > Do not specify this parameter in the request if the video source is not a live stream.
	//
	// example:
	//
	// rtmp://guide.aliyundoc.com/caster/4a82a3d1b7f0462ea37348366201****?auth_key=1608953344-0-0-ac8c628078541d7055a170ec59a5****
	LiveStreamUrl *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	// The ID that is used to identify the position of the video source.
	//
	// Define the reference numbers in the layout. Each reference number is associated with only one resource. The value of this parameter must be in the RV[Number] format, where Number is `01 to 99`.
	//
	// example:
	//
	// RV01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The ID of the material from the media library.
	//
	// **
	//
	// **Important*	- This parameter takes effect and is required only if the video source is a material.
	//
	// If you query the configurations of the production studio by calling the [DescribeCasterConfig](https://help.aliyun.com/document_detail/60259.html) operation, obtain the value of the response parameter UrgentMaterialId.
	//
	// > The value of the UrgentMaterialId parameter is the ID of the material from the media library.
	//
	// example:
	//
	// f080575eb5f4427684fc0715159a****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval between presentation timestamp (PTS) callbacks. Unit: milliseconds.
	//
	// example:
	//
	// 2000
	PtsCallbackInterval *int32  `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of playbacks after the first playback is complete. Valid values:
	//
	// **
	//
	// **Important*	- This parameter takes effect only if the video source is a file.
	//
	// 	- **0**: specifies that the video source is played only once. This is the default value.
	//
	// 	- **-1**: specifies that the video source is played in loop mode.
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
	// The URL of the VOD file.
	//
	// **
	//
	// **Important*	- This parameter takes effect only if the video source is a file that is not from the media library.
	//
	// > The VOD file must be in the MP4, FLV, or TS format.
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
