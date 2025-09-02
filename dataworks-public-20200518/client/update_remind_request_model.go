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
	// The intervals at which alert notifications are sent. Unit: seconds. Minimum value: 1200. Default value: 1800.
	//
	// example:
	//
	// 1800
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The notification method. Valid values:
	//
	// 	- MAIL: Alert notifications are sent by email.
	//
	// 	- SMS: Alert notifications are sent by text message.
	//
	// 	- PHONE: Alert notifications are sent by phone call. You can use this notification method only in DataWorks Professional Edition or more advanced editions.
	//
	// 	- DINGROBOTS: Alert notifications are sent by DingTalk message. You can use this notification method only if the RobotUrls parameter is configured.
	//
	// 	- WEBHOOKS (WeCom or Lark chatbot): Alert notifications are sent by WeCom or Lark message. You can use this notification method only if the Webhooks parameter is configured.
	//
	// Multiple notification methods are separated by commas (,).
	//
	// example:
	//
	// SMS,MAIL
	AlertMethods *string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty"`
	// The value format required by this parameter varies based on the value that you specify for the AlertUnit parameter. Take note of the following items:
	//
	// 	- If the AlertUnit parameter is set to OWNER, leave this parameter empty.
	//
	// 	- If the AlertUnit parameter is set to OTHER, set this parameter to the unique ID (UID) of the specified user. You can specify multiple UIDs. Separate them with commas (,). A maximum of 10 UIDs can be specified for receiving alert notifications.
	//
	// example:
	//
	// 9527952795279527
	AlertTargets *string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty"`
	// The recipient to whom alert notifications are sent. Valid values: OWNER and OTHER. The value OWNER indicates that alert notifications are sent to the object owner. The value OTHER indicates that alert notifications are sent to a specified user.
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
	// The ID of the baseline to which the custom alert rule is applied. A maximum of 5 baselines can be specified for a custom alert rule. You can specify multiple IDs. Separate multiple IDs with commas (,). This parameter takes effect when you set the RemindUnit parameter to BASELINE.
	//
	// example:
	//
	// 1,2,3
	BaselineIds *string `json:"BaselineIds,omitempty" xml:"BaselineIds,omitempty"`
	// The ID of the workflow to which the custom alert rule is applied. A maximum of 5 workflows can be specified for a custom alert rule. You can specify multiple IDs. Separate multiple IDs with commas (,). This parameter takes effect when you set the RemindUnit parameter to BIZPROCESS.
	//
	// example:
	//
	// 1,2,3
	BizProcessIds *string `json:"BizProcessIds,omitempty" xml:"BizProcessIds,omitempty"`
	// The details of the conditions that trigger an alert.
	//
	// 	- If the RemindType parameter is set to FINISHED, leave this parameter empty.
	//
	// 	- If the RemindType parameter is set to UNFINISHED, set this parameter to key-value pairs. Example: {"hour":23,"minu":59}. Valid values of hour: [0,47]. Valid values of minu: [0,59].
	//
	// 	- If the RemindType parameter is set to ERROR, leave this parameter empty.
	//
	// 	- If the RemindType parameter is set to CYCLE_UNFINISHED, set this parameter to key-value pairs in the JSON format. Example: {"1":"05:50","2":"06:50","3":"07:50","4":"08:50","5":"09:50","6":"10:50","7":"11:50","8":"12:50","9":"13:50","10":"14:50","11":"15:50","12":"16:50","13":"17:50","14":"18:50","15":"19:50","16":"20:50","17":"21:50","18":"22:50","19":"23:50","20":"24:50","21":"25:50"}. A key in the JSON string indicates the sequence number of a cycle. Valid values of keys: 1 to 288. A value in the JSON string indicates the time in point when a monitored instance times out in the relevant cycle. Values must be in the format of hh:mm. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// 	- If the RemindType parameter is set to TIMEOUT, set this parameter to the timeout period. Unit: seconds. Example: 1800. This indicates that an alert notification is sent if the running duration of a monitored instance exceeds 30 minutes.
	//
	// example:
	//
	// {"hour":23,"minu":59}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end of the period during which no alert notifications are sent. Specify the time in the hh:mm format. Valid values of hh: [0,23]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 08:00
	DndEnd *string `json:"DndEnd,omitempty" xml:"DndEnd,omitempty"`
	// The maximum number of alerts. Valid values: 1 to 10. Default value: 3.
	//
	// example:
	//
	// 3
	MaxAlertTimes *int32 `json:"MaxAlertTimes,omitempty" xml:"MaxAlertTimes,omitempty"`
	// The ID of the node to which the custom alert rule is applied. A maximum of 50 nodes can be specified for a custom alert rule. You can specify multiple IDs. Separate multiple IDs with commas (,). This parameter takes effect when you set the RemindUnit parameter to NODE.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The ID of the workspace to which the custom alert rule is applied. You can specify only one workspace for a custom alert rule. This parameter takes effect when you set the RemindUnit parameter to PROJECT.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The custom alert rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	RemindId *int64 `json:"RemindId,omitempty" xml:"RemindId,omitempty"`
	// The name of the custom alert rule. The name cannot exceed 128 characters in length.
	//
	// example:
	//
	// update_remindname
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The condition that triggers the alert rule. Valid values:
	//
	// 	- FINISHED: The system monitors an instance when it starts to run and sends an alert notification after the running of the instance is complete.
	//
	// 	- UNFINISHED: The system monitors an instance when it starts to run and sends an alert notification if the instance is still running at the specified point in time.
	//
	// 	- ERROR: The system monitors an instance when it starts to run and sends an alert notification if an error occurs.
	//
	// 	- CYCLE_UNFINISHED: The system sends an alert notification if a monitored instance is still running at the end of the specified cycle. In most cases, you can configure this trigger condition for node instances that are scheduled to run by hour.
	//
	// 	- TIMEOUT: The system monitors an instance when it starts to run and sends an alert notification if the instance is still running after the specified period ends. In most cases, you can configure this trigger condition to monitor the running duration of node instances.
	//
	// For more information, see [Manage custom alert rules](https://help.aliyun.com/document_detail/138172.html).
	//
	// example:
	//
	// FINISHED
	RemindType *string `json:"RemindType,omitempty" xml:"RemindType,omitempty"`
	// The type of the object to which the custom alert rule is applied. Valid values:
	//
	// 	- NODE
	//
	// 	- BASELINE
	//
	// 	- PROJECT
	//
	// 	- BIZPROCESS
	//
	// example:
	//
	// NODE
	RemindUnit *string `json:"RemindUnit,omitempty" xml:"RemindUnit,omitempty"`
	// The webhook URL of the DingTalk chatbot. You can specify multiple webhook URLs. Separate multiple webhook URLs with commas (,). If this parameter is set to undefined, the specified webhook URLs are cleared.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=******************************
	RobotUrls *string `json:"RobotUrls,omitempty" xml:"RobotUrls,omitempty"`
	// Specifies whether to enable the alert rule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	UseFlag *bool `json:"UseFlag,omitempty" xml:"UseFlag,omitempty"`
	// The webhook URL of the WeCom or Lark chatbot. You can specify multiple webhook URLs. Separate multiple webhook URLs with commas (,). The value of AlertMethods must include WEBHOOKS. If this parameter is set to undefined, the specified webhook URLs are cleared.
	//
	// Only DataWorks Enterprise Edition supports this parameter. The webhook URL-based alerting feature is supported in the following regions: China (Shanghai), China (Chengdu), China (Zhangjiakou), China (Beijing), China (Hangzhou), China (Shenzhen), China (Hong Kong), Germany (Frankfurt), and Singapore.
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
