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
	Message        *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
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
	AgentInstallErrorCode *string `json:"AgentInstallErrorCode,omitempty" xml:"AgentInstallErrorCode,omitempty"`
	AutoInstall           *bool   `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoongCollectorStatus  *string `json:"LoongCollectorStatus,omitempty" xml:"LoongCollectorStatus,omitempty"`
	LoongCollectorVersion *string `json:"LoongCollectorVersion,omitempty" xml:"LoongCollectorVersion,omitempty"`
	OsMonitorConfig       *string `json:"OsMonitorConfig,omitempty" xml:"OsMonitorConfig,omitempty"`
	OsMonitorErrorCode    *string `json:"OsMonitorErrorCode,omitempty" xml:"OsMonitorErrorCode,omitempty"`
	OsMonitorErrorDetail  *string `json:"OsMonitorErrorDetail,omitempty" xml:"OsMonitorErrorDetail,omitempty"`
	OsMonitorStatus       *string `json:"OsMonitorStatus,omitempty" xml:"OsMonitorStatus,omitempty"`
	OsMonitorVersion      *string `json:"OsMonitorVersion,omitempty" xml:"OsMonitorVersion,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
