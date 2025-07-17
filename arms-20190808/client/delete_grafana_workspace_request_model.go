// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrafanaWorkspaceId(v string) *DeleteGrafanaWorkspaceRequest
	GetGrafanaWorkspaceId() *string
	SetRegionId(v string) *DeleteGrafanaWorkspaceRequest
	GetRegionId() *string
}

type DeleteGrafanaWorkspaceRequest struct {
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// grafana-rnglkcdrntlhk0****
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// The region ID. Default value: cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGrafanaWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaWorkspaceRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *DeleteGrafanaWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGrafanaWorkspaceRequest) SetGrafanaWorkspaceId(v string) *DeleteGrafanaWorkspaceRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceRequest) SetRegionId(v string) *DeleteGrafanaWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGrafanaWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
