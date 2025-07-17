// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIntegrationDetail interface {
	dara.Model
	String() string
	GoString() string
	SetDataSources(v []*GrafanaWorkspaceIntegrationDataSource) *GrafanaWorkspaceIntegrationDetail
	GetDataSources() []*GrafanaWorkspaceIntegrationDataSource
	SetIntegrationId(v string) *GrafanaWorkspaceIntegrationDetail
	GetIntegrationId() *string
	SetStatus(v string) *GrafanaWorkspaceIntegrationDetail
	GetStatus() *string
}

type GrafanaWorkspaceIntegrationDetail struct {
	DataSources   []*GrafanaWorkspaceIntegrationDataSource `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	IntegrationId *string                                  `json:"integrationId,omitempty" xml:"integrationId,omitempty"`
	Status        *string                                  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GrafanaWorkspaceIntegrationDetail) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIntegrationDetail) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIntegrationDetail) GetDataSources() []*GrafanaWorkspaceIntegrationDataSource {
	return s.DataSources
}

func (s *GrafanaWorkspaceIntegrationDetail) GetIntegrationId() *string {
	return s.IntegrationId
}

func (s *GrafanaWorkspaceIntegrationDetail) GetStatus() *string {
	return s.Status
}

func (s *GrafanaWorkspaceIntegrationDetail) SetDataSources(v []*GrafanaWorkspaceIntegrationDataSource) *GrafanaWorkspaceIntegrationDetail {
	s.DataSources = v
	return s
}

func (s *GrafanaWorkspaceIntegrationDetail) SetIntegrationId(v string) *GrafanaWorkspaceIntegrationDetail {
	s.IntegrationId = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDetail) SetStatus(v string) *GrafanaWorkspaceIntegrationDetail {
	s.Status = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDetail) Validate() error {
	return dara.Validate(s)
}
