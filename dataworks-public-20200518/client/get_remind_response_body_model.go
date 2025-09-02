// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRemindResponseBodyData) *GetRemindResponseBody
	GetData() *GetRemindResponseBodyData
	SetErrorCode(v string) *GetRemindResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetRemindResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetRemindResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetRemindResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRemindResponseBody
	GetSuccess() *bool
}

type GetRemindResponseBody struct {
	// The details of the custom alert rule.
	Data *GetRemindResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
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
	// 0000-ABCD-EFGH-IJKLMNOPQ
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRemindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBody) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBody) GetData() *GetRemindResponseBodyData {
	return s.Data
}

func (s *GetRemindResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRemindResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRemindResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRemindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRemindResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRemindResponseBody) SetData(v *GetRemindResponseBodyData) *GetRemindResponseBody {
	s.Data = v
	return s
}

func (s *GetRemindResponseBody) SetErrorCode(v string) *GetRemindResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRemindResponseBody) SetErrorMessage(v string) *GetRemindResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetRemindResponseBody) SetHttpStatusCode(v int32) *GetRemindResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRemindResponseBody) SetRequestId(v string) *GetRemindResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRemindResponseBody) SetSuccess(v bool) *GetRemindResponseBody {
	s.Success = &v
	return s
}

