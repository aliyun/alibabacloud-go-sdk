// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *CreatePrometheusAlertRuleRequest
	GetAlertName() *string
	SetAnnotations(v string) *CreatePrometheusAlertRuleRequest
	GetAnnotations() *string
	SetClusterId(v string) *CreatePrometheusAlertRuleRequest
	GetClusterId() *string
	SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleRequest
	GetDispatchRuleId() *int64
	SetDuration(v string) *CreatePrometheusAlertRuleRequest
	GetDuration() *string
	SetExpression(v string) *CreatePrometheusAlertRuleRequest
	GetExpression() *string
	SetLabels(v string) *CreatePrometheusAlertRuleRequest
	GetLabels() *string
	SetMessage(v string) *CreatePrometheusAlertRuleRequest
	GetMessage() *string
	SetNotifyType(v string) *CreatePrometheusAlertRuleRequest
	GetNotifyType() *string
	SetRegionId(v string) *CreatePrometheusAlertRuleRequest
	GetRegionId() *string
	SetType(v string) *CreatePrometheusAlertRuleRequest
	GetType() *string
}

type CreatePrometheusAlertRuleRequest struct {
	// This parameter is required.
	AlertName   *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// This parameter is required.
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// This parameter is required.
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// This parameter is required.
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *CreatePrometheusAlertRuleRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *CreatePrometheusAlertRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePrometheusAlertRuleRequest) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *CreatePrometheusAlertRuleRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreatePrometheusAlertRuleRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreatePrometheusAlertRuleRequest) GetLabels() *string {
	return s.Labels
}

func (s *CreatePrometheusAlertRuleRequest) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusAlertRuleRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *CreatePrometheusAlertRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrometheusAlertRuleRequest) GetType() *string {
	return s.Type
}

func (s *CreatePrometheusAlertRuleRequest) SetAlertName(v string) *CreatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetAnnotations(v string) *CreatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetClusterId(v string) *CreatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDuration(v string) *CreatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetExpression(v string) *CreatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetLabels(v string) *CreatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetMessage(v string) *CreatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetNotifyType(v string) *CreatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetRegionId(v string) *CreatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetType(v string) *CreatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
