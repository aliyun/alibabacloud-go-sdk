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
	// The minimum alert interval, in seconds. Minimum value: 1200. Default value: 1800.
	//
	// example:
	//
	// 1800
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The alert method. Valid values:
	//
	// - MAIL: email.
	//
	// - SMS: text message.
	//
	// <props="intl">The regions that support SMS alerts are Singapore, Malaysia (Kuala Lumpur), and Germany (Frankfurt).
	//
	// <props="china">- PHONE: phone call. Only DataWorks Professional Edition and higher editions are supported.
	//
	// - Webhooks (WeCom or Lark chatbot). This alert method takes effect only after the Webhooks parameter is configured.
	//
	// - DINGROBOTS: DingTalk chatbot.
	//
	// Separate multiple alert methods with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS,MAIL
	AlertMethods *string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty"`
	// - When AlertUnit (alert recipient) is set to OWNER (node owner), pass an empty value.
	//
	// - When AlertUnit (alert recipient) is set to OTHER (specified user), pass the Alibaba Cloud UIDs of the specified users. Separate multiple Alibaba Cloud UIDs with commas (,). A maximum of 10 UIDs are supported.
	//
	// example:
	//
	// 9527952795279527
	AlertTargets *string `json:"AlertTargets,omitempty" xml:"AlertTargets,omitempty"`
	// The granularity of the alert recipient. Valid values: OWNER (node owner) and OTHER (specified user).
	//
	// This parameter is required.
	//
	// example:
	//
	// OWNER
	AlertUnit *string `json:"AlertUnit,omitempty" xml:"AlertUnit,omitempty"`
	// The IDs of the baselines to monitor when RemindUnit (object type) is set to BASELINE (baseline). Separate multiple IDs with commas (,). A maximum of 5 baselines can be monitored by a single rule.
	//
	// example:
	//
	// 1,2,3
	BaselineIds *string `json:"BaselineIds,omitempty" xml:"BaselineIds,omitempty"`
	// The IDs of the business processes to monitor when RemindUnit (object type) is set to BIZPROCESS (business process). Separate multiple business process IDs with commas (,). A maximum of 5 business processes can be monitored by a single rule.
	//
	// example:
	//
	// 1,2,3
	BizProcessIds *string `json:"BizProcessIds,omitempty" xml:"BizProcessIds,omitempty"`
	// The descriptions for different trigger conditions are as follows:
	//
	// - When RemindType (trigger condition) is set to FINISHED (completed), pass an empty value.
	//
	// - When RemindType (trigger condition) is set to UNFINISHED (not completed), pass parameter in the format of {"hour":23,"minu":59}. Valid values of hour: [0,47\\]. Valid values of minu: [0,59\\].
	//
	// - When RemindType (trigger condition) is set to ERROR (error), pass an empty value.
	//
	// - When RemindType (trigger condition) is set to CYCLE_UNFINISHED (cycle not completed), pass parameter in the format of {"1":"05:50","2":"06:50","3":"07:50","4":"08:50","5":"09:50","6":"10:50","7":"11:50","8":"12:50","9":"13:50","10":"14:50","11":"15:50","12":"16:50","13":"17:50","14":"18:50","15":"19:50","16":"20:50","17":"21:50","18":"22:50","19":"23:50","20":"24:50","21":"25:50"}. The key in the JSON character string is the cycle number. Valid values: [1,288\\]. The value is the not-completed time for the corresponding cycle, in the hh:mm format. Valid values of hh: [0,47\\]. Valid values of mm: [0,59\\].
	//
	// - When RemindType (trigger condition) is set to TIMEOUT (running timeout), pass parameter as a value such as 1800, in seconds. This means that an alert is triggered if the running time exceeds 30 minutes from the start of execution.
	//
	// example:
	//
	// {"hour":"23","minu":"59"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end time of the do-not-disturb period, in the hh:mm format. Valid values of hh: [0,23\\]. Valid values of mm: [0,59\\].
	//
	// example:
	//
	// 08:00
	DndEnd *string `json:"DndEnd,omitempty" xml:"DndEnd,omitempty"`
	// The maximum number of alerts. Minimum value: 1. Maximum value: 10. Default value: 3.
	//
	// example:
	//
	// 2
	MaxAlertTimes *int32 `json:"MaxAlertTimes,omitempty" xml:"MaxAlertTimes,omitempty"`
	// The IDs of the nodes to monitor when RemindUnit (object type) is set to NODE (node). Separate multiple IDs with commas (,). A maximum of 50 nodes can be monitored by a single rule.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The ID of the workspace to monitor when RemindUnit (object type) is set to PROJECT (workspace). A single rule can monitor only one workspace.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the custom rule. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_error_remind
	RemindName *string `json:"RemindName,omitempty" xml:"RemindName,omitempty"`
	// The trigger condition. Valid values: FINISHED (completed), UNFINISHED (not completed), ERROR (error), CYCLE_UNFINISHED (cycle not completed), and TIMEOUT (running timeout).
	//
	// This parameter is required.
	//
	// example:
	//
	// FINISHED
	RemindType *string `json:"RemindType,omitempty" xml:"RemindType,omitempty"`
	// The type of the object. Valid values: NODE (node), BASELINE (baseline), PROJECT (workspace), and BIZPROCESS (business process).
	//
	// This parameter is required.
	//
	// example:
	//
	// NODE
	RemindUnit *string `json:"RemindUnit,omitempty" xml:"RemindUnit,omitempty"`
	// The webhook URLs of DingTalk chatbots. Separate multiple webhook URLs with commas (,).
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=******************************
	RobotUrls *string `json:"RobotUrls,omitempty" xml:"RobotUrls,omitempty"`
	// The webhook URLs of WeCom or Lark chatbots. Separate multiple webhook URLs with commas (,). The alertMethods parameter must include the WEBHOOKS alert method.
	//
	// Only DataWorks Enterprise Edition is supported.
	//
	// Available regions: China (Shanghai), China (Chengdu), China (Zhangjiakou), China (Beijing), China (Hangzhou), China (Shenzhen), Hong Kong (China), Germany (Frankfurt), and Singapore.
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
