// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRemindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertInterval(v int32) *CreateRemindRequest
	GetAlertInterval() *int32
	SetAlertMethods(v string) *CreateRemindRequest
	GetAlertMethods() *string
	SetAlertTargets(v string) *CreateRemindRequest
	GetAlertTargets() *string
	SetAlertUnit(v string) *CreateRemindRequest
	GetAlertUnit() *string
	SetBaselineIds(v string) *CreateRemindRequest
	GetBaselineIds() *string
	SetBizProcessIds(v string) *CreateRemindRequest
	GetBizProcessIds() *string
	SetDetail(v string) *CreateRemindRequest
	GetDetail() *string
	SetDndEnd(v string) *CreateRemindRequest
	GetDndEnd() *string
	SetMaxAlertTimes(v int32) *CreateRemindRequest
	GetMaxAlertTimes() *int32
	SetNodeIds(v string) *CreateRemindRequest
	GetNodeIds() *string
	SetProjectId(v int64) *CreateRemindRequest
	GetProjectId() *int64
	SetRemindName(v string) *CreateRemindRequest
	GetRemindName() *string
	SetRemindType(v string) *CreateRemindRequest
	GetRemindType() *string
	SetRemindUnit(v string) *CreateRemindRequest
	GetRemindUnit() *string
	SetRobotUrls(v string) *CreateRemindRequest
	GetRobotUrls() *string
	SetWebhooks(v string) *CreateRemindRequest
	GetWebhooks() *string
}

type CreateRemindRequest struct {
	// The minimum interval at which alerts are reported. Unit: seconds. Minimum value: 1200. Default value: 1800.
	//
	// example:
	//
	// 1800
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The notification method. Valid values:
	//
	// 	- MAIL: Alert notifications are sent by email.
	//
	// 	- SMS: Alert notifications are sent by text message. Alert notifications can be sent by text message only in the Singapore, Malaysia (Kuala Lumpur), and Germany (Frankfurt) regions.
	//
	// 	- WEBHOOKS (WeCom or Lark chatbot): Alert notifications are sent by WeCom or Lark message. If you want to use this notification method, you must configure the Webhooks parameter.
	//
	// You can specify multiple notification methods. Separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS,MAIL
	AlertMethods *string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty"`
	// 	- If the AlertUnit parameter is set to OWNER, leave this parameter empty.
	//
	// 	- If the AlertUnit parameter is set to OTHER, set this parameter to the ID of the Alibaba Cloud account used by the specified user. You can specify multiple IDs. Separate multiple IDs with commas (,). You can specify a maximum of 10 IDs.
	//
	// example:
	//
	// 9527952795279527
	AlertTargets *string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty"`
	// The recipient of the alert. Valid values: OWNER and OTHER. The value OWNER indicates the node owner. The value OTHER indicates a specified user.
	//
	// This parameter is required.
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
	// The ID of the baseline to which the custom alert rule is applied. This parameter takes effect when the RemindUnit parameter is set to BASELINE. You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of five baselines can be specified for a custom alert rule.
	//
	// example:
	//
	// 1,2,3
	BaselineIds *string `json:"BaselineIds,omitempty" xml:"BaselineIds,omitempty"`
	// The ID of the workflow to which the custom alert rule is applied. This parameter takes effect when the RemindUnit parameter is set to BIZPROCESS. You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of five workflows can be specified for a custom alert rule.
	//
	// example:
	//
	// 1,2,3
	BizProcessIds *string `json:"BizProcessIds,omitempty" xml:"BizProcessIds,omitempty"`
	// The details of the conditions that trigger an alert.
	//
	// 	- If the RemindType parameter is set to FINISHED, leave this parameter empty.
	//
	// 	- If the RemindType parameter is set to UNFINISHED, configure this parameter as key-value pairs. Example: {"hour":23,"minu":59}. Valid values of hour: [0,47]. Valid values of minu: [0,59].
	//
	// 	- If the RemindType parameter is set to ERROR, leave this parameter empty.
	//
	// 	- If the RemindType parameter is set to CYCLE_UNFINISHED, configure this parameter as key-value pairs. Example: {"1":"05:50","2":"06:50","3":"07:50","4":"08:50","5":"09:50","6":"10:50","7":"11:50","8":"12:50","9":"13:50","10":"14:50","11":"15:50","12":"16:50","13":"17:50","14":"18:50","15":"19:50","16":"20:50","17":"21:50","18":"22:50","19":"23:50","20":"24:50","21":"25:50"}. The key indicates the ID of the cycle. Valid values: [1,288]. The value indicates the timeout period of the node that is running in the cycle. Specify the value in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// 	- If the RemindType parameter is set to TIMEOUT, set this parameter to the timeout period. Unit: seconds. Example: 1800. This value indicates that an alert is reported if the node has run for more than 30 minutes.
	//
	// example:
	//
	// {"hour":23,"minu":59}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end time of the quiet hours. Specify the time in the hh:mm format. Valid values of hh: [0,23]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 08:00
	DndEnd *string `json:"DndEnd,omitempty" xml:"DndEnd,omitempty"`
	// The maximum number of alerts. Valid values: 1 to 10. Default value: 3.
	//
	// example:
	//
	// 2
	MaxAlertTimes *int32 `json:"MaxAlertTimes,omitempty" xml:"MaxAlertTimes,omitempty"`
	// The ID of the node to which the custom alert rule is applied. This parameter takes effect when the RemindUnit parameter is set to NODE. You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of 50 nodes can be specified for a custom alert rule.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The ID of the workspace to which the custom alert rule is applied. This parameter takes effect when the RemindUnit parameter is set to PROJECT. You can specify only one workspace for a custom alert rule.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the custom alert rule. The name cannot exceed 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_error_remind
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The conditions that trigger an alert. Valid values: FINISHED, UNFINISHED, ERROR, CYCLE_UNFINISHED, and TIMEOUT.
	//
	// This parameter is required.
	//
	// example:
	//
	// FINISHED
	RemindType *string `json:"RemindType,omitempty" xml:"RemindType,omitempty"`
	// The type of the object to which the custom alert rule is applied. Valid values: NODE, BASELINE, PROJECT, and BIZPROCESS. The value NODE indicates a node. The value BASELINE indicates a baseline. The value PROJECT indicates a workspace. The value BIZPROCESS indicates a workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// NODE
	RemindUnit *string `json:"RemindUnit,omitempty" xml:"RemindUnit,omitempty"`
	// The webhook URL of the DingTalk chatbot. You can specify multiple webhook URLs. Separate multiple webhook URLs with commas (,).
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=******************************
	RobotUrls *string `json:"RobotUrls,omitempty" xml:"RobotUrls,omitempty"`
	// The webhook URL of the WeCom or Lark chatbot. You can specify multiple webhook URLs. Separate multiple webhook URLs with commas (,). You must specify WEBHOOKS for AlertMethods.
	//
	// Only DataWorks Enterprise Edition supports this parameter. The webhook URL-based alerting feature is supported in the following regions: China (Shanghai), China (Chengdu), China (Zhangjiakou), China (Beijing), China (Hangzhou), China (Shenzhen), China (Hong Kong), Germany (Frankfurt), and Singapore.
	//
	// example:
	//
	// https://open.feishu.cn/open-apis/bot/v2/hook/*******
	Webhooks *string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty"`
}

