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
	// The IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The unique ID of the address.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The attribution information of the address.
	//
	// example:
	//
	// 当前版本不支持此参数，不会返回地址归属信息。
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The switchover mode for the address when a health check detects an exception:
	//
	// - auto: Automatic mode. The system determines whether to stop or resume DNS resolution for the address based on health check results. DNS resolution is stopped if the address is abnormal and is resumed if the address becomes normal.
	//
	// - manual: Manual mode. You manually control the address status. If the address is set to abnormal, DNS resolution is stopped and is not resumed even if the health check result is normal. If the address is set to normal, DNS resolution is performed. An alert is triggered but DNS resolution is not stopped if a health check detects an exception.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// The availability status of the address:
	//
	// - available: The address is available.
	//
	// - unavailable: The address is unavailable.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The time when the address was created.
	//
	// example:
	//
	// 2024-03-23T13:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UNIX timestamp when the address was created.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The enabled status of the address:
	//
	// enable: The address is enabled.
	//
	// disable: The address is disabled.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health determination condition for the address:
	//
	// - any_ok: At least one health check probe is normal.
	//
	// - p30_ok: At least 30% of health check probes are normal.
	//
	// - p50_ok: At least 50% of health check probes are normal.
	//
	// - p70_ok: At least 70% of health check probes are normal.
	//
	// - all_ok: All health check probes are normal.
	//
	// example:
	//
	// p50_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check result of the address:
	//
	// - ok: All health check tasks that are associated with the address are normal.
	//
	// - ok_alert: Some health check tasks that are associated with the address are abnormal, but the address is still considered normal.
	//
	// - ok_no_monitor: The address is not associated with any health check tasks.
	//
	// - exceptional: Some or all health check tasks that are associated with the address are abnormal, and the address is considered abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string                                         `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HealthTasks  *DescribeCloudGtmAddressResponseBodyHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	// The availability status of the address that is set when the switchover mode is manual:
	//
	// - available: The address is available. DNS resolution is performed for the address. If a health check detects an exception, an alert is triggered but DNS resolution is not stopped.
	//
	// - unavailable: The address is unavailable. DNS resolution is stopped for the address and is not resumed even if the health check result is normal.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	// The name of the address.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the address. Valid values:
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
	// The time when the address configuration was last modified.
	//
	// example:
	//
	// 2024-03-29T13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The UNIX timestamp when the address was last modified.
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
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	TemplateId    *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName  *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
