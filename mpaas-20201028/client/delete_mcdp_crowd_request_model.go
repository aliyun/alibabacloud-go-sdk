// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteMcdpCrowdRequest
	GetAppId() *string
	SetMpaasMappcenterMcdpCrowdDeleteJsonStr(v string) *DeleteMcdpCrowdRequest
	GetMpaasMappcenterMcdpCrowdDeleteJsonStr() *string
	SetTenantId(v string) *DeleteMcdpCrowdRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *DeleteMcdpCrowdRequest
	GetWorkspaceId() *string
}

type DeleteMcdpCrowdRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MpaasMappcenterMcdpCrowdDeleteJsonStr *string `json:"MpaasMappcenterMcdpCrowdDeleteJsonStr,omitempty" xml:"MpaasMappcenterMcdpCrowdDeleteJsonStr,omitempty"`
	TenantId                              *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId                           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteMcdpCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpCrowdRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcdpCrowdRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMcdpCrowdRequest) GetMpaasMappcenterMcdpCrowdDeleteJsonStr() *string {
	return s.MpaasMappcenterMcdpCrowdDeleteJsonStr
}

func (s *DeleteMcdpCrowdRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteMcdpCrowdRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteMcdpCrowdRequest) SetAppId(v string) *DeleteMcdpCrowdRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMcdpCrowdRequest) SetMpaasMappcenterMcdpCrowdDeleteJsonStr(v string) *DeleteMcdpCrowdRequest {
	s.MpaasMappcenterMcdpCrowdDeleteJsonStr = &v
	return s
}

func (s *DeleteMcdpCrowdRequest) SetTenantId(v string) *DeleteMcdpCrowdRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteMcdpCrowdRequest) SetWorkspaceId(v string) *DeleteMcdpCrowdRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteMcdpCrowdRequest) Validate() error {
	return dara.Validate(s)
}
