// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGrafanaWorkspaceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateGrafanaWorkspaceVersionRequest
	GetAliyunLang() *string
	SetGrafanaVersion(v string) *UpdateGrafanaWorkspaceVersionRequest
	GetGrafanaVersion() *string
	SetGrafanaWorkspaceId(v string) *UpdateGrafanaWorkspaceVersionRequest
	GetGrafanaWorkspaceId() *string
	SetRegionId(v string) *UpdateGrafanaWorkspaceVersionRequest
	GetRegionId() *string
}

type UpdateGrafanaWorkspaceVersionRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The Grafana version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.0.x
	GrafanaVersion *string `json:"GrafanaVersion,omitempty" xml:"GrafanaVersion,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-4xl3g******
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

func (s UpdateGrafanaWorkspaceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGrafanaWorkspaceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGrafanaWorkspaceVersionRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateGrafanaWorkspaceVersionRequest) GetGrafanaVersion() *string {
	return s.GrafanaVersion
}

func (s *UpdateGrafanaWorkspaceVersionRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *UpdateGrafanaWorkspaceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateGrafanaWorkspaceVersionRequest) SetAliyunLang(v string) *UpdateGrafanaWorkspaceVersionRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionRequest) SetGrafanaVersion(v string) *UpdateGrafanaWorkspaceVersionRequest {
	s.GrafanaVersion = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionRequest) SetGrafanaWorkspaceId(v string) *UpdateGrafanaWorkspaceVersionRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionRequest) SetRegionId(v string) *UpdateGrafanaWorkspaceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateGrafanaWorkspaceVersionRequest) Validate() error {
	return dara.Validate(s)
}
