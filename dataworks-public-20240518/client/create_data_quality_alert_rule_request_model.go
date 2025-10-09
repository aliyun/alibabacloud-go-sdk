// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *CreateDataQualityAlertRuleRequest
	GetCondition() *string
	SetNotification(v *CreateDataQualityAlertRuleRequestNotification) *CreateDataQualityAlertRuleRequest
	GetNotification() *CreateDataQualityAlertRuleRequestNotification
	SetProjectId(v int64) *CreateDataQualityAlertRuleRequest
	GetProjectId() *int64
	SetTarget(v *CreateDataQualityAlertRuleRequestTarget) *CreateDataQualityAlertRuleRequest
	GetTarget() *CreateDataQualityAlertRuleRequestTarget
}

type CreateDataQualityAlertRuleRequest struct {
	// The alert condition of the data quality monitoring rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// results.any { r -> r.status == \\"fail\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The list of alert channels.
	//
	// This parameter is required.
	Notification *CreateDataQualityAlertRuleRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The monitored target of the data quality monitoring rule.
	//
	// This parameter is required.
	Target *CreateDataQualityAlertRuleRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s CreateDataQualityAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleRequest) GetCondition() *string {
	return s.Condition
}

func (s *CreateDataQualityAlertRuleRequest) GetNotification() *CreateDataQualityAlertRuleRequestNotification {
	return s.Notification
}

func (s *CreateDataQualityAlertRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityAlertRuleRequest) GetTarget() *CreateDataQualityAlertRuleRequestTarget {
	return s.Target
}

func (s *CreateDataQualityAlertRuleRequest) SetCondition(v string) *CreateDataQualityAlertRuleRequest {
	s.Condition = &v
	return s
}

func (s *CreateDataQualityAlertRuleRequest) SetNotification(v *CreateDataQualityAlertRuleRequestNotification) *CreateDataQualityAlertRuleRequest {
	s.Notification = v
	return s
}

func (s *CreateDataQualityAlertRuleRequest) SetProjectId(v int64) *CreateDataQualityAlertRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityAlertRuleRequest) SetTarget(v *CreateDataQualityAlertRuleRequestTarget) *CreateDataQualityAlertRuleRequest {
	s.Target = v
	return s
}

func (s *CreateDataQualityAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityAlertRuleRequestNotification struct {
	// The list of alert channels. You can set both `Email` and `Sms` at the same time. In other cases, only one channel can be set.
	//
	// This parameter is required.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The alert recipients.
	//
	// This parameter is required.
	Receivers []*CreateDataQualityAlertRuleRequestNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
}

func (s CreateDataQualityAlertRuleRequestNotification) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleRequestNotification) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleRequestNotification) GetChannels() []*string {
	return s.Channels
}

func (s *CreateDataQualityAlertRuleRequestNotification) GetReceivers() []*CreateDataQualityAlertRuleRequestNotificationReceivers {
	return s.Receivers
}

func (s *CreateDataQualityAlertRuleRequestNotification) SetChannels(v []*string) *CreateDataQualityAlertRuleRequestNotification {
	s.Channels = v
	return s
}

func (s *CreateDataQualityAlertRuleRequestNotification) SetReceivers(v []*CreateDataQualityAlertRuleRequestNotificationReceivers) *CreateDataQualityAlertRuleRequestNotification {
	s.Receivers = v
	return s
}

func (s *CreateDataQualityAlertRuleRequestNotification) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityAlertRuleRequestNotificationReceivers struct {
	// Additional configurations required for the alert recipients. When ReceiverType is DingdingUrl, you can set `{"atAll":true}` to mention all members.
	//
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of alert recipients.
	//
	// 	- AliUid
	//
	// 	- WebhookUrl
	//
	// 	- DingdingUrl
	//
	// 	- WeixinUrl
	//
	// 	- FeishuUrl
	//
	// 	- TaskOwner
	//
	// 	- DataQualityScanOwner
	//
	// 	- ShiftSchedule
	//
	// This parameter is required.
	//
	// example:
	//
	// TaskOwner
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The value of alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s CreateDataQualityAlertRuleRequestNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleRequestNotificationReceivers) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) SetExtension(v string) *CreateDataQualityAlertRuleRequestNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) SetReceiverType(v string) *CreateDataQualityAlertRuleRequestNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) SetReceiverValues(v []*string) *CreateDataQualityAlertRuleRequestNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *CreateDataQualityAlertRuleRequestNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityAlertRuleRequestTarget struct {
	// The list of monitored target IDs. Currently, only one ID can be set.
	//
	// This parameter is required.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The type of the monitored target. Only DataQualityScan is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// DataQualityScan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityAlertRuleRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleRequestTarget) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleRequestTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *CreateDataQualityAlertRuleRequestTarget) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityAlertRuleRequestTarget) SetIds(v []*int64) *CreateDataQualityAlertRuleRequestTarget {
	s.Ids = v
	return s
}

func (s *CreateDataQualityAlertRuleRequestTarget) SetType(v string) *CreateDataQualityAlertRuleRequestTarget {
	s.Type = &v
	return s
}

func (s *CreateDataQualityAlertRuleRequestTarget) Validate() error {
	return dara.Validate(s)
}
