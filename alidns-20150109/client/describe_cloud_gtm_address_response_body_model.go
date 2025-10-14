// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeCloudGtmAddressResponseBody
	GetAddress() *string
	SetAddressId(v string) *DescribeCloudGtmAddressResponseBody
	GetAddressId() *string
	SetAttributeInfo(v string) *DescribeCloudGtmAddressResponseBody
	GetAttributeInfo() *string
	SetAvailableMode(v string) *DescribeCloudGtmAddressResponseBody
	GetAvailableMode() *string
	SetAvailableStatus(v string) *DescribeCloudGtmAddressResponseBody
	GetAvailableStatus() *string
	SetCreateTime(v string) *DescribeCloudGtmAddressResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeCloudGtmAddressResponseBody
	GetCreateTimestamp() *int64
	SetEnableStatus(v string) *DescribeCloudGtmAddressResponseBody
	GetEnableStatus() *string
	SetHealthJudgement(v string) *DescribeCloudGtmAddressResponseBody
	GetHealthJudgement() *string
	SetHealthStatus(v string) *DescribeCloudGtmAddressResponseBody
	GetHealthStatus() *string
	SetHealthTasks(v *DescribeCloudGtmAddressResponseBodyHealthTasks) *DescribeCloudGtmAddressResponseBody
	GetHealthTasks() *DescribeCloudGtmAddressResponseBodyHealthTasks
	SetManualAvailableStatus(v string) *DescribeCloudGtmAddressResponseBody
	GetManualAvailableStatus() *string
	SetName(v string) *DescribeCloudGtmAddressResponseBody
	GetName() *string
	SetRemark(v string) *DescribeCloudGtmAddressResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeCloudGtmAddressResponseBody
	GetRequestId() *string
	SetType(v string) *DescribeCloudGtmAddressResponseBody
	GetType() *string
	SetUpdateTime(v string) *DescribeCloudGtmAddressResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeCloudGtmAddressResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeCloudGtmAddressResponseBody struct {
	// IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The address ID. This ID uniquely identifies the address.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Address ownership information.
	//
	// example:
	//
	// The current version does not support passing this parameter, please do not input the parameter.
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The failover method that is used if the address fails health checks. Valid values:
	//
	// 	- auto: the automatic mode. The system determines whether to return an address based on the health check results. If the address fails health checks, the system does not return the address. If the address passes health checks, the system returns the address.
	//
	// 	- manual: the manual mode. If an address is in the unavailable state, the address is not returned for Domain Name System (DNS) requests even if the address passes health checks. If an address is in the available state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// Address availability status:
	//
	// - available: Available
	//
	// - unavailable: Unavailable
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Address creation time.
	//
	// example:
	//
	// 2024-03-23T13:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creation time (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Indicates the current enabled status of the address:
	//
	// enabled: enabled state
	//
	// disabled: disabled state
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health status of the address. Valid values:
	//
	// 	- any_ok: The health check results of at least one health check template are normal.
	//
	// 	- p30_ok: The health check results of at least 30% of health check templates are normal.
	//
	// 	- p50_ok: The health check results of at least 50% of health check templates are normal.
	//
	// 	- p70_ok: The health check results of at least 70% of health check templates are normal.
	//
	// 	- all_ok: The health check results of all health check templates are normal.
	//
	// example:
	//
	// p50_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check state of the address. Valid values:
	//
	// 	- ok: The address passes all health checks of the referenced health check templates.
	//
	// 	- ok_alert: The address fails some health checks of the referenced health check templates but the address is deemed normal.
	//
	// 	- ok_no_monitor: The address does not reference a health check template.
	//
	// 	- exceptional: The address fails some or all health checks of the referenced health check templates and the address is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The health check tasks referenced by the address.
	HealthTasks *DescribeCloudGtmAddressResponseBodyHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	// The availability state of the address when AvailableMode is set to manual. Valid values:
	//
	// 	- available: The address is normal. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is abnormal. In this state, the address is not returned for DNS requests even if the address passes health checks.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	// Address name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Remarks.
	//
	// example:
	//
	// test1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Address type:
	//
	// - IPv4
	//
	// - IPv6
	//
	// - domain
	//
	// example:
	//
	// IPv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The last modification time of the address configuration.
	//
	// example:
	//
	// 2024-03-29T13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Modified time (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeCloudGtmAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressResponseBody) GetAddress() *string {
	return s.Address
}

