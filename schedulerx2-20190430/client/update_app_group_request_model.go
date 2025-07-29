// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v int32) *UpdateAppGroupRequest
	GetAppVersion() *int32
	SetDescription(v string) *UpdateAppGroupRequest
	GetDescription() *string
	SetGroupId(v string) *UpdateAppGroupRequest
	GetGroupId() *string
	SetMaxConcurrency(v int32) *UpdateAppGroupRequest
	GetMaxConcurrency() *int32
	SetMonitorConfigJson(v string) *UpdateAppGroupRequest
	GetMonitorConfigJson() *string
	SetMonitorContactsJson(v string) *UpdateAppGroupRequest
	GetMonitorContactsJson() *string
	SetNamespace(v string) *UpdateAppGroupRequest
	GetNamespace() *string
	SetRegionId(v string) *UpdateAppGroupRequest
	GetRegionId() *string
}

type UpdateAppGroupRequest struct {
	// The application version. 1: Basic version, 2: Professional version.
	//
	// example:
	//
	// 2
	AppVersion *int32 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the application. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of concurrent instances. Default value: 1. A value of 1 specifies that if the last triggered instance is running, the next instance is not triggered even if the scheduled point in time for running the next instance is reached.
	//
	// example:
	//
	// 1
	MaxConcurrency *int32 `json:"MaxConcurrency,omitempty" xml:"MaxConcurrency,omitempty"`
	// The configuration of the alert. The value is a JSON string. For more information about this parameter, see **Additional information about request parameters**.
	//
	// example:
	//
	// {
	//
	//     "sendChannel": "ding,sms,mail,phone",
	//
	//     "alarmType": "Contacts",
	//
	//     "webhookIsAtAll": false
	//
	// }
	MonitorConfigJson *string `json:"MonitorConfigJson,omitempty" xml:"MonitorConfigJson,omitempty"`
	// The configuration of alert contacts. The value is a JSON string.
	//
	// example:
	//
	// [{"userName":"Tom","userPhone":"89756******"},{"userName":"Bob","ding":"http://www.example.com"}]
	MonitorContactsJson *string `json:"MonitorContactsJson,omitempty" xml:"MonitorContactsJson,omitempty"`
	// The ID of the namespace. You can obtain the ID of the namespace on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppGroupRequest) GetAppVersion() *int32 {
	return s.AppVersion
}

func (s *UpdateAppGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateAppGroupRequest) GetMaxConcurrency() *int32 {
	return s.MaxConcurrency
}

func (s *UpdateAppGroupRequest) GetMonitorConfigJson() *string {
	return s.MonitorConfigJson
}

func (s *UpdateAppGroupRequest) GetMonitorContactsJson() *string {
	return s.MonitorContactsJson
}

func (s *UpdateAppGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateAppGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAppGroupRequest) SetAppVersion(v int32) *UpdateAppGroupRequest {
	s.AppVersion = &v
	return s
}

func (s *UpdateAppGroupRequest) SetDescription(v string) *UpdateAppGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppGroupRequest) SetGroupId(v string) *UpdateAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateAppGroupRequest) SetMaxConcurrency(v int32) *UpdateAppGroupRequest {
	s.MaxConcurrency = &v
	return s
}

func (s *UpdateAppGroupRequest) SetMonitorConfigJson(v string) *UpdateAppGroupRequest {
	s.MonitorConfigJson = &v
	return s
}

func (s *UpdateAppGroupRequest) SetMonitorContactsJson(v string) *UpdateAppGroupRequest {
	s.MonitorContactsJson = &v
	return s
}

func (s *UpdateAppGroupRequest) SetNamespace(v string) *UpdateAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateAppGroupRequest) SetRegionId(v string) *UpdateAppGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
