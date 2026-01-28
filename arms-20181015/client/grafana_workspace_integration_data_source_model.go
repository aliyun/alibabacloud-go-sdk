// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceIntegrationDataSource interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *GrafanaWorkspaceIntegrationDataSource
	GetClusterType() *string
	SetDatasourceId(v string) *GrafanaWorkspaceIntegrationDataSource
	GetDatasourceId() *string
	SetDatasourceName(v string) *GrafanaWorkspaceIntegrationDataSource
	GetDatasourceName() *string
	SetDatasourceUrl(v string) *GrafanaWorkspaceIntegrationDataSource
	GetDatasourceUrl() *string
	SetDescription(v string) *GrafanaWorkspaceIntegrationDataSource
	GetDescription() *string
	SetExploreUrl(v string) *GrafanaWorkspaceIntegrationDataSource
	GetExploreUrl() *string
	SetFolderUrl(v string) *GrafanaWorkspaceIntegrationDataSource
	GetFolderUrl() *string
	SetRegionId(v string) *GrafanaWorkspaceIntegrationDataSource
	GetRegionId() *string
	SetStatus(v string) *GrafanaWorkspaceIntegrationDataSource
	GetStatus() *string
	SetType(v string) *GrafanaWorkspaceIntegrationDataSource
	GetType() *string
}

type GrafanaWorkspaceIntegrationDataSource struct {
	ClusterType    *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	DatasourceId   *string `json:"datasourceId,omitempty" xml:"datasourceId,omitempty"`
	DatasourceName *string `json:"datasourceName,omitempty" xml:"datasourceName,omitempty"`
	DatasourceUrl  *string `json:"datasourceUrl,omitempty" xml:"datasourceUrl,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	ExploreUrl     *string `json:"exploreUrl,omitempty" xml:"exploreUrl,omitempty"`
	FolderUrl      *string `json:"folderUrl,omitempty" xml:"folderUrl,omitempty"`
	RegionId       *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GrafanaWorkspaceIntegrationDataSource) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceIntegrationDataSource) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetClusterType() *string {
	return s.ClusterType
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetDatasourceUrl() *string {
	return s.DatasourceUrl
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetDescription() *string {
	return s.Description
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetExploreUrl() *string {
	return s.ExploreUrl
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetFolderUrl() *string {
	return s.FolderUrl
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetRegionId() *string {
	return s.RegionId
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetStatus() *string {
	return s.Status
}

func (s *GrafanaWorkspaceIntegrationDataSource) GetType() *string {
	return s.Type
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetClusterType(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.ClusterType = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetDatasourceId(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.DatasourceId = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetDatasourceName(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.DatasourceName = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetDatasourceUrl(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.DatasourceUrl = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetDescription(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.Description = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetExploreUrl(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.ExploreUrl = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetFolderUrl(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.FolderUrl = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetRegionId(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.RegionId = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetStatus(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.Status = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) SetType(v string) *GrafanaWorkspaceIntegrationDataSource {
	s.Type = &v
	return s
}

func (s *GrafanaWorkspaceIntegrationDataSource) Validate() error {
	return dara.Validate(s)
}
