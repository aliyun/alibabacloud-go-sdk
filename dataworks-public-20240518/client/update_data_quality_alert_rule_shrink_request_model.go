// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityAlertRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *UpdateDataQualityAlertRuleShrinkRequest
	GetCondition() *string
	SetId(v int64) *UpdateDataQualityAlertRuleShrinkRequest
	GetId() *int64
	SetNotificationShrink(v string) *UpdateDataQualityAlertRuleShrinkRequest
	GetNotificationShrink() *string
	SetProjectId(v int64) *UpdateDataQualityAlertRuleShrinkRequest
	GetProjectId() *int64
	SetTargetShrink(v string) *UpdateDataQualityAlertRuleShrinkRequest
	GetTargetShrink() *string
}

type UpdateDataQualityAlertRuleShrinkRequest struct {
	// The alert condition of the data quality monitoring rule.
	//
	// example:
	//
	// results.any { r -> r.status == \\"fail\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Alert notification configurations.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The monitored target of the data quality monitoring rule.
	TargetShrink *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s UpdateDataQualityAlertRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) GetCondition() *string {
	return s.Condition
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) GetTargetShrink() *string {
	return s.TargetShrink
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) SetCondition(v string) *UpdateDataQualityAlertRuleShrinkRequest {
	s.Condition = &v
	return s
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) SetId(v int64) *UpdateDataQualityAlertRuleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) SetNotificationShrink(v string) *UpdateDataQualityAlertRuleShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) SetProjectId(v int64) *UpdateDataQualityAlertRuleShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) SetTargetShrink(v string) *UpdateDataQualityAlertRuleShrinkRequest {
	s.TargetShrink = &v
	return s
}

func (s *UpdateDataQualityAlertRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
