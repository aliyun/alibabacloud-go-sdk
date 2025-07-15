// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterVideoResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCasterVideoResourcesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCasterVideoResourcesResponseBody
	GetTotal() *int32
	SetVideoResources(v *DescribeCasterVideoResourcesResponseBodyVideoResources) *DescribeCasterVideoResourcesResponseBody
	GetVideoResources() *DescribeCasterVideoResourcesResponseBodyVideoResources
}

type DescribeCasterVideoResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF60DB6A-7FD6-426E-9288-122CC1A52FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The input sources.
	VideoResources *DescribeCasterVideoResourcesResponseBodyVideoResources `json:"VideoResources,omitempty" xml:"VideoResources,omitempty" type:"Struct"`
}

func (s DescribeCasterVideoResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterVideoResourcesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterVideoResourcesResponseBody) GetVideoResources() *DescribeCasterVideoResourcesResponseBodyVideoResources {
	return s.VideoResources
}

func (s *DescribeCasterVideoResourcesResponseBody) SetRequestId(v string) *DescribeCasterVideoResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBody) SetTotal(v int32) *DescribeCasterVideoResourcesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBody) SetVideoResources(v *DescribeCasterVideoResourcesResponseBodyVideoResources) *DescribeCasterVideoResourcesResponseBody {
	s.VideoResources = v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterVideoResourcesResponseBodyVideoResources struct {
	VideoResource []*DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource `json:"VideoResource,omitempty" xml:"VideoResource,omitempty" type:"Repeated"`
}

func (s DescribeCasterVideoResourcesResponseBodyVideoResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponseBodyVideoResources) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResources) GetVideoResource() []*DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	return s.VideoResource
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResources) SetVideoResource(v []*DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) *DescribeCasterVideoResourcesResponseBodyVideoResources {
	s.VideoResource = v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResources) Validate() error {
	return dara.Validate(s)
}

type DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource struct {
	// The offset of the position where the system starts to read the video resource. This parameter takes effect only if the input source is a video file. Unit: milliseconds.
	//
	// A value **greater than 0*	- indicates an offset from the first frame.
	//
	// example:
	//
	// 1000
	BeginOffset *int32 `json:"BeginOffset,omitempty" xml:"BeginOffset,omitempty"`
	// The offset of the position where the system stops reading the video file. This parameter takes effect only if the input source is a video file. Unit: milliseconds.
	//
	// 	- A value greater than **0*	- indicates an offset from the first frame.
	//
	// 	- A value smaller than **0*	- indicates an offset from the last frame.
	//
	// example:
	//
	// 10000
	EndOffset *int32 `json:"EndOffset,omitempty" xml:"EndOffset,omitempty"`
	// The source URL.
	//
	// example:
	//
	// http://guide.aliyundoc.com/example.org
	FlvUrl *string `json:"FlvUrl,omitempty" xml:"FlvUrl,omitempty"`
	// The image ID.
	//
	// example:
	//
	// a089175eb5f4427684fc0715159a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image URL.
	//
	// example:
	//
	// http://learn.aliyundoc.com/AppName/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The URL of the live stream.
	//
	// example:
	//
	// rtmp://abcLive/appName/b5447c21fcfe444c9e9b6f7ba208****
	LiveStreamUrl *string `json:"LiveStreamUrl,omitempty" xml:"LiveStreamUrl,omitempty"`
	// The position of the video resource.
	//
	// example:
	//
	// RV01
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	// The material ID.
	//
	// example:
	//
	// d2c429cd907742ee8f6e76465ad3****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	// The interval between presentation timestamp (PTS) callbacks. If you set the value to 0, the PTS callback is disabled. This parameter is returned only when the video resource is a video-on-demand (VOD) file.
	//
	// example:
	//
	// 0
	PtsCallbackInterval *int32 `json:"PtsCallbackInterval,omitempty" xml:"PtsCallbackInterval,omitempty"`
	// The number of playback times after the first playback is complete. This parameter takes effect only when the input source is a video file. Valid values:
	//
	// 	- **0*	- (default): The video file is played only once.
	//
	// 	- **-1**: The video file is played in loop mode.
	//
	// example:
	//
	// 0
	RepeatNum *int32 `json:"RepeatNum,omitempty" xml:"RepeatNum,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// resource-Name1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The URL of the VOD file.
	//
	// This parameter is returned only when the video resource is an MP4, FLV, or TS file that is not from the media library.
	//
	// example:
	//
	// http://developer.aliyundoc.com/caster1.flv
	VodUrl *string `json:"VodUrl,omitempty" xml:"VodUrl,omitempty"`
}

func (s DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetBeginOffset() *int32 {
	return s.BeginOffset
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetEndOffset() *int32 {
	return s.EndOffset
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetFlvUrl() *string {
	return s.FlvUrl
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetLiveStreamUrl() *string {
	return s.LiveStreamUrl
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetLocationId() *string {
	return s.LocationId
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetMaterialId() *string {
	return s.MaterialId
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetPtsCallbackInterval() *int32 {
	return s.PtsCallbackInterval
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetRepeatNum() *int32 {
	return s.RepeatNum
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) GetVodUrl() *string {
	return s.VodUrl
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetBeginOffset(v int32) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.BeginOffset = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetEndOffset(v int32) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.EndOffset = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetFlvUrl(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.FlvUrl = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetImageId(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.ImageId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetImageUrl(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.ImageUrl = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetLiveStreamUrl(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.LiveStreamUrl = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetLocationId(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.LocationId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetMaterialId(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.MaterialId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetPtsCallbackInterval(v int32) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.PtsCallbackInterval = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetRepeatNum(v int32) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.RepeatNum = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetResourceId(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetResourceName(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.ResourceName = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) SetVodUrl(v string) *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource {
	s.VodUrl = &v
	return s
}

func (s *DescribeCasterVideoResourcesResponseBodyVideoResourcesVideoResource) Validate() error {
	return dara.Validate(s)
}
