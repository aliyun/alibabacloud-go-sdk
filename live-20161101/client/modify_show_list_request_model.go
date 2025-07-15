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
	// The episode of the highest priority.
	//
	// >  You can configure this parameter only before the playback of the episode list starts.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	HighPriorityShowId *string `json:"HighPriorityShowId,omitempty" xml:"HighPriorityShowId,omitempty"`
	// The time at which the episode of the highest priority is played. Format: yyyy-MM-dd\\"T\\"HH:mm:ss.
	//
	// >  You can configure this parameter only before the episode list starts playing.\\
	//
	// After you configure this parameter, when the specified point in time is reached, any episode that is playing stops and the episode of the highest priority in the episode list starts to play.
	//
	// example:
	//
	// 2021-11-23T12:30:00
	HighPriorityShowStartTime *string `json:"HighPriorityShowStartTime,omitempty" xml:"HighPriorityShowStartTime,omitempty"`
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of additional times the episode list is played.
	//
	// >
	//
	// 	- The RepeatTimes parameter specifies the number of repetitions. For example, if you set the value to **0**, the episode list is played **once**. If you set the value to **1**, the episode list is played **twice**.********
	//
	// 	- If you set the value to -1, the episode list is repeated indefinitely.
	//
	// example:
	//
	// 5
	RepeatTimes *int32 `json:"RepeatTimes,omitempty" xml:"RepeatTimes,omitempty"`
	// The ID of the episode for which you want to change the position in the playlist.
	//
	// >  You can call the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) or [DescribeShowList](https://help.aliyun.com/document_detail/2848054.html) operation and check the value of the response parameter ShowId to obtain the ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
	// The position of the episode in the episode list. If you want to change the position of an episode in a playlist, place the ID of the episode in **Spot**.
	//
	// >  The value must be greater than or equal to 0 and less than or equal to the total number of episodes in the playlist.
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
