// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterEpisodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *ModifyCasterEpisodeRequest
	GetCasterId() *string
	SetComponentId(v []*string) *ModifyCasterEpisodeRequest
	GetComponentId() []*string
	SetEndTime(v string) *ModifyCasterEpisodeRequest
	GetEndTime() *string
	SetEpisodeId(v string) *ModifyCasterEpisodeRequest
	GetEpisodeId() *string
	SetEpisodeName(v string) *ModifyCasterEpisodeRequest
	GetEpisodeName() *string
	SetOwnerId(v int64) *ModifyCasterEpisodeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCasterEpisodeRequest
	GetRegionId() *string
	SetResourceId(v string) *ModifyCasterEpisodeRequest
	GetResourceId() *string
	SetStartTime(v string) *ModifyCasterEpisodeRequest
	GetStartTime() *string
	SetSwitchType(v string) *ModifyCasterEpisodeRequest
	GetSwitchType() *string
}

type ModifyCasterEpisodeRequest struct {
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
	// The components. Components in the production studio are listed from the bottom to the top in an array. When the production studio switches to another video resource, the components are also switched.
	//
	// 	- This parameter takes effect and is required only when the EpisodeType parameter is set to **Component**.
	//
	// 	- This parameter is optional when the EpisodeType parameter is set to **Resource**. In this case, if this parameter is specified, the components are bound to and switched together with video resources.
	//
	// >  The variable N specifies the sequence number of the component. For example, ComponentId.1 specifies the ID of the first component and ComponentId.2 specifies the ID of the second component.
	//
	// example:
	//
	// ["16A96B9A-F203-4EC5-8E43-CB92E68F****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The time when the episode ends. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T10:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the episode. It is included in the response when you call the [AddCasterEpisode](~~94745#doc-api-live-AddCasterEpisode~~ "Adds an episode to a production studio.") operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938623****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// episode_name_1
	EpisodeName *string `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the video resource. It is included in the response when you call the [AddCasterVideoResource](~~60250#doc-api-live-AddCasterVideoResource~~ "Adds a video resource to a production studio.") operation.
	//
	// 	- This parameter takes effect and is required only when the EpisodeType is set to **Resource**.
	//
	// 	- If the EpisodeType parameter is set to **Component**, this parameter is invalid.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E683****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The time when the episode starts. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The policy for switching episodes. This parameter takes effect only when the EpisodeType parameter is set to **Resource**. Valid values:
	//
	// 	- **TimeFirst**: The episode starts when the preceding episode ends and ends when the next episode starts. If no next episode exists, the episode keeps repeating until a new episode is added or the production studio stops. This parameter must be set to TimeFirst when the video resource is a live stream.
	//
	// 	- **ContentFirst**: The episode starts and ends as scheduled.
	//
	// example:
	//
	// TimeFirst
	SwitchType *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s ModifyCasterEpisodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterEpisodeRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterEpisodeRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterEpisodeRequest) GetComponentId() []*string {
	return s.ComponentId
}

func (s *ModifyCasterEpisodeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyCasterEpisodeRequest) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *ModifyCasterEpisodeRequest) GetEpisodeName() *string {
	return s.EpisodeName
}

func (s *ModifyCasterEpisodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCasterEpisodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCasterEpisodeRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ModifyCasterEpisodeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyCasterEpisodeRequest) GetSwitchType() *string {
	return s.SwitchType
}

func (s *ModifyCasterEpisodeRequest) SetCasterId(v string) *ModifyCasterEpisodeRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetComponentId(v []*string) *ModifyCasterEpisodeRequest {
	s.ComponentId = v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetEndTime(v string) *ModifyCasterEpisodeRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetEpisodeId(v string) *ModifyCasterEpisodeRequest {
	s.EpisodeId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetEpisodeName(v string) *ModifyCasterEpisodeRequest {
	s.EpisodeName = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetOwnerId(v int64) *ModifyCasterEpisodeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetRegionId(v string) *ModifyCasterEpisodeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetResourceId(v string) *ModifyCasterEpisodeRequest {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetStartTime(v string) *ModifyCasterEpisodeRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) SetSwitchType(v string) *ModifyCasterEpisodeRequest {
	s.SwitchType = &v
	return s
}

func (s *ModifyCasterEpisodeRequest) Validate() error {
	return dara.Validate(s)
}
