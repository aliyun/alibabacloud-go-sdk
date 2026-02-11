// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertProvider(v string) *ListAlertTemplatesRequest
	GetAlertProvider() *string
	SetLabels(v string) *ListAlertTemplatesRequest
	GetLabels() *string
	SetName(v string) *ListAlertTemplatesRequest
	GetName() *string
	SetRegionId(v string) *ListAlertTemplatesRequest
	GetRegionId() *string
	SetStatus(v bool) *ListAlertTemplatesRequest
	GetStatus() *bool
	SetTemplateProvider(v string) *ListAlertTemplatesRequest
	GetTemplateProvider() *string
	SetType(v string) *ListAlertTemplatesRequest
	GetType() *string
}

type ListAlertTemplatesRequest struct {
	AlertProvider *string `json:"AlertProvider,omitempty" xml:"AlertProvider,omitempty"`
	Labels        *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateProvider *string `json:"TemplateProvider,omitempty" xml:"TemplateProvider,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertTemplatesRequest) GetAlertProvider() *string {
	return s.AlertProvider
}

func (s *ListAlertTemplatesRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListAlertTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAlertTemplatesRequest) GetStatus() *bool {
	return s.Status
}

func (s *ListAlertTemplatesRequest) GetTemplateProvider() *string {
	return s.TemplateProvider
}

func (s *ListAlertTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListAlertTemplatesRequest) SetAlertProvider(v string) *ListAlertTemplatesRequest {
	s.AlertProvider = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetLabels(v string) *ListAlertTemplatesRequest {
	s.Labels = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetName(v string) *ListAlertTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetRegionId(v string) *ListAlertTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetStatus(v bool) *ListAlertTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetTemplateProvider(v string) *ListAlertTemplatesRequest {
	s.TemplateProvider = &v
	return s
}

func (s *ListAlertTemplatesRequest) SetType(v string) *ListAlertTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListAlertTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
