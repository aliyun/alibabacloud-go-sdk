// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *UpdateDataQualityAlertRuleRequest
	GetCondition() *string
	SetId(v int64) *UpdateDataQualityAlertRuleRequest
	GetId() *int64
	SetNotification(v *UpdateDataQualityAlertRuleRequestNotification) *UpdateDataQualityAlertRuleRequest
	GetNotification() *UpdateDataQualityAlertRuleRequestNotification
	SetProjectId(v int64) *UpdateDataQualityAlertRuleRequest
	GetProjectId() *int64
	SetTarget(v *UpdateDataQualityAlertRuleRequestTarget) *UpdateDataQualityAlertRuleRequest
	GetTarget() *UpdateDataQualityAlertRuleRequestTarget
}

type UpdateDataQualityAlertRuleRequest struct {
	// example:
	//
	// results.any { r -> r.status == \\"fail\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// 105412
	Id           *int64                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Notification *UpdateDataQualityAlertRuleRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// example:
	//
	// 1000
	ProjectId *int64                                   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Target    *UpdateDataQualityAlertRuleRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
}

func (s UpdateDataQualityAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleRequest) GetCondition() *string {
	return s.Condition
}

func (s *UpdateDataQualityAlertRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityAlertRuleRequest) GetNotification() *UpdateDataQualityAlertRuleRequestNotification {
	return s.Notification
}

func (s *UpdateDataQualityAlertRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityAlertRuleRequest) GetTarget() *UpdateDataQualityAlertRuleRequestTarget {
	return s.Target
}

func (s *UpdateDataQualityAlertRuleRequest) SetCondition(v string) *UpdateDataQualityAlertRuleRequest {
	s.Condition = &v
	return s
}

func (s *UpdateDataQualityAlertRuleRequest) SetId(v int64) *UpdateDataQualityAlertRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityAlertRuleRequest) SetNotification(v *UpdateDataQualityAlertRuleRequestNotification) *UpdateDataQualityAlertRuleRequest {
	s.Notification = v
	return s
}

func (s *UpdateDataQualityAlertRuleRequest) SetProjectId(v int64) *UpdateDataQualityAlertRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityAlertRuleRequest) SetTarget(v *UpdateDataQualityAlertRuleRequestTarget) *UpdateDataQualityAlertRuleRequest {
	s.Target = v
	return s
}

func (s *UpdateDataQualityAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityAlertRuleRequestNotification struct {
	// This parameter is required.
	Channels  []*string                                                 `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	Receivers []*UpdateDataQualityAlertRuleRequestNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
}

func (s UpdateDataQualityAlertRuleRequestNotification) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleRequestNotification) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleRequestNotification) GetChannels() []*string {
	return s.Channels
}

func (s *UpdateDataQualityAlertRuleRequestNotification) GetReceivers() []*UpdateDataQualityAlertRuleRequestNotificationReceivers {
	return s.Receivers
}

func (s *UpdateDataQualityAlertRuleRequestNotification) SetChannels(v []*string) *UpdateDataQualityAlertRuleRequestNotification {
	s.Channels = v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestNotification) SetReceivers(v []*UpdateDataQualityAlertRuleRequestNotificationReceivers) *UpdateDataQualityAlertRuleRequestNotification {
	s.Receivers = v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestNotification) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityAlertRuleRequestNotificationReceivers struct {
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TaskOwner
	ReceiverType   *string   `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s UpdateDataQualityAlertRuleRequestNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleRequestNotificationReceivers) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) SetExtension(v string) *UpdateDataQualityAlertRuleRequestNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) SetReceiverType(v string) *UpdateDataQualityAlertRuleRequestNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) SetReceiverValues(v []*string) *UpdateDataQualityAlertRuleRequestNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityAlertRuleRequestTarget struct {
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// example:
	//
	// DataQualityScan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityAlertRuleRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleRequestTarget) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleRequestTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *UpdateDataQualityAlertRuleRequestTarget) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityAlertRuleRequestTarget) SetIds(v []*int64) *UpdateDataQualityAlertRuleRequestTarget {
	s.Ids = v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestTarget) SetType(v string) *UpdateDataQualityAlertRuleRequestTarget {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityAlertRuleRequestTarget) Validate() error {
	return dara.Validate(s)
}
