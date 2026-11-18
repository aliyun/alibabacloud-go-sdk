// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstantSiteMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateInstantSiteMonitorRequest
	GetAddress() *string
	SetAgentGroup(v string) *CreateInstantSiteMonitorRequest
	GetAgentGroup() *string
	SetIspCities(v string) *CreateInstantSiteMonitorRequest
	GetIspCities() *string
	SetOptionsJson(v string) *CreateInstantSiteMonitorRequest
	GetOptionsJson() *string
	SetRandomIspCity(v int32) *CreateInstantSiteMonitorRequest
	GetRandomIspCity() *int32
	SetRegionId(v string) *CreateInstantSiteMonitorRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateInstantSiteMonitorRequest
	GetTaskName() *string
	SetTaskType(v string) *CreateInstantSiteMonitorRequest
	GetTaskType() *string
}

type CreateInstantSiteMonitorRequest struct {
	// The URL or IP address of the detection task.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The type of the detection point. Valid values: PC and MOBILE. PC indicates a fixed broadband detection point. MOBILE indicates a mobile detection point. Default value: PC.
	//
	// example:
	//
	// PC
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// The information about the detection points. If this parameter is left empty, the system randomly selects three detection points.
	//
	// The value is a JSON array. Example: [{"city":"546","isp":"465", "type":"IDC"},{"city":"572","isp":"465", "type":"LASTMILE"},{"city":"738","isp":"465"}], which correspond to Beijing, Hangzhou, and Qingdao.
	//
	// type indicates the type of the detection point. When AgentGroup is set to PC, valid values are IDC and LASTMILE. IDC indicates a detection point deployed in a data center. LASTMILE indicates a detection point deployed in a residential area that uses a fixed broadband network. type is optional and defaults to IDC. When AgentGroup is set to MOBILE, you do not need to specify type.
	//
	// For more information about how to obtain detection point information, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html).
	//
	// > You must specify either IspCities or RandomIspCity.
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCities *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	// The advanced extension options for the protocol type of the detection task. Different protocol types correspond to different extension options.
	//
	// example:
	//
	// {"time_out":5000}
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The number of detection points.
	//
	// > - You must specify either IspCities or RandomIspCity. - If you specify RandomIspCity, IspCities automatically becomes invalid.
	//
	// example:
	//
	// 1
	RandomIspCity *int32  `json:"RandomIspCity,omitempty" xml:"RandomIspCity,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the detection task.
	//
	// <props="china">The name must be 4 to 100 characters in length and can contain letters, digits, underscores (_), and Chinese characters.
	//
	//
	// <props="intl">The name must be 4 to 100 characters in length and can contain letters, digits, and underscores (_).
	//
	// <props="partner">The name must be 4 to 100 characters in length and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the detection task. Valid values: HTTP, PING, TCP, UDP, and DNS.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateInstantSiteMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstantSiteMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateInstantSiteMonitorRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateInstantSiteMonitorRequest) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *CreateInstantSiteMonitorRequest) GetIspCities() *string {
	return s.IspCities
}

func (s *CreateInstantSiteMonitorRequest) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *CreateInstantSiteMonitorRequest) GetRandomIspCity() *int32 {
	return s.RandomIspCity
}

func (s *CreateInstantSiteMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstantSiteMonitorRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateInstantSiteMonitorRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateInstantSiteMonitorRequest) SetAddress(v string) *CreateInstantSiteMonitorRequest {
	s.Address = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetAgentGroup(v string) *CreateInstantSiteMonitorRequest {
	s.AgentGroup = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetIspCities(v string) *CreateInstantSiteMonitorRequest {
	s.IspCities = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetOptionsJson(v string) *CreateInstantSiteMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetRandomIspCity(v int32) *CreateInstantSiteMonitorRequest {
	s.RandomIspCity = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetRegionId(v string) *CreateInstantSiteMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetTaskName(v string) *CreateInstantSiteMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) SetTaskType(v string) *CreateInstantSiteMonitorRequest {
	s.TaskType = &v
	return s
}

func (s *CreateInstantSiteMonitorRequest) Validate() error {
	return dara.Validate(s)
}
