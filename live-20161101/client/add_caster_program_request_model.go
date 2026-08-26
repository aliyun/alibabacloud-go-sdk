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
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster operation](https://help.aliyun.com/document_detail/2848009.html), check the CasterId value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, navigate to **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the production studio name.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The program list information.
	//
	// This parameter is required.
	Episode []*AddCasterProgramRequestEpisode `json:"Episode,omitempty" xml:"Episode,omitempty" type:"Repeated"`
	OwnerId *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The component list. Elements are arranged from bottom to top in order.
	//
	// 	Notice: This parameter is valid and required when Episode.N.EpisodeType is set to **Component**.
	//
	//
	//  When the node type is **Resource**, this indicates that the component is bound to the video source and switches synchronously.
	//
	// example:
	//
	// [ "a2b8e671-2fe5-4642-a2ec-bf931826****",  "a2b8e671-2fe5-4642-a2ec-28374657****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The end time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC). This parameter is required. If not specified, MissingParameter is returned.
	//
	// example:
	//
	// 2016-06-29T10:02:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The program name.
	//
	// example:
	//
	// program_name_1
	EpisodeName *string `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	// The node type. Valid values:
	//
	//
	//
	// - **Resource**: video source. If you select Resource, you must also set the request parameters Episode.N.ResourceId and Episode.N.SwitchType.
	//
	// - **Component**: component. If you select Component, you must also set the request parameter Episode.N.ComponentId.N.
	//
	//
	// >
	//
	// > - When Resource is selected and the referenced resource contains a VodUrl (video-on-demand file), EndTime - StartTime cannot exceed the actual playback duration (in seconds) of the VOD file. Otherwise, InvalidParameter.EndTime is returned.
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	// The video source ID.
	//
	// 	Notice: This parameter is valid and required when Episode.N.EpisodeType is set to **Resource**.
	//
	//
	//
	//  This parameter is not applicable when Episode.N.EpisodeType is set to **Component**.
	//
	// If you added the video source by calling the [AddCasterVideoResource operation](https://help.aliyun.com/document_detail/60250.html), check the ResourceId value returned by the AddCasterVideoResource operation.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time. Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC). This parameter is required. If not specified, MissingParameter is returned.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The switch policy. Valid values:
	//
	// 	Notice: This parameter is valid only when Episode.N.EpisodeType is set to **Resource**.
	//
	//
	//
	// - **TimeFirst**: time first. Live video sources can only use the time first policy.
	//
	// - **ContentFirst**: content first.
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
