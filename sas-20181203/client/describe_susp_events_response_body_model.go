// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeSuspEventsResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeSuspEventsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSuspEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSuspEventsResponseBody
	GetRequestId() *string
	SetSuspEvents(v []*DescribeSuspEventsResponseBodySuspEvents) *DescribeSuspEventsResponseBody
	GetSuspEvents() []*DescribeSuspEventsResponseBodySuspEvents
	SetTotalCount(v int32) *DescribeSuspEventsResponseBody
	GetTotalCount() *int32
}

type DescribeSuspEventsResponseBody struct {
	// The number of entries on the current page in a paging query.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// 0D6E20E4-8326-1D03-A553-2182BE9E82F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security alert information.
	SuspEvents []*DescribeSuspEventsResponseBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
	// The total number of security alerts.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSuspEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSuspEventsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSuspEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSuspEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventsResponseBody) GetSuspEvents() []*DescribeSuspEventsResponseBodySuspEvents {
	return s.SuspEvents
}

func (s *DescribeSuspEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSuspEventsResponseBody) SetCount(v int32) *DescribeSuspEventsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetCurrentPage(v int32) *DescribeSuspEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetPageSize(v int32) *DescribeSuspEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetRequestId(v string) *DescribeSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetSuspEvents(v []*DescribeSuspEventsResponseBodySuspEvents) *DescribeSuspEventsResponseBody {
	s.SuspEvents = v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetTotalCount(v int32) *DescribeSuspEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) Validate() error {
	if s.SuspEvents != nil {
		for _, item := range s.SuspEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSuspEventsResponseBodySuspEvents struct {
	// Indicates whether this is an offline analysis alert.
	//
	// example:
	//
	// true
	Advanced *bool `json:"Advanced,omitempty" xml:"Advanced,omitempty"`
	// The name of the security alert.
	//
	// example:
	//
	// login_common_location
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	// The alert name.
	//
	// example:
	//
	// Login with unusual location
	AlarmEventNameDisplay *string `json:"AlarmEventNameDisplay,omitempty" xml:"AlarmEventNameDisplay,omitempty"`
	// The alerting event type.
	//
	// example:
	//
	// Unusual Logon
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// The alerting event type.
	//
	// example:
	//
	// Unusual Logon
	AlarmEventTypeDisplay *string `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	// The unique ID of the alert event.
	//
	// example:
	//
	// 8df914418f****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The name of the application to which the security alert belongs.
	//
	// example:
	//
	// pro-deploy-tibasic
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Indicates whether automatic defense is enabled.
	//
	// example:
	//
	// true
	AutoBreaking *bool `json:"AutoBreaking,omitempty" xml:"AutoBreaking,omitempty"`
	// Indicates whether the alert event can be handled online, such as quarantine. Valid values:
	//
	// - **true**: Online handling is supported.
	//
	// - **false**: Online handling is not supported.
	//
	// example:
	//
	// true
	CanBeDealOnLine *bool `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	// Indicates whether the false positive mark can be unmarked. Valid values:
	//
	// example:
	//
	// false
	CanCancelFault *bool `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	// Indicates whether critical event protection mode is enabled for the server. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// false
	ContainHwMode *bool `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	// The container ID.
	//
	// example:
	//
	// container_1648601865161_14925_02_000****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The container image ID.
	//
	// example:
	//
	// sha256:2e5a3b0ae5f452b3cb458789a9a7542ef40035a84318469a8528c5e444db1****
	ContainerImageId *string `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	// The container image name.
	//
	// example:
	//
	// centos7_apache:v1.0.1
	ContainerImageName *string `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	// The data source. You can ignore this parameter.
	//
	// example:
	//
	// aegis_suspicious_****
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The download URL of the file deep analysis report.
	DeepAnalyzeReportUrl *string `json:"DeepAnalyzeReportUrl,omitempty" xml:"DeepAnalyzeReportUrl,omitempty"`
	// The description of the impact of the security alert.
	//
	// example:
	//
	// webshell
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The details of the security alert.
	Details []*DescribeSuspEventsResponseBodySuspEventsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The discovery source. This parameter is invalid.
	//
	// example:
	//
	// linux
	DetectSource *string `json:"DetectSource,omitempty" xml:"DetectSource,omitempty"`
	// Indicates whether cloud sandbox detection is supported. Valid values:
	//
	// example:
	//
	// true
	DisplaySandboxResult *bool `json:"DisplaySandboxResult,omitempty" xml:"DisplaySandboxResult,omitempty"`
	// The remarks of the security alert.
	EventNotes []*DescribeSuspEventsResponseBodySuspEventsEventNotes `json:"EventNotes,omitempty" xml:"EventNotes,omitempty" type:"Repeated"`
	// The status of the security alert. Valid values:
	//
	// example:
	//
	// 8
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The subtype of the security alert.
	//
	// example:
	//
	// login_common_location
	EventSubType *string `json:"EventSubType,omitempty" xml:"EventSubType,omitempty"`
	// Indicates whether the security alert has tracing information. Valid values:
	//
	// example:
	//
	// true
	HasTraceInfo *bool `json:"HasTraceInfo,omitempty" xml:"HasTraceInfo,omitempty"`
	// The unique ID of the alert event.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image UUID.
	//
	// example:
	//
	// 70489fb520cea585ad9761d5a842****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The ID of the asset instance affected by the security alert.
	//
	// example:
	//
	// i-9dp6dwsxdl9z5u1e2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the associated instance.
	//
	// example:
	//
	// nginx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the associated instance.
	//
	// example:
	//
	// 1.2.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 100.100.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The Kubernetes cluster ID.
	//
	// example:
	//
	// c517b37e1401e4961b3951863a49a****
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// The name of the Kubernetes cluster.
	//
	// example:
	//
	// test-daily
	K8sClusterName *string `json:"K8sClusterName,omitempty" xml:"K8sClusterName,omitempty"`
	// The Kubernetes namespace.
	//
	// example:
	//
	// default
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// The Kubernetes node ID.
	//
	// example:
	//
	// i-bp14a1ay8e0aa9t0****
	K8sNodeId *string `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	// The Kubernetes node name.
	//
	// example:
	//
	// N/A
	K8sNodeName *string `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	// The Kubernetes pod name.
	//
	// example:
	//
	// myapp-pod
	K8sPodName *string `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	// Indicates whether large model analysis tags are supported. Valid values:
	//
	// - **true**: Supported.
	//
	// - **false**: Not supported.
	//
	// example:
	//
	// true
	LargeModel *bool `json:"LargeModel,omitempty" xml:"LargeModel,omitempty"`
	// The most recent time when the security alert occurred.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	LastTime *string `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The timestamp of the last occurrence, in milliseconds.
	//
	// example:
	//
	// 1631699497000
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The risk level of the alert event. Valid values:
	//
	// - **serious**: critical
	//
	// - **suspicious**: suspicious
	//
	// - **remind**: reminder
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The status of the malicious behavior defense rule. Valid values:
	//
	// example:
	//
	// open
	MaliciousRuleStatus *string `json:"MaliciousRuleStatus,omitempty" xml:"MaliciousRuleStatus,omitempty"`
	// The collection of security alert tags.
	MarkList []*string `json:"MarkList,omitempty" xml:"MarkList,omitempty" type:"Repeated"`
	// The advanced whitelisting rules.
	//
	// example:
	//
	// <strong>1.</strong>&nbsp&nbsppath&nbsp&nbspcontain&nbsp&nbsp232&nbsp&nbsp
	MarkMisRules *string `json:"MarkMisRules,omitempty" xml:"MarkMisRules,omitempty"`
	// The full name of the security alert.
	//
	// example:
	//
	// Unusual Logon-Login with unusual location
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the security alert first occurred.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// The timestamp of the first occurrence, in milliseconds.
	//
	// example:
	//
	// 1631699497000
	OccurrenceTimeStamp *int64 `json:"OccurrenceTimeStamp,omitempty" xml:"OccurrenceTimeStamp,omitempty"`
	// The error code of the security alert operation.
	//
	// example:
	//
	// kill_and_quara.Success
	OperateErrorCode *string `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	// The remarks of the security alert operation.
	//
	// example:
	//
	// success
	OperateMsg *string `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	// The timestamp of the security alert operation, in milliseconds.
	//
	// example:
	//
	// 1631699497000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The product edition that supports the detection of this security alert. Valid values:
	//
	// example:
	//
	// 1
	SaleVersion *string `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	// The IDs of the security alerts associated with this security alert.
	//
	// example:
	//
	// 270789
	SecurityEventIds *string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty"`
	// The Alibaba Cloud account ID that generated the alert.
	//
	// example:
	//
	// 196072141348****
	SourceAliUid *int64 `json:"SourceAliUid,omitempty" xml:"SourceAliUid,omitempty"`
	// The attack stage.
	//
	// example:
	//
	// "["authority_maintenance"]"
	Stages *string `json:"Stages,omitempty" xml:"Stages,omitempty"`
	// The operation types supported by the alert. Valid values:
	//
	// - **AI.false_positive**: suspected false positive.
	//
	// - **AI.real_attack**: real attack.
	//
	// - **AI.Insufficient_information_to_evaluate**: unable to determine.
	//
	// example:
	//
	// AI.real_attack
	SupportOperateCode *string `json:"SupportOperateCode,omitempty" xml:"SupportOperateCode,omitempty"`
	// The display name of the attack stage.
	TacticItems []*DescribeSuspEventsResponseBodySuspEventsTacticItems `json:"TacticItems,omitempty" xml:"TacticItems,omitempty" type:"Repeated"`
	// The unique key of the security alert.
	//
	// example:
	//
	// e17e****
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// The unique identifier of the associated instance.
	//
	// example:
	//
	// bf6b30d3-eea8-4924-9f0a-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c2051775877374cccbf68af596e6****
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAdvanced() *bool {
	return s.Advanced
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAlarmEventName() *string {
	return s.AlarmEventName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAlarmEventNameDisplay() *string {
	return s.AlarmEventNameDisplay
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAlarmEventType() *string {
	return s.AlarmEventType
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAlarmEventTypeDisplay() *string {
	return s.AlarmEventTypeDisplay
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAppName() *string {
	return s.AppName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetAutoBreaking() *bool {
	return s.AutoBreaking
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetCanBeDealOnLine() *bool {
	return s.CanBeDealOnLine
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetCanCancelFault() *bool {
	return s.CanCancelFault
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetContainHwMode() *bool {
	return s.ContainHwMode
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetContainerImageId() *string {
	return s.ContainerImageId
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetContainerImageName() *string {
	return s.ContainerImageName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetDataSource() *string {
	return s.DataSource
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetDeepAnalyzeReportUrl() *string {
	return s.DeepAnalyzeReportUrl
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetDesc() *string {
	return s.Desc
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetDetails() []*DescribeSuspEventsResponseBodySuspEventsDetails {
	return s.Details
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetDetectSource() *string {
	return s.DetectSource
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetDisplaySandboxResult() *bool {
	return s.DisplaySandboxResult
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetEventNotes() []*DescribeSuspEventsResponseBodySuspEventsEventNotes {
	return s.EventNotes
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetEventStatus() *int32 {
	return s.EventStatus
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetEventSubType() *string {
	return s.EventSubType
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetHasTraceInfo() *bool {
	return s.HasTraceInfo
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetId() *int64 {
	return s.Id
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetK8sClusterName() *string {
	return s.K8sClusterName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetK8sNodeId() *string {
	return s.K8sNodeId
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetK8sNodeName() *string {
	return s.K8sNodeName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetK8sPodName() *string {
	return s.K8sPodName
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetLargeModel() *bool {
	return s.LargeModel
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetLastTime() *string {
	return s.LastTime
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetLevel() *string {
	return s.Level
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetMaliciousRuleStatus() *string {
	return s.MaliciousRuleStatus
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetMarkList() []*string {
	return s.MarkList
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetMarkMisRules() *string {
	return s.MarkMisRules
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetName() *string {
	return s.Name
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetOccurrenceTimeStamp() *int64 {
	return s.OccurrenceTimeStamp
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetOperateErrorCode() *string {
	return s.OperateErrorCode
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetOperateMsg() *string {
	return s.OperateMsg
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetSaleVersion() *string {
	return s.SaleVersion
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetSecurityEventIds() *string {
	return s.SecurityEventIds
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetSourceAliUid() *int64 {
	return s.SourceAliUid
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetStages() *string {
	return s.Stages
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetSupportOperateCode() *string {
	return s.SupportOperateCode
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetTacticItems() []*DescribeSuspEventsResponseBodySuspEventsTacticItems {
	return s.TacticItems
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSuspEventsResponseBodySuspEvents) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAdvanced(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.Advanced = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventNameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmEventTypeDisplay(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAppName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.AppName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetAutoBreaking(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.AutoBreaking = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanBeDealOnLine(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetCanCancelFault(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainHwMode(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainerId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainerId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainerImageId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetContainerImageName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDataSource(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDeepAnalyzeReportUrl(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DeepAnalyzeReportUrl = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDesc(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Desc = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDetails(v []*DescribeSuspEventsResponseBodySuspEventsDetails) *DescribeSuspEventsResponseBodySuspEvents {
	s.Details = v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDetectSource(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.DetectSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetDisplaySandboxResult(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.DisplaySandboxResult = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventNotes(v []*DescribeSuspEventsResponseBodySuspEventsEventNotes) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventNotes = v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventStatus(v int32) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetEventSubType(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.EventSubType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetHasTraceInfo(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.HasTraceInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetId(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetImageUuid(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ImageUuid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetInstanceId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.InstanceId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetInstanceName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetInternetIp(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetIntranetIp(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sClusterId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sClusterName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sClusterName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sNamespace(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sNodeId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sNodeName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetK8sPodName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.K8sPodName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLargeModel(v bool) *DescribeSuspEventsResponseBodySuspEvents {
	s.LargeModel = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLastTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLastTimeStamp(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.LastTimeStamp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetLevel(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetMaliciousRuleStatus(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.MaliciousRuleStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetMarkList(v []*string) *DescribeSuspEventsResponseBodySuspEvents {
	s.MarkList = v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetMarkMisRules(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.MarkMisRules = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetName(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOccurrenceTime(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOccurrenceTimeStamp(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.OccurrenceTimeStamp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateErrorCode(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateMsg(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetOperateTime(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.OperateTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetSaleVersion(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.SaleVersion = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetSecurityEventIds(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.SecurityEventIds = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetSourceAliUid(v int64) *DescribeSuspEventsResponseBodySuspEvents {
	s.SourceAliUid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetStages(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Stages = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetSupportOperateCode(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.SupportOperateCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetTacticItems(v []*DescribeSuspEventsResponseBodySuspEventsTacticItems) *DescribeSuspEventsResponseBodySuspEvents {
	s.TacticItems = v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetUniqueInfo(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetUuid(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) SetClusterId(v string) *DescribeSuspEventsResponseBodySuspEvents {
	s.ClusterId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEvents) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EventNotes != nil {
		for _, item := range s.EventNotes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TacticItems != nil {
		for _, item := range s.TacticItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSuspEventsResponseBodySuspEventsDetails struct {
	// The display name of the security alert.
	//
	// example:
	//
	// Login with unusual location
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// The alerting event type.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The path where the security alert occurred.
	//
	// example:
	//
	// /etc/crontab
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The path where the security alert occurred.
	//
	// example:
	//
	// /etc/crontab
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEventsDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) GetType() *string {
	return s.Type
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) GetValue() *string {
	return s.Value
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetNameDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetType(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValue(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) SetValueDisplay(v string) *DescribeSuspEventsResponseBodySuspEventsDetails {
	s.ValueDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeSuspEventsResponseBodySuspEventsEventNotes struct {
	// The remark information.
	//
	// example:
	//
	// Test
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The event record ID.
	//
	// example:
	//
	// 123
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// The event record time.
	//
	// example:
	//
	// 2018-09-26 01:51:01
	NoteTime *string `json:"NoteTime,omitempty" xml:"NoteTime,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEventsEventNotes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEventsEventNotes) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) GetNote() *string {
	return s.Note
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) GetNoteId() *int64 {
	return s.NoteId
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) GetNoteTime() *string {
	return s.NoteTime
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) SetNote(v string) *DescribeSuspEventsResponseBodySuspEventsEventNotes {
	s.Note = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) SetNoteId(v int64) *DescribeSuspEventsResponseBodySuspEventsEventNotes {
	s.NoteId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) SetNoteTime(v string) *DescribeSuspEventsResponseBodySuspEventsEventNotes {
	s.NoteTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsEventNotes) Validate() error {
	return dara.Validate(s)
}

type DescribeSuspEventsResponseBodySuspEventsTacticItems struct {
	// The ATT&CK tactic name.
	//
	// example:
	//
	// Malicious scripts-Malicious script code execution
	TacticDisplayName *string `json:"TacticDisplayName,omitempty" xml:"TacticDisplayName,omitempty"`
	// The ATT&CK attack stage information.
	//
	// example:
	//
	// TA0001
	TacticId *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
}

func (s DescribeSuspEventsResponseBodySuspEventsTacticItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodySuspEventsTacticItems) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodySuspEventsTacticItems) GetTacticDisplayName() *string {
	return s.TacticDisplayName
}

func (s *DescribeSuspEventsResponseBodySuspEventsTacticItems) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeSuspEventsResponseBodySuspEventsTacticItems) SetTacticDisplayName(v string) *DescribeSuspEventsResponseBodySuspEventsTacticItems {
	s.TacticDisplayName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsTacticItems) SetTacticId(v string) *DescribeSuspEventsResponseBodySuspEventsTacticItems {
	s.TacticId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodySuspEventsTacticItems) Validate() error {
	return dara.Validate(s)
}
