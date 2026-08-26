// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShowIntoShowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *AddShowIntoShowListRequest
	GetCasterId() *string
	SetDuration(v int64) *AddShowIntoShowListRequest
	GetDuration() *int64
	SetLiveInputType(v int32) *AddShowIntoShowListRequest
	GetLiveInputType() *int32
	SetOwnerId(v int64) *AddShowIntoShowListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddShowIntoShowListRequest
	GetRegionId() *string
	SetRepeatTimes(v int32) *AddShowIntoShowListRequest
	GetRepeatTimes() *int32
	SetResourceId(v string) *AddShowIntoShowListRequest
	GetResourceId() *string
	SetResourceType(v string) *AddShowIntoShowListRequest
	GetResourceType() *string
	SetResourceUrl(v string) *AddShowIntoShowListRequest
	GetResourceUrl() *string
	SetShowName(v string) *AddShowIntoShowListRequest
	GetShowName() *string
	SetSpot(v int32) *AddShowIntoShowListRequest
	GetSpot() *int32
	SetIsBatchMode(v bool) *AddShowIntoShowListRequest
	GetIsBatchMode() *bool
	SetShowList(v []*AddShowIntoShowListRequestShowList) *AddShowIntoShowListRequest
	GetShowList() []*AddShowIntoShowListRequestShowList
}

type AddShowIntoShowListRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId value returned by the CreateCaster operation.
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
	// The playback duration of a single show. Unit: seconds.
	//
	// > - You can set only one of **RepeatTimes*	- and **Duration**.
	//
	// > - This parameter is required when ResourceType is set to live.
	//
	// example:
	//
	// 20
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The custom type tag.
	//
	// example:
	//
	// 1
	LiveInputType *int32 `json:"LiveInputType,omitempty" xml:"LiveInputType,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of times a single show is repeated. Default value: 0.
	//
	// > - You can set only one of **RepeatTimes*	- and **Duration**.- RepeatTimes specifies the number of repetitions. For example, -1 means infinite repetition, 0 means the show is repeated 0 times (played once), 1 means the show is repeated 1 time (played twice), and so on.
	//
	// example:
	//
	// 0
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// The VOD file ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - live: live streaming resource.
	//
	// - vod: video-on-demand resource.
	//
	// - pic: image resource.
	//
	// >- When using video-on-demand (VOD) resources, use managed Bucket resources first. Resources in your own Bucket may expire. If you use resources in your own Bucket, check the resource validity period.
	//
	// - Live files support live streaming resources and third-party URLs.
	//
	// - VOD files support video-on-demand resources, image resources, and third-party URLs.
	//
	// - When using live streaming resources, you must also specify the Duration parameter.
	//
	// example:
	//
	// vod
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The URL of the playback file.
	ResourceUrl *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
	// The show name.
	//
	// example:
	//
	// liveShow****
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The position in the playlist where the show is inserted. Positions start from 0. By default, the show is added to the end of the current playlist.
	//
	// example:
	//
	// 1
	Spot *int32 `json:"Spot,omitempty" xml:"Spot,omitempty"`
	// Specifies whether to add shows to the playlist in batch. Valid values:
	//
	// - true: Batch addition.
	//
	// - false: Single addition.
	//
	// >If this parameter is not specified or left empty, single addition is used.
	//
	// example:
	//
	// false
	IsBatchMode *bool `json:"isBatchMode,omitempty" xml:"isBatchMode,omitempty"`
	// The list of show resources to add. Each resource has independent parameters such as showName and resourceUrl.
	ShowList []*AddShowIntoShowListRequestShowList `json:"showList,omitempty" xml:"showList,omitempty" type:"Repeated"`
}

func (s AddShowIntoShowListRequest) String() string {
	return dara.Prettify(s)
}

func (s AddShowIntoShowListRequest) GoString() string {
	return s.String()
}

func (s *AddShowIntoShowListRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddShowIntoShowListRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *AddShowIntoShowListRequest) GetLiveInputType() *int32 {
	return s.LiveInputType
}

func (s *AddShowIntoShowListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddShowIntoShowListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddShowIntoShowListRequest) GetRepeatTimes() *int32 {
	return s.RepeatTimes
}

func (s *AddShowIntoShowListRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddShowIntoShowListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddShowIntoShowListRequest) GetResourceUrl() *string {
	return s.ResourceUrl
}

func (s *AddShowIntoShowListRequest) GetShowName() *string {
	return s.ShowName
}

func (s *AddShowIntoShowListRequest) GetSpot() *int32 {
	return s.Spot
}

func (s *AddShowIntoShowListRequest) GetIsBatchMode() *bool {
	return s.IsBatchMode
}

func (s *AddShowIntoShowListRequest) GetShowList() []*AddShowIntoShowListRequestShowList {
	return s.ShowList
}

func (s *AddShowIntoShowListRequest) SetCasterId(v string) *AddShowIntoShowListRequest {
	s.CasterId = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetDuration(v int64) *AddShowIntoShowListRequest {
	s.Duration = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetLiveInputType(v int32) *AddShowIntoShowListRequest {
	s.LiveInputType = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetOwnerId(v int64) *AddShowIntoShowListRequest {
	s.OwnerId = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetRegionId(v string) *AddShowIntoShowListRequest {
	s.RegionId = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetRepeatTimes(v int32) *AddShowIntoShowListRequest {
	s.RepeatTimes = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetResourceId(v string) *AddShowIntoShowListRequest {
	s.ResourceId = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetResourceType(v string) *AddShowIntoShowListRequest {
	s.ResourceType = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetResourceUrl(v string) *AddShowIntoShowListRequest {
	s.ResourceUrl = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetShowName(v string) *AddShowIntoShowListRequest {
	s.ShowName = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetSpot(v int32) *AddShowIntoShowListRequest {
	s.Spot = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetIsBatchMode(v bool) *AddShowIntoShowListRequest {
	s.IsBatchMode = &v
	return s
}

func (s *AddShowIntoShowListRequest) SetShowList(v []*AddShowIntoShowListRequestShowList) *AddShowIntoShowListRequest {
	s.ShowList = v
	return s
}

func (s *AddShowIntoShowListRequest) Validate() error {
	if s.ShowList != nil {
		for _, item := range s.ShowList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddShowIntoShowListRequestShowList struct {
	// The playback duration of a single show. Unit: seconds.
	//
	// > - You can set only one of **repeatTimes*	- and **duration**.
	//
	// > - This parameter is required when resourceType is set to live.
	//
	// example:
	//
	// 20
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The custom type tag.
	//
	// example:
	//
	// 1
	LiveInputType *int32 `json:"liveInputType,omitempty" xml:"liveInputType,omitempty"`
	// The number of times a single show is repeated. Default value: 0.
	//
	// >- You can set only one of **repeatTimes*	- and **duration**.
	//
	// - repeatTimes specifies the number of repetitions. For example, 0 means the show is repeated 0 times (played once), 1 means the show is repeated 1 time (played twice), and so on.
	//
	// example:
	//
	// 0
	RepeatTimes *int32 `json:"repeatTimes,omitempty" xml:"repeatTimes,omitempty"`
	// The VOD file ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Valid values:
	//
	// - live: live streaming resource.
	//
	// - vod: video-on-demand resource.
	//
	// - pic: image resource.
	//
	// >- When using video-on-demand (VOD) resources, use managed Bucket resources first. Resources in your own Bucket may expire. If you use resources in your own Bucket, check the resource validity period.
	//
	// - Live files support live streaming resources and third-party URLs.
	//
	// - VOD files support video-on-demand resources, image resources, and third-party URLs.
	//
	// - When using live streaming resources, you must also specify the duration parameter.
	//
	// example:
	//
	// vod
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The URL of the playback file.
	ResourceUrl *string `json:"resourceUrl,omitempty" xml:"resourceUrl,omitempty"`
	// The show name.
	//
	// example:
	//
	// liveShow****
	ShowName *string `json:"showName,omitempty" xml:"showName,omitempty"`
}

func (s AddShowIntoShowListRequestShowList) String() string {
	return dara.Prettify(s)
}

func (s AddShowIntoShowListRequestShowList) GoString() string {
	return s.String()
}

func (s *AddShowIntoShowListRequestShowList) GetDuration() *int64 {
	return s.Duration
}

func (s *AddShowIntoShowListRequestShowList) GetLiveInputType() *int32 {
	return s.LiveInputType
}

func (s *AddShowIntoShowListRequestShowList) GetRepeatTimes() *int32 {
	return s.RepeatTimes
}

func (s *AddShowIntoShowListRequestShowList) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddShowIntoShowListRequestShowList) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddShowIntoShowListRequestShowList) GetResourceUrl() *string {
	return s.ResourceUrl
}

func (s *AddShowIntoShowListRequestShowList) GetShowName() *string {
	return s.ShowName
}

func (s *AddShowIntoShowListRequestShowList) SetDuration(v int64) *AddShowIntoShowListRequestShowList {
	s.Duration = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) SetLiveInputType(v int32) *AddShowIntoShowListRequestShowList {
	s.LiveInputType = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) SetRepeatTimes(v int32) *AddShowIntoShowListRequestShowList {
	s.RepeatTimes = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) SetResourceId(v string) *AddShowIntoShowListRequestShowList {
	s.ResourceId = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) SetResourceType(v string) *AddShowIntoShowListRequestShowList {
	s.ResourceType = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) SetResourceUrl(v string) *AddShowIntoShowListRequestShowList {
	s.ResourceUrl = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) SetShowName(v string) *AddShowIntoShowListRequestShowList {
	s.ShowName = &v
	return s
}

func (s *AddShowIntoShowListRequestShowList) Validate() error {
	return dara.Validate(s)
}
