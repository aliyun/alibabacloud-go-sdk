// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityAlertRule(v *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) *GetDataQualityAlertRuleResponseBody
	GetDataQualityAlertRule() *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule
	SetRequestId(v string) *GetDataQualityAlertRuleResponseBody
	GetRequestId() *string
}

type GetDataQualityAlertRuleResponseBody struct {
	// Data quality monitoring alert rules.
	DataQualityAlertRule *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule `json:"DataQualityAlertRule,omitempty" xml:"DataQualityAlertRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115****159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleResponseBody) GetDataQualityAlertRule() *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule {
	return s.DataQualityAlertRule
}

func (s *GetDataQualityAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityAlertRuleResponseBody) SetDataQualityAlertRule(v *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) *GetDataQualityAlertRuleResponseBody {
	s.DataQualityAlertRule = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBody) SetRequestId(v string) *GetDataQualityAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAlertRuleResponseBodyDataQualityAlertRule struct {
	// Alert conditions of the data quality monitoring alert rule.
	//
	// example:
	//
	// results.any { r -> r.status == \\"fail\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The ID of the data quality monitoring alert rule.
	//
	// example:
	//
	// 21045
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Alert notification settings.
	Notification *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The project ID.
	//
	// example:
	//
	// 90912
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The monitored target of the data quality monitoring alert rule.
	Target *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) GetNotification() *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification {
	return s.Notification
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) GetTarget() *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget {
	return s.Target
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) SetCondition(v string) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule {
	s.Condition = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) SetId(v int64) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule {
	s.Id = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) SetNotification(v *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule {
	s.Notification = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) SetProjectId(v int64) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) SetTarget(v *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule {
	s.Target = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRule) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification struct {
	// In Channels, you can set both Email and Sms at the same time. In other cases, only one channel can be set.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The alert recipients.
	Receivers []*GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) GetChannels() []*string {
	return s.Channels
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) GetReceivers() []*GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers {
	return s.Receivers
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) SetChannels(v []*string) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification {
	s.Channels = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) SetReceivers(v []*GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification {
	s.Receivers = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotification) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers struct {
	// Additional configurations required for the alert recipients. When ReceiverType is DingdingUrl, you can set `{"atAll":true}` to mention all members.
	//
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of alert recipients.
	//
	// 	- ShiftSchedule
	//
	// 	- WebhookUrl
	//
	// 	- FeishuUrl
	//
	// 	- TaskOwner
	//
	// 	- WeixinUrl
	//
	// 	- DingdingUrl
	//
	// 	- DataQualityScanOwner
	//
	// 	- AliUid
	//
	// example:
	//
	// TaskOwner
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The value of alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) SetExtension(v string) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) SetReceiverType(v string) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) SetReceiverValues(v []*string) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget struct {
	// The list of monitored target IDs.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The type of the monitored target. Only `DataQualityScan` is supported.
	//
	// example:
	//
	// DataQualityScan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) GoString() string {
	return s.String()
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) GetType() *string {
	return s.Type
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) SetIds(v []*int64) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget {
	s.Ids = v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) SetType(v string) *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget {
	s.Type = &v
	return s
}

func (s *GetDataQualityAlertRuleResponseBodyDataQualityAlertRuleTarget) Validate() error {
	return dara.Validate(s)
}
