// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v *SearchCloudGtmAddressesResponseBodyAddresses) *SearchCloudGtmAddressesResponseBody
	GetAddresses() *SearchCloudGtmAddressesResponseBodyAddresses
	SetPageNumber(v int32) *SearchCloudGtmAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCloudGtmAddressesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchCloudGtmAddressesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCloudGtmAddressesResponseBody
	GetTotalPages() *int32
}

type SearchCloudGtmAddressesResponseBody struct {
	// The addresses.
	Addresses *SearchCloudGtmAddressesResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// Current page number, starting from **1**, default is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of **100*	- and a default of **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of address entries that meet the query conditions.
	//
	// example:
	//
	// 15
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchCloudGtmAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBody) GetAddresses() *SearchCloudGtmAddressesResponseBodyAddresses {
	return s.Addresses
}

func (s *SearchCloudGtmAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCloudGtmAddressesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCloudGtmAddressesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCloudGtmAddressesResponseBody) SetAddresses(v *SearchCloudGtmAddressesResponseBodyAddresses) *SearchCloudGtmAddressesResponseBody {
	s.Addresses = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetPageNumber(v int32) *SearchCloudGtmAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetPageSize(v int32) *SearchCloudGtmAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetRequestId(v string) *SearchCloudGtmAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetTotalItems(v int32) *SearchCloudGtmAddressesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) SetTotalPages(v int32) *SearchCloudGtmAddressesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressesResponseBodyAddresses struct {
	Address []*SearchCloudGtmAddressesResponseBodyAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressesResponseBodyAddresses) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddresses) GetAddress() []*SearchCloudGtmAddressesResponseBodyAddressesAddress {
	return s.Address
}

func (s *SearchCloudGtmAddressesResponseBodyAddresses) SetAddress(v []*SearchCloudGtmAddressesResponseBodyAddressesAddress) *SearchCloudGtmAddressesResponseBodyAddresses {
	s.Address = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddresses) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressesResponseBodyAddressesAddress struct {
	// IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// ID of the address, unique identifier for the address.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Address ownership information, not supported in the current version.
	//
	// example:
	//
	// The parameter should be none.
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
	// The availability state of the address when AvailableMode is set to manual. Valid values:
	//
	// 	- available: The address is normal. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is abnormal. In this state, the address is not returned for DNS requests even if the address passes health checks.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Creation time of the address.
	//
	// example:
	//
	// 2024-03-23T13:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creation time of the address (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Current activation status of the address:
	//
	// - enable: Enabled status
	//
	// - disable: Disabled status
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
	// 	- ok_alert: The address fails some health checks of the referenced health check templates, but the address is deemed available.
	//
	// 	- ok_no_monitor: The address does not reference any health check template.
	//
	// 	- exceptional: The address fails some or all health checks of the referenced health check templates, and the address is deemed unavailable.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The health check tasks.
	HealthTasks *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
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
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Address type:
	//
	// - IPv4: ipv4 address - IPv6: ipv6 address - domain: domain name
	//
	// example:
	//
	// IPv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The last modification time of the address.
	//
	// example:
	//
	// 2024-03-29T13:20Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modification time of the address (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddress) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetHealthTasks() *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks {
	return s.HealthTasks
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetName() *string {
	return s.Name
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetType() *string {
	return s.Type
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAddress(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Address = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAddressId(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAttributeInfo(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAvailableMode(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetAvailableStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetCreateTime(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetCreateTimestamp(v int64) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetEnableStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetHealthJudgement(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetHealthStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetHealthTasks(v *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.HealthTasks = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetManualAvailableStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetName(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Name = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetRemark(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetType(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.Type = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetUpdateTime(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) SetUpdateTimestamp(v int64) *SearchCloudGtmAddressesResponseBodyAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddress) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks struct {
	HealthTask []*SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) GetHealthTask() []*SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	return s.HealthTask
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) SetHealthTask(v []*SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks {
	s.HealthTask = v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasks) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask struct {
	// The state of the health check task. Valid values:
	//
	// 	- ok: The task is normal.
	//
	// 	- alert: The task has an alert.
	//
	// 	- no_data: No data is available. In most cases, the health check task is newly created and no data is collected.
	//
	// example:
	//
	// ok
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The target service port for health check probes. When the health check protocol is set to Ping, configuration of the service port is not supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template associated with the address.
	//
	// example:
	//
	// mtp-895180524251002880
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Health check template name.
	//
	// example:
	//
	// IPv4-Ping
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetMonitorStatus(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.MonitorStatus = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetPort(v int32) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateId(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateName(v string) *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *SearchCloudGtmAddressesResponseBodyAddressesAddressHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}
