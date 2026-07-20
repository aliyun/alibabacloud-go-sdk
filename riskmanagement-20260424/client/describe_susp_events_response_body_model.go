// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSuspEventsResponseBody
	GetCode() *string
	SetData(v *DescribeSuspEventsResponseBodyData) *DescribeSuspEventsResponseBody
	GetData() *DescribeSuspEventsResponseBodyData
	SetMessage(v string) *DescribeSuspEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSuspEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSuspEventsResponseBody
	GetSuccess() *bool
}

type DescribeSuspEventsResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeSuspEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 67BD8435-6624-5484-A75D-170231B51615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSuspEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSuspEventsResponseBody) GetData() *DescribeSuspEventsResponseBodyData {
	return s.Data
}

func (s *DescribeSuspEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSuspEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSuspEventsResponseBody) SetCode(v string) *DescribeSuspEventsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetData(v *DescribeSuspEventsResponseBodyData) *DescribeSuspEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetMessage(v string) *DescribeSuspEventsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetRequestId(v string) *DescribeSuspEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) SetSuccess(v bool) *DescribeSuspEventsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSuspEventsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSuspEventsResponseBodyData struct {
	Body *DescribeSuspEventsResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s DescribeSuspEventsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodyData) GetBody() *DescribeSuspEventsResponseBodyDataBody {
	return s.Body
}

func (s *DescribeSuspEventsResponseBodyData) SetBody(v *DescribeSuspEventsResponseBodyDataBody) *DescribeSuspEventsResponseBodyData {
	s.Body = v
	return s
}

func (s *DescribeSuspEventsResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSuspEventsResponseBodyDataBody struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AD2345D1-A498-58AF-97C0-88940AF87CB7
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuspEvents []*DescribeSuspEventsResponseBodyDataBodySuspEvents `json:"SuspEvents,omitempty" xml:"SuspEvents,omitempty" type:"Repeated"`
	// example:
	//
	// 72
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSuspEventsResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodyDataBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSuspEventsResponseBodyDataBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSuspEventsResponseBodyDataBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSuspEventsResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventsResponseBodyDataBody) GetSuspEvents() []*DescribeSuspEventsResponseBodyDataBodySuspEvents {
	return s.SuspEvents
}

