// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSyntheticTaskDetailResponseBody
	GetRequestId() *string
	SetTaskDetail(v *GetSyntheticTaskDetailResponseBodyTaskDetail) *GetSyntheticTaskDetailResponseBody
	GetTaskDetail() *GetSyntheticTaskDetailResponseBodyTaskDetail
}

type GetSyntheticTaskDetailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4D6C358A-A58B-4F4B-94CE-F5AAF023****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the task.
	TaskDetail *GetSyntheticTaskDetailResponseBodyTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Struct"`
}

func (s GetSyntheticTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSyntheticTaskDetailResponseBody) GetTaskDetail() *GetSyntheticTaskDetailResponseBodyTaskDetail {
	return s.TaskDetail
}

func (s *GetSyntheticTaskDetailResponseBody) SetRequestId(v string) *GetSyntheticTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBody) SetTaskDetail(v *GetSyntheticTaskDetailResponseBodyTaskDetail) *GetSyntheticTaskDetailResponseBody {
	s.TaskDetail = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBody) Validate() error {
	if s.TaskDetail != nil {
		if err := s.TaskDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSyntheticTaskDetailResponseBodyTaskDetail struct {
	// The list of common parameters.
	CommonParam *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam `json:"CommonParam,omitempty" xml:"CommonParam,omitempty" type:"Struct"`
	// The file download task.
	Download *GetSyntheticTaskDetailResponseBodyTaskDetailDownload `json:"Download,omitempty" xml:"Download,omitempty" type:"Struct"`
	// The frequency.
	ExtendInterval *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval `json:"ExtendInterval,omitempty" xml:"ExtendInterval,omitempty" type:"Struct"`
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
	// example:
	//
	// 20
	IntervalTime *int64 `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	// The interval type. Valid values:
	//
	// 	- 0: daily
	//
	// 	- 1: custom frequency
	//
	// example:
	//
	// 0
	IntervalType *int64 `json:"IntervalType,omitempty" xml:"IntervalType,omitempty"`
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
	IpType *int64 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The detection points.
	MonitorList []*GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList `json:"MonitorList,omitempty" xml:"MonitorList,omitempty" type:"Repeated"`
	// The detection points.
	//
	// example:
	//
	// 12
	MonitorListString *string `json:"MonitorListString,omitempty" xml:"MonitorListString,omitempty"`
	// The browser test task.
	Nav *GetSyntheticTaskDetailResponseBodyTaskDetailNav `json:"Nav,omitempty" xml:"Nav,omitempty" type:"Struct"`
	// The network synthetic monitoring task.
	Net *GetSyntheticTaskDetailResponseBodyTaskDetailNet `json:"Net,omitempty" xml:"Net,omitempty" type:"Struct"`
	// The synthetic monitoring task of the API performance type.
	Protocol *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol `json:"Protocol,omitempty" xml:"Protocol,omitempty" type:"Struct"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 19584
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// net-test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 1.  3: web page performance - IE
	//
	// 2.  34: web page performance - Chrome
	//
	// 3.  0: network quality
	//
	// 4.  40: file download
	//
	// 5.  7: API performance
	//
	// example:
	//
	// 0
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The URL for synthetic monitoring.
	//
	// example:
	//
	// www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetail) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetCommonParam() *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	return s.CommonParam
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetDownload() *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	return s.Download
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetExtendInterval() *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	return s.ExtendInterval
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetIntervalTime() *int64 {
	return s.IntervalTime
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetIntervalType() *int64 {
	return s.IntervalType
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetIpType() *int64 {
	return s.IpType
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetMonitorList() []*GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList {
	return s.MonitorList
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetMonitorListString() *string {
	return s.MonitorListString
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetNav() *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	return s.Nav
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetNet() *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	return s.Net
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetProtocol() *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	return s.Protocol
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetTaskName() *string {
	return s.TaskName
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetTaskType() *int64 {
	return s.TaskType
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) GetUrl() *string {
	return s.Url
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetCommonParam(v *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.CommonParam = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetDownload(v *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.Download = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetExtendInterval(v *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.ExtendInterval = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetIntervalTime(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.IntervalTime = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetIntervalType(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.IntervalType = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetIpType(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.IpType = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetMonitorList(v []*GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.MonitorList = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetMonitorListString(v string) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.MonitorListString = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetNav(v *GetSyntheticTaskDetailResponseBodyTaskDetailNav) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.Nav = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetNet(v *GetSyntheticTaskDetailResponseBodyTaskDetailNet) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.Net = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetProtocol(v *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.Protocol = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetTaskId(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.TaskId = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetTaskName(v string) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.TaskName = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetTaskType(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.TaskType = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) SetUrl(v string) *GetSyntheticTaskDetailResponseBodyTaskDetail {
	s.Url = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetail) Validate() error {
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
	if s.Nav != nil {
		if err := s.Nav.Validate(); err != nil {
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

type GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam struct {
	// The identifier of the alert.
	//
	// example:
	//
	// 1
	AlarmFlag *int64 `json:"AlarmFlag,omitempty" xml:"AlarmFlag,omitempty"`
	// The list of alerts.
	AlertList []*GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList `json:"AlertList,omitempty" xml:"AlertList,omitempty" type:"Repeated"`
	// The ID of the alert identifier.
	//
	// example:
	//
	// 1
	AlertNotifierId *string `json:"AlertNotifierId,omitempty" xml:"AlertNotifierId,omitempty"`
	// The ID of the alert policy.
	//
	// example:
	//
	// 1
	AlertPolicyId *string `json:"AlertPolicyId,omitempty" xml:"AlertPolicyId,omitempty"`
	// The monitoring samples.
	//
	// example:
	//
	// 1
	MonitorSamples *string `json:"MonitorSamples,omitempty" xml:"MonitorSamples,omitempty"`
	// The start time of the execution.
	//
	// example:
	//
	// 1664427128
	StartExecutionTime *string `json:"StartExecutionTime,omitempty" xml:"StartExecutionTime,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GetAlarmFlag() *int64 {
	return s.AlarmFlag
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GetAlertList() []*GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList {
	return s.AlertList
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GetAlertNotifierId() *string {
	return s.AlertNotifierId
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GetAlertPolicyId() *string {
	return s.AlertPolicyId
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GetMonitorSamples() *string {
	return s.MonitorSamples
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) GetStartExecutionTime() *string {
	return s.StartExecutionTime
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) SetAlarmFlag(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	s.AlarmFlag = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) SetAlertList(v []*GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	s.AlertList = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) SetAlertNotifierId(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	s.AlertNotifierId = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) SetAlertPolicyId(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	s.AlertPolicyId = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) SetMonitorSamples(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	s.MonitorSamples = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) SetStartExecutionTime(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam {
	s.StartExecutionTime = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParam) Validate() error {
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

type GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList struct {
	// The low-risk alert.
	//
	// example:
	//
	// 1
	GeneralAlert *string `json:"GeneralAlert,omitempty" xml:"GeneralAlert,omitempty"`
	// Indicates whether the condition is essential.
	//
	// example:
	//
	// 0
	IsCritical *string `json:"IsCritical,omitempty" xml:"IsCritical,omitempty"`
	// The alert name.
	//
	// example:
	//
	// alert-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Critical-level alert.
	//
	// example:
	//
	// 1
	SeriousAlert *string `json:"SeriousAlert,omitempty" xml:"SeriousAlert,omitempty"`
	// Greater than or less than.
	//
	// example:
	//
	// 1
	Symbols *string `json:"Symbols,omitempty" xml:"Symbols,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) GetGeneralAlert() *string {
	return s.GeneralAlert
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) GetIsCritical() *string {
	return s.IsCritical
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) GetName() *string {
	return s.Name
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) GetSeriousAlert() *string {
	return s.SeriousAlert
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) GetSymbols() *string {
	return s.Symbols
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) SetGeneralAlert(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList {
	s.GeneralAlert = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) SetIsCritical(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList {
	s.IsCritical = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) SetName(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList {
	s.Name = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) SetSeriousAlert(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList {
	s.SeriousAlert = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) SetSymbols(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList {
	s.Symbols = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailCommonParamAlertList) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailDownload struct {
	// The timeout period of the file download task.
	//
	// example:
	//
	// 200
	ConnectionTimeout *int64 `json:"ConnectionTimeout,omitempty" xml:"ConnectionTimeout,omitempty"`
	// The items to be ignored in a certificate error. Multiple values are concatenated with vertical bars (|).
	//
	// example:
	//
	// char
	DownloadCustomHeaderContent *string `json:"DownloadCustomHeaderContent,omitempty" xml:"DownloadCustomHeaderContent,omitempty"`
	// The custom host. Valid values:
	//
	// 	- 1: round robin
	//
	// 	- 0: random
	//
	// example:
	//
	// 0
	DownloadCustomHost *int64 `json:"DownloadCustomHost,omitempty" xml:"DownloadCustomHost,omitempty"`
	// The custom IP address of the host. Multiple IP addresses are separated with commas (,).
	//
	// example:
	//
	// 168.23.45.1
	DownloadCustomHostIp *string `json:"DownloadCustomHostIp,omitempty" xml:"DownloadCustomHostIp,omitempty"`
	// The kernel type. Valid values:
	//
	// 	- 1: curl
	//
	// 	- 0: WinInet
	//
	// example:
	//
	// 1
	DownloadKernel *int64 `json:"DownloadKernel,omitempty" xml:"DownloadKernel,omitempty"`
	// Indicates whether redirection is supported.
	//
	// example:
	//
	// 1
	DownloadRedirect *int64 `json:"DownloadRedirect,omitempty" xml:"DownloadRedirect,omitempty"`
	// The file size. Unit: KB.
	//
	// example:
	//
	// 20
	DownloadTransmissionSize *int64 `json:"DownloadTransmissionSize,omitempty" xml:"DownloadTransmissionSize,omitempty"`
	// The monitoring duration.
	//
	// example:
	//
	// 12
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// The QUIC protocol type. Valid values:
	//
	// 	- 1: HTTP/1
	//
	// 	- 2: HTTP/2
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
	// test
	ValidateKeywords *string `json:"ValidateKeywords,omitempty" xml:"ValidateKeywords,omitempty"`
	// The method that is used to verify the response content. Valid values:
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
	// The whitelisted objects that are used to avoid DNS hijacking. Format: `<domain name>:<objects>`.
	//
	// >  WAP networks do not support hijacking.
	//
	// example:
	//
	// [{\\"src\\":\\"211.154.166.174\\"}]
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailDownload) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetConnectionTimeout() *int64 {
	return s.ConnectionTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetDownloadCustomHeaderContent() *string {
	return s.DownloadCustomHeaderContent
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetDownloadCustomHost() *int64 {
	return s.DownloadCustomHost
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetDownloadCustomHostIp() *string {
	return s.DownloadCustomHostIp
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetDownloadKernel() *int64 {
	return s.DownloadKernel
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetDownloadRedirect() *int64 {
	return s.DownloadRedirect
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetDownloadTransmissionSize() *int64 {
	return s.DownloadTransmissionSize
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetQuickProtocol() *string {
	return s.QuickProtocol
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetValidateKeywords() *string {
	return s.ValidateKeywords
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetVerifyWay() *int64 {
	return s.VerifyWay
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) GetWhiteList() *string {
	return s.WhiteList
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetConnectionTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.ConnectionTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetDownloadCustomHeaderContent(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.DownloadCustomHeaderContent = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetDownloadCustomHost(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.DownloadCustomHost = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetDownloadCustomHostIp(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.DownloadCustomHostIp = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetDownloadKernel(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.DownloadKernel = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetDownloadRedirect(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.DownloadRedirect = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetDownloadTransmissionSize(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.DownloadTransmissionSize = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetMonitorTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.MonitorTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetQuickProtocol(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.QuickProtocol = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetValidateKeywords(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.ValidateKeywords = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetVerifyWay(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.VerifyWay = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) SetWhiteList(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailDownload {
	s.WhiteList = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailDownload) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval struct {
	// The day on which synthetic monitoring is performed. Valid values:
	//
	// 	- \\-1: every day
	//
	// 	- 0: Sunday
	//
	// 	- 1: Monday
	//
	// 	- 2: Tuesday
	//
	// 	- 3: Wednesday
	//
	// 	- 4: Thursday
	//
	// 	- 5: Friday
	//
	// 	- 6: Saturday
	Days []*int64 `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	// The minute at which synthetic monitoring ends.
	//
	// example:
	//
	// 20
	EndMinute *int64 `json:"EndMinute,omitempty" xml:"EndMinute,omitempty"`
	// The time when synthetic monitoring ends. Format: `yyyy-MM-dd HH`.
	//
	// example:
	//
	// 2022-05-03 11:40
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The hour at which synthetic monitoring ends.
	//
	// example:
	//
	// 12
	Endhour *int64 `json:"Endhour,omitempty" xml:"Endhour,omitempty"`
	// The hour at which synthetic monitoring starts.
	//
	// example:
	//
	// 9
	StartHour *int64 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
	// The minute at which synthetic monitoring starts.
	//
	// example:
	//
	// 20
	StartMinute *int64 `json:"StartMinute,omitempty" xml:"StartMinute,omitempty"`
	// The time when synthetic monitoring starts. Format: yyyy-MM-dd HH.
	//
	// example:
	//
	// 2022-02-26 11:40
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetDays() []*int64 {
	return s.Days
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetEndMinute() *int64 {
	return s.EndMinute
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetEndhour() *int64 {
	return s.Endhour
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetStartHour() *int64 {
	return s.StartHour
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetStartMinute() *int64 {
	return s.StartMinute
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetDays(v []*int64) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.Days = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetEndMinute(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.EndMinute = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetEndTime(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.EndTime = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetEndhour(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.Endhour = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetStartHour(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.StartHour = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetStartMinute(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.StartMinute = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) SetStartTime(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval {
	s.StartTime = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailExtendInterval) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList struct {
	// The city code.
	//
	// example:
	//
	// 110100
	CityCode *int64 `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The type of the detection point.
	//
	// example:
	//
	// IDC
	MonitorType *int64 `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// The ID of the network service.
	//
	// example:
	//
	// 12001
	NetServiceId *int64 `json:"NetServiceId,omitempty" xml:"NetServiceId,omitempty"`
	// The number of times that the system sends detection requests.
	//
	// example:
	//
	// 20
	SendCount *int64 `json:"SendCount,omitempty" xml:"SendCount,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) GetCityCode() *int64 {
	return s.CityCode
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) GetMonitorType() *int64 {
	return s.MonitorType
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) GetNetServiceId() *int64 {
	return s.NetServiceId
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) GetSendCount() *int64 {
	return s.SendCount
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) SetCityCode(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList {
	s.CityCode = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) SetMonitorType(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList {
	s.MonitorType = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) SetNetServiceId(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList {
	s.NetServiceId = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) SetSendCount(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList {
	s.SendCount = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailMonitorList) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailNav struct {
	// The DNS whitelist.
	//
	// example:
	//
	// 119.119.53.156
	DnsHijackWhitelist *string `json:"DnsHijackWhitelist,omitempty" xml:"DnsHijackWhitelist,omitempty"`
	// The element blacklist.
	//
	// example:
	//
	// test
	ElementBlacklist *string `json:"ElementBlacklist,omitempty" xml:"ElementBlacklist,omitempty"`
	// Indicates whether ActiveX is executed. Valid values:
	//
	// 	- 3: yes
	//
	// 	- 0: no
	//
	// >  Only IE elements support this parameter.
	//
	// example:
	//
	// 3
	ExecuteActiveX *int64 `json:"ExecuteActiveX,omitempty" xml:"ExecuteActiveX,omitempty"`
	// Indicates whether the applet is executed. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	ExecuteApplet *int64 `json:"ExecuteApplet,omitempty" xml:"ExecuteApplet,omitempty"`
	// Indicates whether scripts are executed. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// >  Only IE elements support this parameter.
	//
	// example:
	//
	// 1
	ExecuteScript *int64 `json:"ExecuteScript,omitempty" xml:"ExecuteScript,omitempty"`
	// Indicates whether invalid IP addresses are excluded. Valid values:
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
	// 12
	FlowHijackJumpTimes *int64 `json:"FlowHijackJumpTimes,omitempty" xml:"FlowHijackJumpTimes,omitempty"`
	// The tag that is used in DNS hijacking.
	//
	// example:
	//
	// target
	FlowHijackLogo *string `json:"FlowHijackLogo,omitempty" xml:"FlowHijackLogo,omitempty"`
	// The monitoring timeout period.
	//
	// example:
	//
	// 20
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// Indicates whether the screen is automatically scrolled up and down to load a page. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	NavAutomaticScrolling *int64 `json:"NavAutomaticScrolling,omitempty" xml:"NavAutomaticScrolling,omitempty"`
	// Indicates whether a custom header is created. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: A custom header is created for the first packet.
	//
	// 	- 2: A custom header is created for all packets.
	//
	// example:
	//
	// 0
	NavCustomHeader *string `json:"NavCustomHeader,omitempty" xml:"NavCustomHeader,omitempty"`
	// The format of the custom header. Multiple fields are separated with vertical bars (|).
	//
	// example:
	//
	// content
	NavCustomHeaderContent *string `json:"NavCustomHeaderContent,omitempty" xml:"NavCustomHeaderContent,omitempty"`
	// The custom host mode. Valid values:
	//
	// 	- 1: round robin
	//
	// 	- 0: random
	//
	// example:
	//
	// 1
	NavCustomHost *int64 `json:"NavCustomHost,omitempty" xml:"NavCustomHost,omitempty"`
	// The custom IP address of the host. Multiple IP addresses are separated with commas (,).
	//
	// example:
	//
	// 119.119.53.156/32
	NavCustomHostIp *string `json:"NavCustomHostIp,omitempty" xml:"NavCustomHostIp,omitempty"`
	// Indicates whether caching is disabled. Valid values:
	//
	// 	- 1: Caching is disabled.
	//
	// 	- 0: Caching is enabled.
	//
	// example:
	//
	// 1
	NavDisableCache *int64 `json:"NavDisableCache,omitempty" xml:"NavDisableCache,omitempty"`
	// Indicates whether compression is disabled. Valid values:
	//
	// 	- 0: Compression is enabled.
	//
	// 	- 1: Compression is disabled.
	//
	// example:
	//
	// 0
	NavDisableCompression *int64 `json:"NavDisableCompression,omitempty" xml:"NavDisableCompression,omitempty"`
	// Indicates whether certificate errors are ignored during certificate verification in the SSL handshake. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// example:
	//
	// 1
	NavIgnoreCertificateError *int64 `json:"NavIgnoreCertificateError,omitempty" xml:"NavIgnoreCertificateError,omitempty"`
	// Indicates whether redirection is enabled. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	NavRedirect *int64 `json:"NavRedirect,omitempty" xml:"NavRedirect,omitempty"`
	// Indicates whether the elements on the page are returned.
	//
	// 	- 1: no. The basic document data is returned.
	//
	// 	- 2: yes. All document data is returned.
	//
	// example:
	//
	// 1
	NavReturnElement *int64 `json:"NavReturnElement,omitempty" xml:"NavReturnElement,omitempty"`
	// The page tampering.
	//
	// example:
	//
	// content
	PageTampering *string `json:"PageTampering,omitempty" xml:"PageTampering,omitempty"`
	// The process ID.
	//
	// example:
	//
	// ssh
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The domain name of the QUIC request element.
	//
	// example:
	//
	// www.example.com
	QuicDomain *string `json:"QuicDomain,omitempty" xml:"QuicDomain,omitempty"`
	// The QUIC version. Default value: 0. Valid values:
	//
	// 	- 35
	//
	// 	- 39
	//
	// 	- 43
	//
	// 	- 44
	//
	// >  Only Chrome elements support this parameter.
	//
	// example:
	//
	// 0
	QuicVersion *int64 `json:"QuicVersion,omitempty" xml:"QuicVersion,omitempty"`
	// Indicates whether request headers are returned. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: The headers of base documents are returned.
	//
	// 	- 2: All headers are returned.
	//
	// example:
	//
	// 0
	RequestHeader *int64 `json:"RequestHeader,omitempty" xml:"RequestHeader,omitempty"`
	// The time threshold that is used to define a slow element. Unit: seconds.
	//
	// example:
	//
	// 30
	SlowElementThreshold *int64 `json:"SlowElementThreshold,omitempty" xml:"SlowElementThreshold,omitempty"`
	// The blacklist for string verification.
	//
	// example:
	//
	// test
	VerifyStringBlacklist *string `json:"VerifyStringBlacklist,omitempty" xml:"VerifyStringBlacklist,omitempty"`
	// The whitelist for string verification.
	//
	// example:
	//
	// test
	VerifyStringWhitelist *string `json:"VerifyStringWhitelist,omitempty" xml:"VerifyStringWhitelist,omitempty"`
	// The timeout period of waiting for the monitoring to complete.
	//
	// example:
	//
	// 20
	WaitCompletionTime *int64 `json:"WaitCompletionTime,omitempty" xml:"WaitCompletionTime,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailNav) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailNav) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetDnsHijackWhitelist() *string {
	return s.DnsHijackWhitelist
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetElementBlacklist() *string {
	return s.ElementBlacklist
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetExecuteActiveX() *int64 {
	return s.ExecuteActiveX
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetExecuteApplet() *int64 {
	return s.ExecuteApplet
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetExecuteScript() *int64 {
	return s.ExecuteScript
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetFilterInvalidIP() *int64 {
	return s.FilterInvalidIP
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetFlowHijackJumpTimes() *int64 {
	return s.FlowHijackJumpTimes
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetFlowHijackLogo() *string {
	return s.FlowHijackLogo
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavAutomaticScrolling() *int64 {
	return s.NavAutomaticScrolling
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavCustomHeader() *string {
	return s.NavCustomHeader
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavCustomHeaderContent() *string {
	return s.NavCustomHeaderContent
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavCustomHost() *int64 {
	return s.NavCustomHost
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavCustomHostIp() *string {
	return s.NavCustomHostIp
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavDisableCache() *int64 {
	return s.NavDisableCache
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavDisableCompression() *int64 {
	return s.NavDisableCompression
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavIgnoreCertificateError() *int64 {
	return s.NavIgnoreCertificateError
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavRedirect() *int64 {
	return s.NavRedirect
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetNavReturnElement() *int64 {
	return s.NavReturnElement
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetPageTampering() *string {
	return s.PageTampering
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetQuicDomain() *string {
	return s.QuicDomain
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetQuicVersion() *int64 {
	return s.QuicVersion
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetRequestHeader() *int64 {
	return s.RequestHeader
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetSlowElementThreshold() *int64 {
	return s.SlowElementThreshold
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetVerifyStringBlacklist() *string {
	return s.VerifyStringBlacklist
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetVerifyStringWhitelist() *string {
	return s.VerifyStringWhitelist
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) GetWaitCompletionTime() *int64 {
	return s.WaitCompletionTime
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetDnsHijackWhitelist(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.DnsHijackWhitelist = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetElementBlacklist(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.ElementBlacklist = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetExecuteActiveX(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.ExecuteActiveX = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetExecuteApplet(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.ExecuteApplet = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetExecuteScript(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.ExecuteScript = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetFilterInvalidIP(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.FilterInvalidIP = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetFlowHijackJumpTimes(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.FlowHijackJumpTimes = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetFlowHijackLogo(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.FlowHijackLogo = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetMonitorTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.MonitorTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavAutomaticScrolling(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavAutomaticScrolling = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavCustomHeader(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavCustomHeader = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavCustomHeaderContent(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavCustomHeaderContent = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavCustomHost(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavCustomHost = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavCustomHostIp(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavCustomHostIp = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavDisableCache(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavDisableCache = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavDisableCompression(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavDisableCompression = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavIgnoreCertificateError(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavIgnoreCertificateError = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavRedirect(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavRedirect = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetNavReturnElement(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.NavReturnElement = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetPageTampering(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.PageTampering = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetProcessName(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.ProcessName = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetQuicDomain(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.QuicDomain = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetQuicVersion(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.QuicVersion = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetRequestHeader(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.RequestHeader = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetSlowElementThreshold(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.SlowElementThreshold = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetVerifyStringBlacklist(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.VerifyStringBlacklist = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetVerifyStringWhitelist(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.VerifyStringWhitelist = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) SetWaitCompletionTime(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNav {
	s.WaitCompletionTime = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNav) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailNet struct {
	// Indicates whether the data is displayed in the DIG format. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	NetDigSwitch *int64 `json:"NetDigSwitch,omitempty" xml:"NetDigSwitch,omitempty"`
	// The NS server.
	//
	// example:
	//
	// 189.12.32.124
	NetDnsNs *string `json:"NetDnsNs,omitempty" xml:"NetDnsNs,omitempty"`
	// The DNS query method. Valid values:
	//
	// 	- 1: recursive
	//
	// 	- 2: iterative
	//
	// example:
	//
	// 1
	NetDnsQueryMethod *string `json:"NetDnsQueryMethod,omitempty" xml:"NetDnsQueryMethod,omitempty"`
	// The type of the DNS server. Valid values:
	//
	// 	- 0: ipv4
	//
	// 	- 1: ipv6
	//
	// 	- 2: A version is automatically selected.
	//
	// example:
	//
	// 0
	NetDnsServer *int64 `json:"NetDnsServer,omitempty" xml:"NetDnsServer,omitempty"`
	// Indicates whether DNS test is enabled. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	NetDnsSwitch *int64 `json:"NetDnsSwitch,omitempty" xml:"NetDnsSwitch,omitempty"`
	// The timeout period of DNS requests.
	//
	// example:
	//
	// 10
	NetDnsTimeout *string `json:"NetDnsTimeout,omitempty" xml:"NetDnsTimeout,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- 0 : ICMP
	//
	// 	- 1 : TCP
	//
	// example:
	//
	// 0
	NetIcmpActive *int64 `json:"NetIcmpActive,omitempty" xml:"NetIcmpActive,omitempty"`
	// Indicates whether packets are split. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	NetIcmpDataCut *int64 `json:"NetIcmpDataCut,omitempty" xml:"NetIcmpDataCut,omitempty"`
	// The interval at which the synthetic monitoring task is executed.
	//
	// example:
	//
	// 10
	NetIcmpInterval *int64 `json:"NetIcmpInterval,omitempty" xml:"NetIcmpInterval,omitempty"`
	// The number of packets.
	//
	// example:
	//
	// 10
	NetIcmpNum *int64 `json:"NetIcmpNum,omitempty" xml:"NetIcmpNum,omitempty"`
	// The packet size.
	//
	// example:
	//
	// 30
	NetIcmpSize *int64 `json:"NetIcmpSize,omitempty" xml:"NetIcmpSize,omitempty"`
	// Indicates whether ICMP test is enabled. Valid values:
	//
	// 	- 0: no.
	//
	// 	- 1: yes. If you set this parameter to 1, you must also set the Icmp parameter.
	//
	// example:
	//
	// 0
	NetIcmpSwitch *int64 `json:"NetIcmpSwitch,omitempty" xml:"NetIcmpSwitch,omitempty"`
	// The monitoring timeout period.
	//
	// example:
	//
	// 20
	NetIcmpTimeout *int64 `json:"NetIcmpTimeout,omitempty" xml:"NetIcmpTimeout,omitempty"`
	// The maximum number of active detection points.
	//
	// example:
	//
	// 30
	NetTraceRouteNum *int64 `json:"NetTraceRouteNum,omitempty" xml:"NetTraceRouteNum,omitempty"`
	// Indicates whether Tracert test is enabled. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes. If you set this parameter to 1, you must also set the Tracert parameter.
	//
	// example:
	//
	// 0
	NetTraceRouteSwitch *int64 `json:"NetTraceRouteSwitch,omitempty" xml:"NetTraceRouteSwitch,omitempty"`
	// The monitoring timeout period. Valid values: 0 to 300. Unit: seconds.
	//
	// example:
	//
	// 20
	NetTraceRouteTimeout *int64 `json:"NetTraceRouteTimeout,omitempty" xml:"NetTraceRouteTimeout,omitempty"`
	// The whitelisted objects that are used to avoid DNS hijacking. Format: `<domain name>:<objects>`.
	//
	// >  WAP networks do not support hijacking.
	//
	// example:
	//
	// 119.119.53.156/32
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailNet) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailNet) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetDigSwitch() *int64 {
	return s.NetDigSwitch
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetDnsNs() *string {
	return s.NetDnsNs
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetDnsQueryMethod() *string {
	return s.NetDnsQueryMethod
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetDnsServer() *int64 {
	return s.NetDnsServer
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetDnsSwitch() *int64 {
	return s.NetDnsSwitch
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetDnsTimeout() *string {
	return s.NetDnsTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpActive() *int64 {
	return s.NetIcmpActive
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpDataCut() *int64 {
	return s.NetIcmpDataCut
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpInterval() *int64 {
	return s.NetIcmpInterval
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpNum() *int64 {
	return s.NetIcmpNum
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpSize() *int64 {
	return s.NetIcmpSize
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpSwitch() *int64 {
	return s.NetIcmpSwitch
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetIcmpTimeout() *int64 {
	return s.NetIcmpTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetTraceRouteNum() *int64 {
	return s.NetTraceRouteNum
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetTraceRouteSwitch() *int64 {
	return s.NetTraceRouteSwitch
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetNetTraceRouteTimeout() *int64 {
	return s.NetTraceRouteTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) GetWhiteList() *string {
	return s.WhiteList
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetDigSwitch(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetDigSwitch = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetDnsNs(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetDnsNs = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetDnsQueryMethod(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetDnsQueryMethod = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetDnsServer(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetDnsServer = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetDnsSwitch(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetDnsSwitch = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetDnsTimeout(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetDnsTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpActive(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpActive = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpDataCut(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpDataCut = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpInterval(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpInterval = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpNum(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpNum = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpSize(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpSize = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpSwitch(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpSwitch = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetIcmpTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetIcmpTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetTraceRouteNum(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetTraceRouteNum = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetTraceRouteSwitch(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetTraceRouteSwitch = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetNetTraceRouteTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.NetTraceRouteTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) SetWhiteList(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailNet {
	s.WhiteList = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailNet) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailProtocol struct {
	// The encoding format. Valid values:
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
	// The custom host. Valid values:
	//
	// 	- 1: round robin
	//
	// 	- 0: random
	//
	// example:
	//
	// 1
	CustomHost *int64 `json:"CustomHost,omitempty" xml:"CustomHost,omitempty"`
	// The custom IP address of the host. Multiple IP addresses are separated with commas (,).
	//
	// example:
	//
	// 119.119.53.156
	CustomHostIp *string `json:"CustomHostIp,omitempty" xml:"CustomHostIp,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 20
	ProtocolConnectionTimeout *int64 `json:"ProtocolConnectionTimeout,omitempty" xml:"ProtocolConnectionTimeout,omitempty"`
	// The timeout period of API performance monitoring. Unit: seconds.
	//
	// example:
	//
	// 30
	ProtocolMonitorTimeout *int64 `json:"ProtocolMonitorTimeout,omitempty" xml:"ProtocolMonitorTimeout,omitempty"`
	// The size of the received data. This parameter is returned when **VerifyWay*	- is set to 2.
	//
	// example:
	//
	// 30
	ReceivedDataSize *int64 `json:"ReceivedDataSize,omitempty" xml:"ReceivedDataSize,omitempty"`
	// The request content, including the header and body.
	RequestContent *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent `json:"RequestContent,omitempty" xml:"RequestContent,omitempty" type:"Struct"`
	// The verification string.
	//
	// example:
	//
	// list
	VerifyContent *string `json:"VerifyContent,omitempty" xml:"VerifyContent,omitempty"`
	// The method that is used to verify the response content. Valid values:
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

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetCharacterEncoding() *int64 {
	return s.CharacterEncoding
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetCustomHost() *int64 {
	return s.CustomHost
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetCustomHostIp() *string {
	return s.CustomHostIp
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetProtocolConnectionTimeout() *int64 {
	return s.ProtocolConnectionTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetProtocolMonitorTimeout() *int64 {
	return s.ProtocolMonitorTimeout
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetReceivedDataSize() *int64 {
	return s.ReceivedDataSize
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetRequestContent() *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent {
	return s.RequestContent
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetVerifyContent() *string {
	return s.VerifyContent
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) GetVerifyWay() *int64 {
	return s.VerifyWay
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetCharacterEncoding(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.CharacterEncoding = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetCustomHost(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.CustomHost = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetCustomHostIp(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.CustomHostIp = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetProtocolConnectionTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.ProtocolConnectionTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetProtocolMonitorTimeout(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.ProtocolMonitorTimeout = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetReceivedDataSize(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.ReceivedDataSize = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetRequestContent(v *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.RequestContent = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetVerifyContent(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.VerifyContent = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) SetVerifyWay(v int64) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol {
	s.VerifyWay = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocol) Validate() error {
	if s.RequestContent != nil {
		if err := s.RequestContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent struct {
	// The content of the request body.
	Body *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The request header.
	Header []*GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
	// The request method. Valid values:
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

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) GetBody() *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody {
	return s.Body
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) GetHeader() []*GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader {
	return s.Header
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) GetMethod() *string {
	return s.Method
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) SetBody(v *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent {
	s.Body = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) SetHeader(v []*GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent {
	s.Header = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) SetMethod(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent {
	s.Method = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContent) Validate() error {
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

type GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody struct {
	// The data content. This parameter is returned when Mode is set to form-data.
	Formdata *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata `json:"Formdata,omitempty" xml:"Formdata,omitempty" type:"Struct"`
	// The language used when Mode is set to raw. Valid values:
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
	// xml
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The type of the content. Valid values:
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
	// The data content. This parameter is returned when **Mode*	- is set to **raw**.
	//
	// example:
	//
	// content
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The URL of the body content.
	Urlencoded *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded `json:"Urlencoded,omitempty" xml:"Urlencoded,omitempty" type:"Struct"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) GetFormdata() *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata {
	return s.Formdata
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) GetLanguage() *string {
	return s.Language
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) GetMode() *string {
	return s.Mode
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) GetRaw() *string {
	return s.Raw
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) GetUrlencoded() *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded {
	return s.Urlencoded
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) SetFormdata(v *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody {
	s.Formdata = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) SetLanguage(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody {
	s.Language = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) SetMode(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody {
	s.Mode = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) SetRaw(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody {
	s.Raw = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) SetUrlencoded(v *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody {
	s.Urlencoded = v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBody) Validate() error {
	if s.Formdata != nil {
		if err := s.Formdata.Validate(); err != nil {
			return err
		}
	}
	if s.Urlencoded != nil {
		if err := s.Urlencoded.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata struct {
	// The key of the **form-data**.
	//
	// example:
	//
	// appId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the form-data.
	//
	// example:
	//
	// 3425
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) GetKey() *string {
	return s.Key
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) GetValue() *string {
	return s.Value
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) SetKey(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata {
	s.Key = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) SetValue(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata {
	s.Value = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyFormdata) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded struct {
	// The tag key.
	//
	// example:
	//
	// appId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 11080
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) GetKey() *string {
	return s.Key
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) GetValue() *string {
	return s.Value
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) SetKey(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded {
	s.Key = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) SetValue(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded {
	s.Value = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentBodyUrlencoded) Validate() error {
	return dara.Validate(s)
}

type GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader struct {
	// The key of the header in the request parameters.
	//
	// example:
	//
	// regionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the header in the request parameters.
	//
	// example:
	//
	// cn-hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) GetKey() *string {
	return s.Key
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) GetValue() *string {
	return s.Value
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) SetKey(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader {
	s.Key = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) SetValue(v string) *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader {
	s.Value = &v
	return s
}

func (s *GetSyntheticTaskDetailResponseBodyTaskDetailProtocolRequestContentHeader) Validate() error {
	return dara.Validate(s)
}