func (s CreateRemindRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRemindRequest) GoString() string {
	return s.String()
}

func (s *CreateRemindRequest) GetAlertInterval() *int32 {
	return s.AlertInterval
}

func (s *CreateRemindRequest) GetAlertMethods() *string {
	return s.AlertMethods
}

func (s *CreateRemindRequest) GetAlertTargets() *string {
	return s.AlertTargets
}

func (s *CreateRemindRequest) GetAlertUnit() *string {
	return s.AlertUnit
}

func (s *CreateRemindRequest) GetBaselineIds() *string {
	return s.BaselineIds
}

func (s *CreateRemindRequest) GetBizProcessIds() *string {
	return s.BizProcessIds
}

func (s *CreateRemindRequest) GetDetail() *string {
	return s.Detail
}

func (s *CreateRemindRequest) GetDndEnd() *string {
	return s.DndEnd
}

func (s *CreateRemindRequest) GetMaxAlertTimes() *int32 {
	return s.MaxAlertTimes
}

func (s *CreateRemindRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *CreateRemindRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateRemindRequest) GetRemindName() *string {
	return s.RemindName
}

func (s *CreateRemindRequest) GetRemindType() *string {
	return s.RemindType
}

func (s *CreateRemindRequest) GetRemindUnit() *string {
	return s.RemindUnit
}

func (s *CreateRemindRequest) GetRobotUrls() *string {
	return s.RobotUrls
}

func (s *CreateRemindRequest) GetWebhooks() *string {
	return s.Webhooks
}

func (s *CreateRemindRequest) SetAlertInterval(v int32) *CreateRemindRequest {
	s.AlertInterval = &v
	return s
}

func (s *CreateRemindRequest) SetAlertMethods(v string) *CreateRemindRequest {
	s.AlertMethods = &v
	return s
}

func (s *CreateRemindRequest) SetAlertTargets(v string) *CreateRemindRequest {
	s.AlertTargets = &v
	return s
}

func (s *CreateRemindRequest) SetAlertUnit(v string) *CreateRemindRequest {
	s.AlertUnit = &v
	return s
}

func (s *CreateRemindRequest) SetBaselineIds(v string) *CreateRemindRequest {
	s.BaselineIds = &v
	return s
}

func (s *CreateRemindRequest) SetBizProcessIds(v string) *CreateRemindRequest {
	s.BizProcessIds = &v
	return s
}

func (s *CreateRemindRequest) SetDetail(v string) *CreateRemindRequest {
	s.Detail = &v
	return s
}

func (s *CreateRemindRequest) SetDndEnd(v string) *CreateRemindRequest {
	s.DndEnd = &v
	return s
}

func (s *CreateRemindRequest) SetMaxAlertTimes(v int32) *CreateRemindRequest {
	s.MaxAlertTimes = &v
	return s
}

func (s *CreateRemindRequest) SetNodeIds(v string) *CreateRemindRequest {
	s.NodeIds = &v
	return s
}

func (s *CreateRemindRequest) SetProjectId(v int64) *CreateRemindRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateRemindRequest) SetRemindName(v string) *CreateRemindRequest {
	s.RemindName = &v
	return s
}

func (s *CreateRemindRequest) SetRemindType(v string) *CreateRemindRequest {
	s.RemindType = &v
	return s
}

func (s *CreateRemindRequest) SetRemindUnit(v string) *CreateRemindRequest {
	s.RemindUnit = &v
	return s
}

func (s *CreateRemindRequest) SetRobotUrls(v string) *CreateRemindRequest {
	s.RobotUrls = &v
	return s
}

func (s *CreateRemindRequest) SetWebhooks(v string) *CreateRemindRequest {
	s.Webhooks = &v
	return s
}

func (s *CreateRemindRequest) Validate() error {
	return dara.Validate(s)
}
