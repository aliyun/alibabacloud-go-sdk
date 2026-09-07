// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrafanaWorkspaceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *ListGrafanaWorkspaceAccountRequest
	GetAliyunLang() *string
	SetGrafanaWorkspaceId(v string) *ListGrafanaWorkspaceAccountRequest
	GetGrafanaWorkspaceId() *string
	SetRegionId(v string) *ListGrafanaWorkspaceAccountRequest
	GetRegionId() *string
}

type ListGrafanaWorkspaceAccountRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-qzm4mumf401
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGrafanaWorkspaceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGrafanaWorkspaceAccountRequest) GoString() string {
	return s.String()
}

func (s *ListGrafanaWorkspaceAccountRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *ListGrafanaWorkspaceAccountRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *ListGrafanaWorkspaceAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGrafanaWorkspaceAccountRequest) SetAliyunLang(v string) *ListGrafanaWorkspaceAccountRequest {
	s.AliyunLang = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountRequest) SetGrafanaWorkspaceId(v string) *ListGrafanaWorkspaceAccountRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountRequest) SetRegionId(v string) *ListGrafanaWorkspaceAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ListGrafanaWorkspaceAccountRequest) Validate() error {
	return dara.Validate(s)
}
