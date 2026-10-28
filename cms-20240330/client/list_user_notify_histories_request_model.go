// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserNotifyHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *ListUserNotifyHistoriesRequest
	GetBatchId() *string
	SetBizSource(v string) *ListUserNotifyHistoriesRequest
	GetBizSource() *string
	SetChannel(v string) *ListUserNotifyHistoriesRequest
	GetChannel() *string
	SetChannelCategory(v string) *ListUserNotifyHistoriesRequest
	GetChannelCategory() *string
	SetFrom(v int32) *ListUserNotifyHistoriesRequest
	GetFrom() *int32
	SetGroup(v string) *ListUserNotifyHistoriesRequest
	GetGroup() *string
	SetIncidentId(v string) *ListUserNotifyHistoriesRequest
	GetIncidentId() *string
	SetIncludeNoChannel(v bool) *ListUserNotifyHistoriesRequest
	GetIncludeNoChannel() *bool
	SetIsDesc(v bool) *ListUserNotifyHistoriesRequest
	GetIsDesc() *bool
	SetMaxResults(v int32) *ListUserNotifyHistoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserNotifyHistoriesRequest
	GetNextToken() *string
	SetNotifyStrategyId(v string) *ListUserNotifyHistoriesRequest
	GetNotifyStrategyId() *string
	SetNotifyStrategyName(v string) *ListUserNotifyHistoriesRequest
	GetNotifyStrategyName() *string
	SetPartitionKey(v string) *ListUserNotifyHistoriesRequest
	GetPartitionKey() *string
	SetReceiver(v string) *ListUserNotifyHistoriesRequest
	GetReceiver() *string
	SetRuleId(v string) *ListUserNotifyHistoriesRequest
	GetRuleId() *string
	SetRuleName(v string) *ListUserNotifyHistoriesRequest
	GetRuleName() *string
	SetSendType(v string) *ListUserNotifyHistoriesRequest
	GetSendType() *string
	SetSeverity(v string) *ListUserNotifyHistoriesRequest
	GetSeverity() *string
	SetSubscriptionId(v string) *ListUserNotifyHistoriesRequest
	GetSubscriptionId() *string
	SetSubscriptionName(v string) *ListUserNotifyHistoriesRequest
	GetSubscriptionName() *string
	SetSuccess(v bool) *ListUserNotifyHistoriesRequest
	GetSuccess() *bool
	SetTo(v int32) *ListUserNotifyHistoriesRequest
	GetTo() *int32
	SetWorkspace(v string) *ListUserNotifyHistoriesRequest
	GetWorkspace() *string
}