func (s *GetRemindResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyData struct {
	// The minimum interval at which alerts are reported. Unit: seconds.
	//
	// example:
	//
	// 1800
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The notification method.
	AlertMethods []*string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty" type:"Repeated"`
	// The description of the alert recipient. For more information about alert recipients, see the Receivers parameter.
	AlertTargets []*string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty" type:"Repeated"`
	// The type of the alert recipient. For more information about alert recipient types, see the Receivers parameter. Valid values:
	//
	// 	- OWNER: the task owner
	//
	// 	- OTHER: specified personnel
	//
	// 	- SHIFT_SCHEDULE: personnel in a shift schedule
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
	// The IDs of the nodes that are added to a whitelist.
	AllowNodes []*int64 `json:"AllowNodes,omitempty" xml:"AllowNodes,omitempty" type:"Repeated"`
	// The baselines to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is BASELINE.
	Baselines []*GetRemindResponseBodyDataBaselines `json:"Baselines,omitempty" xml:"Baselines,omitempty" type:"Repeated"`
	// The workflows to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is BIZPROCESS.
	BizProcesses []*GetRemindResponseBodyDataBizProcesses `json:"BizProcesses,omitempty" xml:"BizProcesses,omitempty" type:"Repeated"`
	// 	- If the value of the RemindType parameter is FINISHED, this parameter is left empty.
	//
	// 	- If the value of the RemindType parameter is UNFINISHED, the trigger conditions are returned as key-value pairs. Example: {"hour":23,"minu":59}. Valid values of hour: [0,47]. Valid values of minu: [0,59].
	//
	// 	- If the value of the RemindType parameter is ERROR, this parameter is left empty.
	//
	// 	- If the value of the RemindType parameter is CYCLE_UNFINISHED, the trigger conditions are returned as key-value pairs. Example: {"1":"05:50","2":"06:50","3":"07:50","4":"08:50","5":"09:50","6":"10:50","7":"11:50","8":"12:50","9":"13:50","10":"14:50","11":"15:50","12":"16:50","13":"17:50","14":"18:50","15":"19:50","16":"20:50","17":"21:50","18":"22:50","19":"23:50","20":"24:50","21":"25:50"}. The key indicates the ID of the cycle. Valid values: [1,288]. The value indicates the timeout period of the node that is running in the cycle. Specify the value in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// 	- If the value of the RemindType parameter is TIMEOUT, the timeout period is returned. Unit: seconds. Example: 1800. This value indicates that an alert is reported if the node has run for more than 30 minutes.
	//
	// example:
	//
	// {"hour":23,"minu":59}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end time of the quiet hours. The value is in the hh:mm format. Valid values of hh: [0,23]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 08:00
	DndEnd *string `json:"DndEnd,omitempty" xml:"DndEnd,omitempty"`
	// The start time of the quiet hours. The value is in the hh:mm format. Valid values of hh: [0,23]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 00:00
	DndStart *string `json:"DndStart,omitempty" xml:"DndStart,omitempty"`
	// The ID of the Alibaba Cloud account used by the creator of the custom alert rule.
	//
	// example:
	//
	// 9527951795****
	Founder *string `json:"Founder,omitempty" xml:"Founder,omitempty"`
	// The maximum number of alerts.
	//
	// example:
	//
	// 3
	MaxAlertTimes *int32 `json:"MaxAlertTimes,omitempty" xml:"MaxAlertTimes,omitempty"`
	// The nodes to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is NODE.
	Nodes []*GetRemindResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The workspaces to which the custom alert rule is applied. This parameter is returned if the value of the RemindUnit parameter is PROJECT.
	Projects []*GetRemindResponseBodyDataProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The information about the alert recipients.
	Receivers []*GetRemindResponseBodyDataReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The custom alert rule ID.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// Rule name
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The conditions that trigger an alert. Valid values: FINISHED, UNFINISHED, ERROR, CYCLE_UNFINISHED, and TIMEOUT.
	//
	// example:
	//
	// FINISHED
	RemindType *string `json:"RemindType,omitempty" xml:"RemindType,omitempty"`
	// The type of the object to which the custom alert rule is applied. Valid values: NODE, BASELINE, PROJECT, and BIZPROCESS. The value NODE indicates a node. The value BASELINE indicates a baseline. The value PROJECT indicates a workspace. The value BIZPROCESS indicates a workflow.
	//
	// example:
	//
	// NODE
	RemindUnit *string `json:"RemindUnit,omitempty" xml:"RemindUnit,omitempty"`
	// The webhook URLs of the DingTalk chatbots.
	Robots []*GetRemindResponseBodyDataRobots `json:"Robots,omitempty" xml:"Robots,omitempty" type:"Repeated"`
	// Indicates whether the custom alert rule is enabled. Valid values: true and false.
	//
	// example:
	//
	// true
	Useflag *bool `json:"Useflag,omitempty" xml:"Useflag,omitempty"`
	// WebHook URL
	Webhooks []*string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty" type:"Repeated"`
}

func (s GetRemindResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyData) GetAlertInterval() *int32 {
	return s.AlertInterval
}

func (s *GetRemindResponseBodyData) GetAlertMethods() []*string {
	return s.AlertMethods
}

func (s *GetRemindResponseBodyData) GetAlertTargets() []*string {
	return s.AlertTargets
}

func (s *GetRemindResponseBodyData) GetAlertUnit() *string {
	return s.AlertUnit
}

func (s *GetRemindResponseBodyData) GetAllowNodes() []*int64 {
	return s.AllowNodes
}

func (s *GetRemindResponseBodyData) GetBaselines() []*GetRemindResponseBodyDataBaselines {
	return s.Baselines
}

func (s *GetRemindResponseBodyData) GetBizProcesses() []*GetRemindResponseBodyDataBizProcesses {
	return s.BizProcesses
}

func (s *GetRemindResponseBodyData) GetDetail() *string {
	return s.Detail
}

func (s *GetRemindResponseBodyData) GetDndEnd() *string {
	return s.DndEnd
}

func (s *GetRemindResponseBodyData) GetDndStart() *string {
	return s.DndStart
}

func (s *GetRemindResponseBodyData) GetFounder() *string {
	return s.Founder
}

func (s *GetRemindResponseBodyData) GetMaxAlertTimes() *int32 {
	return s.MaxAlertTimes
}

func (s *GetRemindResponseBodyData) GetNodes() []*GetRemindResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetRemindResponseBodyData) GetProjects() []*GetRemindResponseBodyDataProjects {
	return s.Projects
}

func (s *GetRemindResponseBodyData) GetReceivers() []*GetRemindResponseBodyDataReceivers {
	return s.Receivers
}

func (s *GetRemindResponseBodyData) GetRemindId() *int64 {
	return s.RemindId
}

func (s *GetRemindResponseBodyData) GetRemindName() *string {
	return s.RemindName
}

func (s *GetRemindResponseBodyData) GetRemindType() *string {
	return s.RemindType
}

func (s *GetRemindResponseBodyData) GetRemindUnit() *string {
	return s.RemindUnit
}

func (s *GetRemindResponseBodyData) GetRobots() []*GetRemindResponseBodyDataRobots {
	return s.Robots
}

func (s *GetRemindResponseBodyData) GetUseflag() *bool {
	return s.Useflag
}

