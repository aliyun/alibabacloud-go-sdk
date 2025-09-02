// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAlertMessageResponseBodyData) *GetAlertMessageResponseBody
	GetData() *GetAlertMessageResponseBodyData
	SetErrorCode(v string) *GetAlertMessageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAlertMessageResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetAlertMessageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetAlertMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAlertMessageResponseBody
	GetSuccess() *bool
}

type GetAlertMessageResponseBody struct {
	// The information about returned alerts.
	Data *GetAlertMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetAlertMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponseBody) GetData() *GetAlertMessageResponseBodyData {
	return s.Data
}

func (s *GetAlertMessageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAlertMessageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAlertMessageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAlertMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAlertMessageResponseBody) SetData(v *GetAlertMessageResponseBodyData) *GetAlertMessageResponseBody {
	s.Data = v
	return s
}

func (s *GetAlertMessageResponseBody) SetErrorCode(v string) *GetAlertMessageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAlertMessageResponseBody) SetErrorMessage(v string) *GetAlertMessageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAlertMessageResponseBody) SetHttpStatusCode(v int32) *GetAlertMessageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAlertMessageResponseBody) SetRequestId(v string) *GetAlertMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertMessageResponseBody) SetSuccess(v bool) *GetAlertMessageResponseBody {
	s.Success = &v
	return s
}

func (s *GetAlertMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAlertMessageResponseBodyData struct {
	// The alert ID.
	//
	// example:
	//
	// 123
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The sending status of the notification. Valid values:
	//
	// 	- READY_TO_SEND: The notification is waiting to be sent.
	//
	// 	- SEND_FAIL: The notification fails to be sent.
	//
	// 	- SEND_SUCCESS: The notification is sent.
	//
	// 	- SEND_OVERLIMIT: The number of notifications that are sent exceeds the upper limit.
	//
	// example:
	//
	// READY_TO_SEND
	AlertMessageStatus *string `json:"AlertMessageStatus,omitempty" xml:"AlertMessageStatus,omitempty"`
	// The notification method. Valid values:
	//
	// 	- MAIL.
	//
	// 	- SMS.
	//
	// 	- PHONE. Only DataWorks Professional Edition and more advanced editions support the PHONE notification method.
	//
	// example:
	//
	// SMS
	AlertMethod *string `json:"AlertMethod,omitempty" xml:"AlertMethod,omitempty"`
	// The time when the alert was reported.
	//
	// example:
	//
	// 1553524393000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the alert recipient.
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
	Instances []*GetAlertMessageResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The nodes returned for different alert sources. The nodes that form a loop are returned if the value of the Source parameter is NODE_CYCLE_ALERT. The nodes that are isolated are returned if the value of the Source parameter is NODE_LONELY_ALERT.
	Nodes []*GetAlertMessageResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
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
	SlaAlert *GetAlertMessageResponseBodyDataSlaAlert `json:"SlaAlert,omitempty" xml:"SlaAlert,omitempty" type:"Struct"`
	// The type of the alert. Valid values:
	//
	// 	- REMIND_ALERT: The alert is a custom alert.
	//
	// 	- TOPIC_ALERT: The alert is an event alert.
	//
	// 	- SLA_ALERT: The alert is a baseline alert.
	//
	// 	- NODE_CYCLE_ALERT: The alert is reported for a node dependency loop.
	//
	// 	- NODE_LONELY_ALERT: The alert is reported for isolated nodes.
	//
	// example:
	//
	// REMIND_ALERT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The events that triggered alerts. This parameter is returned if the value of the Source parameter is TOPIC_ALERT. This parameter is left empty if the value of the Source parameter is not TOPIC_ALERT.
	Topics []*GetAlertMessageResponseBodyDataTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s GetAlertMessageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponseBodyData) GetAlertId() *int64 {
	return s.AlertId
}

func (s *GetAlertMessageResponseBodyData) GetAlertMessageStatus() *string {
	return s.AlertMessageStatus
}

func (s *GetAlertMessageResponseBodyData) GetAlertMethod() *string {
	return s.AlertMethod
}

func (s *GetAlertMessageResponseBodyData) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *GetAlertMessageResponseBodyData) GetAlertUser() *string {
	return s.AlertUser
}

