// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySiteMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *ModifySiteMonitorRequest
	GetAddress() *string
	SetAlertIds(v string) *ModifySiteMonitorRequest
	GetAlertIds() *string
	SetCustomSchedule(v string) *ModifySiteMonitorRequest
	GetCustomSchedule() *string
	SetInterval(v string) *ModifySiteMonitorRequest
	GetInterval() *string
	SetIntervalUnit(v string) *ModifySiteMonitorRequest
	GetIntervalUnit() *string
	SetIspCities(v string) *ModifySiteMonitorRequest
	GetIspCities() *string
	SetOptionsJson(v string) *ModifySiteMonitorRequest
	GetOptionsJson() *string
	SetRegionId(v string) *ModifySiteMonitorRequest
	GetRegionId() *string
	SetTaskId(v string) *ModifySiteMonitorRequest
	GetTaskId() *string
	SetTaskName(v string) *ModifySiteMonitorRequest
	GetTaskName() *string
}

type ModifySiteMonitorRequest struct {
	// The URL or IP address that is monitored by the task.
	//
	// example:
	//
	// http://www.aliyun.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the alert rule. You can call the DescribeMetricRuleList operation to query the IDs of existing alert rules in CloudMonitor. For more information, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// example:
	//
	// 49f7c317-7645-4cc9-94fd-ea42e122****
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	// The custom detection period. You can only select a time period from Monday to Sunday for detection.
	//
	// example:
	//
	// {"start_hour":0,"end_hour":24, "days":[0], "time_zone":"Local"}
	CustomSchedule *string `json:"CustomSchedule,omitempty" xml:"CustomSchedule,omitempty"`
	// The interval at which detection requests are sent. Valid values: 1, 5, and 15. Unit: minutes. Default value: 1.
	//
	// example:
	//
	// 1
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The unit of the metric.
	//
	// Unit: milliseconds.
	//
	// example:
	//
	// ms
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// The information of the detection points. The value is a JSON array. Example: `[{"city":"546","isp":"465"},{"city":"572","isp":"465"},{"city":"738","isp":"465"}]`. The values of the `city` field indicate Beijing, Hangzhou, and Qingdao.
	//
	// > You can call the DescribeSiteMonitorISPCityList operation to query the detection points. For more information, see [DescribeSiteMonitorISPCityList](https://help.aliyun.com/document_detail/115045.html). If you leave this parameter empty, the system randomly selects three detection points.
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
	// The ID of the site monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2c8dbdf9-a3ab-46a1-85a4-f094965e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the site monitoring task. The name must be 4 to 100 characters in length, and can contain letters, digits, and underscores (_).
	//
	// example:
	//
	// HanZhou_ECS2
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ModifySiteMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySiteMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifySiteMonitorRequest) GetAddress() *string {
	return s.Address
}

func (s *ModifySiteMonitorRequest) GetAlertIds() *string {
	return s.AlertIds
}

func (s *ModifySiteMonitorRequest) GetCustomSchedule() *string {
	return s.CustomSchedule
}

func (s *ModifySiteMonitorRequest) GetInterval() *string {
	return s.Interval
}

func (s *ModifySiteMonitorRequest) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *ModifySiteMonitorRequest) GetIspCities() *string {
	return s.IspCities
}

func (s *ModifySiteMonitorRequest) GetOptionsJson() *string {
	return s.OptionsJson
}

func (s *ModifySiteMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySiteMonitorRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifySiteMonitorRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifySiteMonitorRequest) SetAddress(v string) *ModifySiteMonitorRequest {
	s.Address = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetAlertIds(v string) *ModifySiteMonitorRequest {
	s.AlertIds = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetCustomSchedule(v string) *ModifySiteMonitorRequest {
	s.CustomSchedule = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetInterval(v string) *ModifySiteMonitorRequest {
	s.Interval = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetIntervalUnit(v string) *ModifySiteMonitorRequest {
	s.IntervalUnit = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetIspCities(v string) *ModifySiteMonitorRequest {
	s.IspCities = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetOptionsJson(v string) *ModifySiteMonitorRequest {
	s.OptionsJson = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetRegionId(v string) *ModifySiteMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetTaskId(v string) *ModifySiteMonitorRequest {
	s.TaskId = &v
	return s
}

func (s *ModifySiteMonitorRequest) SetTaskName(v string) *ModifySiteMonitorRequest {
	s.TaskName = &v
	return s
}

func (s *ModifySiteMonitorRequest) Validate() error {
	return dara.Validate(s)
}
