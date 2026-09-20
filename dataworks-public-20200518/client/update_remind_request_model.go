// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRemindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertInterval(v int32) *UpdateRemindRequest
	GetAlertInterval() *int32
	SetAlertMethods(v string) *UpdateRemindRequest
	GetAlertMethods() *string
	SetAlertTargets(v string) *UpdateRemindRequest
	GetAlertTargets() *string
	SetAlertUnit(v string) *UpdateRemindRequest
	GetAlertUnit() *string
	SetBaselineIds(v string) *UpdateRemindRequest
	GetBaselineIds() *string
	SetBizProcessIds(v string) *UpdateRemindRequest
	GetBizProcessIds() *string
	SetDetail(v string) *UpdateRemindRequest
	GetDetail() *string
	SetDndEnd(v string) *UpdateRemindRequest
	GetDndEnd() *string
	SetMaxAlertTimes(v int32) *UpdateRemindRequest
	GetMaxAlertTimes() *int32
	SetNodeIds(v string) *UpdateRemindRequest
	GetNodeIds() *string
	SetProjectId(v int64) *UpdateRemindRequest
	GetProjectId() *int64
	SetRemindId(v int64) *UpdateRemindRequest
	GetRemindId() *int64
	SetRemindName(v string) *UpdateRemindRequest
	GetRemindName() *string
	SetRemindType(v string) *UpdateRemindRequest
	GetRemindType() *string
	SetRemindUnit(v string) *UpdateRemindRequest
	GetRemindUnit() *string
	SetRobotUrls(v string) *UpdateRemindRequest
	GetRobotUrls() *string
	SetUseFlag(v bool) *UpdateRemindRequest
	GetUseFlag() *bool
	SetWebhooks(v string) *UpdateRemindRequest
	GetWebhooks() *string
}

