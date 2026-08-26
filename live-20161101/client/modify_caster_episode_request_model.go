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
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned in the response.
	//
	// - If you created the production studio in the LIVE console, find the ID on the Cloud Production Studio page. To go to the page, choose **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio**.
	//
	// > The name of a production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The IDs of the components. The components are arranged from bottom to top and are switched in sync with the video source.
	//
	// - This parameter is required and takes effect only if EpisodeType is set to **Component**.
	//
	// - If EpisodeType is set to **Resource**, this parameter specifies the components that are attached to the video source and switched in sync.
	//
	// > N specifies the Nth component ID. For example, ComponentId.1 specifies the first component ID and ComponentId.2 specifies the second component ID.
	//
	// example:
	//
	// ["16A96B9A-F203-4EC5-8E43-CB92E68F****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The end time. The time must be in UTC. The format is *yyyy-MM-dd*T*HH:mm:ss*Z.
	//
	// example:
	//
	// 2016-06-29T10:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the episode.
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
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the video source.
	//
	// - This parameter is required and takes effect only if EpisodeType is set to **Resource**.
	//
	// - This parameter is not available if EpisodeType is set to **Component**.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E683****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time. The time must be in UTC. The format is *yyyy-MM-dd*T*HH:mm:ss*Z.
	//
	// example:
	//
	// 2016-06-29T09:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The switch policy. This parameter takes effect only if EpisodeType is set to **Resource**.
	//
	// - **TimeFirst**: time-priority. This is the only policy available for live stream video sources.
	//
	// - **ContentFirst**: content-priority.
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
