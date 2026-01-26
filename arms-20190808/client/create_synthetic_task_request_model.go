// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonParam(v *CreateSyntheticTaskRequestCommonParam) *CreateSyntheticTaskRequest
	GetCommonParam() *CreateSyntheticTaskRequestCommonParam
	SetDownload(v *CreateSyntheticTaskRequestDownload) *CreateSyntheticTaskRequest
	GetDownload() *CreateSyntheticTaskRequestDownload
	SetExtendInterval(v *CreateSyntheticTaskRequestExtendInterval) *CreateSyntheticTaskRequest
	GetExtendInterval() *CreateSyntheticTaskRequestExtendInterval
	SetIntervalTime(v string) *CreateSyntheticTaskRequest
	GetIntervalTime() *string
	SetIntervalType(v string) *CreateSyntheticTaskRequest
	GetIntervalType() *string
	SetIpType(v int64) *CreateSyntheticTaskRequest
	GetIpType() *int64
	SetMonitorList(v []*CreateSyntheticTaskRequestMonitorList) *CreateSyntheticTaskRequest
	GetMonitorList() []*CreateSyntheticTaskRequestMonitorList
	SetNavigation(v *CreateSyntheticTaskRequestNavigation) *CreateSyntheticTaskRequest
	GetNavigation() *CreateSyntheticTaskRequestNavigation
	SetNet(v *CreateSyntheticTaskRequestNet) *CreateSyntheticTaskRequest
	GetNet() *CreateSyntheticTaskRequestNet
	SetProtocol(v *CreateSyntheticTaskRequestProtocol) *CreateSyntheticTaskRequest
	GetProtocol() *CreateSyntheticTaskRequestProtocol
	SetRegionId(v string) *CreateSyntheticTaskRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateSyntheticTaskRequest
	GetTaskName() *string
	SetTaskType(v int64) *CreateSyntheticTaskRequest
	GetTaskType() *int64
	SetUpdateTask(v bool) *CreateSyntheticTaskRequest
	GetUpdateTask() *bool
	SetUrl(v string) *CreateSyntheticTaskRequest
	GetUrl() *string
}

type CreateSyntheticTaskRequest struct {
	// The common parameters.
	CommonParam *CreateSyntheticTaskRequestCommonParam `json:"CommonParam,omitempty" xml:"CommonParam,omitempty" type:"Struct"`
	// The file download task.
	Download *CreateSyntheticTaskRequestDownload `json:"Download,omitempty" xml:"Download,omitempty" type:"Struct"`
	// The frequency.
	ExtendInterval *CreateSyntheticTaskRequestExtendInterval `json:"ExtendInterval,omitempty" xml:"ExtendInterval,omitempty" type:"Struct"`
	// The interval at which synthetic monitoring is performed. Unit: minutes. Valid values:
	//
	// 	- 1
	//
	// 	- 5
	//
	// 	- 10
	//
	// 	- 15
	//
	// 	- 20
	//
	// 	- 30
	//
	// 	- 60
	//
	// 	- 120
	//
	// 	- 180
	//
	// 	- 240
	//
	// 	- 360
	//
	// 	- 480
	//
	// 	- 720
	//
	// 	- 1440
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	IntervalTime *string `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	// The interval type.
	//
	// 	- 0: daily
	//
	// 	- 1: custom frequency
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IntervalType *string `json:"IntervalType,omitempty" xml:"IntervalType,omitempty"`
	// The IP address type:
	//
	// 	- 0: an automatic IP address
	//
	// 	- 1: IPv4
	//
	// 	- 2: IPv6
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IpType *int64 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The list of monitoring points.
	//
	// This parameter is required.
	MonitorList []*CreateSyntheticTaskRequestMonitorList `json:"MonitorList,omitempty" xml:"MonitorList,omitempty" type:"Repeated"`
	// The monitoring items that are associated with the browse tasks.
	Navigation *CreateSyntheticTaskRequestNavigation `json:"Navigation,omitempty" xml:"Navigation,omitempty" type:"Struct"`
	// The network synthetic monitoring task.
	Net *CreateSyntheticTaskRequestNet `json:"Net,omitempty" xml:"Net,omitempty" type:"Struct"`
	// The API performance synthetic monitoring task.
	Protocol *CreateSyntheticTaskRequestProtocol `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Struct"`
	// The ID of the region in which the application is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the task. To update a synthetic monitoring task, enter the task name and set the **UpdateTask*	- parameter to **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Network synthetic monitoring task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the monitoring task. Valid values:
	//
	// 1.  3: web page performance - IE
	//
	// 2.  34: web Page Performance - Chrome
	//
	// 3.  0: network quality
	//
	// 4.  40: file download
	//
	// 5.  7:API performance
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Specifies whether to update existing synthetic monitoring tasks.
	//
	// 	- true: updates existing synthetic monitoring tasks.
	//
	// 	- false: creates new synthetic monitoring tasks.
	//
	// example:
	//
	// false
	UpdateTask *bool `json:"UpdateTask,omitempty" xml:"UpdateTask,omitempty"`
	// The URL for synthetic monitoring.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequest) GetCommonParam() *CreateSyntheticTaskRequestCommonParam {
	return s.CommonParam
}

func (s *CreateSyntheticTaskRequest) GetDownload() *CreateSyntheticTaskRequestDownload {
	return s.Download
}

func (s *CreateSyntheticTaskRequest) GetExtendInterval() *CreateSyntheticTaskRequestExtendInterval {
	return s.ExtendInterval
}

func (s *CreateSyntheticTaskRequest) GetIntervalTime() *string {
	return s.IntervalTime
}

func (s *CreateSyntheticTaskRequest) GetIntervalType() *string {
	return s.IntervalType
}

func (s *CreateSyntheticTaskRequest) GetIpType() *int64 {
	return s.IpType
}

func (s *CreateSyntheticTaskRequest) GetMonitorList() []*CreateSyntheticTaskRequestMonitorList {
	return s.MonitorList
}

func (s *CreateSyntheticTaskRequest) GetNavigation() *CreateSyntheticTaskRequestNavigation {
	return s.Navigation
}

func (s *CreateSyntheticTaskRequest) GetNet() *CreateSyntheticTaskRequestNet {
	return s.Net
}

func (s *CreateSyntheticTaskRequest) GetProtocol() *CreateSyntheticTaskRequestProtocol {
	return s.Protocol
}

func (s *CreateSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSyntheticTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSyntheticTaskRequest) GetTaskType() *int64 {
	return s.TaskType
}

func (s *CreateSyntheticTaskRequest) GetUpdateTask() *bool {
	return s.UpdateTask
}