type UpdateRemindRequest struct {
	// The alert interval, in seconds. Minimum value: 1200. Default value: 1800.
	//
	// example:
	//
	// 1800
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The alert notification method. Valid values:
	//
	// - MAIL
	//
	// - SMS
	//
	// - PHONE. Only DataWorks Professional Edition and higher support phone alerts.
	//
	// - DINGROBOTS (DingTalk chatbot). This method takes effect only after the RobotUrls parameter is configured.
	//
	// - Webhooks (WeCom or Lark chatbot). This method takes effect only after the Webhooks parameter is configured.
	//
	// Separate multiple alert methods with commas (,).
	//
	// example:
	//
	// SMS,MAIL
	AlertMethods *string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty"`
	// The configuration details for different alert recipients:
	//
	// - When AlertUnit is set to OWNER (node owner), the configuration is left empty.
	//
	// - When AlertUnit is set to OTHER (specified user), set this parameter to the Alibaba Cloud UIDs of the specified users. Separate multiple UIDs with commas (,). You can specify up to 10 users to receive alerts.
	//
	// example:
	//
	// 9527952795279527
	AlertTargets *string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty"`
	// The recipient of the alert. Valid values:
	//
	// - OWNER: the node owner.
	//
	// - OTHER: a specified user.
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
	// The baseline IDs when the monitored object is a baseline. A rule can monitor up to 5 baselines. Separate multiple baseline IDs with commas (,).
	//
	// This parameter takes effect only when RemindUnit is set to BASELINE.
	//
	// example:
	//
	// 1,2,3
	BaselineIds *string `json:"BaselineIds,omitempty" xml:"BaselineIds,omitempty"`
	// The business process IDs when the monitored object is a business process. A rule can monitor up to 5 business processes. Separate multiple business process IDs with commas (,).
	//
	// This parameter takes effect only when RemindUnit is set to BIZPROCESS.
	//
	// example:
	//
	// 1,2,3
	BizProcessIds *string `json:"BizProcessIds,omitempty" xml:"BizProcessIds,omitempty"`
	// The configuration details for different trigger conditions:
	//
	// - When RemindType (trigger condition) is set to FINISHED, the configuration is left empty.
	//
	// - When RemindType (trigger condition) is set to UNFINISHED, the configuration format is {"hour":23,"minu":59}. Valid values of hour: [0,47\\]. Valid values of minu: [0,59\\].
	//
	// - When RemindType (trigger condition) is set to ERROR, the configuration is left empty.
	//
	// - When RemindType (trigger condition) is set to CYCLE_UNFINISHED (cycle unfinished), the configuration format is {"1":"05:50","2":"06:50","3":"07:50","4":"08:50","5":"09:50","6":"10:50","7":"11:50","8":"12:50","9":"13:50","10":"14:50","11":"15:50","12":"16:50","13":"17:50","14":"18:50","15":"19:50","16":"20:50","17":"21:50","18":"22:50","19":"23:50","20":"24:50","21":"25:50"}.
	//
	// The key in the JSON string is the cycle number. Valid values: [1,288\\]. The value is the unfinished time for the corresponding cycle in the format hh:mm. Valid values of hh: [0,47\\]. Valid values of mm: [0,59\\].
	//
	// - When RemindType (trigger condition) is set to TIMEOUT, the configuration format is 1800, in seconds. This means an alert is triggered if the instance has been running for more than 30 minutes.
	//
	// example:
	//
	// {"hour":23,"minu":59}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end time of the do-not-disturb period. Alerts are not sent before this time. Format: hh:mm. Valid values of hh: [0,23\\]. Valid values of mm: [0,59\\].
	//
	// example:
	//
	// 08:00
	DndEnd *string `json:"DndEnd,omitempty" xml:"DndEnd,omitempty"`
	// The maximum number of alerts. Valid values: [1,10\\]. Default value: 3.
	//
	// example:
	//
	// 3
	MaxAlertTimes *int32 `json:"MaxAlertTimes,omitempty" xml:"MaxAlertTimes,omitempty"`
	// The node IDs when the monitored object is a node. A rule can monitor up to 50 nodes. Separate multiple node IDs with commas (,).
	//
	// This parameter takes effect only when RemindUnit is set to NODE.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The workspace ID when the monitored object is a workspace. A rule can monitor only one workspace.
	//
	// This parameter takes effect only when RemindUnit is set to PROJECT.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the custom rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	// The name of the custom rule. The name cannot exceed 128 characters in length.
	//
	// example:
	//
	// update_remindname
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The condition that triggers the alert rule. Valid values:
	//
	// - FINISHED: The system monitors the instance from the start time and sends an alert when the node runs successfully.
	//
	// - UNFINISHED: The system monitors the instance from the start time and sends an alert if the node has not finished running by the specified target time.
	//
	// - ERROR: The system monitors the instance from the start time and sends an alert when the node encounters an error.
	//
	// - CYCLE_UNFINISHED: The system sends an alert if the instance has not finished running within the specified cycle. This is typically used to monitor instances that run on an hourly cycle.
	//
	// - TIMEOUT: The system monitors the instance from the start time and sends an alert if the node has not finished running after the specified duration. This is typically used to monitor the running duration of instances.
	//
	// For more information about alert trigger conditions, see [Custom rules](https://help.aliyun.com/document_detail/138172.html).
	//
	// example:
	//
	// FINISHED
	RemindType *string `json:"RemindType,omitempty" xml:"RemindType,omitempty"`
	// The type of the monitored object. Valid values:
	//
	// - NODE
	//
	// - BASELINE
	//
	// - PROJECT (workspace)
	//
	// - BIZPROCESS (business process)
	//
	// example:
	//
	// NODE
	RemindUnit *string `json:"RemindUnit,omitempty" xml:"RemindUnit,omitempty"`
	// The webhook URLs of DingTalk group chatbots. Separate multiple webhook URLs with commas (,).
	//
	// When the parameter settings are set to undefined, the system clears the DingTalk chatbot webhook URLs.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=******************************
	RobotUrls *string `json:"RobotUrls,omitempty" xml:"RobotUrls,omitempty"`
	// Specifies whether to enable the alert rule. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	UseFlag *bool `json:"UseFlag,omitempty" xml:"UseFlag,omitempty"`
	// The webhook URLs of WeCom or Lark chatbots. Separate multiple webhook URLs with commas (,). The alertMethods parameter must include the WEBHOOKS alerting method. When the parameter is set to undefined, the system clears the webhook URLs.
	//
	// Only DataWorks Enterprise Edition is supported.
	//
	// Active regions: China (Shanghai), China (Chengdu), China (Zhangjiakou), China (Beijing), China (Hangzhou), China (Shenzhen), Hong Kong (China), Germany (Frankfurt), Asia-Pacific Southeast 1 (Singapore).
	//
	// example:
	//
	// https://open.feishu.cn/open-apis/bot/v2/hook/*******
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s UpdateRemindRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRemindRequest) GoString() string {
	return s.String()
}

