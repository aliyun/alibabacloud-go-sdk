// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *ModifyCasterProgramRequest
	GetCasterId() *string
	SetEpisode(v []*ModifyCasterProgramRequestEpisode) *ModifyCasterProgramRequest
	GetEpisode() []*ModifyCasterProgramRequestEpisode
	SetOwnerId(v int64) *ModifyCasterProgramRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCasterProgramRequest
	GetRegionId() *string
}

type ModifyCasterProgramRequest struct {
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
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The episodes.
	//
	// This parameter is required.
	Episode  []*ModifyCasterProgramRequestEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" type:"Repeated"`
	OwnerId  *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyCasterProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyCasterProgramRequest) GetEpisode() []*ModifyCasterProgramRequestEpisode {
	return s.Episode
}

func (s *ModifyCasterProgramRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCasterProgramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCasterProgramRequest) SetCasterId(v string) *ModifyCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyCasterProgramRequest) SetEpisode(v []*ModifyCasterProgramRequestEpisode) *ModifyCasterProgramRequest {
	s.Episode = v
	return s
}

func (s *ModifyCasterProgramRequest) SetOwnerId(v int64) *ModifyCasterProgramRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCasterProgramRequest) SetRegionId(v string) *ModifyCasterProgramRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCasterProgramRequest) Validate() error {
	if s.Episode != nil {
		for _, item := range s.Episode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCasterProgramRequestEpisode struct {
	// The components. Components in the production studio are listed from the bottom to the top in an array. When the production studio switches to another video resource, the components are also switched.
	//
	// 	- This parameter is required and available only when EpisodeType is set to **Component**.
	//
	// 	- This parameter is optional when EpisodeType is set to **Resource**. This indicates that the components are bound to and switched together with video resources.
	//
	// example:
	//
	// ["a2b8e671-2fe5-4642-a2ec-bf93888****" ]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The end time of the episode. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T10:04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the episode. If the episode was added by calling the [AddCasterEpisode](https://help.aliyun.com/document_detail/2848068.html) operation, check the value of the response parameter EpisodeId to obtain the ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938887****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// program_name_2
	EpisodeName *string `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	// The type of the episode. Valid values:
	//
	// 	- **Resource**: a video resource
	//
	// 	- **Component**: a component
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	// The ID of the video resource. If the video resource was added by calling the [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html) operation, check the value of the response parameter ResourceId to obtain the ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938887****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time of the episode. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T09:02:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The policy for switching episodes. This parameter takes effect only when EpisodeType is set to **Resource**. Valid values:
	//
	// 	- **TimeFirst**: The episode starts when the previous episode ends and ends when the next episode starts. If no next episode exists, the episode keeps repeating until a new episode is added or the production studio stops. This value is required for live video resources.
	//
	// 	- **ContentFirst**: The episode starts and ends as scheduled.
	//
	// example:
	//
	// TimeFirst
	SwitchType *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s ModifyCasterProgramRequestEpisode) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterProgramRequestEpisode) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramRequestEpisode) GetComponentId() []*string {
	return s.ComponentId
}

func (s *ModifyCasterProgramRequestEpisode) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyCasterProgramRequestEpisode) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *ModifyCasterProgramRequestEpisode) GetEpisodeName() *string {
	return s.EpisodeName
}

func (s *ModifyCasterProgramRequestEpisode) GetEpisodeType() *string {
	return s.EpisodeType
}

func (s *ModifyCasterProgramRequestEpisode) GetResourceId() *string {
	return s.ResourceId
}

func (s *ModifyCasterProgramRequestEpisode) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyCasterProgramRequestEpisode) GetSwitchType() *string {
	return s.SwitchType
}

func (s *ModifyCasterProgramRequestEpisode) SetComponentId(v []*string) *ModifyCasterProgramRequestEpisode {
	s.ComponentId = v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEndTime(v string) *ModifyCasterProgramRequestEpisode {
	s.EndTime = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEpisodeId(v string) *ModifyCasterProgramRequestEpisode {
	s.EpisodeId = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEpisodeName(v string) *ModifyCasterProgramRequestEpisode {
	s.EpisodeName = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetEpisodeType(v string) *ModifyCasterProgramRequestEpisode {
	s.EpisodeType = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetResourceId(v string) *ModifyCasterProgramRequestEpisode {
	s.ResourceId = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetStartTime(v string) *ModifyCasterProgramRequestEpisode {
	s.StartTime = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) SetSwitchType(v string) *ModifyCasterProgramRequestEpisode {
	s.SwitchType = &v
	return s
}

func (s *ModifyCasterProgramRequestEpisode) Validate() error {
	return dara.Validate(s)
}