func (s *CreateSyntheticTaskRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateSyntheticTaskRequest) SetCommonParam(v *CreateSyntheticTaskRequestCommonParam) *CreateSyntheticTaskRequest {
	s.CommonParam = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetDownload(v *CreateSyntheticTaskRequestDownload) *CreateSyntheticTaskRequest {
	s.Download = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetExtendInterval(v *CreateSyntheticTaskRequestExtendInterval) *CreateSyntheticTaskRequest {
	s.ExtendInterval = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetIntervalTime(v string) *CreateSyntheticTaskRequest {
	s.IntervalTime = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetIntervalType(v string) *CreateSyntheticTaskRequest {
	s.IntervalType = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetIpType(v int64) *CreateSyntheticTaskRequest {
	s.IpType = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetMonitorList(v []*CreateSyntheticTaskRequestMonitorList) *CreateSyntheticTaskRequest {
	s.MonitorList = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetNavigation(v *CreateSyntheticTaskRequestNavigation) *CreateSyntheticTaskRequest {
	s.Navigation = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetNet(v *CreateSyntheticTaskRequestNet) *CreateSyntheticTaskRequest {
	s.Net = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetProtocol(v *CreateSyntheticTaskRequestProtocol) *CreateSyntheticTaskRequest {
	s.Protocol = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetRegionId(v string) *CreateSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetTaskName(v string) *CreateSyntheticTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetTaskType(v int64) *CreateSyntheticTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetUpdateTask(v bool) *CreateSyntheticTaskRequest {
	s.UpdateTask = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetUrl(v string) *CreateSyntheticTaskRequest {
	s.Url = &v
	return s
}

func (s *CreateSyntheticTaskRequest) Validate() error {
	if s.CommonParam != nil {
		if err := s.CommonParam.Validate(); err != nil {
			return err
		}
	}
	if s.Download != nil {
		if err := s.Download.Validate(); err != nil {
			return err
		}
	}
	if s.ExtendInterval != nil {
		if err := s.ExtendInterval.Validate(); err != nil {
			return err
		}
	}
	if s.MonitorList != nil {
		for _, item := range s.MonitorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Navigation != nil {
		if err := s.Navigation.Validate(); err != nil {
			return err
		}
	}
	if s.Net != nil {
		if err := s.Net.Validate(); err != nil {
			return err
		}
	}
	if s.Protocol != nil {
		if err := s.Protocol.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSyntheticTaskRequestCommonParam struct {
	// Specifies whether to create an alert rule.
	//
	// 	- 1: creates an alert.
	//
	// 	- 0: does not create an alert.
	//
	// example:
	//
	// 1
	AlarmFlag *string `json:"AlarmFlag,omitempty" xml:"AlarmFlag,omitempty"`
	// The alert parameters.
	AlertList []*CreateSyntheticTaskRequestCommonParamAlertList `json:"AlertList,omitempty" xml:"AlertList,omitempty" type:"Repeated"`
	// The ID of the alert recipient. Separate multiple recipients with commas (,).
	//
	// example:
	//
	// 123
	AlertNotifierId *string `json:"AlertNotifierId,omitempty" xml:"AlertNotifierId,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// 1234
	AlertPolicyId *string `json:"AlertPolicyId,omitempty" xml:"AlertPolicyId,omitempty"`
	// Specifies whether to evenly distribute monitoring samples. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	MonitorSamples *int64 `json:"MonitorSamples,omitempty" xml:"MonitorSamples,omitempty"`
	// The time when execution starts.
	//
	// example:
	//
	// 2022-07-20 10
	StartExecutionTime *int64 `json:"StartExecutionTime,omitempty" xml:"StartExecutionTime,omitempty"`
}

func (s CreateSyntheticTaskRequestCommonParam) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestCommonParam) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestCommonParam) GetAlarmFlag() *string {
	return s.AlarmFlag
}

func (s *CreateSyntheticTaskRequestCommonParam) GetAlertList() []*CreateSyntheticTaskRequestCommonParamAlertList {
	return s.AlertList
}

func (s *CreateSyntheticTaskRequestCommonParam) GetAlertNotifierId() *string {
	return s.AlertNotifierId
}

func (s *CreateSyntheticTaskRequestCommonParam) GetAlertPolicyId() *string {
	return s.AlertPolicyId
}

func (s *CreateSyntheticTaskRequestCommonParam) GetMonitorSamples() *int64 {
	return s.MonitorSamples
}

func (s *CreateSyntheticTaskRequestCommonParam) GetStartExecutionTime() *int64 {
	return s.StartExecutionTime
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlarmFlag(v string) *CreateSyntheticTaskRequestCommonParam {
	s.AlarmFlag = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlertList(v []*CreateSyntheticTaskRequestCommonParamAlertList) *CreateSyntheticTaskRequestCommonParam {
	s.AlertList = v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlertNotifierId(v string) *CreateSyntheticTaskRequestCommonParam {
	s.AlertNotifierId = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlertPolicyId(v string) *CreateSyntheticTaskRequestCommonParam {
	s.AlertPolicyId = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetMonitorSamples(v int64) *CreateSyntheticTaskRequestCommonParam {
	s.MonitorSamples = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetStartExecutionTime(v int64) *CreateSyntheticTaskRequestCommonParam {
	s.StartExecutionTime = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) Validate() error {
	if s.AlertList != nil {
		for _, item := range s.AlertList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSyntheticTaskRequestCommonParamAlertList struct {
	// Specifies whether the condition must be met.
	//
	// example:
	//
	// true
	IsCritical *int64 `json:"IsCritical,omitempty" xml:"IsCritical,omitempty"`
	// The name of the alert rule.
	//
	// For network synthetic monitoring, use the following names:
	//
	// 	- Latency: PING_SET
	//
	// 	- Packet loss rate: PING_LOST_RATE
	//
	// 	- Hijacking: HIJACKPER
	//
	// example:
	//
	// PING_SET
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies how the condition is evaluated. Valid values:
	//
	// 	- 1: greater than
	//
	// 	- 0: less than
	//
	// example:
	//
	// 1
	Symbols *int64 `json:"Symbols,omitempty" xml:"Symbols,omitempty"`
}

func (s CreateSyntheticTaskRequestCommonParamAlertList) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestCommonParamAlertList) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) GetIsCritical() *int64 {
	return s.IsCritical
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) GetName() *string {
	return s.Name
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) GetSymbols() *int64 {
	return s.Symbols
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) SetIsCritical(v int64) *CreateSyntheticTaskRequestCommonParamAlertList {
	s.IsCritical = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) SetName(v string) *CreateSyntheticTaskRequestCommonParamAlertList {
	s.Name = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) SetSymbols(v int64) *CreateSyntheticTaskRequestCommonParamAlertList {
	s.Symbols = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestDownload struct {
	// The connection timeout period.
	//
	// example:
	//
	// 200
	ConnectionTimeout *float64 `json:"ConnectionTimeout,omitempty" xml:"ConnectionTimeout,omitempty"`
	// The items to be ignored in a certificate error. Pass the values of the check boxes that are separated with vertical bars (|).
	//
	// example:
	//
	// Host:www.example.com|Referer:www.example.com
	DownloadCustomHeaderContent *string `json:"DownloadCustomHeaderContent,omitempty" xml:"DownloadCustomHeaderContent,omitempty"`
	// The custom host mode.
	//
	// 	- 1: round robin
	//
	// 	- 0: random
	//
	// example:
	//
	// 1
	DownloadCustomHost *int64 `json:"DownloadCustomHost,omitempty" xml:"DownloadCustomHost,omitempty"`
	// The custom host IP address. You can enter multiple IP addresses. Separate the IP addresses with commas (,).
	//
	// example:
	//
	// ipv4:192.168.2.1,192.168.2.5:img.a.com|192.168.2.1[8080]:img.a.com
	DownloadCustomHostIp *string `json:"DownloadCustomHostIp,omitempty" xml:"DownloadCustomHostIp,omitempty"`
	// The items to be ignored in a certificate error. Pass the values of the check boxes that are separated with vertical bars (|).
	//
	// example:
	//
	// 1|2|4
	DownloadIgnoreCertificateError *string `json:"DownloadIgnoreCertificateError,omitempty" xml:"DownloadIgnoreCertificateError,omitempty"`
	// The kernel type.
	//
	// 	- 1: curl
	//
	// 	- 0: WinInet
	//
	// example:
	//
	// 1
	DownloadKernel *int64 `json:"DownloadKernel,omitempty" xml:"DownloadKernel,omitempty"`
	// Specifies whether to support redirection.
	//
	// example:
	//
	// 0
	DownloadRedirection *int64 `json:"DownloadRedirection,omitempty" xml:"DownloadRedirection,omitempty"`
	// The size of the download file. Unit: KB.
	//
	// example:
	//
	// 10240
	DownloadTransmissionSize *int64 `json:"DownloadTransmissionSize,omitempty" xml:"DownloadTransmissionSize,omitempty"`
	// The monitoring duration.
	//
	// example:
	//
	// 30
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// The QUIC protocol type.
	//
	// 	- 1: http1
	//
	// 	- 2: http2
	//
	// 	- 3: http3
	//
	// example:
	//
	// 1
	QuickProtocol *string `json:"QuickProtocol,omitempty" xml:"QuickProtocol,omitempty"`
	// The keyword that is used in verification.
	//
	// example:
	//
	// keyword
	ValidateKeywords *string `json:"ValidateKeywords,omitempty" xml:"ValidateKeywords,omitempty"`
	// The verification method.
	//
	// 	- 0: no verification
	//
	// 	- 1: string verification
	//
	// 	- 2: MD5 verification
	//
	// example:
	//
	// 0
	VerifyWay *int64 `json:"VerifyWay,omitempty" xml:"VerifyWay,omitempty"`
	// The whitelist for DNS hijacking.
	//
	// example:
	//
	// [{\\"src\\":\\"211.154.166.174\\"}]
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s CreateSyntheticTaskRequestDownload) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestDownload) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestDownload) GetConnectionTimeout() *float64 {
	return s.ConnectionTimeout
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadCustomHeaderContent() *string {
	return s.DownloadCustomHeaderContent
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadCustomHost() *int64 {
	return s.DownloadCustomHost
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadCustomHostIp() *string {
	return s.DownloadCustomHostIp
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadIgnoreCertificateError() *string {
	return s.DownloadIgnoreCertificateError
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadKernel() *int64 {
	return s.DownloadKernel
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadRedirection() *int64 {
	return s.DownloadRedirection
}

func (s *CreateSyntheticTaskRequestDownload) GetDownloadTransmissionSize() *int64 {
	return s.DownloadTransmissionSize
}

func (s *CreateSyntheticTaskRequestDownload) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *CreateSyntheticTaskRequestDownload) GetQuickProtocol() *string {
	return s.QuickProtocol
}

func (s *CreateSyntheticTaskRequestDownload) GetValidateKeywords() *string {
	return s.ValidateKeywords
}

func (s *CreateSyntheticTaskRequestDownload) GetVerifyWay() *int64 {
	return s.VerifyWay
}

func (s *CreateSyntheticTaskRequestDownload) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateSyntheticTaskRequestDownload) SetConnectionTimeout(v float64) *CreateSyntheticTaskRequestDownload {
	s.ConnectionTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadCustomHeaderContent(v string) *CreateSyntheticTaskRequestDownload {
	s.DownloadCustomHeaderContent = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadCustomHost(v int64) *CreateSyntheticTaskRequestDownload {
	s.DownloadCustomHost = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadCustomHostIp(v string) *CreateSyntheticTaskRequestDownload {
	s.DownloadCustomHostIp = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadIgnoreCertificateError(v string) *CreateSyntheticTaskRequestDownload {
	s.DownloadIgnoreCertificateError = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadKernel(v int64) *CreateSyntheticTaskRequestDownload {
	s.DownloadKernel = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadRedirection(v int64) *CreateSyntheticTaskRequestDownload {
	s.DownloadRedirection = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetDownloadTransmissionSize(v int64) *CreateSyntheticTaskRequestDownload {
	s.DownloadTransmissionSize = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetMonitorTimeout(v int64) *CreateSyntheticTaskRequestDownload {
	s.MonitorTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetQuickProtocol(v string) *CreateSyntheticTaskRequestDownload {
	s.QuickProtocol = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetValidateKeywords(v string) *CreateSyntheticTaskRequestDownload {
	s.ValidateKeywords = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetVerifyWay(v int64) *CreateSyntheticTaskRequestDownload {
	s.VerifyWay = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) SetWhiteList(v string) *CreateSyntheticTaskRequestDownload {
	s.WhiteList = &v
	return s
}

func (s *CreateSyntheticTaskRequestDownload) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestExtendInterval struct {
	// The day on which synthetic monitoring is performed.
	Days []*int64 `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The hour at which synthetic monitoring ends.
	//
	// example:
	//
	// 23
	EndHour *int64 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The minute at which synthetic monitoring ends.
	//
	// example:
	//
	// 00
	EndMinute *int64 `json:"EndMinute,omitempty" xml:"EndMinute,omitempty"`
	// The time when synthetic monitoring ends. The format is `yyyy-MM-dd HH`.
	//
	// example:
	//
	// 2022-08-20 10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The hour at which synthetic monitoring starts.
	//
	// example:
	//
	// 00
	StartHour *int64 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
	// The minute at which synthetic monitoring starts.
	//
	// example:
	//
	// 00
	StartMinute *int64 `json:"StartMinute,omitempty" xml:"StartMinute,omitempty"`
	// The time when synthetic monitoring starts. The format is `yyyy-MM-dd HH`.
	//
	// example:
	//
	// 2022-07-20 10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateSyntheticTaskRequestExtendInterval) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestExtendInterval) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetDays() []*int64 {
	return s.Days
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetEndHour() *int64 {
	return s.EndHour
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetEndMinute() *int64 {
	return s.EndMinute
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetStartHour() *int64 {
	return s.StartHour
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetStartMinute() *int64 {
	return s.StartMinute
}

func (s *CreateSyntheticTaskRequestExtendInterval) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetDays(v []*int64) *CreateSyntheticTaskRequestExtendInterval {
	s.Days = v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetEndHour(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.EndHour = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetEndMinute(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.EndMinute = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetEndTime(v string) *CreateSyntheticTaskRequestExtendInterval {
	s.EndTime = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetStartHour(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.StartHour = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetStartMinute(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.StartMinute = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetStartTime(v string) *CreateSyntheticTaskRequestExtendInterval {
	s.StartTime = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestMonitorList struct {
	// The ID of the city to which the monitoring point belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1100101
	CityCode *int64 `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The carrier type:
	//
	// 	- IDC
	//
	// 	- LastMilie
	//
	// This parameter is required.
	//
	// example:
	//
	// IDC
	MonitorType *int64 `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// The ID of the carrier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18
	NetServiceId *int64 `json:"NetServiceId,omitempty" xml:"NetServiceId,omitempty"`
}

func (s CreateSyntheticTaskRequestMonitorList) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestMonitorList) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestMonitorList) GetCityCode() *int64 {
	return s.CityCode
}

func (s *CreateSyntheticTaskRequestMonitorList) GetMonitorType() *int64 {
	return s.MonitorType
}

func (s *CreateSyntheticTaskRequestMonitorList) GetNetServiceId() *int64 {
	return s.NetServiceId
}

func (s *CreateSyntheticTaskRequestMonitorList) SetCityCode(v int64) *CreateSyntheticTaskRequestMonitorList {
	s.CityCode = &v
	return s
}

func (s *CreateSyntheticTaskRequestMonitorList) SetMonitorType(v int64) *CreateSyntheticTaskRequestMonitorList {
	s.MonitorType = &v
	return s
}

func (s *CreateSyntheticTaskRequestMonitorList) SetNetServiceId(v int64) *CreateSyntheticTaskRequestMonitorList {
	s.NetServiceId = &v
	return s
}

func (s *CreateSyntheticTaskRequestMonitorList) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestNavigation struct {
	// The whitelist for DNS hijacking.
	//
	// example:
	//
	// www.aliyun.com:202.0.3.55|203.3.44.67
	DNSHijackWhiteList *string `json:"DNSHijackWhiteList,omitempty" xml:"DNSHijackWhiteList,omitempty"`
	// The element blacklist.
	//
	// example:
	//
	// *.a.com
	ElementBlacklist *string `json:"ElementBlacklist,omitempty" xml:"ElementBlacklist,omitempty"`
	// Specifies whether to execute ActiveX.
	//
	// 	- 3: yes
	//
	// 	- 0: no
	//
	// >  This parameter is supported only by IE Full Elements.
	//
	// example:
	//
	// 3
	ExecuteActiveX *int64 `json:"ExecuteActiveX,omitempty" xml:"ExecuteActiveX,omitempty"`
	// Specifies whether to run applets.
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// >  This parameter is supported only by IE Full Elements.
	//
	// example:
	//
	// 1
	ExecuteApplication *int64 `json:"ExecuteApplication,omitempty" xml:"ExecuteApplication,omitempty"`
	// Specifies whether to execute scripts.
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// >  This parameter is supported only by IE Full Elements.
	//
	// example:
	//
	// 1
	ExecuteScript *int64 `json:"ExecuteScript,omitempty" xml:"ExecuteScript,omitempty"`
	// Specifies whether to filter invalid IP addresses.
	//
	// 	- 1: no
	//
	// 	- 0: yes
	//
	// example:
	//
	// 1
	FilterInvalidIP *int64 `json:"FilterInvalidIP,omitempty" xml:"FilterInvalidIP,omitempty"`
	// The element that is used in DNS hijacking.
	//
	// example:
	//
	// 50
	FlowHijackJumpTimes *int64 `json:"FlowHijackJumpTimes,omitempty" xml:"FlowHijackJumpTimes,omitempty"`
	// The tag that is used in DNS hijacking.
	//
	// example:
	//
	// test
	FlowHijackLogo *string `json:"FlowHijackLogo,omitempty" xml:"FlowHijackLogo,omitempty"`
	// The timeout period of monitoring. Unit: seconds.
	//
	// example:
	//
	// 20
	MonitorTimeout *string `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// Specifies whether to automatically scroll up and down the screen to load a page.
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	NavAutomaticScrolling *string `json:"NavAutomaticScrolling,omitempty" xml:"NavAutomaticScrolling,omitempty"`
	// The method that is used to customize the header. Valid values:
	//
	// 	- 0: disables the customer header.
	//
	// 	- 1: modifies the first package.
	//
	// 	- 2: modifies all packages.
	//
	// example:
	//
	// 0
	NavCustomHeader *string `json:"NavCustomHeader,omitempty" xml:"NavCustomHeader,omitempty"`
	// The format of the custom header. You can specify multiple fields. Separate the fields with vertical bars (|).
	//
	// example:
	//
	// Host:www.example.com|Referer:www.example.com
	NavCustomHeaderContent *string `json:"NavCustomHeaderContent,omitempty" xml:"NavCustomHeaderContent,omitempty"`
	// The custom host mode.
	//
	// 	- 1: round robin
	//
	// 	- 0: random
	//
	// example:
	//
	// 1
	NavCustomHost *int64 `json:"NavCustomHost,omitempty" xml:"NavCustomHost,omitempty"`
	// The custom host IP address. You can enter multiple IP addresses. Separate the IP addresses with commas (,).
	//
	// example:
	//
	// ipv4:192.168.2.1,192.168.2.5:img.a.com|192.168.2.1[8080]:img.a.com
	NavCustomHostIp *string `json:"NavCustomHostIp,omitempty" xml:"NavCustomHostIp,omitempty"`
	// Specifies whether to disable caching.
	//
	// 	- 1: disable
	//
	// 	- 0: enable
	//
	// example:
	//
	// 1
	NavDisableCache *int64 `json:"NavDisableCache,omitempty" xml:"NavDisableCache,omitempty"`
	// Specifies whether to enable the feature of using the Accept-Encoding field to determine whether to accept compressed files.
	//
	// 	- 1: disable
	//
	// 	- 0: enable
	//
	// example:
	//
	// 1
	NavDisableCompression *string `json:"NavDisableCompression,omitempty" xml:"NavDisableCompression,omitempty"`
	// Specifies whether to ignore certificate errors during certificate verification in the SSL handshake and continue browsing.
	//
	// 	- 1: ignore
	//
	// 	- 0: does not ignore
	//
	// example:
	//
	// 1
	NavIgnoreCertificateError *int64 `json:"NavIgnoreCertificateError,omitempty" xml:"NavIgnoreCertificateError,omitempty"`
	// Specifies whether to continue browsing after a redirection occurs.
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	NavRedirection *int64 `json:"NavRedirection,omitempty" xml:"NavRedirection,omitempty"`
	// Specifies whether to return the elements on the page.
	//
	// 	- 1: no. Returns the basic document data.
	//
	// 	- 2: yes. Returns all document data.
	//
	// example:
	//
	// 2
	NavReturnElement *int64 `json:"NavReturnElement,omitempty" xml:"NavReturnElement,omitempty"`
	// The web page defacement.
	//
	// example:
	//
	// www.example.com:202.0.3.55|203.3.44.67
	PageTamper *string `json:"PageTamper,omitempty" xml:"PageTamper,omitempty"`
	// The process ID.
	//
	// example:
	//
	// ssh
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The domain name of the QUIC request element.
	//
	// >  This parameter is supported by all elements of only Chrome
	//
	// example:
	//
	// www.example.com
	QUICDomain *string `json:"QUICDomain,omitempty" xml:"QUICDomain,omitempty"`
	// The Quick UDP Internet Connections (QUIC) protocol version. Default value: 0. Valid values:
	//
	// *
	//
	// 	- 35
	//
	// 	- 39
	//
	// 	- 43
	//
	// 	- 44
	//
	// >  This parameter is supported by all elements of only Chrome
	//
	// example:
	//
	// 0
	QUICVersion *int64 `json:"QUICVersion,omitempty" xml:"QUICVersion,omitempty"`
	// Specifies whether to return the request header.
	//
	// 	- 0: does not return the response header.
	//
	// 	- 1: returns the basic document header.
	//
	// 	- 2: returns all headers.
	//
	// example:
	//
	// 0
	RequestHeader *int64 `json:"RequestHeader,omitempty" xml:"RequestHeader,omitempty"`
	// The method that is used to return the response header. Valid values:
	//
	// 	- 0: does not return the response header.
	//
	// 	- 1: returns the basic document header.
	//
	// 	- 2: returns all headers.
	//
	// example:
	//
	// 0
	ResponseHeader *int64 `json:"ResponseHeader,omitempty" xml:"ResponseHeader,omitempty"`
	// The time threshold that is used to define a slow element. Unit: seconds.
	//
	// example:
	//
	// 5
	SlowElementThreshold *float64 `json:"SlowElementThreshold,omitempty" xml:"SlowElementThreshold,omitempty"`
	// The blacklist for string verification.
	//
	// example:
	//
	// Regex:*.example|expalme|
	VerifyStringBlacklist *string `json:"VerifyStringBlacklist,omitempty" xml:"VerifyStringBlacklist,omitempty"`
	// The whitelist for string verification.
	//
	// example:
	//
	// Regex:*.example|expalme|
	VerifyStringWhiteList *string `json:"VerifyStringWhiteList,omitempty" xml:"VerifyStringWhiteList,omitempty"`
	// The timeout period of waiting for the monitoring to complete.
	//
	// example:
	//
	// 15
	WaitCompletionTime *float64 `json:"WaitCompletionTime,omitempty" xml:"WaitCompletionTime,omitempty"`
}

func (s CreateSyntheticTaskRequestNavigation) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestNavigation) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestNavigation) GetDNSHijackWhiteList() *string {
	return s.DNSHijackWhiteList
}

func (s *CreateSyntheticTaskRequestNavigation) GetElementBlacklist() *string {
	return s.ElementBlacklist
}

func (s *CreateSyntheticTaskRequestNavigation) GetExecuteActiveX() *int64 {
	return s.ExecuteActiveX
}

func (s *CreateSyntheticTaskRequestNavigation) GetExecuteApplication() *int64 {
	return s.ExecuteApplication
}

func (s *CreateSyntheticTaskRequestNavigation) GetExecuteScript() *int64 {
	return s.ExecuteScript
}

func (s *CreateSyntheticTaskRequestNavigation) GetFilterInvalidIP() *int64 {
	return s.FilterInvalidIP
}

func (s *CreateSyntheticTaskRequestNavigation) GetFlowHijackJumpTimes() *int64 {
	return s.FlowHijackJumpTimes
}

func (s *CreateSyntheticTaskRequestNavigation) GetFlowHijackLogo() *string {
	return s.FlowHijackLogo
}

func (s *CreateSyntheticTaskRequestNavigation) GetMonitorTimeout() *string {
	return s.MonitorTimeout
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavAutomaticScrolling() *string {
	return s.NavAutomaticScrolling
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavCustomHeader() *string {
	return s.NavCustomHeader
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavCustomHeaderContent() *string {
	return s.NavCustomHeaderContent
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavCustomHost() *int64 {
	return s.NavCustomHost
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavCustomHostIp() *string {
	return s.NavCustomHostIp
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavDisableCache() *int64 {
	return s.NavDisableCache
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavDisableCompression() *string {
	return s.NavDisableCompression
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavIgnoreCertificateError() *int64 {
	return s.NavIgnoreCertificateError
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavRedirection() *int64 {
	return s.NavRedirection
}

func (s *CreateSyntheticTaskRequestNavigation) GetNavReturnElement() *int64 {
	return s.NavReturnElement
}

func (s *CreateSyntheticTaskRequestNavigation) GetPageTamper() *string {
	return s.PageTamper
}

func (s *CreateSyntheticTaskRequestNavigation) GetProcessName() *string {
	return s.ProcessName
}

func (s *CreateSyntheticTaskRequestNavigation) GetQUICDomain() *string {
	return s.QUICDomain
}

func (s *CreateSyntheticTaskRequestNavigation) GetQUICVersion() *int64 {
	return s.QUICVersion
}

func (s *CreateSyntheticTaskRequestNavigation) GetRequestHeader() *int64 {
	return s.RequestHeader
}

func (s *CreateSyntheticTaskRequestNavigation) GetResponseHeader() *int64 {
	return s.ResponseHeader
}

func (s *CreateSyntheticTaskRequestNavigation) GetSlowElementThreshold() *float64 {
	return s.SlowElementThreshold
}

func (s *CreateSyntheticTaskRequestNavigation) GetVerifyStringBlacklist() *string {
	return s.VerifyStringBlacklist
}

func (s *CreateSyntheticTaskRequestNavigation) GetVerifyStringWhiteList() *string {
	return s.VerifyStringWhiteList
}

func (s *CreateSyntheticTaskRequestNavigation) GetWaitCompletionTime() *float64 {
	return s.WaitCompletionTime
}

func (s *CreateSyntheticTaskRequestNavigation) SetDNSHijackWhiteList(v string) *CreateSyntheticTaskRequestNavigation {
	s.DNSHijackWhiteList = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetElementBlacklist(v string) *CreateSyntheticTaskRequestNavigation {
	s.ElementBlacklist = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetExecuteActiveX(v int64) *CreateSyntheticTaskRequestNavigation {
	s.ExecuteActiveX = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetExecuteApplication(v int64) *CreateSyntheticTaskRequestNavigation {
	s.ExecuteApplication = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetExecuteScript(v int64) *CreateSyntheticTaskRequestNavigation {
	s.ExecuteScript = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetFilterInvalidIP(v int64) *CreateSyntheticTaskRequestNavigation {
	s.FilterInvalidIP = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetFlowHijackJumpTimes(v int64) *CreateSyntheticTaskRequestNavigation {
	s.FlowHijackJumpTimes = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetFlowHijackLogo(v string) *CreateSyntheticTaskRequestNavigation {
	s.FlowHijackLogo = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetMonitorTimeout(v string) *CreateSyntheticTaskRequestNavigation {
	s.MonitorTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavAutomaticScrolling(v string) *CreateSyntheticTaskRequestNavigation {
	s.NavAutomaticScrolling = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavCustomHeader(v string) *CreateSyntheticTaskRequestNavigation {
	s.NavCustomHeader = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavCustomHeaderContent(v string) *CreateSyntheticTaskRequestNavigation {
	s.NavCustomHeaderContent = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavCustomHost(v int64) *CreateSyntheticTaskRequestNavigation {
	s.NavCustomHost = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavCustomHostIp(v string) *CreateSyntheticTaskRequestNavigation {
	s.NavCustomHostIp = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavDisableCache(v int64) *CreateSyntheticTaskRequestNavigation {
	s.NavDisableCache = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavDisableCompression(v string) *CreateSyntheticTaskRequestNavigation {
	s.NavDisableCompression = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavIgnoreCertificateError(v int64) *CreateSyntheticTaskRequestNavigation {
	s.NavIgnoreCertificateError = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavRedirection(v int64) *CreateSyntheticTaskRequestNavigation {
	s.NavRedirection = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetNavReturnElement(v int64) *CreateSyntheticTaskRequestNavigation {
	s.NavReturnElement = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetPageTamper(v string) *CreateSyntheticTaskRequestNavigation {
	s.PageTamper = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetProcessName(v string) *CreateSyntheticTaskRequestNavigation {
	s.ProcessName = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetQUICDomain(v string) *CreateSyntheticTaskRequestNavigation {
	s.QUICDomain = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetQUICVersion(v int64) *CreateSyntheticTaskRequestNavigation {
	s.QUICVersion = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetRequestHeader(v int64) *CreateSyntheticTaskRequestNavigation {
	s.RequestHeader = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetResponseHeader(v int64) *CreateSyntheticTaskRequestNavigation {
	s.ResponseHeader = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetSlowElementThreshold(v float64) *CreateSyntheticTaskRequestNavigation {
	s.SlowElementThreshold = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetVerifyStringBlacklist(v string) *CreateSyntheticTaskRequestNavigation {
	s.VerifyStringBlacklist = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetVerifyStringWhiteList(v string) *CreateSyntheticTaskRequestNavigation {
	s.VerifyStringWhiteList = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) SetWaitCompletionTime(v float64) *CreateSyntheticTaskRequestNavigation {
	s.WaitCompletionTime = &v
	return s
}

func (s *CreateSyntheticTaskRequestNavigation) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestNet struct {
	// The DNS server.
	//
	// example:
	//
	// 114.114.XX.XX
	NetDNSNs *string `json:"NetDNSNs,omitempty" xml:"NetDNSNs,omitempty"`
	// The DNS query method. Valid values:
	//
	// 	- 1: recursion
	//
	// 	- 2: iteration
	//
	// example:
	//
	// 1
	NetDNSQueryMethod *int64 `json:"NetDNSQueryMethod,omitempty" xml:"NetDNSQueryMethod,omitempty"`
	// The IP address type of the DNS server.
	//
	// 	- 0: IPv4
	//
	// 	- 1: IPv6
	//
	// 	- 2: an automatic IP address
	//
	// example:
	//
	// 0
	NetDNSServer *int64 `json:"NetDNSServer,omitempty" xml:"NetDNSServer,omitempty"`
	// Specifies whether to enable domain name system (DNS) monitoring.
	//
	// 	- 0: Off.
	//
	// 	- 1: On. You must set DNS parameters if you want to enable DNS monitoring.
	//
	// example:
	//
	// 1
	NetDNSSwitch *int64 `json:"NetDNSSwitch,omitempty" xml:"NetDNSSwitch,omitempty"`
	// The timeout period of DNS monitoring. Default value: 5 seconds. Valid values: 0 to 45 seconds.
	//
	// example:
	//
	// 5
	NetDNSTimeout *int64 `json:"NetDNSTimeout,omitempty" xml:"NetDNSTimeout,omitempty"`
	// Specifies whether to display the data in the DIG format. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	NetDigSwitch *int64 `json:"NetDigSwitch,omitempty" xml:"NetDigSwitch,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- 0: ICMP
	//
	// 	- 1: TCP
	//
	// example:
	//
	// 0
	NetICMPActive *int64 `json:"NetICMPActive,omitempty" xml:"NetICMPActive,omitempty"`
	// Specifies whether to split packages.
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	NetICMPDataCut *int64 `json:"NetICMPDataCut,omitempty" xml:"NetICMPDataCut,omitempty"`
	// The interval at which the network synthetic monitoring task is executed. Unit: seconds.
	//
	// example:
	//
	// 1
	NetICMPInterval *int64 `json:"NetICMPInterval,omitempty" xml:"NetICMPInterval,omitempty"`
	// The number of packages.
	//
	// example:
	//
	// 4
	NetICMPNum *int64 `json:"NetICMPNum,omitempty" xml:"NetICMPNum,omitempty"`
	// The package size.
	//
	// example:
	//
	// 32
	NetICMPSize *int64 `json:"NetICMPSize,omitempty" xml:"NetICMPSize,omitempty"`
	// Specifies whether to enable ping monitoring.
	//
	// 	- 0: Off.
	//
	// 	- 1: On. You must set Internet control message protocol (ICMP) parameters if you want to enable ping monitoring.
	//
	// example:
	//
	// 1
	NetICMPSwitch *int64 `json:"NetICMPSwitch,omitempty" xml:"NetICMPSwitch,omitempty"`
	// The timeout period of Ping monitoring.
	//
	// example:
	//
	// 20
	NetICMPTimeout *int64 `json:"NetICMPTimeout,omitempty" xml:"NetICMPTimeout,omitempty"`
	// The maximum number of active monitoring points.
	//
	// example:
	//
	// 20
	NetTraceRouteNum *int64 `json:"NetTraceRouteNum,omitempty" xml:"NetTraceRouteNum,omitempty"`
	// Specifies whether to enable tracert monitoring.
	//
	// 	- 0: Off.
	//
	// 	- 1: On. You must set the tracert parameters if you want to enable tracert monitoring.
	//
	// example:
	//
	// 1
	NetTraceRouteSwitch *int64 `json:"NetTraceRouteSwitch,omitempty" xml:"NetTraceRouteSwitch,omitempty"`
	// The timeout period of tracert monitoring. Valid values: 0 to 300 seconds.
	//
	// example:
	//
	// 60
	NetTraceRouteTimeout *int64 `json:"NetTraceRouteTimeout,omitempty" xml:"NetTraceRouteTimeout,omitempty"`
	// The whitelist for DNS hijacking. The format is `Domain name: Matching rule`.
	//
	// >  Wireless application protocol (WAP) networks do not support DNS hijacking.
	//
	// example:
	//
	// www.aliyun.com:202.0.3.55|203.3.44.67
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s CreateSyntheticTaskRequestNet) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestNet) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestNet) GetNetDNSNs() *string {
	return s.NetDNSNs
}

func (s *CreateSyntheticTaskRequestNet) GetNetDNSQueryMethod() *int64 {
	return s.NetDNSQueryMethod
}

func (s *CreateSyntheticTaskRequestNet) GetNetDNSServer() *int64 {
	return s.NetDNSServer
}

func (s *CreateSyntheticTaskRequestNet) GetNetDNSSwitch() *int64 {
	return s.NetDNSSwitch
}

func (s *CreateSyntheticTaskRequestNet) GetNetDNSTimeout() *int64 {
	return s.NetDNSTimeout
}

func (s *CreateSyntheticTaskRequestNet) GetNetDigSwitch() *int64 {
	return s.NetDigSwitch
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPActive() *int64 {
	return s.NetICMPActive
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPDataCut() *int64 {
	return s.NetICMPDataCut
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPInterval() *int64 {
	return s.NetICMPInterval
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPNum() *int64 {
	return s.NetICMPNum
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPSize() *int64 {
	return s.NetICMPSize
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPSwitch() *int64 {
	return s.NetICMPSwitch
}

func (s *CreateSyntheticTaskRequestNet) GetNetICMPTimeout() *int64 {
	return s.NetICMPTimeout
}

func (s *CreateSyntheticTaskRequestNet) GetNetTraceRouteNum() *int64 {
	return s.NetTraceRouteNum
}

func (s *CreateSyntheticTaskRequestNet) GetNetTraceRouteSwitch() *int64 {
	return s.NetTraceRouteSwitch
}

func (s *CreateSyntheticTaskRequestNet) GetNetTraceRouteTimeout() *int64 {
	return s.NetTraceRouteTimeout
}

func (s *CreateSyntheticTaskRequestNet) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSNs(v string) *CreateSyntheticTaskRequestNet {
	s.NetDNSNs = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSQueryMethod(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSQueryMethod = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSServer(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSServer = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSTimeout(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDigSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDigSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPActive(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPActive = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPDataCut(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPDataCut = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPInterval(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPInterval = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPNum(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPNum = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPSize(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPSize = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPTimeout(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetTraceRouteNum(v int64) *CreateSyntheticTaskRequestNet {
	s.NetTraceRouteNum = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetTraceRouteSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetTraceRouteSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetTraceRouteTimeout(v int64) *CreateSyntheticTaskRequestNet {
	s.NetTraceRouteTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetWhiteList(v string) *CreateSyntheticTaskRequestNet {
	s.WhiteList = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestProtocol struct {
	// The encoding format.
	//
	// 	- 0: UTF-8
	//
	// 	- 1: GBK
	//
	// 	- 2: GB2312
	//
	// 	- 3: Unicode
	//
	// example:
	//
	// 0
	CharacterEncoding *int64 `json:"CharacterEncoding,omitempty" xml:"CharacterEncoding,omitempty"`
	// The custom host mode.
	//
	// 	- 1: round robin
	//
	// 	- 0: random
	//
	// example:
	//
	// 1
	CustomHost *int64 `json:"CustomHost,omitempty" xml:"CustomHost,omitempty"`
	// The custom host IP address. You can enter multiple IP addresses. Separate the IP addresses with commas (,).
	//
	// example:
	//
	// ipv4:192.168.2.1,192.168.2.5:img.a.com|192.168.2.1[8080]:img.a.com
	CustomHostIp *string `json:"CustomHostIp,omitempty" xml:"CustomHostIp,omitempty"`
	// The connection timeout period of the protocol. Unit: seconds.
	//
	// example:
	//
	// 3
	ProtocolConnectionTime *int64 `json:"ProtocolConnectionTime,omitempty" xml:"ProtocolConnectionTime,omitempty"`
	// The timeout period of API performance synthetic monitoring. Unit: seconds.
	//
	// example:
	//
	// 30
	ProtocolMonitorTimeout *string `json:"ProtocolMonitorTimeout,omitempty" xml:"ProtocolMonitorTimeout,omitempty"`
	// The size of the received data. This parameter is required when you set the value of the VerifyWay parameter to 2.
	//
	// example:
	//
	// 500
	ReceivedDataSize *int64 `json:"ReceivedDataSize,omitempty" xml:"ReceivedDataSize,omitempty"`
	// The request content, including the request header and request body.
	RequestContent *CreateSyntheticTaskRequestProtocolRequestContent `json:"RequestContent,omitempty" xml:"RequestContent,omitempty" type:"Struct"`
	// The verification string.
	//
	// example:
	//
	// "code":200
	VerifyContent *string `json:"VerifyContent,omitempty" xml:"VerifyContent,omitempty"`
	// The method that is used to verify the response content.
	//
	// 	- 0: no verification.
	//
	// 	- 1: exact match with the verification string.
	//
	// 	- 2: partial match with the verification string.
	//
	// 	- 3: MD5 verification.
	//
	// example:
	//
	// 0
	VerifyWay *int64 `json:"VerifyWay,omitempty" xml:"VerifyWay,omitempty"`
}

func (s CreateSyntheticTaskRequestProtocol) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestProtocol) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestProtocol) GetCharacterEncoding() *int64 {
	return s.CharacterEncoding
}

func (s *CreateSyntheticTaskRequestProtocol) GetCustomHost() *int64 {
	return s.CustomHost
}

func (s *CreateSyntheticTaskRequestProtocol) GetCustomHostIp() *string {
	return s.CustomHostIp
}

func (s *CreateSyntheticTaskRequestProtocol) GetProtocolConnectionTime() *int64 {
	return s.ProtocolConnectionTime
}

func (s *CreateSyntheticTaskRequestProtocol) GetProtocolMonitorTimeout() *string {
	return s.ProtocolMonitorTimeout
}

func (s *CreateSyntheticTaskRequestProtocol) GetReceivedDataSize() *int64 {
	return s.ReceivedDataSize
}

func (s *CreateSyntheticTaskRequestProtocol) GetRequestContent() *CreateSyntheticTaskRequestProtocolRequestContent {
	return s.RequestContent
}

func (s *CreateSyntheticTaskRequestProtocol) GetVerifyContent() *string {
	return s.VerifyContent
}

func (s *CreateSyntheticTaskRequestProtocol) GetVerifyWay() *int64 {
	return s.VerifyWay
}

func (s *CreateSyntheticTaskRequestProtocol) SetCharacterEncoding(v int64) *CreateSyntheticTaskRequestProtocol {
	s.CharacterEncoding = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetCustomHost(v int64) *CreateSyntheticTaskRequestProtocol {
	s.CustomHost = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetCustomHostIp(v string) *CreateSyntheticTaskRequestProtocol {
	s.CustomHostIp = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetProtocolConnectionTime(v int64) *CreateSyntheticTaskRequestProtocol {
	s.ProtocolConnectionTime = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetProtocolMonitorTimeout(v string) *CreateSyntheticTaskRequestProtocol {
	s.ProtocolMonitorTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetReceivedDataSize(v int64) *CreateSyntheticTaskRequestProtocol {
	s.ReceivedDataSize = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetRequestContent(v *CreateSyntheticTaskRequestProtocolRequestContent) *CreateSyntheticTaskRequestProtocol {
	s.RequestContent = v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetVerifyContent(v string) *CreateSyntheticTaskRequestProtocol {
	s.VerifyContent = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) SetVerifyWay(v int64) *CreateSyntheticTaskRequestProtocol {
	s.VerifyWay = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocol) Validate() error {
	if s.RequestContent != nil {
		if err := s.RequestContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSyntheticTaskRequestProtocolRequestContent struct {
	// The custom body of a request to initiate an API performance synthetic monitoring task.
	Body *CreateSyntheticTaskRequestProtocolRequestContentBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The custom header of a request to initiate an API performance synthetic monitoring task.
	Header []*CreateSyntheticTaskRequestProtocolRequestContentHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
	// The request method.
	//
	// 	- POST
	//
	// 	- GET
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s CreateSyntheticTaskRequestProtocolRequestContent) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestProtocolRequestContent) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) GetBody() *CreateSyntheticTaskRequestProtocolRequestContentBody {
	return s.Body
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) GetHeader() []*CreateSyntheticTaskRequestProtocolRequestContentHeader {
	return s.Header
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) GetMethod() *string {
	return s.Method
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) SetBody(v *CreateSyntheticTaskRequestProtocolRequestContentBody) *CreateSyntheticTaskRequestProtocolRequestContent {
	s.Body = v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) SetHeader(v []*CreateSyntheticTaskRequestProtocolRequestContentHeader) *CreateSyntheticTaskRequestProtocolRequestContent {
	s.Header = v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) SetMethod(v string) *CreateSyntheticTaskRequestProtocolRequestContent {
	s.Method = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContent) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Header != nil {
		for _, item := range s.Header {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSyntheticTaskRequestProtocolRequestContentBody struct {
	// The data that is passed when the **Mode*	- parameter is set to **form-data**.
	FormData []*CreateSyntheticTaskRequestProtocolRequestContentBodyFormData `json:"FormData,omitempty" xml:"FormData,omitempty" type:"Repeated"`
	// The language that is selected when the Mode parameter is set to raw.
	//
	// 	- json
	//
	// 	- xml
	//
	// 	- javascript
	//
	// 	- html
	//
	// 	- text
	//
	// example:
	//
	// json
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The data type of the content.
	//
	// 	- form-data
	//
	// 	- x-www-form-urlencoded
	//
	// 	- raw
	//
	// example:
	//
	// form-data
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The data that is passed when the **Mode*	- parameter is set to **raw**.
	//
	// example:
	//
	// content
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The data that is passed when the **Mode*	- parameter is set to **x-www-form-urlencoded**.
	UrlEncoding []*CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding `json:"UrlEncoding,omitempty" xml:"UrlEncoding,omitempty" type:"Repeated"`
}

func (s CreateSyntheticTaskRequestProtocolRequestContentBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestProtocolRequestContentBody) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) GetFormData() []*CreateSyntheticTaskRequestProtocolRequestContentBodyFormData {
	return s.FormData
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) GetLanguage() *string {
	return s.Language
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) GetMode() *string {
	return s.Mode
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) GetRaw() *string {
	return s.Raw
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) GetUrlEncoding() []*CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding {
	return s.UrlEncoding
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) SetFormData(v []*CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) *CreateSyntheticTaskRequestProtocolRequestContentBody {
	s.FormData = v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) SetLanguage(v string) *CreateSyntheticTaskRequestProtocolRequestContentBody {
	s.Language = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) SetMode(v string) *CreateSyntheticTaskRequestProtocolRequestContentBody {
	s.Mode = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) SetRaw(v string) *CreateSyntheticTaskRequestProtocolRequestContentBody {
	s.Raw = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) SetUrlEncoding(v []*CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) *CreateSyntheticTaskRequestProtocolRequestContentBody {
	s.UrlEncoding = v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBody) Validate() error {
	if s.FormData != nil {
		for _, item := range s.FormData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UrlEncoding != nil {
		for _, item := range s.UrlEncoding {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSyntheticTaskRequestProtocolRequestContentBodyFormData struct {
	// The key of **form-data**.
	//
	// example:
	//
	// appId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of **form-data**.
	//
	// example:
	//
	// 3425
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) GetKey() *string {
	return s.Key
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) GetValue() *string {
	return s.Value
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) SetKey(v string) *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData {
	s.Key = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) SetValue(v string) *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData {
	s.Value = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyFormData) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding struct {
	// The key of **x-www-form-urlencoded**.
	//
	// example:
	//
	// appId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of **x-www-form-urlencoded**.
	//
	// example:
	//
	// 11080
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) GetKey() *string {
	return s.Key
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) GetValue() *string {
	return s.Value
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) SetKey(v string) *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding {
	s.Key = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) SetValue(v string) *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding {
	s.Value = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentBodyUrlEncoding) Validate() error {
	return dara.Validate(s)
}

type CreateSyntheticTaskRequestProtocolRequestContentHeader struct {
	// The key of the request header.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the request header.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSyntheticTaskRequestProtocolRequestContentHeader) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskRequestProtocolRequestContentHeader) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentHeader) GetKey() *string {
	return s.Key
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentHeader) GetValue() *string {
	return s.Value
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentHeader) SetKey(v string) *CreateSyntheticTaskRequestProtocolRequestContentHeader {
	s.Key = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentHeader) SetValue(v string) *CreateSyntheticTaskRequestProtocolRequestContentHeader {
	s.Value = &v
	return s
}

func (s *CreateSyntheticTaskRequestProtocolRequestContentHeader) Validate() error {
	return dara.Validate(s)
}
