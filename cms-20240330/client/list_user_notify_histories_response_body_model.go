// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserNotifyHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHistoryList(v []*ListUserNotifyHistoriesResponseBodyHistoryList) *ListUserNotifyHistoriesResponseBody
	GetHistoryList() []*ListUserNotifyHistoriesResponseBodyHistoryList
	SetMaxResults(v int32) *ListUserNotifyHistoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserNotifyHistoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserNotifyHistoriesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListUserNotifyHistoriesResponseBody
	GetTotal() *int64
}

type ListUserNotifyHistoriesResponseBody struct {
	HistoryList []*ListUserNotifyHistoriesResponseBodyHistoryList `json:"historyList,omitempty" xml:"historyList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 2f8aKz3mNq9p****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 128
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListUserNotifyHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserNotifyHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserNotifyHistoriesResponseBody) GetHistoryList() []*ListUserNotifyHistoriesResponseBodyHistoryList {
	return s.HistoryList
}

func (s *ListUserNotifyHistoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserNotifyHistoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserNotifyHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserNotifyHistoriesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListUserNotifyHistoriesResponseBody) SetHistoryList(v []*ListUserNotifyHistoriesResponseBodyHistoryList) *ListUserNotifyHistoriesResponseBody {
	s.HistoryList = v
	return s
}

func (s *ListUserNotifyHistoriesResponseBody) SetMaxResults(v int32) *ListUserNotifyHistoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBody) SetNextToken(v string) *ListUserNotifyHistoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBody) SetRequestId(v string) *ListUserNotifyHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBody) SetTotal(v int64) *ListUserNotifyHistoriesResponseBody {
	s.Total = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBody) Validate() error {
	if s.HistoryList != nil {
		for _, item := range s.HistoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserNotifyHistoriesResponseBodyHistoryList struct {
	// example:
	//
	// DING
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// 张三,李四
	Contacts *string `json:"contacts,omitempty" xml:"contacts,omitempty"`
	// example:
	//
	// 运维组,值班组
	Groups *string `json:"groups,omitempty" xml:"groups,omitempty"`
	// example:
	//
	// 5d41402abc4b2a76b9719d911017c592
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// example:
	//
	// 3fa85f64-5717-4562-b3fc-2c963f66afa8
	JobId     *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	NoChannel *bool   `json:"noChannel,omitempty" xml:"noChannel,omitempty"`
	// example:
	//
	// 2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// example:
	//
	// 默认通知策略
	NotifyStrategyName *string `json:"notifyStrategyName,omitempty" xml:"notifyStrategyName,omitempty"`
	// example:
	//
	// 13800138000
	Receiver *string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// example:
	//
	// {"errcode":0,"errmsg":"ok"}
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// 3fa85f64-5717-4562-b3fc-2c963f66afa6
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// example:
	//
	// CPU使用率告警
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// example:
	//
	// SUBSCRIBE
	SendType *string `json:"sendType,omitempty" xml:"sendType,omitempty"`
	// example:
	//
	// CRITICAL
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// example:
	//
	// TRIGGERED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d
	SubscriptionId *string `json:"subscriptionId,omitempty" xml:"subscriptionId,omitempty"`
	// example:
	//
	// 默认订阅
	SubscriptionName *string `json:"subscriptionName,omitempty" xml:"subscriptionName,omitempty"`
	Success          *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 1729400000
	Timestamp *int32 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 5884500827389836
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListUserNotifyHistoriesResponseBodyHistoryList) String() string {
	return dara.Prettify(s)
}

func (s ListUserNotifyHistoriesResponseBodyHistoryList) GoString() string {
	return s.String()
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetChannel() *string {
	return s.Channel
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetContacts() *string {
	return s.Contacts
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetGroups() *string {
	return s.Groups
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetIncidentId() *string {
	return s.IncidentId
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetJobId() *string {
	return s.JobId
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetNoChannel() *bool {
	return s.NoChannel
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetNotifyStrategyName() *string {
	return s.NotifyStrategyName
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetReceiver() *string {
	return s.Receiver
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetResult() *string {
	return s.Result
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetRuleId() *string {
	return s.RuleId
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetSendType() *string {
	return s.SendType
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetSeverity() *string {
	return s.Severity
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetStatus() *string {
	return s.Status
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetSubscriptionId() *string {
	return s.SubscriptionId
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetSubscriptionName() *string {
	return s.SubscriptionName
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) GetUserId() *string {
	return s.UserId
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetChannel(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Channel = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetContacts(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Contacts = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetGroups(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Groups = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetIncidentId(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.IncidentId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetJobId(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.JobId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetNoChannel(v bool) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.NoChannel = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetNotifyStrategyId(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.NotifyStrategyId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetNotifyStrategyName(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.NotifyStrategyName = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetReceiver(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Receiver = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetResult(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Result = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetRuleId(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.RuleId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetRuleName(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.RuleName = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetSendType(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.SendType = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetSeverity(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Severity = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetStatus(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Status = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetSubscriptionId(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.SubscriptionId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetSubscriptionName(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.SubscriptionName = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetSuccess(v bool) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Success = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetTimestamp(v int32) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.Timestamp = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) SetUserId(v string) *ListUserNotifyHistoriesResponseBodyHistoryList {
	s.UserId = &v
	return s
}

func (s *ListUserNotifyHistoriesResponseBodyHistoryList) Validate() error {
	return dara.Validate(s)
}