func (s *UpdateRemindRequest) GetAlertInterval() *int32 {
	return s.AlertInterval
}

func (s *UpdateRemindRequest) GetAlertMethods() *string {
	return s.AlertMethods
}

func (s *UpdateRemindRequest) GetAlertTargets() *string {
	return s.AlertTargets
}

func (s *UpdateRemindRequest) GetAlertUnit() *string {
	return s.AlertUnit
}

func (s *UpdateRemindRequest) GetBaselineIds() *string {
	return s.BaselineIds
}

func (s *UpdateRemindRequest) GetBizProcessIds() *string {
	return s.BizProcessIds
}

func (s *UpdateRemindRequest) GetDetail() *string {
	return s.Detail
}

func (s *UpdateRemindRequest) GetDndEnd() *string {
	return s.DndEnd
}

func (s *UpdateRemindRequest) GetMaxAlertTimes() *int32 {
	return s.MaxAlertTimes
}

func (s *UpdateRemindRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *UpdateRemindRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateRemindRequest) GetRemindId() *int64 {
	return s.RemindId
}

func (s *UpdateRemindRequest) GetRemindName() *string {
	return s.RemindName
}

func (s *UpdateRemindRequest) GetRemindType() *string {
	return s.RemindType
}

func (s *UpdateRemindRequest) GetRemindUnit() *string {
	return s.RemindUnit
}

func (s *UpdateRemindRequest) GetRobotUrls() *string {
	return s.RobotUrls
}

func (s *UpdateRemindRequest) GetUseFlag() *bool {
	return s.UseFlag
}

func (s *UpdateRemindRequest) GetWebhooks() *string {
	return s.Webhooks
}

func (s *UpdateRemindRequest) SetAlertInterval(v int32) *UpdateRemindRequest {
	s.AlertInterval = &v
	return s
}

func (s *UpdateRemindRequest) SetAlertMethods(v string) *UpdateRemindRequest {
	s.AlertMethods = &v
	return s
}

func (s *UpdateRemindRequest) SetAlertTargets(v string) *UpdateRemindRequest {
	s.AlertTargets = &v
	return s
}

func (s *UpdateRemindRequest) SetAlertUnit(v string) *UpdateRemindRequest {
	s.AlertUnit = &v
	return s
}

func (s *UpdateRemindRequest) SetBaselineIds(v string) *UpdateRemindRequest {
	s.BaselineIds = &v
	return s
}

func (s *UpdateRemindRequest) SetBizProcessIds(v string) *UpdateRemindRequest {
	s.BizProcessIds = &v
	return s
}

func (s *UpdateRemindRequest) SetDetail(v string) *UpdateRemindRequest {
	s.Detail = &v
	return s
}

func (s *UpdateRemindRequest) SetDndEnd(v string) *UpdateRemindRequest {
	s.DndEnd = &v
	return s
}

func (s *UpdateRemindRequest) SetMaxAlertTimes(v int32) *UpdateRemindRequest {
	s.MaxAlertTimes = &v
	return s
}

func (s *UpdateRemindRequest) SetNodeIds(v string) *UpdateRemindRequest {
	s.NodeIds = &v
	return s
}

func (s *UpdateRemindRequest) SetProjectId(v int64) *UpdateRemindRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateRemindRequest) SetRemindId(v int64) *UpdateRemindRequest {
	s.RemindId = &v
	return s
}

func (s *UpdateRemindRequest) SetRemindName(v string) *UpdateRemindRequest {
	s.RemindName = &v
	return s
}

func (s *UpdateRemindRequest) SetRemindType(v string) *UpdateRemindRequest {
	s.RemindType = &v
	return s
}

func (s *UpdateRemindRequest) SetRemindUnit(v string) *UpdateRemindRequest {
	s.RemindUnit = &v
	return s
}

func (s *UpdateRemindRequest) SetRobotUrls(v string) *UpdateRemindRequest {
	s.RobotUrls = &v
	return s
}

func (s *UpdateRemindRequest) SetUseFlag(v bool) *UpdateRemindRequest {
	s.UseFlag = &v
	return s
}

func (s *UpdateRemindRequest) SetWebhooks(v string) *UpdateRemindRequest {
	s.Webhooks = &v
	return s
}

func (s *UpdateRemindRequest) Validate() error {
	return dara.Validate(s)
}