func (s *DescribeSuspEventsResponseBodyDataBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSuspEventsResponseBodyDataBody) SetCount(v int32) *DescribeSuspEventsResponseBodyDataBody {
	s.Count = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBody) SetCurrentPage(v int32) *DescribeSuspEventsResponseBodyDataBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBody) SetPageSize(v int32) *DescribeSuspEventsResponseBodyDataBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBody) SetRequestId(v string) *DescribeSuspEventsResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBody) SetSuspEvents(v []*DescribeSuspEventsResponseBodyDataBodySuspEvents) *DescribeSuspEventsResponseBodyDataBody {
	s.SuspEvents = v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBody) SetTotalCount(v int32) *DescribeSuspEventsResponseBodyDataBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBody) Validate() error {
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

type DescribeSuspEventsResponseBodyDataBodySuspEvents struct {
	// example:
	//
	// true
	Advanced *bool `json:"Advanced,omitempty" xml:"Advanced,omitempty"`
	// example:
	//
	// 反弹shell_拦截
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	// example:
	//
	// Login with unusual location
	AlarmEventNameDisplay *string `json:"AlarmEventNameDisplay,omitempty" xml:"AlarmEventNameDisplay,omitempty"`
	// example:
	//
	// Unusual Logon
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// example:
	//
	// Unusual Logon
	AlarmEventTypeDisplay *string `json:"AlarmEventTypeDisplay,omitempty" xml:"AlarmEventTypeDisplay,omitempty"`
	// example:
	//
	// 8df914418f****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// example:
	//
	// dfield-cloud-service-prod
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// true
	AutoBreaking *bool `json:"AutoBreaking,omitempty" xml:"AutoBreaking,omitempty"`
	// example:
	//
	// true
	CanBeDealOnLine *bool `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	// example:
	//
	// true
	CanCancelFault *bool `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	// example:
	//
	// c8c87dae64c9947269091f36cfa9adc87
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	ContainHwMode *bool `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	// example:
	//
	// 95878ef8779fae3dd82126812edd910402fc550a72f9bce87e56a4435d018384
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// example:
	//
	// sha256:2e5a3b0ae5f452b3cb458789a9a7542ef40035a84318469a8528c5e444db1****
	ContainerImageId *string `json:"ContainerImageId,omitempty" xml:"ContainerImageId,omitempty"`
	// example:
	//
	// centos7_apache:v1.0.1
	ContainerImageName *string `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	// example:
	//
	// URL
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// example:
	//
	// webshell
	Desc    *string                                                    `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Details []*DescribeSuspEventsResponseBodyDataBodySuspEventsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// -
	DetectSource *string `json:"DetectSource,omitempty" xml:"DetectSource,omitempty"`
	// example:
	//
	// true
	DisplaySandboxResult *bool                                                         `json:"DisplaySandboxResult,omitempty" xml:"DisplaySandboxResult,omitempty"`
	EventNotes           []*DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes `json:"EventNotes,omitempty" xml:"EventNotes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// example:
	//
	// login_common_location
	EventSubType *string `json:"EventSubType,omitempty" xml:"EventSubType,omitempty"`
	// example:
	//
	// true
	HasTraceInfo *bool `json:"HasTraceInfo,omitempty" xml:"HasTraceInfo,omitempty"`
	// example:
	//
	// 3178
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ccdab289-9765-47ef-af50-ba6be09aacd6
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// example:
	//
	// i-9dp6dwsxdl9z5u1e2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// nginx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 8.137.3*.6
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 10.36.*6.149
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// ce3c41ed427794a7bb3d9da4554fc8039
	K8sClusterId *string `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	// example:
	//
	// testName
	K8sClusterName *string `json:"K8sClusterName,omitempty" xml:"K8sClusterName,omitempty"`
	// example:
	//
	// default
	K8sNamespace *string `json:"K8sNamespace,omitempty" xml:"K8sNamespace,omitempty"`
	// example:
	//
	// i-bp14a1ay8e0aa9t0****
	K8sNodeId *string `json:"K8sNodeId,omitempty" xml:"K8sNodeId,omitempty"`
	// example:
	//
	// N/A
	K8sNodeName *string `json:"K8sNodeName,omitempty" xml:"K8sNodeName,omitempty"`
	// example:
	//
	// myapp-pod
	K8sPodName *string `json:"K8sPodName,omitempty" xml:"K8sPodName,omitempty"`
	// example:
	//
	// true
	LargeModel *bool `json:"LargeModel,omitempty" xml:"LargeModel,omitempty"`
	// example:
	//
	// 2018-09-26 01:51:01
	LastTime *string `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// 1631699497000
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// example:
	//
	// remind
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// open
	MaliciousRuleStatus *string   `json:"MaliciousRuleStatus,omitempty" xml:"MaliciousRuleStatus,omitempty"`
	MarkList            []*string `json:"MarkList,omitempty" xml:"MarkList,omitempty" type:"Repeated"`
	// example:
	//
	// <strong>1.</strong>&nbsp&nbsppath&nbsp&nbspcontain&nbsp&nbsp232&nbsp&nbsp
	MarkMisRules *string `json:"MarkMisRules,omitempty" xml:"MarkMisRules,omitempty"`
	// example:
	//
	// Unusual Logon-Login with unusual location
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2018-09-26 01:51:01
	OccurrenceTime *string `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// example:
	//
	// 1631699497000
	OccurrenceTimeStamp *int64 `json:"OccurrenceTimeStamp,omitempty" xml:"OccurrenceTimeStamp,omitempty"`
	// example:
	//
	// kill_and_quara.Success
	OperateErrorCode *string `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	// example:
	//
	// success
	OperateMsg *string `json:"OperateMsg,omitempty" xml:"OperateMsg,omitempty"`
	// example:
	//
	// 1631699497000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// example:
	//
	// 1
	SaleVersion *string `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	// example:
	//
	// 628978308
	SecurityEventIds *string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty"`
	// example:
	//
	// 124075**67406
	SourceAliUid *int64 `json:"SourceAliUid,omitempty" xml:"SourceAliUid,omitempty"`
	// example:
	//
	// "["authority_maintenance"]"
	Stages *string `json:"Stages,omitempty" xml:"Stages,omitempty"`
	// example:
	//
	// AI.false_positive
	SupportOperateCode *string                                                        `json:"SupportOperateCode,omitempty" xml:"SupportOperateCode,omitempty"`
	TacticItems        []*DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems `json:"TacticItems,omitempty" xml:"TacticItems,omitempty" type:"Repeated"`
	// example:
	//
	// 1dfbdf56c5343b63c4854d08ec20e067
	UniqueInfo *string `json:"UniqueInfo,omitempty" xml:"UniqueInfo,omitempty"`
	// example:
	//
	// 9A75F21D3993C0A2B094A4AB132890B2
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEvents) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAdvanced() *bool {
	return s.Advanced
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAlarmEventName() *string {
	return s.AlarmEventName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAlarmEventNameDisplay() *string {
	return s.AlarmEventNameDisplay
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAlarmEventType() *string {
	return s.AlarmEventType
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAlarmEventTypeDisplay() *string {
	return s.AlarmEventTypeDisplay
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAppName() *string {
	return s.AppName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetAutoBreaking() *bool {
	return s.AutoBreaking
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetCanBeDealOnLine() *bool {
	return s.CanBeDealOnLine
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetCanCancelFault() *bool {
	return s.CanCancelFault
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetContainHwMode() *bool {
	return s.ContainHwMode
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetContainerImageId() *string {
	return s.ContainerImageId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetContainerImageName() *string {
	return s.ContainerImageName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetDataSource() *string {
	return s.DataSource
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetDesc() *string {
	return s.Desc
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetDetails() []*DescribeSuspEventsResponseBodyDataBodySuspEventsDetails {
	return s.Details
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetDetectSource() *string {
	return s.DetectSource
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetDisplaySandboxResult() *bool {
	return s.DisplaySandboxResult
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetEventNotes() []*DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes {
	return s.EventNotes
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetEventStatus() *int32 {
	return s.EventStatus
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetEventSubType() *string {
	return s.EventSubType
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetHasTraceInfo() *bool {
	return s.HasTraceInfo
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetId() *int64 {
	return s.Id
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetK8sClusterId() *string {
	return s.K8sClusterId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetK8sClusterName() *string {
	return s.K8sClusterName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetK8sNamespace() *string {
	return s.K8sNamespace
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetK8sNodeId() *string {
	return s.K8sNodeId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetK8sNodeName() *string {
	return s.K8sNodeName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetK8sPodName() *string {
	return s.K8sPodName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetLargeModel() *bool {
	return s.LargeModel
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetLastTime() *string {
	return s.LastTime
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetLevel() *string {
	return s.Level
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetMaliciousRuleStatus() *string {
	return s.MaliciousRuleStatus
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetMarkList() []*string {
	return s.MarkList
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetMarkMisRules() *string {
	return s.MarkMisRules
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetName() *string {
	return s.Name
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetOccurrenceTime() *string {
	return s.OccurrenceTime
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetOccurrenceTimeStamp() *int64 {
	return s.OccurrenceTimeStamp
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetOperateErrorCode() *string {
	return s.OperateErrorCode
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetOperateMsg() *string {
	return s.OperateMsg
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetSaleVersion() *string {
	return s.SaleVersion
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetSecurityEventIds() *string {
	return s.SecurityEventIds
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetSourceAliUid() *int64 {
	return s.SourceAliUid
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetStages() *string {
	return s.Stages
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetSupportOperateCode() *string {
	return s.SupportOperateCode
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetTacticItems() []*DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems {
	return s.TacticItems
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetUniqueInfo() *string {
	return s.UniqueInfo
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAdvanced(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Advanced = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAlarmEventName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAlarmEventNameDisplay(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AlarmEventNameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAlarmEventType(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAlarmEventTypeDisplay(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AlarmEventTypeDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAlarmUniqueInfo(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAppName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AppName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetAutoBreaking(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.AutoBreaking = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetCanBeDealOnLine(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetCanCancelFault(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetClusterId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.ClusterId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetContainHwMode(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetContainerId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.ContainerId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetContainerImageId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.ContainerImageId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetContainerImageName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.ContainerImageName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetDataSource(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.DataSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetDesc(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Desc = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetDetails(v []*DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Details = v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetDetectSource(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.DetectSource = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetDisplaySandboxResult(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.DisplaySandboxResult = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetEventNotes(v []*DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.EventNotes = v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetEventStatus(v int32) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetEventSubType(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.EventSubType = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetHasTraceInfo(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.HasTraceInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetId(v int64) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetImageUuid(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.ImageUuid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetInstanceId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.InstanceId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetInstanceName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetInternetIp(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetIntranetIp(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetK8sClusterId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetK8sClusterName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.K8sClusterName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetK8sNamespace(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.K8sNamespace = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetK8sNodeId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.K8sNodeId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetK8sNodeName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.K8sNodeName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetK8sPodName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.K8sPodName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetLargeModel(v bool) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.LargeModel = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetLastTime(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.LastTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetLastTimeStamp(v int64) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.LastTimeStamp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetLevel(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Level = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetMaliciousRuleStatus(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.MaliciousRuleStatus = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetMarkList(v []*string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.MarkList = v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetMarkMisRules(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.MarkMisRules = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Name = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetOccurrenceTime(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetOccurrenceTimeStamp(v int64) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.OccurrenceTimeStamp = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetOperateErrorCode(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetOperateMsg(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.OperateMsg = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetOperateTime(v int64) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.OperateTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetSaleVersion(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.SaleVersion = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetSecurityEventIds(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.SecurityEventIds = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetSourceAliUid(v int64) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.SourceAliUid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetStages(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Stages = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetSupportOperateCode(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.SupportOperateCode = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetTacticItems(v []*DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.TacticItems = v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetUniqueInfo(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.UniqueInfo = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) SetUuid(v string) *DescribeSuspEventsResponseBodyDataBodySuspEvents {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEvents) Validate() error {
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

type DescribeSuspEventsResponseBodyDataBodySuspEventsDetails struct {
	// example:
	//
	// login with unusual location
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// /etc/crontab
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// /etc/crontab
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) GetType() *string {
	return s.Type
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) GetValue() *string {
	return s.Value
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) SetNameDisplay(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails {
	s.NameDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) SetType(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails {
	s.Type = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) SetValue(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails {
	s.Value = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) SetValueDisplay(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails {
	s.ValueDisplay = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes struct {
	// example:
	//
	// test
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// example:
	//
	// 2859481
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// example:
	//
	// 2018-09-26 01:51:01
	NoteTime *string `json:"NoteTime,omitempty" xml:"NoteTime,omitempty"`
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) GetNote() *string {
	return s.Note
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) GetNoteId() *int64 {
	return s.NoteId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) GetNoteTime() *string {
	return s.NoteTime
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) SetNote(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes {
	s.Note = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) SetNoteId(v int64) *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes {
	s.NoteId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) SetNoteTime(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes {
	s.NoteTime = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsEventNotes) Validate() error {
	return dara.Validate(s)
}

type DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems struct {
	// example:
	//
	// Malicious scripts-Malicious script code execution
	TacticDisplayName *string `json:"TacticDisplayName,omitempty" xml:"TacticDisplayName,omitempty"`
	// example:
	//
	// TA0042
	TacticId *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) GetTacticDisplayName() *string {
	return s.TacticDisplayName
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) GetTacticId() *string {
	return s.TacticId
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) SetTacticDisplayName(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems {
	s.TacticDisplayName = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) SetTacticId(v string) *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems {
	s.TacticId = &v
	return s
}

func (s *DescribeSuspEventsResponseBodyDataBodySuspEventsTacticItems) Validate() error {
	return dara.Validate(s)
}
