// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheusAlertTemplates(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) *ListPrometheusAlertTemplatesResponseBody
	GetPrometheusAlertTemplates() []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates
	SetRequestId(v string) *ListPrometheusAlertTemplatesResponseBody
	GetRequestId() *string
}

type ListPrometheusAlertTemplatesResponseBody struct {
	PrometheusAlertTemplates []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates `json:"PrometheusAlertTemplates,omitempty" xml:"PrometheusAlertTemplates,omitempty" type:"Repeated"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBody) GetPrometheusAlertTemplates() []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	return s.PrometheusAlertTemplates
}

func (s *ListPrometheusAlertTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusAlertTemplatesResponseBody) SetPrometheusAlertTemplates(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) *ListPrometheusAlertTemplatesResponseBody {
	s.PrometheusAlertTemplates = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBody) SetRequestId(v string) *ListPrometheusAlertTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBody) Validate() error {
	if s.PrometheusAlertTemplates != nil {
		for _, item := range s.PrometheusAlertTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates struct {
	AlertName   *string                                                                        `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	Description *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *string                                                                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression  *string                                                                        `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels      []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Type        *string                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Version     *string                                                                        `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetAlertName() *string {
	return s.AlertName
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetAnnotations() []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	return s.Annotations
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetDuration() *string {
	return s.Duration
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetExpression() *string {
	return s.Expression
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetLabels() []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	return s.Labels
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetType() *string {
	return s.Type
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GetVersion() *string {
	return s.Version
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetAlertName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetAnnotations(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetDescription(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Description = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetDuration(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetExpression(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetLabels(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetType(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetVersion(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Version = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) SetName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) SetValue(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) SetName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) SetValue(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) Validate() error {
	return dara.Validate(s)
}
