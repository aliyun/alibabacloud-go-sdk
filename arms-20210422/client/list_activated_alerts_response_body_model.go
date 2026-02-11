// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActivatedAlertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v *ListActivatedAlertsResponseBodyPage) *ListActivatedAlertsResponseBody
	GetPage() *ListActivatedAlertsResponseBodyPage
	SetRequestId(v string) *ListActivatedAlertsResponseBody
	GetRequestId() *string
}

type ListActivatedAlertsResponseBody struct {
	Page      *ListActivatedAlertsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListActivatedAlertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActivatedAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBody) GetPage() *ListActivatedAlertsResponseBodyPage {
	return s.Page
}

func (s *ListActivatedAlertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActivatedAlertsResponseBody) SetPage(v *ListActivatedAlertsResponseBodyPage) *ListActivatedAlertsResponseBody {
	s.Page = v
	return s
}

func (s *ListActivatedAlertsResponseBody) SetRequestId(v string) *ListActivatedAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActivatedAlertsResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListActivatedAlertsResponseBodyPage struct {
	Alerts   []*ListActivatedAlertsResponseBodyPageAlerts `json:"Alerts,omitempty" xml:"Alerts,omitempty" type:"Repeated"`
	Page     *int32                                       `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPage) GetAlerts() []*ListActivatedAlertsResponseBodyPageAlerts {
	return s.Alerts
}

func (s *ListActivatedAlertsResponseBodyPage) GetPage() *int32 {
	return s.Page
}

func (s *ListActivatedAlertsResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListActivatedAlertsResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListActivatedAlertsResponseBodyPage) SetAlerts(v []*ListActivatedAlertsResponseBodyPageAlerts) *ListActivatedAlertsResponseBodyPage {
	s.Alerts = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetPage(v int32) *ListActivatedAlertsResponseBodyPage {
	s.Page = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetPageSize(v int32) *ListActivatedAlertsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetTotal(v int32) *ListActivatedAlertsResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) Validate() error {
	if s.Alerts != nil {
		for _, item := range s.Alerts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActivatedAlertsResponseBodyPageAlerts struct {
	AlertId            *string                                                   `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName          *string                                                   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertType          *string                                                   `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Count              *int32                                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime         *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DispatchRules      []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	EndsAt             *int64                                                    `json:"EndsAt,omitempty" xml:"EndsAt,omitempty"`
	ExpandFields       map[string]interface{}                                    `json:"ExpandFields,omitempty" xml:"ExpandFields,omitempty"`
	IntegrationName    *string                                                   `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationType    *string                                                   `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	InvolvedObjectKind *string                                                   `json:"InvolvedObjectKind,omitempty" xml:"InvolvedObjectKind,omitempty"`
	InvolvedObjectName *string                                                   `json:"InvolvedObjectName,omitempty" xml:"InvolvedObjectName,omitempty"`
	Message            *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Severity           *string                                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartsAt           *int64                                                    `json:"StartsAt,omitempty" xml:"StartsAt,omitempty"`
	Status             *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPageAlerts) String() string {
	return dara.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPageAlerts) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetAlertId() *string {
	return s.AlertId
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetAlertName() *string {
	return s.AlertName
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetAlertType() *string {
	return s.AlertType
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetCount() *int32 {
	return s.Count
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetDispatchRules() []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	return s.DispatchRules
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetEndsAt() *int64 {
	return s.EndsAt
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetExpandFields() map[string]interface{} {
	return s.ExpandFields
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetInvolvedObjectKind() *string {
	return s.InvolvedObjectKind
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetInvolvedObjectName() *string {
	return s.InvolvedObjectName
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetMessage() *string {
	return s.Message
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetSeverity() *string {
	return s.Severity
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetStartsAt() *int64 {
	return s.StartsAt
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) GetStatus() *string {
	return s.Status
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertId(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertId = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertType(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertType = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetCount(v int32) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Count = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetCreateTime(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.CreateTime = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetDispatchRules(v []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules) *ListActivatedAlertsResponseBodyPageAlerts {
	s.DispatchRules = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetEndsAt(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.EndsAt = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetExpandFields(v map[string]interface{}) *ListActivatedAlertsResponseBodyPageAlerts {
	s.ExpandFields = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetIntegrationName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.IntegrationName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetIntegrationType(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.IntegrationType = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetInvolvedObjectKind(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.InvolvedObjectKind = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetInvolvedObjectName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.InvolvedObjectName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetMessage(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Message = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetSeverity(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Severity = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetStartsAt(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.StartsAt = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetStatus(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Status = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) Validate() error {
	if s.DispatchRules != nil {
		for _, item := range s.DispatchRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActivatedAlertsResponseBodyPageAlertsDispatchRules struct {
	RuleId   *int32  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPageAlertsDispatchRules) String() string {
	return dara.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPageAlertsDispatchRules) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) GetRuleId() *int32 {
	return s.RuleId
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) SetRuleId(v int32) *ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	s.RuleId = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) SetRuleName(v string) *ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	s.RuleName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) Validate() error {
	return dara.Validate(s)
}
