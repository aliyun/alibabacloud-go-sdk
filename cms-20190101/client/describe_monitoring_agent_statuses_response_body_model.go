// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMonitoringAgentStatusesResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMonitoringAgentStatusesResponseBody
	GetMessage() *string
	SetNodeStatusList(v *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) *DescribeMonitoringAgentStatusesResponseBody
	GetNodeStatusList() *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList
	SetRequestId(v string) *DescribeMonitoringAgentStatusesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitoringAgentStatusesResponseBody
	GetSuccess() *bool
}

type DescribeMonitoringAgentStatusesResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The host status information.
	NodeStatusList *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList `json:"NodeStatusList,omitempty" xml:"NodeStatusList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6F8371DF-AB81-41B8-9E1B-5493B3FF0E4F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitoringAgentStatusesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitoringAgentStatusesResponseBody) GetNodeStatusList() *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList {
	return s.NodeStatusList
}

func (s *DescribeMonitoringAgentStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitoringAgentStatusesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetCode(v string) *DescribeMonitoringAgentStatusesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetMessage(v string) *DescribeMonitoringAgentStatusesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetNodeStatusList(v *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) *DescribeMonitoringAgentStatusesResponseBody {
	s.NodeStatusList = v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetRequestId(v string) *DescribeMonitoringAgentStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentStatusesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBody) Validate() error {
	if s.NodeStatusList != nil {
		if err := s.NodeStatusList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitoringAgentStatusesResponseBodyNodeStatusList struct {
	NodeStatus []*DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty" type:"Repeated"`
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) GetNodeStatus() []*DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	return s.NodeStatus
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) SetNodeStatus(v []*DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList {
	s.NodeStatus = v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusList) Validate() error {
	if s.NodeStatus != nil {
		for _, item := range s.NodeStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus struct {
	// The error code returned when the CloudMonitor agent is installed. Valid values:
	//
	// 	- Common.Timeout: The installation timed out.
	//
	// 	- Common.SLR: The service-linked role for CloudMonitor is unauthorized.
	//
	// 	- Common.OS: The operating system is not supported.
	//
	// 	- Assist.Invalid: Cloud Assistant is not running.
	//
	// 	- Assist.Invoke: An error occurred when the installation program is started.
	//
	// 	- Assist.Execute: An error occurred when the installation program is running.
	AgentInstallErrorCode *string `json:"AgentInstallErrorCode,omitempty" xml:"AgentInstallErrorCode,omitempty"`
	// Indicates whether the CloudMonitor agent is automatically installed. Valid values:
	//
	// 	- true: The CloudMonitor agent is automatically installed.
	//
	// 	- false: The CloudMonitor agent is not automatically installed.
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-hp3dunahluwajv6f****
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoongCollectorStatus  *string `json:"LoongCollectorStatus,omitempty" xml:"LoongCollectorStatus,omitempty"`
	LoongCollectorVersion *string `json:"LoongCollectorVersion,omitempty" xml:"LoongCollectorVersion,omitempty"`
	// Indicates whether the SysAK monitoring feature is enabled.`` Valid values:
	//
	// 	- `true`: The SysAK monitoring feature is enabled.
	//
	// 	- `false`: the SysAK monitoring feature is disabled.
	//
	// example:
	//
	// {"sysak":true}
	OsMonitorConfig *string `json:"OsMonitorConfig,omitempty" xml:"OsMonitorConfig,omitempty"`
	// The error status of SysOM. Valid values:
	//
	// 	- `install_fail`: SysOM fails to be installed or an unknown error occurs.
	//
	// 	- `install_assist_invalid`: SysOM fails to be installed because the status of Cloud Assistant is invalid.
	//
	// 	- `install_assist_command_fail`: SysOM fails to be installed because the installation command fails to run.
	//
	// 	- `uninstall_fail`: SysOM fails to be uninstalled or an unknown error occurs.
	//
	// 	- `uninstall_assist_invalid`: SysOM fails to be uninstalled because the status of Cloud Assistant is invalid.
	//
	// 	- `uninstall_assist_command_fail`: SysOM fails to be uninstalled because the uninstallation command fails to run.
	//
	// example:
	//
	// install_fail
	OsMonitorErrorCode *string `json:"OsMonitorErrorCode,omitempty" xml:"OsMonitorErrorCode,omitempty"`
	// The details of the execution error. Valid values:
	//
	// 	- `Command.ErrorCode.Fail.Downlaod.REGIN_ID`: Failed to obtain the region ID.
	//
	// 	- `Command.ErrorCode.Fail.Downlaod.SYSAK`: Failed to download the .rpm package of System Analyse Kit (SysAK).
	//
	// 	- `Command.ErrorCode.Fail.Downlaod.CMON_FILE`: Failed to download the CMON file.
	//
	// 	- `Command.ErrorCode.Fail.Downlaod.BTF`: Failed to start SysAK because the BTF file is not found.
	//
	// 	- `Command.ErrorCode.Fail.Start.SYSAK`: Failed to start SysAK due to an unknown error.
	//
	// example:
	//
	// Command.ErrorCode.Fail.Downlaod.REGIN_ID
	OsMonitorErrorDetail *string `json:"OsMonitorErrorDetail,omitempty" xml:"OsMonitorErrorDetail,omitempty"`
	// The status of SysOM. Valid values:
	//
	// 	- installing: SysOM is being installed.
	//
	// 	- running: SysOM is running.
	//
	// 	- stopped: SysOM is stopped.
	//
	// 	- uninstalling: SysOM is being uninstalled.
	//
	// example:
	//
	// running
	OsMonitorStatus *string `json:"OsMonitorStatus,omitempty" xml:"OsMonitorStatus,omitempty"`
	// The SysOM version.
	//
	// example:
	//
	// 1.3.0-12
	OsMonitorVersion *string `json:"OsMonitorVersion,omitempty" xml:"OsMonitorVersion,omitempty"`
	// The status of the CloudMonitor agent. Valid values:
	//
	// 	- running: The CloudMonitor agent is running.
	//
	// 	- stopped: The CloudMonitor agent is stopped.
	//
	// 	- installing: The CloudMonitor agent is being installed.
	//
	// 	- install_faild: The CloudMonitor agent fails to be installed.
	//
	// 	- abnormal: The CloudMonitor agent is not properly installed.
	//
	// 	- not_installed: The CloudMonitor agent is not installed.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetAgentInstallErrorCode() *string {
	return s.AgentInstallErrorCode
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetLoongCollectorStatus() *string {
	return s.LoongCollectorStatus
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetLoongCollectorVersion() *string {
	return s.LoongCollectorVersion
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetOsMonitorConfig() *string {
	return s.OsMonitorConfig
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetOsMonitorErrorCode() *string {
	return s.OsMonitorErrorCode
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetOsMonitorErrorDetail() *string {
	return s.OsMonitorErrorDetail
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetOsMonitorStatus() *string {
	return s.OsMonitorStatus
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetOsMonitorVersion() *string {
	return s.OsMonitorVersion
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetAgentInstallErrorCode(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.AgentInstallErrorCode = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetAutoInstall(v bool) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.AutoInstall = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetInstanceId(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetLoongCollectorStatus(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.LoongCollectorStatus = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetLoongCollectorVersion(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.LoongCollectorVersion = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetOsMonitorConfig(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.OsMonitorConfig = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetOsMonitorErrorCode(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.OsMonitorErrorCode = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetOsMonitorErrorDetail(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.OsMonitorErrorDetail = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetOsMonitorStatus(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.OsMonitorStatus = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetOsMonitorVersion(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.OsMonitorVersion = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) SetStatus(v string) *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus {
	s.Status = &v
	return s
}

func (s *DescribeMonitoringAgentStatusesResponseBodyNodeStatusListNodeStatus) Validate() error {
	return dara.Validate(s)
}