type ListUserNotifyHistoriesRequest struct {
	// The notification batch ID. Used to filter the notification history for notifications delivered in the same batch.
	//
	// example:
	//
	// 3fa85f64-5717-4562-b3fc-2c963f66afa7
	BatchId *string `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This is an internal reserved parameter used only for business source filtering in internal platform scenarios. Do not pass this parameter in external calls.
	//
	// example:
	//
	// arms
	BizSource *string `json:"bizSource,omitempty" xml:"bizSource,omitempty"`
	// The notification channel. Valid values:
	//
	// - ALI_YUNGE_SMS: text message
	//
	// - ALI_YUNGE_CALL: voice call
	//
	// - ALI_YUNGE_MAIL: email
	//
	// - ALI_YUNGE_ALIUID: internal message
	//
	// - DING: DingTalk
	//
	// - DING_COOL_APP: DingTalk Cool App
	//
	// - FEISHU: Lark
	//
	// - WEIXIN: WeCom
	//
	// - SLACK: Slack
	//
	// - TEAMS: Teams
	//
	// - PAGERDUTY: PagerDuty
	//
	// - WEBHOOK: custom webhook
	//
	// - ALIYUN_SLS: Simple Log Service
	//
	// - ALIYUN_FC: Function Compute
	//
	// - ALIYUN_FC3: Function Compute 3.0
	//
	// - ALIYUN_MNS: Message Service
	//
	// - ALIYUN_OOS: CloudOps Orchestration Service
	//
	// - ALIYUN_ESS: elastic scaling
	//
	// - ALIYUN_EB: event bus
	//
	// - ALIYUN_OPENAPI: OpenAPI
	//
	// example:
	//
	// DING
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The channel category. Valid values:
	//
	// - WEBHOOK: webhook-based channels, such as DingTalk, Lark, WeCom, Slack, Teams, and custom webhooks
	//
	// - YUNGE: Yunge communication channels, such as text messages, voice calls, and emails
	//
	// example:
	//
	// WEBHOOK
	ChannelCategory *string `json:"channelCategory,omitempty" xml:"channelCategory,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in seconds. Default value: 6 hours before the `to` parameter value.
	//
	// example:
	//
	// 1729400000
	From *int32 `json:"from,omitempty" xml:"from,omitempty"`
	// The contact group. Used to filter the notification history by contact group name.
	//
	// example:
	//
	// O&M contact group
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The incident ID. Used to filter the notification history associated with a specified incident.
	//
	// example:
	//
	// 5d41402abc4b2a76b9719d911017c592
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// Specifies whether to include notification records that have no actual delivery channel. Valid values:
	//
	// - true: The records are included.
	//
	// - false: The records are not included. Only records that were actually delivered through a channel are returned.
	//
	// Default value: true.
	IncludeNoChannel *bool `json:"includeNoChannel,omitempty" xml:"includeNoChannel,omitempty"`
	// Specifies whether to sort the results in descending order of time. Valid values:
	//
	// - true: The results are sorted in descending order, with the latest records appearing first.
	//
	// - false: The results are sorted in ascending order.
	//
	// Default value: true.
	IsDesc *bool `json:"isDesc,omitempty" xml:"isDesc,omitempty"`
	// The maximum number of entries to return per page. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Used to retrieve the next page of data. It is an opaque string generated and encrypted by the server. You do not need to specify this parameter for the first query. For subsequent queries, pass the token returned in the previous response.
	//
	// example:
	//
	// 2f8aKz3mNq9p****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The notification policy ID. Used to filter the notification history for a specified notification policy. It is a 32-character string.
	//
	// example:
	//
	// 2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// The notification policy name. Used to filter the notification history by notification policy name.
	//
	// example:
	//
	// Default notification policy
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- This is an internal reserved parameter used only for dedicated notification history query paths in internal platform scenarios. Do not pass this parameter in external calls.
	//
	// example:
	//
	// 1234567890123456
	PartitionKey *string `json:"partitionKey,omitempty" xml:"partitionKey,omitempty"`
	// The recipient. Used to filter the notification history by recipient address, such as a phone number, email address, or webhook URL.
	//
	// example:
	//
	// 13800138000
	Receiver *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// The alert rule ID. Used to filter the notification history triggered by a specified alert rule. Exact match is applied.
	//
	// example:
	//
	// 3fa85f64-5717-4562-b3fc-2c963f66afa6
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The alert rule name. Used to filter the notification history by rule name.
	//
	// example:
	//
	// CPU usage alert
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// The notification delivery type. Valid values:
	//
	// - SUBSCRIBE: subscription notifications
	//
	// - DIRECT: direct event notifications
	//
	// - ITSM: ticket or action notifications
	//
	// example:
	//
	// SUBSCRIBE
	SendType *string `json:"sendType,omitempty" xml:"sendType,omitempty"`
	// The alert severity. Valid values:
	//
	// - CRITICAL: critical
	//
	// - ERROR: error
	//
	// - WARNING: warning
	//
	// - INFO: information
	//
	// example:
	//
	// CRITICAL
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The subscription ID. Used to filter the notification history for a specified subscription. It is a 32-character string.
	//
	// example:
	//
	// 1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	// The subscription name. Used to filter the notification history by subscription name.
	//
	// example:
	//
	// Default subscription
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	// Specifies whether to filter by delivery result. Valid values:
	//
	// - true: Only successfully delivered records are returned.
	//
	// - false: Only failed delivery records are returned.
	//
	// If you do not specify this parameter, the results are not filtered by delivery result.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in seconds. Default value: the current time.
	//
	// example:
	//
	// 1729421600
	To *int32 `json:"to,omitempty" xml:"to,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListUserNotifyHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserNotifyHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListUserNotifyHistoriesRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *ListUserNotifyHistoriesRequest) GetBizSource() *string {
	return s.BizSource
}

func (s *ListUserNotifyHistoriesRequest) GetChannel() *string {
	return s.Channel
}

func (s *ListUserNotifyHistoriesRequest) GetChannelCategory() *string {
	return s.ChannelCategory
}

func (s *ListUserNotifyHistoriesRequest) GetFrom() *int32 {
	return s.From
}

func (s *ListUserNotifyHistoriesRequest) GetGroup() *string {
	return s.Group
}

func (s *ListUserNotifyHistoriesRequest) GetIncidentId() *string {
	return s.IncidentId
}

func (s *ListUserNotifyHistoriesRequest) GetIncludeNoChannel() *bool {
	return s.IncludeNoChannel
}

func (s *ListUserNotifyHistoriesRequest) GetIsDesc() *bool {
	return s.IsDesc
}

func (s *ListUserNotifyHistoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserNotifyHistoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserNotifyHistoriesRequest) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *ListUserNotifyHistoriesRequest) GetNotifyStrategyName() *string {
	return s.NotifyStrategyName
}

func (s *ListUserNotifyHistoriesRequest) GetPartitionKey() *string {
	return s.PartitionKey
}

func (s *ListUserNotifyHistoriesRequest) GetReceiver() *string {
	return s.Receiver
}

func (s *ListUserNotifyHistoriesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *ListUserNotifyHistoriesRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ListUserNotifyHistoriesRequest) GetSendType() *string {
	return s.SendType
}

func (s *ListUserNotifyHistoriesRequest) GetSeverity() *string {
	return s.Severity
}

func (s *ListUserNotifyHistoriesRequest) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListUserNotifyHistoriesRequest) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *ListUserNotifyHistoriesRequest) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserNotifyHistoriesRequest) GetTo() *int32 {
	return s.To
}

func (s *ListUserNotifyHistoriesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListUserNotifyHistoriesRequest) SetBatchId(v string) *ListUserNotifyHistoriesRequest {
	s.BatchId = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetBizSource(v string) *ListUserNotifyHistoriesRequest {
	s.BizSource = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetChannel(v string) *ListUserNotifyHistoriesRequest {
	s.Channel = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetChannelCategory(v string) *ListUserNotifyHistoriesRequest {
	s.ChannelCategory = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetFrom(v int32) *ListUserNotifyHistoriesRequest {
	s.From = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetGroup(v string) *ListUserNotifyHistoriesRequest {
	s.Group = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetIncidentId(v string) *ListUserNotifyHistoriesRequest {
	s.IncidentId = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetIncludeNoChannel(v bool) *ListUserNotifyHistoriesRequest {
	s.IncludeNoChannel = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetIsDesc(v bool) *ListUserNotifyHistoriesRequest {
	s.IsDesc = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetMaxResults(v int32) *ListUserNotifyHistoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetNextToken(v string) *ListUserNotifyHistoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetNotifyStrategyId(v string) *ListUserNotifyHistoriesRequest {
	s.NotifyStrategyId = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetNotifyStrategyName(v string) *ListUserNotifyHistoriesRequest {
	s.NotifyStrategyName = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetPartitionKey(v string) *ListUserNotifyHistoriesRequest {
	s.PartitionKey = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetReceiver(v string) *ListUserNotifyHistoriesRequest {
	s.Receiver = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetRuleId(v string) *ListUserNotifyHistoriesRequest {
	s.RuleId = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetRuleName(v string) *ListUserNotifyHistoriesRequest {
	s.RuleName = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetSendType(v string) *ListUserNotifyHistoriesRequest {
	s.SendType = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetSeverity(v string) *ListUserNotifyHistoriesRequest {
	s.Severity = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetSubscriptionId(v string) *ListUserNotifyHistoriesRequest {
	s.SubscriptionId = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetSubscriptionName(v string) *ListUserNotifyHistoriesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetSuccess(v bool) *ListUserNotifyHistoriesRequest {
	s.Success = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetTo(v int32) *ListUserNotifyHistoriesRequest {
	s.To = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) SetWorkspace(v string) *ListUserNotifyHistoriesRequest {
	s.Workspace = &v
	return s
}

func (s *ListUserNotifyHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
