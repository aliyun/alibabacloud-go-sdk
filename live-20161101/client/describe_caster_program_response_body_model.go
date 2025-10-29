// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterProgramResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterProgramResponseBody
	GetCasterId() *string
	SetEpisodes(v *DescribeCasterProgramResponseBodyEpisodes) *DescribeCasterProgramResponseBody
	GetEpisodes() *DescribeCasterProgramResponseBodyEpisodes
	SetProgramEffect(v int32) *DescribeCasterProgramResponseBody
	GetProgramEffect() *int32
	SetProgramName(v string) *DescribeCasterProgramResponseBody
	GetProgramName() *string
	SetRequestId(v string) *DescribeCasterProgramResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCasterProgramResponseBody
	GetTotal() *int32
}

type DescribeCasterProgramResponseBody struct {
	// The ID of the production studio.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The list of the episodes.
	Episodes *DescribeCasterProgramResponseBodyEpisodes `json:"Episodes,omitempty" xml:"Episodes,omitempty" type:"Struct"`
	// Indicates whether carousel playback is enabled.
	//
	// 	- **0**: Carousel playback is disabled.
	//
	// 	- **1**: Carousel playback is enabled.
	//
	// example:
	//
	// 1
	ProgramEffect *int32 `json:"ProgramEffect,omitempty" xml:"ProgramEffect,omitempty"`
	// The name of the episode list.
	//
	// example:
	//
	// programs_name
	ProgramName *string `json:"ProgramName,omitempty" xml:"ProgramName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCasterProgramResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterProgramResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterProgramResponseBody) GetEpisodes() *DescribeCasterProgramResponseBodyEpisodes {
	return s.Episodes
}

func (s *DescribeCasterProgramResponseBody) GetProgramEffect() *int32 {
	return s.ProgramEffect
}

func (s *DescribeCasterProgramResponseBody) GetProgramName() *string {
	return s.ProgramName
}

func (s *DescribeCasterProgramResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCasterProgramResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCasterProgramResponseBody) SetCasterId(v string) *DescribeCasterProgramResponseBody {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterProgramResponseBody) SetEpisodes(v *DescribeCasterProgramResponseBodyEpisodes) *DescribeCasterProgramResponseBody {
	s.Episodes = v
	return s
}

func (s *DescribeCasterProgramResponseBody) SetProgramEffect(v int32) *DescribeCasterProgramResponseBody {
	s.ProgramEffect = &v
	return s
}

func (s *DescribeCasterProgramResponseBody) SetProgramName(v string) *DescribeCasterProgramResponseBody {
	s.ProgramName = &v
	return s
}

func (s *DescribeCasterProgramResponseBody) SetRequestId(v string) *DescribeCasterProgramResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCasterProgramResponseBody) SetTotal(v int32) *DescribeCasterProgramResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCasterProgramResponseBody) Validate() error {
	if s.Episodes != nil {
		if err := s.Episodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterProgramResponseBodyEpisodes struct {
	Episode []*DescribeCasterProgramResponseBodyEpisodesEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" type:"Repeated"`
}

func (s DescribeCasterProgramResponseBodyEpisodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterProgramResponseBodyEpisodes) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseBodyEpisodes) GetEpisode() []*DescribeCasterProgramResponseBodyEpisodesEpisode {
	return s.Episode
}

func (s *DescribeCasterProgramResponseBodyEpisodes) SetEpisode(v []*DescribeCasterProgramResponseBodyEpisodesEpisode) *DescribeCasterProgramResponseBodyEpisodes {
	s.Episode = v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodes) Validate() error {
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

type DescribeCasterProgramResponseBodyEpisodesEpisode struct {
	// The components.
	ComponentIds *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds `json:"ComponentIds,omitempty" xml:"ComponentIds,omitempty" type:"Struct"`
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-06-29T10:02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the episode.
	//
	// example:
	//
	// 1872639A-F203-4EC5-8E43-CB92E68F****
	EpisodeId *string `json:"EpisodeId,omitempty" xml:"EpisodeId,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// program_name_1
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
	// The ID of the video resource.
	//
	// example:
	//
	// 1872639A-F203-4EC5-8E43-CB92E837****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The beginning of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the episode.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The policy for switching episodes. Valid values:
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

func (s DescribeCasterProgramResponseBodyEpisodesEpisode) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterProgramResponseBodyEpisodesEpisode) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetComponentIds() *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds {
	return s.ComponentIds
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetEpisodeId() *string {
	return s.EpisodeId
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetEpisodeName() *string {
	return s.EpisodeName
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetEpisodeType() *string {
	return s.EpisodeType
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) GetSwitchType() *string {
	return s.SwitchType
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetComponentIds(v *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.ComponentIds = v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetEndTime(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.EndTime = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetEpisodeId(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.EpisodeId = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetEpisodeName(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.EpisodeName = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetEpisodeType(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.EpisodeType = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetResourceId(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.ResourceId = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetStartTime(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.StartTime = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetStatus(v int32) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.Status = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) SetSwitchType(v string) *DescribeCasterProgramResponseBodyEpisodesEpisode {
	s.SwitchType = &v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisode) Validate() error {
	if s.ComponentIds != nil {
		if err := s.ComponentIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds struct {
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
}

func (s DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds) GoString() string {
	return s.String()
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds) GetComponentId() []*string {
	return s.ComponentId
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds) SetComponentId(v []*string) *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds {
	s.ComponentId = v
	return s
}

func (s *DescribeCasterProgramResponseBodyEpisodesEpisodeComponentIds) Validate() error {
	return dara.Validate(s)
}