func (s *GetRemindResponseBodyData) GetWebhooks() []*string {
	return s.Webhooks
}

func (s *GetRemindResponseBodyData) SetAlertInterval(v int32) *GetRemindResponseBodyData {
	s.AlertInterval = &v
	return s
}

func (s *GetRemindResponseBodyData) SetAlertMethods(v []*string) *GetRemindResponseBodyData {
	s.AlertMethods = v
	return s
}

func (s *GetRemindResponseBodyData) SetAlertTargets(v []*string) *GetRemindResponseBodyData {
	s.AlertTargets = v
	return s
}

func (s *GetRemindResponseBodyData) SetAlertUnit(v string) *GetRemindResponseBodyData {
	s.AlertUnit = &v
	return s
}

func (s *GetRemindResponseBodyData) SetAllowNodes(v []*int64) *GetRemindResponseBodyData {
	s.AllowNodes = v
	return s
}

func (s *GetRemindResponseBodyData) SetBaselines(v []*GetRemindResponseBodyDataBaselines) *GetRemindResponseBodyData {
	s.Baselines = v
	return s
}

func (s *GetRemindResponseBodyData) SetBizProcesses(v []*GetRemindResponseBodyDataBizProcesses) *GetRemindResponseBodyData {
	s.BizProcesses = v
	return s
}

func (s *GetRemindResponseBodyData) SetDetail(v string) *GetRemindResponseBodyData {
	s.Detail = &v
	return s
}

func (s *GetRemindResponseBodyData) SetDndEnd(v string) *GetRemindResponseBodyData {
	s.DndEnd = &v
	return s
}

func (s *GetRemindResponseBodyData) SetDndStart(v string) *GetRemindResponseBodyData {
	s.DndStart = &v
	return s
}

func (s *GetRemindResponseBodyData) SetFounder(v string) *GetRemindResponseBodyData {
	s.Founder = &v
	return s
}

func (s *GetRemindResponseBodyData) SetMaxAlertTimes(v int32) *GetRemindResponseBodyData {
	s.MaxAlertTimes = &v
	return s
}

func (s *GetRemindResponseBodyData) SetNodes(v []*GetRemindResponseBodyDataNodes) *GetRemindResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetRemindResponseBodyData) SetProjects(v []*GetRemindResponseBodyDataProjects) *GetRemindResponseBodyData {
	s.Projects = v
	return s
}

func (s *GetRemindResponseBodyData) SetReceivers(v []*GetRemindResponseBodyDataReceivers) *GetRemindResponseBodyData {
	s.Receivers = v
	return s
}

func (s *GetRemindResponseBodyData) SetRemindId(v int64) *GetRemindResponseBodyData {
	s.RemindId = &v
	return s
}

func (s *GetRemindResponseBodyData) SetRemindName(v string) *GetRemindResponseBodyData {
	s.RemindName = &v
	return s
}

func (s *GetRemindResponseBodyData) SetRemindType(v string) *GetRemindResponseBodyData {
	s.RemindType = &v
	return s
}

func (s *GetRemindResponseBodyData) SetRemindUnit(v string) *GetRemindResponseBodyData {
	s.RemindUnit = &v
	return s
}

func (s *GetRemindResponseBodyData) SetRobots(v []*GetRemindResponseBodyDataRobots) *GetRemindResponseBodyData {
	s.Robots = v
	return s
}

func (s *GetRemindResponseBodyData) SetUseflag(v bool) *GetRemindResponseBodyData {
	s.Useflag = &v
	return s
}

func (s *GetRemindResponseBodyData) SetWebhooks(v []*string) *GetRemindResponseBodyData {
	s.Webhooks = v
	return s
}

func (s *GetRemindResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyDataBaselines struct {
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
}

func (s GetRemindResponseBodyDataBaselines) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyDataBaselines) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyDataBaselines) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetRemindResponseBodyDataBaselines) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetRemindResponseBodyDataBaselines) SetBaselineId(v int64) *GetRemindResponseBodyDataBaselines {
	s.BaselineId = &v
	return s
}

func (s *GetRemindResponseBodyDataBaselines) SetBaselineName(v string) *GetRemindResponseBodyDataBaselines {
	s.BaselineName = &v
	return s
}

func (s *GetRemindResponseBodyDataBaselines) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyDataBizProcesses struct {
	// The ID of the workflow.
	//
	// example:
	//
	// 9527
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// Business process name
	BizProcessName *string `json:"BizProcessName,omitempty" xml:"BizProcessName,omitempty"`
}

