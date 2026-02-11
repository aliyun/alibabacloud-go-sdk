// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *UpdateAlertTemplateRequest
	GetAnnotations() *string
	SetId(v int64) *UpdateAlertTemplateRequest
	GetId() *int64
	SetLabels(v string) *UpdateAlertTemplateRequest
	GetLabels() *string
	SetMatchExpressions(v string) *UpdateAlertTemplateRequest
	GetMatchExpressions() *string
	SetMessage(v string) *UpdateAlertTemplateRequest
	GetMessage() *string
	SetName(v string) *UpdateAlertTemplateRequest
	GetName() *string
	SetRegionId(v string) *UpdateAlertTemplateRequest
	GetRegionId() *string
	SetRule(v string) *UpdateAlertTemplateRequest
	GetRule() *string
	SetStatus(v bool) *UpdateAlertTemplateRequest
	GetStatus() *bool
	SetType(v string) *UpdateAlertTemplateRequest
	GetType() *string
}

type UpdateAlertTemplateRequest struct {
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// This parameter is required.
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Labels           *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	// This parameter is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// This parameter is required.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAlertTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertTemplateRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *UpdateAlertTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAlertTemplateRequest) GetLabels() *string {
	return s.Labels
}

func (s *UpdateAlertTemplateRequest) GetMatchExpressions() *string {
	return s.MatchExpressions
}

func (s *UpdateAlertTemplateRequest) GetMessage() *string {
	return s.Message
}

func (s *UpdateAlertTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertTemplateRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateAlertTemplateRequest) GetStatus() *bool {
	return s.Status
}

func (s *UpdateAlertTemplateRequest) GetType() *string {
	return s.Type
}

func (s *UpdateAlertTemplateRequest) SetAnnotations(v string) *UpdateAlertTemplateRequest {
	s.Annotations = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetId(v int64) *UpdateAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetLabels(v string) *UpdateAlertTemplateRequest {
	s.Labels = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetMatchExpressions(v string) *UpdateAlertTemplateRequest {
	s.MatchExpressions = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetMessage(v string) *UpdateAlertTemplateRequest {
	s.Message = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetName(v string) *UpdateAlertTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetRegionId(v string) *UpdateAlertTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetRule(v string) *UpdateAlertTemplateRequest {
	s.Rule = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetStatus(v bool) *UpdateAlertTemplateRequest {
	s.Status = &v
	return s
}

func (s *UpdateAlertTemplateRequest) SetType(v string) *UpdateAlertTemplateRequest {
	s.Type = &v
	return s
}

func (s *UpdateAlertTemplateRequest) Validate() error {
	return dara.Validate(s)
}
