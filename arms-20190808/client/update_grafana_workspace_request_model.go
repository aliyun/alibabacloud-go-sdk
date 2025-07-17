// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateGrafanaWorkspaceRequest
	GetAliyunLang() *string
	SetDescription(v string) *UpdateGrafanaWorkspaceRequest
	GetDescription() *string
	SetGrafanaWorkspaceId(v string) *UpdateGrafanaWorkspaceRequest
	GetGrafanaWorkspaceId() *string
	SetGrafanaWorkspaceName(v string) *UpdateGrafanaWorkspaceRequest
	GetGrafanaWorkspaceName() *string
	SetRegionId(v string) *UpdateGrafanaWorkspaceRequest
	GetRegionId() *string
}

type UpdateGrafanaWorkspaceRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The description of the workspace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-4xl3g******
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// testGrafana
	GrafanaWorkspaceName *string `json:"GrafanaWorkspaceName,omitempty" xml:"GrafanaWorkspaceName,omitempty"`
	// The region ID. Default value: `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateGrafanaWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateGrafanaWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGrafanaWorkspaceRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *UpdateGrafanaWorkspaceRequest) GetGrafanaWorkspaceName() *string {
	return s.GrafanaWorkspaceName
}

func (s *UpdateGrafanaWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateGrafanaWorkspaceRequest) SetAliyunLang(v string) *UpdateGrafanaWorkspaceRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateGrafanaWorkspaceRequest) SetDescription(v string) *UpdateGrafanaWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateGrafanaWorkspaceRequest) SetGrafanaWorkspaceId(v string) *UpdateGrafanaWorkspaceRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceRequest) SetGrafanaWorkspaceName(v string) *UpdateGrafanaWorkspaceRequest {
	s.GrafanaWorkspaceName = &v
	return s
}

func (s *UpdateGrafanaWorkspaceRequest) SetRegionId(v string) *UpdateGrafanaWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