func (s *DescribeCloudGtmAddressResponseBody) GetAddressId() *string {
	return s.AddressId
}

func (s *DescribeCloudGtmAddressResponseBody) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *DescribeCloudGtmAddressResponseBody) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *DescribeCloudGtmAddressResponseBody) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmAddressResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmAddressResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmAddressResponseBody) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmAddressResponseBody) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *DescribeCloudGtmAddressResponseBody) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmAddressResponseBody) GetHealthTasks() *DescribeCloudGtmAddressResponseBodyHealthTasks {
	return s.HealthTasks
}

func (s *DescribeCloudGtmAddressResponseBody) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *DescribeCloudGtmAddressResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCloudGtmAddressResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmAddressResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeCloudGtmAddressResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmAddressResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmAddressResponseBody) SetAddress(v string) *DescribeCloudGtmAddressResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetAddressId(v string) *DescribeCloudGtmAddressResponseBody {
	s.AddressId = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetAttributeInfo(v string) *DescribeCloudGtmAddressResponseBody {
	s.AttributeInfo = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetAvailableMode(v string) *DescribeCloudGtmAddressResponseBody {
	s.AvailableMode = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetAvailableStatus(v string) *DescribeCloudGtmAddressResponseBody {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetCreateTime(v string) *DescribeCloudGtmAddressResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetCreateTimestamp(v int64) *DescribeCloudGtmAddressResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetEnableStatus(v string) *DescribeCloudGtmAddressResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetHealthJudgement(v string) *DescribeCloudGtmAddressResponseBody {
	s.HealthJudgement = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetHealthStatus(v string) *DescribeCloudGtmAddressResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetHealthTasks(v *DescribeCloudGtmAddressResponseBodyHealthTasks) *DescribeCloudGtmAddressResponseBody {
	s.HealthTasks = v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetManualAvailableStatus(v string) *DescribeCloudGtmAddressResponseBody {
	s.ManualAvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetName(v string) *DescribeCloudGtmAddressResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetRemark(v string) *DescribeCloudGtmAddressResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetRequestId(v string) *DescribeCloudGtmAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetType(v string) *DescribeCloudGtmAddressResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetUpdateTime(v string) *DescribeCloudGtmAddressResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) SetUpdateTimestamp(v int64) *DescribeCloudGtmAddressResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBody) Validate() error {
	if s.HealthTasks != nil {
		if err := s.HealthTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmAddressResponseBodyHealthTasks struct {
	HealthTask []*DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressResponseBodyHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressResponseBodyHealthTasks) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasks) GetHealthTask() []*DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask {
	return s.HealthTask
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasks) SetHealthTask(v []*DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) *DescribeCloudGtmAddressResponseBodyHealthTasks {
	s.HealthTask = v
	return s
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasks) Validate() error {
	if s.HealthTask != nil {
		for _, item := range s.HealthTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask struct {
	// The state of the health check task. Valid values:
	//
	// 	- ok: The task is normal.
	//
	// 	- alert: An alert is triggered.
	//
	// 	- no_data: No data is available. In most cases, the health check task is newly created and no data is collected.
	//
	// example:
	//
	// ok
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The target service port for health checks. When the Ping protocol is selected for health checks, configuration of the service port is not supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template associated with the address.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the health check template.
	//
	// example:
	//
	// Ping-IPv4
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) SetMonitorStatus(v string) *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) SetPort(v int32) *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) SetTemplateId(v string) *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) SetTemplateName(v string) *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *DescribeCloudGtmAddressResponseBodyHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}
