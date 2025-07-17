// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGrafanaWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetGrafanaWorkspaceRequest
	GetAliyunLang() *string
	SetGrafanaWorkspaceId(v string) *GetGrafanaWorkspaceRequest
	GetGrafanaWorkspaceId() *string
	SetRegionId(v string) *GetGrafanaWorkspaceRequest
	GetRegionId() *string
}

type GetGrafanaWorkspaceRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
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

func (s GetGrafanaWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGrafanaWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetGrafanaWorkspaceRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetGrafanaWorkspaceRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GetGrafanaWorkspaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGrafanaWorkspaceRequest) SetAliyunLang(v string) *GetGrafanaWorkspaceRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetGrafanaWorkspaceRequest) SetGrafanaWorkspaceId(v string) *GetGrafanaWorkspaceRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GetGrafanaWorkspaceRequest) SetRegionId(v string) *GetGrafanaWorkspaceRequest {
	s.RegionId = &v
	return s
}

func (s *GetGrafanaWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
