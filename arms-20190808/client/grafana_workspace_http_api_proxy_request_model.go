// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceHttpApiProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyStr(v string) *GrafanaWorkspaceHttpApiProxyRequest
	GetBodyStr() *string
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceHttpApiProxyRequest
	GetGrafanaWorkspaceId() *string
	SetOrgId(v int64) *GrafanaWorkspaceHttpApiProxyRequest
	GetOrgId() *int64
	SetRegionId(v string) *GrafanaWorkspaceHttpApiProxyRequest
	GetRegionId() *string
}

type GrafanaWorkspaceHttpApiProxyRequest struct {
	// example:
	//
	// example1:
	//
	// {"method":"GET","path":"/api/dashboards/tags"}
	//
	// example2:
	//
	// {
	//
	//   "method": "POST",
	//
	//   "path": "/api/dashboards/db",
	//
	//   "headers": { "Content-Type": "application/json" },
	//
	//   "body": "{\\"dashboard\\":{\\"id\\":null,\\"uid\\":null,\\"title\\":\\"demo\\",\\"schemaVersion\\":16},\\"overwrite\\":false}"
	//
	// }
	//
	// example3:
	//
	// {"method":"GET","path":"/api/datasources","queryParams":{"type":["prometheus","mysql"]}}
	BodyStr *string `json:"BodyStr,omitempty" xml:"BodyStr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// grafana-cn-06f4xyxjo01
	GrafanaWorkspaceId *string `json:"GrafanaWorkspaceId,omitempty" xml:"GrafanaWorkspaceId,omitempty"`
	// example:
	//
	// 1
	OrgId *int64 `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GrafanaWorkspaceHttpApiProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceHttpApiProxyRequest) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) GetBodyStr() *string {
	return s.BodyStr
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) SetBodyStr(v string) *GrafanaWorkspaceHttpApiProxyRequest {
	s.BodyStr = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) SetGrafanaWorkspaceId(v string) *GrafanaWorkspaceHttpApiProxyRequest {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) SetOrgId(v int64) *GrafanaWorkspaceHttpApiProxyRequest {
	s.OrgId = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) SetRegionId(v string) *GrafanaWorkspaceHttpApiProxyRequest {
	s.RegionId = &v
	return s
}

func (s *GrafanaWorkspaceHttpApiProxyRequest) Validate() error {
	return dara.Validate(s)
}
