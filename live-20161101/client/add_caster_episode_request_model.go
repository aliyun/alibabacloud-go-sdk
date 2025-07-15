// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *AddCasterEpisodeRequest
	GetCasterId() *string
	SetComponentId(v []*string) *AddCasterEpisodeRequest
	GetComponentId() []*string
	SetEndTime(v string) *AddCasterEpisodeRequest
	GetEndTime() *string
	SetEpisodeName(v string) *AddCasterEpisodeRequest
	GetEpisodeName() *string
	SetEpisodeType(v string) *AddCasterEpisodeRequest
	GetEpisodeType() *string
	SetOwnerId(v int64) *AddCasterEpisodeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCasterEpisodeRequest
	GetRegionId() *string
	SetResourceId(v string) *AddCasterEpisodeRequest
	GetResourceId() *string
	SetStartTime(v string) *AddCasterEpisodeRequest
	GetStartTime() *string
	SetSwitchType(v string) *AddCasterEpisodeRequest
	GetSwitchType() *string
}

type AddCasterEpisodeRequest struct {
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
	// The components. Components in the production studio are listed from the bottom to the top in an array.
	//
	// If a component was added by calling the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation, check the value of the response parameter ComponentId to obtain the component ID.
	//
	// 	- This parameter takes effect and is required when the EpisodeType parameter is set to **Component**.
	//
	// 	- This parameter is optional when the EpisodeType parameter is set to **Resource**. In this case, if this parameter is specified, the components are bound to and switched together with video resources.
	//
	// >  The variable N specifies the sequence number of the component. For example, **ComponentId.1*	- specifies the ID of the first component and **ComponentId.2*	- specifies the ID of the second component.
	//
	// example:
	//
	// ["a2b8e671-2fe5-4642-a2ec-bf93880e****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The time when the episode ends. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T09:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the episode.
	//
	// example:
	//
	// episode_1
	EpisodeName *string `json:"EpisodeName,omitempty" xml:"EpisodeName,omitempty"`
	// The type of the episode. Valid values:
	//
	// 	- **Resource**: a video resource.
	//
	// 	- **Component**: a component.
	//
	// This parameter is required.
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the video resource.
	//
	// >  This parameter takes effect and is required when the EpisodeType parameter is set to Resource.
	//
	// \\
	//
	// If the video resource was added by calling the [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html) operation, check the value of the response parameter ResourceId to obtain the ID.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The time when the episode starts. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The policy for switching episodes. Valid values:
	//
	// 	- **TimeFirst**: The episode starts when the preceding episode ends and ends when the next episode starts. If no next episode exists, the episode keeps repeating until a new episode is added or the production studio stops.
	//
	// 	- **ContentFirst**: The episode starts and ends as scheduled.
	//
	// This parameter takes effect only when the EpisodeType parameter is set to Resource.
	//
	// >  This parameter must be set to TimeFirst when the video resource is a live stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// TimeFirst
	SwitchType *string `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
}

func (s AddCasterEpisodeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeRequest) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddCasterEpisodeRequest) GetComponentId() []*string {
	return s.ComponentId
}

func (s *AddCasterEpisodeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *AddCasterEpisodeRequest) GetEpisodeName() *string {
	return s.EpisodeName
}

func (s *AddCasterEpisodeRequest) GetEpisodeType() *string {
	return s.EpisodeType
}

func (s *AddCasterEpisodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCasterEpisodeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCasterEpisodeRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddCasterEpisodeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *AddCasterEpisodeRequest) GetSwitchType() *string {
	return s.SwitchType
}

func (s *AddCasterEpisodeRequest) SetCasterId(v string) *AddCasterEpisodeRequest {
	s.CasterId = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetComponentId(v []*string) *AddCasterEpisodeRequest {
	s.ComponentId = v
	return s
}

func (s *AddCasterEpisodeRequest) SetEndTime(v string) *AddCasterEpisodeRequest {
	s.EndTime = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetEpisodeName(v string) *AddCasterEpisodeRequest {
	s.EpisodeName = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetEpisodeType(v string) *AddCasterEpisodeRequest {
	s.EpisodeType = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetOwnerId(v int64) *AddCasterEpisodeRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetRegionId(v string) *AddCasterEpisodeRequest {
	s.RegionId = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetResourceId(v string) *AddCasterEpisodeRequest {
	s.ResourceId = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetStartTime(v string) *AddCasterEpisodeRequest {
	s.StartTime = &v
	return s
}

func (s *AddCasterEpisodeRequest) SetSwitchType(v string) *AddCasterEpisodeRequest {
	s.SwitchType = &v
	return s
}

func (s *AddCasterEpisodeRequest) Validate() error {
	return dara.Validate(s)
}
