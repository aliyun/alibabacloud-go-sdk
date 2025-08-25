// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeVhostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMcubeVhostRequest
	GetAppId() *string
	SetTenantId(v string) *QueryMcubeVhostRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMcubeVhostRequest
	GetWorkspaceId() *string
}

type QueryMcubeVhostRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcubeVhostRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeVhostRequest) GoString() string {
	return s.String()
}

func (s *QueryMcubeVhostRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcubeVhostRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMcubeVhostRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcubeVhostRequest) SetAppId(v string) *QueryMcubeVhostRequest {
	s.AppId = &v
	return s
}

func (s *QueryMcubeVhostRequest) SetTenantId(v string) *QueryMcubeVhostRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMcubeVhostRequest) SetWorkspaceId(v string) *QueryMcubeVhostRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcubeVhostRequest) Validate() error {
	return dara.Validate(s)
}
