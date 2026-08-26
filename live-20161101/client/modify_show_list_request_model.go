// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyShowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *ModifyShowListRequest
	GetCasterId() *string
	SetHighPriorityShowId(v string) *ModifyShowListRequest
	GetHighPriorityShowId() *string
	SetHighPriorityShowStartTime(v string) *ModifyShowListRequest
	GetHighPriorityShowStartTime() *string
	SetOwnerId(v int64) *ModifyShowListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyShowListRequest
	GetRegionId() *string
	SetRepeatTimes(v int32) *ModifyShowListRequest
	GetRepeatTimes() *int32
	SetShowId(v string) *ModifyShowListRequest
	GetShowId() *string
	SetSpot(v int32) *ModifyShowListRequest
	GetSpot() *int32
}

type ModifyShowListRequest struct {
	// The ID of the production studio.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value returned in the response.
	//
	// - If you created the production studio in the LIVE console, find the production studio name on the Cloud Production Studio page. To go to the page, choose **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio**.
	//
	// > The name of the production studio on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The highest-priority show.
	//
	// > This parameter can be configured only before the playlist starts.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	HighPriorityShowId *string `json:"HighPriorityShowId,omitempty" xml:"HighPriorityShowId,omitempty"`
	// The time to play the highest-priority show. The format is yyyy-MM-dd\\"T\\"HH:mm:ss.
	//
	// > This parameter can be configured only before the playlist starts.<br>
	//
	// > After this parameter is configured, the system switches from the currently playing show to the highest-priority show at the specified time.
	//
	// example:
	//
	// 2021-11-23T12:30:00
	HighPriorityShowStartTime *string `json:"HighPriorityShowStartTime,omitempty" xml:"HighPriorityShowStartTime,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of times the playlist loops.
	//
	// > - RepeatTimes specifies the number of repetitions. For example, a value of **0*	- means the playlist is played once without repetition. A value of **1*	- means the playlist is played twice (one initial playback and one repetition).
	//
	// >
	//
	// > - A value of -1 indicates that the playlist loops indefinitely.
	//
	// example:
	//
	// 5
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// The ID of the show whose position in the playlist you want to modify.
	//
	// > Obtain the ShowId value from the response of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) or [DescribeShowList](https://help.aliyun.com/document_detail/2848054.html) operation.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	// The new position of the show in the playlist. The show specified by ShowId is moved to the position specified by **Spot**.
	//
	// > The value must be greater than or equal to 0 and less than or equal to the total number of shows in the playlist.
	//
	// example:
	//
	// 1
	Spot *int32 `json:"Spot,omitempty" xml:"Spot,omitempty"`
}

func (s ModifyShowListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyShowListRequest) GoString() string {
	return s.String()
}

func (s *ModifyShowListRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyShowListRequest) GetHighPriorityShowId() *string {
	return s.HighPriorityShowId
}

func (s *ModifyShowListRequest) GetHighPriorityShowStartTime() *string {
	return s.HighPriorityShowStartTime
}

func (s *ModifyShowListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyShowListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyShowListRequest) GetRepeatTimes() *int32 {
	return s.RepeatTimes
}

func (s *ModifyShowListRequest) GetShowId() *string {
	return s.ShowId
}

func (s *ModifyShowListRequest) GetSpot() *int32 {
	return s.Spot
}

func (s *ModifyShowListRequest) SetCasterId(v string) *ModifyShowListRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyShowListRequest) SetHighPriorityShowId(v string) *ModifyShowListRequest {
	s.HighPriorityShowId = &v
	return s
}

func (s *ModifyShowListRequest) SetHighPriorityShowStartTime(v string) *ModifyShowListRequest {
	s.HighPriorityShowStartTime = &v
	return s
}

func (s *ModifyShowListRequest) SetOwnerId(v int64) *ModifyShowListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyShowListRequest) SetRegionId(v string) *ModifyShowListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyShowListRequest) SetRepeatTimes(v int32) *ModifyShowListRequest {
	s.RepeatTimes = &v
	return s
}

func (s *ModifyShowListRequest) SetShowId(v string) *ModifyShowListRequest {
	s.ShowId = &v
	return s
}

func (s *ModifyShowListRequest) SetSpot(v int32) *ModifyShowListRequest {
	s.Spot = &v
	return s
}

func (s *ModifyShowListRequest) Validate() error {
	return dara.Validate(s)
}
