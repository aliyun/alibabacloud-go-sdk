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
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the CasterId parameter that is returned.
	//
	// - If you create a production studio in the LIVE console, go to the **LIVE Console**> **Production Studio*	- > **Production Studio*	- page to view the ID.
	//
	// > The name of the production studio in the production studio list serves as the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// A list of component IDs. The components are layered from bottom to top in the specified order.
	//
	// If you add a component by calling the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation, check the value of the ComponentId parameter that is returned.
	//
	// - This parameter is required and applies only when the resource type is **Component**.
	//
	// - This parameter is optional when the resource type is **Resource**. If you specify this parameter, the component is attached to the video source and they are switched synchronously.
	//
	// > N specifies the sequence number of a component ID. For example, **ComponentId.1*	- specifies the first component ID and **ComponentId.2*	- specifies the second component ID.
	//
	// example:
	//
	// ["a2b8e671-2fe5-4642-a2ec-bf93880e****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The end time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
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
	// The node type. Valid values:
	//
	// - **Resource**: A video source. If you set this parameter to Resource, you must also specify the ResourceId and SwitchType parameters.
	//
	// - **Component**: A component.
	//
	// This parameter is required.
	//
	// example:
	//
	// Resource
	EpisodeType *string `json:"EpisodeType,omitempty" xml:"EpisodeType,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the video source.
	//
	// 	Notice:
	//
	// This parameter is required and applies only when EpisodeType is set to Resource.
	//
	//
	//
	// If you add a video source by calling the [AddCasterVideoResource](https://help.aliyun.com/document_detail/2848020.html) operation, check the value of the ResourceId parameter that is returned.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The switch policy. Valid values:
	//
	// 	Notice:
	//
	// This parameter applies only when EpisodeType is set to Resource.
	//
	//
	//
	// - **TimeFirst**: Time first.
	//
	// - **ContentFirst**: Content first.
	//
	// > For more information about video sources, see [Add a video source](https://help.aliyun.com/document_detail/66094.html).
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
