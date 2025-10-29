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
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The duration of the episode. Unit: seconds.
	//
	// > You can specify only one of the **RepeatTimes*	- and **Duration*	- parameters.
	//
	// example:
	//
	// 20
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The custom type label.
	//
	// example:
	//
	// 1
	LiveInputType *int32  `json:"LiveInputType,omitempty" xml:"LiveInputType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of times the episode repeats after the first playback is complete. The default value is 0.
	//
	// >
	//
	// 	- You can specify only one of the **RepeatTimes*	- and **Duration*	- parameters. - The RepeatTimes parameter specifies the number of repetitions. For example, if you set the value to -1, the episode is to be played for infinite times. If you set the value to 0, the episode is to be played once. If you set the value to 1, the episode is to be played twice.
	//
	// example:
	//
	// 0
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- live: live stream
	//
	// 	- vod: on-demand video
	//
	// 	- pic: image
	//
	// >
	//
	// 	- When you select media resources from ApsaraVideo VOD, we recommend that you select resources that are stored in hosted OSS buckets. Resources stored in non-hosted OSS buckets have a validity period. Pay attention to the validity if you select resources that are stored in non-hosted OSS buckets. - You can add a live stream from ApsaraVideo Live or by using a third-party URL. - You can add an on-demand video from ApsaraVideo VOD or by using a third-party URL, or add an on-demand image.
	//
	// example:
	//
	// vod
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The URL of the resource.
	ResourceUrl *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// liveShow****
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The position of the episode in the episode list. Position indexes start from 0. By default, the episode is added to the end of the episode list.
	//
	// example:
	//
	// 1
	Spot *int32 `json:"Spot,omitempty" xml:"Spot,omitempty"`
	// Specifies whether to add multiple episodes to the episode list at a time. Valid values:
	//
	// 	- true: adds multiple episodes to the episode list at a time.
	//
	// 	- false: adds a single episode to the episode list.
	//
	// > If you do not specify this parameter or this parameter is left empty, a single episode is to be added to the episode list.
	//
	// example:
	//
	// false
	IsBatchMode *bool `json:"isBatchMode,omitempty" xml:"isBatchMode,omitempty"`
	// The episodes that you want to add to the episode list. Each episode has a unique name and resource URL.
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
	// The duration of the episode. Unit: seconds.
	//
	// >  You can specify only one of the **RepeatTimes*	- and **Duration*	- parameters.
	//
	// example:
	//
	// 20
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The custom type label.
	//
	// example:
	//
	// 1
	LiveInputType *int32 `json:"liveInputType,omitempty" xml:"liveInputType,omitempty"`
	// The number of times the episode repeats after the first playback is complete. Default value: 0.
	//
	// >
	//
	// 	- You can specify only one of the **RepeatTimes*	- and **Duration*	- parameters.
	//
	// 	- The RepeatTimes parameter specifies the number of repetitions. For example, if you set the value to 0, the episode is to be played once. If you set the value to 1, the episode is to be played twice.
	//
	// example:
	//
	// 0
	RepeatTimes *int32 `json:"repeatTimes,omitempty" xml:"repeatTimes,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- live: live stream
	//
	// 	- vod: on-demand video
	//
	// 	- pic: image
	//
	// >
	//
	// 	- When you select media resources from ApsaraVideo VOD, we recommend that you select resources that are stored in hosted OSS buckets. Resources stored in non-hosted OSS buckets have a validity period. Pay attention to the validity if you select resources that are stored in non-hosted OSS buckets.
	//
	// 	- You can add a live stream from ApsaraVideo Live or by using a third-party URL.
	//
	// 	- You can add an on-demand video from ApsaraVideo VOD or by using a third-party URL, or add an on-demand image.
	//
	// example:
	//
	// vod
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The URL of the resource.
	ResourceUrl *string `json:"resourceUrl,omitempty" xml:"resourceUrl,omitempty"`
	// The name of the episode.
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
