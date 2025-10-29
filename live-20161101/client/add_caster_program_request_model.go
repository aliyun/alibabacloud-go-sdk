// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterProgramRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *AddCasterProgramRequest
	GetCasterId() *string
	SetEpisode(v []*AddCasterProgramRequestEpisode) *AddCasterProgramRequest
	GetEpisode() []*AddCasterProgramRequestEpisode
	SetOwnerId(v int64) *AddCasterProgramRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCasterProgramRequest
	GetRegionId() *string
}

type AddCasterProgramRequest struct {
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
	// The information about episodes in the episode list.
	//
	// This parameter is required.
	Episode  []*AddCasterProgramRequestEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" type:"Repeated"`
	OwnerId  *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddCasterProgramRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasterProgramRequest) GoString() string {
	return s.String()
}

func (s *AddCasterProgramRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddCasterProgramRequest) GetEpisode() []*AddCasterProgramRequestEpisode {
	return s.Episode
}

func (s *AddCasterProgramRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCasterProgramRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCasterProgramRequest) SetCasterId(v string) *AddCasterProgramRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterProgramRequest) SetEpisode(v []*AddCasterProgramRequestEpisode) *AddCasterProgramRequest {
	s.Episode = v
	return s
}

func (s *AddCasterProgramRequest) SetOwnerId(v int64) *AddCasterProgramRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCasterProgramRequest) SetRegionId(v string) *AddCasterProgramRequest {
	s.RegionId = &v
	return s
}

func (s *AddCasterProgramRequest) Validate() error {
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

type AddCasterProgramRequestEpisode struct {
	// The components. Components in the production studio are listed from the bottom to the top in an array.
	//
	// >  This parameter is required and takes effect when the Episode.N.EpisodeType parameter is set to Component.
	//
	// This parameter is optional when the Episode.N.EpisodeType parameter is set to **Resource**. In this case, if this parameter is specified, the components are bound to and switched together with video resources.
	//
	// example:
	//
	// [ "a2b8e671-2fe5-4642-a2ec-bf931826****",  "a2b8e671-2fe5-4642-a2ec-28374657****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The end time of the episode. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T10:02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// program_name_1
	EpisodeName *string `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	// The type of the episode.
	//
	// 	- **Resource**: a video resource If you set this parameter to Resource, you must specify the Episode.N.ResourceId and Episode.N.SwitchType parameters.
	//
	// 	- **Component**: a component If you set this parameter to Component, you must specify the Episode.N.ComponentId.N parameter.
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	// The ID of the video resource.
	//
	// >  This parameter takes effect and is required when the Episode.N.EpisodeType parameter is set to Resource.
	//
	// \\
	//
	// This parameter is invalid if you set the Episode.N.EpisodeType parameter to **Component**.
	//
	// If the video resource was added by calling the [AddCasterVideoResource](https://help.aliyun.com/document_detail/60250.html) operation, check the value of the response parameter ResourceId to obtain the ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time of the episode. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The policy for switching episodes. Valid values:
	//
	// >  This parameter takes effect only when the Episode.N.EpisodeType parameter is set to Resource.
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

func (s AddCasterProgramRequestEpisode) String() string {
	return dara.Prettify(s)
}

func (s AddCasterProgramRequestEpisode) GoString() string {
	return s.String()
}

func (s *AddCasterProgramRequestEpisode) GetComponentId() []*string {
	return s.ComponentId
}

func (s *AddCasterProgramRequestEpisode) GetEndTime() *string {
	return s.EndTime
}

func (s *AddCasterProgramRequestEpisode) GetEpisodeName() *string {
	return s.EpisodeName
}

func (s *AddCasterProgramRequestEpisode) GetEpisodeType() *string {
	return s.EpisodeType
}

func (s *AddCasterProgramRequestEpisode) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddCasterProgramRequestEpisode) GetStartTime() *string {
	return s.StartTime
}

func (s *AddCasterProgramRequestEpisode) GetSwitchType() *string {
	return s.SwitchType
}

func (s *AddCasterProgramRequestEpisode) SetComponentId(v []*string) *AddCasterProgramRequestEpisode {
	s.ComponentId = v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetEndTime(v string) *AddCasterProgramRequestEpisode {
	s.EndTime = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetEpisodeName(v string) *AddCasterProgramRequestEpisode {
	s.EpisodeName = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetEpisodeType(v string) *AddCasterProgramRequestEpisode {
	s.EpisodeType = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetResourceId(v string) *AddCasterProgramRequestEpisode {
	s.ResourceId = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetStartTime(v string) *AddCasterProgramRequestEpisode {
	s.StartTime = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) SetSwitchType(v string) *AddCasterProgramRequestEpisode {
	s.SwitchType = &v
	return s
}

func (s *AddCasterProgramRequestEpisode) Validate() error {
	return dara.Validate(s)
}