func (s *GetAlertMessageResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetAlertMessageResponseBodyData) GetInstances() []*GetAlertMessageResponseBodyDataInstances {
	return s.Instances
}

func (s *GetAlertMessageResponseBodyData) GetNodes() []*GetAlertMessageResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetAlertMessageResponseBodyData) GetRemindId() *int64 {
	return s.RemindId
}

func (s *GetAlertMessageResponseBodyData) GetRemindName() *string {
	return s.RemindName
}

func (s *GetAlertMessageResponseBodyData) GetSlaAlert() *GetAlertMessageResponseBodyDataSlaAlert {
	return s.SlaAlert
}

func (s *GetAlertMessageResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetAlertMessageResponseBodyData) GetTopics() []*GetAlertMessageResponseBodyDataTopics {
	return s.Topics
}

func (s *GetAlertMessageResponseBodyData) SetAlertId(v int64) *GetAlertMessageResponseBodyData {
	s.AlertId = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetAlertMessageStatus(v string) *GetAlertMessageResponseBodyData {
	s.AlertMessageStatus = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetAlertMethod(v string) *GetAlertMessageResponseBodyData {
	s.AlertMethod = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetAlertTime(v int64) *GetAlertMessageResponseBodyData {
	s.AlertTime = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetAlertUser(v string) *GetAlertMessageResponseBodyData {
	s.AlertUser = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetContent(v string) *GetAlertMessageResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetInstances(v []*GetAlertMessageResponseBodyDataInstances) *GetAlertMessageResponseBodyData {
	s.Instances = v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetNodes(v []*GetAlertMessageResponseBodyDataNodes) *GetAlertMessageResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetRemindId(v int64) *GetAlertMessageResponseBodyData {
	s.RemindId = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetRemindName(v string) *GetAlertMessageResponseBodyData {
	s.RemindName = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetSlaAlert(v *GetAlertMessageResponseBodyDataSlaAlert) *GetAlertMessageResponseBodyData {
	s.SlaAlert = v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetSource(v string) *GetAlertMessageResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetAlertMessageResponseBodyData) SetTopics(v []*GetAlertMessageResponseBodyDataTopics) *GetAlertMessageResponseBodyData {
	s.Topics = v
	return s
}

func (s *GetAlertMessageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetAlertMessageResponseBodyDataInstances struct {
	// The instance ID.
	//
	// example:
	//
	// 12312312
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
	// The status of the instance. Valid values:
	//
	// 	- NOT_RUN
	//
	// 	- WAIT_TIME
	//
	// 	- WAIT_RESOURCE
	//
	// 	- RUNNING
	//
	// 	- CHECKING
	//
	// 	- CHECKING_CONDITION
	//
	// 	- FAILURE
	//
	// 	- SUCCESS
	//
	// example:
	//
	// NOT_RUN
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAlertMessageResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponseBodyDataInstances) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetAlertMessageResponseBodyDataInstances) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetAlertMessageResponseBodyDataInstances) GetNodeName() *string {
	return s.NodeName
}

func (s *GetAlertMessageResponseBodyDataInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAlertMessageResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *GetAlertMessageResponseBodyDataInstances) SetInstanceId(v int64) *GetAlertMessageResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataInstances) SetNodeId(v int64) *GetAlertMessageResponseBodyDataInstances {
	s.NodeId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataInstances) SetNodeName(v string) *GetAlertMessageResponseBodyDataInstances {
	s.NodeName = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataInstances) SetProjectId(v int64) *GetAlertMessageResponseBodyDataInstances {
	s.ProjectId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataInstances) SetStatus(v string) *GetAlertMessageResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}

type GetAlertMessageResponseBodyDataNodes struct {
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

func (s GetAlertMessageResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponseBodyDataNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetAlertMessageResponseBodyDataNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetAlertMessageResponseBodyDataNodes) GetOwner() *string {
	return s.Owner
}

func (s *GetAlertMessageResponseBodyDataNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAlertMessageResponseBodyDataNodes) SetNodeId(v int64) *GetAlertMessageResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataNodes) SetNodeName(v string) *GetAlertMessageResponseBodyDataNodes {
	s.NodeName = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataNodes) SetOwner(v string) *GetAlertMessageResponseBodyDataNodes {
	s.Owner = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataNodes) SetProjectId(v int64) *GetAlertMessageResponseBodyDataNodes {
	s.ProjectId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}

type GetAlertMessageResponseBodyDataSlaAlert struct {
	// The baseline ID.
	//
	// example:
	//
	// 15142123
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
	// The status of the baseline. Valid values:
	//
	// 	- ERROR
	//
	// 	- SAFE
	//
	// 	- DANGEROUS
	//
	// 	- OVER
	//
	// example:
	//
	// SAFE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAlertMessageResponseBodyDataSlaAlert) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponseBodyDataSlaAlert) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetBaselineOwner() *string {
	return s.BaselineOwner
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) GetStatus() *string {
	return s.Status
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetBaselineId(v int64) *GetAlertMessageResponseBodyDataSlaAlert {
	s.BaselineId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetBaselineName(v string) *GetAlertMessageResponseBodyDataSlaAlert {
	s.BaselineName = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetBaselineOwner(v string) *GetAlertMessageResponseBodyDataSlaAlert {
	s.BaselineOwner = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetBizdate(v int64) *GetAlertMessageResponseBodyDataSlaAlert {
	s.Bizdate = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetInGroupId(v int32) *GetAlertMessageResponseBodyDataSlaAlert {
	s.InGroupId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetProjectId(v int64) *GetAlertMessageResponseBodyDataSlaAlert {
	s.ProjectId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) SetStatus(v string) *GetAlertMessageResponseBodyDataSlaAlert {
	s.Status = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataSlaAlert) Validate() error {
	return dara.Validate(s)
}

type GetAlertMessageResponseBodyDataTopics struct {
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
	// 412431
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// error
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The ID of the Alibaba Cloud account used by the event owner.
	//
	// example:
	//
	// 9527952795****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- IGNORE
	//
	// 	- NEW
	//
	// 	- FIXING
	//
	// 	- RECOVER
	//
	// example:
	//
	// FIXING
	TopicStatus *string `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty"`
}

func (s GetAlertMessageResponseBodyDataTopics) String() string {
	return dara.Prettify(s)
}

func (s GetAlertMessageResponseBodyDataTopics) GoString() string {
	return s.String()
}

func (s *GetAlertMessageResponseBodyDataTopics) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetAlertMessageResponseBodyDataTopics) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetAlertMessageResponseBodyDataTopics) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetAlertMessageResponseBodyDataTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *GetAlertMessageResponseBodyDataTopics) GetTopicOwner() *string {
	return s.TopicOwner
}

func (s *GetAlertMessageResponseBodyDataTopics) GetTopicStatus() *string {
	return s.TopicStatus
}

func (s *GetAlertMessageResponseBodyDataTopics) SetInstanceId(v int64) *GetAlertMessageResponseBodyDataTopics {
	s.InstanceId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataTopics) SetNodeId(v int64) *GetAlertMessageResponseBodyDataTopics {
	s.NodeId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataTopics) SetTopicId(v int64) *GetAlertMessageResponseBodyDataTopics {
	s.TopicId = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataTopics) SetTopicName(v string) *GetAlertMessageResponseBodyDataTopics {
	s.TopicName = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataTopics) SetTopicOwner(v string) *GetAlertMessageResponseBodyDataTopics {
	s.TopicOwner = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataTopics) SetTopicStatus(v string) *GetAlertMessageResponseBodyDataTopics {
	s.TopicStatus = &v
	return s
}

func (s *GetAlertMessageResponseBodyDataTopics) Validate() error {
	return dara.Validate(s)
}
