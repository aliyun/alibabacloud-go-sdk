// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotations(v string) *CreateAlertTemplateRequest
	GetAnnotations() *string
	SetLabels(v string) *CreateAlertTemplateRequest
	GetLabels() *string
	SetMatchExpressions(v string) *CreateAlertTemplateRequest
	GetMatchExpressions() *string
	SetMessage(v string) *CreateAlertTemplateRequest
	GetMessage() *string
	SetName(v string) *CreateAlertTemplateRequest
	GetName() *string
	SetParentId(v string) *CreateAlertTemplateRequest
	GetParentId() *string
	SetRegionId(v string) *CreateAlertTemplateRequest
	GetRegionId() *string
	SetRule(v string) *CreateAlertTemplateRequest
	GetRule() *string
	SetTemplateProvider(v string) *CreateAlertTemplateRequest
	GetTemplateProvider() *string
	SetType(v string) *CreateAlertTemplateRequest
	GetType() *string
}

type CreateAlertTemplateRequest struct {
	Annotations      *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Labels           *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	// This parameter is required.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Rule             *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	TemplateProvider *string `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAlertTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertTemplateRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *CreateAlertTemplateRequest) GetLabels() *string {
	return s.Labels
}

func (s *CreateAlertTemplateRequest) GetMatchExpressions() *string {
	return s.MatchExpressions
}

func (s *CreateAlertTemplateRequest) GetMessage() *string {
	return s.Message
}

func (s *CreateAlertTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertTemplateRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateAlertTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAlertTemplateRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateAlertTemplateRequest) GetTemplateProvider() *string {
	return s.TemplateProvider
}

func (s *CreateAlertTemplateRequest) GetType() *string {
	return s.Type
}

func (s *CreateAlertTemplateRequest) SetAnnotations(v string) *CreateAlertTemplateRequest {
	s.Annotations = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetLabels(v string) *CreateAlertTemplateRequest {
	s.Labels = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetMatchExpressions(v string) *CreateAlertTemplateRequest {
	s.MatchExpressions = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetMessage(v string) *CreateAlertTemplateRequest {
	s.Message = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetName(v string) *CreateAlertTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetParentId(v string) *CreateAlertTemplateRequest {
	s.ParentId = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetRegionId(v string) *CreateAlertTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetRule(v string) *CreateAlertTemplateRequest {
	s.Rule = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetTemplateProvider(v string) *CreateAlertTemplateRequest {
	s.TemplateProvider = &v
	return s
}

func (s *CreateAlertTemplateRequest) SetType(v string) *CreateAlertTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateAlertTemplateRequest) Validate() error {
	return dara.Validate(s)
}
