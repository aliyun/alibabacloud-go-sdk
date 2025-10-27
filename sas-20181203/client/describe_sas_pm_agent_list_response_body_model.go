// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSasPmAgentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSasPmAgentListResponseBody
	GetRequestId() *string
	SetSasPmAgentList(v []*DescribeSasPmAgentListResponseBodySasPmAgentList) *DescribeSasPmAgentListResponseBody
	GetSasPmAgentList() []*DescribeSasPmAgentListResponseBodySasPmAgentList
}

type DescribeSasPmAgentListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8EF3ACC2-9400-5B64-B72D-4A1D35113750
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the information about servers.
	SasPmAgentList []*DescribeSasPmAgentListResponseBodySasPmAgentList `json:"SasPmAgentList,omitempty" xml:"SasPmAgentList,omitempty" type:"Repeated"`
}

func (s DescribeSasPmAgentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSasPmAgentListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSasPmAgentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSasPmAgentListResponseBody) GetSasPmAgentList() []*DescribeSasPmAgentListResponseBodySasPmAgentList {
	return s.SasPmAgentList
}

func (s *DescribeSasPmAgentListResponseBody) SetRequestId(v string) *DescribeSasPmAgentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBody) SetSasPmAgentList(v []*DescribeSasPmAgentListResponseBodySasPmAgentList) *DescribeSasPmAgentListResponseBody {
	s.SasPmAgentList = v
	return s
}

func (s *DescribeSasPmAgentListResponseBody) Validate() error {
	if s.SasPmAgentList != nil {
		for _, item := range s.SasPmAgentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSasPmAgentListResponseBodySasPmAgentList struct {
	// The ID of Cloud Assistant.
	//
	// example:
	//
	// mi-hz034jn***yxhc0
	AliyunAssistId *string `json:"AliyunAssistId,omitempty" xml:"AliyunAssistId,omitempty"`
	// The ID of the CloudMonitor agent.
	//
	// example:
	//
	// 5d5ef6be-54ff-11ed-82cf-8f01475e****
	AliyunMonitorId *string `json:"AliyunMonitorId,omitempty" xml:"AliyunMonitorId,omitempty"`
	// The installation result of Cloud Assistant. Valid values:
	//
	// 	- **0**: SUCCESS
	//
	// 	- **1**: MISSING_PARAM
	//
	// 	- **2**: UNKNOWN_SYSTEM
	//
	// 	- **3**: DOWNLOAD_FAILED
	//
	// 	- **4**: INSTALL_FAILED
	//
	// example:
	//
	// 0
	AssistInstallResult *int32 `json:"AssistInstallResult,omitempty" xml:"AssistInstallResult,omitempty"`
	// The status of Cloud Assistant. Valid values:
	//
	// 	- **0**: installing
	//
	// 	- **1**: installed
	//
	// 	- **2**: installation failed
	//
	// 	- **3**: installation timed out
	//
	// example:
	//
	// 1
	AssistInstallStatus *int32 `json:"AssistInstallStatus,omitempty" xml:"AssistInstallStatus,omitempty"`
	// The installation result of the CloudMonitor agent. Valid values:
	//
	// 	- **0**: failed
	//
	// 	- **1**: successful
	//
	// example:
	//
	// 1
	MonitorInstallResult *int32 `json:"MonitorInstallResult,omitempty" xml:"MonitorInstallResult,omitempty"`
	// The status of the CloudMonitor agent. Valid values:
	//
	// 	- **0**: installation failed
	//
	// 	- **1**: installed
	//
	// example:
	//
	// 1
	MonitorInstallStatus *int32 `json:"MonitorInstallStatus,omitempty" xml:"MonitorInstallStatus,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 87f1724d-075e-48d3-95fd-78c2dd36****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSasPmAgentListResponseBodySasPmAgentList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSasPmAgentListResponseBodySasPmAgentList) GoString() string {
	return s.String()
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetAliyunAssistId() *string {
	return s.AliyunAssistId
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetAliyunMonitorId() *string {
	return s.AliyunMonitorId
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetAssistInstallResult() *int32 {
	return s.AssistInstallResult
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetAssistInstallStatus() *int32 {
	return s.AssistInstallStatus
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetMonitorInstallResult() *int32 {
	return s.MonitorInstallResult
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetMonitorInstallStatus() *int32 {
	return s.MonitorInstallStatus
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetAliyunAssistId(v string) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.AliyunAssistId = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetAliyunMonitorId(v string) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.AliyunMonitorId = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetAssistInstallResult(v int32) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.AssistInstallResult = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetAssistInstallStatus(v int32) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.AssistInstallStatus = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetMonitorInstallResult(v int32) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.MonitorInstallResult = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetMonitorInstallStatus(v int32) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.MonitorInstallStatus = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) SetUuid(v string) *DescribeSasPmAgentListResponseBodySasPmAgentList {
	s.Uuid = &v
	return s
}

func (s *DescribeSasPmAgentListResponseBodySasPmAgentList) Validate() error {
	return dara.Validate(s)
}
