// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateSiteMonitorRequest
	GetAddress() *string
	SetAgentGroup(v string) *CreateSiteMonitorRequest
	GetAgentGroup() *string
	SetAlertIds(v string) *CreateSiteMonitorRequest
	GetAlertIds() *string
	SetCustomSchedule(v string) *CreateSiteMonitorRequest
	GetCustomSchedule() *string
	SetInterval(v string) *CreateSiteMonitorRequest
	GetInterval() *string
	SetIspCities(v string) *CreateSiteMonitorRequest
	GetIspCities() *string
	SetOptionsJson(v string) *CreateSiteMonitorRequest
	GetOptionsJson() *string
	SetRegionId(v string) *CreateSiteMonitorRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateSiteMonitorRequest
	GetTaskName() *string
	SetTaskType(v string) *CreateSiteMonitorRequest
	GetTaskType() *string
	SetVpcConfig(v string) *CreateSiteMonitorRequest
	GetVpcConfig() *string
}

type CreateSiteMonitorRequest struct {
	// The URL or IP address that is monitored by the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
	Address    *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// The ID of the alert rule.
	//
	// For more information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// SystemDefault_acs_ecs_dashboard_InternetOutRate_Percent
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// The custom detection period. You can only select a time period from Monday to Sunday for detection.
	//
	// example:
	//
	// {"start_hour":0,"end_hour":24, "days":[0], "time_zone":"Local"}
	CustomSchedule *string `json:"CustomSchedule,omitempty" xml:"CustomSchedule,omitempty"`
	// The interval at which detection requests are sent.
	//
	// Valid values: 1, 5, 15, 30, and 60. Unit: minutes.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The information of the detection points. If you leave this parameter empty, the system randomly selects three detection points.
	//
	// The value is a JSON array. Example: `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]`. The values of the city field indicate Beijing, Hangzhou, and Qingdao.
	//
	// For information about how to obtain detection points, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html).
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCities *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	// The extended options of the protocol that is used by the site monitoring task. The options vary based on the protocol.
	//
	// example:
	//
	// {"time_out":5000}
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the site monitoring task.
	//
	// The name must be 4 to 100 characters in length, and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// HanZhou_ECS1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The protocol that is used by the site monitoring task.
	//
	// Valid values: HTTP, HTTPS, PING, TCP, UDP, DNS, SMTP, POP3, FTP, and WEBSOCKET.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTPS
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// {"vpcId": "vpc-xxxxxx", "vswitchId": "vsw-xxxxxx", "securityGroupId": "sg-xxxxxx", "region": "cn-beijing"}
	VpcConfig *string `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty"`
}

func (s CreateSiteMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateSiteMonitorRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateSiteMonitorRequest) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *CreateSiteMonitorRequest) GetAlertIds() *string {
	return s.AlertIds
}

func (s *CreateSiteMonitorRequest) GetCustomSchedule() *string {
	return s.CustomSchedule
}

func (s *CreateSiteMonitorRequest) GetInterval() *string {
	return s.Interval
}

func (s *CreateSiteMonitorRequest) GetIspCities() *string {
	return s.IspCities
}

func (s *CreateSiteMonitorRequest) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *CreateSiteMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSiteMonitorRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSiteMonitorRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateSiteMonitorRequest) GetVpcConfig() *string {
	return s.VpcConfig
}

func (s *CreateSiteMonitorRequest) SetAddress(v string) *CreateSiteMonitorRequest {
	s.Address = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetAgentGroup(v string) *CreateSiteMonitorRequest {
	s.AgentGroup = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetAlertIds(v string) *CreateSiteMonitorRequest {
	s.AlertIds = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetCustomSchedule(v string) *CreateSiteMonitorRequest {
	s.CustomSchedule = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetInterval(v string) *CreateSiteMonitorRequest {
	s.Interval = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetIspCities(v string) *CreateSiteMonitorRequest {
	s.IspCities = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetOptionsJson(v string) *CreateSiteMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetRegionId(v string) *CreateSiteMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetTaskName(v string) *CreateSiteMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetTaskType(v string) *CreateSiteMonitorRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSiteMonitorRequest) SetVpcConfig(v string) *CreateSiteMonitorRequest {
	s.VpcConfig = &v
	return s
}

func (s *CreateSiteMonitorRequest) Validate() error {
	return dara.Validate(s)
}
