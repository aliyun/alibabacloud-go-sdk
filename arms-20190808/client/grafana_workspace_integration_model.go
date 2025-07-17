// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIntegration interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceAmount(v int64) *GrafanaWorkspaceIntegration
	GetDatasourceAmount() *int64
	SetIntegrationId(v string) *GrafanaWorkspaceIntegration
	GetIntegrationId() *string
	SetIntegrationName(v string) *GrafanaWorkspaceIntegration
	GetIntegrationName() *string
	SetPreviews(v []*GrafanaWorkspaceIntegrationPreview) *GrafanaWorkspaceIntegration
	GetPreviews() []*GrafanaWorkspaceIntegrationPreview
	SetStatus(v string) *GrafanaWorkspaceIntegration
	GetStatus() *string
	SetSupportRegions(v []*string) *GrafanaWorkspaceIntegration
	GetSupportRegions() []*string
}

type GrafanaWorkspaceIntegration struct {
	DatasourceAmount *int64                                `json:"datasourceAmount,omitempty" xml:"datasourceAmount,omitempty"`
	IntegrationId    *string                               `json:"integrationId,omitempty" xml:"integrationId,omitempty"`
	IntegrationName  *string                               `json:"integrationName,omitempty" xml:"integrationName,omitempty"`
	Previews         []*GrafanaWorkspaceIntegrationPreview `json:"previews,omitempty" xml:"previews,omitempty" type:"Repeated"`
	Status           *string                               `json:"status,omitempty" xml:"status,omitempty"`
	SupportRegions   []*string                             `json:"supportRegions,omitempty" xml:"supportRegions,omitempty" type:"Repeated"`
}

func (s GrafanaWorkspaceIntegration) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIntegration) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIntegration) GetDatasourceAmount() *int64 {
	return s.DatasourceAmount
}

func (s *GrafanaWorkspaceIntegration) GetIntegrationId() *string {
	return s.IntegrationId
}

func (s *GrafanaWorkspaceIntegration) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *GrafanaWorkspaceIntegration) GetPreviews() []*GrafanaWorkspaceIntegrationPreview {
	return s.Previews
}

func (s *GrafanaWorkspaceIntegration) GetStatus() *string {
	return s.Status
}

func (s *GrafanaWorkspaceIntegration) GetSupportRegions() []*string {
	return s.SupportRegions
}

func (s *GrafanaWorkspaceIntegration) SetDatasourceAmount(v int64) *GrafanaWorkspaceIntegration {
	s.DatasourceAmount = &v
	return s
}

func (s *GrafanaWorkspaceIntegration) SetIntegrationId(v string) *GrafanaWorkspaceIntegration {
	s.IntegrationId = &v
	return s
}

func (s *GrafanaWorkspaceIntegration) SetIntegrationName(v string) *GrafanaWorkspaceIntegration {
	s.IntegrationName = &v
	return s
}

func (s *GrafanaWorkspaceIntegration) SetPreviews(v []*GrafanaWorkspaceIntegrationPreview) *GrafanaWorkspaceIntegration {
	s.Previews = v
	return s
}

func (s *GrafanaWorkspaceIntegration) SetStatus(v string) *GrafanaWorkspaceIntegration {
	s.Status = &v
	return s
}

func (s *GrafanaWorkspaceIntegration) SetSupportRegions(v []*string) *GrafanaWorkspaceIntegration {
	s.SupportRegions = v
	return s
}

func (s *GrafanaWorkspaceIntegration) Validate() error {
	return dara.Validate(s)
}
