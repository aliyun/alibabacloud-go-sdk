// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityAlertRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCondition(v string) *CreateDataQualityAlertRuleShrinkRequest
	GetCondition() *string
	SetNotificationShrink(v string) *CreateDataQualityAlertRuleShrinkRequest
	GetNotificationShrink() *string
	SetProjectId(v int64) *CreateDataQualityAlertRuleShrinkRequest
	GetProjectId() *int64
	SetTargetShrink(v string) *CreateDataQualityAlertRuleShrinkRequest
	GetTargetShrink() *string
}

type CreateDataQualityAlertRuleShrinkRequest struct {
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
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	TargetShrink *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s CreateDataQualityAlertRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleShrinkRequest) GetCondition() *string {
	return s.Condition
}

func (s *CreateDataQualityAlertRuleShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateDataQualityAlertRuleShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityAlertRuleShrinkRequest) GetTargetShrink() *string {
	return s.TargetShrink
}

func (s *CreateDataQualityAlertRuleShrinkRequest) SetCondition(v string) *CreateDataQualityAlertRuleShrinkRequest {
	s.Condition = &v
	return s
}

func (s *CreateDataQualityAlertRuleShrinkRequest) SetNotificationShrink(v string) *CreateDataQualityAlertRuleShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateDataQualityAlertRuleShrinkRequest) SetProjectId(v int64) *CreateDataQualityAlertRuleShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityAlertRuleShrinkRequest) SetTargetShrink(v string) *CreateDataQualityAlertRuleShrinkRequest {
	s.TargetShrink = &v
	return s
}

func (s *CreateDataQualityAlertRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
