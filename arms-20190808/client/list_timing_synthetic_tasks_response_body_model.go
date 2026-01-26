// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimingSyntheticTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListTimingSyntheticTasksResponseBody
	GetCode() *int64
	SetData(v *ListTimingSyntheticTasksResponseBodyData) *ListTimingSyntheticTasksResponseBody
	GetData() *ListTimingSyntheticTasksResponseBodyData
	SetMessage(v string) *ListTimingSyntheticTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTimingSyntheticTasksResponseBody
	GetRequestId() *string
}

type ListTimingSyntheticTasksResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *ListTimingSyntheticTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 730E90FE-996A-5638-99F3-4F0F9038CC6C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListTimingSyntheticTasksResponseBody) GetData() *ListTimingSyntheticTasksResponseBodyData {
	return s.Data
}

func (s *ListTimingSyntheticTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTimingSyntheticTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTimingSyntheticTasksResponseBody) SetCode(v int64) *ListTimingSyntheticTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBody) SetData(v *ListTimingSyntheticTasksResponseBodyData) *ListTimingSyntheticTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBody) SetMessage(v string) *ListTimingSyntheticTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBody) SetRequestId(v string) *ListTimingSyntheticTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTimingSyntheticTasksResponseBodyData struct {
	// The queried tasks.
	Items []*ListTimingSyntheticTasksResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyData) GetItems() []*ListTimingSyntheticTasksResponseBodyDataItems {
	return s.Items
}

func (s *ListTimingSyntheticTasksResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListTimingSyntheticTasksResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTimingSyntheticTasksResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListTimingSyntheticTasksResponseBodyData) SetItems(v []*ListTimingSyntheticTasksResponseBodyDataItems) *ListTimingSyntheticTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyData) SetPage(v int32) *ListTimingSyntheticTasksResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyData) SetPageSize(v int32) *ListTimingSyntheticTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyData) SetTotal(v int32) *ListTimingSyntheticTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTimingSyntheticTasksResponseBodyDataItems struct {
	// The general settings.
	CommonSetting *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting `json:"CommonSetting,omitempty" xml:"CommonSetting,omitempty" type:"Struct"`
	// The detection frequency. Valid values: 1m, 5m, 10m, 15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, and 24h.
	//
	// example:
	//
	// 1m
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1671454758000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was modified.
	//
	// example:
	//
	// 1673085633000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The detection point type. 1: PC. 2: mobile device.
	//
	// example:
	//
	// 1
	MonitorCategory *int32 `json:"MonitorCategory,omitempty" xml:"MonitorCategory,omitempty"`
	// The number of detection points.
	//
	// example:
	//
	// 10
	MonitorNum *string `json:"MonitorNum,omitempty" xml:"MonitorNum,omitempty"`
	// The task name.
	//
	// example:
	//
	// AlibabaCloud DNS Task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzgwtq5vxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The task status. CREATING: The task is being created. RUNNING: The task is running. PARTIAL_RUNNING: The task is partially running. STOP: The task is stopped. LIMIT_STOP: The task is stopped due to quota limit. EXCEPTION: The task is abnormal. DELETE: The task is deleted. DELETE_EXCEPTION: An exception occurs while deleting the task.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListTimingSyntheticTasksResponseBodyDataItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 5308a2691f59422c8c3b7aeccec9cd3b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task. Valid values:
	//
	// 1: ICMP. 2: TCP. 3: DNS. 4: HTTP. 5: website speed. 6: file download.
	//
	// example:
	//
	// 1
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The URL for synthetic monitoring.
	//
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetCommonSetting() *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	return s.CommonSetting
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetFrequency() *string {
	return s.Frequency
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetMonitorCategory() *int32 {
	return s.MonitorCategory
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetMonitorNum() *string {
	return s.MonitorNum
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetTags() []*ListTimingSyntheticTasksResponseBodyDataItemsTags {
	return s.Tags
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetTaskType() *int32 {
	return s.TaskType
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) GetUrl() *string {
	return s.Url
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetCommonSetting(v *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.CommonSetting = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetFrequency(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.Frequency = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetGmtCreate(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetGmtModified(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.GmtModified = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetMonitorCategory(v int32) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.MonitorCategory = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetMonitorNum(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.MonitorNum = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetName(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetRegionId(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.RegionId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetResourceGroupId(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetStatus(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetTags(v []*ListTimingSyntheticTasksResponseBodyDataItemsTags) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.Tags = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetTaskId(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.TaskId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetTaskType(v int32) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.TaskType = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) SetUrl(v string) *ListTimingSyntheticTasksResponseBodyDataItems {
	s.Url = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItems) Validate() error {
	if s.CommonSetting != nil {
		if err := s.CommonSetting.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting struct {
	// The custom host settings.
	CustomHost *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost `json:"CustomHost,omitempty" xml:"CustomHost,omitempty" type:"Struct"`
	// A reserved field.
	CustomPrometheusSetting *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting `json:"CustomPrometheusSetting,omitempty" xml:"CustomPrometheusSetting,omitempty" type:"Struct"`
	// The information about the virtual private cloud (VPC). If the destination URL is an Alibaba Cloud internal endpoint, you need to configure a VPC.
	CustomVPCSetting *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting `json:"CustomVPCSetting,omitempty" xml:"CustomVPCSetting,omitempty" type:"Struct"`
	// The IP version. Valid values:
	//
	// 	- 0: A version is automatically selected.
	//
	// 	- 1: IPv4.
	//
	// 	- 2: IPv6.
	//
	// example:
	//
	// 0
	IpType *int32 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// Indicates whether tracing is enabled.
	//
	// example:
	//
	// false
	IsOpenTrace *bool `json:"IsOpenTrace,omitempty" xml:"IsOpenTrace,omitempty"`
	// Indicates whether monitoring samples are evenly distributed. Valid values:
	//
	// 	- 0: No
	//
	// 	- 1: Yes
	//
	// example:
	//
	// 0
	MonitorSamples *int32 `json:"MonitorSamples,omitempty" xml:"MonitorSamples,omitempty"`
	// The type of the client for tracing. Valid values:
	//
	// 	- 0: ARMS agent
	//
	// 	- 1: OpenTelemetry
	//
	// 	- 2: Jaeger
	//
	// example:
	//
	// 1
	TraceClientType *int32 `json:"TraceClientType,omitempty" xml:"TraceClientType,omitempty"`
	// The region to which trace data is reported.
	//
	// example:
	//
	// cn-hangzhou
	XtraceRegion *string `json:"XtraceRegion,omitempty" xml:"XtraceRegion,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetCustomHost() *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost {
	return s.CustomHost
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetCustomPrometheusSetting() *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting {
	return s.CustomPrometheusSetting
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetCustomVPCSetting() *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting {
	return s.CustomVPCSetting
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetIpType() *int32 {
	return s.IpType
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetIsOpenTrace() *bool {
	return s.IsOpenTrace
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetMonitorSamples() *int32 {
	return s.MonitorSamples
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetTraceClientType() *int32 {
	return s.TraceClientType
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) GetXtraceRegion() *string {
	return s.XtraceRegion
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetCustomHost(v *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.CustomHost = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetCustomPrometheusSetting(v *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.CustomPrometheusSetting = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetCustomVPCSetting(v *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.CustomVPCSetting = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetIpType(v int32) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.IpType = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetIsOpenTrace(v bool) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.IsOpenTrace = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetMonitorSamples(v int32) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.MonitorSamples = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetTraceClientType(v int32) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.TraceClientType = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) SetXtraceRegion(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting {
	s.XtraceRegion = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSetting) Validate() error {
	if s.CustomHost != nil {
		if err := s.CustomHost.Validate(); err != nil {
			return err
		}
	}
	if s.CustomPrometheusSetting != nil {
		if err := s.CustomPrometheusSetting.Validate(); err != nil {
			return err
		}
	}
	if s.CustomVPCSetting != nil {
		if err := s.CustomVPCSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost struct {
	// The custom host settings.
	Hosts []*ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The selection mode. Valid values:
	//
	// 	- 0: random
	//
	// 	- 1: polling
	//
	// example:
	//
	// 0
	SelectType *int32 `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) GetHosts() []*ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts {
	return s.Hosts
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) GetSelectType() *int32 {
	return s.SelectType
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) SetHosts(v []*ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost {
	s.Hosts = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) SetSelectType(v int32) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost {
	s.SelectType = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHost) Validate() error {
	if s.Hosts != nil {
		for _, item := range s.Hosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts struct {
	// The destination domain name.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP version. Valid values:
	//
	// 	- 0: A version is automatically selected.
	//
	// 	- 1: IPv4.
	//
	// 	- 2: IPv6.
	//
	// example:
	//
	// 0
	IpType *int32 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The IP address.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) GetDomain() *string {
	return s.Domain
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) GetIpType() *int32 {
	return s.IpType
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) GetIps() []*string {
	return s.Ips
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) SetDomain(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts {
	s.Domain = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) SetIpType(v int32) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts {
	s.IpType = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) SetIps(v []*string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts {
	s.Ips = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomHostHosts) Validate() error {
	return dara.Validate(s)
}

type ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting struct {
	// A reserved field.
	//
	// example:
	//
	// Reserved field
	PrometheusClusterId *string `json:"PrometheusClusterId,omitempty" xml:"PrometheusClusterId,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// Reserved field
	PrometheusClusterRegion *string `json:"PrometheusClusterRegion,omitempty" xml:"PrometheusClusterRegion,omitempty"`
	// A reserved field.
	PrometheusLabels map[string]*string `json:"PrometheusLabels,omitempty" xml:"PrometheusLabels,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) GetPrometheusClusterId() *string {
	return s.PrometheusClusterId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) GetPrometheusClusterRegion() *string {
	return s.PrometheusClusterRegion
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) GetPrometheusLabels() map[string]*string {
	return s.PrometheusLabels
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) SetPrometheusClusterId(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) SetPrometheusClusterRegion(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterRegion = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) SetPrometheusLabels(v map[string]*string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting {
	s.PrometheusLabels = v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomPrometheusSetting) Validate() error {
	return dara.Validate(s)
}

type ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting struct {
	// The region ID.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the client belongs. The security group specifies the inbound and outbound rules of the client for the VPC. You need to allow the security group to which the client belongs to access the security group to which the VPC belongs. Otherwise, the client cannot access resources in the VPC.
	//
	// example:
	//
	// sg-xxxxxxxxxxxxxx
	SecureGroupId *string `json:"SecureGroupId,omitempty" xml:"SecureGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1bcmj81kxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zehbd4dfzahxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) GetSecureGroupId() *string {
	return s.SecureGroupId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) GetVpcId() *string {
	return s.VpcId
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) SetRegionId(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting {
	s.RegionId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) SetSecureGroupId(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting {
	s.SecureGroupId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) SetVSwitchId(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting {
	s.VSwitchId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) SetVpcId(v string) *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting {
	s.VpcId = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsCommonSettingCustomVPCSetting) Validate() error {
	return dara.Validate(s)
}

type ListTimingSyntheticTasksResponseBodyDataItemsTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsTags) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponseBodyDataItemsTags) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsTags) GetKey() *string {
	return s.Key
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsTags) GetValue() *string {
	return s.Value
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsTags) SetKey(v string) *ListTimingSyntheticTasksResponseBodyDataItemsTags {
	s.Key = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsTags) SetValue(v string) *ListTimingSyntheticTasksResponseBodyDataItemsTags {
	s.Value = &v
	return s
}

func (s *ListTimingSyntheticTasksResponseBodyDataItemsTags) Validate() error {
	return dara.Validate(s)
}
