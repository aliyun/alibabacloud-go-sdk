// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAlertMessagesResponseBodyData) *ListAlertMessagesResponseBody
	GetData() *ListAlertMessagesResponseBodyData
	SetErrorCode(v string) *ListAlertMessagesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAlertMessagesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListAlertMessagesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListAlertMessagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAlertMessagesResponseBody
	GetSuccess() *bool
}

type ListAlertMessagesResponseBody struct {
	// The information about returned alerts.
	Data *ListAlertMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlertMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBody) GetData() *ListAlertMessagesResponseBodyData {
	return s.Data
}

func (s *ListAlertMessagesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAlertMessagesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAlertMessagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAlertMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlertMessagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAlertMessagesResponseBody) SetData(v *ListAlertMessagesResponseBodyData) *ListAlertMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListAlertMessagesResponseBody) SetErrorCode(v string) *ListAlertMessagesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAlertMessagesResponseBody) SetErrorMessage(v string) *ListAlertMessagesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAlertMessagesResponseBody) SetHttpStatusCode(v int32) *ListAlertMessagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAlertMessagesResponseBody) SetRequestId(v string) *ListAlertMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlertMessagesResponseBody) SetSuccess(v bool) *ListAlertMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlertMessagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAlertMessagesResponseBodyData struct {
	// The alerts.
	AlertMessages []*ListAlertMessagesResponseBodyDataAlertMessages `json:"AlertMessages,omitempty" xml:"AlertMessages,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of alerts returned.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAlertMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBodyData) GetAlertMessages() []*ListAlertMessagesResponseBodyDataAlertMessages {
	return s.AlertMessages
}

func (s *ListAlertMessagesResponseBodyData) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAlertMessagesResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAlertMessagesResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListAlertMessagesResponseBodyData) SetAlertMessages(v []*ListAlertMessagesResponseBodyDataAlertMessages) *ListAlertMessagesResponseBodyData {
	s.AlertMessages = v
	return s
}

func (s *ListAlertMessagesResponseBodyData) SetPageNumber(v string) *ListAlertMessagesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAlertMessagesResponseBodyData) SetPageSize(v string) *ListAlertMessagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAlertMessagesResponseBodyData) SetTotalCount(v string) *ListAlertMessagesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAlertMessagesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListAlertMessagesResponseBodyDataAlertMessages struct {
	// The alert ID.
	//
	// example:
	//
	// 1234
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The sending status of the notification. Valid values: READY_TO_SEND, SEND_FAIL, SEND_SUCCESS, and SEND_OVERLIMIT. The value READY_TO_SEND indicates that the notification is waiting to be sent. The value SEND_FAIL indicates that the notification fails to be sent. The value SEND_SUCCESS indicates that the notification is sent. The value SEND_OVERLIMIT indicates that the number of notifications that are sent exceeds the upper limit.
	//
	// example:
	//
	// READY_TO_SEND
	AlertMessageStatus *string `json:"AlertMessageStatus,omitempty" xml:"AlertMessageStatus,omitempty"`
	// The notification method. Valid values: MAIL, SMS, and PHONE. Only DataWorks Professional Edition and more advanced editions support the PHONE notification method.
	//
	// example:
	//
	// SMS
	AlertMethod *string `json:"AlertMethod,omitempty" xml:"AlertMethod,omitempty"`
	// The timestamp when the alert was reported.
	//
	// example:
	//
	// 1553531401000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// The ID of the Alibaba Cloud used by the alert recipient.
	//
	// example:
	//
	// 9527952795****
	AlertUser *string `json:"AlertUser,omitempty" xml:"AlertUser,omitempty"`
	// The content of the alert.
	//
	// example:
	//
	// Node error
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The instances that triggered the custom alert rule. This parameter is returned if the value of the Source parameter is REMIND_ALERT. This parameter is left empty if the value of the Source parameter is not REMIND_ALERT.
	Instances []*ListAlertMessagesResponseBodyDataAlertMessagesInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The nodes returned for different alert sources.
	//
	// 	- The nodes that form a loop are returned if the value of the Source parameter is NODE_CYCLE_ALERT.
	//
	// 	- The nodes that are isolated are returned if the value of the Source parameter is NODE_LONELY_ALERT.
	Nodes []*ListAlertMessagesResponseBodyDataAlertMessagesNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The ID of the custom alert rule that was triggered. This parameter is returned if the value of the Source parameter is REMIND_ALERT.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	// The name of the custom alert rule that was triggered. This parameter is returned if the value of the Source parameter is REMIND_ALERT.
	//
	// example:
	//
	// Custom monitoring rule name
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The basic information about the baseline instance that triggered an alert. This parameter is returned if the value of the Source parameter is SLA_ALERT. This parameter is left empty if the value of the Source parameter is not SLA_ALERT.
	SlaAlert *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert `json:"SlaAlert,omitempty" xml:"SlaAlert,omitempty" type:"Struct"`
	// The type of the alert. Valid values: REMIND_ALERT, TOPIC_ALERT, SLA_ALERT, NODE_CYCLE_ALERT, and NODE_LONELY_ALERT. The value REMIND_ALERT indicates that the alert is a custom alert. The value TOPIC_ALERT indicates that the alert is an event alert. The value SLA_ALERT indicates that the alert is a baseline alert. The value NODE_CYCLE_ALERT indicates that the alert is reported for a node dependency loop. The value NODE_LONELY_ALERT indicates that the alert is reported for isolated nodes.
	//
	// example:
	//
	// REMIND_ALERT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The events that triggered alerts. This parameter is returned if the value of the Source parameter is TOPIC_ALERT. This parameter is left empty if the value of the Source parameter is not TOPIC_ALERT.
	Topics []*ListAlertMessagesResponseBodyDataAlertMessagesTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s ListAlertMessagesResponseBodyDataAlertMessages) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBodyDataAlertMessages) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetAlertId() *int64 {
	return s.AlertId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetAlertMessageStatus() *string {
	return s.AlertMessageStatus
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetAlertMethod() *string {
	return s.AlertMethod
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetAlertUser() *string {
	return s.AlertUser
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetContent() *string {
	return s.Content
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetInstances() []*ListAlertMessagesResponseBodyDataAlertMessagesInstances {
	return s.Instances
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetNodes() []*ListAlertMessagesResponseBodyDataAlertMessagesNodes {
	return s.Nodes
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetRemindId() *int64 {
	return s.RemindId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetRemindName() *string {
	return s.RemindName
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetSlaAlert() *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	return s.SlaAlert
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetSource() *string {
	return s.Source
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) GetTopics() []*ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	return s.Topics
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetAlertId(v int64) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.AlertId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetAlertMessageStatus(v string) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.AlertMessageStatus = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetAlertMethod(v string) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.AlertMethod = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetAlertTime(v int64) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.AlertTime = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetAlertUser(v string) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.AlertUser = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetContent(v string) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.Content = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetInstances(v []*ListAlertMessagesResponseBodyDataAlertMessagesInstances) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.Instances = v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetNodes(v []*ListAlertMessagesResponseBodyDataAlertMessagesNodes) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.Nodes = v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetRemindId(v int64) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.RemindId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetRemindName(v string) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.RemindName = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetSlaAlert(v *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.SlaAlert = v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetSource(v string) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.Source = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) SetTopics(v []*ListAlertMessagesResponseBodyDataAlertMessagesTopics) *ListAlertMessagesResponseBodyDataAlertMessages {
	s.Topics = v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessages) Validate() error {
	return dara.Validate(s)
}

type ListAlertMessagesResponseBodyDataAlertMessagesInstances struct {
	// The instance ID.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the instance. Valid values: NOT_RUN, WAIT_TIME, WAIT_RESOURCE, RUNNING, CHECKING, CHECKING_CONDITION, FAILURE, and SUCCESS. The value NOT_RUN indicates that the instance is not run. The value WAIT_TIME indicates that the instance is waiting to be run. The value WAIT_RESOURCE indicates that the instance is waiting for resources. The value RUNNING indicates that the instance is running. The value CHECKING indicates that data quality is being checked for the node for which the instance is generated. The value CHECKING_CONDITION indicates that branch conditions are being checked for the node for which the instance is generated. The value FAILURE indicates that the instance fails to run. The value SUCCESS indicates that the instance is successfully run.
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesInstances) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesInstances) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) GetStatus() *string {
	return s.Status
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) SetInstanceId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesInstances {
	s.InstanceId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) SetNodeId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesInstances {
	s.NodeId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) SetNodeName(v string) *ListAlertMessagesResponseBodyDataAlertMessagesInstances {
	s.NodeName = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) SetProjectId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesInstances {
	s.ProjectId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) SetStatus(v string) *ListAlertMessagesResponseBodyDataAlertMessagesInstances {
	s.Status = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesInstances) Validate() error {
	return dara.Validate(s)
}

type ListAlertMessagesResponseBodyDataAlertMessagesNodes struct {
	// The node ID.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 95279527952****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesNodes) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesNodes) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) GetOwner() *string {
	return s.Owner
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) SetNodeId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesNodes {
	s.NodeId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) SetNodeName(v string) *ListAlertMessagesResponseBodyDataAlertMessagesNodes {
	s.NodeName = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) SetOwner(v string) *ListAlertMessagesResponseBodyDataAlertMessagesNodes {
	s.Owner = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) SetProjectId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesNodes {
	s.ProjectId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesNodes) Validate() error {
	return dara.Validate(s)
}

type ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert struct {
	// The baseline ID.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Baseline name
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The ID of the Alibaba Cloud account used by the baseline owner. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 952795279****
	BaselineOwner *string `json:"BaselineOwner,omitempty" xml:"BaselineOwner,omitempty"`
	// The data timestamp of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The ID of the cycle of the baseline instance. Valid values of the ID of an hour-level cycle: [1,24]. The ID of a day-level cycle is 1.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The ID of the workspace to which the baseline belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGEROUS, and OVER. The value ERROR indicates that no nodes are associated with the baseline, or all nodes associated with the baseline are suspended. The value SAFE indicates that nodes are run before the alert duration begins. The value DANGEROUS indicates that nodes are still running after the alert duration ends but the committed completion time does not arrive. The value OVER indicates that nodes are still running after the committed completion time.
	//
	// example:
	//
	// SAFE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetBaselineName() *string {
	return s.BaselineName
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetBaselineOwner() *string {
	return s.BaselineOwner
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) GetStatus() *string {
	return s.Status
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetBaselineId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.BaselineId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetBaselineName(v string) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.BaselineName = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetBaselineOwner(v string) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.BaselineOwner = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetBizdate(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.Bizdate = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetInGroupId(v int32) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.InGroupId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetProjectId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.ProjectId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) SetStatus(v string) *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert {
	s.Status = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesSlaAlert) Validate() error {
	return dara.Validate(s)
}

type ListAlertMessagesResponseBodyDataAlertMessagesTopics struct {
	// The ID of the instance that triggered the event.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node that triggered the event.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// 9527 error
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The ID of the Alibaba Cloud account used by the event owner.
	//
	// example:
	//
	// 9527952795****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
	// The status of the event. Valid values: IGNORE, NEW, FIXING, and RECOVER. The value IGNORE indicates that the event is ignored. The value NEW indicates that the event is a new event. The value FIXING indicates that the event is being handled. The value RECOVER indicates that the event is handled.
	//
	// example:
	//
	// FIXING
	TopicStatus *string `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty"`
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesTopics) String() string {
	return dara.Prettify(s)
}

func (s ListAlertMessagesResponseBodyDataAlertMessagesTopics) GoString() string {
	return s.String()
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) GetTopicOwner() *string {
	return s.TopicOwner
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) GetTopicStatus() *string {
	return s.TopicStatus
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) SetInstanceId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	s.InstanceId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) SetNodeId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	s.NodeId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) SetTopicId(v int64) *ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	s.TopicId = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) SetTopicName(v string) *ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	s.TopicName = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) SetTopicOwner(v string) *ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	s.TopicOwner = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) SetTopicStatus(v string) *ListAlertMessagesResponseBodyDataAlertMessagesTopics {
	s.TopicStatus = &v
	return s
}

func (s *ListAlertMessagesResponseBodyDataAlertMessagesTopics) Validate() error {
	return dara.Validate(s)
}