func (s GetRemindResponseBodyDataBizProcesses) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyDataBizProcesses) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyDataBizProcesses) GetBizId() *int64 {
	return s.BizId
}

func (s *GetRemindResponseBodyDataBizProcesses) GetBizProcessName() *string {
	return s.BizProcessName
}

func (s *GetRemindResponseBodyDataBizProcesses) SetBizId(v int64) *GetRemindResponseBodyDataBizProcesses {
	s.BizId = &v
	return s
}

func (s *GetRemindResponseBodyDataBizProcesses) SetBizProcessName(v string) *GetRemindResponseBodyDataBizProcesses {
	s.BizProcessName = &v
	return s
}

func (s *GetRemindResponseBodyDataBizProcesses) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyDataNodes struct {
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
	// 9527951795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetRemindResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyDataNodes) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetRemindResponseBodyDataNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetRemindResponseBodyDataNodes) GetOwner() *string {
	return s.Owner
}

func (s *GetRemindResponseBodyDataNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetRemindResponseBodyDataNodes) SetNodeId(v int64) *GetRemindResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *GetRemindResponseBodyDataNodes) SetNodeName(v string) *GetRemindResponseBodyDataNodes {
	s.NodeName = &v
	return s
}

func (s *GetRemindResponseBodyDataNodes) SetOwner(v string) *GetRemindResponseBodyDataNodes {
	s.Owner = &v
	return s
}

func (s *GetRemindResponseBodyDataNodes) SetProjectId(v int64) *GetRemindResponseBodyDataNodes {
	s.ProjectId = &v
	return s
}

func (s *GetRemindResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyDataProjects struct {
	// The workspace ID.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetRemindResponseBodyDataProjects) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyDataProjects) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetRemindResponseBodyDataProjects) SetProjectId(v int64) *GetRemindResponseBodyDataProjects {
	s.ProjectId = &v
	return s
}

func (s *GetRemindResponseBodyDataProjects) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyDataReceivers struct {
	// The alert recipient.
	AlertTargets []*string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty" type:"Repeated"`
	// The type of the alert recipient. For more information about alert recipients, see the Receivers parameter. Valid values:
	//
	// 	- OWNER: indicate the task owner.
	//
	// 	- OTHER: indicates specified personnel.
	//
	// 	- SHIFT_SCHEDULE: indicates personnel in a shift schedule.
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
}

func (s GetRemindResponseBodyDataReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyDataReceivers) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyDataReceivers) GetAlertTargets() []*string {
	return s.AlertTargets
}

func (s *GetRemindResponseBodyDataReceivers) GetAlertUnit() *string {
	return s.AlertUnit
}

func (s *GetRemindResponseBodyDataReceivers) SetAlertTargets(v []*string) *GetRemindResponseBodyDataReceivers {
	s.AlertTargets = v
	return s
}

func (s *GetRemindResponseBodyDataReceivers) SetAlertUnit(v string) *GetRemindResponseBodyDataReceivers {
	s.AlertUnit = &v
	return s
}

func (s *GetRemindResponseBodyDataReceivers) Validate() error {
	return dara.Validate(s)
}

type GetRemindResponseBodyDataRobots struct {
	// Indicates whether all group members are notified when the alert notification is sent to a DingTalk group. Valid values: true and false.
	//
	// example:
	//
	// true
	AtAll *bool `json:"AtAll,omitempty" xml:"AtAll,omitempty"`
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=******************************
	WebUrl *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s GetRemindResponseBodyDataRobots) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponseBodyDataRobots) GoString() string {
	return s.String()
}

func (s *GetRemindResponseBodyDataRobots) GetAtAll() *bool {
	return s.AtAll
}

func (s *GetRemindResponseBodyDataRobots) GetWebUrl() *string {
	return s.WebUrl
}

func (s *GetRemindResponseBodyDataRobots) SetAtAll(v bool) *GetRemindResponseBodyDataRobots {
	s.AtAll = &v
	return s
}

func (s *GetRemindResponseBodyDataRobots) SetWebUrl(v string) *GetRemindResponseBodyDataRobots {
	s.WebUrl = &v
	return s
}

func (s *GetRemindResponseBodyDataRobots) Validate() error {
	return dara.Validate(s)
}
