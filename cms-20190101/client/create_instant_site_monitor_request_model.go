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
	// The URL or IP address that you want to test.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The detection points. If you leave this parameter empty, the system randomly selects three detection points.
	//
	// The value is a `JSON array`. Example: {"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}. The values of the city field indicate Beijing, Hangzhou, and Qingdao.
	//
	// For information about how to obtain detection points, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html).
	//
	// > You must specify one of the `IspCities` and `RandomIspCity` parameters.
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCities *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	// The extended options of the protocol that is used by the instant test task. The options vary based on the protocol.
	//
	// example:
	//
	// {"time_out":5000}
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	// The number of detection points.
	//
	// >
	//
	// 	- You must specify one of the `IspCities` and `RandomIspCity` parameters. If you specify the `RandomIspCity` parameter, the `IspCities` parameter automatically becomes invalid.
	//
	// example:
	//
	// 1
	RandomIspCity *int32  `json:"RandomIspCity,omitempty" xml:"RandomIspCity,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the instant test task.
	//
	// The name must be 4 to 100 characters in length, and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the instant test task. Valid values: HTTP, PING, TCP, UDP, and DNS.
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
