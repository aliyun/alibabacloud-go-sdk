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
	// - If you create a production studio by calling the [CreateCaster]() operation, use the value of the CasterId parameter that is returned in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The program information.
	//
	// This parameter is required.
	Episode []*ModifyCasterProgramRequestEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" type:"Repeated"`
	OwnerId *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The list of component IDs. The components are layered from bottom to top in the order they are listed. The components are switched in sync with the video source.
	//
	// - This parameter is required and takes effect only when the node type is **Component**.
	//
	// - If the node type is **Resource**, the components are attached to the video source and switched in sync.
	//
	// example:
	//
	// ["a2b8e671-2fe5-4642-a2ec-bf93888****" ]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The end time. The time is in UTC. The format is *yyyy-MM-dd*T*HH:mm:ss*Z.
	//
	// example:
	//
	// 2016-06-29T10:04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The program ID. If you add a program for the production studio by calling the [AddCasterEpisode]() operation, use the value of the EpisodeId parameter that is returned in the response.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938887****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The program name.
	//
	// example:
	//
	// program_name_2
	EpisodeName *string `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	// The program type. Valid values:
	//
	// - **Resource**: video source.
	//
	// - **Component**: component.
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	// The ID of the video source. If you add a video source for the production studio by calling the [AddCasterVideoResource]() operation, use the value of the ResourceId parameter that is returned in the response.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf938887****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time. The time is in UTC. The format is *yyyy-MM-dd*T*HH:mm:ss*Z.
	//
	// example:
	//
	// 2016-06-29T09:02:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The switch policy. This parameter is valid only when the node type is **Resource**.
	//
	// - **TimeFirst**: time-first. This is the only valid policy for live stream video sources.
	//
	// - **ContentFirst**: content-first.
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
