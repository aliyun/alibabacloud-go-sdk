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
	// The URL or IP address of the monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The detection point group used by the monitoring task. Different options correspond to different detection point types. PC corresponds to PC clients, MOBIEL corresponds to mobile clients, and FC corresponds to VPC internal network detection. If this parameter is not specified, the default value is PC.
	//
	// example:
	//
	// PC
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// The ID of the alert rule.
	//
	// For more information about how to obtain the ID of an existing alert rule in CloudMonitor, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// SystemDefault_acs_ecs_dashboard_InternetOutRate_Percent
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// The custom detection period. You can only select a specific time period from Monday to Sunday for detection.
	//
	// example:
	//
	// {"start_hour":0,"end_hour":24, "days":[0], "time_zone":"Local"}
	CustomSchedule *string `json:"CustomSchedule,omitempty" xml:"CustomSchedule,omitempty"`
	// The monitoring frequency.
	//
	// Valid values: 1 minute, 5 minutes, 15 minutes, 30 minutes, and 60 minutes.
	//
	// Default value: 1 minute.
	//
	// example:
	//
	// 1
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The probe information. If this parameter is left empty, the system randomly selects three detection points.
	//
	// The value is a JSON array. Example: `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]`, which correspond to Beijing, Hangzhou, and Qingdao.
	//
	// For more information about how to obtain detection point information, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html).
	//
	// example:
	//
	// [{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]
	IspCities *string `json:"IspCities,omitempty" xml:"IspCities,omitempty"`
	// The advanced extension options for the protocol type of the monitoring task. Different protocol types correspond to different extension options.
	//
	// example:
	//
	// {"time_out":5000}
	OptionsJson *string `json:"OptionsJson,omitempty" xml:"OptionsJson,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the monitoring task.
	//
	// The name must be 4 to 100 characters in length and can contain letters, digits, underscores (_), and Chinese characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// HanZhou_ECS1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Site monitoring type of site monitoring task.
	//
	// Valid values: HTTP(S), PING, TCP, UDP, DNS, SMTP, POP3, FTP, and WEBSOCKET.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The network information required when creating a VPC internal network monitoring task. You must provide a JSON string that contains the following four fields: vpcId, vswitchId, securityGroupId, and region.
	//
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
